---
name: version-visualization
description: Use when visualizing version hierarchies, displaying version trees, or showing version group structures in text format. Covers SDK, CLI, and MCP access paths for version visualization.
argument-hint: <version-visualization-task>
---

# Version Visualization Skill

## When to Use

- User needs to display version structure as a text tree
- User needs to visualize how versions are grouped and organized
- User needs to generate a human-readable overview of a version collection
- User needs a summary view showing only group-level information
- User is debugging version management logic or presenting version info in CLI

## Quick Start

### SDK (Go)

```go
versionList := versions.NewVersions("1.0.0", "1.0.1", "1.1.0", "2.0.0")
// Detailed tree view (max 5 items per group)
versions.VisualizeVersions(versionList, os.Stdout, 5)

// Summary view (groups only)
versions.VisualizeVersionGroups(versionList, os.Stdout)
```

### CLI

```bash
# Visualize versions from arguments
versions visualize 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1

# Visualize with truncation
versions visualize --max-items 3 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1

# Show only group summary
versions visualize --groups 1.0.0 1.0.1 1.1.0 2.0.0

# Read from file
versions visualize --from-file versions.txt
```

### MCP

```json
{
  "tool": "version_visualize",
  "arguments": {
    "versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.0.1"],
    "max_items_per_group": 5
  }
}
```

## API Reference -- SDK

### VisualizeVersions

**func VisualizeVersions(versions []*Version, w io.Writer, maxItems int)**

Renders a detailed tree view of versions grouped by major version. Shows individual versions with optional release times. The `maxItems` parameter controls truncation per group: 0 means show all versions.

Output structure:
```
版本总数: 8
版本组数: 2

┌─ 版本组: 1 (3个版本)
├── 1.0.0 (发布时间: 2024-01-01)
├── 1.0.1 (发布时间: 2024-02-01)
└── ...还有1个版本未显示

┌─ 版本组: 2 (3个版本)
├── 2.0.0
├── 2.0.1
└── 2.1.0
```

### VisualizeVersionGroups

**func VisualizeVersionGroups(versions []*Version, w io.Writer)**

Renders a summary tree showing only group-level information (group ID and version count). Useful for large collections where showing every version would be impractical.

Output structure:
```
版本总数: 8
版本组数: 5

├─ 1 (3个版本组, 共3个版本)
│  ├─ 1.0 (1个版本)
│  └─ 1.1 (2个版本)
└─ 2 (3个版本组, 共3个版本)
```

## CLI Commands

### `versions visualize`

Render a text tree visualization of versions, showing their grouping and hierarchy.

```bash
versions visualize <version1> <version2> ... <versionN>
```

**Flags:**
- `--max-items <n>` -- Maximum number of versions to display per group (0 = show all, default: 0)
- `--groups` -- Show only group-level summary (no individual versions)
- `--from-file <path>` -- Read versions from a file instead of arguments

**Examples:**
```bash
# Full visualization with all versions
versions visualize 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1 2.1.0 3.0.0-alpha 3.0.0-beta

# Limit to 2 versions per group
versions visualize --max-items 2 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1

# Show only group summary (no individual versions)
versions visualize --groups 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1

# Visualize versions from a file
versions visualize --from-file versions.txt

# Combine options: group summary from file
versions visualize --groups --from-file releases.txt
```

## MCP Tools

### `version_visualize`

Generate a text tree visualization of versions, showing their grouping and hierarchy.

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `versions` | array of strings | Yes | List of version strings to visualize |
| `max_items_per_group` | integer | No | Max versions per group (0 = show all, default: 0) |
| `groups_only` | boolean | No | Show only group-level summary, no individual versions (default: false) |

**Example request:**
```json
{
  "tool": "version_visualize",
  "arguments": {
    "versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.0.1", "3.0.0-alpha", "3.0.0-beta"],
    "max_items_per_group": 2
  }
}
```

