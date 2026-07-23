# go-ai-first-kit

`go-ai-first-kit` cria projetos Go com maturidade AI-first desde o primeiro commit. Extraído da governança do `govolt`, ele empacota `AGENTS.md`, skills, Cursor rules, autopilot dual (`phase_autopilot` + `interactive_sdd_autopilot`), SDD dual-spec, backlog governado, governança de release/changelog/ADR, docs operacionais, CI gates e scripts de compliance.

Pré-requisito: Go 1.26.4+.

## Uso rápido

```bash
go install ./cmd/gakit
gakit create ./myapp --slug myapp --title "My App" --module github.com/acme/myapp --description "AI-first app" --author "Acme"
gakit diagnose --path ./myapp --report-only
```

Após criar, abra `docs/README.md` dentro do projeto gerado: ele é a central de navegação, com catálogo de capacidades (`docs/ai/capabilities.md`) e jornadas humanas (`docs/journeys/`).

Alternativa sem CLI: cole o conteúdo de `prompt-bootstrap-go-ai-first.md` em um agente com permissão de escrita e informe os parâmetros solicitados. Fallback legado: `./scripts/init.sh /caminho/para/novo-projeto`.

## O que o kit executa e o que o agente opera

- A CLI executa: criação (`create`), inspeção (`template list`) e diagnóstico (`diagnose`, sete pilares com score ponderado).
- O projeto gerado executa: `make check-compliance`, lint, vet, testes, race, secrets scan e govulncheck.
- Agentes operam: os autopilots, o SDD dual-spec, o backlog e a governança de release. São workflows governados por runbooks e reports versionados — não há runner hospedado.
- O starter não contém capability de produto: `pkg/`, API, OpenAPI e `examples/` nascem apenas sob spec aprovada.

## Documentação

- [docs/README.md](docs/README.md): índice por objetivo.
- [docs/capabilities.md](docs/capabilities.md): catálogo de capacidades, tipos de verificação e limites.
- [docs/adoption-guide.md](docs/adoption-guide.md): sequência verificável de adoção.
- [docs/cli-reference.md](docs/cli-reference.md): comandos e flags da CLI.
- [docs/how-it-works.md](docs/how-it-works.md): arquitetura do kit.

## Trade-off

Este kit é intencionalmente rígido: prioriza rastreabilidade, dual-spec, diagnóstico e repetibilidade sobre velocidade informal. Evite-o para spikes descartáveis ou projetos que não desejam enforcement em CI.
