# Output Templates

## 1. Release Decision
- Decision: `NO_RELEASE | PRERELEASE_TAG | STABLE_TAG | HOTFIX_TAG | HOLD_RELEASE`
- Recommended version: `vX.Y.Z` or `vX.Y.Z-alpha.N | vX.Y.Z-beta.N | vX.Y.Z-rc.N`
- Classification: `patch | minor | major | no-release`
- Prerelease stage: `N/A | alpha.N | beta.N | rc.N`
- Prerelease stage rationale:

## 2. Breaking Change Assessment
- Public surface touched:
- Breaking?: `yes | no | possible`
- Why:

## 3. Checklist Status
| Check | Status | Notes |
| --- | --- | --- |
| go test ./... | PASS/FAIL/UNVERIFIED | |
| go test -race ./... | PASS/FAIL/UNVERIFIED | |
| Relevant smoke/conformance | PASS/FAIL/UNVERIFIED | |
| OpenAPI/docs alignment | PASS/FAIL/UNVERIFIED | |
| Prerelease stage justified | PASS/FAIL/N/A | |

## 4. Changelog Draft
### Added
- ...
### Changed
- ...
### Fixed
- ...
### Deprecated
- ...
### Removed
- ...

## 5. Release Notes Draft
### Summary
...

### Prerelease Stage
...

### Consumer Impact
...

### Upgrade Notes
...

### Risks
...

## 6. Distribution Recommendation
- Path:
- Reason:
- Suggested tag/branch action:
