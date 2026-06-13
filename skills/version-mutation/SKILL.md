---
name: version-mutation
description: Use when modifying version numbers — bumping, setting components, stripping suffixes, or building from parts. Covers BumpMajor/Minor/Patch, With* immutable modifications, Core, and VersionBuilder via SDK, CLI, and MCP.
argument-hint: <version-mutation-task>
---

# Version Mutation Skill

## When to Use

- User needs to bump a version number (major, minor, or patch)
- User needs to change a specific component of a version (prefix, suffix, major, minor, patch, numbers)
- User needs to strip a version's suffix to get the core version
- User needs to construct a version from individual parts

## Quick Start

**SDK:**
```go
v := versions.NewVersion("1.2.3-beta1")
v.BumpMinor()          // 1.3.0
v.Core()               // 1.2.3
v.WithPrefix("v")      // v1.2.3-beta1
```

**CLI:**
```bash
versions bump 1.2.3 --minor       # 1.3.0
versions core v1.2.3-beta1        # v1.2.3
versions set-prefix 1.2.3 v       # v1.2.3
versions build --major 1 --minor 2 --patch 3  # 1.2.3
```

**MCP:**
```
version_bump(version_string="1.2.3", bump_type="minor")
version_core(version_string="v1.2.3-beta1")
version_build(major=1, minor=2, patch=3)
```

## API Reference — SDK

### Bump Operations

All bump methods clear the suffix and return a new Version:

| Method | Example | Result |
|--------|---------|--------|
| `BumpMajor()` | 1.2.3 → | 2.0.0 |
| `BumpMinor()` | 1.2.3 → | 1.3.0 |
| `BumpPatch()` | 1.2.3 → | 1.2.4 |

### Immutable Modification (With* methods)

All `With*` methods return a **new** Version — the original is never modified:

| Method | Signature | Description |
|--------|-----------|-------------|
| `WithPrefix(prefix string)` | `*Version` | Change prefix (e.g. "" → "v") |
| `WithSuffix(suffix string)` | `*Version` | Change suffix (e.g. "" → "-beta1") |
| `WithMajor(major int)` | `*Version` | Change Major number |
| `WithMinor(minor int)` | `*Version` | Change Minor number |
| `WithPatch(patch int)` | `*Version` | Change Patch number |
| `WithNumbers(numbers []int)` | `*Version` | Replace all version numbers |
| `WithPublicTime(t time.Time)` | `*Version` | Set release time |

### Core

**func (v *Version) Core() *Version**

Returns a new Version with the suffix removed. E.g. `v1.2.3-beta1` → `v1.2.3`.

### VersionBuilder

```go
builder := versions.NewVersionBuilder()
builder.Prefix("v")
builder.Major(1)
builder.Minor(2)
builder.Patch(3)
builder.Suffix("-beta1")
builder.Numbers([]int{1, 2, 3, 4})  // overrides Major/Minor/Patch
v := builder.Build()
```

## CLI Commands

### Bump

```bash
versions bump 1.2.3 --major    # 2.0.0
versions bump 1.2.3 --minor    # 1.3.0
versions bump 1.2.3 --patch    # 1.2.4
```

### Core (strip suffix)

```bash
versions core v1.2.3-beta1     # v1.2.3
```

### Set (immutable modification)

```bash
versions set-prefix 1.2.3 v           # v1.2.3
versions set-suffix 1.2.3 -- -beta1   # 1.2.3-beta1
versions set-major 1.2.3 2            # 2.2.3
versions set-minor 1.2.3 5            # 1.5.3
versions set-patch 1.2.3 9            # 1.2.9
versions set-numbers v1.2.3 4,5,6     # v4.5.6
```

### Build

```bash
versions build --prefix v --major 1 --minor 2 --patch 3    # v1.2.3
versions build --numbers 1,2,3,4                            # 1.2.3.4
versions build --prefix v --major 1 --suffix -alpha1        # v1.0.0-alpha1
```

## MCP Tools

| Tool | Parameters | Description |
|------|-----------|-------------|
| `version_bump` | `version_string`, `bump_type` (major/minor/patch) | Bump version |
| `version_core` | `version_string` | Strip suffix |
| `version_build` | `prefix`, `major`, `minor`, `patch`, `suffix` | Build version |

## Code Examples

```go
package main

import (
    "fmt"
    "github.com/scagogogo/versions-skills"
)

func main() {
    v := versions.NewVersion("1.2.3-beta1")

    // Bump operations (clear suffix)
    fmt.Println(v.BumpMajor())  // 2.0.0
    fmt.Println(v.BumpMinor())  // 1.3.0
    fmt.Println(v.BumpPatch())  // 1.2.4
    fmt.Println(v)               // 1.2.3-beta1 (original unchanged!)

    // Core (strip suffix)
    fmt.Println(v.Core())        // 1.2.3

    // Immutable modifications
    v2 := v.WithPrefix("v")
    fmt.Println(v2)  // v1.2.3-beta1

    v3 := v.WithSuffix("")
    fmt.Println(v3)  // 1.2.3

    v4 := v.WithMajor(2)
    fmt.Println(v4)  // 2.2.3-beta1

    v5 := v.WithNumbers([]int{4, 5, 6})
    fmt.Println(v5)  // 4.5.6-beta1

    // Builder pattern
    built := versions.NewVersionBuilder().
        Prefix("v").
        Major(2).
        Minor(0).
        Patch(0).
        Build()
    fmt.Println(built)  // v2.0.0
}
```

## Important Notes

- All `With*` and `Bump*` methods are **immutable** — they return a new Version, leaving the original unchanged
- `Bump*` methods also **clear the suffix** (e.g. `1.2.3-beta1.BumpPatch()` → `1.2.4`, not `1.2.4-beta1`)
- `With*` methods **preserve the suffix** unless you explicitly change it
- `WithNumbers` replaces ALL version numbers, overriding `Major/Minor/Patch` in the builder
- CLI `set-*` commands require `--` before suffixes starting with `-`: `versions set-suffix 1.2.3 -- -beta1`
- The `VersionBuilder` is the most flexible construction method — use `Numbers()` for arbitrary-segment versions like `1.2.3.4.5`
