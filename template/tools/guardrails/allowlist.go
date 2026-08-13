package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Exception is one allowlist entry shared by all guardrail allowlists.
type Exception struct {
	Path          string
	Symbol        string
	Rule          string
	Justification string
	Owner         string
	ExpiresAt     string
	Ref           string
	line          int
}

// loadAllowlist parses and validates a .guardrails/*.yaml allowlist using a
// stdlib-only line reader (the format is intentionally flat). A missing or
// malformed file is a hard error; it is never treated as "no exceptions",
// because deleting the allowlist must not be a way to pass the gate.
// requireSymbol enforces the symbol field for symbol-scoped allowlists.
func loadAllowlist(path string, requireSymbol bool) ([]Exception, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("allowlist %s: cannot open (%v); create it with `exceptions: []` if there are none", path, err)
	}
	defer f.Close()

	var (
		entries   []Exception
		cur       *Exception
		inList    bool
		sawHeader bool
	)
	flush := func() error {
		if cur == nil {
			return nil
		}
		if err := validateException(path, *cur, requireSymbol); err != nil {
			return err
		}
		entries = append(entries, *cur)
		cur = nil
		return nil
	}

	scanner := bufio.NewScanner(f)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		raw := scanner.Text()
		trimmed := strings.TrimSpace(raw)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if strings.HasPrefix(trimmed, "exceptions:") {
			inList = true
			sawHeader = true
			// support inline "exceptions: []"
			if strings.Contains(trimmed, "[]") {
				inList = false
			}
			continue
		}
		if !inList {
			// top-level keys (e.g. version:) are ignored
			continue
		}
		if strings.HasPrefix(trimmed, "- ") {
			if err := flush(); err != nil {
				return nil, err
			}
			cur = &Exception{line: lineNo}
			trimmed = strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
			if trimmed == "" {
				continue
			}
		}
		if cur == nil {
			return nil, fmt.Errorf("allowlist %s:%d: field outside of a list item: %q", path, lineNo, raw)
		}
		key, value, ok := splitKV(trimmed)
		if !ok {
			return nil, fmt.Errorf("allowlist %s:%d: malformed line: %q", path, lineNo, raw)
		}
		switch key {
		case "path":
			cur.Path = value
		case "symbol":
			cur.Symbol = value
		case "rule":
			cur.Rule = value
		case "justification":
			cur.Justification = value
		case "owner":
			cur.Owner = value
		case "expires_at":
			cur.ExpiresAt = value
		case "ref":
			cur.Ref = value
		default:
			return nil, fmt.Errorf("allowlist %s:%d: unknown field %q", path, lineNo, key)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("allowlist %s: read error: %v", path, err)
	}
	if err := flush(); err != nil {
		return nil, err
	}
	if !sawHeader {
		return nil, fmt.Errorf("allowlist %s: missing 'exceptions:' section", path)
	}
	return entries, nil
}

func splitKV(s string) (string, string, bool) {
	idx := strings.Index(s, ":")
	if idx < 0 {
		return "", "", false
	}
	key := strings.TrimSpace(s[:idx])
	value := strings.TrimSpace(s[idx+1:])
	value = strings.Trim(value, "\"'")
	return key, value, key != ""
}

func validateException(path string, e Exception, requireSymbol bool) error {
	missing := []string{}
	if e.Path == "" {
		missing = append(missing, "path")
	}
	if e.Rule == "" {
		missing = append(missing, "rule")
	}
	if e.Justification == "" {
		missing = append(missing, "justification")
	}
	if e.Owner == "" {
		missing = append(missing, "owner")
	}
	if e.ExpiresAt == "" {
		missing = append(missing, "expires_at")
	}
	if e.Ref == "" {
		missing = append(missing, "ref")
	}
	if requireSymbol && e.Symbol == "" {
		missing = append(missing, "symbol")
	}
	if len(missing) > 0 {
		return fmt.Errorf("allowlist %s:%d: entry %q missing required field(s): %s", path, e.line, e.Path, strings.Join(missing, ", "))
	}
	exp, err := time.Parse("2006-01-02", e.ExpiresAt)
	if err != nil {
		return fmt.Errorf("allowlist %s:%d: entry %q has invalid expires_at %q (want YYYY-MM-DD)", path, e.line, e.Path, e.ExpiresAt)
	}
	today := time.Now().UTC().Truncate(24 * time.Hour)
	if exp.Before(today) {
		return fmt.Errorf("allowlist %s:%d: entry %q expired at %s (refactor or renew with justification)", path, e.line, e.Path, e.ExpiresAt)
	}
	return nil
}

// allowKey builds the lookup key used to match findings to allowlist entries.
func allowKey(path, symbol string) string {
	if symbol == "" {
		return path
	}
	return path + "::" + symbol
}

func allowlistSet(entries []Exception) map[string]bool {
	set := make(map[string]bool, len(entries))
	for _, e := range entries {
		set[allowKey(e.Path, e.Symbol)] = true
	}
	return set
}

// runAllowlistPaths validates an allowlist and prints its path values, one per
// line. Used by bash gates (e.g. package-clean) so allowlist validation stays
// centralized in Go.
func runAllowlistPaths(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: guardrails allowlist-paths <file>")
	}
	entries, err := loadAllowlist(args[0], false)
	if err != nil {
		return err
	}
	for _, e := range entries {
		fmt.Println(e.Path)
	}
	return nil
}
