---
name: version-range-query
description: Use when querying versions within a specific range, implementing version constraints, or checking version compatibility. Covers SDK, CLI, and MCP access paths for range queries.
argument-hint: <version-range-query-task>
---

# Version Range Query Skill

## When to Use

- User needs to find all versions between a start and end version
- User needs to implement version constraint logic (e.g., ">=1.0.0 and <2.0.0")
- User needs to check if a version falls within a specific range
- User needs to check if a version is between two other versions
- User is implementing dependency version resolution or compatibility checks

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
// Simple between check
v := versions.NewVersion("1.5.0")
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")
fmt.Println(v.IsBetween(low, high))  // true

// Range query across groups
sortedGroups := versions.NewSortedVersionGroups(versionList)
start := tuple.New2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
end := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
result := sortedGroups.QueryRange(start, end)
```

### CLI

```bash
# Query versions in a range
versions range 1.0.0 2.0.0 --include-start --include-end

# Check if a version is between two others
versions check 1.5.0 --between-low 1.0.0 --between-high 2.0.0
```

### MCP

```json
{
  "tool": "version_range_query",
  "arguments": {
    "versions": ["1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0"],
    "start": "1.0.0",
    "end": "2.0.0",
    "include_start": true,
    "include_end": false
  }
}
```

## API Reference -- SDK

### ContainsPolicy

```go
type ContainsPolicy int

