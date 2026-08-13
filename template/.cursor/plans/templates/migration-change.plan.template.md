# Template: Migration Change

Frontmatter esperado ao copiar para um plano real:

```md
---
id: <plan-id>
status: DRAFT_NOT_EXECUTABLE
risk: high
area: migration
---
```

# Objetivo

Adicionar ou alterar schema via migrations versionadas, com par up/down aplicavel.

# Escopo permitido

- `migrations/` (par up/down)
- repositorios e adapters de persistencia correspondentes

# Fora de escopo

- DDL em construtor, handler ou codigo de aplicacao.
- Migration destrutiva sem spec, backfill e backup.

# Contratos afetados

- Spec de dados (construcao + diagnostico).
- Isolamento por {{DOMAIN_ACTOR}} quando o schema for multi-tenant.

# Plano de implementacao

1. Confirmar spec de dados e impacto em leitura/escrita existentes.
2. Criar o par up/down idempotente.
3. Ajustar repositorios e adapters.
4. Validar aplicacao e reversao contra banco limpo e banco com dados.

# Testes obrigatorios

- Aplicar e reverter a migration em ambiente local.
- `make test`
- `make guardrails`

# Rollback

- Aplicar a migration down e reverter o adapter no mesmo commit.

# Stop conditions

- Mudanca destrutiva sem spec aprovada.
- Falta de plano de backfill para dado existente.
- Reversao nao testada.
