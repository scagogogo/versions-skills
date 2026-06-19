---
name: cli-operations
description: Compose the versions CLI in shell pipelines, scripts, and CI/CD — output formats, exit codes, input methods, and command patterns.
argument-hint: <cli-pattern-or-command>
---

# CLI Operations

> **Prerequisite:** See `/installation` skill for CLI binary setup.

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

## Output Format Reference

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
