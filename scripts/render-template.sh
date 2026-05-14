#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 2 ]]; then
  echo "usage: scripts/render-template.sh <template-dir> <target-dir>" >&2
  exit 1
fi

TEMPLATE_DIR="$1"
TARGET_DIR="$2"

read -r -p "Project slug (Go package safe, e.g. myapp): " PROJECT_SLUG
read -r -p "Project title (e.g. My App): " PROJECT_TITLE
read -r -p "Go module path (e.g. github.com/acme/myapp): " MODULE_PATH
read -r -p "Project description: " PROJECT_DESCRIPTION
read -r -p "Author name: " AUTHOR_NAME
read -r -p "License name [MIT]: " LICENSE_NAME
LICENSE_NAME="${LICENSE_NAME:-MIT}"
read -r -p "Upstream/reference name [none]: " UPSTREAM_NAME
UPSTREAM_NAME="${UPSTREAM_NAME:-none}"

export PROJECT_SLUG PROJECT_TITLE MODULE_PATH PROJECT_DESCRIPTION AUTHOR_NAME LICENSE_NAME UPSTREAM_NAME
export UPSTREAM_NAME_LOWER="$(printf '%s' "$UPSTREAM_NAME" | tr '[:upper:]' '[:lower:]')"
export UPSTREAM_OPS_NAME="${UPSTREAM_NAME}Ops"
export DOMAIN_ACTOR="domain actor"
export DOMAIN_ACTOR_TITLE="Domain Actor"
export EXTERNAL_SYSTEM="External System"
export EXTERNAL_SYSTEM_LOWER="external system"
export DOMAIN_ENTITY_SET="domain entities"

python3 - "$TEMPLATE_DIR" "$TARGET_DIR" <<'PY'
from pathlib import Path
import os
import sys

src = Path(sys.argv[1])
dst = Path(sys.argv[2])
keys = [
    "PROJECT_SLUG",
    "PROJECT_TITLE",
    "MODULE_PATH",
    "PROJECT_DESCRIPTION",
    "AUTHOR_NAME",
    "LICENSE_NAME",
    "UPSTREAM_NAME",
    "UPSTREAM_NAME_LOWER",
    "UPSTREAM_OPS_NAME",
    "DOMAIN_ACTOR",
    "DOMAIN_ACTOR_TITLE",
    "EXTERNAL_SYSTEM",
    "EXTERNAL_SYSTEM_LOWER",
    "DOMAIN_ENTITY_SET",
]
repls = {key: os.environ[key] for key in keys}


def render(value: str) -> str:
    for key, replacement in repls.items():
        value = value.replace("{{" + key + "}}", replacement)
    return value


for path in sorted(src.rglob("*")):
    rel = path.relative_to(src)
    out = dst / Path(*[render(part) for part in rel.parts])
    if path.is_dir():
        out.mkdir(parents=True, exist_ok=True)
        continue

    out.parent.mkdir(parents=True, exist_ok=True)
    data = path.read_bytes()
    try:
        text = data.decode("utf-8")
    except UnicodeDecodeError:
        out.write_bytes(data)
    else:
        out.write_text(render(text), encoding="utf-8")
    out.chmod(path.stat().st_mode & 0o777)
PY
