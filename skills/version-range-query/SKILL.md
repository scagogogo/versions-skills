---
name: version-range-query
description: Query versions within a specific range, check if a version falls between two boundaries, or implement version constraints with inclusion/exclusion policies. Use when finding versions in an interval, dependency resolution, or compatibility checks.
argument-hint: <start> <end> [version1 version2 ...]
---

# Version Range Query

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- You need to find all versions between a start and end version
- You need to implement interval-based version filtering with inclusion/exclusion control
- You need to check if a single version falls between two boundaries
- You need cross-group range queries on sorted version groups
- You are implementing dependency version resolution or compatibility checks

## Decision Tree

```
Need to query versions within a range?
├─ Single version check (yes/no)?
│   ├─ Inclusive both ends?              → Version.IsBetween(low, high)
│   └─ Custom inclusion policy?          → version_range_query with check_version
├─ Filter a version list by range?
│   ├─ Simple inclusive/exclusive?       → versions range <start> <end> [versions...]
│   └─ Programmatic with tuples?         → SortedVersionGroups.QueryRange() / VersionGroup.QueryRangeVersions()
├─ From a file?                          → versions range --from-file <path>
└─ Open-ended range (one bound only)?    → IsBetween with nil / use "999.0.0" as far end
```

## Task Patterns

### Check if a single version is between two bounds

**Goal:** Is `"1.5.0"` between `"1.0.0"` and `"2.0.0"`?

**SDK approach:**
```go
v := versions.NewVersion("1.5.0")
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")
inRange := v.IsBetween(low, high) // true (inclusive both ends)

// Open-ended: pass nil to skip one bound
v.IsBetween(nil, high)   // true (no lower bound)
v.IsBetween(low, nil)    // true (no upper bound)
```

**CLI approach:**
```bash
versions check 1.5.0 --between-low 1.0.0 --between-high 2.0.0  # exit 0
```

**MCP approach:**
```json
{
  "tool": "version_range_query",
  "arguments": {"check_version": "1.5.0", "low": "1.0.0", "high": "2.0.0"}
}
```

### Query a version list for all versions in a range

**Goal:** From `["1.0.0", "1.1.0", "1.5.0", "2.0.0", "2.1.0", "3.0.0"]`, find all in `[1.0.0, 2.0.0)`.

**SDK approach:**
```go
versionList := versions.NewVersions("1.0.0", "1.1.0", "1.5.0", "2.0.0", "2.1.0", "3.0.0")
sortedGroups := versions.NewSortedVersionGroups(versionList)
start := tuple.New2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
end := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
result := sortedGroups.QueryRange(start, end)
// result = [1.0.0, 1.1.0, 1.5.0]
```

**CLI approach:**
```bash
versions range 1.0.0 2.0.0 1.0.0 1.1.0 1.5.0 2.0.0 2.1.0 3.0.0
# Default: include-start=true, include-end=false → [1.0.0, 2.0.0)
# Output: 1.0.0 1.1.0 1.5.0

# Inclusive both ends:
versions range 1.0.0 2.0.0 --include-end 1.0.0 1.1.0 1.5.0 2.0.0 2.1.0
# Output: 1.0.0 1.1.0 1.5.0 2.0.0
```

**MCP approach:**
```json
{
  "tool": "version_range_query",
  "arguments": {
    "versions": ["1.0.0", "1.1.0", "1.5.0", "2.0.0", "2.1.0", "3.0.0"],
    "start": "1.0.0",
    "end": "2.0.0",
    "include_start": true,
    "include_end": false
  }
}
```

### Range query within a single version group

**Goal:** In the `"1.0"` group, find versions in `[1.0.1, 1.0.3)`.

**SDK approach:**
```go
groupMap := versions.Group(versionList)
group := groupMap["1.0"]
gStart := tuple.New2(versions.NewVersion("1.0.1"), versions.ContainsPolicyYes)
gEnd := tuple.New2(versions.NewVersion("1.0.3"), versions.ContainsPolicyNo)
result := group.QueryRangeVersions(gStart, gEnd)
// result = [1.0.1, 1.0.2]
```

