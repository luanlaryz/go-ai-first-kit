# Version Decision Tree

## Step 1: Did public behavior change?
- No -> likely `NO_RELEASE` or at most roll into next scheduled release notes.
- Yes -> continue.

## Step 2: Is any consumer action required to keep working?
- Yes -> breaking change -> `major`.
- No -> continue.

## Step 3: Is this new capability or only a fix?
- New compatible capability -> `minor`.
- Compatible correction/fix -> `patch`.

## Step 4: Is release maturity incomplete?
- Yes -> convert recommendation to `PRERELEASE_TAG` and choose a stage:
  - `alpha.N` for early technical feedback or unstable shape.
  - `beta.N` for controlled integration with core behavior present and known gaps documented.
  - `rc.N` for a stable-release candidate where only localized fixes are expected before final tag.
- No -> continue.

## Step 5: Is urgency exceptional on stable line?
- Yes -> `HOTFIX_TAG`.
- No -> `STABLE_TAG`.

## Additional checks
- If only spec/report/README/feature-matrix changed without consumer-visible behavior: `NO_RELEASE`.
- If OpenAPI changed because actual runtime changed, classify from runtime impact, not from file type alone.
- If starter/getting-started path changed materially for new users, treat as public-surface impact.
- Prerelease stage never lowers SemVer impact; breaking changes remain breaking.
