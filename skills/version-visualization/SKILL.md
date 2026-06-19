---
name: version-visualization
description: Render version hierarchies as text trees for debugging, inspection, or reporting.
argument-hint: <version-list or file-path>
---

# Version Visualization

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- Displaying version structure as an ASCII tree for inspection
- Visualizing how versions are grouped and nested
- Generating a summary view showing only group-level information
- Debugging version collections or presenting version info in terminal output

## Decision Tree

```
Need to show individual versions in a tree?
  → Use VisualizeVersions / version_visualize (without --groups)
Collection has many versions and you need a high-level overview?
  → Use VisualizeVersionGroups / version_visualize with groups_only=true
Need to write output to a buffer or file instead of stdout?
  → Use SDK with io.Writer (bytes.Buffer, os.File)
Reading versions from a file?
  → Use --from-file flag (CLI) or ReadVersionsFromFile (SDK) first
```

## Task Patterns

### Render a detailed version tree

**Goal:** Show all versions grouped by major version with individual entries.

**SDK approach:**
```go
versionList := versions.NewVersions("1.0.0", "1.0.1", "1.1.0", "2.0.0")
versions.VisualizeVersions(versionList, os.Stdout, 0) // 0 = no truncation
```

**CLI approach:**
```bash
versions visualize 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1
```

**MCP approach:**
```json
{
  "tool": "version_visualize",
  "arguments": {
    "versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.0.1"],
    "max_items_per_group": 0
  }
}
```

### Render a group-level summary only

**Goal:** Show only the group hierarchy without individual version entries.

**SDK approach:**
```go
versions.VisualizeVersionGroups(versionList, os.Stdout)
```

**CLI approach:**
```bash
versions visualize --groups 1.0.0 1.0.1 1.1.0 2.0.0
```

**MCP approach:**
```json
{
  "tool": "version_visualize",
  "arguments": {
    "versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0"],
    "groups_only": true
  }
}
```

### Visualize with truncation per group

**Goal:** Limit the number of versions shown per group to avoid overwhelming output.

**SDK approach:**
```go
versions.VisualizeVersions(versionList, os.Stdout, 3) // max 3 per group
```

**CLI approach:**
```bash
versions visualize --max-items 3 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1 2.1.0
```

**MCP approach:**
```json
{
  "tool": "version_visualize",
  "arguments": {
    "versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.0.1", "2.1.0"],
    "max_items_per_group": 3
  }
}
```

### Write visualization to a buffer

**Goal:** Capture the tree output as a string for further processing.

**SDK approach:**
```go
var buf bytes.Buffer
versions.VisualizeVersions(versionList, &buf, 0)
output := buf.String()
```

**CLI approach:**
```bash
versions visualize 1.0.0 1.1.0 2.0.0 > tree.txt
```

### Read versions from a file and visualize

**Goal:** Load versions from a text file and render the tree.

**CLI approach:**
```bash
versions visualize --from-file versions.txt
versions visualize --groups --from-file releases.txt
```

## Cross-References

- [[version-file-operations]] — read version lists from files before visualizing
- [[version-grouping]] — Group() is what powers the tree structure
- [[version-sorting]] — sort versions before visualization for predictable output

## Important Notes

- **Output uses Unicode box-drawing** characters (├──, └──, ┌─) — ensure terminal supports UTF-8.
- **Release time is only shown** when `PublicTime` is non-zero (set it before calling VisualizeVersions).
- **maxItems = 0 means show all** versions per group with no truncation.
- **Output labels are in Chinese** ("版本总数" = total versions, "版本组" = version group, "发布时间" = release time).
- **Both SDK functions accept any `io.Writer`** — use `os.Stdout`, `bytes.Buffer`, or `os.File`.
- **MCP response** includes both `visualization` (the tree string) and `total_versions` / `total_groups` counts.
