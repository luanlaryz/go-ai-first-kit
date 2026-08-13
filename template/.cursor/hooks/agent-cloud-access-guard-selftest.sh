#!/usr/bin/env bash
# Selftest for the agent cloud access guard.
#
# Proves the classifier decides by target, never allows without proof of
# dev/local, keeps ordinary tooling untouched, and fails closed.
#
# The decision matrix is exercised against a FIXTURE config in a temp root, not
# against the project's real `.cursor/hooks/cloud-access-config.json`: the
# selftest must hold regardless of which markers a given project configured.
# The shipped default (empty marker lists) is asserted separately.
set -uo pipefail

root="$(git rev-parse --show-toplevel 2>/dev/null || pwd)"
GUARD_PY="$root/.cursor/hooks/agent-cloud-access-guard.py"
GUARD_SH="$root/.cursor/hooks/agent-cloud-access-guard.sh"
failures=0

command -v python3 >/dev/null 2>&1 || { echo "python3 is required for this selftest" >&2; exit 1; }
[ -f "$GUARD_PY" ] || { echo "guard not found: $GUARD_PY" >&2; exit 1; }

tmp_configured="$(mktemp -d)"
tmp_unconfigured="$(mktemp -d)"
trap 'rm -rf "$tmp_configured" "$tmp_unconfigured"' EXIT

mkdir -p "$tmp_configured/.cursor/hooks" "$tmp_unconfigured/.cursor/hooks"
cat >"$tmp_configured/.cursor/hooks/cloud-access-config.json" <<'JSON'
{
  "prod_markers": ["acme-prod", "eks-prod", "111111111111"],
  "dev_markers": ["acme-dev", "eks-dev", "222222222222"],
  "local_markers": ["localhost", "127.0.0.1", "localstack", ":4566"]
}
JSON
cat >"$tmp_unconfigured/.cursor/hooks/cloud-access-config.json" <<'JSON'
{
  "prod_markers": [],
  "dev_markers": [],
  "local_markers": ["localhost", "127.0.0.1", "localstack", ":4566"]
}
JSON

event() {
  python3 -c 'import json,sys; print(json.dumps({"command": sys.argv[1], "cwd": "."}))' "$1"
}

permission() {
  python3 -c 'import json,sys
try:
    print(json.load(sys.stdin).get("permission", "<none>"))
except Exception:
    print("<invalid>")'
}

# Runs the classifier against a given fake project root, with AWS target env
# cleared unless the caller re-adds it through EXTRA_ENV.
decide_in() {
  local fake_root="$1" cmd="$2"
  event "$cmd" \
    | env -u AWS_PROFILE -u AWS_DEFAULT_PROFILE -u AWS_ENDPOINT_URL -u AWS_REGION \
        CURSOR_PROJECT_ROOT="$fake_root" ${EXTRA_ENV:-} python3 "$GUARD_PY" 2>/dev/null \
    | permission
}

expect() {
  local label="$1" want="$2" cmd="$3"
  local got
  got="$(decide_in "$tmp_configured" "$cmd")"
  if [ "$got" = "$want" ]; then
    printf 'ok   %-56s -> %s\n' "$label" "$got"
  else
    printf 'FAIL %-56s -> got=%s want=%s\n' "$label" "$got" "$want" >&2
    failures=$((failures + 1))
  fi
}

expect_env() {
  local label="$1" want="$2" cmd="$3" assignment="$4"
  local got
  got="$(EXTRA_ENV="$assignment" decide_in "$tmp_configured" "$cmd")"
  if [ "$got" = "$want" ]; then
    printf 'ok   %-56s -> %s\n' "$label" "$got"
  else
    printf 'FAIL %-56s -> got=%s want=%s\n' "$label" "$got" "$want" >&2
    failures=$((failures + 1))
  fi
}

expect_unconfigured() {
  local label="$1" want="$2" cmd="$3"
  local got
  got="$(decide_in "$tmp_unconfigured" "$cmd")"
  if [ "$got" = "$want" ]; then
    printf 'ok   %-56s -> %s\n' "$label" "$got"
  else
    printf 'FAIL %-56s -> got=%s want=%s\n' "$label" "$got" "$want" >&2
    failures=$((failures + 1))
  fi
}

echo "== mutacao em producao = deny =="
expect "rds modify prod" deny "aws rds modify-db-instance --profile acme-prod --db-instance-identifier x"
expect "secretsmanager put prod" deny "aws secretsmanager put-secret-value --profile acme-prod --secret-id app-secret"
expect "sqs purge conta prod" deny "aws sqs purge-queue --queue-url https://sqs.us-east-1.amazonaws.com/111111111111/q.fifo"
expect "kubectl patch prod" deny "kubectl --context eks-prod -n app patch cronjob x -p {}"
expect "kubectl exec prod" deny "kubectl --context eks-prod -n app exec deploy/api -- sh"

