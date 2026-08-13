# Cursor Plan Templates

Templates de plano para as mudancas mais comuns em `{{PROJECT_SLUG}}`.

## Como usar

1. Copie o conteudo do template adequado para um novo arquivo em `.cursor/plans/<nome>_<hash>.plan.md`.
2. Preencha o frontmatter real (`id`, `status`, `risk`, `area`) e todas as secoes.
3. Prepare o gate de review com [`skills/28-plan-review-autopilot/SKILL.md`](../../../skills/28-plan-review-autopilot/SKILL.md) antes de executar.
4. Rode `make verify-plan-reviews` e so execute apos veredito `APPROVED` ou `APPROVED_WITH_CHANGES`.

## Por que `.plan.template.md`

Os templates usam a extensao `.plan.template.md` e mantem o frontmatter de exemplo dentro de blocos de
codigo. Isso evita que `scripts/verify-plan-review.sh` (que varre `.cursor/plans/*.plan.md`) trate um
template como plano executavel sem review valido.

## Templates disponiveis

- `adapter-port.plan.template.md`: adapter de borda mapeando sistema externo para DTO canonico.
- `openapi-change.plan.template.md`: mudanca em contrato publicado.
- `migration-change.plan.template.md`: mudanca de schema com migrations versionadas.
- `runtime-change.plan.template.md`: mudanca de comportamento no runtime interno.
- `worker-change.plan.template.md`: mudanca em consumo assincrono, retry, DLQ ou idempotencia.
- `golden-case.plan.template.md`: novo caso golden ou de conformidade.
