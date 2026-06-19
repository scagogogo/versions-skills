---
name: version-check
description: Check boolean properties of version numbers — type checks (IsBeta, IsStable, IsRC, etc.) and comparison predicates (IsNewerThan, IsOlderThan, IsBetween). Use for CI/CD conditionals, shell script exit-code checks, and programmatic version filtering.
argument-hint: <version-string> --<flag>
---

# Version Check

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- You need to check if a version is a specific type (alpha, beta, RC, stable, dev, snapshot, etc.)
- You need boolean comparison results (is newer? is older? is between?)
- You are building CI/CD conditionals based on version properties
- You need exit-code-based checks for shell scripts (`if versions check --stable $V; then ...`)
- You need to filter versions by type or comparison predicates

## Decision Tree

```
Need a yes/no answer about a version?
├─ Type check (what kind of version)?
│   ├─ Has any suffix?                → IsPrerelease() / versions check --prerelease
│   ├─ No suffix (release)?           → IsStable() / versions check --stable
│   ├─ Specific suffix type?          → IsAlpha/IsBeta/IsRC/IsDev/etc. / versions check --<type>
│   ├─ All type flags at once?        → version_info (MCP) / versions info (CLI)
│   └─ All numbers are zero?          → IsZero() / versions check --zero
├─ Comparison check (relative to another version)?
│   ├─ Is v1 newer than v2?           → IsNewerThan() / versions check --newer v2 v1
│   ├─ Is v1 older than v2?           → IsOlderThan() / versions check --older v2 v1
│   ├─ Are they equal?                → Equals() / versions check --equal v1 v2
│   └─ Is v between low and high?     → IsBetween() / versions check --between-low L --between-high H V
└─ Need suffix weight for ordering?   → SuffixWeight() / versions suffix-weight
```

## Task Patterns

### Check version type in CI/CD pipeline

**Goal:** Only deploy if the version is stable (no prerelease suffix).

**SDK approach:**
```go
v := versions.NewVersion(os.Getenv("VERSION"))
if v.IsStable() {
    deploy()
}
```

**CLI approach:**
```bash
if versions check --stable "$VERSION"; then
    deploy
fi
```

**MCP approach:**
```json
{"tool": "version_info", "arguments": {"version_string": "1.2.3"}}
```
Check `IsStable` in the response.

### Check if a version is a specific prerelease type

**Goal:** Determine if `"1.0.0-beta2"` is a beta (not just any prerelease).

**SDK approach:**
```go
v := versions.NewVersion("1.0.0-beta2")
isBeta := v.IsBeta()         // true
isAlpha := v.IsAlpha()       // false (beta != alpha)
isRC := v.IsRC()             // false
isPrerelease := v.IsPrerelease() // true (any suffix)
subVersion := v.SubVersion() // 2
```

**CLI approach:**
```bash
versions check --beta 1.0.0-beta2      # exit 0
versions check --alpha 1.0.0-beta2     # exit 1
versions check --rc 1.0.0-beta2        # exit 1
```

**MCP approach:**
```json
{"tool": "version_info", "arguments": {"version_string": "1.0.0-beta2"}}
```

### Check comparison predicates

**Goal:** Determine if `"2.0.0"` is newer than `"1.0.0"`.

**SDK approach:**
```go
v1 := versions.NewVersion("1.0.0")
v2 := versions.NewVersion("2.0.0")
v2.IsNewerThan(v1)  // true
v1.IsOlderThan(v2)  // true
v1.Equals(v2)       // false
```

**CLI approach:**
```bash
versions check --newer 1.0.0 2.0.0     # exit 0 (2.0.0 > 1.0.0)
versions check --older 2.0.0 1.0.0     # exit 0 (1.0.0 < 2.0.0)
versions check --equal 1.0.0 1.0.0     # exit 0
```

**MCP approach:**
```json
{"tool": "version_compare", "arguments": {"version1": "2.0.0", "version2": "1.0.0"}}
```
Check `result == 1` in the response.

### Check if a version is in a range

**Goal:** Determine if `"1.5.0"` is between `"1.0.0"` and `"2.0.0"`.

**SDK approach:**
```go
v := versions.NewVersion("1.5.0")
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")
inRange := v.IsBetween(low, high) // true (inclusive both ends)
```

**CLI approach:**
```bash
versions check --between-low 1.0.0 --between-high 2.0.0 1.5.0  # exit 0
```

**MCP approach:**
```json
{
  "tool": "version_range_query",
  "arguments": {"check_version": "1.5.0", "low": "1.0.0", "high": "2.0.0"}
}
```

### Check suffix weight for ordering

**Goal:** Get the semantic weight of `"1.0.0-alpha1"` to compare with other suffixes.

**SDK approach:**
```go
v := versions.NewVersion("1.0.0-alpha1")
weight := v.SuffixWeight()                    // 100
isAlpha := weight == versions.SuffixWeightAlpha // true
name := weight.String()                       // "alpha"
```

**CLI approach:**
```bash
versions suffix-weight 1.0.0-alpha1   # alpha (100)
versions suffix-weight 1.0.0           # unknown (0)
```

**MCP approach:**
```json
{"tool": "version_parse", "arguments": {"version_string": "1.0.0-alpha1"}}
```
Check `suffix_weight` in the response.

### Check all type flags at once (for comprehensive filtering)

**Goal:** Get every Is* flag for a version in a single call.

**SDK approach:**
```go
v := versions.NewVersion("v1.2.3-beta1")
// Call each Is* method as needed:
// v.IsPrerelease(), v.IsStable(), v.IsDev(), v.IsAlpha(),
// v.IsBeta(), v.IsRC(), v.IsSnapshot(), v.IsMilestone(),
// v.IsNightly(), v.IsFinal(), v.IsGA(), v.IsPre(),
// v.IsRelease(), v.IsSP(), v.IsPost(), v.IsZero()
```

**CLI approach:**
```bash
versions info v1.2.3-beta1
```

**MCP approach:**
```json
{"tool": "version_info", "arguments": {"version_string": "v1.2.3-beta1"}}
```

## Cross-References

- [[version-parsing]] — for parsing version strings before checking
- [[version-comparison]] — for the underlying CompareTo logic behind IsNewerThan/IsOlderThan
- [[version-constraints]] — for constraint-based matching (more expressive than simple checks)
- [[version-range-query]] — for range-based queries on version collections

## Important Notes

- `IsStable()` and `IsPrerelease()` are exact opposites: a version is either stable (no suffix) or prerelease (has suffix)
- Type checks are suffix-based: `"1.2.3-beta1"` is `IsBeta()` but NOT `IsAlpha()` — they are mutually exclusive
- `IsPre()` and `IsPrerelease()` are different: `IsPre()` matches the specific `-pre*` suffix type, while `IsPrerelease()` matches ANY non-empty suffix
- `IsStable()` and `IsRelease()` are different: `IsStable()` means "no suffix", `IsRelease()` means "has explicit `-release` suffix"
- `IsZero()` checks if all version numbers are zero — `"0.0.0"` has valid numbers so `IsValid()` is true
- CLI `check` exit codes: 0 = true, 1 = false — ideal for shell conditionals
- `IsBetween(low, high)` is inclusive on both bounds; pass `nil` for open-ended checks
