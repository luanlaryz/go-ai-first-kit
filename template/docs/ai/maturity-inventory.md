# Inventário de Maturidade AI — {{PROJECT_TITLE}}

Este documento apresenta a baseline AI-first entregue pelo starter de
`{{PROJECT_SLUG}}`. Ele é um mapa de artefatos versionados, não uma
certificação de produção e não substitui specs, reports ou decisões humanas.

Ele responde "qual é o estado da baseline". Para "o que cada capacidade faz e
como ativá-la", use o [catálogo de capacidades](capabilities.md); para o fluxo
de execução de uma tarefa, use o [guia de operação](ops-guide.md).

## Baseline entregue

### Contrato para agentes e regras de execução

- [AGENTS.md](../../AGENTS.md) define missão, fontes de verdade, workflow
  obrigatório e formato de resposta.
- [.cursor/rules/](../../.cursor/rules/) aplica regras modulares para baseline,
  dual-spec, autopilot, intake e stop conditions.
- [skills/00-skill-index/SKILL.md](../../skills/00-skill-index/SKILL.md)
  direciona o agente para a skill adequada antes de editar.

### Spec Driven Development e evidência

- [specs/](../../specs/) contém missão, non-goals, feature matrix, arquitetura
  e o par inicial de construção/diagnóstico.
- [skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md](../../skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md)
  orienta a decisão entre amendar uma spec ou abrir uma nova trilha dual-spec.
- [docs/ai/task-input-format.md](task-input-format.md) padroniza a entrada de
  tarefas; [docs/ai/ai-contribution-contract.md](ai-contribution-contract.md)
  define as regras mínimas de contribuição assistida por AI.

### Autopilot, backlog e decisões

- [automation/](../../automation/) separa o `phase_autopilot`, orientado pelo
  roadmap, do `interactive_sdd_autopilot`, orientado por uma solicitação.
- [docs/backlog/Backlog.md](../backlog/Backlog.md) e
  [skills/26-backlog-item-intake/SKILL.md](../../skills/26-backlog-item-intake/SKILL.md)
  fornecem intake, triagem e classificação governada.
- [docs/decisions/](../decisions/) e as políticas de
  [release](../release-versioning-policy.md) mantêm decisões, changelog e
  critérios de entrega rastreáveis.

### Enforcement, qualidade e segurança

- [scripts/check-compliance.sh](../../scripts/check-compliance.sh),
  [Makefile](../../Makefile), [.github/workflows/ci.yml](../../.github/workflows/ci.yml)
  e [.pre-commit-config.yaml](../../.pre-commit-config.yaml) materializam
  checks de formatação, lint, vet, compliance, segurança e testes.
- [test/security/](../../test/security/) contém a baseline mínima de testes de
  segurança.
- [skills/24-agent-readiness-governance/SKILL.md](../../skills/24-agent-readiness-governance/SKILL.md)
  define como usar readiness como evidência auxiliar, sem substituir o gate da
  trilha.

## Evidência a produzir

O starter fornece processo e estrutura; cada projeto deve produzir evidências
da sua própria evolução:

1. registrar specs de construção e diagnóstico para novas trilhas;
2. gerar reports em `docs/reports/` com classificação e decisão explícitas;
3. atualizar o backlog, ADRs e changelog quando o escopo exigir;
4. executar os checks aplicáveis do `Makefile` e do CI;
5. executar `gakit diagnose --path <projeto>` para avaliar os pilares
   AI-first do estado atual.

Relatórios de readiness, segurança ou diagnóstico pertencem ao projeto que os
produziu. Não copie reports de outro repositório como evidência local.

## Gaps e limites

- O starter não implementa uma capability de produto, API pública, provider,
  serviço hospedado ou control plane.
- `pkg/`, `internal/`, `examples/` e um contrato OpenAPI devem ser criados
  apenas quando uma spec aprovada os exigir. Até lá, alguns pilares do
  diagnóstico podem apontar esses itens como ausentes.
- A baseline de segurança é mínima; uma suite ampla de regressão de safety,
  observabilidade de produção e operação hospedada não são declaradas como
  entregues.
- O inventário não autoriza alterar código fora de spec nem avançar um
  autopilot sem report `PASS`.

## Como manter este inventário honesto

Ao alterar esta baseline, atualize este documento somente com links para
artefatos realmente presentes no repositório renderizado. Descreva capacidades
futuras como gaps ou trabalho a especificar; nunca como entrega concluída.
