#!/usr/bin/env bash
# Selftest de scripts/check-parallel-collision.sh.
#
# Hermetico: monta repositorios temporarios e usa BASE_REF para apontar a uma
# branch local, entao a logica central e exercitada sem rede. As fontes que
# dependem de rede (PRs abertos, refs remotas) degradam para WARN e sao cobertas
# pelo uso real no hook e no pre-pr.
set -uo pipefail

SCRIPT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/check-parallel-collision.sh"
[ -x "$SCRIPT" ] || { echo "selftest: FAIL: $SCRIPT nao executavel" >&2; exit 1; }

pass=0
fail=0
tmproot="$(mktemp -d)"
trap 'rm -rf "$tmproot"' EXIT

# make_repo <dir>: repo com BLG-0100 no backlog e specs/100 na branch `base`.
make_repo() {
    local dir="$1"
    mkdir -p "$dir/docs/backlog" "$dir/specs" "$dir/migrations"
    git -C "$dir" init -q -b base
    git -C "$dir" config user.email selftest@example.com
    git -C "$dir" config user.name selftest
    cat >"$dir/docs/backlog/Backlog.md" <<'EOF'
# Backlog

## Itens

### BLG-0100 - item existente

- **Status**: `done`

EOF
    : >"$dir/specs/100-fixture.md"
    git -C "$dir" add -A
    git -C "$dir" commit -q -m "base"
}

# run_case <nome> <exit esperado> <dir> [detached]
# detached=1 reproduz o checkout do CI de pull_request, onde --abbrev-ref HEAD
# devolve "HEAD": sem identidade propria o gate nao pode acusar a si mesmo.
run_case() {
    local name="$1" expected="$2" dir="$3" detached="${4:-0}" out status
    if [ "$detached" -eq 1 ]; then
        git -C "$dir" checkout -q --detach HEAD
    fi
    out="$(cd "$dir" && BASE_REF=base GITHUB_HEAD_REF= "$SCRIPT" 2>&1)"
    status=$?
    if [ "$status" -eq "$expected" ]; then
        printf 'selftest: PASS %s (exit %d)\n' "$name" "$status"
        pass=$((pass + 1))
    else
        printf 'selftest: FAIL %s: exit %d, esperado %d\n%s\n' "$name" "$status" "$expected" "$out" >&2
        fail=$((fail + 1))
    fi
}

# 1) branch sem ID sequencial novo -> passa
d="$tmproot/no-new"; make_repo "$d"
git -C "$d" checkout -q -b feat/nada
echo "nota" >>"$d/docs/backlog/Backlog.md"
git -C "$d" commit -qam "sem id novo"
run_case "branch sem ID sequencial novo" 0 "$d"

# 2) item novo e livre -> passa
d="$tmproot/item-livre"; make_repo "$d"
git -C "$d" checkout -q -b feat/blg0101
printf '\n### BLG-0101 - novo\n\n- **Status**: `in_progress`\n' >>"$d/docs/backlog/Backlog.md"
git -C "$d" commit -qam "item novo livre"
run_case "item novo e livre na base" 0 "$d"

# 3) item com ID ja existente na base -> colisao
d="$tmproot/item-colide"; make_repo "$d"
git -C "$d" checkout -q -b feat/blg0100
printf '\n### BLG-0100 - duplicado\n\n- **Status**: `in_progress`\n' >>"$d/docs/backlog/Backlog.md"
git -C "$d" commit -qam "item duplicado"
run_case "item com ID ja ocupado na base" 1 "$d"

# 4) reescrever o heading de um item existente nao e reserva de ID novo
d="$tmproot/item-reescrito"; make_repo "$d"
git -C "$d" checkout -q -b docs/retitle
python3 - "$d/docs/backlog/Backlog.md" <<'PY'
import sys
from pathlib import Path
p = Path(sys.argv[1])
p.write_text(
    p.read_text(encoding="utf-8").replace(
        "### BLG-0100 - item existente", "### BLG-0100 - titulo corrigido"
    ),
    encoding="utf-8",
)
PY
git -C "$d" commit -qam "corrige titulo do item existente"
run_case "reescrita de heading existente nao e ID novo" 0 "$d"

# 5) spec nova e livre -> passa
d="$tmproot/spec-livre"; make_repo "$d"
git -C "$d" checkout -q -b feat/spec110
: >"$d/specs/110-nova.md"
git -C "$d" add -A && git -C "$d" commit -qam "spec nova"
run_case "spec nova e livre na base" 0 "$d"

# 6) spec com numero ja usado -> colisao
d="$tmproot/spec-colide"; make_repo "$d"
git -C "$d" checkout -q -b feat/spec100
: >"$d/specs/100-outra-coisa.md"
git -C "$d" add -A && git -C "$d" commit -qam "spec duplicada"
run_case "spec com numero ja ocupado" 1 "$d"

# 7) migration com numero ja usado -> colisao
d="$tmproot/mig-colide"; make_repo "$d"
: >"$d/migrations/0001_base.up.sql"
git -C "$d" add -A && git -C "$d" commit -qam "migration base"
git -C "$d" checkout -q -b feat/mig0001
: >"$d/migrations/0001_outra.up.sql"
git -C "$d" add -A && git -C "$d" commit -qam "migration duplicada"
run_case "migration com numero ja ocupado" 1 "$d"

# 8) remover item do backlog nao conta como ID novo
d="$tmproot/remocao"; make_repo "$d"
git -C "$d" checkout -q -b docs/limpeza
printf '# Backlog\n\n## Itens\n\n' >"$d/docs/backlog/Backlog.md"
git -C "$d" commit -qam "remove item"
run_case "remocao de item nao e ID novo" 0 "$d"

# 9) HEAD detached (checkout do CI de pull_request): ID novo e livre segue passando
d="$tmproot/detached-livre"; make_repo "$d"
git -C "$d" checkout -q -b feat/blg0120
printf '\n### BLG-0120 - novo\n\n- **Status**: `in_progress`\n' >>"$d/docs/backlog/Backlog.md"
git -C "$d" commit -qam "item novo em detached"
run_case "HEAD detached com ID livre" 0 "$d" 1

# 10) HEAD detached ainda detecta colisao na base, que nao depende de identidade
d="$tmproot/detached-colide"; make_repo "$d"
git -C "$d" checkout -q -b feat/blg0100dup
printf '\n### BLG-0100 - duplicado\n\n- **Status**: `in_progress`\n' >>"$d/docs/backlog/Backlog.md"
git -C "$d" commit -qam "item duplicado em detached"
run_case "HEAD detached ainda detecta colisao na base" 1 "$d" 1

printf 'selftest: %d PASS, %d FAIL\n' "$pass" "$fail"
[ "$fail" -eq 0 ] || exit 1
