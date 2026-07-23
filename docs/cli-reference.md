# Referência da CLI gakit

Comandos e flags efetivamente suportados. Fonte de verdade: `internal/cli/`.

## Flags globais

- `--no-color`: desabilita cores ANSI na saída.
- `--verbose`: mostra detalhes adicionais (em `create`, imprime os próximos comandos sugeridos).

`gakit` sem argumentos mostra uma dica rápida; `gakit help` mostra todos os comandos.

## gakit create

Renderiza o template embarcado em um diretório de destino.

```bash
gakit create ./myapp \
  --slug myapp \
  --title "My App" \
  --module github.com/acme/myapp \
  --description "AI-first app" \
  --author "Acme"
```

Flags:

- `[target]` ou `--target`: diretório de destino (obrigatório em `--non-interactive`).
- `--template`: template embarcado a usar (hoje apenas `default`).
- `--slug`: slug Go-safe do projeto (`{{PROJECT_SLUG}}`).
- `--title`: nome humano (`{{PROJECT_TITLE}}`).
- `--module`: módulo Go escrito no `go.mod` (`{{MODULE_PATH}}`).
- `--description`: descrição objetiva (`{{PROJECT_DESCRIPTION}}`).
- `--author`: autor/mantenedor inicial.
- `--license`: licença (default `MIT`).
- `--upstream`: referência conceitual de upstream (default `none`; com `none`, as regras de paridade upstream do projeto gerado ficam explicitamente não aplicáveis).
- `--non-interactive`: falha se algum parâmetro obrigatório estiver ausente, em vez de perguntar.
- `--force`: permite renderizar em diretório não vazio.

Sem `--non-interactive`, parâmetros ausentes são coletados interativamente. Ao final o comando renomeia `go.mod.tmpl` para `go.mod`, aplica permissão executável em scripts e inicializa git.

Jornada relacionada: [adoption-guide.md](adoption-guide.md).

## gakit template list

Inspeciona o conteúdo do template embarcado antes de gerar.

```bash
gakit template list
gakit template list --tree
gakit template list --json
gakit template list --files-only
gakit template list --dirs-only
```

Flags:

- `--template`: template a inspecionar (default `default`).
- `--tree`: imprime como árvore.
- `--json`: imprime o inventário em JSON indentado.
- `--files-only` / `--dirs-only`: filtram a listagem (mutuamente exclusivos).

## gakit diagnose

Avalia a maturidade AI-first de um projeto em sete pilares (pesos em [capabilities.md](capabilities.md)).

```bash
gakit diagnose --path .
gakit diagnose --path . --min-score 80 --report-only --json --out ./reports
gakit diagnose --path . --plan-prompt --plan-prompt-only --out ./reports
```

Flags:

- `--path` (obrigatório): diretório do projeto a diagnosticar.
- `--min-score`: score global mínimo; abaixo dele o comando retorna exit code 1 — útil como gate em CI.
- `--out`: diretório para salvar relatórios (default: diretório atual).
- `--report-only`: salva o relatório sem perguntar ao final.
- `--json`: salva também a versão JSON do relatório.
- `--plan-prompt`: imprime um prompt para criar um plano de correção SDD dual-spec a partir dos achados.
- `--plan-prompt-only`: salva o prompt de plano em Markdown sem perguntar.

Interpretação: score por pilar é `0.6 * coverage + 0.4 * quality`. Em um starter recém-gerado, os pilares Hexagonal e OpenAPI reportam ausências por design; ver [adoption-guide.md](adoption-guide.md).

## gakit version

Imprime versão, commit e data de build.
