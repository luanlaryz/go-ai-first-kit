# README Checklist

Use this checklist when creating or reviewing the root `README.md` or any sub-module README.

## Root README.md

### Required sections

- [ ] **Project name and description**: one paragraph explaining what {{PROJECT_SLUG}} is and its relationship to {{UPSTREAM_NAME}}.
- [ ] **Installation**: `go get` command with the module path.
- [ ] **Quick start**: minimal compilable code snippet showing the most common use case (e.g., creating and running an agent).
- [ ] **Documentation link**: link to pkg.go.dev for full API reference.
- [ ] **Contributing**: link to `CONTRIBUTING.md`.
- [ ] **License**: license type and link to `LICENSE` file.

### Recommended sections

- [ ] **Badges**: CI status, Go version, license badge, Go Report Card.
- [ ] **Roadmap**: link to specs or high-level milestones.
- [ ] **Non-scope**: brief mention of what {{PROJECT_SLUG}} does NOT cover ({{UPSTREAM_OPS_NAME}}) with link to `specs/001-non-goals.md`.
- [ ] **Examples**: list of examples under `examples/` with one-line descriptions.

### Rules

1. Quick-start code must compile without importing `internal/*`.
2. Do not duplicate godoc content; link to pkg.go.dev instead.
3. Keep the README concise; detailed API documentation belongs in `doc.go` files.
4. Update the README in the same PR when public API changes affect user-facing instructions.
5. Badge URLs must point to stable endpoints (GitHub Actions, pkg.go.dev, goreportcard).

## Example README (examples/*/README.md)

See [example_readme_template.md](example_readme_template.md) for the full template.

### Minimum content

- [ ] **Title**: name of the example.
- [ ] **Purpose**: one sentence explaining what the example demonstrates.
- [ ] **How to run**: exact command (`go run ./examples/<name>`).
- [ ] **Expected output**: what the user should see in the terminal.
- [ ] **Concepts**: list of {{PROJECT_SLUG}} concepts demonstrated, with links to relevant `pkg/*` godoc or specs.
