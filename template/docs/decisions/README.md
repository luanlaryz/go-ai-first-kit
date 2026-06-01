# Architecture Decision Records (ADRs)

Este diretorio contem os ADRs (Architecture Decision Records) governantes do `{{PROJECT_SLUG}}`.

Cada ADR documenta uma decisao tecnica ou de governanca com contexto, opcoes consideradas, decisao adotada e consequencias.

## Formato

ADRs seguem template simplificado (MADR-like):

- Titulo
- Status (`Proposed` / `Accepted` / `Deprecated` / `Superseded by ...`)
- Contexto
- Decisao
- Opcoes consideradas (com pros e cons)
- Consequencias (positivas e negativas)
- Criterios para reabrir
- Referencias

## Numeracao

ADRs sao numerados sequencialmente em formato `NNNN-titulo-kebab-case.md`. A faixa atual:

- `0001`: registra a propria pratica de ADRs (bootstrap).
- `0002-0009`: reservadas para ADRs de bootstrap futuros.
- `0010+`: ADRs documentando decisoes registradas via `interactive_sdd_autopilot` ou `phase_autopilot`.

## Indice

- [0001 - Record Architecture Decisions](0001-record-architecture-decisions.md) - `Accepted`

## Quando criar um ADR

Crie um ADR quando:

- Houver decisao de arquitetura nao-obvia que afete contrato publico ou comportamento operacional.
- Houver decisao de governanca (release, ownership, deprecation, security policy).
- Houver decisao de tecnologia que limite caminhos futuros.
- Houver decisao explicitamente requisitada por spec, plano ou stop condition.

Decisoes pequenas ou que sao consequencia direta de uma spec aprovada nao exigem ADR; o registro vai na propria spec ou no report da trilha.

## Quando atualizar um ADR

ADRs sao **imutaveis** apos `Accepted`. Para revisitar uma decisao:

1. Criar novo ADR com proximo numero sequencial.
2. Marcar status do novo ADR com referencia ao antigo: `Status: Accepted (supersedes 0001)`.
3. Atualizar status do antigo para `Superseded by 0NNN`.

Excecao: ADRs em `Proposed` podem ser editados ate aprovacao. Apos `Accepted`, sao imutaveis.
