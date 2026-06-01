# Changelog

## [Unreleased]

### Added
- CLI `gakit` com comandos `create`, `diagnose` e `version`.
- Diagnóstico ponderado por pilares com saída terminal, Markdown e JSON opcional.
- Bootstrap inicial do `go-ai-first-kit`, extraído da governança AI-first do `govolt`.
- Camada de backlog governado no `template/`: `skills/26-backlog-item-intake/` (SKILL, agents e referências) e `docs/backlog/Backlog.md` canônico.
- Governança de release/changelog/ADR no `template/`: `docs/decisions/` (sistema de ADRs + seed `0001`), `docs/release-versioning-policy.md`, `docs/release-notes-policy.md`, `docs/release-checklist.md` e `CHANGELOG.md` no formato Keep a Changelog + SemVer.
- `template/docs/plans/.gitkeep` e `template/docs/reports/.gitkeep` para artefatos versionados de plano e report.
- Diagnóstico do diff `govolt → kit` em `docs/govolt-sync-diagnosis.md`.

### Changed
- `gakit diagnose` passa a pontuar a camada de backlog (pilar AI-first), os ADRs e a política de release (pilar Governança) e o changelog Keep a Changelog (pilar Documentação).
- `template/AGENTS.md` e `template/skills/00-skill-index/SKILL.md` religados ao backlog, ADRs e governança de release.
- `template/.editorconfig` alinhado ao `govolt` (regras para Go, Makefile, YAML/JSON e shell).
