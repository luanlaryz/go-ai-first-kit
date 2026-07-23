# AI Contribution Contract

Este contrato define a baseline operacional expandida para contribuicoes assistidas por AI na `{{PROJECT_SLUG}}`.
Ele complementa `AGENTS.md` e e aplicado pelo enforcement real do repositorio: `scripts/check-compliance.sh`, o template de PR e o CI.

## Regras normativas

1. Antes de editar codigo, ler `AGENTS.md`, a spec governante, a spec de diagnostico correspondente quando existir, `skills/00-skill-index/SKILL.md` e as skills aplicaveis.
2. Toda tarefa deve chegar com `objetivo`, `specs lidas` com a spec governante identificada, `arquivos em escopo`, `entregaveis minimos`, `validacoes obrigatorias`, `regras` e `formato de saida`.
3. Nenhuma feature pode ser implementada sem spec. Se a spec estiver ausente, incompleta ou ambigua, a execucao deve parar e o gap precisa ser declarado.
4. Mudanca de comportamento exige teste no mesmo patch, ou excecao formal explicitamente registrada em `docs/ai/compliance-exceptions.md`.
5. Mudanca de API publica exige atualizacao da spec correspondente, dos testes de contrato afetados e declaracao do impacto em docs ou superficie publica.
6. `pkg/*` deve continuar livre de dependencias em `internal/*`, exceto pela ponte publica permitida em `pkg/app`.
7. Docs, prompts, briefs, scripts e CI nao podem afirmar capacidade que o repositorio nao entrega de fato.
8. Toda entrega deve declarar, no minimo: `objetivo`, `specs lidas`, `skills aplicadas`, `arquivos alterados`, `comandos executados`, `testes executados` e `gaps restantes`.
9. Quando houver impacto em docs operacionais ou superficie publica, esse impacto deve ser declarado de forma explicita na tarefa, no PR e na documentacao alterada.
10. O checker nao pode depender de excecao silenciosa. Toda excecao ativa de compliance, teste ou enforcement deve ser auditavel no artefato formal do repositorio.

## Uso esperado

- Use `docs/ai/task-input-format.md` como formato padrao de entrada.
- Use `docs/ai/maturity-inventory.md` para localizar a baseline entregue e seus limites, sem tratá-lo como substituto de specs.
- Use `docs/ai/prompts/` como base de prompt operacional, nao como substituto da leitura das specs.
- Use `docs/ai/briefs/` para briefs reais e rastreaveis por spec.
- Use `docs/ai/compliance-exceptions.md` quando alguma exigencia obrigatoria precisar de excecao formal temporaria.

## Nao aceitavel

- implementar feature fora de spec
- deixar marcador vazio em assets operacionais
- tratar gap conhecido como capacidade implementada
- omitir spec governante, testes executados ou gaps restantes quando a mudanca tocar comportamento observavel
- usar excecao informal em comentario, script ou narrativa de PR sem registro auditavel
