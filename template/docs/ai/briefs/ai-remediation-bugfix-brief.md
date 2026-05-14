# Brief real - AI remediation ou bugfix

## Objetivo

Corrigir um bug ou gap operacional ja diagnosticado, restaurando aderencia a spec com a menor mudanca correta e com evidencia real de fechamento.

## Specs governantes

- `AGENTS.md`
- spec governante da correcao
- spec de diagnostico que evidenciou o problema, quando existir
- `docs/ai/ai-contribution-contract.md`
- `docs/ai/compliance-exceptions.md`, quando houver excecao ativa relacionada

## Escopo desta entrega

- reproduzir ou citar a evidencia do problema antes da correcao
- limitar a mudanca ao bug ou gap diagnosticado
- atualizar testes, checker, docs ou CI quando eles fizerem parte da causa ou da confirmacao do fechamento
- declarar o que foi resolvido e o que continua pendente

## Fora do escopo

- usar remediation para abrir feature nova
- remover exigencia de compliance sem justificativa arquitetural
- esconder pendencia atras de excecao nao auditavel

## Evidencias esperadas

- reproducao ou referencia concreta do problema
- arquivos alterados e comandos executados
- testes que confirmam a correcao
- gaps restantes ou excecoes formais ainda ativas
