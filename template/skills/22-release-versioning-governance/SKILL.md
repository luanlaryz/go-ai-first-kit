---
name: release-versioning-governance
description: Govern versioning and release decisions for {{PROJECT_SLUG}}. Use when an agent needs to decide whether a change warrants a new release, classify patch/minor/major impact, detect breaking changes in the public surface, prepare release checklist, changelog and release notes, and recommend the correct distribution path such as no release, prerelease, stable tag or hotfix.
---

# Release Versioning Governance
Goal: make release, versioning and distribution decisions for `{{PROJECT_SLUG}}` consistently and transparently, with explicit SemVer reasoning and hard gates before any stable release.

## When to Use
- Evaluating whether a merged or proposed change should produce a new release.
- Classifying a change as `patch`, `minor`, `major`, `prerelease`, `hotfix`, or `no release`.
- Preparing release checklist, changelog, release notes, distribution recommendation, or compatibility assessment.
- Auditing whether a candidate release is safe to publish without compromising stable consumers.

## Non-negotiables
1. Use SemVer as the public release contract.
2. Treat `pkg/*`, documented HTTP contracts, and published OpenAPI as public surface.
3. Treat `internal/*` as non-public unless surfaced through `pkg/*`, HTTP, OpenAPI, starter flow, or documented behavior.
4. Never recommend mutating an already published release; publish a new version instead.
5. Stable release requires passing release gates; if evidence is incomplete, recommend `NOT READY` or `prerelease`.
6. Do not classify by intent alone; classify by observable impact on consumers.
7. For `{{PROJECT_SLUG}}`, phase reports and parity diagnostics are release evidence, not optional context.

## Public Surface Rules
Use these rules before classifying impact:
- **Public API:** exported contracts in `pkg/*`, OpenAPI-described endpoints, starter-facing documented setup flow, documented config and CLI/demo usage that the repo treats as supported.
- **Non-public API:** `internal/*`, test helpers, temporary build seams, comments-only changes, reports, and internal file moves with no public behavior impact.
- **Potentially public behavior:** example/starter flows, `/docs`, `/openapi.yaml`, documented response envelopes, SSE event shapes, trace/read surfaces if documented as part of the framework.

## Release Classification Rules
Read [release_policy.md](resources/release_policy.md) and [version_decision_tree.md](resources/version_decision_tree.md) before deciding.

### `no release`
Choose when:
- only docs/reports/spec reconciliation changed
- comments or internal-only refactors changed
- tests changed without public behavior change
- examples changed without affecting documented supported flow

### `patch`
Choose when:
- bug fix preserves public API and documented behavior
- validation, error mapping, docs surface, or runtime defect is corrected compatibly
- OpenAPI/doc bundle is corrected to match already-existing behavior

### `minor`
Choose when:
- new backward-compatible capability is added
- new optional endpoint, config, trace surface, starter path, or documented public contract is added compatibly
- existing behavior is extended without breaking prior consumers

### `major`
Choose when:
- exported symbol removal or signature change in `pkg/*`
- incompatible HTTP or OpenAPI contract change
- renamed/removed documented endpoint
- incompatible behavior change in starter or canonical setup flow
- changed semantics that force consumer code changes even if symbols remain

### `prerelease`
Choose when:
- feature is real but baseline is not yet stable enough for general stable release
- diagnostics say capability is present but not yet fully closed
- API exists but you want consumer trial without stability promise
- recommended tags: `vX.Y.Z-alpha.N`, `vX.Y.Z-beta.N`, `vX.Y.Z-rc.N`
- `alpha.N`: early technical feedback or unstable shape
- `beta.N`: mostly complete candidate for controlled integration with known gaps documented
- `rc.N`: stable-release candidate where only localized fixes are expected before final tag

Prerelease stages describe candidate maturity only. They do not hide breaking changes, missing gates, or known incompatibilities.

### `hotfix`
Choose when:
- urgent fix targets latest stable line
- change is minimal and clearly `patch`
- risk of waiting for next normal release is unacceptable

## Required Inputs
When applying this skill, gather at least:
- change summary or diff summary
- touched files/modules
- whether public surface changed
- current latest released version or intended baseline
- relevant phase/spec/report evidence if the repo uses phase gates for this change

If any of the above is missing, say what is missing and classify provisionally.

## Required Outputs
Always produce all of the following sections:
1. **Release decision**
2. **Version bump recommendation**
3. **Why this is not lower/higher impact**
4. **Breaking change assessment**
5. **Release checklist status**
6. **Changelog draft**
7. **Release notes draft**
8. **Recommended distribution path**
9. **Residual risks / blockers**

Use the templates in [output_templates.md](resources/output_templates.md).

## Release Gates
Apply [release_checklist.md](resources/release_checklist.md) strictly.

At minimum, before recommending a stable release, verify:
- `go test ./...`
- `go test -race ./...`
- relevant smoke/conformance checks for changed surface
- OpenAPI/docs alignment if HTTP surface changed
- feature matrix/reports not contradicting the proposed release story
- no open blocker-level gap for the affected capability

If any gate is unknown, explicitly mark as `UNVERIFIED`.

## Distribution Recommendation Rules
Recommend exactly one path:
- `NO_RELEASE`
- `PRERELEASE_TAG`
- `STABLE_TAG`
- `HOTFIX_TAG`
- `HOLD_RELEASE`

Map them as follows:
- docs/spec/report only => usually `NO_RELEASE`
- compatible completed feature with green gates => `STABLE_TAG`
- incomplete but useful capability => `PRERELEASE_TAG`
- urgent patch on stable line => `HOTFIX_TAG`
- blockers or unverified gates => `HOLD_RELEASE`

## {{PROJECT_SLUG}} Specific Rules
1. Use phase diagnostics as evidence for release maturity.
2. If a phase/report says `NOT READY`, do not recommend stable release for that track.
3. If a phase closes a capability but known S2/S3 gaps remain, reflect them in risk and release notes.
4. If change only reconciles specs/docs without new consumer behavior, prefer `NO_RELEASE`.
5. If OpenAPI, `/docs`, starter, or public traces change materially, treat them as public-facing changes.
6. For `v0.x.y`, breaking changes are still breaking changes; do not hide them just because major is zero.
7. Recommend `v1.0.0` only when public API/starter/docs/release policy are intentionally declared stable.

## Checklists
**Before deciding**
- [ ] Identify touched public surface.
- [ ] Identify current and target version line.
- [ ] Identify whether change is behavior, docs-only, or compatibility-affecting.
- [ ] Load release policy, decision tree, checklist, and templates.

**Before stable recommendation**
- [ ] Core tests green.
- [ ] Race tests green.
- [ ] Changed capability diagnostics not blocked.
- [ ] Changelog and release notes drafted.
- [ ] Distribution path chosen.

## Definition of Done
A release recommendation is complete only when:
- SemVer classification is explicit.
- Public-surface impact is explained.
- Checklist status is explicit (`PASS`, `FAIL`, `UNVERIFIED`).
- Changelog and release notes are drafted.
- Distribution path is explicit.
- Residual blockers are named honestly.

## Minimal Examples
- Docs/spec reconciliation only -> `NO_RELEASE`, maybe roll into next minor release notes.
- New backward-compatible HTTP endpoint -> `minor`, usually `STABLE_TAG` if gates are green.
- Renamed exported `pkg/*` symbol -> `major` or defer to next major line.
- Urgent bug fix in stable persistence path -> `patch`, `HOTFIX_TAG` if urgent.
