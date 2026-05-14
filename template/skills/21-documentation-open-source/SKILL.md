---
name: documentation-open-source
description: Godoc for pkg/* packages, README maintenance, CONTRIBUTING guide, CHANGELOG updates, example READMEs, and spec-to-doc traceability. Use when creating or changing public packages, examples, or any user-facing documentation artifact.
---

# Open Source Documentation
Goal: ensure every public surface of {{PROJECT_SLUG}} has traceable, idiomatic, and up-to-date documentation.

## When to Use
- Creating or modifying a package under `pkg/*` (godoc required).
- Changing public API signatures, types, or observable behavior (README, examples, and spec must follow).
- Adding a new example under `examples/` (example README and smoke test required).
- Preparing a release or recording a behavior change (CHANGELOG entry required).
- Onboarding contributors or accepting first external PR (CONTRIBUTING.md must be current).

## Non-negotiables
1. Every package in `pkg/*` must have a `doc.go` with valid godoc following Go conventions.
2. Godoc opening sentence must start with the declared name (`Package X ...`, `Type Y ...`, `Func Z ...`).
3. Root `README.md` must contain: project description, installation, quick-start snippet, link to pkg.go.dev, and license.
4. Every example under `examples/` must have its own `README.md` with purpose, how to run, and expected output.
5. Public API changes require documentation updates in the same PR as the code change.
6. `CONTRIBUTING.md` must exist and cover: local setup, commit conventions, PR workflow, and how to run tests.
7. `CHANGELOG.md` must be updated for release-worthy changes (integrated with git-cliff per skill 15).
8. Every documentation artifact must be traceable to at least one spec.

## Do / Don't
- **Do** update docs in the same PR as the code change.
- **Do** use a dedicated `doc.go` file for packages with long documentation.
- **Do** validate markdown links and formatting before merge.
- **Do** keep examples compilable and runnable (`go run` or `go test`).
- **Do** reference `pkg.go.dev` from README instead of duplicating API details.
- **Don't** duplicate content between README and godoc; README links, godoc details the API.
- **Don't** document `internal/*` types or behavior in public-facing docs.
- **Don't** create documentation that is not traceable to a spec.
- **Don't** leave an example without a README or without execution instructions.

## Interfaces / Contracts
- `doc.go` template and conventions: [doc_go_template.md](resources/doc_go_template.md).
- Root and sub-module README checklist: [readme_checklist.md](resources/readme_checklist.md).
- Example README template: [example_readme_template.md](resources/example_readme_template.md).
- Expected file locations: `pkg/*/doc.go`, `README.md`, `CONTRIBUTING.md`, `CHANGELOG.md`, `examples/*/README.md`.

## Checklists
**Before**
- [ ] Identify which documentation artifacts are affected by the change.
- [ ] Check if the affected package already has a `doc.go`.
- [ ] Confirm which spec governs the change for traceability.

**During**
- [ ] Write or update `doc.go` for every touched `pkg/*` package.
- [ ] Update root `README.md` if public API surface changed.
- [ ] Write or update example `README.md` for affected examples.
- [ ] Add `CHANGELOG.md` entry if the change is release-worthy.

**After**
- [ ] Validate markdown links (no broken references).
- [ ] Confirm all affected examples compile and run.
- [ ] Verify spec traceability is explicit in PR description.
- [ ] Run `go doc` locally to confirm godoc renders correctly.

## Definition of Done
- All touched `pkg/*` packages have valid `doc.go` with proper opening sentences.
- Root `README.md` reflects current public API if it changed.
- Affected examples have `README.md` and compile successfully.
- `CHANGELOG.md` updated for release-worthy changes.
- `CONTRIBUTING.md` exists and is current.
- No public documentation references `internal/*` types directly.

## Minimal Examples
- Well-structured `doc.go` (reference model: `pkg/memory/doc.go`):
  ```go
  // Package tool defines the contract for executable tools and the
  // helpers used to compose toolkits in {{PROJECT_SLUG}}.
  //
  // # Defining a tool
  //
  // Implement the [Tool] interface and register it with [Registry.Register].
  //
  // # Composing toolkits
  //
  // Use [NewToolkit] to group related tools under a shared namespace.
  package tool
  ```
- Example README (`examples/basic-agent/README.md`):
  ```markdown
  # Basic Agent

  Demonstrates creating and running an agent with a single tool.

  ## How to run

      go run ./examples/basic-agent

  ## Expected output

      Agent replied: Hello from {{PROJECT_SLUG}}!

  ## Concepts

  - Agent creation via `pkg/agent`
  - Tool registration via `pkg/tool`
  ```
- CHANGELOG entry:
  ```markdown
  ## [Unreleased]
  ### Added
  - `pkg/memory`: FileStore with atomic writes and restart persistence.
  ```
