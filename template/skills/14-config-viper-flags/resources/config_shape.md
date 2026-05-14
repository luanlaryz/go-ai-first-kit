# Config Shape

Use Viper loading order: env vars > config file > defaults.

```yaml
app:
  env: dev|staging|prod
  http_port: 8080
  grpc_port: 9090
  log_level: info
  feature_flags:
    global:
      streaming_default: true
    {{DOMAIN_ACTOR}}_overrides:
      {{DOMAIN_ACTOR}}-123:
        streaming_default: false
        provider_override: "openai:gpt-4o"
redis:
  url: redis://localhost:6379
postgres:
  url: postgres://postgres:postgres@localhost:5432/app?sslmode=disable
sqs:
  queue_url: http://localstack:4566/000000000000/cmd-{{DOMAIN_ACTOR}}-conversation.fifo
telemetry:
  otel_collector: http://otel-collector:4317
  metrics_port: 2112
```

## Feature Flag Access
```go
type FlagStore interface {
    Enabled(ctx context.Context, {{DOMAIN_ACTOR}}ID, flag string) bool
}
```
- Flags stored in Postgres `control.feature_flags` table, cached in Redis for 60s.

## Secrets
- Pull from AWS SSM or Doppler; never commit to repo.
