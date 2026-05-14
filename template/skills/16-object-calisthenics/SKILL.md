---
name: object-calisthenics-go-core
description: Apply Object Calisthenics adapted to Go in pkg/* and internal/* to reduce complexity, improve readability, and preserve hexagonal boundaries.
---

# Object Calisthenics for Go Core
Goal: reduce accidental complexity in public contracts and runtime internals while preserving architectural boundaries.

## When to Use
- Changing or reviewing `pkg/*` and `internal/*`.
- Refactoring long functions, nested conditionals, or bulky structs.
- Creating input/output types that carry structured identifiers.

## Non-negotiables
1. Max 1 indentation level per function; use guard clauses and early returns.
2. Avoid `else`; return early and keep the happy path flat.
3. Keep functions small (target <= 20 lines of logic; refactor when > 30).
4. No primitive obsession for identifiers; use strong types when clarity benefits.
5. Prefer max 3 function arguments; use input structs for larger inputs.
6. Avoid god structs; split by responsibility and compose explicitly.
7. Encapsulate collections with domain rules (example: `Messages`).
8. `pkg/*` must not import `internal/*` except through the authorized bridge (`pkg/app`).

## Do / Don't
- **Do** extract helpers named by intent (`validateTenantScope`, `buildCommandEnvelope`).
- **Do** centralize invariant checks near domain types.
- **Do** keep side effects at the edges (ports/adapters).
- **Don't** pass raw `string` for tenant identifiers across layers.
- **Don't** mix validation, orchestration, and infrastructure in one method.
- **Don't** add conditional trees when polymorphism or map-based dispatch is clearer.

## Interfaces / Contracts
- Canonical typed identifiers:
  ```go
  type {{DOMAIN_ACTOR_TITLE}}ID string
  type ConversationID string
  type RequestID string
  type Seq int64
  ```
- Input object for use cases:
  ```go
  type SendMessageInput struct {
      {{DOMAIN_ACTOR_TITLE}}ID       {{DOMAIN_ACTOR_TITLE}}ID
      ConversationID ConversationID
      RequestID      RequestID
      Body           string
  }
  ```
- Collection wrapper with invariant:
  ```go
  type Messages struct {
      items []Message
  }
  ```

## Checklists
**Before**
- [ ] Identify functions with nested control flow or too many args.
- [ ] Identify primitive IDs that should become strong types.
- [ ] Confirm invariant ownership (domain type vs use case).

**During**
- [ ] Flatten control flow with guard clauses.
- [ ] Extract helpers by intention, not by technical detail.
- [ ] Keep tenant scope (`{{DOMAIN_ACTOR}}_id`) explicit in all core inputs.

**After**
- [ ] Re-read function top-to-bottom: happy path is obvious.
- [ ] Ensure signatures are minimal and typed.
- [ ] Confirm no adapter imports leaked into `domain`/`application`.

## Definition of Done
- Core code is flatter, smaller, and easier to scan.
- Inputs/IDs are typed and tenant-safe.
- Domain invariants are explicit and tested.
- No hexagonal boundary violations.

## Minimal Examples
- Replace:
  ```go
  if req.{{DOMAIN_ACTOR_TITLE}}ID == "" {
      return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
  } else {
      // continue
  }
  ```
  With:
  ```go
  if req.{{DOMAIN_ACTOR_TITLE}}ID == "" {
      return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
  }
  // continue
  ```
- Prefer `SendMessageInput` over `SendMessage(ctx, {{DOMAIN_ACTOR}}ID, conversationID, requestID, body, source, locale string)`.
- Full before/after snippets: [examples_go.md](resources/examples_go.md).
