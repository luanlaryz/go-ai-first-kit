package gakit

import "embed"

// TemplateFS contains the AI-first project template used by the gakit CLI.
//
//go:embed all:template
var TemplateFS embed.FS

const TemplateRoot = "template"
