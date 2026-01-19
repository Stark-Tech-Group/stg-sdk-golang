# stg-sdk-golang: Go 1.25 Upgrade Plan

## Overview

Upgrade stg-sdk-golang from Go 1.23 to Go 1.25.0 as part of the platform-wide Go version upgrade initiative.

**Jira Ticket:** [OP-2875](https://controlfreak.atlassian.net/browse/OP-2875)
**Branch:** `feature/OP-2875-go-1.25` from `dev`
**Parent Initiative:** Go 1.25 Platform Upgrade (19/25 complete)

---

## Critical Context

**This is a shared SDK library used by 9+ downstream services:**
- stark-event-message-bus
- stark-subscription-service
- notification-service
- stark-babbage
- iot-hub-message-bus
- stark-telemetry-message-bus
- stark-permission-service
- stark-search-service
- And others...

**After this upgrade is merged, a new release must be tagged (v2.1.0) so downstream services can update their dependencies.**

---

## Current State

| Item | Value |
|------|-------|
| Go Version | 1.23 |
| Type | Library (no Dockerfile) |
| CGO Required | No (pure Go) |
| Latest Release | v2.0.2 |
| Test Files | 15 |
| Test Duration | ~1.9s |

---

## Files to Update

### 1. go.mod
```diff
- go 1.23
+ go 1.25.0
```

### 2. .github/workflows/*.yml (if needed)
Update Go version in CI workflows from `^1.23` to `^1.25` or `1.25`.

---

## Dependencies to Update

### Direct Dependencies
| Package | Current | Target | Notes |
|---------|---------|--------|-------|
| Azure/azure-event-hubs-go/v3 | v3.6.1 | v3.6.2 | Azure Event Hubs |
| go-playground/validator/v10 | v10.15.4 | latest | Input validation |
| lib/pq | v1.10.9 | v1.10.9 | PostgreSQL (already latest) |
| sirupsen/logrus | v1.9.3 | v1.9.3 | Logging (already latest) |
| stretchr/testify | v1.8.4 | v1.11.1 | Testing |

### Indirect Dependencies (Security Critical)
| Package | Current | Target | Notes |
|---------|---------|--------|-------|
| golang.org/x/crypto | v0.13.0 | v0.47.0 | **Security patches** |
| golang.org/x/net | v0.15.0 | v0.49.0 | **Security patches** |
| golang.org/x/sys | v0.12.0 | latest | System calls |
| golang.org/x/text | v0.13.0 | latest | Text processing |

---

## Execution Checklist

### Phase 1: Setup
- [x] Create Jira ticket OP-2875
- [x] Create branch from dev
- [x] Create /plans directory
- [x] Create upgrade plan

### Phase 2: Version Updates
- [x] Update go.mod to Go 1.25.0
- [x] Run `go get -u ./...` to update dependencies
- [x] Run `go mod tidy`
- [x] Run `go fmt ./...` to fix formatting (11 files)
- [x] Verify go.sum updated

### Phase 3: Verification
- [x] Run `go vet ./...` - no issues
- [x] Run `go test ./...` - 4 packages pass
- [x] Run `go build ./cmd/stg-sdk-golang` - success

### Phase 4: Ship
- [ ] Commit changes
- [ ] Push to remote
- [ ] Create PR to dev
- [ ] Update master plan

### Phase 5: Release (Post-Merge)
- [ ] Merge PR to dev
- [ ] CI/CD will auto-tag release on build
- [ ] Create Jira tickets for downstream services to update SDK dependency

---

## Technical Notes

### Pure Go Library
This is a library project with no CGO dependencies:
- No Dockerfile to update
- No platform-specific build flags
- Simple `go build ./cmd/stg-sdk-golang` for verification

### Azure SDK Compatibility
The Azure dependencies show `+incompatible` markers:
- `github.com/Azure/azure-sdk-for-go v68.0.0+incompatible`
- `github.com/Azure/go-autorest v14.2.0+incompatible`

These work but indicate older module patterns. Test thoroughly after update.

### Downstream Impact
After releasing v2.1.0, downstream services should update:
```bash
go get github.com/Stark-Tech-Group/stg-sdk-golang@v2.1.0
go mod tidy
```

---

## Rollback Plan

If issues arise:
1. Revert go.mod to Go 1.23
2. Run `go mod tidy`
3. Do not tag a new release

---

## Success Criteria

- [ ] Go 1.25.0 in go.mod
- [ ] All tests pass (4 packages)
- [ ] Build succeeds
- [ ] No breaking API changes
- [ ] CI/CD creates release after merge

---

## Downstream Services to Update

After this SDK is released, create Jira tickets for these services to update their SDK dependency:
- stark-event-message-bus
- stark-subscription-service
- notification-service
- stark-babbage
- iot-hub-message-bus
- stark-telemetry-message-bus
- stark-permission-service
- stark-search-service
