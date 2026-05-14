# Prompt Operacional - Diagnostico

Use este prompt quando a tarefa principal for diagnosticar estado real, auditar aderencia a spec ou produzir um relatorio de confirmacao.

```text
Leia primeiro:
- AGENTS.md
- a spec de diagnostico governante
- a spec de construcao correspondente
- os docs, scripts, testes e workflows citados pelas specs

Objetivo:
Diagnosticar o estado real do repositorio contra a spec, sem superestimar maturidade, cobertura ou enforcement.

Regras de execucao:
- classifique cada dimensao exigida pela spec como PASS, PARTIAL, FAIL ou BLOCKED
- cite evidencias concretas por arquivo e, quando necessario, por comando executado
- diferencie gap historico, gap remanescente e regressao introduzida
- se algum item nao se aplicar diretamente, registre a justificativa arquitetural
- nao trate ausencia de artefato obrigatorio como detalhe menor

Formato de saida:
1. resumo executivo
2. avaliacao por dimensao
3. evidencias por arquivo
4. gaps remanescentes
5. decisao final
```
