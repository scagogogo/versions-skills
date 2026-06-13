---
name: mcp-operations
description: Use when invoking version operations via the versions MCP server for AI tool use. Provides expert guidance on using MCP tools for version parsing, comparison, sorting, grouping, constraint checking, and more.
argument-hint: <mcp-tool-or-task>
---

# MCP Operations Skill

## When to Use

- AI agent needs to perform version operations as tools via MCP protocol
- User wants to integrate version capabilities into an AI workflow
- User is building an AI-powered dependency analysis or version management system
- User needs programmatic version operations without direct Go SDK usage

## Server Setup

Install the MCP server binary:

```bash
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

Configure in Claude Code `settings.json`:

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

For SSE mode (network access):

```bash
versions-mcp --transport sse --port 8080
```

## Tool Reference

### Parsing & Validation

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_parse` | `version_string` | Parse version into components |
| `version_validate` | `version_string` | Validate version string |
| `version_info` | `version_string` | Full info with all Is* checks |

### Comparison

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_compare` | `version1`, `version2` | Compare two versions |

### Sorting & Filtering

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_sort` | `versions`, `descending` | Sort version list |
| `version_filter` | `versions`, `stable`, `prerelease`, `major`, `minor`, `patch`, `prefix`, `suffix`, `constraint` | Filter by conditions |

### Grouping & Range Queries

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_group` | `versions` | Group by VersionNumbers |
| `version_range_query` | `start`, `end`, `versions`, `include_start`, `include_end` | Query versions in range |

### Constraints

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_constraint_check` | `expression`, `version`, `type` | Check constraint satisfaction |

### Min/Max

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_min` | `versions` | Find minimum version |
| `version_max` | `versions` | Find maximum version |
| `version_latest_stable` | `versions` | Find latest stable version |
| `version_latest_prerelease` | `versions` | Find latest prerelease version |

### Set Operations

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_unique` | `versions` | Remove duplicates |
| `version_set_operation` | `operation`, `set_a`, `set_b` | Difference/intersection/union |

### Construction

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_build` | `prefix`, `major`, `minor`, `patch`, `suffix` | Build version string |
| `version_bump` | `version_string`, `bump_type` | Bump version (major/minor/patch) |
| `version_core` | `version_string` | Strip suffix |

### File I/O

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_read_file` | `filepath` | Read versions from file |
| `version_write_file` | `filepath`, `versions` | Write sorted versions to file |

### Visualization

| Tool | Key Parameters | Description |
|------|---------------|-------------|
| `version_visualize` | `versions`, `max_items_per_group`, `groups_only` | Text tree visualization |

## Common Patterns

```
# Parse a version
version_parse(version_string="v1.2.3-beta1")

# Compare two versions
version_compare(version1="1.2.3", version2="2.0.0")

# Sort versions
version_sort(versions=["2.0.0", "1.0.0", "1.5.0"])

# Filter stable versions
version_filter(versions=["1.0.0-alpha", "1.0.0", "2.0.0"], stable=true)

# Check constraint
version_constraint_check(expression=">=1.0.0,<2.0.0", version="1.5.0")

# Range query
version_range_query(start="1.0.0", end="3.0.0", versions=["1.0.0", "1.5.0", "2.0.0", "3.0.0", "4.0.0"])

# Group versions
version_group(versions=["1.0.0", "1.0.0-alpha", "2.0.0"])

# Bump version
version_bump(version_string="1.2.3", bump_type="patch")

# Set difference
version_set_operation(operation="difference", set_a=["1.0.0", "2.0.0"], set_b=["2.0.0", "3.0.0"])
```

## Important Notes

- All tool names are prefixed with `version_` to avoid collision with other MCP servers
- The `versions` parameter is always a JSON array of strings, e.g. `["1.0.0", "2.0.0"]`
- The `type` parameter in `version_constraint_check` accepts `set` (default, comma-separated AND) or `union` (||-separated OR)
- The `operation` parameter in `version_set_operation` accepts `difference`, `intersection`, or `union`
- The `bump_type` parameter in `version_bump` accepts `major`, `minor`, or `patch`
- All tools return JSON-formatted results with structured data
- `version_read_file` and `version_write_file` operate on the server's local filesystem
