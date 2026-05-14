# go-ai-first-kit

`go-ai-first-kit` é um kit para criar projetos Go com maturidade AI-first desde o primeiro commit. Ele foi extraído da infraestrutura de governança do `govolt` e empacota skills, Cursor rules, `AGENTS.md`, autopilot dual, SDD dual-spec, docs operacionais, CI gates, prompts e scripts de compliance.

## O que ele entrega

- `prompt-bootstrap-go-ai-first.md` auto-contido para colar em Cursor, Claude Code, Codex CLI ou outro agente.
- `template/` completo com infraestrutura parametrizada por `{{PROJECT_SLUG}}`, `{{MODULE_PATH}}` e demais placeholders.
- `scripts/init.sh` para renderizar o template em um diretório de projeto novo.
- `scripts/build-master-prompt.sh` para regenerar o prompt único a partir do template.

## Uso rápido: prompt único

1. Abra `prompt-bootstrap-go-ai-first.md`.
2. Cole o conteúdo em um agente com permissão para escrever arquivos.
3. Informe os parâmetros solicitados pelo prompt.
4. Deixe o agente criar o projeto, rodar validações e reportar evidência.

## Uso via script

```bash
./scripts/init.sh /caminho/para/novo-projeto
```

## Trade-off

Este kit é intencionalmente rígido: ele prioriza rastreabilidade, dual-spec, diagnóstico e repetibilidade sobre velocidade informal.