const (
    ContainsPolicyNone ContainsPolicy = iota  // Unspecified (defaults to inclusive)
    ContainsPolicyYes                          // Include boundary [ ]
    ContainsPolicyNo                           // Exclude boundary ( )
)
```

Used to control whether range boundaries are included or excluded. Mathematically: ContainsPolicyYes = `[` or `]`, ContainsPolicyNo = `(` or `)`. ContainsPolicyNone defaults to ContainsPolicyYes (inclusive) in most contexts.

### Version.IsBetween

**func (x *Version) IsBetween(low, high *Version) bool**

Checks if the current version is between two boundary versions (inclusive on both ends). Returns true if `low <= x <= high`. If low or high is nil, the corresponding boundary check is skipped.

```go
v := versions.NewVersion("1.5.0")
fmt.Println(v.IsBetween(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0")))  // true
fmt.Println(v.IsBetween(nil, versions.NewVersion("2.0.0")))  // true (no lower bound)
fmt.Println(v.IsBetween(versions.NewVersion("3.0.0"), nil))  // false (1.5.0 < 3.0.0)
```

### VersionGroup.QueryRangeVersions

**func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**

Queries versions within a range inside a single group. Returns sorted results. Import `github.com/golang-infrastructure/go-tuple` for tuple creation.

### SortedVersionGroups.QueryRange

**func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version**

Cross-group range query. Searches across all relevant version groups. Returns sorted results. Returns nil if the start version's group does not exist in the collection.

## CLI Commands

### `versions range`

Query all versions that fall within a specified range.

```bash
versions range <start> <end> [version1 version2 ...]
```

**Flags:**
- `--include-start` -- Include the start boundary version (default: true)
- `--include-end` -- Include the end boundary version (default: false)
- `--from-file <path>` -- Read versions from a file instead of arguments

**Examples:**
```bash
# Find versions in [1.0.0, 2.0.0)
versions range 1.0.0 2.0.0 1.0.0 1.1.0 1.5.0 2.0.0 2.1.0 3.0.0
# Output: 1.0.0 1.1.0 1.5.0

# Find versions in [1.0.0, 2.0.0] (inclusive on both ends)
versions range 1.0.0 2.0.0 --include-end 1.0.0 1.1.0 1.5.0 2.0.0 2.1.0
# Output: 1.0.0 1.1.0 1.5.0 2.0.0

# Use version file as input
versions range 1.0.0 2.0.0 --from-file versions.txt
```

### `versions check`

Check version properties, including between-ness.

```bash
versions check <version> --between-low <low> --between-high <high>
```

**Flags:**
- `--between-low <version>` -- Lower bound for IsBetween check
- `--between-high <version>` -- Upper bound for IsBetween check

**Examples:**
```bash
# Check if 1.5.0 is between 1.0.0 and 2.0.0
versions check 1.5.0 --between-low 1.0.0 --between-high 2.0.0
# Output: true

# Check if 3.0.0 is between 1.0.0 and 2.0.0
versions check 3.0.0 --between-low 1.0.0 --between-high 2.0.0
# Output: false
```

## MCP Tools

### `version_range_query`

Query versions within a specified range or check if a version falls between two boundaries.

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `versions` | array of strings | Yes | List of version strings to search |
| `start` | string | No | Start boundary for range query |
| `end` | string | No | End boundary for range query |
| `include_start` | boolean | No | Include start boundary (default: true) |
| `include_end` | boolean | No | Include end boundary (default: false) |
| `check_version` | string | No | Version to check for between-ness |
| `low` | string | No | Lower bound for between-ness check |
| `high` | string | No | Upper bound for between-ness check |

**Example: Range query**
```json
{
  "tool": "version_range_query",
  "arguments": {
    "versions": ["1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0"],
    "start": "1.0.0",
    "end": "2.0.0",
    "include_start": true,
    "include_end": false
  }
}
```

**Example: Between-ness check**
```json
{
  "tool": "version_range_query",
  "arguments": {
    "check_version": "1.5.0",
    "low": "1.0.0",
    "high": "2.0.0"
  }
}
```

## Code Examples (SDK)

### Simple Between Check with IsBetween

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    v := versions.NewVersion("1.5.0")
    low := versions.NewVersion("1.0.0")
    high := versions.NewVersion("2.0.0")

    fmt.Println(v.IsBetween(low, high))  // true (1.0.0 <= 1.5.0 <= 2.0.0)

    // Boundary values are inclusive
    fmt.Println(low.IsBetween(low, high))  // true
    fmt.Println(high.IsBetween(low, high))  // true

    // Nil boundaries are ignored
    fmt.Println(v.IsBetween(nil, high))  // true (no lower bound check)
    fmt.Println(v.IsBetween(low, nil))   // true (1.5.0 >= 1.0.0, no upper bound)
}
```

### Range Query with ContainsPolicy

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
    "github.com/golang-infrastructure/go-tuple"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0")
    sortedGroups := versions.NewSortedVersionGroups(versionList)

    // Query [1.0.0, 2.0.0) -- include 1.0.0, exclude 2.0.0
    start := tuple.New2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
    end := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
    result := sortedGroups.QueryRange(start, end)
    fmt.Printf("Versions in [1.0.0, 2.0.0): %d\n", len(result))
    for _, v := range result {
        fmt.Println(v.Raw)
    }
    // Output: 1.0.0, 1.1.0

    // Query [2.0.0, infinity) -- all versions >= 2.0.0
    start2 := tuple.New2(versions.NewVersion("2.0.0"), versions.ContainsPolicyYes)
    end2 := tuple.New2(versions.NewVersion("999.0.0"), versions.ContainsPolicyNo)
    result2 := sortedGroups.QueryRange(start2, end2)
    fmt.Printf("Versions >= 2.0.0: %d\n", len(result2))
}
```

### Single Group Range Query

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
    "github.com/golang-infrastructure/go-tuple"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.0.1", "1.0.2", "1.0.3", "2.0.0")
    groupMap := versions.Group(versionList)

    // Range query within the "1.0" group
    if group, ok := groupMap["1.0"]; ok {
        gStart := tuple.New2(versions.NewVersion("1.0.1"), versions.ContainsPolicyYes)
        gEnd := tuple.New2(versions.NewVersion("1.0.3"), versions.ContainsPolicyNo)
        rangeResult := group.QueryRangeVersions(gStart, gEnd)
        fmt.Printf("Versions in [1.0.1, 1.0.3): %v\n", rangeResult)
        // Output: [1.0.1, 1.0.2]
    }
}
```

## Important Notes

- **SDK**: tuple.New2 creates a Tuple2 -- import `github.com/golang-infrastructure/go-tuple`
- **SDK**: ContainsPolicyNone defaults to ContainsPolicyYes (inclusive) in most contexts
- **SDK**: QueryRange on SortedVersionGroups returns nil if the start version's group does not exist
- **SDK**: IsBetween is inclusive on both ends (`low <= x <= high`); nil boundaries are skipped
- **SDK**: For open-ended range queries, use a very large version as the end boundary (e.g., "999.0.0")
- **SDK**: Range queries always return pre-sorted results (ascending order)
- **CLI**: `--include-start` defaults to true, `--include-end` defaults to false, matching the common `[start, end)` convention
- **MCP**: The same tool supports both range queries (with start/end) and between-ness checks (with check_version/low/high)
- **All paths**: Boundary inclusion semantics are consistent: `ContainsPolicyYes` = `[` or `]`, `ContainsPolicyNo` = `(` or `)`