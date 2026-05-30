---
name: version-parsing
description: Use when parsing, validating, or extracting components from version strings like "1.2.3", "v1.2.3-beta1". Provides expert guidance on using the Go versions SDK for version parsing.
argument-hint: <version-string-or-task>
---

# Version Parsing Skill

## When to Use

- User needs to parse a version string into structured components
- User needs to validate if a string is a valid version number
- User needs to extract prefix, numbers, or suffix from a version string
- User is working with semver, pre-release, or custom version formats in Go

## API Reference

### Core Functions

**NewVersion(versionStr string) *Version**
Creates a Version object from a string. Never panics — returns an object even for invalid strings (check with IsValid()).

**NewVersionE(versionStr string) (*Version, error)**
Creates a Version with error return. Returns ErrVersionInvalid for invalid strings.

**NewVersions(versionStringSlice ...string) []*Version**
Batch-create multiple Version objects.

### Version Struct

```go
type Version struct {
    Raw            string         // Original string, e.g. "v1.2.3-beta1"
    PublicTime     time.Time      // Release time
    VersionNumbers VersionNumbers // Number part, e.g. [1,2,3]
    Prefix         VersionPrefix  // Prefix, e.g. "v"
    Suffix         VersionSuffix  // Suffix, e.g. "-beta1"
}
```

### Validation

**func (v *Version) IsValid() bool**
Returns true if the version contains at least one number in VersionNumbers.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions"
)

func main() {
    // Parse a version string
    v := versions.NewVersion("v1.2.3-rc1")
    fmt.Printf("Prefix: %s\n", v.Prefix)           // "v"
    fmt.Printf("Numbers: %v\n", v.VersionNumbers)   // [1 2 3]
    fmt.Printf("Suffix: %s\n", v.Suffix)            // "-rc1"

    // Validate a version
    valid := versions.NewVersion("1.2.3")
    invalid := versions.NewVersion("not-a-version")
    fmt.Println(valid.IsValid())    // true
    fmt.Println(invalid.IsValid())  // false

    // Error-checking parse
    v, err := versions.NewVersionE("")
    if err != nil {
        fmt.Println("Error:", err)  // "version invalid"
    }
}
```

## Important Notes

- NewVersion never returns nil — always check IsValid() for invalid inputs
- Supports arbitrary number of segments: "1.2.3.4.5" works
- Leading zeros are stripped: "1.02.003" becomes [1,2,3]
- Pure alphabetic strings like "abc" return invalid Version (empty VersionNumbers)
- VersionPrefix is a string type — call IsEmpty() to check for no prefix
- VersionSuffix is a string type — call IsEmpty() to check for no suffix
