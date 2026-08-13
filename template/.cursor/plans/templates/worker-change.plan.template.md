# Template: Worker Change

Frontmatter esperado ao copiar para um plano real:

```md
---
id: <plan-id>
status: DRAFT_NOT_EXECUTABLE
risk: high
area: worker
---
```

# Objetivo

Alterar consumo assincrono preservando idempotencia, ordenacao declarada, retry e destino de falha.

# Escopo permitido

- worker e adapters de fila correspondentes
- `test/conformance` e fixtures de mensagem

# Fora de escopo

- Consumo sem retry, sem DLQ e sem observabilidade.
- Efeito colateral nao idempotente em reprocessamento.

# Contratos afetados

- Contrato da mensagem e sua chave de idempotencia.
- Politica de retry, DLQ e replay.

# Plano de implementacao

1. Declarar a chave de idempotencia e o efeito de reprocessar a mesma mensagem.
2. Definir retry, backoff e destino de falha antes de implementar.
3. Implementar o handler e a instrumentacao (log, metrica, trace).
4. Provar reprocessamento seguro com a mesma mensagem duas vezes.

# Testes obrigatorios

- `go test ./test/conformance/...`
- `make test`
- `make race`
- `make guardrails`

# Rollback

- Reverter handler e configuracao de fila; mensagens em DLQ precisam de plano de replay.

# Stop conditions

- Idempotencia nao demonstravel.
- Ausencia de destino de falha (DLQ) ou de plano de replay.
- Ordenacao exigida pela spec nao garantida pelo transporte.
