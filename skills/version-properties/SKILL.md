---
name: version-properties
description: Extract segments, sub-version, suffix weight, pure prefix, or group ID from a parsed version.
argument-hint: <property-to-extract>
---

# Version Properties

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- Extracting numeric segments from a version string (`[1, 2, 3]`)
- Getting the sub-version number embedded in a prerelease suffix (`beta2` -> `2`)
- Determining the semantic weight of a suffix for ordering (`rc` = 400)
- Getting the clean prefix without trailing delimiters (`"curl-"` -> `"curl"`)
- Getting the group ID used by `Group()` for version grouping

## Decision Tree

```
Need all components at once?
  → Use version_info (MCP) or version info (CLI)
Need numeric segments only?
  → Use Segments() / Segments64()
Need suffix semantic weight?
  → Use SuffixWeight() or GetSuffixWeight()
Need the group key for Group()?
  → Use BuildGroupID()
Need a clean prefix string?
  → Use Prefix.PurePrefix()
```

## Task Patterns

### Get all numeric segments

**Goal:** Extract the version number components as a slice.

**SDK approach:**
```go
v := versions.NewVersion("v1.2.3-rc2")
segs := v.Segments()    // []int{1, 2, 3}
major := v.Major()      // 1
minor := v.Minor()      // 2
patch := v.Patch()      // 3
```

**CLI approach:**
```bash
versions segments v1.2.3-rc2   # {"segments": [1, 2, 3]}
```

**MCP approach:**
```
version_info(version_string="v1.2.3-rc2")
# Response includes: major, minor, patch, version_numbers
```

### Determine suffix type and weight

**Goal:** Classify the prerelease suffix and get its ordering weight.

**SDK approach:**
```go
v := versions.NewVersion("1.2.3-rc1")
weight := v.SuffixWeight()  // SuffixWeight(400)
name := weight.String()     // "rc"
// Package-level: versions.GetSuffixWeight("-beta3") -> beta (200)
```

**CLI approach:**
```bash
versions suffix-weight 1.2.3-rc1
# {"suffix_weight": "rc", "weight_value": 400}
```

**MCP approach:**
```
version_info(version_string="1.2.3-rc1")
# Response includes: suffix_weight
```

### Extract sub-version number from suffix

**Goal:** Get the numeric part of a prerelease suffix (e.g., `beta2` -> `2`).

**SDK approach:**
```go
v := versions.NewVersion("1.2.3-beta2")
sub := v.SubVersion()  // 2
// Returns 0 if suffix has no number (e.g., "-beta")
```

**CLI approach:**
```bash
versions sub-version 1.2.3-beta2   # {"sub_version": 2}
```

**MCP approach:**
```
version_info(version_string="1.2.3-beta2")
```

### Get the group ID for grouping

**Goal:** Obtain the key used by `Group()` to cluster versions.

**SDK approach:**
```go
v := versions.NewVersion("v1.2.3-beta1")
gid := v.BuildGroupID()  // "1.2.3" — prefix + suffix stripped
// Versions "1.2.3" and "1.2.3-alpha" both have group ID "1.2.3"
```

**CLI approach:**
```bash
versions group-id v1.2.3-beta1   # {"group_id": "1.2.3"}
```

**MCP approach:**
```
version_info(version_string="v1.2.3-beta1")
# Response includes: group_id
```

### Clean a prefix string

**Goal:** Strip trailing delimiter characters from a prefix.

**SDK approach:**
```go
v := versions.NewVersion("curl-7.85.0")
pure := v.Prefix.PurePrefix()  // "curl" (strips trailing "-")
```

**CLI approach:**
```bash
versions pure-prefix curl-7.85.0
# {"prefix": "curl-", "pure_prefix": "curl"}
```

### Deep-copy a version

**Goal:** Create an independent copy safe for mutation.

**SDK approach:**
```go
cloned := v.Clone()  // deep copy — modifying cloned does not affect v
```

**CLI approach:**
```bash
versions clone v1.2.3
```

## Cross-References

- [[version-check]] — boolean checks on version type (IsBeta, IsStable, etc.)
- [[version-grouping]] — Group() uses BuildGroupID() for clustering
- [[version-comparison]] — CompareTo uses SuffixWeight for ordering
- [[version-mutation]] — modify version components after inspecting them

## Important Notes

- **Suffix weight order**: dev(50) < snapshot(60) < nightly(70) < alpha(100) < beta(200) < milestone(300) < rc(400) < final/release/ga(500) < sp(600) < patch(700) < post(800)
- **SubVersion() returns 0** if the suffix has no numeric part (`"-beta"` -> `0`, `"-beta2"` -> `2`).
- **PurePrefix() strips trailing `-`, `.`, `_`** — use it to extract a clean package name.
- **BuildGroupID()** is the key for `Group()` — versions differing only by suffix share the same group.
- **Clone() creates a fully independent deep copy** — safe for concurrent modification.
- **RawString()** returns the original parsed string; **String()** returns a different (JSON) format.
