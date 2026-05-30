---
name: version-range-query
description: Use when querying versions within a specific range, implementing version constraints, or checking version compatibility. Provides expert guidance on using the Go versions SDK for range queries.
argument-hint: <version-range-query-task>
---

# Version Range Query Skill

## When to Use

- User needs to find all versions between a start and end version
- User needs to implement version constraint logic (e.g., ">=1.0.0 and <2.0.0")
- User needs to check if a version falls within a specific range
- User is implementing dependency version resolution or compatibility checks

## API Reference

### ContainsPolicy

```go
type ContainsPolicy int

const (
    ContainsPolicyNone ContainsPolicy = iota  // Unspecified
    ContainsPolicyYes                          // Include boundary
    ContainsPolicyNo                           // Exclude boundary
)
```

Used to control whether range boundaries are included or excluded. Mathematically: ContainsPolicyYes = [ or ], ContainsPolicyNo = ( or ).

### VersionGroup.QueryRangeVersions

**func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**

Queries versions within a range inside a single group. Returns sorted results.

### SortedVersionGroups.QueryRange

**func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**

Cross-group range query. Searches across all relevant version groups. Returns sorted results.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
    "github.com/golang-infrastructure/go-tuple"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0")
    sortedGroups := versions.NewSortedVersionGroups(versionList)

    // Query [1.0.0, 2.0.0) — include 1.0.0, exclude 2.0.0
    start := tuple.New2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
    end := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
    result := sortedGroups.QueryRange(start, end)
    fmt.Printf("Versions in [1.0.0, 2.0.0): %d\n", len(result))
    for _, v := range result {
        fmt.Println(v.Raw)
    }
    // Output: 1.0.0, 1.1.0

    // Query [2.0.0, ∞) — all versions >= 2.0.0
    start2 := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyYes)
    end2 := tuple.New2(versions.NewVersion("999.0.0"), versions.ContainsPolicyNo)
    result2 := sortedGroups.QueryRange(start2, end2)
    fmt.Printf("Versions >= 2.0.0: %d\n", len(result2))

    // Single group range query
    groupMap := versions.Group(versionList)
    for _, group := range groupMap {
        gStart := tuple.New2(versions.NewVersion("0.0.0"), versions.ContainsPolicyYes)
        gEnd := tuple.New2(versions.NewVersion("99.0.0"), versions.ContainsPolicyYes)
        rangeResult := group.QueryRangeVersions(gStart, gEnd)
        fmt.Printf("Group %s: %d versions in range\n", group.ID(), len(rangeResult))
    }
}
```

## Important Notes

- tuple.New2 creates a Tuple2 — import "github.com/golang-infrastructure/go-tuple"
- ContainsPolicyNone defaults to ContainsPolicyYes (inclusive) in most contexts
- QueryRange on SortedVersionGroups returns nil if the start version's group doesn't exist
- For open-ended ranges, use a very large version as the end boundary
- Range queries return pre-sorted results (ascending order)
