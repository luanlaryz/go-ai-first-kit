package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"sort"
	"strings"
)

// Public route security gate (conservative, AST-based).
//
// Scope and limitations: this models a Gin-style router in
// internal/adapters/http/router.go by tracking New() roots, Group(...) chains
// and Use(...) middleware. It flags any public route not behind approved
// middleware (auth/tenant/rate-limit) unless the route is allowlisted. It does
// NOT fully model dynamic routing (routes built from variables or loops), and it
// does not trace where the tenant identity comes from: that part stays normative
// in .cursor/rules/http-security-tenant.mdc. OpenAPI security is checked
// textually.
//
// Both inputs are optional: a project without an HTTP adapter or without a
// published contract passes without configuration.

var httpMethods = map[string]bool{
	"GET": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true, "HEAD": true, "OPTIONS": true,
}

// publicRoutePrefix is the route namespace treated as publicly reachable.
var publicRoutePrefix = envOrDefault("PUBLIC_ROUTE_PREFIX", "/v1")

func envOrDefault(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

type routerGroup struct {
	parent string
	prefix string
	mws    []string
}

func approvedMiddleware(name string) bool {
	lower := strings.ToLower(name)
	return strings.Contains(lower, "auth") ||
		strings.Contains(lower, "tenant") ||
		strings.Contains(lower, "ratelimit") ||
		strings.Contains(lower, "rate_limit")
}

func exprCalleeName(e ast.Expr) string {
	if call, ok := e.(*ast.CallExpr); ok {
		return calleeName(call)
	}
	return ""
}

func stringLit(e ast.Expr) (string, bool) {
	lit, ok := e.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", false
	}
	return strings.Trim(lit.Value, "\"`"), true
}

func selRecvAndName(e ast.Expr) (string, string, bool) {
	sel, ok := e.(*ast.SelectorExpr)
	if !ok {
		return "", "", false
	}
	recv, ok := sel.X.(*ast.Ident)
	if !ok {
		return "", "", false
	}
	return recv.Name, sel.Sel.Name, true
}

type routeReg struct {
	recv   string
	method string
	path   string
}

func runPublicRoute() error {
	allow, err := loadAllowlist(".guardrails/public-route-exceptions.yaml", false)
	if err != nil {
		return err
	}
	allowed := allowlistSet(allow)

	var violations []string

	routerFile := "internal/adapters/http/router.go"
	if _, statErr := os.Stat(routerFile); statErr == nil {
		groups := map[string]*routerGroup{}
		var routes []routeReg

		fset := token.NewFileSet()
		file, perr := parseGoFile(fset, routerFile)
		if perr != nil {
			return perr
		}
		ast.Inspect(file, func(n ast.Node) bool {
			// group definitions: x := gin.New()  /  x := y.Group("p", mw...)
			if as, ok := n.(*ast.AssignStmt); ok && len(as.Lhs) == 1 && len(as.Rhs) == 1 {
				lhs, lok := as.Lhs[0].(*ast.Ident)
				call, cok := as.Rhs[0].(*ast.CallExpr)
				if lok && cok {
					if name := exprCalleeName(as.Rhs[0]); name == "New" || name == "Default" {
						groups[lhs.Name] = &routerGroup{parent: "", prefix: ""}
					} else if recv, sel, ok := selRecvAndName(call.Fun); ok && sel == "Group" {
						g := &routerGroup{parent: recv}
						if len(call.Args) > 0 {
							if p, ok := stringLit(call.Args[0]); ok {
								g.prefix = p
							}
						}
						for _, a := range call.Args[1:] {
							g.mws = append(g.mws, exprCalleeName(a))
						}
						groups[lhs.Name] = g
					}
				}
			}
			// calls: x.Use(...) / x.METHOD("path", ...)
			if call, ok := n.(*ast.CallExpr); ok {
				if recv, sel, ok := selRecvAndName(call.Fun); ok {
					if sel == "Use" {
						if g := groups[recv]; g != nil {
							for _, a := range call.Args {
								g.mws = append(g.mws, exprCalleeName(a))
							}
						}
					} else if httpMethods[strings.ToUpper(sel)] {
						if len(call.Args) > 0 {
							if p, ok := stringLit(call.Args[0]); ok {
								routes = append(routes, routeReg{recv: recv, method: strings.ToUpper(sel), path: p})
							}
						}
					}
				}
			}
			return true
		})

		resolvePrefix := func(recv string) string {
			parts := []string{}
			seen := map[string]bool{}
			for recv != "" && groups[recv] != nil && !seen[recv] {
				seen[recv] = true
				g := groups[recv]
				if g.prefix != "" {
					parts = append([]string{g.prefix}, parts...)
				}
				recv = g.parent
			}
			return strings.Join(parts, "")
		}
		protectedChain := func(recv string) bool {
			seen := map[string]bool{}
			for recv != "" && groups[recv] != nil && !seen[recv] {
				seen[recv] = true
				g := groups[recv]
				for _, mw := range g.mws {
					if approvedMiddleware(mw) {
						return true
					}
				}
				recv = g.parent
			}
			return false
		}

		for _, r := range routes {
			full := resolvePrefix(r.recv) + r.path
			if !strings.HasPrefix(full, publicRoutePrefix) {
				continue
			}
			if protectedChain(r.recv) {
				continue
			}
			key := r.method + " " + full
			if allowed[allowKey(key, "")] {
				continue
			}
			violations = append(violations, fmt.Sprintf("public route without auth/tenant/rate-limit: %q (allowlist it in .guardrails/public-route-exceptions.yaml only with an expiring debt ref, or put it behind approved middleware)", key))
		}
	}

	// Published contract: must declare securitySchemes (else allowlisted debt).
	openapi := "api/openapi.yaml"
	if data, rerr := os.ReadFile(openapi); rerr == nil {
		if !strings.Contains(string(data), "securitySchemes:") {
			if !allowed[allowKey(openapi, "")] {
				violations = append(violations, fmt.Sprintf("%s declares no securitySchemes for the public API (define schemes + per-operation security, or allowlist as expiring debt)", openapi))
			}
		}
	}

	if len(violations) > 0 {
		sort.Strings(violations)
		return fmt.Errorf("public-route security violations:\n  - %s", strings.Join(violations, "\n  - "))
	}
	fmt.Printf("public-route: %s routes covered by auth/tenant/rate-limit or audited allowlist; contract security present or allowlisted\n", publicRoutePrefix)
	return nil
}