**CLI approach:**
Not directly available for single-group range — use full `versions range` and filter by group separately.

**MCP approach:**
Not directly available — use `version_range_query` with the full list, then filter by group.

### Read versions from a file and query range

**Goal:** Read versions from `releases.txt` and find those in `[2.0.0, 3.0.0]`.

**SDK approach:**
```go
// Use version_read_file to read, then QueryRange
```

**CLI approach:**
```bash
versions range 2.0.0 3.0.0 --from-file releases.txt --include-end
```

**MCP approach:**
```json
{"tool": "version_read_file", "arguments": {"file_path": "releases.txt"}}
```
Then pass the result to `version_range_query`.

### Open-ended range query (all versions >= 2.0.0)

**Goal:** Find all versions from `"2.0.0"` onward.

**SDK approach:**
```go
start := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyYes)
end := tuple.New2(versions.NewVersion("999.0.0"), versions.ContainsPolicyNo)
result := sortedGroups.QueryRange(start, end)
```

**CLI approach:**
```bash
# Use a high sentinel version as the end
versions range 2.0.0 999.0.0 --include-end 1.0.0 2.0.0 2.1.0 3.0.0
```

**MCP approach:**
```json
{
  "tool": "version_range_query",
  "arguments": {
    "versions": ["1.0.0", "2.0.0", "2.1.0", "3.0.0"],
    "start": "2.0.0",
    "end": "999.0.0",
    "include_start": true,
    "include_end": true
  }
}
```

## API Reference

### SDK

```go
// Simple between check (inclusive both ends, nil boundaries skipped)
func (x *Version) IsBetween(low, high *Version) bool

// Range query within a single version group
func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version

// Cross-group range query across all groups
func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version

// Boundary inclusion policy
type ContainsPolicy int
const (
    ContainsPolicyNone ContainsPolicy = iota  // defaults to inclusive
    ContainsPolicyYes                          // [ or ]
    ContainsPolicyNo                           // ( or )
)
```

### CLI Commands

```bash
versions range <start> <end> [versions...]    # range query
  --include-start        # include start boundary (default: true)
  --include-end          # include end boundary (default: false)
  --from-file <path>     # read versions from file

versions check <v> --between-low <low> --between-high <high>  # between check
```

### MCP Tools

| Tool | Key Arguments | Returns |
|------|--------------|---------|
| `version_range_query` | `versions[]`, `start`, `end`, `include_start`, `include_end` | matching versions (sorted) |
| `version_range_query` | `check_version`, `low`, `high` | `{between: bool}` |

## ContainsPolicy Reference

| Policy | Symbol | Meaning |
|--------|--------|---------|
| `ContainsPolicyYes` | `[` / `]` | Include boundary version |
| `ContainsPolicyNo` | `(` / `)` | Exclude boundary version |
| `ContainsPolicyNone` | (defaults to Yes) | Unspecified — treated as inclusive |

Tuple construction (requires `github.com/golang-infrastructure/go-tuple`):
```go
start := tuple.New2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
end := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
```

## Cross-References

- [[version-comparison]] — for the IsBetween check and underlying CompareTo logic
- [[version-constraints]] — for constraint expression matching (caret, tilde, wildcards)
- [[version-grouping]] — for grouping versions before cross-group range queries
- [[version-sorting]] — range query results are always returned sorted

## Important Notes

- **QueryRange on SortedVersionGroups returns nil** if the start version's group does not exist in the collection — always nil-check the result
- **IsBetween is inclusive on both ends** (`low <= x <= high`); nil boundaries are skipped
- **For open-ended range queries in SDK**, use a very large version as the end boundary (e.g., `"999.0.0"`)
- **Range queries always return pre-sorted results** in ascending order
- **CLI defaults**: `--include-start` defaults to true, `--include-end` defaults to false — matching the common `[start, end)` convention
- **ContainsPolicyNone defaults to ContainsPolicyYes** (inclusive) in most SDK contexts
- **Import `github.com/golang-infrastructure/go-tuple`** for `tuple.New2` in SDK range queries
- **MCP supports both modes**: range queries (start/end) and between-ness checks (check_version/low/high) in the same `version_range_query` tool
