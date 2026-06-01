# Changelog and Release Notes Policy

Politica curta para manter `CHANGELOG.md` e release notes coerentes com a classificacao SemVer do `{{PROJECT_SLUG}}`.

Traceabilidade:

- `docs/release-versioning-policy.md`
- `skills/15-devx-ci-precommit-changelog/SKILL.md`
- `skills/21-documentation-open-source/SKILL.md`
- `skills/22-release-versioning-governance/SKILL.md`

## 1. `CHANGELOG.md`

O repositorio usa Keep a Changelog com secao `[Unreleased]`.

Regras:

1. toda mudanca release-worthy deve deixar rastro em `CHANGELOG.md`
2. a entrada deve refletir o impacto real ao consumidor
3. docs-only ou spec-only podem permanecer em `[Unreleased]` sem forcar release proprio
4. nao usar changelog para esconder breaking change sob texto neutro
5. quando uma mudanca e `major`, isso deve aparecer de forma explicita na linguagem da entrada

Secoes preferenciais:

- `Added`
- `Changed`
- `Deprecated`
- `Removed`
- `Fixed`
- `Security`

## 2. Release notes

Todo release diferente de `NO_RELEASE` deve ter release notes curtas com:

1. versao recomendada
2. classificacao SemVer
3. stage de prerelease (`alpha.N`, `beta.N`, `rc.N`) quando a distribuicao for `PRERELEASE_TAG`
4. resumo executivo
5. mudancas publicas principais
6. avaliacao de compatibilidade
7. status dos gates relevantes
8. riscos, limites conhecidos ou itens `UNVERIFIED`
9. instrucao de upgrade quando aplicavel

Quando as release notes forem de prerelease, elas devem dizer se o candidato e `alpha`, `beta` ou `rc`, por que esse stage foi escolhido e quais riscos impedem ou ainda nao justificam `STABLE_TAG`.

## 3. Relacao entre changelog e release notes

Os dois artefatos devem contar a mesma historia sobre:

- o que mudou
- se houve ou nao breaking change
- se o release e estavel, prerelease ou hotfix
- qual stage de prerelease foi escolhido, quando aplicavel
- quais riscos ainda existem

Se changelog, release notes, feature matrix e reports divergirem, o artefato menos aderente ao comportamento real deve ser corrigido.

## 4. Modelos minimos

### Changelog

```md
## [Unreleased]

### Changed
- Release governance docs adicionadas para formalizar classificacao SemVer, checklist de release e politica de changelog/release notes. (spec governante)
```

### Release notes

```md
# Release Notes - vX.Y.Z

- Classification: `patch | minor | major`
- Distribution: `STABLE_TAG | PRERELEASE_TAG | HOTFIX_TAG`
- Prerelease stage: `alpha.N | beta.N | rc.N | N/A`
- Public surface touched: `pkg/*`, HTTP, OpenAPI, starter, docs surface

## Summary
- _resumo curto_

## Compatibility
- _impacto compativel ou breaking_

## Gates
- `go test ./...`: `PASS | FAIL | UNVERIFIED`
- `go test -race ./...`: `PASS | FAIL | UNVERIFIED`
- `smoke/conformance`: `PASS | FAIL | UNVERIFIED`

## Residual Risks
- _riscos ou blockers_
```

## 5. Alinhamento com skills

Esta policy fica alinhada com:

- skill 15 para higiene de changelog e processo de release
- skill 21 para rastreabilidade documental
- skill 22 para classificacao SemVer, checklist e distribuicao
