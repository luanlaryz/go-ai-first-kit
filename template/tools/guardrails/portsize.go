package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

const (
	portWarnMethods  = 3
	portErrorMethods = 5
)

// interfaceMethodCount counts explicit methods (named fields with a function
// type); embedded interfaces are not counted as methods.
func interfaceMethodCount(it *ast.InterfaceType) int {
	if it.Methods == nil {
		return 0
	}
	count := 0
	for _, field := range it.Methods.List {
		if len(field.Names) == 0 {
			continue // embedded interface
		}
		if _, ok := field.Type.(*ast.FuncType); ok {
			count += len(field.Names)
		}
	}
	return count
}

func runPortSize() error {
	allow, err := loadAllowlist(".guardrails/port-size-exceptions.yaml", true)
	if err != nil {
		return err
	}
	allowed := allowlistSet(allow)

	files, err := listProductionGoFiles("pkg", "internal/domain", "internal/application")
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
		ast.Inspect(file, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			it, ok := ts.Type.(*ast.InterfaceType)
			if !ok {
				return true
			}
			name := ts.Name.Name
			count := interfaceMethodCount(it)
			switch {
			case count > portErrorMethods:
				if allowed[allowKey(f, name)] {
					return true
				}
				errors = append(errors, fmt.Sprintf("%s:%s has %d methods (limit %d)", f, name, count, portErrorMethods))
			case count > portWarnMethods:
				warnings = append(warnings, fmt.Sprintf("%s:%s has %d methods (warn %d)", f, name, count, portWarnMethods))
			}
			return true
		})
	}
	if len(warnings) > 0 {
		fmt.Printf("port-size: %d interfaces in warn band (>%d methods)\n", len(warnings), portWarnMethods)
	}
	if len(errors) > 0 {
		return fmt.Errorf("port-size violations (limit %d methods):\n  - %s\n  prefer 1-3 methods; add a justified, expiring entry to .guardrails/port-size-exceptions.yaml or split the port", portErrorMethods, strings.Join(errors, "\n  - "))
	}
	fmt.Printf("port-size: scanned %d files; no errors\n", len(files))
	return nil
}
