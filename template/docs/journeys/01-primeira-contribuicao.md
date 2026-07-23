# Jornada 01 — Primeira contribuição

Para quem acabou de gerar ou clonar o `{{PROJECT_SLUG}}` e vai fazer a primeira mudança.

## 1. Validar a baseline

```bash
make setup
make check-compliance
make test
```

Evidência esperada: `check-compliance: ok` e testes verdes. Se algo falhar aqui, a baseline foi alterada ou o ambiente não atende Go 1.26.4+; corrija antes de qualquer contribuição.

## 2. Ler o mínimo normativo

Nesta ordem:

1. [AGENTS.md](../../AGENTS.md) — constituição: fontes de verdade, workflow obrigatório e formato de resposta.
2. [skills/00-skill-index/SKILL.md](../../skills/00-skill-index/SKILL.md) — roteamento para a skill que governa sua mudança.
3. A spec que cobre sua mudança em [specs/](../../specs/). Em um projeto recém-gerado existem as specs 000, 001, 010, 020 e o par 680/681 da fase 0.
4. [docs/ai/capabilities.md](../ai/capabilities.md) — o que existe, o que é processo e o que ainda não existe.

## 3. Estruturar a tarefa

Use [docs/ai/task-input-format.md](../ai/task-input-format.md): objetivo, specs lidas, arquivos em escopo, entregáveis mínimos, validações e formato de saída. Se a mudança não tiver spec governante, pare e registre o gap — não implemente fora de spec.

## 4. Implementar e validar

- Edite apenas arquivos no escopo declarado.
- Mudança de comportamento exige teste no mesmo patch.
- Antes de commitar código Go: `make fmt`.
- Valide com os comandos aplicáveis de [docs/ai/ops-guide.md](../ai/ops-guide.md) (`make fmt-check`, `make lint`, `make vet`, `make test`, `make race`, `make test-security`).

## 5. Abrir o PR

Use `scripts/create-pr-from-template.sh` para gerar o corpo a partir de [.github/PULL_REQUEST_TEMPLATE.md](../../.github/PULL_REQUEST_TEMPLATE.md). O CI valida o corpo com `make check-pr-body`.

## Encerramento

- Evidência esperada: PR com specs lidas, skills aplicadas, comandos e testes executados, e gaps restantes declarados.
- Bloqueia quando: falta spec governante, um check obrigatório falha, ou a mudança exige decisão humana de arquitetura.
- Próxima ação humana: revisar com o checklist de [docs/feature-request-lifecycle.md](../feature-request-lifecycle.md); se a mudança merecer trilha completa, seguir a [jornada 03](03-backlog-e-trilha-sdd.md).
