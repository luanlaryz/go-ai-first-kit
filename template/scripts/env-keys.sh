#!/usr/bin/env bash
# Lists which variables exist in local env files WITHOUT exposing their values.
#
# Counterpart of the env read guard (.cursor/hooks/env-read-guard.py): an agent
# that needs to know whether a key is configured should use this instead of
# asking for approval to read the file.
set -uo pipefail

shopt -s nullglob 2>/dev/null || true

found=0
for file in .env .env.* */.env */.env.*; do
  case "$file" in
    *.example|*.sample|*.template) continue ;;
  esac
  [ -f "$file" ] || continue
  found=1
  printf '%s\n' "$file"
  # Print key names only. Values never reach stdout.
  awk -F= '
    /^[[:space:]]*#/ { next }
    /^[[:space:]]*$/ { next }
    /=/ {
      key = $1
      sub(/^[[:space:]]*(export[[:space:]]+)?/, "", key)
      gsub(/[[:space:]]+$/, "", key)
      if (key != "") printf "  %s\n", key
    }
  ' "$file"
done

if [ "$found" -eq 0 ]; then
  echo "env-keys: nenhum arquivo .env local encontrado"
fi
