# Prompt Operacional - Remediation ou bugfix

Use este prompt quando a tarefa principal for corrigir um bug, fechar um gap diagnosticado ou remediar um desvio de compliance sem reabrir escopo amplo.

```text
Leia primeiro:
- AGENTS.md
- a spec governante da correcao
- a spec de diagnostico que evidenciou o gap, quando existir
- os arquivos afetados pelo bug ou desvio
- skills/00-skill-index/SKILL.md
- as skills aplicaveis ao escopo

Objetivo:
Corrigir o problema com a menor mudanca correta, restaurando aderencia a spec e deixando evidencia auditavel da remediacao.

Regras de execucao:
- reproduza o problema ou descreva a evidencia concreta antes de editar
- mantenha a correcao dentro do escopo diagnosticado; nao use remediation para introduzir feature nova
- atualize testes, checker, docs ou CI no mesmo patch quando eles fizerem parte da causa ou da evidencia da correcao
- se algum requisito obrigatorio continuar pendente, registre o gap remanescente e a justificativa
- use `docs/ai/compliance-exceptions.md` apenas quando uma excecao formal continuar necessaria apos a remediacao

Validacao minima:
- execute os checks que reproduzem ou confirmam a correcao
- declare claramente o que ficou resolvido e o que ainda permanece aberto

Formato de saida:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- gaps restantes
```
