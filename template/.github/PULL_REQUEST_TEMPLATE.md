## Objetivo

Descreva o resultado observavel desta mudanca.

## Backlog

Item governado desta mudanca, ou `none` quando o PR nao toca `pkg/**`, `internal/**`, `api/**` nem `migrations/**`.

- `BLG-NNNN` (`docs/backlog/Backlog.md`), com `Status` em `ready_for_implementation`, `in_progress` ou `done`

## Specs lidas

- [ ] `AGENTS.md`
- [ ] spec(s) governante(s) citadas abaixo

Liste as specs e secoes aplicaveis. Escopo governado exige o par dual-spec (construcao `NNN` + diagnostico `NNN+1`):

- `specs/...`

## Autopilot

Plano revisado desta mudanca, ou `none` fora do escopo governado.

- Plano: `.cursor/plans/<id>.plan.md`
- Veredito: `automation/PLAN_REVIEWS/<id>.md`
- `make verify-plan-reviews`: `PASS`

## Regressao

Cobertura regressiva desta mudanca, ou `none` fora do escopo governado.

- `BLG-NNNN` + cenario alvo em `test/<pacote>::<cenario>`

## Skills aplicadas

- `skills/...`

## Agent Readiness

- Report: `none` ou `docs/reports/agent-readiness/...`
- Source: `none` ou comando/report `@kodus/agent-readiness`
- {{PROJECT_TITLE}}-filtered result: `READY`, `READY_WITH_WARNINGS`, `NOT_READY` ou `not applicable`
- PR impact: `BLOCKS_PR`, `DOES_NOT_BLOCK_PR` ou `not applicable`

Regra de bloqueio: `optional_for_{{PROJECT_SLUG}}` e `out_of_scope_for_{{PROJECT_SLUG}}` nao bloqueiam; apenas `worth_it_for_{{PROJECT_SLUG}}` que afete esta PR e tenha `PR impact = BLOCKS_PR` bloqueia merge.

Achados relevantes, recomendacoes opcionais e achados fora de escopo:

- `none`, ou liste apenas o que afeta esta PR

## Arquivos alterados

- liste os arquivos principais alterados neste PR

## Impacto em docs e superficie publica

- `none`, ou descreva o impacto relevante quando existir

## Comandos executados

- liste comandos executados alem dos checks padrao, quando houver

## Testes executados

- [ ] `make fmt-check`
- [ ] `make lint`
- [ ] `make vet`
- [ ] `make check-compliance`
- [ ] `make test`
- [ ] testes adicionais relevantes

## Excecoes formais

- `none`, ou referencie `docs/ai/compliance-exceptions.md`

## Gaps restantes

- descreva gaps, tradeoffs ou itens fora do escopo

## Checklist

- [ ] Li `AGENTS.md` e as specs governantes antes de editar
- [ ] Atualizei testes ou declarei excecao formal quando a mudanca tocou comportamento
- [ ] Atualizei docs e spec correspondente quando a mudanca tocou API publica ou fluxo operacional
- [ ] Nao deixei marcador vazio em assets operacionais
- [ ] Declarei impacto em docs ou superficie publica, ou marquei `none`
- [ ] A descricao deste PR lista gaps restantes de forma explicita
