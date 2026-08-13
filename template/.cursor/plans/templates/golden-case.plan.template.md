# Template: Golden Case

Frontmatter esperado ao copiar para um plano real:

```md
---
id: <plan-id>
status: DRAFT_NOT_EXECUTABLE
risk: low
area: tests
---
```

# Objetivo

Adicionar um caso golden ou de conformidade que fixe comportamento observavel ja especificado.

# Escopo permitido

- `test/conformance`, `test/fixtures`

# Fora de escopo

- Alterar codigo de producao para fazer o golden passar.
- Golden que congele comportamento nao especificado (isso transforma bug em contrato).

# Contratos afetados

- Spec que define o comportamento sendo fixado.

# Plano de implementacao

1. Citar a secao da spec que define o comportamento esperado.
2. Criar a fixture de entrada, redigindo qualquer dado sensivel.
3. Gerar o golden e revisar o conteudo linha a linha.
4. Confirmar que o teste falha se o comportamento regredir.

# Testes obrigatorios

- `go test ./test/conformance/...`
- `make test`

# Rollback

- Remover fixture e golden adicionados.

# Stop conditions

- Comportamento observado divergindo da spec: corrigir a spec ou o codigo antes de fixar o golden.
- Fixture exigiria dado sensivel real.
