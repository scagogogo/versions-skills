---
name: version-sorting
description: Use when sorting a list of version numbers in natural order. Covers SDK, CLI, and MCP access paths for version sorting operations.
argument-hint: <version-sorting-task>
---

# Version Sorting Skill

## When to Use

- User needs to sort a list of version strings or Version objects
- User needs to find the latest or oldest version in a collection
- User needs to display versions in ascending or descending order
- User is implementing version selection UI or dependency resolution
- User wants to sort versions read from a file

## Installation

### SDK (Go library)

```bash
go get github.com/scagogogo/versions-skills
```

### CLI binary

**Option A: Download from GitHub Releases (Recommended)**

Pre-built binaries for Linux, macOS, Windows, FreeBSD, OpenBSD, and NetBSD on amd64, arm64, arm, 386, mips, mips64, mips64le, ppc64, ppc64le, s390x, and riscv64. Linux packages: deb, rpm, apk.

```bash
# Linux (amd64)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_linux_amd64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# macOS (arm64 / Apple Silicon)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_darwin_arm64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# Or install via package manager (Linux only):
# Debian/Ubuntu: dpkg -i versions_{VERSION}_linux_amd64.deb
# RHEL/Fedora:   rpm -i versions_{VERSION}_linux_amd64.rpm
# Alpine:        apk add versions_{VERSION}_linux_amd64.apk
```

> Replace `{VERSION}` with the latest release tag (e.g. `0.2.0`). See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all available platforms and the current version.

**Option B: Install via Go**

```bash
go install github.com/scagogogo/versions-skills/cmd/versions@latest
```

### MCP server

**Option A: Download from GitHub Releases (Recommended)**

```bash
# Linux (amd64)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions-mcp_{VERSION}_linux_amd64.tar.gz | tar xz
chmod +x versions-mcp && sudo mv versions-mcp /usr/local/bin/

# macOS (arm64 / Apple Silicon)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions-mcp_{VERSION}_darwin_arm64.tar.gz | tar xz
chmod +x versions-mcp && sudo mv versions-mcp /usr/local/bin/
```

> Replace `{VERSION}` with the latest release tag. See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all platforms.

**Option B: Install via Go**

```bash
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

## Quick Start

### SDK (Go)

```go
sorted := versions.SortVersionStringSlice([]string{"2.0.0", "1.0.0", "1.10.0", "1.2.0"})
// Result: ["1.0.0", "1.2.0", "1.10.0", "2.0.0"]
```

### CLI

```bash
# Sort version strings (ascending)
versions sort 1.0.0 1.10.0 1.2.0 2.0.0

# Sort in descending order
versions sort --desc 1.0.0 1.10.0 1.2.0 2.0.0

# Sort from a file
versions sort --from-file versions.txt
```

### MCP

```json
{
  "tool": "version_sort",
  "arguments": {
    "versions": ["2.0.0", "1.0.0", "1.10.0", "1.2.0"],
    "descending": false
  }
}
```

## API Reference -- SDK

### SortVersionStringSlice

**func SortVersionStringSlice(versionStringSlice []string) []string**

Sorts a string slice of version numbers. Parses each string, sorts by Version.CompareTo rules, returns sorted strings. Does not modify the original slice.

### SortVersionSlice

**func SortVersionSlice(versions []*Version) []*Version**

Sorts Version object slice. Uses group-based algorithm: groups by major version, sorts groups, sorts within each group, merges results. Does not modify the original slice.

### VersionSlice

**type VersionSlice []*Version**

An ordered collection type that implements sort.Interface. Allows direct use of `sort.Sort()` on version slices without closures.

Methods:
- **Len() int** -- returns the number of versions
- **Less(i, j int) bool** -- returns true if version at index i is older than version at j
- **Swap(i, j int)** -- swaps versions at indices i and j

```go
slice := versions.VersionSlice(versions.NewVersions("2.0.0", "1.0.0", "1.5.0"))
sort.Sort(slice)  // slice is now sorted ascending
```

### SortVersionGroupMap

**func SortVersionGroupMap(versionGroupMap map[string]*VersionGroup) []*VersionGroup**

Converts a version group map to a sorted slice of VersionGroup objects.

### SortVersionGroupSlice

**func SortVersionGroupSlice(groupSlice []*VersionGroup)**

In-place sort of a VersionGroup slice. Modifies the input slice directly.

## CLI Commands

### `versions sort`

Sort version strings in natural (semantic) order.

```bash
versions sort <version1> <version2> ... <versionN>
```

**Flags:**
- `--desc` -- Sort in descending order (latest first)
- `--from-file <path>` -- Read versions from a file instead of arguments (one version per line)

**Examples:**
```bash
# Basic ascending sort
versions sort 1.0.0 1.10.0 1.2.0 2.0.0
# Output: 1.0.0 1.2.0 1.10.0 2.0.0

