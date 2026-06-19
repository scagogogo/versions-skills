---
name: cli-operations
description: Compose the versions CLI in shell pipelines, scripts, and CI/CD — output formats, exit codes, input methods, and command patterns.
argument-hint: <cli-pattern-or-command>
---

# CLI Operations

> **Setup:** See `/installation` for one-time CLI binary install.  
> **Layers:** CLI (shell) — this skill covers CLI-specific patterns. For domain logic, see the corresponding version-* skill.

## When to Use

- Running version operations from shell scripts or CI/CD pipelines
- Composing version commands in shell pipelines with `jq` or other tools
- Using exit-code-based checks for conditional logic (`if versions check ...`)
- Choosing between JSON, table, or text output formats
- Understanding the three input methods: positional args, `--from-file`, and stdin

## Decision Tree

```
Need to pipe output to jq or another tool?
  → Use -q (quiet mode) to strip the JSON envelope
Need exit-code-based conditionals (CI)?
  → Use versions check or versions validate
Need to read versions from a file?
  → Use --from-file flag or stdin pipe
Need human-readable output?
  → Use --format table or --format text
Need to combine multiple operations?
  → Pipe commands together or use shell composition
```

## Task Patterns

### Use exit codes for CI conditionals

**Goal:** Branch on version properties in shell scripts using exit codes.

```bash
# Exit 0 = true, Exit 1 = false
if versions check --stable "$VERSION"; then
    echo "Deploying stable release"
    deploy-production
fi

if versions check --beta "$VERSION"; then
    echo "Beta version — deploying to staging only"
    deploy-staging
fi

# Validate a version string
versions validate "1.2.3-beta1" && echo "Valid" || echo "Invalid"
```

### Pipe output to jq

**Goal:** Extract specific fields from JSON output for scripting.

```bash
# Quiet mode strips the envelope — outputs raw data
versions sort 3.0.0 1.0.0 2.0.0 -q | jq '.[0]'        # "1.0.0"
versions parse v1.2.3 -q | jq '.major'                   # 1
versions info 1.2.3-beta1 -q | jq '{stable, prerelease}' # {"stable":false,"prerelease":true}

# Without -q, the envelope wraps data:
versions parse v1.2.3 | jq '.data.major'                 # 1
```

### Choose output format

**Goal:** Select the right output format for the audience.

```bash
versions sort 1.0.0 2.0.0 1.5.0 --format json   # AI/script consumption (default)
versions sort 1.0.0 2.0.0 1.5.0 --format table  # Human-readable table
versions sort 1.0.0 2.0.0 1.5.0 --format text   # Indented plain text
```

### Feed versions via stdin

**Goal:** Pipe version strings into a command from another process.

```bash
# Stdin is read when no positional arguments are given
cat versions.txt | versions sort
echo -e "2.0.0\n1.0.0\n1.5.0" | versions sort --desc
versions read versions.txt | versions filter --stable
```

### Use --from-file for file input

**Goal:** Read versions from a file without a pipe.

```bash
versions sort --from-file versions.txt
versions sort --from-file versions.txt --desc
versions group --from-file versions.txt
versions range 1.0.0 3.0.0 --from-file versions.txt
```

### Compose multi-step pipelines

**Goal:** Chain multiple version operations together.

```bash
# Sort then filter: get sorted stable versions
versions sort 3.0.0-alpha 1.0.0 2.0.0-beta 2.0.0 | versions filter --stable

# Read from file, filter, then visualize
versions read versions.txt | versions filter --prerelease | versions visualize

# Bump version in CI and tag
NEXT=$(versions bump 1.2.3 --patch -q)
git tag "v$NEXT"

# Count stable versions
versions count --stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0 -q
```

### Parse with custom delimiters

**Goal:** Handle non-standard version separators.

```bash
versions parse --delimiters "_-" curl-7_85_0
```

## API Reference

### Global Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `json` | Output format: `json`, `table`, or `text` |
| `--quiet` | `-q` | false | Quiet mode: output data only, no JSON envelope |
| `--version` | `-v` | — | Print binary version |

### Output Format Reference

Default JSON envelope structure:
```json
{
  "command": "parse",
  "success": true,
  "data": { ... },
  "error": ""
}
```

Quiet mode (`-q`) outputs only `data` — ideal for piping to `jq`.

### Input Methods

All commands that accept version lists support three input methods:

1. **Positional arguments**: `versions sort 2.0.0 1.0.0 1.5.0`
2. **File input**: `versions sort --from-file versions.txt`
3. **Stdin pipe**: `cat versions.txt | versions sort`

Stdin is only read when no positional arguments are provided.

### Complete Command Reference

#### Parsing & Validation

```bash
versions parse <v>                              # structured JSON output
versions parse --delimiters "_-" <v>            # custom delimiters
versions validate <v>                           # exit 0 = valid, 1 = invalid
versions info <v>                               # all Is* flags + segments
```

#### Property Checks (exit 0 = true, 1 = false)

