---
name: version-sorting
description: Use when sorting a list of version numbers in natural order. Provides expert guidance on using the Go versions SDK for version sorting.
argument-hint: <version-sorting-task>
---

# Version Sorting Skill

## When to Use

- User needs to sort a list of version strings or Version objects
- User needs to find the latest or oldest version in a collection
- User needs to display versions in ascending/descending order
- User is implementing version selection UI or dependency resolution

## API Reference

### SortVersionStringSlice

**func SortVersionStringSlice(versionStringSlice []string) []string**

Sorts a string slice of version numbers. Parses each string, sorts by Version.CompareTo rules, returns sorted strings.

### SortVersionSlice

**func SortVersionSlice(versions []*Version) []*Version**

Sorts Version object slice. Uses group-based algorithm: groups by major version, sorts groups, sorts within each group, merges results.

### SortVersionGroupMap / SortVersionGroupSlice

**func SortVersionGroupMap(versionGroupMap map[string]*VersionGroup) []*VersionGroup**
**func SortVersionGroupSlice(groupSlice []*VersionGroup)**

Utility functions for sorting version group collections.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // Sort version strings
    unsorted := []string{"2.0.0", "1.0.0", "1.10.0", "1.2.0", "v1.5.0"}
    sorted := versions.SortVersionStringSlice(unsorted)
    for _, v := range sorted {
        fmt.Println(v)
    }
    // Output: 1.0.0, 1.2.0, v1.5.0, 1.10.0, 2.0.0

    // Sort Version objects
    versionList := versions.NewVersions("2.0.0", "1.0.0", "1.10.0")
    sortedVersions := versions.SortVersionSlice(versionList)
    for _, v := range sortedVersions {
        fmt.Println(v.Raw)
    }
    // Output: 1.0.0, 1.10.0, 2.0.0

    // Find latest version
    latest := sortedVersions[len(sortedVersions)-1]
    fmt.Printf("Latest: %s\n", latest.Raw)
}
```

## Important Notes

- SortVersionStringSlice creates new Version objects internally — O(n) extra space
- SortVersionSlice uses a group-based algorithm for stable, semantically correct ordering
- "1.10.0" correctly sorts after "1.2.0" (numeric, not alphabetic)
- Pre-release versions sort before their release counterparts
- Neither function modifies the original input slice
