---
name: config-viper-flags
description: Manage configuration via Viper, load env-aware settings, and operate feature flags per {{DOMAIN_ACTOR}}.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Config & Feature Flags
Goal: centralize configuration (Viper) and {{DOMAIN_ACTOR}}-specific feature flags without hardcoding per-environment differences.

## When to Use
- Changing config structure, adding options, or wiring Viper usage.
- Implementing feature flag evaluation or overrides.
- Handling secrets, env var parsing, or config-driven behavior.

## Non-negotiables
1. Viper loads defaults, config file, then env vars (highest priority).
2. Config objects passed via dependency injection; no global `viper.Get` outside setup.
3. Feature flags stored in Postgres + cached in Redis (60s TTL) with per-{{DOMAIN_ACTOR}} overrides.
4. {{DOMAIN_ACTOR_TITLE}}-level toggles evaluated in application layer before calling providers.
5. Secrets fetched from secure store (SSM/Doppler) not committed.

## Do / Don't
- **Do** create typed config structs covering all modules; avoid map[string]any.
- **Do** expose config via `pkg/config` package returning immutable copies.
- **Do** document new fields in [config_shape.md](resources/config_shape.md).
- **Don't** read env vars deep inside business logic.
- **Don't** use feature flags for sensitive security controls (deploy gating instead).
- **Don't** rely on default values silently; log config summary on startup.

## Interfaces / Contracts
- Config struct example:
  ```go
  type AppConfig struct {
      Env   string
      HTTP  HTTPConfig
      Redis RedisConfig
      Flags FlagConfig
  }
  ```
- Feature flag store: `flagStore.Enabled(ctx, {{DOMAIN_ACTOR}}ID, "streaming_default")`.
- Config shape documented in resource file.

## Checklists
**Before coding**
- [ ] Decide which module(s) need new config knobs.
- [ ] Define defaults + env overrides.
- [ ] Plan migration path for existing deployments.

**During**
- [ ] Update config structs + Viper bindings.
- [ ] Add validation logic (panic on missing required fields) during startup.
- [ ] Implement flag evaluation paths with caching + metrics.

**After**
- [ ] Update docs/resources + sample env files.
- [ ] Add tests ensuring config parsing works (use `t.Setenv`).
- [ ] Bump config version in release/changelog.

## Definition of Done
- Config struct + Viper wiring compile and pass tests.
- Feature flags accessible per {{DOMAIN_ACTOR}} with redis cache + fallback.
- Secrets handled securely.
- Documented shape + usage instructions.

## Minimal Examples
- Startup: `cfg := config.Load()` -> `server := app.NewServer(cfg)`.
- Flag check: `if !flagStore.Enabled(ctx, {{DOMAIN_ACTOR}}ID, "streaming_default") { disableStreaming() }`.
