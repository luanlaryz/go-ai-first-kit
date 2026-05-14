# Guia de operacao AI - {{PROJECT_SLUG}}

Base normativa atual:

- `AGENTS.md`
- `specs/010-feature-matrix.md`
- `specs/348-ai-operational-maturity-baseline-recovery.md`
- `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md`
- `specs/350-ai-operational-maturity-expansion.md`
- `specs/351-ai-operational-maturity-expansion-diagnosis.md`
- `specs/352-ai-historical-spec-convergence.md`
- `specs/353-ai-historical-spec-convergence-diagnosis.md`
- `specs/342-ai-development-remediation-closure.md`
- `specs/343-ai-development-remediation-closure-diagnosis.md`
- `specs/344-ai-development-doc-coherence-cleanup.md`
- `specs/345-ai-development-doc-coherence-cleanup-diagnosis.md`
- `specs/346-ai-development-doc-truthfulness-closure.md`
- `specs/347-ai-development-doc-truthfulness-closure-diagnosis.md`
- `docs/ai/spec-lineage.md`

Este guia descreve apenas os artefatos de AI development realmente versionados no repositorio em 2026-04-09.

## 1. Iniciar tarefa com Codex ou Cursor

1. Ler `AGENTS.md` e a spec governante da mudanca.
2. Ler `specs/010-feature-matrix.md` quando a tarefa tocar status, cobertura ou paridade.
3. Ler `skills/00-skill-index/SKILL.md` e as skills aplicaveis ao escopo.
4. Usar `docs/ai/task-input-format.md` como formato padrao de entrada.
5. Reutilizar os prompts oficiais de `docs/ai/prompts/` e os briefs versionados em `docs/ai/briefs/` quando fizer sentido.
6. Exigir que a resposta do agente liste objetivo, specs lidas, skills aplicadas, arquivos alterados, comandos executados, testes executados e gaps restantes.

Estado atual de assets operacionais:

- existe `docs/ai/ai-contribution-contract.md`
- existe `docs/ai/task-input-format.md`
- existe `docs/ai/compliance-exceptions.md` com caminho formal para excecoes auditaveis
- existe `docs/ai/prompts/` com prompts reais para implementacao, review, diagnostico e remediation/bugfix
- existe `docs/ai/briefs/` com briefs reais para implementacao, diagnostico e remediation/bugfix
- existe `.cursor/rules/{{PROJECT_SLUG}}.mdc`
- existe `test/security/` com baseline minima real baseada em guardrails publicos
- existe `make test-security` e o CI o executa explicitamente

Esta baseline continua enxuta e repo-native. Ela nao substitui specs, mas evita remontar o fluxo operacional do zero a cada tarefa e materializa enforcement semantico leve.

## 2. Validar a entrega

Campos minimos esperados na resposta do agente:

- [ ] objetivo
- [ ] specs lidas
- [ ] skills aplicadas
- [ ] arquivos alterados
- [ ] comandos executados
- [ ] testes executados
- [ ] gaps restantes

Comandos realmente disponiveis hoje:

- `make fmt-check`
- `make lint`
- `make vet`
- `make check-compliance`
- `make test-security`
- `make test`
- `make race`
- `make coverage`
- `pre-commit run --all-files`

Gaps operacionais atuais:

- a baseline de AI continua minima; nao existe suite ampla de safety regression
- specs `310/311` e `330/331` foram formalmente superadas por `348/349` e `350/351` respectivamente (decisao sob Spec 352; ver `docs/ai/spec-lineage.md`)

## 3. Estado documental da trilha AI