**Example response:**
```json
{
  "visualization": "版本总数: 7\n版本组数: 3\n\n┌─ 版本组: 1 (3个版本)\n├── 1.0.0\n├── 1.0.1\n└── ...还有1个版本未显示\n\n┌─ 版本组: 2 (2个版本)\n├── 2.0.0\n├── 2.0.1\n\n┌─ 版本组: 3 (2个版本)\n├── 3.0.0-alpha\n├── 3.0.0-beta",
  "total_versions": 7,
  "total_groups": 3
}
```

**Example: Groups only**
```json
{
  "tool": "version_visualize",
  "arguments": {
    "versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.0.1"],
    "groups_only": true
  }
}
```

## Code Examples (SDK)

### Detailed Visualization with Truncation

```go
package main

import (
    "os"
    "time"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions(
        "1.0.0", "1.0.1", "1.1.0",
        "2.0.0", "2.0.1", "2.1.0",
        "3.0.0-alpha", "3.0.0-beta",
    )

    // Set release times (optional)
    for i, v := range versionList {
        v.PublicTime = time.Now().AddDate(0, -i, 0)
    }

    // Detailed visualization -- max 2 per group
    versions.VisualizeVersions(versionList, os.Stdout, 2)
    // Output:
    // 版本总数: 8
    // 版本组数: 3
    //
    // ┌─ 版本组: 1 (3个版本)
    // ├── 1.0.0 (发布时间: 2024-01-01)
    // ├── 1.0.1 (发布时间: 2024-02-01)
    // └── ...还有1个版本未显示
    //
    // ┌─ 版本组: 2 (3个版本)
    // ...
}
```

### Summary Visualization (Groups Only)

```go
package main

import (
    "os"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions(
        "1.0.0", "1.0.1", "1.1.0",
        "2.0.0", "2.0.1",
    )

    // Summary visualization -- shows group hierarchy only
    versions.VisualizeVersionGroups(versionList, os.Stdout)
    // Output:
    // 版本总数: 5
    // 版本组数: 3
    //
    // ├─ 1 (2个版本组, 共3个版本)
    // │  ├─ 1.0 (2个版本)
    // │  └─ 1.1 (1个版本)
    // └─ 2 (1个版本组, 共2个版本)
    //    └─ 2.0 (2个版本)
}
```

### Write Visualization to a Buffer

```go
package main

import (
    "bytes"
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")

    // Write to a buffer instead of stdout
    var buf bytes.Buffer
    versions.VisualizeVersions(versionList, &buf, 0)

    // Use the output as a string
    output := buf.String()
    fmt.Println(output)
}
```

### Show All Versions (No Truncation)

```go
package main

import (
    "os"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("1.0.0", "1.0.1", "1.0.2")

    // maxItems = 0 means show all versions (no truncation)
    versions.VisualizeVersions(versionList, os.Stdout, 0)
}
```

## Important Notes

- **All paths**: Output uses Chinese labels ("版本总数" = total versions, "版本组" = version group, "发布时间" = release time)
- **All paths**: Uses Unicode box-drawing characters for tree structure (├──, └──, ┌─)
- **SDK**: Release time is only shown when PublicTime is non-zero (set via `v.PublicTime = ...`)
- **SDK**: maxItems = 0 means show all versions (no truncation)
- **SDK**: VisualizeVersionGroups shows group hierarchy, not individual versions
- **SDK**: Both functions accept any io.Writer -- os.Stdout, bytes.Buffer, or os.File
- **CLI**: `--max-items 0` shows all versions; any positive integer truncates per group
- **CLI**: `--groups` flag switches to the summary view (equivalent to VisualizeVersionGroups)
- **CLI**: `--from-file` reads one version per line from the specified file
- **MCP**: `max_items_per_group` = 0 shows all versions; positive values truncate per group
- **MCP**: `groups_only` = true switches to group-level summary view