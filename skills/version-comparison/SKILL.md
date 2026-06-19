---
name: version-comparison
description: Compare two versions to determine which is newer, older, or if they are equal. Use when implementing update-check logic, dependency resolution, or any ordering decision between versions.
argument-hint: <version1> <version2>
---

# Version Comparison

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

## When to Use

- You need to determine which of two versions is newer/older
- You need to check if a version falls within a range (between low and high)
- You are implementing update-checking or dependency version logic
- You need to compare pre-release vs release versions (alpha < beta < rc < stable)
- You need batch comparison: is a version newer/older/equal than a target?

## Decision Tree

```
Two versions to compare?
├─ Need detailed ordering (-1/0/1)?     → CompareTo() / versions compare / version_compare
├─ Need boolean yes/no answer?          → IsNewerThan / IsOlderThan / Equals
├─ Need range check (between bounds)?   → IsBetween() / versions check --between-low/--high
├─ Need suffix-level comparison?        → VersionSuffix.CompareTo()
└─ Need number-only comparison?         → VersionNumbers.CompareTo()
```

## Task Patterns

### Determine which version is newer

**Goal:** Compare `"1.2.3"` and `"2.0.0"` — which is newer?

**SDK approach:**
```go
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("2.0.0")
result := v1.CompareTo(v2)    // -1 (v1 older)
isNewer := v2.IsNewerThan(v1) // true
```

**CLI approach:**
```bash
versions compare 1.2.3 2.0.0     # {"result": -1, ...}
versions check --newer 1.0.0 2.0.0  # exit 0 = true
```

**MCP approach:**
```json
{"tool": "version_compare", "arguments": {"version1": "1.2.3", "version2": "2.0.0"}}
```

### Check if a version is between two bounds

**Goal:** Is `"1.5.0"` between `"1.0.0"` and `"2.0.0"`?

**SDK approach:**
```go
v := versions.NewVersion("1.5.0")
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")
inRange := v.IsBetween(low, high) // true (inclusive both ends)
// Pass nil for open-ended: v.IsBetween(nil, high) = no lower bound
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

### Compare pre-release vs release versions

**Goal:** Determine that `"1.0.0-beta"` is older than `"1.0.0"`.

**SDK approach:**
```go
beta := versions.NewVersion("1.0.0-beta")
release := versions.NewVersion("1.0.0")
beta.IsOlderThan(release) // true — empty suffix > any suffix

alpha := versions.NewVersion("1.0.0-alpha")
alpha.IsOlderThan(beta)   // true — alpha(100) < beta(200) by weight
```

**CLI approach:**
```bash
versions compare 1.0.0-beta 1.0.0    # result: -1 (beta is older)
versions compare 1.0.0-alpha 1.0.0-beta  # result: -1 (alpha < beta)
```

**MCP approach:**
```json
{"tool": "version_compare", "arguments": {"version1": "1.0.0-beta", "version2": "1.0.0"}}
```

### Compare version number arrays directly

**Goal:** Compare `[1,2,3]` vs `[1,2,0]` without involving full Version objects.

**SDK approach:**
```go
nums1 := versions.NewVersionNumbers([]int{1, 2, 3})
nums2 := versions.NewVersionNumbers([]int{1, 2, 0})
result := nums1.CompareTo(nums2) // 1 (nums1 > nums2)
```

**CLI approach:**
Not available — use full version comparison (`versions compare 1.2.3 1.2.0`).

**MCP approach:**
Not available — use `version_compare` with full version strings.

### Compare suffixes by semantic weight

**Goal:** Determine that `"-alpha2"` is older than `"-beta1"`.

**SDK approach:**
```go
s1 := versions.VersionSuffix("-alpha2")
s2 := versions.VersionSuffix("-beta1")
s1.CompareTo(s2) // -1: alpha(100) < beta(200)
// Same type with sub-version: "-alpha1" < "-alpha3" (1 < 3)
```

**CLI approach:**
Not directly available — use full version comparison.

**MCP approach:**
Not directly available — use `version_compare` with full version strings.

## API Reference

### SDK — Core Comparison Methods

```go
// CompareTo returns: -1 (x < target), 0 (equal), 1 (x > target)
// Comparison order: VersionNumbers → Suffix → PublicTime → Raw string
func (x *Version) CompareTo(target *Version) int

