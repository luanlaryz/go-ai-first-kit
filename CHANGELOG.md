# Changelog

## [Unreleased]

### Added
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
