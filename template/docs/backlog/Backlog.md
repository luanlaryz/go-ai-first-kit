# {{PROJECT_TITLE}} Backlog

## Objetivo

Este arquivo é o backlog canônico de triagem para itens governados por SDD no `{{PROJECT_SLUG}}`. Ele registra pedidos, gaps, bugs, diagnósticos, recomendações e ideias depois de passarem pela triagem da skill [`skills/26-backlog-item-intake/SKILL.md`](../../skills/26-backlog-item-intake/SKILL.md).

Este backlog não substitui specs, reports, testes ou decisão humana. Todo item implementável exige spec de construção e spec de diagnóstico antes de qualquer implementação.

## Escopo e fontes

Fontes em ordem de precedência:

1. [AGENTS.md](../../AGENTS.md)
2. [specs/000-project-mission.md](../../specs/000-project-mission.md), [specs/001-non-goals.md](../../specs/001-non-goals.md), [specs/010-feature-matrix.md](../../specs/010-feature-matrix.md), [specs/020-repository-architecture.md](../../specs/020-repository-architecture.md)
3. Reports em `docs/reports/**` e fontes auxiliares em `docs/backlog/**`

Regras de escopo:

- `docs/backlog/Backlog.md` não substitui specs, reports, testes ou decisão humana.
- Não migrar fontes históricas em lote para itens `BLG-*`.
- Não registrar capability de {{UPSTREAM_OPS_NAME}}, hosted service, Admin UI, control plane, RBAC, billing, quotas ou responsabilidade de aplicação consumidora como item implementável.
- Todo item implementável deve exigir spec de construção e spec de diagnóstico antes de qualquer implementação.

## Lifecycle do backlog

Estados canônicos:

- `candidate`: entrada registrada, ainda sem triagem completa.
- `ready_for_spec`: recorte suficiente para criar ou amendar specs.
- `ready_for_implementation`: specs governantes existem e diagnóstico está definido.
- `blocked`: falta decisão, evidência, spec ou ambiente.
- `in_progress`: trilha SDD em execução.
- `done`: item concluído com report `PASS`.
- `rejected`: fora de escopo, duplicado ou incompatível com non-goals.

## Critérios de classificação

A classificação de cada item (status, tipo, prioridade, severidade, valor, complexidade, risco, decisão de escopo, impacto em API pública e risco de breaking change) segue [`skills/26-backlog-item-intake/references/backlog-classification-rules.md`](../../skills/26-backlog-item-intake/references/backlog-classification-rules.md).

## Índice de priorização

Liste aqui apenas itens `BLG-*` reais, ordenados por prioridade (`P0` a `P3`), depois bloqueados, rejeitados e concluídos. Dentro da mesma prioridade, prefira maior risco e maior valor.

Nenhum item `BLG-*` registrado ainda.

## Itens

Insira cada item usando [`skills/26-backlog-item-intake/references/backlog-item-template.md`](../../skills/26-backlog-item-intake/references/backlog-item-template.md). O primeiro item recebe o ID `BLG-0001`.

Nenhum item `BLG-*` registrado ainda.

## Backlogs auxiliares e fontes históricas

Cite aqui backlogs herdados, planos importados ou relatórios de pesquisa que ainda não foram triados. Eles não viram itens implementáveis sem triagem individual.

Nenhuma fonte auxiliar registrada ainda.

## Itens candidatos a triagem

Liste candidatos ainda não promovidos a `BLG-*`. Cada candidato deve registrar fonte, recorte, decisão atual, motivo para não criar `BLG-*` ainda e próxima ação de triagem.

Nenhum candidato registrado ainda.
