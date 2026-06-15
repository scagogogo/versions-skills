---
name: version-check
description: Use when checking boolean properties of version numbers — IsBeta? IsStable? IsNewerThan? IsBetween? Covers all Is* type checks and comparison predicates via SDK, CLI, and MCP.
argument-hint: <version-check-task>
---

# Version Check Skill

## When to Use

- User needs to check if a version is a specific type (alpha, beta, RC, stable, etc.)
- User needs boolean comparison results (is newer? is older? is between?)
- User is building CI/CD conditionals based on version properties
- User needs exit-code-based checks for shell scripts

## Installation

### SDK (Go library)

```bash
go get github.com/scagogogo/versions-skills
```

### CLI binary

**Option A: Download from GitHub Releases (Recommended)**

Pre-built binaries for Linux, macOS, Windows, FreeBSD, OpenBSD, and NetBSD on amd64, arm64, arm, 386, mips, mips64, mips64le, ppc64, ppc64le, s390x, and riscv64. Linux packages: deb, rpm, apk.

```bash
# Linux (amd64)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_linux_amd64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# macOS (arm64 / Apple Silicon)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_darwin_arm64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# Or install via package manager (Linux only):
# Debian/Ubuntu: dpkg -i versions_{VERSION}_linux_amd64.deb
# RHEL/Fedora:   rpm -i versions_{VERSION}_linux_amd64.rpm
# Alpine:        apk add versions_{VERSION}_linux_amd64.apk
```

> Replace `{VERSION}` with the latest release tag (e.g. `0.2.0`). See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all available platforms and the current version.

**Option B: Install via Go**

```bash
go install github.com/scagogogo/versions-skills/cmd/versions@latest
```

### MCP server

**Option A: Download from GitHub Releases (Recommended)**

```bash
# Linux (amd64)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions-mcp_{VERSION}_linux_amd64.tar.gz | tar xz
chmod +x versions-mcp && sudo mv versions-mcp /usr/local/bin/

# macOS (arm64 / Apple Silicon)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions-mcp_{VERSION}_darwin_arm64.tar.gz | tar xz
chmod +x versions-mcp && sudo mv versions-mcp /usr/local/bin/
```

> Replace `{VERSION}` with the latest release tag. See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all platforms.

**Option B: Install via Go**

```bash
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

**SDK:**
```go
v := versions.NewVersion("1.2.3-beta1")
v.IsBeta()       // true
v.IsPrerelease() // true
v.IsStable()     // false
```

**CLI:**
```bash
versions check --beta 1.2.3-beta1     # exit 0 (true)
versions check --stable 1.2.3-beta1   # exit 1 (false)
versions check --newer 1.0.0 2.0.0    # exit 0 (true)
```

**MCP:**
```
version_info(version_string="1.2.3-beta1")  # returns all Is* flags
```

## API Reference — SDK

### Type Check Methods

All return `bool`. A version is exactly one of these types based on its suffix:

| Method | Suffix Pattern | Weight |
|--------|---------------|--------|
| `IsDev()` | `-dev*` | 50 |
| `IsSnapshot()` | `-snapshot*` | 60 |
| `IsNightly()` | `-nightly*` | 70 |
| `IsAlpha()` | `-alpha*` | 100 |
| `IsBeta()` | `-beta*` | 200 |
| `IsMilestone()` | `-milestone*` | 300 |
| `IsRC()` | `-rc*` | 400 |
| `IsFinal()` | `-final*` | 500 |
| `IsGA()` | `-ga*` | 500 |
| `IsRelease()` | `-release*` | 500 |
| `IsPre()` | `-pre*` | — |
| `IsSP()` | `-sp*` | 600 |
| `IsPost()` | `-post*` | 800 |

### Composite Checks

| Method | Returns true when |
|--------|-------------------|
| `IsPrerelease()` | Suffix is non-empty (any suffix = prerelease) |
| `IsStable()` | Suffix is empty (no suffix = stable) |
| `IsZero()` | All version numbers are zero |

### Comparison Predicates

| Method | Signature | Description |
|--------|-----------|-------------|
| `IsNewerThan` | `(target *Version) bool` | Is this version newer than target? |
| `IsOlderThan` | `(target *Version) bool` | Is this version older than target? |
| `Equals` | `(target *Version) bool` | Is this version equal to target? |
| `IsBetween` | `(low, high *Version) bool` | Is this version in [low, high) range? |

### Comparison Order

`CompareTo` compares in this order: **VersionNumbers → Suffix → PublicTime → Raw string**

## CLI Commands

### `versions check <version-string>`

Returns JSON result and uses exit code (0=true, 1=false):

```bash
# Type checks
versions check --prerelease 1.2.3-alpha
versions check --stable 1.2.3
versions check --dev 1.2.3-dev
versions check --alpha 1.2.3-alpha1
versions check --beta 1.2.3-beta2
versions check --rc 1.2.3-rc1
versions check --snapshot 1.2.3-snapshot
versions check --milestone 1.2.3-m1
versions check --nightly 1.2.3-nightly
versions check --final 1.2.3-final
versions check --ga 1.2.3-ga
versions check --pre 1.2.3-pre
versions check --release 1.2.3-release
versions check --sp 1.2.3-sp1
versions check --post 1.2.3-post
versions check --zero 0.0.0

