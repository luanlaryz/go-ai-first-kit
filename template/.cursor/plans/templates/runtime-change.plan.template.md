# Template: Runtime Change

Frontmatter esperado ao copiar para um plano real:

```md
---
id: <plan-id>
status: DRAFT_NOT_EXECUTABLE
risk: medium
area: runtime
---
```

# Objetivo

Alterar comportamento observavel do runtime interno preservando a fronteira `pkg/` <-> `internal/`.

# Escopo permitido

- `internal/runtime` e pacotes internos afetados
- `pkg/**` somente quando a spec autorizar mudanca de API publica

# Fora de escopo

- Expor tipo de `internal/*` em assinatura exportada.
- Nova dependencia externa sem decisao registrada.

# Contratos afetados

- Spec de construcao do modulo e sua spec de diagnostico.
- Contrato publico quando a mudanca alcancar `pkg/**`.

# Plano de implementacao

1. Declarar o comportamento antes e depois, de forma observavel.
2. Confirmar se a mudanca exige atualizacao de API publica.
3. Implementar preservando propagacao de `context.Context` e cancelamento.
4. Cobrir sucesso, falha e cancelamento no teste.

# Testes obrigatorios

- `make test`
- `make race`
- `make guardrails`

# Rollback

- Reverter o pacote alterado; API publica nao deve ter mudado se o plano era interno.

# Stop conditions

- Mudanca exigiria quebrar contrato publico sem aprovacao.
- Comportamento novo nao coberto por spec.
