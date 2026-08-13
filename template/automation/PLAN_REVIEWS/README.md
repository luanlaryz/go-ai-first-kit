# Plan Reviews

Este diretorio guarda os **vereditos de review de plano**. Um plano em `.cursor/plans/*.plan.md` so e
executavel quando existe aqui um artefato correspondente e o bloco `review:` do plano aponta para ele
com o hash correto.

Contrato completo: [`skills/27-dual-model-plan-review/SKILL.md`](../../skills/27-dual-model-plan-review/SKILL.md).
Preparacao operacional: [`skills/28-plan-review-autopilot/SKILL.md`](../../skills/28-plan-review-autopilot/SKILL.md).

## Nomeacao

- Primeiro review de um plano: `<plan-id>.md`, onde `<plan-id>` e o nome do plano sem `.plan.md`.
- Reviews posteriores do mesmo plano: `<plan-id>-v2.md`, `<plan-id>-v3.md`, e assim por diante.
- Artefato mergeado e **append-only**: nunca edite um veredito publicado; crie o sucessor versionado.

## Por que o hash existe

O campo `plan_sha256` cobre o **corpo** do plano (tudo depois do frontmatter). Se o corpo mudar depois
do review, o hash deixa de casar e `scripts/verify-plan-review.sh` reprova. Isso e o que impede o
padrao "aprova um plano pequeno, depois cresce o plano".

Por isso a preparacao do review nunca altera o corpo: `scripts/review-plan.py` escreve somente o
bloco `review:` no frontmatter.

## Comandos

```bash
make review-plan PLAN=.cursor/plans/<id>.plan.md              # prepara artefato
make review-plan-turn-based PLAN=.cursor/plans/<id>.plan.md   # prepara em modo turn-based
make verify-plan-reviews                                       # valida todos os planos alterados
scripts/verify-plan-review.sh --plan .cursor/plans/<id>.plan.md # valida um plano, estrito
```

`make verify-plan-reviews` nao chama modelo e nao exige API key: ele valida artefatos. E isso que
permite rodar o gate em CI.

## O que um veredito precisa conter

No minimo: identidade do plano, caminho de review (`DUAL_AUTOMATED`, `DUAL_TURN_BASED` ou
`OPERATOR_FALLBACK`), modelos planner e reviewer, data com timezone, TTL, o `plan_sha256` verbatim e o
resultado do checklist. Template copiavel em
[`skills/27-dual-model-plan-review/resources/verdict_template.md`](../../skills/27-dual-model-plan-review/resources/verdict_template.md).
