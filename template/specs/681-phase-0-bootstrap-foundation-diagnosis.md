# Spec 681: Phase 0 - Bootstrap Foundation Diagnosis

## 1. Objetivo

Definir a auditoria da Phase 0 de `{{PROJECT_SLUG}}`, confirmando que a baseline AI-first foi criada com evidência real.

## 2. Pergunta principal

A baseline AI-first está pronta para operar agentes de desenvolvimento Go com specs, skills, automation, reports e CI?

## 3. Escopo

Verificar arquivos obrigatórios, JSON de automation, scripts executáveis, docs/ai e prompts.

## 4. Checklist obrigatório

- [ ] `AGENTS.md` existe.
- [ ] `skills/00-skill-index/SKILL.md` existe.
- [ ] `automation/AUTOPILOT.md` existe.
- [ ] `automation/INTERACTIVE_AUTOPILOT.md` existe.
- [ ] `.cursor/rules/` existe.
- [ ] `.github/PULL_REQUEST_TEMPLATE.md` existe.
- [ ] `scripts/check-compliance.sh` existe e é executável.

## 5. Classificação

- `PASS`: todos os itens obrigatórios atendidos e checks verdes.
- `PARTIAL`: baseline existe mas algum eixo material está incompleto.
- `FAIL`: baseline não opera.
- `BLOCKED`: falta decisão humana ou dependência externa.

## 6. Saída obrigatória

`docs/reports/phase-0-bootstrap-foundation-report.md`

## 7. Critérios de aceitação

O report deve concluir `READY FOR FIRST FEATURE TRAIL` ou `NOT READY FOR FIRST FEATURE TRAIL`.
