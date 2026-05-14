# Example README Template

Use this template when creating a `README.md` for any example under `examples/`.

---

## Template

```markdown
# <Example Name>

<One sentence explaining what this example demonstrates.>

## Prerequisites

- Go 1.22+ installed
- <Any other prerequisites, or remove this section if none>

## How to run

    go run ./examples/<example-name>

## Expected output

    <Paste the exact terminal output the user should see.>

## Concepts

- <Concept 1>: brief explanation, link to `pkg/<package>` or spec.
- <Concept 2>: brief explanation, link to `pkg/<package>` or spec.
```

---

## Rules

1. **Title** must match the directory name in human-readable form (e.g., `basic-agent` becomes `Basic Agent`).
2. **Purpose sentence** must be concrete: "Demonstrates X using Y", not "Shows how things work".
3. **How to run** must use the exact `go run` command from the repository root.
4. **Expected output** must reflect what a clean run produces. If output varies (e.g., timestamps), use placeholders like `<timestamp>`.
5. **Concepts** must link to the relevant `pkg/*` package or spec so the reader can dive deeper.
6. Do not include implementation details or internal package references.
7. Keep the README short; the example code is the primary documentation.

## Reference

See `examples/demo-app/README.md` for an existing example in the repository.
