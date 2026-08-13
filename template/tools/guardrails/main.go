// Command guardrails implements deterministic, production-only structural
// guardrails (architecture boundaries, function/port size, ignored errors,
// public route security). Allowlists live in .guardrails/*.yaml and are
// validated (required fields, ISO expiry, ref) by a stdlib-only parser, so the
// tool builds without adding any dependency to the module.
//
// Every check skips roots that do not exist, so a freshly rendered project
// passes and the gates activate as the codebase grows.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: guardrails <architecture|funcsize|portsize|ignored-errors|public-route|allowlist-paths> [args]")
		os.Exit(2)
	}
	cmd := os.Args[1]
	args := os.Args[2:]
	var err error
	switch cmd {
	case "architecture":
		err = runArchitecture()
	case "funcsize":
		err = runFuncSize()
	case "portsize":
		err = runPortSize()
	case "ignored-errors":
		err = runIgnoredErrors()
	case "public-route":
		err = runPublicRoute()
	case "allowlist-paths":
		err = runAllowlistPaths(args)
	default:
		fmt.Fprintf(os.Stderr, "guardrails: unknown subcommand %q\n", cmd)
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "guardrails %s: FAIL\n%v\n", cmd, err)
		os.Exit(1)
	}
	// allowlist-paths emits data on stdout; keep it clean for shell consumers.
	if cmd != "allowlist-paths" {
		fmt.Printf("guardrails %s: ok\n", cmd)
	}
}

// excludedDir reports whether a directory must be skipped for production scans.
func excludedDir(name string) bool {
	switch name {
	case "vendor", "third_party", "tools", "testdata", ".git", "node_modules":
		return true
	}
	return false
}

// listProductionGoFiles walks the given roots and returns production *.go files,
// excluding tests, vendored/generated trees and generated files.
func listProductionGoFiles(roots ...string) ([]string, error) {
	var files []string
	for _, root := range roots {
		if _, err := os.Stat(root); err != nil {
			continue
		}
		walkErr := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				if excludedDir(d.Name()) {
					return filepath.SkipDir
				}
				return nil
			}
			if !strings.HasSuffix(path, ".go") {
				return nil
			}
			if strings.HasSuffix(path, "_test.go") {
				return nil
			}
			gen, gerr := isGenerated(path)
			if gerr != nil {
				return gerr
			}
			if gen {
				return nil
			}
			files = append(files, path)
			return nil
		})
		if walkErr != nil {
			return nil, walkErr
		}
	}
	return files, nil
}

// isGenerated reports whether a Go file carries the standard generated marker.
func isGenerated(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	const marker = "// Code generated"
	head := data
	if len(head) > 2048 {
		head = head[:2048]
	}
	for _, line := range strings.Split(string(head), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, marker) && strings.Contains(line, "DO NOT EDIT") {
			return true, nil
		}
	}
	return false, nil
}

func parseGoFile(fset *token.FileSet, path string) (*ast.File, error) {
	return parser.ParseFile(fset, path, nil, parser.ParseComments)
}

// receiverName returns the bare receiver type name for a method declaration.
func receiverName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return ""
	}
	switch t := fn.Recv.List[0].Type.(type) {
	case *ast.StarExpr:
		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name
		}
	case *ast.Ident:
		return t.Name
	}
	return ""
}

// funcSymbol returns "Recv.Name" for methods and "Name" for free functions.
func funcSymbol(fn *ast.FuncDecl) string {
	if recv := receiverName(fn); recv != "" {
		return recv + "." + fn.Name.Name
	}
	return fn.Name.Name
}
