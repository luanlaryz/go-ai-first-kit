# Guia de adoção

Use quando o projeto será mantido por humanos e agentes, mudanças precisam ser auditáveis e o time quer specs, diagnosis specs, reports e gates explícitos.

## Fluxo recomendado

```bash
go install ./cmd/gakit
gakit create ./novo-projeto
gakit diagnose --path ./novo-projeto --min-score 80 --report-only
```

Use `gakit diagnose --path .` como gate de adoção em projetos existentes. O report ajuda a priorizar lacunas de DX, AI-first, security, arquitetura, OpenAPI, documentação e governança.

Evite para spikes descartáveis ou projetos que não desejam enforcement em CI.
