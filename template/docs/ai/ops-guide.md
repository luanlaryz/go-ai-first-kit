# Guia de operação AI — {{PROJECT_TITLE}}

Este guia orienta contribuições assistidas por AI no `{{PROJECT_SLUG}}`. Leia
primeiro [AGENTS.md](../../AGENTS.md); ele e as specs existentes são as fontes
normativas. Este documento, assim como o
[catálogo de capacidades](capabilities.md) e o
[inventário de maturidade](maturity-inventory.md), apenas torna o fluxo mais
fácil de localizar. Jornadas humanas passo a passo estão em
[docs/journeys/](../journeys/README.md).

## 1. Iniciar uma tarefa

1. Leia `AGENTS.md`, a spec que governa a mudança e
   `skills/00-skill-index/SKILL.md`.
2. Use [task-input-format.md](task-input-format.md) para estruturar o pedido.
3. Para uma nova trilha, decida entre amendar uma spec e criar o par
   construção/diagnóstico com a skill de spec architect.
4. Use os prompts em [prompts/](prompts/) e briefs em [briefs/](briefs/) como
   apoio operacional, nunca como substitutos de uma spec.
5. Declare objetivo, specs lidas, skills aplicadas, arquivos alterados,
   comandos, testes e gaps restantes ao concluir.

## 2. Artefatos disponíveis

- [capabilities.md](capabilities.md) cataloga cada capacidade, como ativá-la,
  a evidência esperada e os limites.
- [ai-contribution-contract.md](ai-contribution-contract.md) define o contrato
  mínimo de contribuição assistida por AI.
- [compliance-exceptions.md](compliance-exceptions.md) é o único caminho para
  exceções auditáveis.
- [spec-lineage.md](spec-lineage.md) explica como tratar as specs iniciais e
  futuras sem inventar histórico.
- [maturity-inventory.md](maturity-inventory.md) mapeia a baseline entregue e
  os limites do starter.
- `.cursor/rules/`, `skills/`, `automation/`, `docs/backlog/`,
  `docs/decisions/` e `docs/reports/` compõem a infraestrutura operacional.

## 3. Validar a entrega

Quando aplicáveis ao escopo, execute:

- `make fmt-check`
- `make lint`
- `make vet`
- `make check-compliance`
- `make test-security`
- `make test`
- `make race`
- `make coverage`
- `pre-commit run --all-files`

O report da trilha, não apenas testes verdes, decide o avanço de um autopilot.
Use `gakit diagnose --path <projeto>` para obter sinais adicionais sobre a
baseline AI-first e seus gaps.

## 4. Limites conhecidos

- O starter fornece uma baseline mínima de safety; uma suite ampla de
  regressão, operação hospedada e observabilidade de produção não estão
  declaradas como entregues.
- Specs, reports e capacidades de produto devem nascer no projeto quando uma
  necessidade aprovada os exigir. Paths ausentes são gaps a especificar, não
  evidência de trabalho já concluído.
- O projeto recém-gerado ainda não contém uma API, OpenAPI, `pkg/`,
  `internal/` ou `examples`; crie essas superfícies somente sob uma spec
  aprovada.
