---
name: version-file-operations
description: Read version lists from files, write sorted versions to files, or stream versions from any io.Reader.
argument-hint: <file-path>
---

# Version File Operations

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- Reading version numbers from a text file (one per line)
- Writing a sorted version list to a file
- Processing version lists from dependency lock files or release logs
- Using file-based version data as input for sorting, grouping, or filtering
- Building CI/CD tools that read or write version files

## Decision Tree

```
Need to read and parse versions from a file?
  → Use ReadVersionsFromFile / version_read_file
Need raw strings without parsing overhead?
  → Use ReadVersionsStringFromFile / version_read_file with parse=false
Need to read from a non-file source (HTTP, pipe)?
  → Use ReadVersionsFromReader
Need to write versions to a file sorted?
  → Use WriteVersionsToFile / version_write_file
Need file input for another operation (sort, group, filter)?
  → Use --from-file flag (CLI) or read-then-operate (SDK)
```

## Task Patterns

### Read and parse versions from a file

**Goal:** Load a file, parse each line as a Version, and filter out invalid entries.

**SDK approach:**
```go
versionList, err := versions.ReadVersionsFromFile("versions.txt")
// Each line parsed via NewVersion; check v.IsValid() to filter
```

**CLI approach:**
```bash
versions read versions.txt
```

**MCP approach:**
```json
{
  "tool": "version_read_file",
  "arguments": { "filepath": "versions.txt", "parse": true }
}
```

### Read raw strings without parsing

**Goal:** Get version strings from a file without the overhead of parsing into Version objects.

**SDK approach:**
```go
rawStrings, err := versions.ReadVersionsStringFromFile("versions.txt")
```

**CLI approach:**
```bash
versions read-strings versions.txt
```

**MCP approach:**
```json
{
  "tool": "version_read_file",
  "arguments": { "filepath": "versions.txt", "parse": false }
}
```

### Write sorted versions to a file

**Goal:** Take a list of versions, sort them, and write to a file.

**SDK approach:**
```go
versionList := versions.NewVersions("2.0.0", "1.0.0", "1.1.0")
err := versions.WriteVersionsToFile(versionList, "sorted.txt")
// File is automatically sorted before writing
```

**CLI approach:**
```bash
versions write --output sorted.txt 2.0.0 1.0.0 1.1.0
```

**MCP approach:**
```json
{
  "tool": "version_write_file",
  "arguments": {
    "filepath": "sorted.txt",
    "versions": ["2.0.0", "1.0.0", "1.1.0"]
  }
}
```

### Read from any io.Reader (streaming)

**Goal:** Parse versions from a non-file source like an HTTP response or string buffer.

**SDK approach:**
```go
data := strings.NewReader("1.0.0\n1.1.0\n2.0.0\n")
versionList, err := versions.ReadVersionsFromReader(data)
```

### Use file input for other operations

**Goal:** Feed a file's contents directly into sort, group, or range commands.

**CLI approach:**
```bash
versions sort --from-file versions.txt
versions group --from-file versions.txt
versions range 1.0.0 3.0.0 --from-file versions.txt
```

### Full read-filter-write pipeline

**Goal:** Read versions from a file, filter out invalid ones, and write back sorted.

**SDK approach:**
```go
versionList, _ := versions.ReadVersionsFromFile("input.txt")
var valid []*versions.Version
for _, v := range versionList {
    if v.IsValid() {
        valid = append(valid, v)
    }
}
versions.WriteVersionsToFile(valid, "output.txt")
```

## File Format

```
# Comments start with # and are ignored
1.0.0
1.0.1

# Blank lines are also ignored
1.1.0-beta
1.1.0
```

## Cross-References

- [[version-sorting]] — sort versions after reading from a file
- [[version-grouping]] — group versions loaded from a file
- [[version-filtering]] — filter versions read from a file
- [[version-visualization]] — visualize versions loaded from a file

## Important Notes

- **File format**: one version per line; `#` comments; blank lines ignored; leading/trailing whitespace trimmed.
- **ReadVersionsFromFile uses NewVersion** (not NewVersionE) — invalid lines become Version objects with `IsValid() == false`. Always check validity after reading.
- **WriteVersionsToFile sorts before writing** — the output file is always in ascending order.
- **WriteVersionsToFile uses os.WriteFile with permission 0644**.
- **For very large files**, consider streaming with ReadVersionsFromReader instead of ReadVersionsFromFile (which reads the entire file into memory).
- **CLI `--from-file` flag** is available on `sort`, `sort-strings`, `group`, and `range` commands.
- **MCP `version_read_file`** supports both parsed and raw output via the `parse` boolean parameter.