// Boolean convenience methods
func (x *Version) IsNewerThan(target *Version) bool   // CompareTo(target) > 0
func (x *Version) IsOlderThan(target *Version) bool   // CompareTo(target) < 0
func (x *Version) Equals(target *Version) bool         // CompareTo(target) == 0

// Range check — inclusive on both bounds (low <= x <= high)
// Pass nil to skip a bound
func (x *Version) IsBetween(low, high *Version) bool
```

### SDK — Sub-Component Comparison

```go
// Compare number arrays element-by-element, left to right
// Shorter arrays are less than longer ones when shared prefix matches
func (x VersionNumbers) CompareTo(target []int) int

// Compare suffixes by semantic weight:
// dev(50) < snapshot(60) < nightly(70) < alpha(100) < beta(200)
// < milestone(300) < rc(400) < final/release/ga(500) < sp(600)
// < patch(700) < post(800) < empty/release(∞)
// Unknown suffixes sort after known suffixes, alphabetically
func (x VersionSuffix) CompareTo(target VersionSuffix) int
```

### SDK — VersionNumbers Constructor

```go
func NewVersionNumbers(nums []int) VersionNumbers
```

### CLI Commands

```bash
# Compare two versions — JSON output with result (-1/0/1) and description
versions compare <version1> <version2>

# Boolean comparison checks — exit 0 = true, exit 1 = false
versions check --newer <target> <version>                   # IsNewerThan
versions check --older <target> <version>                   # IsOlderThan
versions check --equal <target> <version>                   # Equals
versions check --between-low <low> --between-high <high> <v> # IsBetween
```

**Examples:**
```bash
versions compare 1.2.3 2.0.0           # {"result": -1, "description": "..."}
versions check --newer 1.0.0 2.0.0     # exit 0 (true: 2.0.0 > 1.0.0)
versions check --older 2.0.0 1.0.0     # exit 0 (true: 1.0.0 < 2.0.0)
versions check --equal 1.0.0 1.0.0     # exit 0 (true)
```

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_compare` | `version1: string`, `version2: string` | `{result: int, description: string}` |
| `version_range_query` | `check_version: string`, `low: string`, `high: string` | between-ness result |

## Cross-References

- [[version-parsing]] — for parsing version strings before comparison
- [[version-check]] — for boolean type checks and CI/CD conditionals
- [[version-sorting]] — for sorting collections using comparison logic
- [[version-range-query]] — for querying ranges of versions
- [[version-constraints]] — for constraint-based version matching

## Important Notes

- **Comparison order: VersionNumbers → Suffix → PublicTime → Raw string** (Suffix is before PublicTime, not after)
- **`1.0` and `1.0.0` are NOT equal** — shorter number arrays are less than longer ones when the shared prefix matches: `[1,0] < [1,0,0]`
- **Empty suffix (release) is always greater than any non-empty suffix (prerelease)**: `1.0.0 > 1.0.0-beta > 1.0.0-alpha`
- **Suffix comparison uses semantic weight, not alphabetical order**: dev(50) < alpha(100) < beta(200) < rc(400) < release(500)
- **Prefix does not affect CompareTo** — `v1.0.0` and `1.0.0` are treated the same on numbers/suffix/time
- **IsBetween is inclusive on both bounds** (`low <= x <= high`); pass `nil` for open-ended checks
- **PublicTime is only compared when BOTH versions have non-zero PublicTime** — if one is zero, this step is skipped
