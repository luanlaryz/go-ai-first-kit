# Catálogo de capacidades — {{PROJECT_TITLE}}

Este catálogo mapeia as capacidades operacionais do `{{PROJECT_SLUG}}`: o que cada uma entrega, como ativá-la e qual evidência confirma seu uso. Ele é um guia de navegação; as fontes normativas continuam sendo [AGENTS.md](../../AGENTS.md) e as specs em [specs/](../../specs/).

## Como ler este catálogo

Tipos de verificação:

- `automático`: existe comando ou script executável cuja saída confirma a capacidade.
- `processual`: workflow governado operado por um agente (humano ou AI); a evidência são specs, reports e arquivos de estado versionados — não existe daemon nem runner hospedado.
- `guia`: orientação normativa (skills, políticas, checklists); nada de código ou infraestrutura foi instalado por ela.

Estados:

- `baseline entregue`: o artefato existe neste repositório recém-renderizado.
- `evidência a produzir`: o processo existe, mas este projeto ainda precisa gerar a própria evidência (reports, specs novas, itens de backlog).
- `fora de escopo`: intencionalmente não entregue; ver a seção final.

## Capacidades

### Governança para agentes

- Tipo: `guia` com enforcement parcial `automático`. Estado: `baseline entregue`.
- Artefatos: [AGENTS.md](../../AGENTS.md), [.cursor/rules/](../../.cursor/rules/), [.cursor/hooks/](../../.cursor/hooks/) (gofmt pós-edição), [skills/00-skill-index/SKILL.md](../../skills/00-skill-index/SKILL.md).
- Como ativar: leia `AGENTS.md` e o índice de skills antes de qualquer edição; o roteamento por intenção está na tabela do índice.
- Evidência: PRs declarando specs lidas e skills aplicadas, conforme o template de PR.
- Limites: as skills de infraestrutura (02–04, 06–12, 14, 18–19) são guidance herdada para quando o produto exigir aquela stack; nenhuma dependência de Redis, SQS, Postgres, Gin ou gRPC está instalada.

### Spec Driven Development dual-spec

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: [specs/000-project-mission.md](../../specs/000-project-mission.md), [specs/001-non-goals.md](../../specs/001-non-goals.md), [specs/010-feature-matrix.md](../../specs/010-feature-matrix.md), [specs/020-repository-architecture.md](../../specs/020-repository-architecture.md), par [specs/680-phase-0-bootstrap-foundation.md](../../specs/680-phase-0-bootstrap-foundation.md) / [specs/681-phase-0-bootstrap-foundation-diagnosis.md](../../specs/681-phase-0-bootstrap-foundation-diagnosis.md); skills [05](../../skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md) e [25](../../skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md).
- Como ativar: toda trilha nova exige spec de construção e spec de diagnóstico antes de implementar; a decisão entre amendar e criar segue a skill 05.
- Evidência a produzir: as duas specs da trilha, versionadas, mais o report que as valida.

### Autopilot de roadmap (`phase_autopilot`)

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: [automation/AUTOPILOT.md](../../automation/AUTOPILOT.md), [automation/RUNBOOK.md](../../automation/RUNBOOK.md), `automation/ROADMAP.json` (fase 0), `automation/PHASE_STATE.json`, [automation/STOP_CONDITIONS.md](../../automation/STOP_CONDITIONS.md).
- Como ativar: use o prompt mestre da seção 15 do runbook; o agente executa somente a fase indicada por `PHASE_STATE.json`.
- Evidência a produzir: report da fase no caminho exato declarado em `ROADMAP.json`, com `classification` e `decision` iguais ao gate.
- Limites: não é um serviço; é uma máquina de estados operada por agente. Testes verdes sem report não autorizam avanço.
- Jornada: [docs/journeys/02-fase-do-roadmap.md](../journeys/02-fase-do-roadmap.md).

