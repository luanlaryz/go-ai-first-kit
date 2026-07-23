---
name: multi-provider-model-routing
description: Configure and read per-{{DOMAIN_ACTOR}} provider/model routing with Postgres source of truth, Redis cache, and JSONB params validation.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Multi-provider Model Routing
Goal: allow each {{DOMAIN_ACTOR}} to select providers/models/params safely while keeping Postgres authoritative and Redis fast.

## When to Use
- Updating {{DOMAIN_ACTOR}} model selection, routing logic, or cache behavior.
- Implementing new provider integrations or validation rules.
- Building admin APIs for {{DOMAIN_ACTOR}} model configuration.

## Non-negotiables
1. Postgres table `control.{{DOMAIN_ACTOR}}_model_config` is the single source of truth.
2. Redis cache TTL <=5m; always write-through/invalidate after DB updates.
3. Provider/model combos validated against allowlist prior to persisting.
4. JSONB `params` validated in application layer; store as-is after validation.
5. Config retrieval returns defaults when {{DOMAIN_ACTOR}}-specific row missing.

## Do / Don't
- **Do** wrap DB writes in transactions when multiple tables touched (e.g., audit log).
- **Do** log config loads with `{{DOMAIN_ACTOR}}_id`, `provider`, `model` (no secrets).
- **Do** guard concurrency by using `SELECT ... FOR UPDATE` when editing via admin API.
- **Don't** let workers fetch provider config per message without caching.
- **Don't** hardcode provider parameters in code; keep in JSONB per {{DOMAIN_ACTOR}}.
- **Don't** skip validation just because params are flexible.

## Interfaces / Contracts
- Schema + cache described in [model_config_schema.md](resources/model_config_schema.md).
- Port example:
  ```go
  type {{DOMAIN_ACTOR_TITLE}}ModelConfigStore interface {
      Get(ctx context.Context, {{DOMAIN_ACTOR}}ID string) ({{DOMAIN_ACTOR_TITLE}}ModelConfig, error)
      Update(ctx context.Context, cfg {{DOMAIN_ACTOR_TITLE}}ModelConfig) error
  }
  ```
- Validation contract: `ValidateProviderParams(provider string, params map[string]any) error` per provider module.

## Checklists
**Before coding**
- [ ] List providers/models affected; confirm allowlist updates.
- [ ] Decide default fallback behavior.
- [ ] Plan cache invalidation strategy.

**During**
- [ ] Fetch config via Redis first, then DB.
- [ ] Validate JSONB structure; reject unknown keys when necessary.
- [ ] Merge feature flags (skill 14) for dynamic overrides.

**After**
- [ ] Add unit/integration tests covering cache hit/miss and invalid configs.
- [ ] Document migrations if schema changed.
- [ ] Update Grafana panels for config load latency/cache hit ratio.

## Definition of Done
- Config API/use case returns correct provider/model per {{DOMAIN_ACTOR}}.
- Cache coherence proven via tests or TTL instrumentation.
- Validation errors clear and localized.
- Downstream consumers (ADK core, LLM adapters) read typed config.

## Minimal Examples
- Cache miss flow: `cfg, err := store.Get(ctx, {{DOMAIN_ACTOR}}ID)` -> miss -> query Postgres -> store in Redis -> return typed struct.
- Admin update: begin tx, update row, insert audit entry, delete Redis key, commit.
