---
name: version-grouping
description: Use when grouping versions by major/minor version number, managing version collections by their version series. Provides expert guidance on using the Go versions SDK for version grouping.
argument-hint: <version-grouping-task>
---

# Version Grouping Skill

## When to Use

- User needs to group versions by their major version number
- User needs to organize a large collection of versions into version series
- User needs to find all versions in a specific major version (e.g., all 1.x versions)
- User is building a version selector UI with hierarchical version choices

## API Reference

### Group Function

**func Group(versions []*Version) map[string]*VersionGroup**

Groups versions by their VersionNumbers' BuildGroupID (full number sequence). Returns a map of groupID → VersionGroup.

### VersionGroup Type

```go
type VersionGroup struct {
    GroupVersionNumbers VersionNumbers
    VersionMap          map[string]*Version
}
```

**Key Methods:**

- **NewVersionGroup(groupVersionNumbers VersionNumbers) *VersionGroup** — create a new group
- **NewVersionGroupFromVersions(versions []*Version) *VersionGroup** — create from version array
- **func (x *VersionGroup) Add(v *Version) bool** — add version to group
- **func (x *VersionGroup) Contains(v *Version) bool** — check if version exists
- **func (x *VersionGroup) ID() string** — get group ID (e.g., "1.2")
- **func (x *VersionGroup) Versions() []*Version** — get all versions (unordered)
- **func (x *VersionGroup) SortVersions() []*Version** — get sorted versions
- **func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version** — range query within group

### SortedVersionGroups Type

```go
type SortedVersionGroups struct { ... }
```

**Key Methods:**

- **func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups** — create sorted groups
- **func (x *SortedVersionGroups) GroupIDs() []string** — get sorted group ID list
- **func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version** — cross-group range query

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
    "github.com/golang-infrastructure/go-tuple"
)

func main() {
    // Group versions
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0")
    groupMap := versions.Group(versionList)
    for groupID, group := range groupMap {
        fmt.Printf("Group %s: %d versions\n", groupID, len(group.Versions()))
    }

    // Create sorted groups and list IDs
    sortedGroups := versions.NewSortedVersionGroups(versionList)
    fmt.Printf("Group IDs: %v\n", sortedGroups.GroupIDs())

    // Find latest in a group
    if group, ok := groupMap["1.0"]; ok {
        sorted := group.SortVersions()
        latest := sorted[len(sorted)-1]
        fmt.Printf("Latest in group 1.0: %s\n", latest.Raw)
    }
}
```

## Important Notes

- Group() uses BuildGroupID() as the key — this is the full version number string (e.g., "1.2.3"), NOT just the major version
- VersionMap is keyed by Raw string — duplicates overwrite
- Versions() returns unordered results; use SortVersions() for ordered
- SortedVersionGroups pre-sorts on construction — efficient for repeated queries
