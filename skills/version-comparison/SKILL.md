---
name: version-comparison
description: Use when comparing version numbers, checking which version is newer/older, or determining version ordering. Provides expert guidance on using the Go versions SDK for version comparison.
argument-hint: <version-comparison-task>
---

# Version Comparison Skill

## When to Use

- User needs to compare two version numbers (which is newer/older)
- User needs to sort versions by their numeric value
- User needs to implement update checking or dependency version logic
- User is working with version comparison in Go

## API Reference

### Version.CompareTo

**func (x *Version) CompareTo(target *Version) int**

Comparison order:
1. VersionNumbers (digit-by-digit, left to right; missing positions treated as 0)
2. PublicTime (earlier = smaller)
3. Suffix (alphabetic comparison)
4. Raw string (alphabetic fallback)

Returns: -1 (x < target), 0 (equal), 1 (x > target)

### VersionNumbers.CompareTo

**func (x VersionNumbers) CompareTo(target []int) int**

Compares number arrays element-by-element. Shorter array is treated as padded with zeros.

Returns: negative (x < target), 0 (equal), positive (x > target)

### VersionSuffix.CompareTo

**func (x VersionSuffix) CompareTo(target VersionSuffix) int**

Compares suffix strings alphabetically. Empty suffix > any non-empty suffix (release > pre-release).

### VersionPrefix

VersionPrefix is a string type with an IsEmpty() method. Prefix comparison is alphabetical and case-sensitive.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // Basic comparison
    v1 := versions.NewVersion("1.2.3")
    v2 := versions.NewVersion("1.3.0")
    if v1.CompareTo(v2) < 0 {
        fmt.Println("1.2.3 is older than 1.3.0")
    }

    // Pre-release vs release
    beta := versions.NewVersion("1.0.0-beta")
    release := versions.NewVersion("1.0.0")
    if beta.CompareTo(release) < 0 {
        fmt.Println("Pre-release is older than release")
    }

    // Different length versions
    short := versions.NewVersion("1.0")
    full := versions.NewVersion("1.0.0")
    fmt.Println(short.CompareTo(full))  // 0 (treated as equal)

    // Number parts comparison
    nums1 := versions.NewVersionNumbers([]int{1, 2, 0})
    nums2 := versions.NewVersionNumbers([]int{1, 2, 3})
    if nums1.CompareTo(nums2) < 0 {
        fmt.Println("1.2.0 < 1.2.3")
    }
}
```

## Important Notes

- CompareTo treats "1.0" and "1.0.0" as equal (missing digits = 0)
- Prefix comparison is case-sensitive: "v1.0" != "V1.0"
- Empty suffix (release version) is greater than any non-empty suffix (pre-release)
- PublicTime is only compared when both versions have it set (non-zero)
- The comparison implements the compare_anything.Comparable interface
