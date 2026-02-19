# OP-3118: Case-Sensitive Filters Are Difficult

## Problem Statement

Text filters across all Clarity datagrid admin lists are case-sensitive. When a user types "smith" in the Last Name filter, it does not match "Smith" because the backend uses PostgreSQL `LIKE` (case-sensitive). The reporter also requests substring/wildcard matching (e.g., "brook" matching "Alsbrooks").

## Root Cause Analysis

### Filter Data Flow
```
Clarity Datagrid filter input ("smith")
  -> QueryModel.toCriteria() [query-model.ts:210-222]
  -> QueryCriteria.startsWith(field, value) [query-model.ts:356-361]
  -> V2httpService.listByQueryModel() encodes as HTTP param: lastName=<sw>smith
  -> Backend stg-sdk-golang operatorMap translates <sw> to SQL: LIKE 'smith%'
  -> PostgreSQL LIKE is case-sensitive -> no results for "Smith"
```

### Affected Components (9 total with server-side filtering)

| Component | File | Affected Text Fields |
|---|---|---|
| **UserList** | `modules/admin/user/user-list/user-list.component.ts` | username, ref, email, firstName, lastName |
| **SiteList** | `shared/components/site/site-list/site-list.component.ts` | name, ref, profileName, geoAddress1, geoAddress2, geoCity, geoPostalCode, geoStateCode |
| **ProfileList** | `shared/components/profile/profile-list/profile-list.component.ts` | name, ref, description |
| **EquipList** | `modules/equip/components/equip-list/equip-list.component.ts` | siteName, name, ref, description |
| **PointList** | `modules/point/components/point-list/point-list.component.ts` | equipName, name, ref, description |
| **UridList** | `modules/admin/urid/urid-list/urid-list.component.ts` | name, ref, description, category |
| **EquipTypeList** | `modules/admin/equip-type/equip-type-list/equip-type-list.component.ts` | name, description, code |
| **EquipTypeConfigList** | `modules/admin/equip-type-config/equip-type-config-list/equip-type-config-list.component.ts` | name, description |
| **BranchList** | `modules/admin/branch-table/branch-list/branch-list.component.ts` | name, ref, description |

### Bonus Bugs Found During Analysis

1. **Silent filter failures** - Several `[clrDgField]` values are NOT in `_defaultStartsWithFields` or any other field list:
   - `equipTypeName` in EquipList - silently dropped (user types, nothing happens)
   - `unit`, `pointTypeName`, `pointUridName` in PointList - silently dropped
   - `equipTypeConfig` in EquipTypeList - wrong field name, should be `equipTypeConfigName`
   - `enabled`, `phoneNumber` in UserList - not in any field list
2. **Console noise** - Unmatched filters log `console.log('unable to convert filter to criteria')` with no user feedback

---

## Proposed Solution

### Approach: Backend `ILIKE` via `<sw>` operator change

**Recommended**: Modify the `<sw>` operator in `stg-sdk-golang` to use PostgreSQL `ILIKE` instead of `LIKE`. This makes ALL existing text filters case-insensitive with **zero frontend changes** required.

#### Why this approach?
- **Minimal change surface** - One line in the SDK, zero lines in stark-web
- **Universal fix** - Every existing `<sw>` filter becomes case-insensitive automatically
- **No API contract change** - `<sw>` parameter name stays the same
- **PostgreSQL `ILIKE` performance** - Equivalent to `LIKE` when a `gin_trgm` index exists; for prefix matching, it can use btree indexes with `text_pattern_ops`

#### Backend Change (stg-sdk-golang)

**Repo**: `/Volumes/data/projects/tsp/stg-sdk-golang`
**File**: `pkg/starkapi/query-params.go`

Two lines to change in `parameterizedClause()`:

**Line 411** (`<sw>` / startLike):
```go
// BEFORE:
return fmt.Sprintf("%s like $%d", p.Column, seedIndex+1)
// AFTER:
return fmt.Sprintf("%s ilike $%d", p.Column, seedIndex+1)
```

**Line 417** (`<ew>` / endLike):
```go
// BEFORE:
return fmt.Sprintf("%s like $%d", p.Column, seedIndex+1)
// AFTER:
return fmt.Sprintf("%s ilike $%d", p.Column, seedIndex+1)
```

**Test updates** in `pkg/starkapi/query-params_test.go`:
- `TestQueryParams_StartLike` (line 598): Update expected SQL from `like` to `ilike`
- `TestQueryParams_EndLike` (line 609): Update expected SQL from `like` to `ilike`
- `TestQueryParams_StartLikeWithEventType` (line 620): Update expected SQL from `like` to `ilike`
- `TestQueryParams_EndLikeWithEventType` (line 632): Update expected SQL from `like` to `ilike`

#### Consuming Service Update (stark-asset-api)

**Repo**: `/Volumes/data/projects/tsp/stark-asset-api`

After the SDK change is merged and tagged:
- Update `go.mod` to reference the new SDK version
- No code changes needed in asset-api itself — it delegates to the SDK's `BuildParameterizedQuery()`

#### Optional: Add `contains` operator for substring matching

