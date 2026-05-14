#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

python3 - "$ROOT_DIR" <<'PY'
from pathlib import Path
import sys

root = Path(sys.argv[1])
template = root / "template"
out = root / "prompt-bootstrap-go-ai-first.md"

header = """# Prompt Bootstrap Go AI-First

Você é um agente de engenharia criando um projeto Go novo com baseline AI-first completa.

## Parâmetros Obrigatórios

Antes de escrever arquivos, colete:

- `PROJECT_SLUG`: slug do projeto e pacote raiz Go seguro, por exemplo `myapp`.
- `PROJECT_TITLE`: nome humano do projeto, por exemplo `My App`.
- `MODULE_PATH`: módulo Go, por exemplo `github.com/acme/myapp`.
- `PROJECT_DESCRIPTION`: descrição objetiva do projeto.
- `AUTHOR_NAME`: autor/mantenedor inicial.
- `LICENSE_NAME`: licença, default `MIT`.
- `UPSTREAM_NAME`: referência/upstream opcional, default `none`.

## Regras Obrigatórias

1. Leia este prompt inteiro antes de criar arquivos.
2. Crie a árvore exatamente conforme os blocos `<file path="...">`.
3. Substitua placeholders `{{...}}` em paths e conteúdos.
4. Aplique permissão executável em scripts `.sh` e `.cursor/hooks/*.sh`.
5. Se `go.mod.tmpl` existir, renomeie para `go.mod` após renderizar.
6. Rode `make check-compliance`, `make fmt-check`, `make vet` e `make test` quando possível.
7. Responda com parâmetros usados, arquivos criados, comandos executados, validações passadas e gaps restantes.

## Arquivos

"""

parts = [header]
for path in sorted(template.rglob("*")):
    if path.is_dir():
        continue
    try:
        text = path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        continue
    rel = path.relative_to(template).as_posix()
    parts.append(f'<file path="{rel}">\n')
    parts.append(text)
    if not text.endswith("\n"):
        parts.append("\n")
    parts.append("</file>\n\n")

out.write_text("".join(parts), encoding="utf-8")
print(f"wrote {out}")
PY