echo "== leitura de segredo em producao = ask =="
expect "get-secret-value prod" ask "aws secretsmanager get-secret-value --profile acme-prod --secret-id app-secret"
expect "kubectl get secret prod" ask "kubectl --context eks-prod -n app get secret app-secret -o yaml"
expect "ssm get-parameter prod" ask "aws ssm get-parameter --profile acme-prod --name /x --with-decryption"

echo "== leitura em dev/local = allow =="
expect "get-secret-value dev" allow "aws secretsmanager get-secret-value --profile acme-dev --secret-id app-secret"
expect "sqs attrs local" allow "aws sqs get-queue-attributes --endpoint-url http://localhost:4566 --queue-url http://localhost:4566/000000000000/q"
expect "kubectl get deployment dev" allow "kubectl --context arn:aws:eks:us-east-1:222222222222:cluster/eks-dev -n app get deployment api -o json"

echo "== alvo ambiguo nunca vira allow =="
expect "secret read sem alvo" ask "aws secretsmanager get-secret-value --secret-id app-secret"
expect "leitura comum sem alvo" ask "aws sqs get-queue-attributes --queue-url https://example.invalid/q"
expect "mutacao sem alvo" deny "aws secretsmanager put-secret-value --secret-id app-secret"
expect "sts get-caller-identity" allow "aws sts get-caller-identity"

echo "== alvo resolvido pelo ambiente do hook =="
expect_env "env prod + secret read" ask "aws secretsmanager get-secret-value --secret-id x" "AWS_PROFILE=acme-prod"
expect_env "env prod + mutacao" deny "aws secretsmanager put-secret-value --secret-id x" "AWS_PROFILE=acme-prod"
expect_env "env dev + secret read" allow "aws secretsmanager get-secret-value --secret-id x" "AWS_PROFILE=acme-dev"
expect_env "env endpoint local" allow "aws sqs get-queue-attributes --queue-url q" "AWS_ENDPOINT_URL=http://localhost:4566"

echo "== kubectl sem --context = deny com auto-correcao =="
expect "kubectl get sem context" deny "kubectl -n app get pods"
expect "kubectl logs sem context" deny "kubectl logs -n app deploy/api"
msg="$(event "kubectl -n app get pods" \
  | env -u AWS_PROFILE CURSOR_PROJECT_ROOT="$tmp_configured" python3 "$GUARD_PY" 2>/dev/null \
  | python3 -c 'import json,sys; print(json.load(sys.stdin).get("agent_message",""))')"
case "$msg" in
  *--context*) printf 'ok   %-56s -> mensagem sugere --context\n' "auto-correcao presente" ;;
  *) printf 'FAIL %-56s -> agent_message sem --context\n' "auto-correcao presente" >&2; failures=$((failures + 1)) ;;
esac

echo "== config default (sem marcadores) e restritiva, nao permissiva =="
expect_unconfigured "mutacao sem config = deny" deny "aws s3 rm s3://bucket/key"
expect_unconfigured "leitura sem config = ask" ask "aws s3 ls s3://bucket"
expect_unconfigured "local ainda liberado" allow "aws sqs list-queues --endpoint-url http://localhost:4566"

echo "== vazao preservada: ferramental comum e nao-invocacao =="
expect "entrypoint de make" allow "make test"
expect "runner python de teste" allow "python3 scripts/smoke.py --opt-in"
expect "psql leitura" allow "psql \"\$DATABASE_URL\" -c 'select 1'"
expect "kubectl como dado em loop" allow "for value in python3 kubectl git; do command -v \$value; done"
expect "shutil.which kubectl" allow "python3 -c 'import shutil; print(shutil.which(\"kubectl\"))'"
expect "git status" allow "git status --short"
expect "rg em path com aws" allow "rg -n 'queue' internal/adapters/aws/queue.go"

echo "== fail-closed =="
if [ -x "$GUARD_SH" ]; then
  bad="$(printf 'not json at all' | "$GUARD_SH" 2>/dev/null | permission)"
  if [ "$bad" = "ask" ]; then
    printf 'ok   %-56s -> %s\n' "evento invalido nao produz allow" "$bad"
  else
    printf 'FAIL %-56s -> got=%s want=ask\n' "evento invalido nao produz allow" "$bad" >&2
    failures=$((failures + 1))
  fi
else
  printf 'FAIL %-56s -> wrapper nao executavel\n' "fail-closed via wrapper" >&2
  failures=$((failures + 1))
fi

if [ "$failures" -ne 0 ]; then
  echo "agent-cloud-access-guard-selftest: FAIL ($failures caso(s))" >&2
  exit 1
fi
echo "agent-cloud-access-guard-selftest: ok"
