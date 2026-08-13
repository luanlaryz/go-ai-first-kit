package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

// Calls whose error genuinely matters: audit trails, persistence, queue writes
// and IO. Discarding those errors turns a failure into silence.
var ignoredErrorKeywords = []string{
	"audit", "capture", "record", "publish", "persist", "save", "send", "enqueue", "write",
}

func ignoredErrorKeyword(name string) string {
	lower := strings.ToLower(name)
	for _, kw := range ignoredErrorKeywords {
		if strings.Contains(lower, kw) {
			return kw
		}
	}
	return ""
}

func calleeName(call *ast.CallExpr) string {
	switch fn := call.Fun.(type) {
	case *ast.Ident:
		return fn.Name
	case *ast.SelectorExpr:
		return fn.Sel.Name
	}
	return ""
}

func runIgnoredErrors() error {
	allow, err := loadAllowlist(".guardrails/ignored-error-exceptions.yaml", true)
	if err != nil {
		return err
	}
	allowed := allowlistSet(allow)

	files, err := listProductionGoFiles("pkg", "internal", "cmd")
	if err != nil {
		return err
	}
	sort.Strings(files)

	fset := token.NewFileSet()
	var violations []string
	for _, f := range files {
		file, err := parseGoFile(fset, f)
		if err != nil {
			return err
		}
		ast.Inspect(file, func(n ast.Node) bool {
			assign, ok := n.(*ast.AssignStmt)
			if !ok || len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
				return true
			}
			lhs, ok := assign.Lhs[0].(*ast.Ident)
			if !ok || lhs.Name != "_" {
				return true
			}
			call, ok := assign.Rhs[0].(*ast.CallExpr)
			if !ok {
				return true
			}
			name := calleeName(call)
			kw := ignoredErrorKeyword(name)
			if kw == "" {
				return true
			}
			if allowed[allowKey(f, name)] {
				return true
			}
			line := fset.Position(assign.Pos()).Line
			violations = append(violations, fmt.Sprintf("%s:%d: ignored error from %q (matches %q); audit/persistence/queue/IO errors must be handled", f, line, name, kw))
			return true
		})
	}
	if len(violations) > 0 {
		return fmt.Errorf("ignored-error violations:\n  - %s\n  handle the error, or add a justified, expiring entry to .guardrails/ignored-error-exceptions.yaml with an inline `// best-effort:` comment", strings.Join(violations, "\n  - "))
	}
	fmt.Printf("ignored-errors: scanned %d files; none unhandled outside allowlist\n", len(files))
	return nil
}
