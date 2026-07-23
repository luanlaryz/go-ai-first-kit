# Como o kit funciona

O kit tem três partes: a CLI `gakit`, um `template/` parametrizado e um prompt master que embute o conteúdo desse template em blocos `<file path="...">`.

## CLI gakit

`gakit` é a interface primária do projeto. Referência completa de comandos e flags em [cli-reference.md](cli-reference.md).

- `gakit create` usa o `template/` embarcado via `go:embed`, renderiza placeholders em paths e conteúdos, renomeia `go.mod.tmpl` para `go.mod`, aplica permissões executáveis e inicializa git.
- `gakit template list` inspeciona o template embarcado sem gerar nada.
- `gakit diagnose --path <dir>` cria um inventário do projeto alvo e avalia sete pilares: DX, AI-first, Security, Hexagonal Architecture, OpenAPI, Documentação e Governança.
- O diagnóstico calcula score por pilar com `0.6 * coverage + 0.4 * quality`, aplica pesos e pode persistir relatório Markdown e JSON.
- Os pilares AI-first, Governança e Documentação reconhecem a camada de backlog governado, o inventário de maturidade, o catálogo de capacidades, os ADRs em `docs/decisions/` e o changelog no formato Keep a Changelog.

## Template e prompt master

O `template/` é a fonte única do que um projeto gerado contém. O `prompt-bootstrap-go-ai-first.md` é gerado a partir dele por `scripts/build-master-prompt.sh` e validado por teste (`--check`): CLI e prompt produzem a mesma árvore.

## Camadas geradas

1. `AGENTS.md` como constituição do repositório.
2. `skills/` como roteamento operacional para agentes.
3. `.cursor/rules/` e `.cursor/hooks/`.
4. `automation/` com `phase_autopilot` e `interactive_sdd_autopilot`.
5. `specs/` com specs base e phase 0 dual-spec.
6. `docs/ai/` com contrato, inventário de maturidade, catálogo de capacidades, prompts e briefs.
7. `docs/README.md` (central de navegação) e `docs/journeys/` (jornadas humanas).
8. `.github/`, `scripts/`, `Makefile` e pre-commit.
9. `docs/backlog/Backlog.md` com `skills/26-backlog-item-intake/` para intake governado de backlog.
10. `docs/decisions/` (ADRs) e `docs/release-versioning-policy.md`, `docs/release-notes-policy.md` e `docs/release-checklist.md` para governança de release, versionamento e changelog.

## Os dois autopilots

Ambos são workflows governados operados por agentes seguindo runbooks versionados — não há daemon nem runner hospedado.

- `phase_autopilot` executa um roadmap fechado: `automation/ROADMAP.json` define as fases e gates, `automation/PHASE_STATE.json` guarda o estado, e o report de cada fase é a fonte de verdade para avanço.
- `interactive_sdd_autopilot` transforma um pedido de mudança (feature, bug, docs, refactor) em trilha SDD: intake, requisitos, decisão de spec, dual-spec, implementação, diagnóstico, report e gate, com estado em `automation/INTERACTIVE_STATE.json`.

O modo interativo é aditivo: nunca altera `ROADMAP.json` nem `PHASE_STATE.json`. Em ambos, testes verdes sem report não autorizam avanço.
