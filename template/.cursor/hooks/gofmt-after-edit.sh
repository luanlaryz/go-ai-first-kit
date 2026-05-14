#!/bin/bash
input=$(cat)
filepath=$(echo "$input" | jq -r '.path // empty')

if [[ -n "$filepath" && "$filepath" == *.go ]]; then
  gofmt -w "$filepath" 2>/dev/null
fi

echo '{}'
exit 0