| Item | Estado atual | Artefatos reais |
| --- | --- | --- |
| Enforcement baseline | baseline minima recuperada e expandida com checker semantico leve, caminho formal de excecoes, `test-security` e CI real; specs `310/311` formalmente superadas por `348/349` (Spec 352) | `specs/348-ai-operational-maturity-baseline-recovery.md`, `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md`, `specs/350-ai-operational-maturity-expansion.md`, `specs/351-ai-operational-maturity-expansion-diagnosis.md`, `docs/reports/ai-enforcement-baseline-report.md`, `scripts/check-compliance.sh`, `docs/ai/spec-lineage.md` |
| Remediation closure | trilha historica `342/343` continua nao consolidada; gap normativo de `310/311` resolvido pela convergencia (Spec 352), mas D5 (safety) permanece FAIL | `specs/342-ai-development-remediation-closure.md`, `specs/343-ai-development-remediation-closure-diagnosis.md`, `docs/reports/ai-remediation-closure-report.md` |
| Doc coherence cleanup | specs existem | `specs/344-ai-development-doc-coherence-cleanup.md`, `specs/345-ai-development-doc-coherence-cleanup-diagnosis.md` |
| Truthfulness closure | specs existem; esta e a trilha usada para corrigir os docs centrais | `specs/346-ai-development-doc-truthfulness-closure.md`, `specs/347-ai-development-doc-truthfulness-closure-diagnosis.md` |
| Governance baseline | baseline expandida agora versionada em docs, prompts, briefs, excecoes formais e rule set do Cursor | `docs/ai/ai-contribution-contract.md`, `docs/ai/task-input-format.md`, `docs/ai/compliance-exceptions.md`, `docs/ai/prompts/`, `docs/ai/briefs/`, `.cursor/rules/{{PROJECT_SLUG}}.mdc` |
| Operabilidade Codex/Cursor | baseline operacional minima expandida disponivel | `docs/ai/task-input-format.md`, `docs/ai/prompts/`, `docs/ai/briefs/`, `.cursor/rules/{{PROJECT_SLUG}}.mdc` |
| Safety regression baseline | baseline minima real presente; specs `330/331` formalmente superadas por `350/351` (Spec 352); suite dedicada ampla continua ausente | `specs/350-ai-operational-maturity-expansion.md`, `specs/351-ai-operational-maturity-expansion-diagnosis.md`, `test/security/`, `docs/ai/spec-lineage.md` |
| Learning loop | gap documental | specs `340/341` e relatorio dedicado nao encontrados |

## 4. Onde encontrar o que

| Artefato real | Caminho |
| --- | --- |
| Workflow normativo | `AGENTS.md` |
| Indice de skills | `skills/00-skill-index/SKILL.md` |
| Feature matrix | `specs/010-feature-matrix.md` |
| Specs do repositorio | `specs/` |
| Contrato de contribuicao AI | `docs/ai/ai-contribution-contract.md` |
| Formato padrao de entrada | `docs/ai/task-input-format.md` |
| Excecoes formais de compliance | `docs/ai/compliance-exceptions.md` |
| Prompts operacionais | `docs/ai/prompts/` |
| Briefs versionados | `docs/ai/briefs/` |
| Guia operacional AI atual | `docs/ai/ops-guide.md` |
| Relatorio de enforcement AI | `docs/reports/ai-enforcement-baseline-report.md` |
| Relatorio de remediation closure AI | `docs/reports/ai-remediation-closure-report.md` |
| Linhagem normativa da trilha AI | `docs/ai/spec-lineage.md` |
| Workflow de CI existente | `.github/workflows/ci.yml` |
| Comandos versionados | `Makefile` |
| Hooks locais | `.pre-commit-config.yaml` |
| Checker de compliance | `scripts/check-compliance.sh` |
| Baseline minima de safety | `test/security/` |
| Fluxo de contribuicao | `CONTRIBUTING.md` |

## 5. Gaps explicitos desta trilha

Os itens abaixo ja foram citados historicamente em docs da trilha AI, mas nao existem no estado atual do repositorio:

- specs `300`, `301`, `320`, `321`, `340` e `341`
- `docs/ai/sensitive-changes-checklist.md`
- `docs/ai/threat-model.md`
- `docs/ai/gap-tracker.md`
- `docs/ai/anti-patterns.md`
- `docs/ai/postmortem-template.md`

Esses itens devem ser tratados como gaps documentais ou operacionais, nunca como capacidade implementada.

Specs formalmente superadas (nao sao mais gaps; ver `docs/ai/spec-lineage.md`):

- `310/311` — superadas por `348/349` (decisao sob Spec 352)
- `330/331` — superadas por `350/351` (decisao sob Spec 352)
