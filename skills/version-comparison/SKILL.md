---
name: version-comparison
description: Use when comparing version numbers, checking which version is newer/older, or determining version ordering. Covers SDK, CLI, and MCP paths for version comparison.
argument-hint: <version-comparison-task>
---

# Version Comparison Skill

## When to Use

- Comparing two version numbers to determine which is newer/older
- Checking if a version satisfies a range (between low and high)
- Sorting or ordering versions by numeric value
- Implementing update-checking or dependency version logic
- Pre-release vs release comparison (alpha < beta < rc < stable)
- Batch comparison: checking if a version is newer/older/equal than a target

## Quick Start

### SDK

```go
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("2.0.0")

// Simple boolean checks
v1.IsOlderThan(v2)   // true
v1.IsNewerThan(v2)   // false
v1.Equals(v2)        // false

// Range check
v := versions.NewVersion("1.5.0")
v.IsBetween(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))  // true
```

### CLI

```bash
# Direct comparison
versions compare 1.2.3 2.0.0

# Boolean checks (exit code 0=true, 1=false)
versions check --newer 1.0.0 2.0.0
versions check --older 2.0.0 1.0.0
versions check --equal 1.0.0 1.0.0

# Range check
versions check --between-low 1.0.0 --between-high 3.0.0 2.0.0
```

### MCP

```json
{
  "tool": "version_compare",
  "arguments": {
    "version1": "1.2.3",
    "version2": "2.0.0"
  }
}
```

Returns: `{ "result": -1, "description": "1.2.3 旧于 2.0.0" }`

## API Reference -- SDK

### Version.CompareTo

**func (x *Version) CompareTo(target *Version) int**

Compares two versions using the following priority order:

1. **VersionNumbers** -- digit-by-digit, left to right; shorter array is less than longer when shared prefix matches (e.g., `1.0` < `1.0.0`)
2. **Suffix** -- empty suffix (release) > any non-empty suffix (pre-release); known suffixes compared by semantic weight; unknown suffixes compared alphabetically
3. **PublicTime** -- earlier publication time = smaller (only compared when both versions have non-zero PublicTime)
4. **Raw string** -- alphabetical fallback when all above are equal

Returns: `-1` (x < target), `0` (equal), `1` (x > target)

### Version.IsNewerThan

**func (x *Version) IsNewerThan(target *Version) bool**

Equivalent to `CompareTo(target) > 0`. Returns true when x is strictly greater than target.

### Version.IsOlderThan

**func (x *Version) IsOlderThan(target *Version) bool**

Equivalent to `CompareTo(target) < 0`. Returns true when x is strictly less than target.

### Version.Equals

**func (x *Version) Equals(target *Version) bool**

Equivalent to `CompareTo(target) == 0`. Returns true when versions are equal per the full comparison order.

### Version.IsBetween

**func (x *Version) IsBetween(low, high *Version) bool**

Returns true when `low <= x <= high` (inclusive on both bounds). Pass `nil` for either bound to skip that side of the check.

### VersionNumbers.CompareTo

**func (x VersionNumbers) CompareTo(target []int) int**

Compares number arrays element-by-element (left to right). When the shared prefix is equal, the longer array is considered greater: `[1,0] < [1,0,0]`. Returns negative (x < target), 0 (equal), positive (x > target).

### VersionSuffix.CompareTo

**func (x VersionSuffix) CompareTo(target VersionSuffix) int**

Compares suffixes using the semantic weight system:

- Known suffixes (alpha, beta, rc, etc.) are compared by weight: dev(50) < snapshot(60) < nightly(70) < alpha(100) < beta(200) < milestone(300) < rc(400) < final/release/ga(500) < sp(600) < patch(700) < post(800)
- Same-weight suffixes with sub-version numbers (e.g., `-alpha1` vs `-alpha2`) compare by the numeric sub-version
- Unknown suffixes sort after known suffixes; unknown-vs-unknown uses alphabetical comparison
- Empty suffix (release version) is always greater than any non-empty suffix

## CLI Commands

### compare

```bash
versions compare <version1> <version2>
```

Returns JSON with `result` (-1/0/1) and `description`. Examples:

```bash
versions compare 1.2.3 2.0.0     # result: -1  (1.2.3 is older)
versions compare 2.0.0 1.0.0     # result: 1   (2.0.0 is newer)
versions compare 1.0.0 1.0.0     # result: 0   (equal)
```

### check (comparison flags)

```bash
versions check --<flag> <target-version> <version-to-check>
```

Returns JSON with boolean `result` and `description`. Exit code 0 = true, 1 = false.

| Flag | Meaning | SDK equivalent |
|------|---------|---------------|
| `--newer <v>` | Is the given version newer than v? | `IsNewerThan()` |
| `--older <v>` | Is the given version older than v? | `IsOlderThan()` |
| `--equal <v>` | Is the given version equal to v? | `Equals()` |
| `--between-low <v>` + `--between-high <v>` | Is the given version in range? | `IsBetween()` |

Examples:

