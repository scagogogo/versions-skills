---
name: cli-operations
description: Use when running version operations via the versions CLI tool. Provides expert guidance on using the command-line interface for version parsing, comparison, sorting, grouping, constraint checking, property checking, mutation, and more.
argument-hint: <cli-command-or-task>
---

# CLI Operations Skill

## When to Use

- User needs to perform version operations from the command line or in scripts
- User wants to integrate version operations into shell pipelines or CI/CD
- User needs JSON-formatted output for AI consumption or programmatic processing
- User is working with version files and needs batch processing

## Installation

### Option 1: Download from GitHub Releases (Recommended)

Pre-built binaries are available for Linux, macOS, Windows, FreeBSD, OpenBSD, and NetBSD on amd64, arm64, arm, 386, and more architectures.

```bash
# Download the latest release for your platform
# Visit https://github.com/scagogogo/versions-skills/releases/latest

# Linux amd64 example:
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_0.1.0_linux_amd64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# macOS arm64 (Apple Silicon) example:
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_0.1.0_darwin_arm64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# Windows: download the .zip from the releases page and extract
```

### Option 2: Install via Go

```bash
go install github.com/scagogogo/versions-skills/cmd/versions@latest
```

### Option 3: Install from deb/rpm/apk package

```bash
# Debian/Ubuntu:
sudo dpkg -i versions_0.1.0_linux_amd64.deb

# RHEL/CentOS/Fedora:
sudo rpm -i versions_0.1.0_linux_amd64.rpm

# Alpine:
sudo apk add versions_0.1.0_linux_amd64.apk
```

## Global Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `json` | Output format: `json`, `table`, or `text` |
| `--quiet` | `-q` | false | Quiet mode: output data only, no envelope |
| `--version` | `-v` | — | Print binary version |

## Input Methods

All commands that accept version lists support three input methods:

1. **Positional arguments**: `versions sort 2.0.0 1.0.0 1.5.0`
2. **File input**: `versions sort --from-file versions.txt`
3. **Stdin pipe**: `cat versions.txt | versions sort`

## Command Reference

### Parsing & Validation

| Command | Description |
|---------|-------------|
| `versions parse <v>` | Parse and show components (`--delimiters` for custom separators) |
| `versions validate <v>` | Strict validation (exit 0=valid / 1=invalid) |
| `versions info <v>` | Full info with all Is* checks |

### Property Checks (exit 0=true / 1=false)

| Command | Description |
|---------|-------------|
| `versions check --prerelease <v>` | Is prerelease version? |
| `versions check --stable <v>` | Is stable version? |
| `versions check --dev <v>` | Is dev version? |
| `versions check --alpha <v>` | Is Alpha version? |
| `versions check --beta <v>` | Is Beta version? |
| `versions check --rc <v>` | Is RC version? |
| `versions check --snapshot <v>` | Is snapshot version? |
| `versions check --milestone <v>` | Is milestone version? |
| `versions check --nightly <v>` | Is nightly build? |
| `versions check --final <v>` | Is Final version? |
| `versions check --ga <v>` | Is GA version? |
| `versions check --pre <v>` | Is Pre version? |
| `versions check --release <v>` | Is Release version? |
| `versions check --sp <v>` | Is SP version? |
| `versions check --post <v>` | Is Post version? |
| `versions check --zero <v>` | Is zero value? |
| `versions check --newer <target> <v>` | Is newer than target? |
| `versions check --older <target> <v>` | Is older than target? |
| `versions check --equal <target> <v>` | Equals target? |
| `versions check --between-low <lo> --between-high <hi> <v>` | Is between range? |

### Property Access

| Command | Description |
|---------|-------------|
| `versions segments <v>` | Get numeric segments array |
| `versions sub-version <v>` | Get sub-version number from suffix |
| `versions suffix-weight <v>` | Get semantic weight of suffix |
| `versions pure-prefix <v>` | Get prefix without trailing delimiters |
| `versions group-id <v>` | Get group ID (e.g. v1.2.3-beta → "1.2.3") |
| `versions clone <v>` | Deep copy a version |

### Comparison

| Command | Description |
|---------|-------------|
| `versions compare <v1> <v2>` | Compare two versions (-1/0/1) |

### Sorting & Filtering

| Command | Description |
|---------|-------------|
| `versions sort [versions...]` | Sort ascending (`--desc` for descending) |
| `versions sort-strings [versions...]` | Sort returning raw strings |
| `versions filter [versions...]` | Filter by conditions (see filter flags below) |
| `versions count [versions...]` | Count matching versions |
| `versions partition [versions...]` | Split into matched/unmatched groups |

**Filter flags**: `--stable`, `--prerelease`, `--major N`, `--minor N`, `--patch N`, `--prefix`, `--suffix`, `--constraint EXPR`, `--constraint-type single|set|union`

