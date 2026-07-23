# Contributing

Contribuições — humanas ou assistidas por AI — seguem o mesmo fluxo governado.

## Antes de começar

1. Leia [AGENTS.md](AGENTS.md) e [skills/00-skill-index/SKILL.md](skills/00-skill-index/SKILL.md).
2. Prepare o ambiente: `make setup` e confirme a baseline com `make check-compliance` e `make test`.
3. Primeira contribuição? Siga [docs/journeys/01-primeira-contribuicao.md](docs/journeys/01-primeira-contribuicao.md).

## Abrindo uma tarefa

Use [docs/ai/task-input-format.md](docs/ai/task-input-format.md) para estruturar o pedido: objetivo, specs lidas, arquivos em escopo, entregáveis mínimos, validações obrigatórias e formato de saída. O contrato mínimo de contribuição está em [docs/ai/ai-contribution-contract.md](docs/ai/ai-contribution-contract.md).

Nenhuma feature entra sem spec governante. Se a spec não existir ou estiver ambígua, registre o gap em vez de implementar.

## Validando e abrindo PR

- Rode os checks aplicáveis: `make fmt-check`, `make lint`, `make vet`, `make test`, `make race`, `make test-security` e `pre-commit run --all-files`.
- Abra o PR com `scripts/create-pr-from-template.sh`; o corpo segue [.github/PULL_REQUEST_TEMPLATE.md](.github/PULL_REQUEST_TEMPLATE.md) e é validado no CI por `make check-pr-body`.
- Se alguma validação obrigatória não se aplicar, registre a exceção formal em [docs/ai/compliance-exceptions.md](docs/ai/compliance-exceptions.md) — exceções informais não valem.
