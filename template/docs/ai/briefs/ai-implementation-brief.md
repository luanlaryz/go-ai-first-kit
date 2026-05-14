# Brief real - AI implementation

## Objetivo

Executar uma mudanca implementativa assistida por AI com rastreabilidade por spec, validacao minima real e declaracao explicita de gaps restantes.

## Specs governantes

- `AGENTS.md`
- spec governante da tarefa
- spec de diagnostico correspondente, quando existir
- `docs/ai/ai-contribution-contract.md`
- `docs/ai/task-input-format.md`

## Escopo desta entrega

- ler primeiro a spec governante e os arquivos diretamente afetados
- implementar a menor mudanca correta dentro do escopo pedido
- atualizar testes no mesmo patch quando houver mudanca de comportamento
- sincronizar docs operacionais, prompts, briefs, checker, Makefile ou CI quando o fluxo operacional for parte da mudanca

## Fora do escopo

- introduzir feature fora de spec
- reabrir arquitetura sem necessidade observavel
- declarar capacidade operacional sem artefato correspondente

## Evidencias esperadas

- arquivos alterados listados de forma objetiva
- comandos e testes executados declarados de forma rastreavel
- impacto em docs ou superficie publica declarado quando existir
- gaps restantes registrados de forma explicita
