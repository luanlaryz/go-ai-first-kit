---
name: tenant-auth-quotas
description: Apply multi-tenant ({{DOMAIN_ACTOR}}) context, JWT/API key auth, CRM-agnostic adapters, and per-{{DOMAIN_ACTOR}} quotas/rate limits.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Tenant Auth & Quotas
Goal: guarantee every request is authenticated, scoped, and throttled per {{DOMAIN_ACTOR}} with CRM adapters isolated.

## When to Use
- Building or modifying auth middleware, tenant context propagation, CRM adapters, or per-{{DOMAIN_ACTOR}} quotas.
- Touching Redis rate limit keys or tenant-aware config.
- Handling {{DOMAIN_ACTOR}} onboarding/offboarding logic.

## Non-negotiables
1. Every inbound request yields a `TenantContext` (see resource file) stored in `context.Context`.
2. Auth supports both JWT and API key; pick via `Authorization` scheme.
3. Rate limiting enforced before calling application layer; tokens stored per {{DOMAIN_ACTOR}}.
4. CRM-specific code stays in adapters and receives sanitized tenant context copies.
5. Logs, metrics, traces include `{{DOMAIN_ACTOR}}_id`, `auth_method`, `rate_limit_tier`.

## Do / Don't
- **Do** short-circuit unauthorized traffic with 401/403 before hex layers.
- **Do** expose derived flags (features, tiers) via context to use cases.
- **Do** cache JWKS per {{DOMAIN_ACTOR}} with TTL + background refresh.
- **Don't** hardcode {{DOMAIN_ACTOR}}-specific behavior; use config skill for overrides.
- **Don't** bypass rate limits in tests without explicit helper toggles.
- **Don't** log raw API keys or JWTs.

## Interfaces / Contracts
- Middleware signature: `func InjectTenantContext(ctx context.Context, headers http.Header) (context.Context, error)`.
- `TenantContext` schema lives in [tenant_context_contract.md](resources/tenant_context_contract.md).
- Rate limit store port example:
  ```go
  type RateLimiter interface {
      Allow(ctx context.Context, {{DOMAIN_ACTOR}}ID, key string, limit int, window time.Duration) (allowed bool, remaining int, reset time.Time, err error)
  }
  ```

## Checklists
**Before coding**
- [ ] Identify auth methods affected (JWT, API key, both).
- [ ] Determine required scopes/tiers for feature.
- [ ] Confirm rate limit thresholds per {{DOMAIN_ACTOR}} (default + overrides).

**During**
- [ ] Validate token signatures and expiration.
- [ ] Populate context with canonical IDs + correlation metadata.
- [ ] Apply Redis quota mutation atomically with scripts or Lua.

**After**
- [ ] Add unit tests covering allowed, throttled, and unauthorized cases.
- [ ] Update OpenAPI/GRPC docs to reflect auth scheme.
- [ ] Verify logs redact secrets.

## Definition of Done
- Tenant context available to downstream handlers/use cases.
- Rate limit metrics exported per {{DOMAIN_ACTOR}} tier.
- CRM adapters receive only sanitized, necessary attributes.
- Error responses include traceable IDs and skill-aligned messages.

## Minimal Examples
- HTTP middleware verifying JWT, storing `TenantContext`, passing to Gin handlers via `ctx.Request.Context()`.
- CLI script creating API key row with hashed secret + tier flag, then invalidating Redis caches.
