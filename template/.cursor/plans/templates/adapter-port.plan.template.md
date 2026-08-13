# Template: Adapter Port

Frontmatter esperado ao copiar para um plano real:

```md
---
id: <plan-id>
status: DRAFT_NOT_EXECUTABLE
risk: medium
area: adapters
---
```

# Objetivo

Implementar ou portar um adapter de borda para {{EXTERNAL_SYSTEM}}, mapeando o payload do sistema
externo para o DTO canonico do core.

# Escopo permitido

- `internal/adapters/<adapter>`
- `test/fixtures/<sistema-externo>`, `test/conformance`

# Fora de escopo

- Tipo do sistema externo vazando para o dominio ou para a camada de aplicacao.
- Chamada real ao sistema externo em teste.

# Contratos afetados

- Spec do adapter (construcao + diagnostico).
- Contrato canonico de entrada do core.

# Plano de implementacao

1. Confirmar o payload externo e o mapeamento canonico na spec.
2. Criar fixtures de borda a partir de payload real redigido.
3. Implementar o mapeamento no adapter.
4. Provar que nenhum tipo do sistema externo alcanca `pkg/**` nem o core.

# Testes obrigatorios

- `go test ./test/conformance/...`
- `make test`
- `make guardrails`

# Rollback

- Reverter adapter e fixtures adicionados; o core nao deve ter mudado.

# Stop conditions

- Contrato do sistema externo indisponivel ou ambiguo.
- Mapeamento exigiria expor tipo externo em API publica.
