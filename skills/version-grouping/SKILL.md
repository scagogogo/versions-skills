---
name: version-grouping
description: Group versions by their version number series (major.minor). Use when you need to organize a large collection of versions into groups, find versions within a specific series, or query group-level aggregates (latest, oldest, stable, prerelease).
argument-hint: <version1> <version2> ... <versionN>
---

# Version Grouping

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

## When to Use

- You need to organize a large collection of versions into version series
- You need to find all versions in a specific major.minor group (e.g., all 1.0.x versions)
- You need the latest, oldest, stable, or prerelease version within a group
- You need to navigate groups by index or check group membership
- You are building a version selector UI with hierarchical version choices

## Decision Tree

```
Need to organize versions into groups?
├─ Get all groups?                    → Group() / versions group / version_group
├─ Need sorted groups?                → NewSortedVersionGroups()
├─ List group IDs only?               → SortedVersionGroups.GroupIDs() / versions group-ids
├─ Get a specific group by ID?        → SortedVersionGroups.Get("1.0") / versions group --id "1.0"
├─ Find latest in a group?            → VersionGroup.GetLatest() / versions group-latest
├─ Find latest stable in a group?     → VersionGroup.LatestStable() / versions group-latest-stable
├─ Check if version exists in group?  → VersionGroup.Contains() / versions group-contains
└─ Get group ID for a single version? → Version.BuildGroupID() / versions group-id
```

## Task Patterns

### Group versions by their number series

**Goal:** Group `["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.1.0"]` into series.

**SDK approach:**
```go
versionList := versions.NewVersions("1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.1.0")
sortedGroups := versions.NewSortedVersionGroups(versionList)
for _, id := range sortedGroups.GroupIDs() {
    group := sortedGroups.Get(id)
    fmt.Printf("Group %s: %d versions\n", id, group.Count())
}
```

**CLI approach:**
```bash
versions group 1.0.0 1.0.1 1.1.0 2.0.0 2.1.0
versions group-ids 1.0.0 1.0.1 1.1.0 2.0.0 2.1.0
```

**MCP approach:**
```json
{"tool": "version_group", "arguments": {"versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.1.0"]}}
```

### Find latest version in a group

**Goal:** Get the newest version in the `"1.0"` group.

**SDK approach:**
```go
sortedGroups := versions.NewSortedVersionGroups(versionList)
group := sortedGroups.Get("1.0")
if group != nil {
    latest := group.GetLatest() // returns nil if group is empty
    fmt.Println(latest.RawString())
}
```

**CLI approach:**
```bash
versions group-latest --group-id "1.0" 1.0.0 1.0.1 1.0.2
```

**MCP approach:**
```json
{"tool": "version_group", "arguments": {"versions": ["1.0.0", "1.0.1", "1.0.2"], "group_id": "1.0", "operation": "latest"}}
```

### Find latest stable and prerelease in a group

**Goal:** In group `"2.0"`, find the latest stable and latest prerelease separately.

**SDK approach:**
```go
group := sortedGroups.Get("2.0")
stable := group.LatestStable()         // nil if none
prerelease := group.LatestPrerelease() // nil if none
```

**CLI approach:**
```bash
versions group-latest-stable --group-id "2.0" 2.0.0 2.0.1-rc1 2.0.1
versions group-latest-prerelease --group-id "2.0" 2.0.0 2.0.1-rc1 2.0.1
```

**MCP approach:**
```json
{"tool": "version_group", "arguments": {"versions": ["2.0.0", "2.0.1-rc1", "2.0.1"], "group_id": "2.0", "operation": "latest_stable"}}
```

### Manipulate a version group (add, remove, check, filter)

**Goal:** Build a group, add versions, check membership, filter by predicate.

**SDK approach:**
```go
group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.0.0", "1.0.1"))
group.Add(versions.NewVersion("1.0.3"))
group.Remove(versions.NewVersion("1.0.0"))
exists := group.Contains(versions.NewVersion("1.0.1")) // true

// Filter with a predicate
stableOnly := group.Filter(func(v *versions.Version) bool {
    return v.IsStable()
})
```

**CLI approach:**
```bash
versions group-contains --group-id "1.0" --version "1.0.1" 1.0.0 1.0.1 1.0.2
```

**MCP approach:**
```json
{"tool": "version_group", "arguments": {"versions": ["1.0.0", "1.0.1", "1.0.2"], "group_id": "1.0", "operation": "contains", "target_version": "1.0.1"}}
```

### Navigate sorted groups by index

**Goal:** Access the first and last group by position.

**SDK approach:**
```go
sortedGroups := versions.NewSortedVersionGroups(versionList)
first := sortedGroups.At(0)                         // nil if empty
last := sortedGroups.At(sortedGroups.Len() - 1)     // nil if empty
hasGroup := sortedGroups.Contains("1.0")            // true/false
allVersions := sortedGroups.Versions()              // all versions across groups, sorted
```

**CLI approach:**
Not directly available — use `versions group` to see all groups.

**MCP approach:**
Not directly available — use `version_group` to see all groups.

### Get the group ID for a single version

**Goal:** Determine that `"1.2.3"` belongs to group `"1.2"`.

**SDK approach:**
```go
v := versions.NewVersion("1.2.3")
groupId := v.BuildGroupID() // "1.2.3" — full number string as key
```

