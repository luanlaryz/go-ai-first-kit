# Contributing

Contribuições — humanas ou assistidas por AI — seguem o mesmo fluxo governado.

## Antes de começar

1. Leia [AGENTS.md](AGENTS.md) e [skills/00-skill-index/SKILL.md](skills/00-skill-index/SKILL.md).
2. Prepare o ambiente: `make setup` e confirme a baseline com `make check-compliance` e `make test`.
3. Primeira contribuição? Siga [docs/journeys/01-primeira-contribuicao.md](docs/journeys/01-primeira-contribuicao.md).

### Dependências de desenvolvimento

Além do Go e do `git`:

- `python3`: usado pelos hooks de segurança em `.cursor/hooks/`, pelo gate de review de plano (`scripts/review-plan.py`) e pelos gates que fazem parsing estruturado. Ausência de `python3` faz os hooks degradarem fail-closed (`ask`) e reprova os gates que dependem dele — nunca passa em silêncio.
- `gh` (GitHub CLI): opcional, mas sem ele o gate de mudança governada não consegue provar que não existe PR aberto, e a varredura de colisão entre sessões paralelas perde uma das três fontes.
- `pre-commit`: para rodar `pre-commit run --all-files`.

## Abrindo uma tarefa

Use [docs/ai/task-input-format.md](docs/ai/task-input-format.md) para estruturar o pedido: objetivo, specs lidas, arquivos em escopo, entregáveis mínimos, validações obrigatórias e formato de saída. O contrato mínimo de contribuição está em [docs/ai/ai-contribution-contract.md](docs/ai/ai-contribution-contract.md).

Nenhuma feature entra sem spec governante. Se a spec não existir ou estiver ambígua, registre o gap em vez de implementar.

## Validando e abrindo PR

- Rode `make pre-pr` com exit 0 antes de `git push` ou `gh pr create`. Ele agrega formatação, lint, vet, `make guardrails`, testes, race, baseline de segurança e os gates de governança. Checks parciais não substituem o alvo agregado.
- Se a mudança toca `pkg/**`, `internal/**`, `api/**` ou `migrations/**`, ela é governada: exige item `BLG-NNNN` no backlog, par dual-spec e plano com review aprovado. Comece por [skills/29-governed-change-workflow/SKILL.md](skills/29-governed-change-workflow/SKILL.md).
- Assuma sessões paralelas: reserve `BLG-NNNN`, `specs/NNN` e `migrations/NNNN` no último momento antes do commit e siga [skills/31-parallel-session-coordination/SKILL.md](skills/31-parallel-session-coordination/SKILL.md).
- Abra o PR com `scripts/create-pr-from-template.sh`; o corpo segue [.github/PULL_REQUEST_TEMPLATE.md](.github/PULL_REQUEST_TEMPLATE.md) e é validado no CI por `make check-pr-body`.
- Se alguma validação obrigatória não se aplicar, registre a exceção formal em [docs/ai/compliance-exceptions.md](docs/ai/compliance-exceptions.md) — exceções informais não valem.
