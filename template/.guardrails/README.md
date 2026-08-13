# Guardrail allowlists

Cada arquivo aqui e a lista de excecoes de um gate estrutural. O ponto central:
**excecao tem prazo**. Uma entrada com `expires_at` no passado nao e ignorada nem
avisada, ela reprova o gate. Divida tecnica registrada aqui e divida com data de
vencimento, nao permissao permanente.

## Formato

```yaml
version: 1
exceptions:
  - path: internal/adapters/http/router.go
    symbol: Router.Register
    rule: function-size
    justification: tabela de rotas declarativa; quebrar reduziria legibilidade
    owner: nome-ou-time
    expires_at: 2026-12-31
    ref: BLG-0042
```

Campos obrigatorios em toda entrada: `path`, `rule`, `justification`, `owner`,
`expires_at`, `ref`. Gates com escopo por simbolo (`function-size`, `port-size`,
`ignored-error`) exigem tambem `symbol`.

Um arquivo ausente ou malformado e **erro duro**, nunca "sem excecoes": apagar a
allowlist nao pode ser um jeito de passar no gate. Se nao houver excecoes, o
arquivo existe com `exceptions: []`.

## Arquivos

| Arquivo | Gate | Escopo da chave |
|---|---|---|
| `function-size-exceptions.yaml` | `guardrails funcsize` | `path` + `symbol` |
| `port-size-exceptions.yaml` | `guardrails portsize` | `path` + `symbol` |
| `ignored-error-exceptions.yaml` | `guardrails ignored-errors` | `path` + `symbol` (nome da funcao chamada) |
| `public-route-exceptions.yaml` | `guardrails public-route` | `path` (`METODO /rota` ou arquivo de contrato) |
| `package-exceptions.yaml` | `scripts/check-package-clean.sh` | prefixo de `path` |
| `governed-change-exceptions.yaml` | `scripts/check-governed-change.sh` | glob de `path` |

## Verificacao

```bash
make guardrails                                    # roda todos os gates
go run ./tools/guardrails allowlist-paths .guardrails/package-exceptions.yaml
```
