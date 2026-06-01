# 0001 - Record Architecture Decisions

- **Status**: `Accepted`
- **Data**: bootstrap

## Contexto

O `{{PROJECT_SLUG}}` segue Spec Driven Development rigoroso e governanca AI-first. Decisoes de arquitetura, governanca de release, ownership e tecnologia precisam de um registro estavel, rastreavel e imutavel que sobreviva a conversas, prompts e mudancas de equipe.

Specs descrevem o contrato de um modulo, mas nem toda decisao transversal (escolha de tecnologia, fronteira de escopo, politica de release) cabe em uma spec de modulo. Sem um registro dedicado, o "porque" de decisoes importantes se perde.

## Decisao

Adotar Architecture Decision Records (ADRs) em `docs/decisions/`, seguindo o formato e as regras de numeracao descritos em [README.md](README.md).

ADRs sao imutaveis apos `Accepted`; revisitar uma decisao exige um novo ADR que supersede o anterior.

## Opcoes consideradas

- **ADRs versionados no repo (escolhida)**: historico imutavel, rastreavel por commit, proximo de codigo e specs.
  - Pros: simples, duravel, revisavel em PR, sem dependencia externa.
  - Cons: exige disciplina de numeracao e de imutabilidade.
- **Registrar decisoes apenas em specs**: evita um diretorio novo.
  - Pros: menos artefatos.
  - Cons: mistura contrato de modulo com decisao transversal; decisoes de governanca ficam sem lar.
- **Wiki ou ferramenta externa**: centraliza fora do repo.
  - Pros: edicao facil.
  - Cons: drift em relacao ao codigo, sem revisao por PR, sem imutabilidade garantida.

## Consequencias

Positivas:

- Decisoes transversais ficam rastreaveis e auditaveis.
- O "porque" de cada decisao sobrevive a mudancas de contexto.
- Release governance e fronteiras de escopo passam a ter registro normativo.

Negativas:

- Exige disciplina para criar e superseder ADRs corretamente.
- Adiciona um passo de processo para decisoes nao-obvias.

## Criterios para reabrir

- Mudanca material no fluxo de governanca ou no formato de decisao do projeto.
- Necessidade de integrar ADRs com outra ferramenta de rastreabilidade.

## Referencias

- [docs/decisions/README.md](README.md)
- `AGENTS.md`
- `skills/22-release-versioning-governance/SKILL.md`
