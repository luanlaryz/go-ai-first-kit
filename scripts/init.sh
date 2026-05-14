#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

if [[ $# -ne 1 ]]; then
  echo "usage: scripts/init.sh <target-dir>" >&2
  exit 1
fi

TARGET_DIR="$1"
mkdir -p "$TARGET_DIR"

"$ROOT_DIR/scripts/render-template.sh" "$ROOT_DIR/template" "$TARGET_DIR"

if [[ -f "$TARGET_DIR/go.mod.tmpl" && ! -f "$TARGET_DIR/go.mod" ]]; then
  mv "$TARGET_DIR/go.mod.tmpl" "$TARGET_DIR/go.mod"
fi

if [[ ! -d "$TARGET_DIR/.git" ]]; then
  git -C "$TARGET_DIR" init -b main
fi

chmod +x "$TARGET_DIR"/scripts/*.sh 2>/dev/null || true
chmod +x "$TARGET_DIR"/.cursor/hooks/*.sh 2>/dev/null || true

echo "AI-first Go project rendered at $TARGET_DIR"
