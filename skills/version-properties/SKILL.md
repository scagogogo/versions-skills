---
name: version-properties
description: Extract segments, sub-version, suffix weight, pure prefix, or group ID from a parsed version.
argument-hint: <property-to-extract>
---

# Version Properties

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

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

| Layer | Approach |
|-------|----------|
| SDK | `v.Segments()`, `v.Major()`, `v.Minor()`, `v.Patch()` |
| CLI | `versions segments v1.2.3-rc2` |
| MCP | `{"tool": "version_info", "arguments": {"version_string": "v1.2.3-rc2"}}` → major, minor, patch, version_numbers |

### Determine suffix type and weight

**Goal:** Classify the prerelease suffix and get its ordering weight.

| Layer | Approach |
|-------|----------|
| SDK | `v.SuffixWeight()` → SuffixWeight(400); `versions.GetSuffixWeight("-beta3")` → beta (200) |
| CLI | `versions suffix-weight 1.2.3-rc1` |
| MCP | `{"tool": "version_info", "arguments": {"version_string": "1.2.3-rc1"}}` → suffix_weight field |

### Extract sub-version number from suffix

**Goal:** Get the numeric part of a prerelease suffix (e.g., `beta2` -> `2`).

| Layer | Approach |
|-------|----------|
| SDK | `v.SubVersion()` → 2 (returns 0 if suffix has no number) |
| CLI | `versions sub-version 1.2.3-beta2` |
| MCP | `{"tool": "version_info", "arguments": {"version_string": "1.2.3-beta2"}}` |

### Get the group ID for grouping

**Goal:** Obtain the key used by `Group()` to cluster versions.

| Layer | Approach |
|-------|----------|
| SDK | `v.BuildGroupID()` → `"1.2.3"` from `"v1.2.3-beta1"` |
| CLI | `versions group-id v1.2.3-beta1` |
| MCP | `{"tool": "version_info", "arguments": {"version_string": "v1.2.3-beta1"}}` → group_id field |

### Clean a prefix string

**Goal:** Strip trailing delimiter characters from a prefix.

| Layer | Approach |
|-------|----------|
| SDK | `v.Prefix.PurePrefix()` → `"curl"` from `"curl-"` |
| CLI | `versions pure-prefix curl-7.85.0` |
| MCP | `{"tool": "version_info", "arguments": {"version_string": "curl-7.85.0"}}` |

### Deep-copy a version

**Goal:** Create an independent copy safe for mutation.

| Layer | Approach |
|-------|----------|
| SDK | `cloned := v.Clone()` — deep copy, modifying cloned does not affect v |
| CLI | `versions clone v1.2.3` |
| MCP | Use `version_parse` and reconstruct with `version_build` |

## API Reference

### SDK — Segment Accessors

```go
v.Major() int           // VersionNumbers[0] or 0
v.Minor() int           // VersionNumbers[1] or 0
v.Patch() int           // VersionNumbers[2] or 0
v.Segments() []int      // all VersionNumbers
v.Segments64() []int64  // as int64
```

### SDK — Suffix Properties

```go
v.SubVersion() int                    // numeric from suffix (e.g. "beta2" → 2, "-beta" → 0)
v.SuffixWeight() SuffixWeight         // semantic ordering weight on the Version
versions.GetSuffixWeight(s string) SuffixWeight  // package-level: weight of any suffix string
```

Suffix weight ordering:

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

### SDK — Prefix, Group, Clone, Serialization

```go
v.Prefix.PurePrefix() string   // prefix without trailing delimiters, e.g. "curl-" → "curl"
v.BuildGroupID() string        // group key, e.g. "v1.2.3-beta1" → "1.2.3"
v.Clone() *Version             // deep copy, independent of original
v.RawString() string           // original parsed string (human-readable, unlike String())
```

### CLI Commands

```bash
versions segments <v>            # {"segments": [1, 2, 3]}
versions sub-version <v>         # {"sub_version": 2}
versions suffix-weight <v>       # {"suffix_weight": "rc", "weight_value": 400}
versions pure-prefix <v>         # {"prefix": "curl-", "pure_prefix": "curl"}
versions group-id <v>            # {"group_id": "1.2.3"}
versions clone <v>               # full version object as JSON
```

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_info` | `version_string` | raw, valid, major, minor, patch, version_numbers, prefix, suffix, suffix_weight, group_id, all Is* flags |

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
