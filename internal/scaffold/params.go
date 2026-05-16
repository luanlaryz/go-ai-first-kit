package scaffold

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	slugPattern   = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)
	modulePattern = regexp.MustCompile(`^[^/\s]+/[^/\s]+/[^/\s]+`)
)

type Params struct {
	ProjectSlug        string
	ProjectTitle       string
	ModulePath         string
	ProjectDescription string
	AuthorName         string
	LicenseName        string
	UpstreamName       string
}

func (p *Params) ApplyDefaults() {
	if strings.TrimSpace(p.LicenseName) == "" {
		p.LicenseName = "MIT"
	}
	if strings.TrimSpace(p.UpstreamName) == "" {
		p.UpstreamName = "none"
	}
}

func (p Params) Validate() error {
	var missing []string
	if strings.TrimSpace(p.ProjectSlug) == "" {
		missing = append(missing, "slug")
	}
	if strings.TrimSpace(p.ProjectTitle) == "" {
		missing = append(missing, "title")
	}
	if strings.TrimSpace(p.ModulePath) == "" {
		missing = append(missing, "module")
	}
	if strings.TrimSpace(p.ProjectDescription) == "" {
		missing = append(missing, "description")
	}
	if strings.TrimSpace(p.AuthorName) == "" {
		missing = append(missing, "author")
	}
	if len(missing) > 0 {
		return fmt.Errorf("parametros obrigatorios ausentes: %s", strings.Join(missing, ", "))
	}
	if !slugPattern.MatchString(p.ProjectSlug) {
		return errors.New("slug deve ser Go-safe: comece com letra minuscula e use apenas letras minusculas, numeros ou underscore")
	}
	if !modulePattern.MatchString(p.ModulePath) {
		return errors.New("module deve parecer um modulo Go completo, por exemplo github.com/acme/myapp")
	}
	return nil
}
