---
name: version-grouping
description: Use when grouping versions by major/minor version number, managing version collections by their version series. Covers SDK, CLI, and MCP access paths for version grouping.
argument-hint: <version-grouping-task>
---

# Version Grouping Skill

## When to Use

- User needs to group versions by their major/minor version number
- User needs to organize a large collection of versions into version series
- User needs to find all versions in a specific major version (e.g., all 1.x versions)
- User needs the latest, oldest, stable, or prerelease version within a group
- User is building a version selector UI with hierarchical version choices

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
versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0")
sortedGroups := versions.NewSortedVersionGroups(versionList)
fmt.Println(sortedGroups.GroupIDs())  // ["1.0", "1.1", "2.0", "2.1", "3.0"]
```

### CLI

```bash
# Group versions and display
versions group 1.0.0 1.1.0 2.0.0 2.1.0 3.0.0

# List all group IDs
versions group-ids 1.0.0 1.1.0 2.0.0 2.1.0 3.0.0

# Get latest version in a group
versions group-latest --group-id "1.0" 1.0.0 1.0.1 1.0.2
```

### MCP

```json
{
  "tool": "version_group",
  "arguments": {
    "versions": ["1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0"]
  }
}
```

## API Reference -- SDK

### Group

**func Group(versions []*Version) map[string]*VersionGroup**

Groups versions by their VersionNumbers' BuildGroupID (full number sequence). Returns a map of groupID to VersionGroup.

### VersionGroup Type

```go
type VersionGroup struct {
    GroupVersionNumbers VersionNumbers
    VersionMap          map[string]*Version
}
```

**Core Methods:**
- **NewVersionGroup(groupVersionNumbers VersionNumbers) *VersionGroup** -- create a new group
- **NewVersionGroupFromVersions(versions []*Version) *VersionGroup** -- create from version array
- **func (x *VersionGroup) Add(v *Version) bool** -- add version to group; returns true if version already existed
- **func (x *VersionGroup) Contains(v *Version) bool** -- check if version exists in group
- **func (x *VersionGroup) ID() string** -- get group ID (e.g., "1.2")
- **func (x *VersionGroup) Versions() []*Version** -- get all versions (unordered)
- **func (x *VersionGroup) SortVersions() []*Version** -- get sorted versions (ascending)
- **func (x *VersionGroup) CompareTo(target *VersionGroup) int** -- compare two groups

**Convenience Methods:**
- **func (x *VersionGroup) GetLatest() *Version** -- return the newest version in the group; nil if empty
- **func (x *VersionGroup) GetOldest() *Version** -- return the oldest version in the group; nil if empty
- **func (x *VersionGroup) Count() int** -- return the number of versions in the group
- **func (x *VersionGroup) StableVersions() []*Version** -- return all stable (no suffix) versions
- **func (x *VersionGroup) PrereleaseVersions() []*Version** -- return all prerelease versions
- **func (x *VersionGroup) LatestStable() *Version** -- return the newest stable version; nil if none
- **func (x *VersionGroup) LatestPrerelease() *Version** -- return the newest prerelease version; nil if none
- **func (x *VersionGroup) Remove(v *Version) bool** -- remove a version; returns true if it existed
- **func (x *VersionGroup) Filter(predicate func(*Version) bool) []*Version** -- filter versions with a predicate

**Range Query:**
- **func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version** -- range query within group; returns sorted results

### SortedVersionGroups Type

```go
type SortedVersionGroups struct { ... }
```

**Methods:**
- **func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups** -- create sorted groups from versions
- **func (x *SortedVersionGroups) GroupIDs() []string** -- get sorted group ID list
- **func (x *SortedVersionGroups) Get(groupID string) *VersionGroup** -- get group by ID; nil if not found
- **func (x *SortedVersionGroups) At(index int) *VersionGroup** -- get group by index; nil if out of bounds
- **func (x *SortedVersionGroups) Contains(groupID string) bool** -- check if a group ID exists
- **func (x *SortedVersionGroups) Len() int** -- return number of groups
- **func (x *SortedVersionGroups) Versions() []*Version** -- return all versions across all groups, sorted
- **func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version** -- cross-group range query

## CLI Commands

### `versions group`

Group versions and display their structure.

```bash
versions group <version1> <version2> ... <versionN>
```

**Flags:**
- `--id <groupID>` -- Show only the specified group

**Example:**
```bash
versions group 1.0.0 1.0.1 1.1.0 2.0.0 2.1.0
versions group --id "1.0" 1.0.0 1.0.1 1.1.0 2.0.0
```

### `versions group-ids`

List all version group IDs.

```bash
versions group-ids <version1> <version2> ... <versionN>
```

### `versions group-latest`

Get the latest version in a specific group.

```bash
versions group-latest --group-id <id> <version1> <version2> ...
```

### `versions group-oldest`

Get the oldest version in a specific group.

```bash
versions group-oldest --group-id <id> <version1> <version2> ...
```

### `versions group-stable`

Get all stable versions in a specific group.

```bash
versions group-stable --group-id <id> <version1> <version2> ...
```

### `versions group-prerelease`

Get all prerelease versions in a specific group.

```bash
versions group-prerelease --group-id <id> <version1> <version2> ...
```

### `versions group-latest-stable`

Get the latest stable version in a specific group.

```bash
versions group-latest-stable --group-id <id> <version1> <version2> ...
```

### `versions group-latest-prerelease`

Get the latest prerelease version in a specific group.

```bash
versions group-latest-prerelease --group-id <id> <version1> <version2> ...
```

### `versions group-contains`

Check if a specific version exists in a group.

```bash
versions group-contains --group-id <id> --version <version> <version1> <version2> ...
```

### `versions group-id`

Get the group ID for a single version.

```bash
versions group-id <version>
```

**Example:**
```bash
versions group-id 1.2.3
# Output: 1.2
```

## MCP Tools

### `version_group`

Group versions by their version number series and query group information.

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `versions` | array of strings | Yes | List of version strings to group |
| `group_id` | string | No | Filter to a specific group ID |
| `operation` | string | No | Operation: "list" (default), "latest", "oldest", "stable", "prerelease", "latest_stable", "latest_prerelease", "contains" |
| `target_version` | string | No | Version to check (for "contains" operation) |

**Example request:**
```json
{
  "tool": "version_group",
  "arguments": {
    "versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.1.0"],
    "group_id": "1.0",
    "operation": "latest"
  }
}
```

**Example response:**
```json
{
  "group_id": "1.0",
  "versions": ["1.0.0", "1.0.1"],
  "latest": "1.0.1"
}
```

## Code Examples (SDK)

### Basic Grouping

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0", "3.0.0")
    groupMap := versions.Group(versionList)
    for groupID, group := range groupMap {
        fmt.Printf("Group %s: %d versions\n", groupID, group.Count())
    }
}
```

