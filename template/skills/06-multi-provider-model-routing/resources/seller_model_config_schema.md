# {{DOMAIN_ACTOR}}_model_config Schema

Table: `control.{{DOMAIN_ACTOR}}_model_config`

| Column              | Type        | Notes |
|---------------------|-------------|-------|
| {{DOMAIN_ACTOR}}_id           | text PK     | matches tenant context |
| provider            | text        | e.g., `openai`, `anthropic`, `vertex` |
| model               | text        | provider-specific model id |
| params              | jsonb       | arbitrary provider config (temperature, top_p, tool hints) |
| max_output_tokens   | int         | guardrail |
| tools               | jsonb       | list of enabled tools |
| updated_at          | timestamptz | audit |
| updated_by          | text        | human or system |

## Redis Cache
- Key: `{{DOMAIN_ACTOR}}:model-config:{{{DOMAIN_ACTOR}}_id}`
- TTL: 5 minutes (refresh early on misses or critical updates).
- Value: JSON copy of table row.

## API Validation Rules
- Validate provider/model against allowlist.
- Validate `params` schema per provider at application layer (not DB constraint).
- Default fallback row stored under {{DOMAIN_ACTOR}} `default` per environment.

## Multi-provider Routing
1. Fetch config from Redis; on miss, query Postgres and warm cache.
2. Merge feature flags (skill 14) before returning to caller.
3. Provide typed struct:
```go
type {{DOMAIN_ACTOR_TITLE}}ModelConfig struct {
    {{DOMAIN_ACTOR_TITLE}}ID string
    Provider string
    Model    string
    Params   map[string]any
    MaxOutputTokens int
    Tools   []string
}
```
