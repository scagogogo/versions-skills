---
name: version-sorting
description: Sort a list of version strings or Version objects in natural (semantic) order. Use when you need to order versions, find the latest/oldest in a collection, or display versions in ascending/descending order.
argument-hint: <version1> <version2> ... <versionN>
---

# Version Sorting

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- You have a list of version strings and need them in natural order
- You need to find the latest or oldest version in a collection
- You need to display versions in ascending or descending order
- You are implementing version selection UI or dependency resolution
- You need to sort versions read from a file

## Decision Tree

```
Need to sort versions?
├─ Input is []string?                → SortVersionStringSlice() / versions sort / version_sort
├─ Input is []*Version?              → SortVersionSlice() or VersionSlice + sort.Sort()
├─ Need descending order?            → SortVersionStringSlice then reverse / --desc / descending:true
├─ Input is from a file?             → versions sort --from-file <path>
├─ Need to sort version groups?      → SortVersionGroupMap() or SortVersionGroupSlice()
└─ Need latest/oldest after sort?    → Sort ascending, then take first (oldest) or last (latest)
```

## Task Patterns

### Sort version strings in ascending order

**Goal:** Sort `["2.0.0", "1.0.0", "1.10.0", "1.2.0"]` to `["1.0.0", "1.2.0", "1.10.0", "2.0.0"]`.

**SDK approach:**
```go
sorted := versions.SortVersionStringSlice([]string{"2.0.0", "1.0.0", "1.10.0", "1.2.0"})
// sorted = ["1.0.0", "1.2.0", "1.10.0", "2.0.0"]
// Original slice is NOT modified
```

**CLI approach:**
```bash
versions sort 2.0.0 1.0.0 1.10.0 1.2.0
# Output: 1.0.0 1.2.0 1.10.0 2.0.0
```

**MCP approach:**
```json
{"tool": "version_sort", "arguments": {"versions": ["2.0.0", "1.0.0", "1.10.0", "1.2.0"]}}
```

### Sort in descending order (latest first)

**Goal:** Get `["2.0.0", "1.10.0", "1.2.0", "1.0.0"]`.

**SDK approach:**
```go
sorted := versions.SortVersionStringSlice(versions)
// Then reverse in place or iterate backward
for i := len(sorted) - 1; i >= 0; i-- {
    fmt.Println(sorted[i])
}
```

**CLI approach:**
```bash
versions sort --desc 2.0.0 1.0.0 1.10.0 1.2.0
# Output: 2.0.0 1.10.0 1.2.0 1.0.0
```

**MCP approach:**
```json
{"tool": "version_sort", "arguments": {"versions": ["2.0.0", "1.0.0", "1.10.0", "1.2.0"], "descending": true}}
```

### Sort Version objects and find latest/oldest

**Goal:** Find the newest and oldest version in a collection of Version objects.

**SDK approach:**
```go
versionList := versions.NewVersions("2.0.0", "1.0.0", "1.10.0")
sortedVersions := versions.SortVersionSlice(versionList)
latest := sortedVersions[len(sortedVersions)-1]  // 2.0.0
oldest := sortedVersions[0]                       // 1.0.0
```

**CLI approach:**
```bash
versions sort --desc 2.0.0 1.0.0 1.10.0 | head -1  # latest
versions sort 2.0.0 1.0.0 1.10.0 | head -1          # oldest
```

**MCP approach:**
```json
{"tool": "version_sort", "arguments": {"versions": ["2.0.0", "1.0.0", "1.10.0"], "descending": true}}
```

### Sort versions from a file

**Goal:** Read versions from `releases.txt` and sort them.

**SDK approach:**
```go
// Use version_read_file to read, then sort
// See [[version-read-write]] for file I/O
```

**CLI approach:**
```bash
versions sort --from-file releases.txt
versions sort --desc --from-file releases.txt
```

**MCP approach:**
```json
{"tool": "version_read_file", "arguments": {"file_path": "releases.txt"}}
```
Then pass the result to `version_sort`.

### Use VersionSlice with standard library sort

**Goal:** Use Go's `sort.Sort()` directly on a version slice.

**SDK approach:**
```go
slice := versions.VersionSlice(versions.NewVersions("3.0.0", "1.0.0", "2.0.0"))
sort.Sort(slice) // no closures needed — VersionSlice implements sort.Interface
// slice is now [1.0.0, 2.0.0, 3.0.0]
```

**CLI approach:**
Not applicable — CLI handles sorting internally.

**MCP approach:**
Not applicable — MCP handles sorting internally.

### Sort version groups

**Goal:** Sort a map of version groups into a deterministic order.

**SDK approach:**
```go
groupMap := versions.Group(versionList)
sortedGroups := versions.SortVersionGroupMap(groupMap)
for _, g := range sortedGroups {
    fmt.Println(g.ID()) // groups in sorted order
}
```

**CLI approach:**
```bash
versions group 1.0.0 1.1.0 2.0.0 2.1.0 3.0.0
```

**MCP approach:**
```json
{"tool": "version_group", "arguments": {"versions": ["1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0"]}}
```

## Cross-References

- [[version-comparison]] — for the underlying CompareTo logic used by sorting
- [[version-parsing]] — for parsing version strings before sorting
- [[version-grouping]] — for grouping versions before sorting groups
- [[version-range-query]] — for querying sorted ranges

## Important Notes

- `SortVersionStringSlice` and `SortVersionSlice` do NOT modify the original input — they return new slices
- `SortVersionGroupSlice` DOES modify the input slice in-place — unlike the other sort functions
- `VersionSlice` implements `sort.Interface` — use `sort.Sort()` directly, no closures needed
- `"1.10.0"` correctly sorts after `"1.2.0"` — numeric comparison, not alphabetical
- Pre-release versions sort before their release counterparts: `"1.0.0-beta"` < `"1.0.0"`
- CLI `--from-file` reads one version per line, ignores blank lines and `#` comments
- MCP `descending` defaults to `false` (ascending order)
