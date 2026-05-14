# Brief real - Spec 348 baseline recovery

## Objetivo

Recuperar a baseline minima de maturidade operacional da trilha de AI development da `{{PROJECT_SLUG}}`, trocando marcadores vazios e gaps abertos por mecanismos versionados, executaveis e auditaveis.

## Specs governantes

- `specs/348-ai-operational-maturity-baseline-recovery.md`
- `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md`
- `AGENTS.md`
- `specs/020-repository-architecture.md`
- `specs/060-guardrails.md`

## Escopo desta entrega

- criar assets operacionais minimos em `docs/ai/`
- versionar regra do Cursor, PR template e config de `pre-commit`
- adicionar checker simples e auditavel
- substituir targets vazios no `Makefile`
- alinhar `.github/workflows/ci.yml` aos targets reais
- recuperar testes basicos em `pkg/types` e `pkg/guardrail`
- atualizar apenas os docs centrais estritamente necessarios

## Fora do escopo

- suite completa de safety regression
- policy engine avancado
- expansao ampla de prompts e briefs
- reabertura de toda a trilha historica 300-347

## Evidencias esperadas

- `make lint`, `make vet` e `make check-compliance` executaveis
- `.pre-commit-config.yaml` usando hooks locais simples
- prompts e briefs reutilizaveis, sem texto de enchimento
- `_test.go` presentes em `pkg/types` e `pkg/guardrail`
- docs centrais alinhados ao novo estado real do repo
