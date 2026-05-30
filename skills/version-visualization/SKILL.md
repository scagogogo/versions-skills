---
name: version-visualization
description: Use when visualizing version hierarchies, displaying version trees, or showing version group structures in text format. Provides expert guidance on using the Go versions SDK for version visualization.
argument-hint: <version-visualization-task>
---

# Version Visualization Skill

## When to Use

- User needs to display version structure as a text tree
- User needs to visualize how versions are grouped and organized
- User needs to generate a human-readable overview of a version collection
- User is debugging version management logic or presenting version info in CLI

## API Reference

### VisualizeVersions

**func VisualizeVersions(versions []*Version, w io.Writer, maxItemsPerGroup int)**

Renders a detailed tree view of versions grouped by major version. Shows individual versions with optional release times. maxItemsPerGroup controls truncation (0 = show all).

### VisualizeVersionGroups

**func VisualizeVersionGroups(versions []*Version, w io.Writer)**

Renders a summary tree showing only group-level information (group ID and version count). Useful for large collections.

## Code Examples

```go
package main

import (
    "os"
    "time"
    "github.com/scagogogo/versions"
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

    // Detailed visualization — max 2 per group
    versions.VisualizeVersions(versionList, os.Stdout, 2)
    // Output:
    // 版本总数: 8
    // 版本组数: 2
    //
    // ┌─ 版本组: 1 (3个版本)
    // ├── 1.0.0 (发布时间: 2024-01-01)
    // ├── 1.0.1 (发布时间: 2024-02-01)
    // └── ...还有1个版本未显示
    //
    // ┌─ 版本组: 2 ...

    // Summary visualization
    versions.VisualizeVersionGroups(versionList, os.Stdout)
    // Output:
    // 版本总数: 8
    // 版本组数: 5
    //
    // ├─ 1 (3个版本组, 共3个版本)
    // │  ├─ 1.0 (1个版本)
    // │  └─ 1.1 (2个版本)
    // └─ 2 (3个版本组, 共3个版本)
}
```

## Important Notes

- Output is in Chinese (中文) — version labels like "版本总数", "版本组" etc.
- Uses Unicode box-drawing characters for tree structure (├──, └──, ┌─)
- Release time is only shown when PublicTime is non-zero
- maxItemsPerGroup = 0 means show all versions (no truncation)
- VisualizeVersionGroups shows group hierarchy, not individual versions
- Write to any io.Writer — os.Stdout, bytes.Buffer, or file
