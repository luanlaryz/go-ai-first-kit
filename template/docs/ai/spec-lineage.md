# Linhagem de Specs AI — {{PROJECT_TITLE}}

Este starter não reproduz a história de specs de outro repositório. A
linhagem começa nos arquivos presentes em `specs/` e cresce junto com o
projeto. Nunca trate o número, o nome ou o report de uma spec ausente como
evidência de implementação.

## Specs iniciais

- `specs/000-project-mission.md` define a missão e as regras globais.
- `specs/001-non-goals.md` define limites explícitos de escopo.
- `specs/010-feature-matrix.md` registra cobertura, prioridade, risco e
  paridade quando aplicável.
- `specs/020-repository-architecture.md` define fronteiras de arquitetura.
- `specs/680-phase-0-bootstrap-foundation.md` e
  `specs/681-phase-0-bootstrap-foundation-diagnosis.md` formam o par inicial
  de construção e diagnóstico do roadmap.

## Como criar a próxima trilha

1. Leia [AGENTS.md](../../AGENTS.md) e a skill de spec architect.
2. Decida se a mudança pertence a uma spec existente ou exige nova trilha.
3. Para uma trilha nova, crie uma spec de construção e outra de diagnóstico
   antes de implementar comportamento.
4. Registre os paths no estado e no report que governam a trilha.
5. Atualize a feature matrix, backlog ou ADR quando a mudança alterar esses
   contratos.

Specs futuras são candidatas a criar e só se tornam normativas após existirem
como arquivos versionados.

## Evidência e decisões

O enforcement real está nos seguintes artefatos:

- `scripts/check-compliance.sh`
- `.pre-commit-config.yaml`
- `.github/workflows/ci.yml`
- `.github/PULL_REQUEST_TEMPLATE.md`
- `docs/ai/ai-contribution-contract.md`
- `docs/ai/compliance-exceptions.md`
- `test/security/`

Use `docs/reports/` para reports produzidos pelo próprio projeto e
`docs/decisions/` para decisões de arquitetura ou governança não óbvias.
