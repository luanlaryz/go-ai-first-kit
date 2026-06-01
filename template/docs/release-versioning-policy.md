# Release Versioning Policy

Normatiza a governanca de release e versionamento do `{{PROJECT_SLUG}}`.

Traceabilidade:

- `specs/020-repository-architecture.md`
- `skills/22-release-versioning-governance/SKILL.md`
- `skills/15-devx-ci-precommit-changelog/SKILL.md`

## 1. Objetivo

Definir como o `{{PROJECT_SLUG}}` classifica mudancas, protege consumidores e recomenda um caminho de distribuicao sem maquiar compatibilidade.

## 2. Unidade de versionamento

A unidade de versionamento e o modulo Go do repositorio.

A biblioteca segue SemVer:

- `major`: mudanca publica incompativel
- `minor`: adicao publica retrocompativel
- `patch`: correcao publica retrocompativel

## 3. Linha atual

A baseline atual permanece em `v0.x.y`.

Regras para `v0.x.y`:

- breaking change continua sendo breaking change
- a classificacao continua sendo `major`
- nao e permitido esconder ruptura sob `patch` ou `minor`
- `v1.0.0` so pode ser recomendado por trilha futura com evidencia diagnostica explicita de estabilidade intencional da API publica, starter, docs e governanca de release

## 4. Superficie publica

Contam como superficie publica:

1. contratos exportados em `pkg/*`
2. `pkg/app` como ponte publica para o runtime
3. comportamento observavel documentado por spec aprovada e ja implementado
4. erros publicos, envelopes, semantica de cancelamento e ordem de execucao quando fizerem parte do contrato
5. endpoints HTTP documentados como suportados
6. schemas e contratos OpenAPI publicados pelo repositorio
7. `/docs` e `/openapi.yaml` quando publicados como surface oficial de documentacao
8. starter flow, examples oficiais suportados e configuracao documentada como caminho canonico

Nao contam como superficie publica, por si so:

1. `internal/*`
2. `test/*`, fixtures e helpers de teste
3. comments-only changes
4. reports historicos e notas auxiliares sem contrato publico proprio
5. refactors internos sem impacto observavel
6. specs ainda nao implementadas

## 5. Classificacao de impacto

### `no release`

Use quando a mudanca:

- altera apenas specs, reports, docs, comments ou feature matrix
- altera apenas `internal/*` sem impacto observavel
- altera apenas testes, smoke checks ou fixtures
- reconcilia OpenAPI, changelog ou docs com comportamento ja existente

### `patch`

Use quando a mudanca:

- corrige bug sem quebrar a superficie publica
- corrige validacao, erro, docs surface ou runtime de forma compativel
- corrige OpenAPI ou docs para refletir comportamento ja existente
- endurece comportamento sem exigir adaptacao do consumidor

### `minor`

Use quando a mudanca:

- adiciona capability nova e retrocompativel
- adiciona endpoint, tipo exportado, configuracao suportada ou ponto de extensao sem quebrar consumidores
- amplia comportamento existente de forma opcional e compativel

### `major`

Use quando a mudanca:

- remove, renomeia ou altera assinatura exportada em `pkg/*`
- muda endpoint, payload, schema, SSE ou contrato OpenAPI de forma incompativel
- muda o starter flow ou setup canonico de forma que force ajuste do consumidor
- preserva simbolos, mas muda semantica observavel a ponto de quebrar consumidores

## 6. Prerelease

Use prerelease quando a mudanca e real e utilizavel, mas ainda nao merece promessa de estabilidade ampla.

Tags recomendadas:

- `vX.Y.Z-alpha.N`
- `vX.Y.Z-beta.N`
- `vX.Y.Z-rc.N`

Estagios normativos:

- `alpha.N`: candidato inicial para feedback tecnico, validacao antecipada ou formato ainda instavel. Use quando a capability existe e pode ser exercitada, mas ainda pode mudar shape, contrato de API, payload, docs ou comportamento observavel antes de uma linha estavel.
- `beta.N`: candidato mais completo para integracao controlada. Use quando o comportamento principal ja esta presente, os contratos esperados estao documentados, os gaps conhecidos estao listados e a expectativa e refinamento localizado, nao redesenho amplo.
- `rc.N`: candidato a release estavel. Use quando nao ha mudanca funcional ampla planejada antes da tag final, os gates relevantes estao verdes ou justificados, e so se esperam fixes localizados, ajustes documentais ou correcao de regressao encontrada na validacao final.

`alpha`, `beta` e `rc` sao estagios de maturidade do candidato, nao substitutos para a classificacao SemVer. Eles nao podem mascarar breaking change, gate ausente, incompatibilidade conhecida ou blocker aberto. Se a mudanca for breaking, ela continua sendo `major`; se um gate critico estiver ausente, a recomendacao deve registrar `UNVERIFIED`, `PRERELEASE_TAG` ou `HOLD_RELEASE` conforme o risco.

Use `PRERELEASE_TAG` quando houver qualquer uma das condicoes abaixo:

1. capability nova ainda sem fechamento diagnostico suficiente
2. gates relevantes em `UNVERIFIED`
3. gaps conhecidos que nao bloqueiam experimentacao, mas ainda bloqueiam estabilidade
4. necessidade de feedback antes da linha estavel

## 7. Hotfix

Use `HOTFIX_TAG` somente quando todas as condicoes forem verdadeiras:

1. existe linha estavel ja publicada
2. o problema e consumidor-facing e urgente
3. a mudanca cabe claramente em `patch`
4. o blast radius e pequeno e entendido
5. o risco de esperar o proximo release normal e inaceitavel

Hotfix nao serve para:

- capability nova
- breaking change
- ajuste interno sem urgencia externa material

## 8. Recomendacao de distribuicao

Toda avaliacao deve produzir exatamente um caminho:

- `NO_RELEASE`
- `PRERELEASE_TAG`
- `STABLE_TAG`
- `HOTFIX_TAG`
- `HOLD_RELEASE`

Mapeamento operacional:

- docs/spec/report only -> `NO_RELEASE`
- capability compativel e gates verdes -> `STABLE_TAG`
- capability util, mas ainda incompleta ou parcialmente verificada -> `PRERELEASE_TAG`
- fix urgente em linha estavel -> `HOTFIX_TAG`
- blockers ou gates criticos ausentes -> `HOLD_RELEASE`

## 9. Gates minimos para release estavel

Antes de recomendar `STABLE_TAG`, verificar:

- `go test ./...`
- `go test -race ./...`
- smoke checks e conformance relevantes
- alinhamento entre docs, OpenAPI e comportamento real quando a surface correspondente mudar
- ausencia de blocker aberto para a capability afetada
- coerencia entre feature matrix, reports e historia contada pelo release

Se algum gate nao puder ser verificado, registrar `UNVERIFIED`.

## 10. Imutabilidade

Versao publicada nao pode ser alterada em lugar.

Toda correcao posterior exige:

- nova versao
- novo tag
- changelog e release notes coerentes com a nova decisao

## 11. Alinhamento com a skill 22

Esta policy operacionaliza no repositorio as mesmas fronteiras usadas pela skill `22-release-versioning-governance`:

- SemVer-first
- `pkg/*` e HTTP/OpenAPI documentados como superficie publica
- `internal/*` como nao-publico salvo promocao explicita
- `NO_RELEASE`, `PRERELEASE_TAG`, `STABLE_TAG`, `HOTFIX_TAG` e `HOLD_RELEASE`
- uso de reports e diagnosticos como evidencia de maturidade
