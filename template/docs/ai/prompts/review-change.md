# Prompt Operacional - Review de mudanca

Use este prompt quando a tarefa principal for revisar codigo, docs, testes, CI ou enforcement antes de merge.

```text
Leia primeiro:
- AGENTS.md
- a spec governante da mudanca revisada
- a spec de diagnostico correspondente, quando existir
- os arquivos alterados e os testes relacionados
- skills/00-skill-index/SKILL.md
- as skills aplicaveis ao escopo revisado

Objetivo:
Revisar a mudanca contra a spec, priorizando bugs, regressao comportamental, riscos operacionais, gaps de teste e drift documental.

Regras de execucao:
- apresente findings primeiro, ordenados por severidade
- cite arquivos e comportamento observavel, nao opinioes vagas
- trate ausencia de spec, teste ou doc obrigatorio como finding real quando afetar criterio de aceite
- diferencie bug real, risco potencial e gap residual
- se nao houver findings, diga isso explicitamente e registre riscos ou coberturas ainda limitadas

Formato de saida:
1. findings
2. perguntas ou premissas abertas
3. resumo curto da mudanca
```
