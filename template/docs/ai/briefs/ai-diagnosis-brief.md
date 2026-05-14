# Brief real - AI diagnosis

## Objetivo

Diagnosticar o estado real de uma trilha, patch ou artefato operacional contra a spec governante, produzindo classificacao auditavel e sem drift aspiracional.

## Specs governantes

- `AGENTS.md`
- spec de diagnostico da trilha
- spec de construcao correspondente
- `docs/ai/ai-contribution-contract.md`

## Escopo desta entrega

- classificar cada dimensao relevante como PASS, PARTIAL, FAIL ou BLOCKED
- sustentar conclusoes com evidencia por arquivo e, quando necessario, por comando
- separar gap historico, gap remanescente e regressao atual
- registrar justificativa arquitetural quando algum item nao se aplicar diretamente

## Fora do escopo

- reescrever a spec durante o diagnostico
- tratar ausencia de artefato obrigatorio como detalhe menor
- inferir maturidade nao sustentada por script, teste ou artefato versionado

## Evidencias esperadas

- resumo executivo
- avaliacao por dimensao
- evidencias por arquivo
- gaps remanescentes
- decisao final rastreavel
