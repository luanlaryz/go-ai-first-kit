---
name: solid-go-hexagonal-ports
description: Apply SOLID in Go with hexagonal architecture using small ports, explicit adapters, and strict dependency direction for {{PROJECT_SLUG}}.
---

# SOLID for Go + Ports/Adapters
Goal: enforce SOLID with Go idioms so public contracts (`pkg/*`) and runtime internals (`internal/*`) stay stable, testable, and independent from infrastructure.

## When to Use
- Creating or changing public interfaces in `pkg/*`.
- Implementing new adapters (memory backends, tool adapters, server transports).
- Reviewing contracts for responsibility leaks across `pkg/` and `internal/` boundaries.

## Non-negotiables
1. DIP: public contracts depend on interfaces only; runtime wiring belongs to `pkg/app` and `internal/runtime`.
2. ISP: ports are small (1-3 methods) and defined by consumer needs.
3. SRP: public types handle contracts; runtime handles orchestration; adapters integrate external systems.
4. OCP: new backends arrive as new adapters; core logic stays unchanged.
5. LSP: fakes/mocks preserve invariants, errors, and expected behavior contracts.
6. `pkg/*` never imports `internal/*` except through `pkg/app`.
7. Contracts keep context explicit via `context.Context` as the standard boundary.

## Do / Don't
- **Do** split broad interfaces into focused ports.
- **Do** return classified errors from public contracts, not provider-specific details.
- **Do** keep adapter mapping isolated at boundaries.
- **Don't** put orchestration logic in transport handlers.
- **Don't** make a shared mega-port for unrelated capabilities.
- **Don't** branch core logic by provider name; use adapter polymorphism.

## Interfaces / Contracts
- Consumer-owned port:
  ```go
  type Store interface {
      Load(ctx context.Context, sessionID string) (Snapshot, error)
      Save(ctx context.Context, sessionID string, delta Delta) error
  }
  ```
- Engine boundary:
  ```go
  type Engine interface {
      Run(ctx context.Context, agent *Agent, req Request) (Response, error)
      Stream(ctx context.Context, agent *Agent, req Request) (Stream, error)
  }
  ```
- Error contract: fakes must return `ErrNotFound`, `ErrInvalidConfig`, or wrapped errors equivalent to real adapters.

## Checklists
**Before**
- [ ] Decide which layer owns the new behavior (public contract / runtime / adapter).
- [ ] Define minimal ports from consumer needs.
- [ ] Identify invariant/error contract to preserve in fakes.

**During**
- [ ] Keep interfaces in consumer package, not provider package.
- [ ] Ensure adapter only translates protocol/SDK concerns.

**After**
- [ ] Add unit tests with fakes for contract behavior.
- [ ] Verify new adapter works by wiring changes only in `pkg/app` or `cmd/*`.
- [ ] Confirm no public package imports internal packages directly.

## Definition of Done
- Ports are minimal, cohesive, and consumer-driven.
- New backends can be added without core edits.
- Fakes behave like production adapters for success and failures.
- Dependency direction remains hexagonal and compile-time safe.

## Minimal Examples
- Adding a memory backend: implement `memory.Store` in adapter package; public contract unaware of backend specifics.
- Split interface:
  ```go
  type Engine interface { Run(ctx context.Context, agent *Agent, req Request) (Response, error) }
  type HistorySink interface { Append(ctx context.Context, entry HistoryEntry) error }
  ```
- Review checklist: [ports_checklist.md](resources/ports_checklist.md).
