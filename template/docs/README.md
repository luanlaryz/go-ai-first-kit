# Documentação do {{PROJECT_TITLE}}

Central de navegação da documentação. As fontes normativas são [AGENTS.md](../AGENTS.md) e as specs em [specs/](../specs/); tudo aqui é apoio de descoberta e operação.

## Comece por papel

### Quero pedir uma mudança (autor de solicitação)

1. Estruture o pedido com [ai/task-input-format.md](ai/task-input-format.md).
2. Entenda o ciclo humano em [feature-request-lifecycle.md](feature-request-lifecycle.md).
3. Para pedidos que viram trilha SDD, use os modelos de [interactive-sdd-autopilot.md](interactive-sdd-autopilot.md).
4. Para registrar sem implementar agora, use o backlog: [journeys/03-backlog-e-trilha-sdd.md](journeys/03-backlog-e-trilha-sdd.md).

### Vou implementar (agente ou pessoa)

1. Leia [AGENTS.md](../AGENTS.md) e [skills/00-skill-index/SKILL.md](../skills/00-skill-index/SKILL.md).
2. Localize a capacidade e seus limites em [ai/capabilities.md](ai/capabilities.md).
3. Siga o fluxo de execução em [ai/ops-guide.md](ai/ops-guide.md).
4. Primeira contribuição? [journeys/01-primeira-contribuicao.md](journeys/01-primeira-contribuicao.md).

### Vou revisar (revisor de PR ou de trilha)

1. Checklist humano de aceite em [feature-request-lifecycle.md](feature-request-lifecycle.md).
2. Contrato mínimo em [ai/ai-contribution-contract.md](ai/ai-contribution-contract.md).
3. Exceções só valem se registradas em [ai/compliance-exceptions.md](ai/compliance-exceptions.md).

### Cuido de release e decisões (mantenedor)

1. [journeys/04-release-e-decisoes.md](journeys/04-release-e-decisoes.md).
2. Políticas: [release-versioning-policy.md](release-versioning-policy.md), [release-notes-policy.md](release-notes-policy.md), [release-checklist.md](release-checklist.md).
3. Decisões não óbvias viram ADR em [decisions/](decisions/).

### Cuido do backlog (triagem)

1. [backlog/Backlog.md](backlog/Backlog.md) é o backlog canônico.
2. Todo intake passa pela skill [26-backlog-item-intake](../skills/26-backlog-item-intake/SKILL.md).

## Mapa da documentação

- [ai/capabilities.md](ai/capabilities.md): catálogo de capacidades, tipos de verificação e limites.
- [ai/maturity-inventory.md](ai/maturity-inventory.md): baseline entregue, evidência a produzir e gaps.
- [ai/ops-guide.md](ai/ops-guide.md): como executar e validar uma tarefa.
- [journeys/README.md](journeys/README.md): jornadas humanas passo a passo.
- [ai/spec-lineage.md](ai/spec-lineage.md): como as specs nascem e evoluem.
- [interactive-sdd-autopilot.md](interactive-sdd-autopilot.md) e [feature-request-lifecycle.md](feature-request-lifecycle.md): trilhas interativas.
- `reports/` e `plans/`: artefatos de evidência produzidos pelo próprio projeto.

## O que este starter ainda não contém

`pkg/`, `internal/`, `examples/`, API pública e contrato OpenAPI não existem em um projeto recém-gerado: são criados apenas sob spec aprovada. A lista completa de limites está em [ai/capabilities.md](ai/capabilities.md) e [ai/maturity-inventory.md](ai/maturity-inventory.md).
