# Guia de adoção

Use quando o projeto será mantido por humanos e agentes, mudanças precisam ser auditáveis e o time quer specs, diagnosis specs, reports e gates explícitos.

## Fluxo recomendado

```bash
go install ./cmd/gakit
gakit create ./novo-projeto
gakit diagnose --path ./novo-projeto --min-score 80 --report-only
```

Use `gakit diagnose --path .` como gate de adoção em projetos existentes. O report ajuda a priorizar lacunas de DX, AI-first, security, arquitetura, OpenAPI, documentação e governança.

Os projetos gerados já vêm com backlog governado (`docs/backlog/Backlog.md` + `skills/26-backlog-item-intake/`) e governança de release/changelog/ADR (`docs/decisions/`, `docs/release-versioning-policy.md`, `docs/release-notes-policy.md`, `docs/release-checklist.md`), prontos para enforcement por specs, reports e gates.

Evite para spikes descartáveis ou projetos que não desejam enforcement em CI.