**CLI approach:**
```bash
versions group-id 1.2.3     # Output: 1.2
```

**MCP approach:**
Not directly available — use `version_parse` and extract the BuildGroupID.

## API Reference

### SDK — Grouping Functions

```go
// Group versions by BuildGroupID (full number sequence).
// Returns a map of groupID → *VersionGroup.
func Group(versions []*Version) map[string]*VersionGroup

// Create a pre-sorted version groups structure from a version list.
// Efficient for repeated queries.
func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups
```

### SDK — VersionGroup Type

```go
type VersionGroup struct {
    GroupVersionNumbers VersionNumbers
    VersionMap          map[string]*Version
}

// Constructors
func NewVersionGroup(groupVersionNumbers VersionNumbers) *VersionGroup
func NewVersionGroupFromVersions(versions []*Version) *VersionGroup

// Core methods
func (x *VersionGroup) Add(v *Version) bool           // add version; true if already existed
func (x *VersionGroup) Contains(v *Version) bool      // membership check
func (x *VersionGroup) Remove(v *Version) bool        // remove; true if it existed
func (x *VersionGroup) ID() string                    // group ID (e.g. "1.2")
func (x *VersionGroup) CompareTo(target *VersionGroup) int
func (x *VersionGroup) Count() int                    // number of versions in group

// Version access
func (x *VersionGroup) Versions() []*Version          // all versions (unordered)
func (x *VersionGroup) SortVersions() []*Version      // sorted versions (ascending)
func (x *VersionGroup) GetLatest() *Version           // newest; nil if empty
func (x *VersionGroup) GetOldest() *Version           // oldest; nil if empty

// Filtering
func (x *VersionGroup) StableVersions() []*Version    // all stable (no suffix) versions
func (x *VersionGroup) PrereleaseVersions() []*Version // all prerelease versions
func (x *VersionGroup) LatestStable() *Version        // newest stable; nil if none
func (x *VersionGroup) LatestPrerelease() *Version    // newest prerelease; nil if none
func (x *VersionGroup) Filter(predicate func(*Version) bool) []*Version

// Range query within group (requires go-tuple)
func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version
```

### SDK — SortedVersionGroups Type

```go
type SortedVersionGroups struct { /* ... */ }

func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups

func (x *SortedVersionGroups) GroupIDs() []string                           // sorted group ID list
func (x *SortedVersionGroups) Get(groupID string) *VersionGroup             // nil if not found
func (x *SortedVersionGroups) At(index int) *VersionGroup                   // nil if out of bounds
func (x *SortedVersionGroups) Contains(groupID string) bool                 // group existence check
func (x *SortedVersionGroups) Len() int                                     // number of groups
func (x *SortedVersionGroups) Versions() []*Version                         // all versions across all groups, sorted
func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version
```

### CLI Commands

```bash
# Group versions and display structure
versions group <version1> <version2> ... <versionN>
versions group --id <groupID> <version1> <version2> ...

# List all group IDs
versions group-ids <version1> <version2> ... <versionN>

# Get the group ID for a single version
versions group-id <version>

# Group-level queries (all require --group-id)
versions group-latest --group-id <id> <version1> <version2> ...
versions group-oldest --group-id <id> <version1> <version2> ...
versions group-stable --group-id <id> <version1> <version2> ...
versions group-prerelease --group-id <id> <version1> <version2> ...
versions group-latest-stable --group-id <id> <version1> <version2> ...
versions group-latest-prerelease --group-id <id> <version1> <version2> ...
versions group-contains --group-id <id> --version <v> <version1> <version2> ...
```

**Examples:**
```bash
versions group-id 1.2.3                              # Output: 1.2
versions group-latest --group-id "1.0" 1.0.0 1.0.1 1.0.2
versions group-contains --group-id "1.0" --version "1.0.1" 1.0.0 1.0.1 1.0.2
```

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_group` | `versions: string[]`, `group_id?: string`, `operation?: string`, `target_version?: string` | group structure / query result |

**Operations:** `"list"` (default), `"latest"`, `"oldest"`, `"stable"`, `"prerelease"`, `"latest_stable"`, `"latest_prerelease"`, `"contains"`

## Cross-References

- [[version-sorting]] — for sorting groups with SortVersionGroupMap / SortVersionGroupSlice
- [[version-comparison]] — for the CompareTo logic behind GetLatest / GetOldest
- [[version-range-query]] — for QueryRangeVersions within a group
- [[version-check]] — for filtering with IsStable / IsPrerelease predicates

## Important Notes

- **SDK Group() uses BuildGroupID() as the key** — this is the full version number string (e.g., `"1.2.3"`), NOT just the major version. Check your SDK version for exact key format.
- **VersionMap is keyed by Raw string** — duplicate raw strings overwrite each other in the map
- **Versions() returns unordered results** within a group; use SortVersions() for ordered output
- **GetLatest / GetOldest return nil if the group is empty** — always check for nil
- **LatestStable / LatestPrerelease return nil if no matching versions exist**
- **Remove returns false if the version was not in the group**
- **SortedVersionGroups pre-sorts on construction** — efficient for repeated queries
- **CLI --group-id targets a specific group** across all group subcommands
- **MCP operation parameter** determines what information to return
