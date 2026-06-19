---
name: version-file-operations
description: Read version lists from files, write sorted versions to files, or stream versions from any io.Reader.
argument-hint: <file-path>
---

# Version File Operations

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

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

| Layer | Approach |
|-------|----------|
| SDK | `versionList, err := versions.ReadVersionsFromFile("versions.txt")` |
| CLI | `versions read versions.txt` |
| MCP | `{"tool": "version_read_file", "arguments": {"filepath": "versions.txt", "parse": true}}` |

### Read raw strings without parsing

**Goal:** Get version strings from a file without the overhead of parsing into Version objects.

| Layer | Approach |
|-------|----------|
| SDK | `rawStrings, err := versions.ReadVersionsStringFromFile("versions.txt")` |
| CLI | `versions read-strings versions.txt` |
| MCP | `{"tool": "version_read_file", "arguments": {"filepath": "versions.txt", "parse": false}}` |

### Write sorted versions to a file

**Goal:** Take a list of versions, sort them, and write to a file.

| Layer | Approach |
|-------|----------|
| SDK | `err := versions.WriteVersionsToFile(versionList, "sorted.txt")` |
| CLI | `versions write --output sorted.txt 2.0.0 1.0.0 1.1.0` |
| MCP | `{"tool": "version_write_file", "arguments": {"filepath": "sorted.txt", "versions": ["2.0.0", "1.0.0", "1.1.0"]}}` |

### Read from any io.Reader (streaming)

**Goal:** Parse versions from a non-file source like an HTTP response or string buffer.

| Layer | Approach |
|-------|----------|
| SDK | `data := strings.NewReader("1.0.0\n1.1.0\n2.0.0\n"); versionList, err := versions.ReadVersionsFromReader(data)` |
| CLI | `cat versions.txt \| versions sort` (stdin pipe) |
| MCP | Not directly supported — use `version_read_file` for server-local files |

### Use file input for other operations

**Goal:** Feed a file's contents directly into sort, group, or range commands.

| Layer | Approach |
|-------|----------|
| SDK | Read first, then pass to operation function |
| CLI | `versions sort --from-file versions.txt`, `versions group --from-file versions.txt`, `versions range 1.0.0 3.0.0 --from-file versions.txt` |
| MCP | Read with `version_read_file`, then pass result to other tools |

### Full read-filter-write pipeline

**Goal:** Read versions from a file, filter out invalid ones, and write back sorted.

| Layer | Approach |
|-------|----------|
| SDK | Read → iterate checking `IsValid()` → `WriteVersionsToFile(valid, "output.txt")` |
| CLI | `versions read versions.txt \| versions filter --stable \| ...` (chain commands) |
| MCP | `version_read_file` → filter client-side → `version_write_file` |

## API Reference

### SDK Functions

```go
// Read and parse versions from a file (uses NewVersion, invalid → IsValid()==false)
func ReadVersionsFromFile(filepath string) ([]*Version, error)

// Read raw strings from a file (no parsing overhead)
func ReadVersionsStringFromFile(filepath string) ([]string, error)

// Read and parse versions from any io.Reader (streaming)
func ReadVersionsFromReader(reader io.Reader) ([]*Version, error)

// Write sorted versions to a file (sorts before writing, perm 0644)
func WriteVersionsToFile(versions []*Version, filepath string) error
```

### File Format

```
# Comments start with # and are ignored
1.0.0
1.0.1

# Blank lines are also ignored
1.1.0-beta
1.1.0
```

### CLI Commands

```bash
versions read <filepath>                  # read and parse, display each version
versions read-strings <filepath>          # read raw strings, no parsing
versions write --output <filepath> <v...> # write sorted versions to file
```

`--from-file` flag is available on: `sort`, `sort-strings`, `group`, `range`.

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_read_file` | `filepath` (string), `parse?` (bool, default true) | parsed versions or raw strings |
| `version_write_file` | `filepath` (string), `versions` ([]string) | `success`, `filepath`, `versions_written` |

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
- **MCP file tools operate on the server's local filesystem**, not the client's.
