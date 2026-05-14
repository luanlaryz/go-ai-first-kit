---
name: sqs-fifo-idempotency
description: Operate the SQS FIFO queues with {{DOMAIN_ACTOR}}:conversation grouping, content dedup, DLQ policies, and inbox-based idempotency.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# SQS FIFO & Idempotency
Goal: ensure command processing is exactly-once using SQS FIFO ordering + Postgres inbox ledger and solid DLQ handling.

## When to Use
- Modifying queue producers/consumers, dedup keys, or DLQ logic.
- Building replay tooling or monitoring.
- Changing inbox schema or processing semantics.

## Non-negotiables
1. `MessageGroupId = {{DOMAIN_ACTOR}}:{{{DOMAIN_ACTOR}}_id}:conversation:{conversation_id}`.
2. `MessageDeduplicationId` derived from deterministic key (hash of {{DOMAIN_ACTOR}} + conversation + command).
3. Inbox table guarded by unique constraint on `command_id` (or dedup key) + `{{DOMAIN_ACTOR}}_id`.
4. No manual deletes from queue without persisting state; always let worker ack.
5. DLQ policy matches [dlq_policy.md](resources/dlq_policy.md).

## Do / Don't
- **Do** include tracing attributes (`request_id`, `trace_id`).
- **Do** extend visibility timeout if downstream call > default.
- **Do** store result/outcome in Postgres for auditing.
- **Don't** rely solely on SQS dedup for idempotency; inbox must exist.
- **Don't** log full payloads with PII; use hashes/ids.
- **Don't** perform destructive replays without backup.

## Interfaces / Contracts
- Inbox repository sample:
  ```go
  type InboxStore interface {
      Insert(ctx context.Context, cmd CommandEnvelope) (inserted bool, err error)
      Complete(ctx context.Context, commandID string, status string) error
  }
  ```
- Replay command script references DLQ policy resource.

## Checklists
**Before coding**
- [ ] Confirm dedup key format and retention (5 minutes default for FIFO).
- [ ] Ensure inbox table indexes support new queries.
- [ ] Validate IAM permissions for queue access.

**During**
- [ ] Wrap SQS send/receive with retries/backoff.
- [ ] On worker start, warm up metrics for queue depth/backlog.
- [ ] Test failure paths (throwing error vs panic) to ensure retries happen.

**After**
- [ ] Update runbooks for DLQ replay if semantics changed.
- [ ] Add integration tests via LocalStack covering dedup collisions.
- [ ] Validate CloudWatch alarms triggered appropriately.

## Definition of Done
- Commands processed exactly once per {{DOMAIN_ACTOR}} conversation even under retries.
- DLQ/backlog metrics visible and alarms configured.
- Replay tooling documented and tested.
- Inbox table consistent with queue schema.

## Minimal Examples
- Dedup key builder: `fmt.Sprintf("%s:%s:%s:%s", {{DOMAIN_ACTOR}}ID, conversationID, commandType, payload.MessageID)` hashed with SHA256.
- Inbox upsert pattern: `INSERT ... ON CONFLICT DO NOTHING` returning bool to skip duplicate work.
