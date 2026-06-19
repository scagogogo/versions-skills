---
name: version-visualization
description: Render version hierarchies as text trees for debugging, inspection, or reporting.
argument-hint: <version-list or file-path>
---

# Version Visualization

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

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

| Layer | Approach |
|-------|----------|
| SDK | `versions.VisualizeVersions(versionList, os.Stdout, 0)` — 0 = no truncation |
| CLI | `versions visualize 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1` |
| MCP | `{"tool": "version_visualize", "arguments": {"versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0", "2.0.1"], "max_items_per_group": 0}}` |

### Render a group-level summary only

**Goal:** Show only the group hierarchy without individual version entries.

| Layer | Approach |
|-------|----------|
| SDK | `versions.VisualizeVersionGroups(versionList, os.Stdout)` |
| CLI | `versions visualize --groups 1.0.0 1.0.1 1.1.0 2.0.0` |
| MCP | `{"tool": "version_visualize", "arguments": {"versions": ["1.0.0", "1.0.1", "1.1.0", "2.0.0"], "groups_only": true}}` |

### Visualize with truncation per group

**Goal:** Limit the number of versions shown per group to avoid overwhelming output.

| Layer | Approach |
|-------|----------|
| SDK | `versions.VisualizeVersions(versionList, os.Stdout, 3)` — max 3 per group |
| CLI | `versions visualize --max-items 3 1.0.0 1.0.1 1.1.0 2.0.0 2.0.1 2.1.0` |
| MCP | `{"tool": "version_visualize", "arguments": {"versions": [...], "max_items_per_group": 3}}` |

### Write visualization to a buffer

**Goal:** Capture the tree output as a string for further processing.

| Layer | Approach |
|-------|----------|
| SDK | `var buf bytes.Buffer; versions.VisualizeVersions(versionList, &buf, 0); output := buf.String()` |
| CLI | `versions visualize 1.0.0 1.1.0 2.0.0 > tree.txt` |
| MCP | Response includes `visualization` (string) and `total_versions` / `total_groups` counts |

### Read versions from a file and visualize

**Goal:** Load versions from a text file and render the tree.

| Layer | Approach |
|-------|----------|
| SDK | `versionList, _ := versions.ReadVersionsFromFile("versions.txt"); versions.VisualizeVersions(versionList, os.Stdout, 0)` |
| CLI | `versions visualize --from-file versions.txt` |
| MCP | Read with `version_read_file` first, then pass result to `version_visualize` |

## API Reference

### SDK Functions

```go
// Detailed tree: individual versions grouped by major version
func VisualizeVersions(versions []*Version, w io.Writer, maxItems int)

// Summary tree: group-level hierarchy only (no individual versions)
func VisualizeVersionGroups(versions []*Version, w io.Writer)
```

- `maxItems = 0` means show all versions per group (no truncation)
- Both functions accept any `io.Writer` — `os.Stdout`, `bytes.Buffer`, or `os.File`
- Release time is only shown when `PublicTime` is non-zero
- Output uses Unicode box-drawing characters (├──, └──, ┌─)
- Labels are in Chinese: "版本总数" (total), "版本组" (group), "发布时间" (release time)

Output structure for `VisualizeVersions`:
```
版本总数: 8
版本组数: 2

┌─ 版本组: 1 (3个版本)
├── 1.0.0 (发布时间: 2024-01-01)
├── 1.0.1 (发布时间: 2024-02-01)
└── ...还有1个版本未显示
```

Output structure for `VisualizeVersionGroups`:
```
版本总数: 8
版本组数: 5

├─ 1 (3个版本组, 共3个版本)
│  ├─ 1.0 (1个版本)
│  └─ 1.1 (2个版本)
└─ 2 (3个版本组, 共3个版本)
```

### CLI Commands

```bash
versions visualize [versions...]              # render tree from positional args
versions visualize --max-items <n> [v...]     # truncate to n per group
versions visualize --groups [v...]            # group-level summary only
versions visualize --from-file <path>         # read versions from file
versions visualize --groups --from-file <path> # combine flags
```

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_visualize` | `versions` ([]string), `max_items_per_group?` (int, default 0), `groups_only?` (bool, default false) | `visualization` (string), `total_versions` (int), `total_groups` (int) |

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
