# Metrics Catalog

| Metric | Type | Labels | Description |
|--------|------|--------|-------------|
| `http_requests_total` | counter | `route`, `method`, `status`, `{{DOMAIN_ACTOR}}_id` | Count of HTTP requests |
| `http_request_duration_seconds` | histogram | `route`, `{{DOMAIN_ACTOR}}_id` | Latency |
| `commands_enqueued_total` | counter | `command_type`, `{{DOMAIN_ACTOR}}_id` | Number of commands queued |
| `command_processing_duration_seconds` | histogram | `command_type`, `{{DOMAIN_ACTOR}}_id` | Worker processing time |
| `redis_operations_total` | counter | `operation`, `result` | Redis ops |
| `redis_operation_duration_seconds` | histogram | `operation` | Redis latency |
| `rate_limit_blocks_total` | counter | `{{DOMAIN_ACTOR}}_id`, `tier` | Requests throttled |
| `adk_step_total` | counter | `{{DOMAIN_ACTOR}}_id`, `step`, `status` | ADK tool usage |
| `adk_run_duration_seconds` | histogram | `{{DOMAIN_ACTOR}}_id` | Agent run durations |
| `grpc_requests_total` | counter | `service`, `method`, `code` | RPC calls |
| `otel_traces_exported_total` | counter | `backend` | Exporter health |
