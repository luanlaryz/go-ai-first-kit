---
name: 29-governed-change-workflow
description: Fluxo mestre obrigatorio para mudancas em escopo governado (pkg/**, internal/**, api/**, migrations/**). Use ao iniciar qualquer feature, fix, refactor ou mudanca de contrato/dados, garantindo BLG-NNNN + par dual-spec + plano revisado antes do PR.
---

# Governed Change Workflow

Goal: garantir que nenhuma mudanca em escopo governado seja desenvolvida ou submetida sem o trio obrigatorio (item de backlog + par dual-spec + plano revisado), declarado no corpo do PR e validado por gate executavel.

## When to Use

- Antes de editar `pkg/**`, `internal/**`, `api/**` ou `migrations/**`.
- Ao receber pedido de feature, fix, refactor ou mudanca de contrato/dados que toque esses caminhos.

## Nao use

- Para mudanca apenas em docs puras, tooling de governanca (`scripts/`, `.cursor/`, `.github/`, `Makefile`) ou specs isoladas fora do escopo governado.

## Required Reading

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `.cursor/rules/governed-change-enforcement.mdc`
4. `skills/26-backlog-item-intake/SKILL.md`
5. `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`
6. `skills/27-dual-model-plan-review/SKILL.md` e `skills/28-plan-review-autopilot/SKILL.md`

## Non-Negotiables

- Escopo governado exige o trio; sem qualquer parte dele, pare e registre o gap (nao avance por narrativa).
- Backlog canonico: `BLG-NNNN` em `docs/backlog/Backlog.md`, com `Status` em `ready_for_implementation`, `in_progress` ou `done` no momento do PR.
- Par dual-spec: spec de construcao `NNN` + spec de diagnostico `NNN+1` terminada em `-diagnosis.md`; decisao amendar vs nova trilha via `skills/05-{{PROJECT_SLUG}}-spec-architect`.
- Plano em `.cursor/plans/*.plan.md` com review `APPROVED`/`APPROVED_WITH_CHANGES` (`skills/27`/`28`); `make verify-plan-reviews` PASS.
- Declarar `## Backlog`, `## Specs lidas`, `## Autopilot` e `## Regressao` no corpo do PR.
- Nao contornar `scripts/check-governed-change.sh`; excecoes apenas via `.guardrails/governed-change-exceptions.yaml` com `expires_at`.

## Workflow

1. Intake: registre ou atualize o item `BLG-NNNN` (`skills/26-backlog-item-intake`).
2. Spec architect: crie ou estenda o par dual-spec (`skills/05-{{PROJECT_SLUG}}-spec-architect`).
3. Autopilot: registre a trilha em `automation/INTERACTIVE_STATE.json` (`interactive_sdd_autopilot`, `automation/INTERACTIVE_RUNBOOK.md`) como estado EFEMERO/LOCAL; nao versione essa mudanca no PR, porque estado operacional transitorio gera conflito entre branches e sessoes. A prova de autopilot exigida pelo gate e o plano revisado, nao o state.
4. Plano: crie `.cursor/plans/<id>.plan.md` e passe pelo gate de review (`skills/27`/`28`; `make verify-plan-reviews`).
5. Implemente conforme as specs; preserve propagacao de `context.Context` e dos identificadores de correlacao.
6. Preencha o corpo do PR (`## Backlog`, `## Specs lidas`, `## Autopilot`, `## Regressao`) e rode `make pre-pr PR_BODY=<arquivo>`.
7. Classifique `PASS/PARTIAL/FAIL/BLOCKED` e registre gaps restantes.

## Reserva de identificadores em sessoes paralelas

`BLG-NNNN`, `specs/NNN` e `migrations/NNNN` sao sequenciais e por isso colidem entre sessoes que trabalham em paralelo. Reserve-os no ultimo momento antes do commit e aplique `skills/31-parallel-session-coordination/SKILL.md`; o gate executavel e `scripts/check-parallel-collision.sh`.

## Definition of Done

- `scripts/check-governed-change.sh` PASS para o PR.
- `BLG-NNNN`, par dual-spec e plano revisado coerentes e cruzados.
- `make pre-pr` PASS; entrega lista skills aplicadas, arquivos alterados, comandos/testes, veredito do plano e gaps.
