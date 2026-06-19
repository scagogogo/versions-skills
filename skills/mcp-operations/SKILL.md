---
name: mcp-operations
description: Invoke version operations via the versions MCP server — tool naming, parameter conventions, server configuration, and batch patterns.
argument-hint: <mcp-tool-or-config-task>
---

# MCP Operations

> **Prerequisite:** See `/installation` skill for MCP server binary setup.

## When to Use

- Invoking version operations as MCP tools from an AI agent
- Configuring the versions MCP server (stdio vs SSE transport)
- Understanding MCP tool naming conventions and parameter formats
- Composing multiple MCP tool calls to accomplish a multi-step task
- Working with the structured JSON responses from MCP tools

## Decision Tree

```
Need to configure the MCP server?
  → Add to settings.json mcpServers block
Need to run the server over network (not local)?
  → Use --transport sse --port <port>
Need to find the right tool for a task?
  → All tools are prefixed version_ — see Tool Catalog below
Need to batch multiple operations?
  → Make sequential tool calls, passing results between them
Need file access from MCP?
  → Use version_read_file / version_write_file (operates on server's filesystem)
```

## Task Patterns

### Configure the MCP server

**Goal:** Register the versions MCP server in Claude Code.

Add to `.claude/settings.json` or `~/.claude/settings.json`:
```json
{
  "mcpServers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}
```

For SSE mode (network-accessible server):
```bash
versions-mcp --transport sse --port 8080
```

### Parse and inspect a version

**Goal:** Get all properties of a version in one call.

```
version_parse(version_string="v1.2.3-beta1")
version_info(version_string="v1.2.3-beta1")
version_validate(version_string="v1.2.3-beta1")
```

Use `version_info` when you need all `Is*` flags at once. Use `version_parse` when you only need structural components. Use `version_validate` for a simple valid/invalid check.

### Compare and sort versions

**Goal:** Compare two versions or sort a list.

```
version_compare(version1="1.2.3", version2="2.0.0")
version_sort(versions=["2.0.0", "1.0.0", "1.5.0"], descending=false)
```

### Filter versions by criteria

**Goal:** Filter a list to only stable, prerelease, or constraint-matching versions.

```
version_filter(versions=["1.0.0-alpha", "1.0.0", "2.0.0-beta", "2.0.0"], stable=true)
version_filter(versions=["1.0.0", "1.5.0", "2.0.0"], constraint=">=1.0.0,<2.0.0")
```

### Find min, max, or latest

**Goal:** Find the extreme version in a list.

```
version_min(versions=["1.0.0", "2.0.0", "1.5.0"])
version_max(versions=["1.0.0", "2.0.0", "1.5.0"])
version_latest_stable(versions=["1.0.0-alpha", "1.0.0", "2.0.0-beta", "2.0.0"])
version_latest_prerelease(versions=["1.0.0-alpha", "1.0.0", "2.0.0-beta", "2.0.0"])
```

### Group versions

**Goal:** Cluster versions by their group ID.

```
version_group(versions=["1.0.0", "1.0.0-alpha", "1.0.1", "2.0.0", "2.0.0-beta"])
```

### Check constraints and ranges

**Goal:** Test if a version satisfies a constraint expression or falls in a range.

```
version_constraint_check(expression=">=1.0.0,<2.0.0", version="1.5.0", type="set")
version_range_query(start="1.0.0", end="3.0.0", versions=["1.0.0", "1.5.0", "2.0.0", "3.0.0", "4.0.0"])
```

### Mutate versions

**Goal:** Bump, build, or strip versions.

```
version_bump(version_string="1.2.3", bump_type="patch")
version_build(prefix="v", major=1, minor=2, patch=3, suffix="-alpha1")
version_core(version_string="v1.2.3-beta1")
```

### Set operations

**Goal:** Remove duplicates or compute set difference/intersection/union.

```
version_unique(versions=["1.0.0", "2.0.0", "1.0.0", "2.0.0"])
version_set_operation(operation="difference", set_a=["1.0.0", "2.0.0"], set_b=["2.0.0", "3.0.0"])
```

### Read and write version files

**Goal:** Read versions from or write versions to files on the MCP server's filesystem.

```
version_read_file(filepath="versions.txt", parse=true)
version_write_file(filepath="sorted.txt", versions=["2.0.0", "1.0.0", "1.1.0"])
```

### Visualize versions

**Goal:** Generate a text tree of version hierarchy.

```
version_visualize(versions=["1.0.0", "1.0.1", "1.1.0", "2.0.0"], max_items_per_group=5)
version_visualize(versions=["1.0.0", "1.0.1", "2.0.0"], groups_only=true)
```

### Batch multiple operations

**Goal:** Compose a multi-step workflow across MCP tool calls.

Example workflow — find the latest stable version in a group:
```
# Step 1: Group versions
version_group(versions=["1.0.0-alpha", "1.0.0", "1.0.1", "2.0.0-beta", "2.0.0"])

# Step 2: From the "1" group, find latest stable
version_latest_stable(versions=["1.0.0-alpha", "1.0.0", "1.0.1"])
```

## Tool Catalog

All tools use the `version_` prefix:

| Category | Tools |
|----------|-------|
| Parse & Validate | `version_parse`, `version_validate`, `version_info` |
| Compare | `version_compare` |
| Sort & Filter | `version_sort`, `version_filter` |
| Group & Range | `version_group`, `version_range_query` |
| Constraints | `version_constraint_check` |
| Min/Max | `version_min`, `version_max`, `version_latest_stable`, `version_latest_prerelease` |
| Set Operations | `version_unique`, `version_set_operation` |
| Mutation | `version_build`, `version_bump`, `version_core` |
| File I/O | `version_read_file`, `version_write_file` |
| Visualization | `version_visualize` |

## Cross-References

- [[cli-operations]] — equivalent operations via the CLI binary
- [[version-check]] — `version_info` returns all Is* flags
- [[version-sorting]] — `version_sort` and `version_latest_*`
- [[version-mutation]] — `version_bump`, `version_build`, `version_core`
- [[version-file-operations]] — `version_read_file`, `version_write_file`
- [[version-visualization]] — `version_visualize`
- [[version-grouping]] — `version_group`
- [[version-constraints]] — `version_constraint_check`
- [[version-range-query]] — `version_range_query`
- [[version-comparison]] — `version_compare`

## Important Notes

- **All tool names are prefixed with `version_`** to avoid collision with other MCP servers.
- **The `versions` parameter is always a JSON array of strings**: `["1.0.0", "2.0.0"]`.
- **`version_constraint_check` type parameter**: `set` (default, comma-separated AND) or `union` (||-separated OR).
- **`version_set_operation` operation parameter**: `difference`, `intersection`, or `union`.
- **`version_bump` bump_type parameter**: `major`, `minor`, or `patch`.
- **`version_read_file` and `version_write_file`** operate on the server's local filesystem, not the client's.
- **All tools return JSON-formatted results** with structured data.
- **For SSE transport**, the server listens on the specified port and accepts HTTP connections.