# Comparison checks
versions check --newer 1.0.0 2.0.0
versions check --older 2.0.0 1.0.0
versions check --equal 1.0.0 1.0.0
versions check --between-low 1.0.0 --between-high 3.0.0 2.0.0
```

### CI/CD Pattern

```bash
# Only deploy if version is stable
if versions check --stable $VERSION; then
  deploy
fi

# Only deploy to production if version is GA
if versions check --ga $VERSION; then
  deploy-production
fi
```

## MCP Tools

Use `version_info` to get all boolean properties at once:

```
version_info(version_string="1.2.3-beta1")
```

Returns a JSON object with all `Is*` flags.

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    // Type checks
    v := versions.NewVersion("1.2.3-rc1")
    fmt.Println(v.IsRC())         // true
    fmt.Println(v.IsPrerelease()) // true
    fmt.Println(v.IsStable())     // false
    fmt.Println(v.IsBeta())       // false (RC ≠ Beta)

    // Suffix weight for ordering
    fmt.Println(v.SuffixWeight()) // rc (400)

    // Comparison predicates
    v1 := versions.NewVersion("1.2.3")
    v2 := versions.NewVersion("2.0.0")
    fmt.Println(v2.IsNewerThan(v1))  // true
    fmt.Println(v1.IsOlderThan(v2))  // true
    fmt.Println(v1.Equals(v1))       // true

    // Range check
    v3 := versions.NewVersion("1.5.0")
    low := versions.NewVersion("1.0.0")
    high := versions.NewVersion("2.0.0")
    fmt.Println(v3.IsBetween(low, high)) // true

    // Zero check
    zero := versions.NewVersion("0.0.0")
    fmt.Println(zero.IsZero())  // false (0.0.0 has valid numbers)
}
```

## Important Notes

- `IsStable()` and `IsPrerelease()` are exact opposites: a version is either stable (no suffix) or prerelease (has suffix)
- Type checks are suffix-based: `1.2.3-beta1` is `IsBeta()` but not `IsAlpha()`, even though both are prerelease
- `IsZero()` checks if all version numbers are zero — `0.0.0` has valid numbers, so `IsValid()` is true but `IsZero()` depends on the implementation
- `IsBetween(low, high)` uses half-open interval: includes `low`, excludes `high` (matches SDK behavior)
- CLI `check` exit codes make it ideal for shell conditionals in CI/CD pipelines
- `IsPre()` and `IsPrerelease()` are different: `IsPre()` matches the `-pre*` suffix type, while `IsPrerelease()` matches any non-empty suffix
