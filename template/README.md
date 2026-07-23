# {{PROJECT_TITLE}}

{{PROJECT_DESCRIPTION}}

## Instalação e primeiro uso

Pré-requisito: Go 1.26.4+.

```bash
make setup
make check-compliance
make test
```

Evidência esperada: `check-compliance: ok` e testes verdes. Depois, se tiver a CLI do kit instalada, rode `gakit diagnose --path . --report-only` para ver a baseline AI-first por pilar.

## AI-first governance

Antes de editar, leia [AGENTS.md](AGENTS.md), [skills/00-skill-index/SKILL.md](skills/00-skill-index/SKILL.md) e a spec governante em [specs/](specs/). O repositório segue Spec Driven Development com dual-spec: nenhuma trilha nova começa sem spec de construção e spec de diagnóstico.

## Navegação

- [docs/README.md](docs/README.md): central de documentação, organizada por papel.
- [docs/ai/capabilities.md](docs/ai/capabilities.md): catálogo de capacidades, como ativá-las e seus limites.
- [docs/journeys/README.md](docs/journeys/README.md): jornadas passo a passo (primeira contribuição, fase do roadmap, backlog e trilha SDD, release).
- [docs/ai/maturity-inventory.md](docs/ai/maturity-inventory.md): baseline entregue, evidência a produzir e gaps.

## O que este starter ainda não contém

Um projeto recém-gerado não tem `pkg/`, `internal/`, `examples/`, API pública nem contrato OpenAPI: essas superfícies nascem apenas sob spec aprovada. Findings de diagnóstico sobre esses itens são gaps declarados por design, não defeitos.