**Count flags**: `--stable`, `--prerelease`, `--major N`, `--minor N`, `--patch N`

**Partition flags**: `--stable`, `--prerelease`

### Grouping & Range Queries

| Command | Description |
|---------|-------------|
| `versions group [versions...]` | Group by VersionNumbers (`--id <id>` for specific group) |
| `versions group-ids [versions...]` | List all group IDs |
| `versions group-latest --group-id <id> [versions...]` | Latest version in group |
| `versions group-oldest --group-id <id> [versions...]` | Oldest version in group |
| `versions group-stable --group-id <id> [versions...]` | Stable versions in group |
| `versions group-prerelease --group-id <id> [versions...]` | Prerelease versions in group |
| `versions group-latest-stable --group-id <id> [versions...]` | Latest stable in group |
| `versions group-latest-prerelease --group-id <id> [versions...]` | Latest prerelease in group |
| `versions group-contains --group-id <id> --version <v> [versions...]` | Check if group contains version |
| `versions range <start> <end> [versions...]` | Query versions in range (`--include-start`, `--include-end`) |

### Constraints

| Command | Description |
|---------|-------------|
| `versions constraint <expr> <v>` | Check if version satisfies constraint (`--type single|set|union`) |
| `versions satisfies <v> <expr>` | Version-centric constraint check (auto-detects type) |

### Min/Max

| Command | Description |
|---------|-------------|
| `versions min [versions...]` | Find minimum version |
| `versions max [versions...]` | Find maximum version |
| `versions latest-stable [versions...]` | Find latest stable version |
| `versions latest-prerelease [versions...]` | Find latest prerelease version |

### Set Operations

| Command | Description |
|---------|-------------|
| `versions unique [versions...]` | Remove duplicates |

### Construction & Mutation

| Command | Description |
|---------|-------------|
| `versions build` | Build version string (`--prefix`, `--major`, `--minor`, `--patch`, `--suffix`, `--numbers 1,2,3`) |
| `versions bump <v> --major/--minor/--patch` | Bump version number |
| `versions core <v>` | Get core version (strip suffix) |
| `versions set-prefix <v> <prefix>` | Immutable: change prefix |
| `versions set-suffix <v> <suffix>` | Immutable: change suffix |
| `versions set-major <v> <n>` | Immutable: change Major |
| `versions set-minor <v> <n>` | Immutable: change Minor |
| `versions set-patch <v> <n>` | Immutable: change Patch |
| `versions set-numbers <v> <n,n,...>` | Immutable: change all numbers |

### File I/O

| Command | Description |
|---------|-------------|
| `versions read <filepath>` | Read versions from file |
| `versions read-strings <filepath>` | Read raw strings without parsing |
| `versions write <filepath> [versions...]` | Write sorted versions to file |

### Visualization

| Command | Description |
|---------|-------------|
| `versions visualize [versions...]` | Visualize version hierarchy (`--max-items N`, `--groups`) |

## Common Patterns

```bash
# Quick property check (useful in CI)
versions check --stable 1.2.3; echo $?     # 0 if stable
versions check --beta 1.2.3-beta1; echo $? # 0 if beta

# Parse with custom delimiters
versions parse --delimiters "_-" curl-7_85_0

# Sort then filter
versions sort 3.0.0 1.0.0 2.0.0 | versions filter --stable

# Bump version in CI
versions bump 1.2.3 --patch  # 1.2.4

# Build version from parts
versions build --prefix v --numbers 1,2,3,4  # v1.2.3.4

# Count stable versions
versions count --stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0

# Partition into stable/prerelease
versions partition --stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0

# Filter by single constraint
versions filter --constraint ">=1.0.0" --constraint-type single 0.5.0 1.0.0 2.0.0

# Get latest stable version in a specific group
versions group-latest-stable --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0

# Quiet mode (data only, no envelope) — ideal for piping
versions sort 3.0.0 1.0.0 2.0.0 -q | jq '.[0]'  # "1.0.0"
```

## Output Format

Default JSON envelope:

```json
{
  "command": "parse",
  "success": true,
  "data": { ... },
  "error": ""
}
```

Quiet mode (`-q`) outputs only the `data` part, no envelope — ideal for piping to `jq`.

Use `--format table` for human-readable tables or `--format text` for indented text.

## Important Notes

- Default output is JSON for AI/script consumption; use `-q` for piping to `jq`
- `validate` returns exit code 0 for valid, 1 for invalid versions
- `check` returns exit code 0 for true, 1 for false — perfect for CI conditionals
- Set operations (`unique`) work on in-memory version lists; `difference`/`intersection`/`union` for file-based set operations are available via SDK and MCP
- `write` command sorts versions before writing
- Stdin is only read when no positional arguments are provided
- `set-*` commands return new versions — the original is never modified
- For suffixes starting with `-`, use `--` separator: `versions set-suffix 1.2.3 -- -beta1`
