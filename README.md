# go-ai-first-kit

`go-ai-first-kit` é um kit para criar projetos Go com maturidade AI-first desde o primeiro commit. Ele foi extraído da infraestrutura de governança do `govolt` e empacota skills, Cursor rules, `AGENTS.md`, autopilot dual, SDD dual-spec, docs operacionais, CI gates, prompts e scripts de compliance.

## O que ele entrega

- CLI `gakit` para criar projetos e diagnosticar maturidade AI-first.
- `prompt-bootstrap-go-ai-first.md` auto-contido para colar em Cursor, Claude Code, Codex CLI ou outro agente.
- `template/` completo com infraestrutura parametrizada por `{{PROJECT_SLUG}}`, `{{MODULE_PATH}}` e demais placeholders.
- Scripts shell legados como fallback para renderizar template e regenerar prompt único.

## Fluxo recomendado: gakit

```bash
go install ./cmd/gakit
gakit create ./myapp --slug myapp --title "My App" --module github.com/acme/myapp --description "AI-first app" --author "Acme"
gakit diagnose --path ./myapp --report-only --json
gakit diagnose --path ./myapp --plan-prompt --plan-prompt-only --out ./reports
```

Comandos principais:

- `gakit help` mostra os comandos e flags disponíveis.
- `gakit create` renderiza o `template/` embarcado, substitui placeholders, renomeia `go.mod.tmpl` para `go.mod`, aplica permissões executáveis e inicializa git.
- `gakit diagnose --path <dir>` avalia DX, AI-first, security, arquitetura hexagonal, OpenAPI, documentação e governança com score ponderado.
- `gakit diagnose --path <dir> --plan-prompt` imprime um prompt opcional para criar plano de correção SDD dual spec a partir dos achados; use `--plan-prompt-only --out ./reports` para persistir o prompt em Markdown sem perguntar.

## Uso rápido: prompt único

1. Abra `prompt-bootstrap-go-ai-first.md`.
2. Cole o conteúdo em um agente com permissão para escrever arquivos.
3. Informe os parâmetros solicitados pelo prompt.
4. Deixe o agente criar o projeto, rodar validações e reportar evidência.

## Fallback legado

```bash
./scripts/init.sh /caminho/para/novo-projeto
```

Os scripts `scripts/init.sh`, `scripts/render-template.sh` e `scripts/build-master-prompt.sh` continuam funcionais para compatibilidade, mas o fluxo primário agora é a CLI `gakit`.

## Trade-off

Este kit é intencionalmente rígido: ele prioriza rastreabilidade, dual-spec, diagnóstico e repetibilidade sobre velocidade informal.
