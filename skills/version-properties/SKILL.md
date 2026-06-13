---
name: version-properties
description: Use when extracting specific properties from a version number — segments, sub-version, suffix weight, pure prefix, group ID. Covers Segments, SubVersion, SuffixWeight, PurePrefix, BuildGroupID, Clone via SDK, CLI, and MCP.
argument-hint: <version-property-task>
---

# Version Properties Skill

## When to Use

- User needs to extract numeric segments from a version string
- User needs the sub-version number from a prerelease suffix (e.g. "beta2" → 2)
- User needs the semantic weight of a suffix for sorting/comparison
- User needs the prefix without trailing delimiters (e.g. "curl-" → "curl")
- User needs the group ID for version grouping (e.g. "v1.2.3-beta1" → "1.2.3")
- User needs to deep-copy a version for independent modification

## Quick Start

**SDK:**
```go
v := versions.NewVersion("v1.2.3-rc2")
v.Segments()       // [1, 2, 3]
v.SubVersion()     // 2
v.SuffixWeight()   // rc (400)
v.BuildGroupID()   // "1.2.3"
```

**CLI:**
```bash
versions segments v1.2.3-rc2        # [1, 2, 3]
versions sub-version v1.2.3-rc2     # 2
versions suffix-weight v1.2.3-rc2   # rc (400)
versions pure-prefix curl-7.85.0    # curl
versions group-id v1.2.3-beta1      # 1.2.3
```

**MCP:**
```
version_info(version_string="v1.2.3-rc2")  # returns all properties
```

## API Reference — SDK

### Segments

**func (v *Version) Segments() []int**

Returns the version number segments as `[]int`. E.g. `"1.2.3"` → `[1, 2, 3]`.

**func (v *Version) Segments64() []int64**

Same as `Segments()` but returns `[]int64` for large version numbers.

### Individual Numbers

| Method | Returns |
|--------|---------|
| `Major() int` | First segment (0 if missing) |
| `Minor() int` | Second segment (0 if missing) |
| `Patch() int` | Third segment (0 if missing) |

### SubVersion

**func (v *Version) SubVersion() int**

Extracts the numeric part from the suffix. E.g. `"-beta2"` → `2`, `"-rc1"` → `1`. Returns `0` if no numeric suffix.

### SuffixWeight

**func (v *Version) SuffixWeight() SuffixWeight**

Returns the semantic weight of the version's suffix, used for comparison ordering:

| Weight | Suffix Type | Value |
|--------|------------|-------|
| 0 | unknown (no suffix / unclassified) | 0 |
| 50 | dev | 50 |
| 60 | snapshot | 60 |
| 70 | nightly | 70 |
| 100 | alpha | 100 |
| 200 | beta | 200 |
| 300 | milestone | 300 |
| 400 | rc | 400 |
| 500 | final / release / ga | 500 |
| 600 | sp | 600 |
| 700 | patch | 700 |
| 800 | post | 800 |

**func GetSuffixWeight(suffix string) SuffixWeight**

Package-level function to get the weight of any suffix string.

### PurePrefix

**func (p VersionPrefix) PurePrefix() string**

Returns the prefix without trailing delimiter characters. E.g. `"curl-"` → `"curl"`, `"v"` → `"v"`.

### BuildGroupID

**func (v *Version) BuildGroupID() string**

Returns the group ID by joining VersionNumbers with `.`. E.g. `"v1.2.3-beta1"` → `"1.2.3"`.

### Clone

**func (v *Version) Clone() *Version**

Returns a deep copy of the version. The clone is completely independent — modifying it does not affect the original.

### RawString

**func (v *Version) RawString() string**

Returns the original version string as parsed. Unlike `String()` which returns JSON, this returns the human-readable form.

## CLI Commands

| Command | Example | Output |
|---------|---------|--------|
| `segments` | `versions segments v1.2.3` | `{"segments": [1, 2, 3]}` |
| `sub-version` | `versions sub-version 1.2.3-beta2` | `{"sub_version": 2}` |
| `suffix-weight` | `versions suffix-weight 1.2.3-rc1` | `{"suffix_weight": "rc", "weight_value": 400}` |
| `pure-prefix` | `versions pure-prefix curl-7.85.0` | `{"prefix": "curl-", "pure_prefix": "curl"}` |
| `group-id` | `versions group-id v1.2.3-beta1` | `{"group_id": "1.2.3"}` |
| `clone` | `versions clone v1.2.3` | Full version object |

## MCP Tools

Use `version_info` to get all properties in one call:

```
version_info(version_string="v1.2.3-rc1")
```

Returns: raw, valid, major, minor, patch, version_numbers, prefix, suffix, suffix_weight, group_id, and all Is* flags.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    v := versions.NewVersion("v1.2.3-rc2")

    // Segments
    fmt.Println(v.Segments())        // [1 2 3]
    fmt.Println(v.Major())           // 1
    fmt.Println(v.Minor())           // 2
    fmt.Println(v.Patch())           // 3

    // Sub-version from suffix
    fmt.Println(v.SubVersion())      // 2

    // Suffix weight for ordering
    w := v.SuffixWeight()
    fmt.Println(w.String())          // "rc"
    fmt.Println(int(w))              // 400

    // Get suffix weight without a version
    w2 := versions.GetSuffixWeight("-beta3")
    fmt.Println(w2.String())         // "beta"

    // Group ID
    fmt.Println(v.BuildGroupID())    // "1.2.3"

    // Pure prefix
    v2 := versions.NewVersion("curl-7.85.0")
    fmt.Println(v2.Prefix.String())       // "curl-"
    fmt.Println(v2.Prefix.PurePrefix())   // "curl"

    // Clone for independent modification
    cloned := v.Clone()
    _ = cloned.WithSuffix("-final")  // cloned modified, v unchanged

    // RawString vs String
    fmt.Println(v.RawString())       // "v1.2.3-rc2"
}
```

## Important Notes

- `Segments()` returns `[]int` while `Segments64()` returns `[]int64` — use the latter for very large version numbers
- `SubVersion()` returns `0` if the suffix has no numeric part (e.g. `"-beta"` → `0`, `"-beta2"` → `2`)
- `SuffixWeight` determines suffix comparison order in `CompareTo`: `dev < snapshot < nightly < alpha < beta < milestone < rc < final/release/ga < sp < patch < post`
- `PurePrefix()` strips trailing `-`, `.`, and `_` from the prefix — useful for package name extraction
- `BuildGroupID()` is the key used by `Group()` to group versions — versions `1.2.3` and `1.2.3-alpha` share group ID `"1.2.3"`
- `Clone()` creates a fully independent deep copy — safe for concurrent modification
- `RawString()` preserves the original string; `String()` returns a different format