```bash
versions check --newer 1.0.0 2.0.0              # true (2.0.0 > 1.0.0)
versions check --older 2.0.0 1.0.0              # true (1.0.0 < 2.0.0)
versions check --equal 1.0.0 1.0.0              # true
versions check --between-low 1.0.0 --between-high 3.0.0 2.0.0  # true
```

## MCP Tools

### version_compare

Compares two version strings and returns the ordering relationship.

**Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `version1` | string | yes | First version string |
| `version2` | string | yes | Second version string |

**Response fields:**

| Field | Type | Description |
|-------|------|-------------|
| `result` | int | -1 (version1 older), 0 (equal), 1 (version1 newer) |
| `description` | string | Human-readable comparison result |

**Example invocation:**

```json
{
  "tool": "version_compare",
  "arguments": {
    "version1": "1.0.0-beta",
    "version2": "1.0.0"
  }
}
```

Response: `{ "v1": "1.0.0-beta", "v2": "1.0.0", "result": -1, "description": "1.0.0-beta 旧于 1.0.0" }`

## Code Examples

### Basic comparison

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    v1 := versions.NewVersion("1.2.3")
    v2 := versions.NewVersion("1.3.0")

    // Using CompareTo directly
    switch v1.CompareTo(v2) {
    case -1:
        fmt.Println("1.2.3 is older than 1.3.0")
    case 0:
        fmt.Println("versions are equal")
    case 1:
        fmt.Println("1.2.3 is newer than 1.3.0")
    }

    // Using convenience methods
    if v1.IsOlderThan(v2) {
        fmt.Println("1.2.3 is older than 1.3.0")
    }
}
```

### Pre-release vs release

```go
beta := versions.NewVersion("1.0.0-beta")
release := versions.NewVersion("1.0.0")

// Empty suffix (release) > any suffix (pre-release)
if beta.IsOlderThan(release) {
    fmt.Println("Pre-release is older than release") // prints
}

// Alpha < Beta by semantic weight
alpha := versions.NewVersion("1.0.0-alpha")
if alpha.IsOlderThan(beta) {
    fmt.Println("Alpha is older than Beta") // prints
}
```

### Different-length version numbers

```go
short := versions.NewVersion("1.0")
full := versions.NewVersion("1.0.0")

// VersionNumbers.CompareTo: longer array is greater when prefix matches
result := short.CompareTo(full)
fmt.Println(result) // -1 (1.0 < 1.0.0 in current implementation)

// Direct number-part comparison
nums1 := versions.NewVersionNumbers([]int{1, 2, 0})
nums2 := versions.NewVersionNumbers([]int{1, 2, 3})
if nums1.CompareTo(nums2) < 0 {
    fmt.Println("1.2.0 < 1.2.3")
}
```

### Range checking

```go
v := versions.NewVersion("1.5.0")
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")

if v.IsBetween(low, high) {
    fmt.Println("1.5.0 is between 1.0.0 and 2.0.0")
}

// Open-ended: nil skips one bound
v.IsBetween(nil, high)   // true if v <= 2.0.0
v.IsBetween(low, nil)    // true if v >= 1.0.0
```

### Suffix weight comparison

```go
s1 := versions.VersionSuffix("-alpha1")
s2 := versions.VersionSuffix("-beta2")

// Known suffixes compared by semantic weight
if s1.CompareTo(s2) < 0 {
    fmt.Println("alpha1 < beta2") // alpha(100) < beta(200)
}

// Same type, different sub-version
s3 := versions.VersionSuffix("-alpha1")
s4 := versions.VersionSuffix("-alpha3")
fmt.Println(s3.CompareTo(s4)) // -1 (sub-version: 1 < 3)

// Unknown suffixes sort after known suffixes
s5 := versions.VersionSuffix("-custom")
fmt.Println(s5.CompareTo(s1)) // 1 (unknown sorts after known)
```

## Important Notes

- **Comparison order is VersionNumbers -> Suffix -> PublicTime -> Raw string**. This is the actual order in the source code; Suffix is compared before PublicTime, not after.
- **`1.0` and `1.0.0` are NOT equal** in the current implementation. `VersionNumbers.CompareTo` treats shorter arrays as less than longer arrays when the shared prefix matches, so `[1,0] < [1,0,0]`.
- **Empty suffix (release) is always greater than any non-empty suffix (pre-release)**. This means `1.0.0 > 1.0.0-beta > 1.0.0-alpha`.
- **Suffix comparison uses semantic weight, not pure alphabetical order**. Known suffixes (alpha, beta, rc, etc.) are sorted by their weight values; unknown suffixes fall after all known suffixes.
- **PublicTime is only compared when both versions have non-zero PublicTime**. If one or both have zero time, this step is skipped.
- **Prefix is NOT part of the CompareTo comparison**. Prefixes like "v" do not affect ordering; `v1.0.0` and `1.0.0` compare equal on the number+suffix+time dimensions and differ only on Raw string.
- **IsBetween is inclusive on both bounds** (`low <= x <= high`). Pass `nil` for either bound to make it open-ended.