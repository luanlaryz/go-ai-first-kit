# Prompt Operacional - Implementacao de tarefa

Use este prompt quando a tarefa principal for implementar ou corrigir algo no repositorio.

```text
Leia primeiro:
- AGENTS.md
- a spec governante da mudanca
- a spec de diagnostico correspondente, quando existir
- specs/010-feature-matrix.md, quando a tarefa tocar status, cobertura, risco ou paridade
- skills/00-skill-index/SKILL.md
- as skills aplicaveis ao escopo
- os arquivos diretamente afetados

Objetivo:
Implementar a tarefa pedida com a menor mudanca correta, mantendo rastreabilidade por spec.

Regras de execucao:
- identifique explicitamente qual spec e qual secao governam a mudanca antes de editar
- nao implemente feature fora de spec; se faltar especificacao, pare e declare o gap
- mantenha fronteiras entre `pkg/*` e `internal/*`
- atualize testes no mesmo patch quando o comportamento mudar
- se algum teste ou check obrigatorio nao se aplicar, registre excecao formal em `docs/ai/compliance-exceptions.md`
- declare impacto em docs ou superficie publica quando houver
- atualize docs operacionais somente quando houver artefato real sustentando a afirmacao
- nao deixe marcadores vazios em Makefile, CI, prompts, briefs, templates ou scripts

Validacao minima:
- execute os comandos mais relevantes para o escopo alterado
- declare claramente o que foi executado, o que falhou e o que ficou pendente

Formato de saida:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- gaps restantes
```
