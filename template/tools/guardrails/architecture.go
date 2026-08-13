package main

import (
	"fmt"
	"go/token"
	"os"
	"strings"
)

// applicationForbidden lists import substrings application code must not use
// (concrete infra, frameworks, vendor SDKs and concrete telemetry). Telemetry
// reaches the application layer through a port wired in the composition root.
var applicationForbidden = []string{
	"/internal/adapters",
	"/internal/adapter",
	"github.com/gin-gonic/gin",
	"github.com/go-chi/chi",
	"github.com/labstack/echo",
	"github.com/jackc/pgx",
	"github.com/redis/go-redis",
	"github.com/aws/aws-sdk",
	"github.com/golang-jwt",
	"github.com/prometheus/client_golang",
	"go.uber.org/zap",
	"go.opentelemetry.io",
	"net/http",
}

// externalSDKs may only be imported from adapters or the composition root.
var externalSDKs = []string{
	"github.com/gin-gonic/gin",
	"github.com/go-chi/chi",
	"github.com/labstack/echo",
	"github.com/jackc/pgx",
	"github.com/redis/go-redis",
	"github.com/aws/aws-sdk",
	"github.com/golang-jwt",
}

func modulePath() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return ""
}

func isStdlibImport(path string) bool {
	first := path
	if i := strings.Index(path, "/"); i >= 0 {
		first = path[:i]
	}
	return !strings.Contains(first, ".")
}

func importsOf(fset *token.FileSet, path string) ([]string, error) {
	file, err := parseGoFile(fset, path)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, imp := range file.Imports {
		out = append(out, strings.Trim(imp.Path.Value, "\""))
	}
	return out, nil
}

func compositionRootOrAdapter(path string) bool {
	return strings.Contains(path, "internal/adapters/") ||
		strings.Contains(path, "internal/adapter/") ||
		strings.Contains(path, "internal/bootstrap/") ||
		strings.HasPrefix(path, "cmd/")
}

// publicBridge is the single package allowed to reach into internal/ from the
// public surface. Everything else in pkg/ must stay independent of internal/.
const publicBridge = "pkg/app/"

func runArchitecture() error {
	mod := modulePath()
	fset := token.NewFileSet()
	var violations []string

	// Domain: stdlib + own domain package only.
	domainFiles, err := listProductionGoFiles("internal/domain")
	if err != nil {
		return err
	}
	if mod != "" {
		domainPkg := mod + "/internal/domain"
		for _, f := range domainFiles {
			imps, err := importsOf(fset, f)
			if err != nil {
				return err
			}
			for _, imp := range imps {
				if isStdlibImport(imp) || imp == domainPkg || strings.HasPrefix(imp, domainPkg+"/") {
					continue
				}
				violations = append(violations, fmt.Sprintf("%s: domain must import stdlib only, found %q", f, imp))
			}
		}
	}

	// Application: denylist of concrete infra/frameworks/SDKs/telemetry.
	appFiles, err := listProductionGoFiles("internal/application")
	if err != nil {
		return err
	}
	for _, f := range appFiles {
		imps, err := importsOf(fset, f)
		if err != nil {
			return err
		}
		for _, imp := range imps {
			for _, bad := range applicationForbidden {
				if strings.Contains(imp, bad) {
					violations = append(violations, fmt.Sprintf("%s: application must not import %q (forbidden: %s)", f, imp, bad))
				}
			}
		}
	}

	// External SDKs only in adapters or composition root (cmd, bootstrap).
	internalFiles, err := listProductionGoFiles("internal", "cmd")
	if err != nil {
		return err
	}
	for _, f := range internalFiles {
		if compositionRootOrAdapter(f) {
			continue
		}
		imps, err := importsOf(fset, f)
		if err != nil {
			return err
		}
		for _, imp := range imps {
			for _, sdk := range externalSDKs {
				if strings.Contains(imp, sdk) {
					violations = append(violations, fmt.Sprintf("%s: external SDK %q only allowed in adapters or composition root (internal/bootstrap, cmd)", f, imp))
				}
			}
		}
	}

	// Public surface: pkg/ must not depend on internal/, except the declared
	// bridge. This is what keeps the exported API free of unexported types.
	pkgFiles, err := listProductionGoFiles("pkg")
	if err != nil {
		return err
	}
	if mod != "" {
		internalPrefix := mod + "/internal"
		for _, f := range pkgFiles {
			if strings.HasPrefix(f, publicBridge) {
				continue
			}
			imps, err := importsOf(fset, f)
			if err != nil {
				return err
			}
			for _, imp := range imps {
				if imp == internalPrefix || strings.HasPrefix(imp, internalPrefix+"/") {
					violations = append(violations, fmt.Sprintf("%s: pkg/ must not import %q; %s is the only allowed public bridge", f, imp, publicBridge))
				}
			}
		}
	}

	if len(violations) > 0 {
		return fmt.Errorf("architecture boundary violations:\n  - %s", strings.Join(violations, "\n  - "))
	}
	fmt.Printf("architecture: scanned %d domain + %d application + %d public files; boundaries intact\n",
		len(domainFiles), len(appFiles), len(pkgFiles))
	return nil
}