To address the wildcard/substring request ("brook" matches "Alsbrooks"), add a new `<cn>` operator:

**Backend** (`query-params.go`):
```go
"<cn>": containsLike, // new operator for ILIKE '%value%'
```

With SQL generation:
```go
p.Value = "%" + p.Value.(string) + "%"
return fmt.Sprintf("%s ilike $%d", p.Column, seedIndex+1)
```

**Frontend** (`query-model.ts`):
```typescript
// New operator string
private static readonly containsString = '<cn>';

// New factory method
public static contains(field: string, value: string): QueryCriteria {
  if (value === null || value === undefined) return null;
  return new QueryCriteria(field, QueryCriteria.containsString, value);
}
```

Then add a `_defaultContainsFields` list in `QueryModel` for fields that should use substring matching, or make `contains` the default for all text fields (replacing `startsWith`).

---

## Implementation Steps

### Phase 1: Case-insensitive `<sw>` (backend SDK change)
1. Open PR in `stg-sdk-golang` to change `LIKE` to `ILIKE` for `<sw>` and `<ew>` operators
2. Verify no existing behavior depends on case-sensitive `LIKE`
3. Test with existing admin lists - no frontend changes needed
4. Bump SDK version in consuming services (api, asset-api)

### Phase 2: Frontend fixes for silent filter failures (stark-web)
1. Add missing fields to `_defaultStartsWithFields` in `query-model.ts`:
   - `equipTypeName` (for EquipList)
   - `unit` (for PointList)
   - `pointTypeName` (for PointList)
   - `pointUridName` (for PointList)
   - `enabled` (for UserList - should be equals, not startsWith)
   - `phoneNumber` (for UserList)
2. Fix `equipTypeConfig` -> `equipTypeConfigName` in EquipTypeList template
3. Add tests to `query-model.spec.ts` for the `toCriteria` method

### Phase 3 (Optional): Contains/substring operator
1. Add `<cn>` operator to backend SDK
2. Add `QueryCriteria.contains()` to frontend
3. Decide per-field: use `startsWith` or `contains` as default
4. Consider adding a UX toggle in datagrid filter header to switch between "starts with" and "contains"

---

## Testing Plan

### Unit Tests (Frontend)
- [ ] `query-model.spec.ts`: Test `toCriteria` maps all `_defaultStartsWithFields` correctly
- [ ] `query-model.spec.ts`: Test that unknown fields don't silently disappear (or at least test the current behavior)
- [ ] `query-model.spec.ts`: Test `startsWith` produces correct operator string
- [ ] `query-model.spec.ts`: Test `contains` produces correct operator string (Phase 3)

### Integration Tests
- [ ] User Management: Filter by "smith" matches "Smith", "SMITH", "smith"
- [ ] Site List: Filter by address fragment, case-insensitive
- [ ] Equip List: Filter by name, case-insensitive
- [ ] Point List: Filter by name, case-insensitive

### Regression
- [ ] Verify numeric/date filters unaffected (they don't use `<sw>`)
- [ ] Verify sort, pagination, offset still work correctly
- [ ] Verify `<eq>`, `<ge>`, `<le>` operators unaffected

---

## Risk Assessment

| Risk | Severity | Mitigation |
|---|---|---|
| `ILIKE` performance on large tables | Low | PostgreSQL `ILIKE` with prefix patterns uses btree indexes with `text_pattern_ops`; for mid-string, `pg_trgm` GIN indexes can be added if needed |
| Behavioral change surprises | Low | Case-insensitive search is strictly more permissive (superset of results), unlikely to break anything |
| SDK version bump coordination | Medium | Requires deployment of updated API services before frontend benefits |

---

## Files to Modify

### Backend (stg-sdk-golang)
- `pkg/starkapi/query-params.go` - Change `LIKE` to `ILIKE` for `<sw>` and `<ew>`

### Frontend (stark-web) - Phase 2
- `ClientApp/src/app/shared/models/query-model.ts` - Add missing fields to `_defaultStartsWithFields`
- `ClientApp/src/app/shared/models/query-model.spec.ts` - Add filter mapping tests
- `ClientApp/src/app/modules/admin/equip-type/equip-type-list/equip-type-list.component.html` - Fix `equipTypeConfig` field name

### Frontend (stark-web) - Phase 3 (Optional)
- `ClientApp/src/app/shared/models/query-model.ts` - Add `contains` operator and `_defaultContainsFields`

---

## Decision Points for Team

1. **Should `<sw>` be globally changed to `ILIKE`?** Or should we add a new `<swi>` operator and leave `<sw>` as-is?
   - Recommendation: Change `<sw>` globally. No use case for case-sensitive prefix matching has been identified.

2. **Should we add substring matching (`<cn>` / contains)?** The ticket specifically mentions wanting "brook" to match "Alsbrooks."
   - Recommendation: Yes, as a follow-up. Start with case-insensitive `startsWith` first.

3. **Should `startsWith` be replaced with `contains` as the default text filter behavior?**
   - Recommendation: Keep `startsWith` as default for performance. Add `contains` as opt-in per field.
