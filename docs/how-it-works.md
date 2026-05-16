# Como o kit funciona

O kit tem três partes: a CLI `gakit`, um `template/` parametrizado e um prompt master que embute o conteúdo desse template em blocos `<file path="...">`.

## CLI gakit

`gakit` é a interface primária do projeto.

- `gakit create` usa o `template/` embarcado via `go:embed`, renderiza placeholders em paths e conteúdos, renomeia `go.mod.tmpl` para `go.mod`, aplica permissões executáveis e inicializa git.
- `gakit diagnose --path <dir>` cria um inventário do projeto alvo e avalia sete pilares: DX, AI-first, Security, Hexagonal Architecture, OpenAPI, Documentação e Governança.
- O diagnóstico calcula score por pilar com `0.6 * coverage + 0.4 * quality`, aplica pesos e pode persistir relatório Markdown e JSON.

## Camadas geradas

1. `AGENTS.md` como constituição do repositório.
2. `skills/` como roteamento operacional para agentes.
3. `.cursor/rules/` e `.cursor/hooks/`.
4. `automation/` com `phase_autopilot` e `interactive_sdd_autopilot`.
5. `specs/` com specs base e phase 0 dual-spec.
6. `docs/ai/` com contrato, prompts e briefs.
7. `.github/`, `scripts/`, `Makefile` e pre-commit.