### Sorted Groups with Convenience Methods

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions(
        "1.0.0", "1.0.1", "1.0.2-beta", "1.0.2",
        "2.0.0", "2.0.1-rc1", "2.0.1",
    )
    sortedGroups := versions.NewSortedVersionGroups(versionList)

    // List all group IDs
    fmt.Printf("Group IDs: %v\n", sortedGroups.GroupIDs())

    // Get a specific group by ID
    if group := sortedGroups.Get("1.0"); group != nil {
        fmt.Printf("Group 1.0 latest: %s\n", group.GetLatest().Raw)
        fmt.Printf("Group 1.0 oldest: %s\n", group.GetOldest().Raw)
        fmt.Printf("Group 1.0 count: %d\n", group.Count())
    }

    // Stable and prerelease filtering
    if group := sortedGroups.Get("2.0"); group != nil {
        fmt.Printf("Stable: %v\n", group.StableVersions())
        fmt.Printf("Prerelease: %v\n", group.PrereleaseVersions())
        fmt.Printf("Latest stable: %s\n", group.LatestStable().Raw)
        fmt.Printf("Latest prerelease: %s\n", group.LatestPrerelease().Raw)
    }
}
```

### Group Manipulation

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.0.1", "1.0.2")
    group := versions.NewVersionGroupFromVersions(versionList)

    // Add a new version
    group.Add(versions.NewVersion("1.0.3"))

    // Remove a version
    group.Remove(versions.NewVersion("1.0.0"))

    // Check membership
    fmt.Println(group.Contains(versions.NewVersion("1.0.1")))  // true
    fmt.Println(group.Contains(versions.NewVersion("1.0.0")))  // false (removed)

    // Filter with a predicate
    onlyStable := group.Filter(func(v *versions.Version) bool {
        return v.IsStable()
    })
    fmt.Printf("Stable versions: %d\n", len(onlyStable))
}
```

### SortedVersionGroups Navigation

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
    sortedGroups := versions.NewSortedVersionGroups(versionList)

    // Access by index
    fmt.Printf("First group: %s\n", sortedGroups.At(0).ID())
    fmt.Printf("Last group: %s\n", sortedGroups.At(sortedGroups.Len()-1).ID())

    // Check existence
    fmt.Println(sortedGroups.Contains("1.0"))  // true
    fmt.Println(sortedGroups.Contains("9.0"))  // false

    // Get all versions sorted
    allSorted := sortedGroups.Versions()
    for _, v := range allSorted {
        fmt.Println(v.Raw)
    }
}
```

## Important Notes

- **SDK**: Group() uses BuildGroupID() as the key -- this is the full version number string (e.g., "1.2.3"), NOT just the major version
- **SDK**: VersionMap is keyed by Raw string -- duplicate raw strings overwrite
- **SDK**: Versions() returns unordered results; use SortVersions() for ordered
- **SDK**: SortedVersionGroups pre-sorts on construction -- efficient for repeated queries
- **SDK**: GetLatest/GetOldest return nil if the group is empty
- **SDK**: LatestStable/LatestPrerelease return nil if no matching versions exist
- **SDK**: Remove returns false if the version was not in the group
- **CLI**: All group subcommands accept `--group-id` to target a specific group
- **MCP**: The `operation` parameter determines what information to return about the group