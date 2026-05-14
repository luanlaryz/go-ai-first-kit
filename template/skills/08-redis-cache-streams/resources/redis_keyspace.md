# Redis Keyspace Map

| Purpose         | Key Pattern                                           | TTL          |
|-----------------|-------------------------------------------------------|--------------|
| {{DOMAIN_ACTOR_TITLE}} config   | `{{DOMAIN_ACTOR}}:model-config:{{{DOMAIN_ACTOR}}_id}`                     | 5m           |
| Rate limit      | `rate:{{DOMAIN_ACTOR}}:{{{DOMAIN_ACTOR}}_id}:window:{minute}`             | 60s          |
| Locks           | `lock:{namespace}:{{{DOMAIN_ACTOR}}_id}:{resource}`             | 15s default  |
| Conversation KV | `conv:state:{{{DOMAIN_ACTOR}}_id}:{conversation_id}`            | 24h          |
| Result store    | `result:{command_id}`                                 | 15m          |
| Redis Stream    | `stream:conversation:{{{DOMAIN_ACTOR}}_id}:{conversation_id}`   | retention 24h|
| Idempotency     | `inbox:{{{DOMAIN_ACTOR}}_id}:{command_id}`                      | 24h          |

## Scripts
- Lua script for rate limit: atomic increment + expiry when first set.
- Lock acquisition: `SET lock ... NX PX 15000` + unique token, release via Lua compare/delete.

## Observability Keys
- `metrics:rate_limit:{{{DOMAIN_ACTOR}}_id}` for aggregated counters.
- Use Redis MONITOR only locally (heavy!).
