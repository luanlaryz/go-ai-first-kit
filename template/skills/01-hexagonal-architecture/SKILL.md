---
name: hexagonal-architecture
description: Maintain the hexagonal (ports-and-adapters) structure for {{PROJECT_SLUG}}. Enforce layered separation between public contracts (pkg/*), runtime internals (internal/*), and composition root (pkg/app).
---

# Hexagonal Architecture Playbook
Goal: keep the codebase layered so features remain testable, extensible, and free from internal leakage.

## When to Use
- Adding or modifying public contracts, runtime logic, adapters, or command binaries.
- Touching wiring/DI for composition root or server entry points.
- Reviewing PRs that blend responsibilities across layers.

## Non-negotiables
1. `pkg/*` contains stable public contracts; `internal/*` contains runtime details.
2. `pkg/app` is the only public package authorized to import `internal/runtime`.
3. Adapters live behind ports; each adapter implements one interface.
4. CMD packages only assemble dependencies, parse config, and start servers.
5. Prefer small interfaces and composition over large abstractions.

## Do / Don't
- **Do** keep structs immutable via constructors or value copies; return domain errors.
- **Do** define ports near consumers (e.g., `agent.Engine`, `memory.Store`).
- **Do** create DTOs in adapters when external schemas differ from domain models.
- **Don't** import adapters into public contract packages.
- **Don't** expose `internal/*` types in exported signatures, errors, or examples.
- **Don't** leak infrastructure details into public API surface.

## Interfaces / Contracts
- Ports should be consumer-driven and tiny (see Effective Go skill). Example:
  ```go
  type Store interface {
      Load(ctx context.Context, sessionID string) (Snapshot, error)
      Save(ctx context.Context, sessionID string, delta Delta) error
  }
  ```
- Wiring contract: `pkg/app` builds the composition root; tests supply fakes via public interfaces.

## Checklists
**Before coding**
- [ ] Identify which layer changes; confirm dependencies only flow inward.
- [ ] Define/adjust port interfaces before touching adapters.
- [ ] Plan DTOs for any new external schemas.

**During**
- [ ] Keep constructors in adapters, injecting via interfaces.
- [ ] Add unit tests per layer (public contract tests, runtime detail tests, wiring smoke tests).

**After**
- [ ] Verify `go test ./...` in touched packages.
- [ ] Update diagrams/docs if new adapter exists.
- [ ] Mention port/interface changes in PR description.

## Definition of Done
- Layer boundaries intact; no forbidden imports.
- Ports documented and implemented with fakes in tests.
- Composition root compiles and respects config skill.
- Public contract tests cover new logic; adapters have coverage or manual verification noted.

## Minimal Examples
- Adding a new memory backend: create adapter implementing `memory.Store` port; public contract unaware of backend specifics.
- New tool: implement `tool.Tool` interface in adapter; register via `pkg/tool.Registry`.
