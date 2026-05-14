# Como o kit funciona

O kit tem duas partes: um `template/` parametrizado e um prompt master que embute o conteúdo desse template em blocos `<file path="...">`.

## Camadas geradas

1. `AGENTS.md` como constituição do repositório.
2. `skills/` como roteamento operacional para agentes.
3. `.cursor/rules/` e `.cursor/hooks/`.
4. `automation/` com `phase_autopilot` e `interactive_sdd_autopilot`.
5. `specs/` com specs base e phase 0 dual-spec.
6. `docs/ai/` com contrato, prompts e briefs.
7. `.github/`, `scripts/`, `Makefile` e pre-commit.
