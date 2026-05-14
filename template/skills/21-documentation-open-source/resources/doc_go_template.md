# doc.go Template and Conventions

## Purpose

Every package under `pkg/*` must have a `doc.go` file that provides godoc-compatible documentation. This file contains only a package-level comment and the `package` clause.

## Structure

```go
// Package <name> <one-sentence summary starting with a verb>.
//
// <Expanded description: what the package provides, its role in {{PROJECT_SLUG}},
// and key abstractions.>
//
// # <Section heading>
//
// <Details about a specific topic: choosing implementations, configuring
// behavior, extending via interfaces, etc.>
//
//   - [TypeOrFunc] short description of the first option.
//
//   - [TypeOrFunc] short description of the second option.
//
// # <Another section heading>
//
// <Guidance for users who want to extend or customize.>
package <name>
```

## Rules

1. **Opening sentence**: must start with `Package <name>` followed by a verb phrase that summarizes the package purpose. This sentence appears in pkg.go.dev package listings.
2. **No blank line** between the comment block and the `package` clause.
3. **Section headings** use `# Heading` (single `#`) inside the comment to create godoc sections (Go 1.19+).
4. **Symbol references** use `[SymbolName]` to create hyperlinks in pkg.go.dev.
5. **Lists** use `//   - ` (three spaces + dash) for indented bullet points rendered by godoc.
6. **Code examples** inside doc comments use indentation (tab or four spaces after `//`).
7. **Keep it focused**: `doc.go` documents the package contract, not implementation details from `internal/*`.
8. **Traceability**: mention the governing spec when the package implements a specific spec requirement (e.g., "as defined in spec 030").

## Reference Model

The canonical example in this repository is `pkg/memory/doc.go`, which demonstrates:

- Opening sentence naming the package
- Contextual description separating conversational memory from working memory
- `# Choosing a backend` section with two listed implementations
- `# Implementing a custom adapter` section for extensibility guidance
- Symbol references via `[Store]`, `[InMemoryStore]`, `[FileStore]`, `[NewFileStore]`

## When NOT to Use doc.go

- Packages with only one or two exported symbols may document directly above those declarations instead of in a separate `doc.go`.
- `internal/*` packages do not require `doc.go` (but inline comments on exported-within-internal symbols are still encouraged).
