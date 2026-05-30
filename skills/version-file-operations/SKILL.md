---
name: version-file-operations
description: Use when reading version numbers from files, processing version lists stored in text files. Provides expert guidance on using the Go versions SDK for file-based version operations.
argument-hint: <file-path-or-task>
---

# Version File Operations Skill

## When to Use

- User needs to read version numbers from a text file
- User needs to process a list of versions stored line-by-line in a file
- User needs to parse version lists from dependency lock files or release logs
- User is building CI/CD tools that read version files

## API Reference

### ReadVersionsFromFile

**func ReadVersionsFromFile(filepath string) ([]*Version, error)**

Reads a file line-by-line, parses each line as a Version object. Ignores empty lines and lines starting with #.

### ReadVersionsStringFromFile

**func ReadVersionsStringFromFile(filepath string) ([]string, error)**

Reads a file line-by-line, returns raw strings. Does NOT parse into Version objects. Ignores empty lines and lines starting with #.

## Code Examples

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/versions"
)

func main() {
    // Read and parse versions from file
    versionList, err := versions.ReadVersionsFromFile("versions.txt")
    if err != nil {
        log.Fatalf("Failed to read: %v", err)
    }
    fmt.Printf("Read %d versions\n", len(versionList))

    // Read raw version strings
    rawStrings, err := versions.ReadVersionsStringFromFile("versions.txt")
    if err != nil {
        log.Fatalf("Failed to read: %v", err)
    }
    for _, s := range rawStrings {
        fmt.Println(s)
    }

    // Filter valid versions after reading
    validVersions := make([]*versions.Version, 0)
    for _, v := range versionList {
        if v.IsValid() {
            validVersions = append(validVersions, v)
        }
    }
    fmt.Printf("Valid versions: %d\n", len(validVersions))
}
```

## File Format

```
# This is a comment — lines starting with # are ignored
1.0.0
1.0.1

# Blank lines are also ignored
1.1.0-beta
1.1.0
```

## Important Notes

- File format: one version per line, # for comments, blank lines ignored
- Leading/trailing whitespace on each line is trimmed
- ReadVersionsFromFile uses NewVersion (not NewVersionE), so invalid lines become Version objects with IsValid() == false
- To handle invalid versions, filter with IsValid() after reading
- File reading uses os.ReadFile — for very large files, consider streaming
