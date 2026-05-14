# Release Checklist

## Required checks
- [ ] Change summary is explicit.
- [ ] Touched files/modules are listed.
- [ ] Public surface impact assessed.
- [ ] Breaking-change assessment completed.
- [ ] Recommended version bump selected.
- [ ] Distribution path selected.
- [ ] Prerelease stage selected and justified when using `PRERELEASE_TAG`.
- [ ] Prerelease stage does not hide breaking change, missing gates, or known incompatibility.
- [ ] `go test ./...` passed or explicitly marked `UNVERIFIED`.
- [ ] `go test -race ./...` passed or explicitly marked `UNVERIFIED`.
- [ ] Relevant smoke/conformance checks passed or explicitly marked `UNVERIFIED`.
- [ ] OpenAPI/docs alignment checked if HTTP/public docs changed.
- [ ] Release notes drafted.
- [ ] Changelog drafted.

## Optional checks
- [ ] Prerelease justification names `alpha.N`, `beta.N`, or `rc.N`.
- [ ] Hotfix urgency justified.
- [ ] Known residual risks listed.
- [ ] Rollback/hold recommendation documented if needed.
