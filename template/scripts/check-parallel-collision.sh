#!/usr/bin/env bash
# Parallel session collision gate (skills/31-parallel-session-coordination, Protocolos B e C).
#
# Detecta os IDs sequenciais que ESTA branch introduz (BLG-NNNN no backlog,
# specs/NNN, migrations/NNNN) e confirma que seguem livres em tres fontes, porque
# nenhuma delas sozinha enxerga tudo:
#
#   1. base ref     -> ID ja mergeado por outra sessao
#   2. PRs abertos  -> ID reservado em PR ainda nao mergeado (invisivel na base)
#   3. refs remotas -> branch publicada cujo nome carrega o ID; o PR pode nao
#                      existir ainda, ou nao citar o ID no texto
#
# Um `max()` do backlog nao resolve: o proximo numero livre na base pode ja estar
# reservado por uma sessao que ainda nao mergeou.
#
# Tambem aplica o Protocolo C: branch homonima ja publicada no remoto com
# trabalho que nao esta em HEAD.
#
# Limitacao declarada: a varredura de PRs e heuristica. Um PR pode reservar um ID
# sem cita-lo em titulo, corpo, branch ou paths. A colisao residual e absorvida
# pelo Protocolo D quando o merge acusar.
#
# Exit 0 = livre (ou nada novo). Exit 1 = colisao: renumere.
set -uo pipefail

BASE_REF="${BASE_REF:-origin/main}"
fail() { printf 'check-parallel-collision: FAIL: %s\n' "$*" >&2; }
warn() { printf 'check-parallel-collision: WARN: %s\n' "$*" >&2; }

git rev-parse --git-dir >/dev/null 2>&1 || { echo "check-parallel-collision: not a git repo" >&2; exit 1; }

# `timeout` e GNU coreutils e nao existe no macOS por padrao. Sem esta resolucao o
# gate degradaria para "fonte indisponivel" em toda maquina Darwin, perdendo em
# silencio as duas fontes que mais importam.
if command -v timeout >/dev/null 2>&1; then
    with_timeout() { timeout "$@"; }
elif command -v gtimeout >/dev/null 2>&1; then
    with_timeout() { gtimeout "$@"; }
else
    with_timeout() { shift; "$@"; }
fi

# Fetch curto: sem isso o veredito e stale, que e pior que nao checar. O nome do
# branch remoto vem do proprio BASE_REF (origin/<branch>).
base_remote="${BASE_REF#origin/}"
if [ "$base_remote" != "$BASE_REF" ]; then
    if ! with_timeout 20 git fetch origin "$base_remote" -q 2>/dev/null; then
        warn "git fetch falhou; comparando com $BASE_REF possivelmente desatualizado"
    fi
fi
git rev-parse --verify -q "$BASE_REF" >/dev/null || { fail "$BASE_REF nao existe localmente"; exit 1; }

# Identidade da propria branch. Ela e o que permite excluir o proprio PR e a
# propria ref da varredura; sem ela, todo ID desta branch aparece "ocupado por si
# mesmo" e o gate reprova trabalho legitimo.
#
# No CI de `pull_request` o checkout fica em HEAD detached e `--abbrev-ref HEAD`
# devolve "HEAD", entao GITHUB_HEAD_REF (nome da branch de origem do PR) tem
# precedencia. Sem nenhuma das duas, a auto-exclusao e impossivel: nesse caso o
# gate degrada e pula as fontes que exigem identidade, em vez de emitir falso
# positivo. Falso positivo bloqueia entrega correta; falso negativo residual e
# absorvido pelo Protocolo D no merge.
branch="${GITHUB_HEAD_REF:-}"
[ -n "$branch" ] || branch="$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo HEAD)"
self_known=1
if [ -z "$branch" ] || [ "$branch" = "HEAD" ]; then
    self_known=0
    warn "branch propria nao identificavel (HEAD detached sem GITHUB_HEAD_REF): PRs abertos e refs remotas nao verificados"
fi

# --- fonte 2: PRs abertos (texto + branch + paths) -------------------------
pr_ids_file="$(mktemp)"; pr_paths_file="$(mktemp)"
trap 'rm -f "$pr_ids_file" "$pr_paths_file"' EXIT
# O PR desta propria branch e excluido: ele cita os IDs que esta branch introduz,
# e conta-lo geraria falso positivo -- no CI de um PR aberto o gate reprovaria
# sempre, e localmente reprovaria todo push posterior ao primeiro.
# `gh --jq` nao aceita --arg; o filtro le a branch por env.SELF.
pr_scan_ok=0
if [ "$self_known" -eq 1 ] && command -v gh >/dev/null 2>&1; then
    if SELF="$branch" gh pr list --state open --json number,title,body,headRefName \
        --jq '.[]|select(.headRefName != env.SELF)|"\(.title) \(.headRefName) \(.body // "")"' \
        >"$pr_ids_file" 2>/dev/null; then
        pr_scan_ok=1
        for n in $(SELF="$branch" gh pr list --state open --json number,headRefName \
            --jq '.[]|select(.headRefName != env.SELF)|.number' 2>/dev/null); do
            gh pr view "$n" --json files --jq '.files[].path' 2>/dev/null
        done | grep -E '^(specs|migrations)/[0-9]+' >>"$pr_paths_file" 2>/dev/null || true
    fi
fi
[ "$pr_scan_ok" -eq 1 ] || warn "gh indisponivel: varredura de PRs abertos pulada (cobertura reduzida)"

# --- fonte 3: refs remotas -------------------------------------------------
# Sem identidade propria a varredura de refs acusaria a propria branch publicada.
remote_refs=""
if [ "$self_known" -eq 1 ]; then
    remote_refs="$(with_timeout 20 git ls-remote --heads origin 2>/dev/null | sed 's#.*refs/heads/##' || true)"
    [ -n "$remote_refs" ] || warn "git ls-remote falhou: branches remotas nao verificadas"
fi
export REMOTE_REFS="$remote_refs"

python3 - "$BASE_REF" "$branch" "$pr_ids_file" "$pr_paths_file" <<'PY'
import os
import re
import subprocess
import sys

base, branch, pr_ids_path, pr_paths_path = sys.argv[1:5]

BACKLOG = "docs/backlog/Backlog.md"
# Item heading in the single-file backlog: `### BLG-0001 - title`.
ITEM_HEADING = r"###\s+BLG-(\d{4})\b"


def git(*a):
    try:
        return subprocess.run(["git", *a], capture_output=True, text=True).stdout
    except Exception:
        return ""


def read(p):
    try:
        with open(p, encoding="utf-8") as fh:
            return fh.read()
    except Exception:
        return ""


pr_text = read(pr_ids_path)
pr_paths = read(pr_paths_path)
remote_refs = os.environ.get("REMOTE_REFS", "")

# IDs introduzidos por esta branch -----------------------------------------
introduced = {"item": set(), "spec": set(), "migration": set()}

# Um heading que aparece como adicionado E removido no mesmo diff e reescrita do
# mesmo item (reordenacao, correcao de titulo), nao reserva de ID novo. Sem essa
# subtracao, todo ajuste de item existente viraria falso positivo -- e falso
# positivo bloqueia entrega correta.
added_items: set[int] = set()
removed_items: set[int] = set()
diff = git("diff", f"{base}...HEAD", "--", BACKLOG)
for line in diff.splitlines():
    m = re.match(rf"^\+{ITEM_HEADING}", line)
    if m:
        added_items.add(int(m.group(1)))
        continue
    m = re.match(rf"^-{ITEM_HEADING}", line)
    if m:
        removed_items.add(int(m.group(1)))
introduced["item"] = added_items - removed_items

for f in git("diff", "--name-only", "--diff-filter=A", f"{base}...HEAD").splitlines():
    f = f.strip()
    m = re.match(r"^specs/(\d{3})-", f)
    if m:
        introduced["spec"].add(int(m.group(1)))
    m = re.match(r"^migrations/(\d{4})[_-]", f)
    if m:
        introduced["migration"].add(int(m.group(1)))

# Fontes de ocupacao --------------------------------------------------------
main_backlog = git("show", f"{base}:{BACKLOG}")
main_items = {int(n) for n in re.findall(rf"^{ITEM_HEADING}", main_backlog, re.M)}

