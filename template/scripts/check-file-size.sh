#!/usr/bin/env bash
set -euo pipefail

WARN_THRESHOLD=${WARN_THRESHOLD:-700}
ERROR_THRESHOLD=${ERROR_THRESHOLD:-1000}
EXCEPTIONS_FILE="${EXCEPTIONS_FILE:-.file-size-exceptions}"

exceptions=()
if [[ -f "$EXCEPTIONS_FILE" ]]; then
  while IFS= read -r line; do
    line="${line%%#*}"
    line="$(echo "$line" | xargs)"
    [[ -n "$line" ]] && exceptions+=("$line")
  done < "$EXCEPTIONS_FILE"
fi

is_excepted() {
  local file="$1"
  for exc in "${exceptions[@]+"${exceptions[@]}"}"; do
    [[ "$file" == "$exc" ]] && return 0
  done
  return 1
}

warnings=0
errors=0

while IFS= read -r file; do
  lines=$(wc -l < "$file" | tr -d ' ')

  if is_excepted "$file"; then
    continue
  fi

  if (( lines > ERROR_THRESHOLD )); then
    echo "ERROR: $file has $lines lines (threshold: $ERROR_THRESHOLD)"
    errors=$((errors + 1))
  elif (( lines > WARN_THRESHOLD )); then
    echo "WARNING: $file has $lines lines (threshold: $WARN_THRESHOLD)"
    warnings=$((warnings + 1))
  fi
done < <(find . -name '*.go' -not -name '*_test.go' -not -path './.git/*' -not -path './vendor/*' | sort)

if (( errors > 0 )); then
  echo ""
  echo "FAIL: $errors file(s) exceed $ERROR_THRESHOLD lines."
  echo "Add justified exceptions to $EXCEPTIONS_FILE or refactor per specs/020 section 8."
  exit 1
fi

if (( warnings > 0 )); then
  echo ""
  echo "INFO: $warnings file(s) exceed $WARN_THRESHOLD lines (review recommended)."
fi

exit 0
