# SDD Phase Template

Use this reference when creating a full paired SDD phase for {{PROJECT_SLUG}}.

## Feature spec template

```md
# Spec <number>: Phase <n> - <title>

## 1. Objetivo

## 2. Motivacao

## 3. Pergunta principal

## 4. Escopo

### Dentro do escopo

### Fora do escopo

## 5. Requisitos de design

## 6. Blocos obrigatorios desta fase

## 7. Requisitos funcionais

## 8. Decisoes obrigatorias de modelagem

## 9. Criterios de aceitacao

## 10. Evidencia obrigatoria

## 11. Fora do escopo tecnico explicito

## 12. Perguntas secundarias
```

## Diagnosis spec template

```md
# Spec <number>: Phase <n> - <title> Diagnosis

## 1. Objetivo

## 2. Pergunta principal

## 3. Escopo

## 4. Checklist obrigatorio

## 5. Classificacao

- PASS
- PARTIAL
- FAIL
- BLOCKED

## 6. Saida obrigatoria

`docs/reports/phase-<n>-<slug>-report.md`

## 7. Criterios de aceitacao
```

## File naming pattern

Prefer:

- `specs/<number>-phase-<n>-<slug>.md`
- `specs/<number>-phase-<n>-<slug>-diagnosis.md`
- `docs/reports/phase-<n>-<slug>-report.md`

If the repo already uses a different numbering range, follow the repo's current numbering.

## Prompt generation pattern

For each phase, generate exactly four prompts:

1. Create feature spec.
2. Implement feature spec.
3. Create diagnosis spec.
4. Execute diagnosis spec.

## Boundary language to reuse

Use these phrases when relevant:

- `nao abrir capability nova`
- `nao reabrir o roadmap principal`
- `tratar {{PROJECT_SLUG}} como lib reutilizavel e versionavel`
- `avaliar evidencia real, nao intencao`
- `diferenciar readiness de framework de readiness da aplicacao`
- `gerar report canonico da fase`
- `concluir com decisao binaria`
