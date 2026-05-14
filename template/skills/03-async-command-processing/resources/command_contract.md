# Command Envelope Contract

```json
{
  "version": 1,
  "{{DOMAIN_ACTOR}}_id": "{{DOMAIN_ACTOR}}-123",
  "conversation_id": "conv-789",
  "command_id": "uuid",
  "command_type": "send_agent_reply",
  "payload": {
    "message_id": "uuid",
    "body": "...",
    "metadata": {
      "language": "en",
      "priority": "default"
    }
  },
  "dedup_key": "{{DOMAIN_ACTOR}}-123:conv-789:send_agent_reply:uuid",
  "created_at": "2026-03-05T12:00:00Z"
}
```

- **SQS FIFO** queue name: `cmd-{{DOMAIN_ACTOR}}-conversation.fifo`.
- **MessageGroupId**: `{{DOMAIN_ACTOR}}:{{{DOMAIN_ACTOR}}_id}:conversation:{conversation_id}`.
- **MessageDeduplicationId**: SHA256 of `dedup_key`.
- **DLQ**: `cmd-{{DOMAIN_ACTOR}}-conversation-dlq` with `maxReceiveCount=3`.
- **Attributes**: `request_id`, `trace_id`, `command_type`.

## Worker Expectations
- Poll via long polling (20s) with visibility timeout >= worker SLA.
- Use idempotent inbox table (Postgres) storing `command_id`, `status`, `processed_at`.
- Confirm processing before deleting message; ack errors by leaving message to retry.
- Publish completion events to Redis Streams when needed for streaming skill.
