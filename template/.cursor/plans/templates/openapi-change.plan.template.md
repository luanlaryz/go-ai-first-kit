# Template: Contract (OpenAPI) Change

Frontmatter esperado ao copiar para um plano real:

```md
---
id: <plan-id>
status: DRAFT_NOT_EXECUTABLE
risk: high
area: contract
---
```

# Objetivo

Alterar contrato publicado (`api/**`) mantendo compatibilidade declarada e clientes tipados em
sincronia.

# Escopo permitido

- `api/**` (spec do contrato)
- handlers e DTOs de transporte correspondentes
- `test/conformance`

# Fora de escopo

- Mudanca de comportamento de dominio sem spec propria.
- Breaking change sem decisao registrada em `docs/decisions/`.

# Contratos afetados

- Contrato publicado e clientes gerados.
- `specs/010-feature-matrix.md` quando a cobertura mudar.

# Plano de implementacao

1. Classificar a mudanca: aditiva, comportamental ou breaking.
2. Atualizar o contrato e regenerar clientes/tipos derivados.
3. Ajustar handlers e validacao de entrada.
4. Cobrir sucesso, erro e limite observavel no teste de conformidade.

# Testes obrigatorios

- `go test ./test/conformance/...`
- `make test`
- `make guardrails`

# Rollback

- Reverter contrato, clientes gerados e handlers no mesmo commit.

# Stop conditions

- Breaking change sem aprovacao registrada.
- Cliente gerado divergindo do contrato apos regeneracao.
