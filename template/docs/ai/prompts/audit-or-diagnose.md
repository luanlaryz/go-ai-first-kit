# Prompt Operacional - Auditoria ou diagnostico

Use este prompt quando a tarefa principal for auditar implementacao, docs, enforcement ou aderencia a spec.

```text
Leia primeiro:
- AGENTS.md
- a spec de diagnostico governante
- a spec de construcao correspondente
- os docs, scripts, testes e workflows citados pela spec

Objetivo:
Auditar o estado real do repositorio contra a spec, sem superestimar capacidade implementada.

Regras de execucao:
- classifique cada dimensao exigida pela spec como PASS, PARTIAL, FAIL ou BLOCKED
- cite evidencias concretas por arquivo e, quando necessario, por comando executado
- diferencie gap historico de gap introduzido pela mudanca atual
- nao trate ausencia de artefato como detalhe menor quando ela for criterio de aceite
- se a implementacao divergir da documentacao, relate a divergencia explicitamente

Formato de saida:
1. resumo executivo
2. avaliacao por dimensao
3. evidencias por arquivo
4. gaps remanescentes
5. decisao final
```
