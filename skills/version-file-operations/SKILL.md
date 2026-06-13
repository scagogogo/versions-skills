---
name: version-file-operations
description: Use when reading version numbers from files, writing version lists to files, or processing version lists stored in text files. Covers SDK, CLI, and MCP access paths for file-based version operations.
argument-hint: <file-path-or-task>
---

# Version File Operations Skill

## When to Use

- User needs to read version numbers from a text file
- User needs to write a sorted version list to a file
- User needs to process a list of versions stored line-by-line in a file
- User needs to parse version lists from dependency lock files or release logs
- User needs to use version data from files as input for sorting, grouping, or other operations
- User is building CI/CD tools that read or write version files

## Quick Start

### SDK (Go)

```go
// Read versions from a file
versionList, err := versions.ReadVersionsFromFile("versions.txt")

// Write sorted versions to a file
err = versions.WriteVersionsToFile(versionList, "output.txt")
```

### CLI

```bash
# Read and display versions from a file
versions read versions.txt

# Read raw version strings (no parsing)
versions read-strings versions.txt

# Write sorted versions to a file
versions write --output output.txt 1.0.0 1.1.0 2.0.0

# Use --from-file as input for other commands
versions sort --from-file versions.txt
```

### MCP

```json
{
  "tool": "version_read_file",
  "arguments": {
    "filepath": "versions.txt"
  }
}
```

## API Reference -- SDK

### ReadVersionsFromFile

**func ReadVersionsFromFile(filepath string) ([]*Version, error)**

Reads a file line-by-line, parses each line as a Version object. Ignores empty lines and lines starting with `#`. Uses NewVersion internally, so invalid lines become Version objects with `IsValid() == false`.

### ReadVersionsStringFromFile

**func ReadVersionsStringFromFile(filepath string) ([]string, error)**

Reads a file line-by-line, returns raw strings. Does NOT parse into Version objects. Ignores empty lines and lines starting with `#`. Useful when you only need the raw strings without parsing overhead.

### ReadVersionsFromReader

**func ReadVersionsFromReader(reader io.Reader) ([]*Version, error)**

Reads versions from any io.Reader (network connections, string buffers, etc.). Same parsing rules as ReadVersionsFromFile. Useful for streaming scenarios where file paths are not available.

### WriteVersionsToFile

**func WriteVersionsToFile(versions []*Version, filepath string) error**

Writes a version list to a file. Each version occupies one line. The versions are sorted before writing, ensuring the output file is in ascending order. Uses os.WriteFile with permission mode 0644.

## CLI Commands

### `versions read`

Read and parse versions from a file, displaying each version with its parsed details.

```bash
versions read <filepath>
```

**Example:**
```bash
versions read versions.txt
# Output (one version per line with parsed info):
# 1.0.0
# 1.0.1
# 1.1.0-beta
# 1.1.0
```

### `versions read-strings`

Read raw version strings from a file without parsing. Useful for inspecting file contents or working with potentially invalid versions.

```bash
versions read-strings <filepath>
```

**Example:**
```bash
versions read-strings versions.txt
# Output (raw strings, no validation):
# 1.0.0
# not-a-version
# 1.1.0-beta
```

### `versions write`

Write version strings to a file. Versions are sorted before writing.

```bash
versions write --output <filepath> <version1> <version2> ...
```

**Flags:**
- `--output <path>` -- Output file path (required)

**Example:**
```bash
versions write --output sorted_versions.txt 2.0.0 1.0.0 1.10.0
# File contents (sorted):
# 1.0.0
# 1.10.0
# 2.0.0
```

### `--from-file` flag (available on many commands)

Several CLI commands accept a `--from-file` flag to read versions from a file instead of passing them as arguments. This flag is available on:

- `versions sort --from-file <path>` -- Sort versions read from a file
- `versions sort-strings --from-file <path>` -- Sort raw version strings from a file
- `versions group --from-file <path>` -- Group versions read from a file
- `versions range --from-file <path>` -- Range query on versions from a file

**Example:**
```bash
# Sort versions stored in a file
versions sort --from-file releases.txt
```

## MCP Tools

### `version_read_file`

Read version strings from a file and parse them into structured version objects.

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `filepath` | string | Yes | Path to the file containing version strings |
| `parse` | boolean | No | Whether to parse into Version objects (default: true). If false, returns raw strings. |

**Example request:**
```json
{
  "tool": "version_read_file",
  "arguments": {
    "filepath": "versions.txt",
    "parse": true
  }
}
```

