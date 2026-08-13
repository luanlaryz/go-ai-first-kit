# Changelog

## [Unreleased]

### Added
- Segurança de acesso do agente no `template/`: hooks `agent-cloud-access-guard` (classifica `aws`/`kubectl` por alvo — conta, profile, contexto, endpoint — em vez de por verbo) e `env-read-guard` (aprovação humana antes de ler `.env`), ambos fail-closed e registrados em `.cursor/hooks.json`; gate de wiring `scripts/check-agent-hooks.sh`, `scripts/env-keys.sh`, `docs/runbooks/agent-cloud-access.md` e rule `agent-cloud-access.mdc`. Marcadores de ambiente ficam em `.cursor/hooks/cloud-access-config.json` e listas vazias são restritivas por design.
- Governança de mudança e de plano no `template/`: `scripts/check-governed-change.sh` (escopo `pkg/**`, `internal/**`, `api/**`, `migrations/**` exige item `BLG-NNNN`, par dual-spec, plano revisado e cobertura regressiva declarada no PR), `scripts/verify-plan-review.sh` (veredito validado por hash do corpo do plano e TTL), `scripts/review-plan.py`, `scripts/check-sdd-compliance.sh`, `automation/PLAN_REVIEWS/`, hook `governed-change-guard` e skills `27-dual-model-plan-review`, `28-plan-review-autopilot`, `29-governed-change-workflow`.
- Coordenação entre sessões paralelas no `template/`: skill `31-parallel-session-coordination` (Protocolos A–F), rule `parallel-session-safety.mdc`, hook `parallel-collision-guard` e `scripts/check-parallel-collision.sh` com selftest hermético — IDs sequenciais são reservados no último momento e cruzados contra a base, PRs abertos e refs remotas.
- Gates estruturais de design em `template/tools/guardrails/` (Go, apenas stdlib): fronteiras arquiteturais incluindo a regra de `pkg/` não depender de `internal/` fora da ponte pública, tamanho de função, tamanho de port, erro descartado em caminho de auditoria/persistência/IO e segurança de rota pública. Allowlists em `.guardrails/*.yaml` com `expires_at` obrigatório: entrada expirada reprova o gate.
- Skill `30-third-party-service-integrations` e rules de qualidade no `template/`: `architecture-boundaries`, `go-design-size`, `observability`, `package-clean`, `http-security-tenant`, `ci-no-silent-skip`, `test-suite-session-isolation`, `regression-coverage-required`, `docs-sync-maintenance`, `sdd-non-negotiable`, `pre-pr-before-push`, `pr-body-gate`.
- Alvos agregados `make guardrails` e `make pre-pr` no `template/`, mais o job `governance` no CI, e templates de plano em `template/.cursor/plans/templates/`.
- Diagnóstico do diff `sac-agents → kit` em `docs/sac-agents-sync-diagnosis.md` e pin da rodada em `docs/sac-agents-sync-baseline.json` — primeira sincronização no sentido downstream → kit.
- Superfície de descoberta documental do kit: `docs/README.md` (índice por objetivo), `docs/capabilities.md` (catálogo de capacidades com tipos de verificação e limites) e `docs/cli-reference.md` (comandos e flags da CLI). É documentação de navegação, não nova capability de produto.
- Central de documentação nos projetos gerados: `template/docs/README.md` (hub por papel), `template/docs/ai/capabilities.md` (catálogo operacional) e `template/docs/journeys/` (jornadas humanas de baseline, roadmap, backlog e release).
- Testes de integridade documental do template: links Markdown locais, referências `specs/*.md` em qualquer documento e referências textuais `Spec NNN` precisam apontar para artefatos renderizados.
- CLI `gakit` com comandos `create`, `diagnose` e `version`.
- Diagnóstico ponderado por pilares com saída terminal, Markdown e JSON opcional.
- Bootstrap inicial do `go-ai-first-kit`, extraído da governança AI-first do `govolt`.
- Camada de backlog governado no `template/`: `skills/26-backlog-item-intake/` (SKILL, agents e referências) e `docs/backlog/Backlog.md` canônico.
- Governança de release/changelog/ADR no `template/`: `docs/decisions/` (sistema de ADRs + seed `0001`), `docs/release-versioning-policy.md`, `docs/release-notes-policy.md`, `docs/release-checklist.md` e `CHANGELOG.md` no formato Keep a Changelog + SemVer.
- `template/docs/plans/.gitkeep` e `template/docs/reports/.gitkeep` para artefatos versionados de plano e report.
- Diagnóstico do diff `govolt → kit` em `docs/govolt-sync-diagnosis.md`.
- Inventário parametrizado de maturidade AI em `template/docs/ai/maturity-inventory.md`.
- Pin de sincronização upstream em `docs/govolt-sync-baseline.json`.

### Changed
- Endurecimentos sobre o artefato de origem, aplicados após revisão do port: as seções do corpo do PR passam a ser escopadas de verdade (sem fallback para o corpo inteiro, que permitia satisfazer `## Backlog` com um id citado em `## Regressao`); `scripts/verify-plan-review.sh` recusa planner igual a reviewer e `OPERATOR_FALLBACK` em plano de alto risco, fechando o caminho do frontmatter escrito à mão; a classificação de risco vive em `scripts/lib/plan_risk.py`, compartilhada entre escritor e verificador e com fronteira de palavra (antes `auth` casava com "author"); o allowlist de mudança governada exige o schema completo; o gate de colisão extrai spec e migration de nomes de branch remotos; `## Regressao` exige `test/<pacote>::<cenário>` concreto; e `check-observability.sh` ignora dependência `// indirect`.
- `gakit diagnose` passa a exigir os guards de credencial, o gate de wiring e o runbook no pilar Security, e a pontuar o gate de mudança governada, o gate de review de plano e as allowlists com expiração no pilar Governança.
- `template/.github/PULL_REQUEST_TEMPLATE.md` ganha as seções `## Backlog`, `## Autopilot` e `## Regressao` que o gate de mudança governada valida; `scripts/check-pr-body.sh` e `scripts/check-compliance.sh` passam a exigi-las.
- CI do `template/` reage a `pull_request: edited`, porque o gate de mudança governada lê o corpo do PR e editar o corpo precisa reexecutar o gate.
- `gakit diagnose` passa a reconhecer o catálogo de capacidades (`docs/ai/capabilities.md`) como sinal de qualidade do pilar AI-first, sem torná-lo gate de compliance.
- `template/docs/feature-request-lifecycle.md` reescrito como guia humano de acompanhamento e aceite; removidas referências a fases, specs e reports históricos não renderizados (Fase 26, Specs 348/350/700/701/726).
- `template/AGENTS.md` alinhado à `specs/000-project-mission.md`: com `--upstream none` (default), as regras de paridade upstream ficam explicitamente não aplicáveis.
- `gakit diagnose` passa a pontuar a camada de backlog (pilar AI-first), os ADRs e a política de release (pilar Governança) e o changelog Keep a Changelog (pilar Documentação).
- `template/AGENTS.md` e `template/skills/00-skill-index/SKILL.md` religados ao backlog, ADRs e governança de release.
- `template/.editorconfig` alinhado ao `govolt` (regras para Go, Makefile, YAML/JSON e shell).
- `gakit diagnose` passa a reconhecer o inventário de maturidade como sinal de qualidade do pilar AI-first, sem torná-lo gate de compliance.
- Template e kit passam a exigir Go 1.26.4.
- Documentação, automations e skills do template deixam de exigir specs e reports que o bootstrap não cria.
- O lint do kit passa a analisar apenas pacotes compiláveis, sem tentar analisar o template ainda não renderizado.
