# Release Checklist

Checklist operacional curto para recomendacoes de release do `{{PROJECT_SLUG}}`.

Traceabilidade:

- `docs/release-versioning-policy.md`
- `skills/22-release-versioning-governance/SKILL.md`
- `skills/22-release-versioning-governance/resources/release_checklist.md`

## 1. Cabecalho obrigatorio

Preencher antes de decidir:

- Candidate version:
- Classification: `no release | patch | minor | major`
- Distribution path: `NO_RELEASE | PRERELEASE_TAG | STABLE_TAG | HOTFIX_TAG | HOLD_RELEASE`
- Prerelease stage, if applicable: `alpha.N | beta.N | rc.N | N/A`
- Public surface touched:
- Current line:
- Evidence sources:

## 2. Gates obrigatorios

Marcar cada item como `PASS`, `FAIL` ou `UNVERIFIED`.

| Gate | Status | Evidencia |
| --- | --- | --- |
| Sumario da mudanca explicito | | |
| Modulos e arquivos tocados listados | | |
| Impacto na superficie publica avaliado | | |
| Avaliacao de breaking change concluida | | |
| Classificacao SemVer escolhida | | |
| Caminho de distribuicao escolhido | | |
| Stage `alpha`/`beta`/`rc` justificado quando `PRERELEASE_TAG` | | |
| `go test ./...` | | |
| `go test -race ./...` | | |
| Smoke/conformance relevantes | | |
| Docs/OpenAPI alinhados quando aplicavel | | |
| Feature matrix/reports sem contradicao material | | |
| Changelog rascunhado | | |
| Release notes rascunhadas | | |
| Riscos residuais listados | | |

## 3. Regras de decisao

1. `STABLE_TAG` exige ausencia de blocker conhecido e gates criticos sem `FAIL`.
2. `PRERELEASE_TAG` pode conviver com `UNVERIFIED`, mas isso deve aparecer nas release notes.
3. `HOTFIX_TAG` exige justificativa de urgencia e escopo de `patch`.
4. Docs/spec/report only tendem a `NO_RELEASE`.
5. Se a evidencia for insuficiente para prometer estabilidade, usar `HOLD_RELEASE` ou `PRERELEASE_TAG`.
6. `alpha`, `beta` e `rc` sao maturidade do candidato; eles nao substituem classificacao SemVer nem escondem breaking change.

## 4. Saida minima

Toda avaliacao deve registrar:

- decisao final
- versao recomendada
- superficie publica tocada
- justificativa do impacto
- checklist preenchido
- riscos ou blockers remanescentes
