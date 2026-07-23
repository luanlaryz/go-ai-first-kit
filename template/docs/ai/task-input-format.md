# Task Input Format

Use este formato quando abrir uma tarefa para Codex, Cursor ou outro agente.
Ele complementa `AGENTS.md` e reduz drift entre prompt, spec e resultado entregue.

## Campos obrigatorios

- `Objetivo`: resultado concreto esperado.
- `Specs lidas`: lista das specs e docs obrigatorios de leitura, com a spec governante identificada explicitamente.
- `Arquivos em escopo`: arquivos ou diretorios que o agente deve inspecionar primeiro.
- `Entregaveis minimos`: artefatos que precisam existir ao final.
- `Impacto esperado em docs ou superficie publica`: declare `none` quando nao houver impacto relevante.
- `Validacoes obrigatorias`: comandos, testes ou checks a executar.
- `Regras`: restricoes de escopo, arquitetura ou operacao.
- `Formato de saida`: campos que a resposta final deve listar.

## Template

```md
Leia primeiro:
- AGENTS.md
- specs/<spec-governante>.md
- specs/<spec-diagnostico-correspondente>.md # quando a trilha tiver spec de diagnostico
- specs/010-feature-matrix.md # quando tocar status, cobertura ou paridade
- arquivos e docs adicionais realmente necessarios

Objetivo:
Descreva o resultado observavel esperado.

Specs lidas:
- AGENTS.md
- specs/<spec-governante>.md
- specs/<spec-diagnostico-correspondente>.md # quando existir
- outros docs realmente necessarios

Arquivos em escopo:
- caminhos que precisam ser lidos antes da implementacao

Entregaveis minimos:
1. artefato ou mudanca 1
2. artefato ou mudanca 2

Impacto esperado em docs ou superficie publica:
- none

Validacoes obrigatorias:
- comando ou teste 1
- comando ou teste 2

Regras:
- limite de escopo
- restricao arquitetural
- criterio de auditabilidade
- se alguma validacao obrigatoria nao se aplicar, registrar excecao formal em `docs/ai/compliance-exceptions.md`

Formato de saida:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- gaps restantes
```

## Quando detalhar mais

- Se houver mudanca de API publica, cite a spec do modulo, os testes de contrato afetados e o impacto em docs ou superficie publica.
- Se houver excecao formal em vez de teste ou check, declare a justificativa no input e vincule `docs/ai/compliance-exceptions.md`.
- Se o escopo tocar CI, docs operacionais, prompts ou briefs, liste os artefatos reais que devem ser sincronizados no mesmo patch.