main_specs = set()
main_migs = set()
for f in git("ls-tree", "-r", "--name-only", base).splitlines():
    f = f.strip()
    m = re.match(r"^specs/(\d{3})-", f)
    if m:
        main_specs.add(int(m.group(1)))
    m = re.match(r"^migrations/(\d{4})[_-]", f)
    if m:
        main_migs.add(int(m.group(1)))

if not any(introduced.values()):
    print("check-parallel-collision: ok (nenhum ID sequencial novo nesta branch)")
    sys.exit(0)

pr_item_ids = {int(n) for n in re.findall(r"(?:BLG-|blg)(\d+)", pr_text)}
pr_spec_ids = set()
pr_mig_ids = set()
for line in pr_paths.splitlines():
    line = line.strip()
    m = re.match(r"^specs/(\d{3})", line)
    if m:
        pr_spec_ids.add(int(m.group(1)))
    m = re.match(r"^migrations/(\d{4})", line)
    if m:
        pr_mig_ids.add(int(m.group(1)))

# Refs remotas: o ID no nome da branch e a chave de deduplicacao entre sessoes.
# A propria branch nao conta como colisao consigo mesma. Nomes de branch tambem
# carregam numero de spec e de migration ("feat/spec-110-foo"), entao os tres
# tipos sao extraidos, nao so o item de backlog.
ref_item_ids = set()
ref_spec_ids = set()
ref_mig_ids = set()
for ref in remote_refs.splitlines():
    ref = ref.strip()
    if not ref or ref == branch:
        continue
    lowered = ref.lower()
    for n in re.findall(r"blg[-_]?(\d+)", lowered):
        ref_item_ids.add(int(n))
    for n in re.findall(r"spec[s]?[-_/]?(\d{3})\b", lowered):
        ref_spec_ids.add(int(n))
    for n in re.findall(r"migration[s]?[-_/]?(\d{4})\b", lowered):
        ref_mig_ids.add(int(n))

SOURCES = {
    "item": ((main_items, base), (pr_item_ids, "PR aberto"), (ref_item_ids, "branch remota")),
    "spec": ((main_specs, base), (pr_spec_ids, "PR aberto"), (ref_spec_ids, "branch remota")),
    "migration": ((main_migs, base), (pr_mig_ids, "PR aberto"), (ref_mig_ids, "branch remota")),
}
LABEL = {
    "item": lambda i: f"BLG-{i:04d}",
    "spec": lambda i: f"specs/{i:03d}",
    "migration": lambda i: f"migrations/{i:04d}",
}

errors = []
for kind, ids in introduced.items():
    for i in sorted(ids):
        where = [name for occupied, name in SOURCES[kind] if i in occupied]
        if where:
            errors.append(f"{LABEL[kind](i)} ja ocupado em: {', '.join(where)}")

if errors:
    for e in errors:
        print(f"check-parallel-collision: FAIL: {e}", file=sys.stderr)
    print(
        "check-parallel-collision: renumere com sort numerico contra as tres fontes, "
        "revalide as referencias cruzadas (backlog, plano, corpo do PR, specs) e renomeie a "
        "branch se ainda nao publicada (skills/31-parallel-session-coordination, Protocolo B).",
        file=sys.stderr,
    )
    sys.exit(1)

summary = ", ".join(f"{k}={sorted(v)}" for k, v in introduced.items() if v)
print(f"check-parallel-collision: ok (IDs livres nas fontes disponiveis; {summary})")
PY
py_status=$?
[ "$py_status" -eq 0 ] || exit "$py_status"

# --- Protocolo C: branch homonima ja publicada -----------------------------
if [ -n "$remote_refs" ] && printf '%s\n' "$remote_refs" | grep -qx -- "$branch"; then
    remote_sha="$(git rev-parse -q --verify "refs/remotes/origin/$branch" 2>/dev/null || true)"
    if [ -n "$remote_sha" ] && [ "$remote_sha" != "$(git rev-parse HEAD)" ]; then
        if ! git merge-base --is-ancestor "$remote_sha" HEAD 2>/dev/null; then
            fail "branch '$branch' ja existe no remoto com trabalho que nao esta em HEAD"
            printf 'check-parallel-collision: integre por rebase e prove o fast-forward antes de publicar (Protocolo D). Force push e proibido.\n' >&2
            exit 1
        fi
    fi
fi

exit 0
