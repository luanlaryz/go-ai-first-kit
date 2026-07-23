# Jornada 04 — Release, changelog e decisões

Para quem vai recomendar um release, atualizar o changelog ou registrar uma decisão de arquitetura. Fontes normativas: [release-versioning-policy.md](../release-versioning-policy.md), [release-notes-policy.md](../release-notes-policy.md), [release-checklist.md](../release-checklist.md) e a skill [22](../../skills/22-release-versioning-governance/SKILL.md).

## 1. Classificar a mudança

Determine a superfície pública tocada e a classificação SemVer (`no release`, `patch`, `minor`, `major`) conforme a política de versionamento. Regras que não se negociam: breaking change é `major` mesmo em `v0.x.y`; docs/spec/report-only tende a `NO_RELEASE`; prerelease (`alpha`/`beta`/`rc`) é maturidade do candidato, não disfarce de incompatibilidade.

## 2. Preencher o checklist

Use [release-checklist.md](../release-checklist.md): cabeçalho (versão candidata, classificação, caminho de distribuição) e gates marcados como `PASS`, `FAIL` ou `UNVERIFIED`, incluindo `go test ./...` e `go test -race ./...`. Gate crítico com `FAIL` impede `STABLE_TAG`; `UNVERIFIED` precisa aparecer nas release notes.

## 3. Atualizar changelog e release notes

Siga [release-notes-policy.md](../release-notes-policy.md): `CHANGELOG.md` no formato Keep a Changelog com seção `[Unreleased]`, e release notes com summary, compatibilidade, gates e riscos residuais.

## 4. Registrar decisões

Decisão de arquitetura ou governança não óbvia vira ADR em [decisions/](../decisions/), seguindo o formato e a numeração de [decisions/README.md](../decisions/README.md). ADR aceito é imutável; revisões criam um ADR novo que o supersede.

## Encerramento

- Evidência esperada: checklist preenchido com decisão final e versão recomendada, changelog atualizado e, quando houver decisão não óbvia, ADR numerado.
- Bloqueia quando: gate crítico em `FAIL`, blocker aberto para a capability afetada, ou evidência insuficiente para prometer estabilidade (use `HOLD_RELEASE` ou `PRERELEASE_TAG`).
- Próxima ação humana: publicar a tag escolhida; se a avaliação revelar gaps, registrá-los pelo backlog na [jornada 03](03-backlog-e-trilha-sdd.md).