```bash
versions check --prerelease <v>
versions check --stable <v>
versions check --dev <v>
versions check --alpha <v>
versions check --beta <v>
versions check --rc <v>
versions check --snapshot <v>
versions check --milestone <v>
versions check --nightly <v>
versions check --final <v>
versions check --ga <v>
versions check --pre <v>
versions check --release <v>
versions check --sp <v>
versions check --post <v>
versions check --zero <v>
versions check --is-valid <v>
versions check --newer <target> <v>
versions check --older <target> <v>
versions check --equal <target> <v>
versions check --between-low <lo> --between-high <hi> <v>
```

#### Property Access

```bash
versions segments <v>              # numeric segments array
versions sub-version <v>           # sub-version number from suffix
versions suffix-weight <v>         # semantic weight of suffix
versions pure-prefix <v>           # prefix without trailing delimiters
versions group-id <v>              # group ID
versions clone <v>                 # deep copy as JSON
```

#### Comparison

```bash
versions compare <v1> <v2>         # -1/0/1
```

#### Sorting & Filtering

```bash
versions sort [versions...]        # ascending; --desc for descending
versions sort-strings [v...]       # sort returning raw strings
versions filter [versions...]      # filter by conditions
versions count [versions...]       # count matching versions
versions partition [versions...]   # split into matched/unmatched groups
```

Filter flags: `--stable`, `--prerelease`, `--major N`, `--minor N`, `--patch N`, `--prefix`, `--suffix`, `--constraint EXPR`, `--constraint-type single|set|union`

Count flags: `--stable`, `--prerelease`, `--major N`, `--minor N`, `--patch N`

Partition flags: `--stable`, `--prerelease`

#### Grouping & Range Queries

```bash
versions group [versions...]                        # group by VersionNumbers
versions group-ids [versions...]                    # list all group IDs
versions group-latest --group-id <id> [v...]        # latest in group
versions group-oldest --group-id <id> [v...]        # oldest in group
versions group-stable --group-id <id> [v...]        # stable in group
versions group-prerelease --group-id <id> [v...]    # prerelease in group
versions group-latest-stable --group-id <id> [v...] # latest stable in group
versions group-latest-prerelease --group-id <id> [v...] # latest prerelease in group
versions group-contains --group-id <id> --version <v> [v...] # check membership
versions range <start> <end> [v...]                 # query in range
```

Range flags: `--include-start`, `--include-end`

#### Constraints

```bash
versions constraint <expr> <v>    # check constraint (--type single|set|union)
versions satisfies <v> <expr>     # version-centric, auto-detects type
```

#### Min/Max

```bash
versions min [versions...]               # minimum version
versions max [versions...]               # maximum version
versions latest-stable [versions...]     # latest stable
versions latest-prerelease [versions...] # latest prerelease
```

#### Set Operations

```bash
versions unique [versions...]            # remove duplicates
```

#### Construction & Mutation

```bash
versions build --prefix v --major 1 --minor 2 --patch 3 --suffix -alpha1
versions build --numbers 1,2,3,4
versions bump <v> --major/--minor/--patch
versions core <v>
versions set-prefix <v> <prefix>
versions set-suffix <v> -- <suffix>
versions set-major <v> <n>
versions set-minor <v> <n>
versions set-patch <v> <n>
versions set-numbers <v> <n,n,...>
```

#### File I/O

```bash
versions read <filepath>             # read and parse
versions read-strings <filepath>     # read raw strings
versions write --output <path> <v...> # write sorted
```

#### Visualization

```bash
versions visualize [versions...]     # text tree (--max-items N, --groups)
```

### Common Patterns

```bash
# Quick property check (CI)
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

# Filter by constraint
versions filter --constraint ">=1.0.0" --constraint-type single 0.5.0 1.0.0 2.0.0

# Get latest stable in a group
versions group-latest-stable --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0

# Quiet mode for piping
versions sort 3.0.0 1.0.0 2.0.0 -q | jq '.[0]'  # "1.0.0"
```

## Cross-References

- [[version-check]] — `versions check` exit-code-based property checks
- [[version-parsing]] — `versions parse` and `versions validate`
- [[version-sorting]] — `versions sort` and `versions sort-strings`
- [[version-mutation]] — `versions bump`, `versions build`, `versions set-*`
- [[version-file-operations]] — `versions read`, `versions write`, `--from-file`
- [[version-visualization]] — `versions visualize`
- [[version-grouping]] — `versions group` and related subcommands
- [[version-constraints]] — `versions constraint` and `versions satisfies`
- [[version-comparison]] — `versions compare`
- [[version-range-query]] — `versions range`

## Important Notes

- **Default output is JSON** for AI/script consumption; use `-q` for piping to `jq`.
- **`validate` returns exit code 0** for valid, 1 for invalid versions.
- **`check` returns exit code 0** for true, 1 for false — perfect for CI conditionals.
- **Stdin is only read** when no positional arguments are provided.
- **`write` command sorts** versions before writing to the file.
- **`set-*` commands return new versions** — the original is never modified (immutable).
- **For suffixes starting with `-`**, use `--` separator: `versions set-suffix 1.2.3 -- -beta1`.
- **Global flags**: `--format` / `-f` (json|table|text), `--quiet` / `-q`, `--version` / `-v`.
