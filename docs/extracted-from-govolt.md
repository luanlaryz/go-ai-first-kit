# Extraído do govolt

Este kit foi montado a partir das camadas de governança observadas no `govolt`: `AGENTS.md`, Cursor rules, hooks, skills, autopilot fixo, autopilot interativo, stop conditions, specs, docs/ai e CI/scripts de compliance.

## Sincronizações posteriores

Camadas trazidas em sincronizações após o bootstrap, parametrizadas e sem conteúdo de domínio do `govolt`:

- Backlog governado: `skills/26-backlog-item-intake/` e `docs/backlog/Backlog.md` (intake, triagem, classificação SDD e prompts autopilot).
- Governança de release/changelog/ADR: `docs/decisions/` (sistema de ADRs MADR-like), `docs/release-versioning-policy.md`, `docs/release-notes-policy.md`, `docs/release-checklist.md` e `CHANGELOG.md` no formato Keep a Changelog + SemVer.

Divergências intencionais: a camada de GitHub Pages do `govolt` (`.github/workflows/pages.yml`, `scripts/check-pages-site.sh`, `make check-pages-site`) é específica daquele repositório e não entra no kit.

O diagnóstico completo do diff `govolt → kit`, com matriz de paridade e tabela de decisão por item, está em [`docs/govolt-sync-diagnosis.md`](govolt-sync-diagnosis.md).
