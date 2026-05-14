# Release Policy for {{PROJECT_SLUG}}

## 1. Versioning Model
Use SemVer:
- `MAJOR`: incompatible public change
- `MINOR`: backward-compatible feature
- `PATCH`: backward-compatible fix

For `v0.x.y`, still classify breakage honestly. The repo may remain in `v0`, but the release recommendation must still flag breaking changes.

## 2. Public Contract Boundaries
Treat as public:
- exported contracts in `pkg/*`
- documented HTTP endpoints
- OpenAPI schemas and envelopes
- starter / getting-started flow when documented as canonical
- documented config and trace surfaces

Treat as non-public unless surfaced indirectly:
- `internal/*`
- smoke/test helpers
- report-only artifacts
- implementation-only seams

## 3. Stable Release Conditions
Recommend stable release only when:
- no blocker gap is open for the changed capability
- validation commands passed or are explicitly confirmed
- docs and OpenAPI match real behavior if relevant
- release notes/changelog are consistent with actual change scope

## 4. Prerelease Conditions
Prefer prerelease when:
- capability is newly added but not yet hardened
- diagnostics are partial or recent
- external feedback is desired before stability promise

Use the stage to describe candidate maturity:
- `alpha.N`: early technical feedback, validation, or unstable shape. Expect API, payload, docs, or behavior refinements before stability.
- `beta.N`: mostly complete candidate for controlled integration. Core behavior is present, contracts are documented, and known gaps are listed.
- `rc.N`: stable-release candidate. No broad functional changes are expected before the final tag; only localized fixes or final documentation adjustments should remain.

Prerelease stages do not replace SemVer classification. A breaking change is still `MAJOR`, and missing gates or known incompatibilities must remain explicit in the release decision.

## 5. Hotfix Conditions
Use hotfix only when all are true:
- targets a released stable line
- urgent consumer-facing defect
- patch-scoped fix
- low blast radius

## 6. Release Immutability Rule
Never recommend changing a published version in place.
Always publish a new version.