### Autopilot interativo (`interactive_sdd_autopilot`)

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: [automation/INTERACTIVE_AUTOPILOT.md](../../automation/INTERACTIVE_AUTOPILOT.md), [automation/INTERACTIVE_RUNBOOK.md](../../automation/INTERACTIVE_RUNBOOK.md), `automation/INTERACTIVE_STATE.json`, [docs/interactive-sdd-autopilot.md](../interactive-sdd-autopilot.md), [docs/feature-request-lifecycle.md](../feature-request-lifecycle.md), skill [23](../../skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md).
- Como ativar: pedido explícito com escopo conhecido e fora de escopo; modelos de pedido em [docs/interactive-sdd-autopilot.md](../interactive-sdd-autopilot.md).
- Evidência a produzir: intake registrado, dual-spec da trilha, report com `PASS` e `INTERACTIVE_STATE.json` coerente.
- Limites: aditivo ao roadmap fixo; nunca altera `ROADMAP.json` nem `PHASE_STATE.json`.
- Jornada: [docs/journeys/03-backlog-e-trilha-sdd.md](../journeys/03-backlog-e-trilha-sdd.md).

### Backlog governado

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: [docs/backlog/Backlog.md](../backlog/Backlog.md) (canônico, vazio por design) e skill [26](../../skills/26-backlog-item-intake/SKILL.md) com templates de item, regras de classificação e prompts de autopilot.
- Como ativar: todo pedido, gap, bug ou ideia passa pela skill 26 antes de virar item `BLG-*`.
- Evidência a produzir: itens `BLG-*` com classificação, evidência e plano; nenhum item vira implementação sem dual-spec.

### ADR, release e changelog

- Tipo: `processual` com `guia`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: [docs/decisions/](../decisions/) (ADR seed 0001), [docs/release-versioning-policy.md](../release-versioning-policy.md), [docs/release-notes-policy.md](../release-notes-policy.md), [docs/release-checklist.md](../release-checklist.md), `CHANGELOG.md`, skill [22](../../skills/22-release-versioning-governance/SKILL.md).
- Como ativar: decisões não óbvias viram ADR; toda recomendação de release preenche o checklist e escolhe um caminho de distribuição.
- Evidência a produzir: ADRs numerados, changelog atualizado e checklists preenchidos por release.
- Jornada: [docs/journeys/04-release-e-decisoes.md](../journeys/04-release-e-decisoes.md).

### Qualidade, segurança e compliance

- Tipo: `automático`. Estado: `baseline entregue`.
- Artefatos: [Makefile](../../Makefile), [scripts/check-compliance.sh](../../scripts/check-compliance.sh), [scripts/check-secrets.sh](../../scripts/check-secrets.sh), [scripts/check-pr-body.sh](../../scripts/check-pr-body.sh), [scripts/check-file-size.sh](../../scripts/check-file-size.sh), [.pre-commit-config.yaml](../../.pre-commit-config.yaml), [.github/workflows/ci.yml](../../.github/workflows/ci.yml), [test/security/](../../test/security/).
- Como ativar: `make check-compliance`, `make fmt-check`, `make lint`, `make vet`, `make test`, `make race`, `make test-security`, `pre-commit run --all-files`.
- Evidência: saída dos comandos; o CI executa os mesmos gates em PR.
- Limites: a baseline de segurança é mínima (suite de presença, secrets scan e govulncheck); não é uma suite ampla de regressão de safety.

### Diagnóstico de maturidade AI-first

- Tipo: `automático` (ferramenta externa `gakit`). Estado: `evidência a produzir`.
- Como ativar: `gakit diagnose --path . --report-only`.
- Evidência: relatório com score por pilar e findings.
- Limites: é sinal de qualidade, não certificação; em um starter recém-gerado, os pilares Hexagonal e OpenAPI apontam ausências por design.

## Fora de escopo

- Capability de produto: API pública, `pkg/`, `internal/`, `examples/` e contrato OpenAPI nascem apenas sob spec aprovada. Até lá, são gaps declarados, não defeitos.
- Serviços hospedados, control planes, dashboards, runners gerenciados e deploy tooling obrigatório.
- Stack das skills de infraestrutura: as skills orientam a implementação futura, não instalam dependências.

## Relação com os demais documentos

- [maturity-inventory.md](maturity-inventory.md): mapa de baseline, evidência a produzir e limites — a visão de estado.
- [ops-guide.md](ops-guide.md): guia de execução de uma tarefa — a visão de fluxo.
- [docs/journeys/README.md](../journeys/README.md): jornadas humanas passo a passo.
- [docs/README.md](../README.md): central de navegação por papel.
