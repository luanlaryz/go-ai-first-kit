# Runbook: acesso de agent a nuvem e credenciais

Este runbook documenta os controles executaveis que limitam o que um agent pode fazer com `aws`,
`kubectl` e arquivos `.env` neste repositorio, e o que esses controles **nao** garantem.

A orientacao normativa correspondente esta em [`.cursor/rules/agent-cloud-access.mdc`](../../.cursor/rules/agent-cloud-access.mdc).

## Controles

| Controle | Evento | Arquivo | Comportamento em falha |
|---|---|---|---|
| Cloud access guard | `beforeShellExecution` (`aws`, `kubectl`) | `.cursor/hooks/agent-cloud-access-guard.sh` | fail-closed (`ask`) |
| Env read guard | `preToolUse` (`Read`) | `.cursor/hooks/env-read-guard.sh` | fail-closed (`ask`) |
| Gate de wiring | `make check-agent-hooks` | `scripts/check-agent-hooks.sh` | falha o build |

O gate de wiring existe porque um controle que pode ser desligado em silencio nao e controle: ele
reprova quando `.cursor/hooks.json` perde um guard, remove `failClosed` ou aponta para script ausente
ou nao executavel.

## Configuracao de ambiente

Os marcadores que classificam o alvo de um comando ficam em
`.cursor/hooks/cloud-access-config.json`:

- `prod_markers`: nomes de profile, contexto e IDs de conta de producao.
- `dev_markers`: idem para desenvolvimento/homologacao.
- `local_markers`: emuladores e enderecos locais (ja preenchido com valores universais).

Listas vazias sao **seguras, nao permissivas**: sem marcadores nada resolve para `dev`, todo alvo fica
`unknown`, mutacao e negada e leitura exige aprovacao. Preencher `dev_markers` e o que libera o fluxo
de desenvolvimento sem aprovacao manual.

## Matriz de decisao

A classificacao e por **alvo** (conta, profile, contexto, endpoint), nunca por verbo isolado: um
comando de leitura apontado para producao vaza tanto quanto uma escrita.

| Acao | `dev` / `local` | `prod` | `unknown` |
|---|---|---|---|
| Mutacao `aws` | `ask` | `deny` | `deny` |
| Leitura de segredo `aws` | `allow` | `ask` | `ask` |
| Leitura comum `aws` | `allow` | `ask` | `ask` |
| `aws sts get-caller-identity` | `allow` | `allow` | `allow` |
| Mutacao `kubectl` | `ask` | `deny` | `deny` |
| `kubectl get secret` | `allow` | `ask` | `ask` |
| Leitura `kubectl` | `allow` | `ask` | `ask` |
| `kubectl` sem `--context` | `deny` | `deny` | `deny` |
| Leitura de `.env` | `ask` | `ask` | `ask` |

`kubectl` sem `--context` e negado independente do alvo porque o `current-context` pode apontar para
producao sem nenhum sinal no comando. A mensagem de bloqueio inclui a auto-correcao.

## Auditoria

Cada decisao diferente de `allow` gera uma linha em `.cursor/hooks/.agent-cloud-access-audit.log` com
data, decisao, ferramenta, verbo e alvo. O comando bruto **nao** e registrado, porque pode conter
credencial embutida. O arquivo e truncado ao passar de 512 KB e nao deve ser versionado.

## Limites declarados

Estes controles reduzem acidente, nao detem um ator malicioso:

- Sao contornaveis por `eval`, `bash -c`, wrapper de PATH ou subprocesso que o hook nao observa.
- Nao cobrem chamadas feitas por bibliotecas dentro de processos de teste ou de aplicacao.
- Nao substituem IAM/RBAC. Se a credencial disponivel permite destruir producao, a fronteira real
  esta errada, e o hook e apenas um lembrete.

## Verificacao

```bash
make check-agent-hooks                              # wiring, failClosed e permissao dos scripts
./.cursor/hooks/agent-cloud-access-guard-selftest.sh # matriz de decisao contra config de fixture
./.cursor/hooks/env-read-guard-selftest.sh           # protecao de .env e fail-closed
make env-keys                                        # chaves de .env sem expor valores
```
