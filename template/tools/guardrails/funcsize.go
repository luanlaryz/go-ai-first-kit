package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

const (
	funcWarnLines  = 40
	funcErrorLines = 80
)

func runFuncSize() error {
	allow, err := loadAllowlist(".guardrails/function-size-exceptions.yaml", true)
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
	var errors, warnings []string
	for _, f := range files {
		file, err := parseGoFile(fset, f)
		if err != nil {
			return err
		}
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Body == nil {
				continue
			}
			start := fset.Position(fn.Pos()).Line
			end := fset.Position(fn.End()).Line
			length := end - start + 1
			sym := funcSymbol(fn)
			switch {
			case length > funcErrorLines:
				if allowed[allowKey(f, sym)] {
					continue
				}
				errors = append(errors, fmt.Sprintf("%s:%s has %d lines (limit %d)", f, sym, length, funcErrorLines))
			case length > funcWarnLines:
				warnings = append(warnings, fmt.Sprintf("%s:%s has %d lines (warn %d)", f, sym, length, funcWarnLines))
			}
		}
	}
	if len(warnings) > 0 {
		fmt.Printf("function-size: %d functions in warn band (>%d)\n", len(warnings), funcWarnLines)
	}
	if len(errors) > 0 {
		return fmt.Errorf("function-size violations (limit %d lines):\n  - %s\n  add a justified, expiring entry to .guardrails/function-size-exceptions.yaml or refactor", funcErrorLines, strings.Join(errors, "\n  - "))
	}
	fmt.Printf("function-size: scanned %d files; no errors\n", len(files))
	return nil
}