**Example response:**
```json
{
  "versions": [
    {"raw": "1.0.0", "valid": true, "major": 1, "minor": 0, "patch": 0},
    {"raw": "1.1.0-beta", "valid": true, "prerelease": true},
    {"raw": "not-a-version", "valid": false}
  ]
}
```

### `version_write_file`

Write a list of version strings to a file. Versions are sorted before writing.

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| `filepath` | string | Yes | Output file path |
| `versions` | array of strings | Yes | List of version strings to write |

**Example request:**
```json
{
  "tool": "version_write_file",
  "arguments": {
    "filepath": "sorted_versions.txt",
    "versions": ["2.0.0", "1.0.0", "1.10.0"]
  }
}
```

**Example response:**
```json
{
  "success": true,
  "filepath": "sorted_versions.txt",
  "versions_written": 3
}
```

## Code Examples (SDK)

### Read and Filter Versions from File

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Read and parse versions from file
    versionList, err := versions.ReadVersionsFromFile("versions.txt")
    if err != nil {
        log.Fatalf("Failed to read: %v", err)
    }
    fmt.Printf("Read %d versions\n", len(versionList))

    // Filter valid versions
    validVersions := make([]*versions.Version, 0)
    for _, v := range versionList {
        if v.IsValid() {
            validVersions = append(validVersions, v)
        }
    }
    fmt.Printf("Valid versions: %d\n", len(validVersions))
}
```

### Read Raw Strings from File

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Read raw version strings (no parsing)
    rawStrings, err := versions.ReadVersionsStringFromFile("versions.txt")
    if err != nil {
        log.Fatalf("Failed to read: %v", err)
    }
    for _, s := range rawStrings {
        fmt.Println(s)
    }
}
```

### Write Sorted Versions to File

```go
package main

import (
    "log"
    "github.com/scagogogo/versions-skills"
)

func main() {
    versionList := versions.NewVersions("2.0.0", "1.0.0", "1.1.0", "3.0.0")

    // WriteVersionsToFile sorts before writing
    err := versions.WriteVersionsToFile(versionList, "sorted_versions.txt")
    if err != nil {
        log.Fatalf("Failed to write: %v", err)
    }
    // File contents:
    // 1.0.0
    // 1.1.0
    // 2.0.0
    // 3.0.0
}
```

### Read from io.Reader (Non-file Sources)

```go
package main

import (
    "fmt"
    "strings"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Read versions from a string (e.g., HTTP response body)
    data := strings.NewReader("1.0.0\n1.1.0\n2.0.0\n")
    versionList, err := versions.ReadVersionsFromReader(data)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Read %d versions from reader\n", len(versionList))
}
```

### Read, Sort, and Write Pipeline

```go
package main

import (
    "log"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Full pipeline: read -> sort -> write
    versionList, err := versions.ReadVersionsFromFile("unsorted.txt")
    if err != nil {
        log.Fatalf("Read failed: %v", err)
    }

    // Filter out invalid versions before writing
    valid := versions.SortVersionSlice(versionList)

    err = versions.WriteVersionsToFile(valid, "sorted.txt")
    if err != nil {
        log.Fatalf("Write failed: %v", err)
    }
}
```

## File Format

```
# This is a comment -- lines starting with # are ignored
1.0.0
1.0.1

# Blank lines are also ignored
1.1.0-beta
1.1.0
```

## Important Notes

- **All paths**: File format is one version per line; `#` for comments; blank lines ignored
- **SDK**: Leading/trailing whitespace on each line is trimmed
- **SDK**: ReadVersionsFromFile uses NewVersion (not NewVersionE), so invalid lines become Version objects with `IsValid() == false`
- **SDK**: ReadVersionsStringFromFile returns raw strings without any parsing or validation
- **SDK**: ReadVersionsFromReader supports any io.Reader -- useful for HTTP responses, strings, pipes
- **SDK**: WriteVersionsToFile sorts versions before writing -- the output file is always in ascending order
- **SDK**: WriteVersionsToFile uses os.WriteFile with permission mode 0644
- **SDK**: For very large files, consider streaming instead of ReadVersionsFromFile (which reads the entire file into memory)
- **CLI**: The `--from-file` flag is available on sort, sort-strings, group, and range commands, allowing file-based input for these operations
- **CLI**: `versions read` parses and validates; `versions read-strings` does not parse
- **MCP**: `version_read_file` supports both parsed and raw string output via the `parse` parameter
- **MCP**: `version_write_file` automatically sorts versions before writing, consistent with the SDK behavior