# Descending sort (latest first)
versions sort --desc 1.0.0 1.10.0 1.2.0 2.0.0
# Output: 2.0.0 1.10.0 1.2.0 1.0.0

# Sort versions from a file
versions sort --from-file releases.txt
```

### `versions sort-strings`

Sort raw version strings (same as `versions sort` but explicitly for string input).

```bash
versions sort-strings <version1> <version2> ... <versionN>
```

**Flags:**
- `--desc` -- Sort in descending order
- `--from-file <path>` -- Read versions from a file

## MCP Tools

### `version_sort`

Sort a list of version strings in semantic order.

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `versions` | array of strings | Yes | List of version strings to sort |
| `descending` | boolean | No | Sort in descending order (default: false) |

**Example request:**
```json
{
  "tool": "version_sort",
  "arguments": {
    "versions": ["2.0.0", "1.0.0", "1.10.0", "1.2.0"],
    "descending": true
  }
}
```

**Example response:**
```json
{
  "sorted_versions": ["2.0.0", "1.10.0", "1.2.0", "1.0.0"]
}
```

## Code Examples (SDK)

### Sort Version Strings

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    unsorted := []string{"2.0.0", "1.0.0", "1.10.0", "1.2.0", "v1.5.0"}
    sorted := versions.SortVersionStringSlice(unsorted)
    for _, v := range sorted {
        fmt.Println(v)
    }
    // Output: 1.0.0, 1.2.0, v1.5.0, 1.10.0, 2.0.0
}
```

### Sort Version Objects and Find Latest

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("2.0.0", "1.0.0", "1.10.0")
    sortedVersions := versions.SortVersionSlice(versionList)

    // Find latest (last element after ascending sort)
    latest := sortedVersions[len(sortedVersions)-1]
    fmt.Printf("Latest: %s\n", latest.Raw)

    // Find oldest (first element)
    oldest := sortedVersions[0]
    fmt.Printf("Oldest: %s\n", oldest.Raw)
}
```

### Use VersionSlice with sort.Sort

```go
package main

import (
    "fmt"
    "sort"
    "github.com/scagogogo/versions-skills"
)

func main() {
    slice := versions.VersionSlice(versions.NewVersions("3.0.0", "1.0.0", "2.0.0"))

    // Direct sort using standard library -- no closures needed
    sort.Sort(slice)

    for _, v := range slice {
        fmt.Println(v.Raw)
    }
    // Output: 1.0.0, 2.0.0, 3.0.0
}
```

### Sort Version Groups

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
    groupMap := versions.Group(versionList)
    sortedGroups := versions.SortVersionGroupMap(groupMap)
    for _, g := range sortedGroups {
        fmt.Printf("Group %s: %d versions\n", g.ID(), g.Count())
    }
}
```

## Important Notes

- **SDK**: SortVersionStringSlice creates new Version objects internally -- O(n) extra space
- **SDK**: SortVersionSlice uses a group-based algorithm for stable, semantically correct ordering
- **SDK**: Neither SortVersionStringSlice nor SortVersionSlice modifies the original input slice
- **SDK**: SortVersionGroupSlice modifies the input slice in-place (unlike the other sort functions)
- **SDK**: VersionSlice implements sort.Interface -- use `sort.Sort()` directly, no closures needed
- **All paths**: "1.10.0" correctly sorts after "1.2.0" (numeric, not alphabetic sorting)
- **All paths**: Pre-release versions sort before their release counterparts (e.g., "1.0.0-beta" < "1.0.0")
- **CLI**: The `--from-file` flag reads one version per line, ignores blank lines and `#` comments
- **MCP**: The `descending` parameter defaults to false (ascending order)