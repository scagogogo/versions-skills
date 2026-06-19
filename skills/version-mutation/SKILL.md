---
name: version-mutation
description: Bump version numbers, modify version components, strip suffixes, or construct versions from parts.
argument-hint: <mutation-task>
---

# Version Mutation

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

## When to Use

- Bumping a version number (major, minor, or patch) for a release
- Changing a specific component: prefix, suffix, major, minor, patch, or all numbers
- Stripping a suffix to get the core version
- Constructing a version from individual parts (prefix, numbers, suffix)

## Decision Tree

```
Need to increment a version for release?
  → Use BumpMajor / BumpMinor / BumpPatch (clears suffix)
Need to change one field while keeping the rest?
  → Use With* methods (immutable, preserves suffix)
Need to strip the suffix?
  → Use Core()
Need to build from scratch?
  → Use VersionBuilder or version_build
```

## Task Patterns

### Bump a version for release

**Goal:** Increment the version number and clear the prerelease suffix.

| Layer | Approach |
|-------|----------|
| SDK | `v := versions.NewVersion("1.2.3-beta1"); next := v.BumpPatch()` |
| CLI | `versions bump 1.2.3-beta1 --patch` |
| MCP | `{"tool": "version_bump", "arguments": {"version_string": "1.2.3-beta1", "bump_type": "patch"}}` |

### Change one component immutably

**Goal:** Modify a single field (prefix, suffix, major, minor, patch) without affecting others.

| Layer | Approach |
|-------|----------|
| SDK | `v.WithPrefix("v")`, `v.WithMajor(2)`, `v.WithSuffix("-beta1")` |
| CLI | `versions set-prefix 1.2.3 v`, `versions set-major 1.2.3 2`, `versions set-suffix 1.2.3 -- -beta1` |
| MCP | `{"tool": "version_build", "arguments": {"prefix": "v", "major": 2, "minor": 0, "patch": 0}}` |

### Strip suffix to get core version

**Goal:** Remove the prerelease suffix, keeping prefix and numbers.

| Layer | Approach |
|-------|----------|
| SDK | `v.Core()` → `"v1.2.3"` from `"v1.2.3-beta1"` |
| CLI | `versions core v1.2.3-beta1` |
| MCP | `{"tool": "version_core", "arguments": {"version_string": "v1.2.3-beta1"}}` |

### Build a version from parts

**Goal:** Construct a version string from individual components.

| Layer | Approach |
|-------|----------|
| SDK | `versions.NewVersionBuilder().Prefix("v").Major(1).Minor(2).Patch(3).Suffix("-alpha1").Build()` |
| CLI | `versions build --prefix v --major 1 --minor 2 --patch 3 --suffix -alpha1` |
| MCP | `{"tool": "version_build", "arguments": {"prefix": "v", "major": 1, "minor": 2, "patch": 3, "suffix": "-alpha1"}}` |

### Replace all version numbers

**Goal:** Change all numeric segments while keeping prefix and suffix.

| Layer | Approach |
|-------|----------|
| SDK | `v.WithNumbers([]int{4, 5, 6})` → `"v4.5.6-beta1"` |
| CLI | `versions set-numbers v1.2.3 4,5,6` |
| MCP | `{"tool": "version_build", "arguments": {"prefix": "v", "numbers": [4, 5, 6], "suffix": "-beta1"}}` |

## API Reference

### SDK — Bump Operations

All bump methods clear the suffix and return a new Version:

```go
v.BumpMajor() *Version  // 1.2.3 → 2.0.0
v.BumpMinor() *Version  // 1.2.3 → 1.3.0
v.BumpPatch() *Version  // 1.2.3 → 1.2.4
```

### SDK — Immutable Modification (With* methods)

All `With*` methods return a **new** Version — the original is never modified:

```go
v.WithPrefix(prefix string) *Version       // change prefix (e.g. "" → "v")
v.WithSuffix(suffix string) *Version       // change suffix (e.g. "" → "-beta1")
v.WithMajor(major int) *Version            // change Major number
v.WithMinor(minor int) *Version            // change Minor number
v.WithPatch(patch int) *Version            // change Patch number
v.WithNumbers(numbers []int) *Version      // replace all version numbers
v.WithPublicTime(t time.Time) *Version     // set release time
v.WithMetadata(m string) *Version          // set build metadata
```

### SDK — Core

```go
v.Core() *Version  // strip suffix, e.g. "v1.2.3-beta1" → "v1.2.3"
```

### SDK — VersionBuilder

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

### CLI Commands

```bash
# Bump
versions bump <v> --major              # 2.0.0
versions bump <v> --minor              # 1.3.0
versions bump <v> --patch              # 1.2.4

# Core (strip suffix)
versions core <v>                      # strip suffix

# Set (immutable modification)
versions set-prefix <v> <prefix>       # change prefix
versions set-suffix <v> -- <suffix>    # change suffix (-- required for -prefix)
versions set-major <v> <n>             # change Major
versions set-minor <v> <n>             # change Minor
versions set-patch <v> <n>             # change Patch
versions set-numbers <v> <n,n,...>     # replace all numbers

# Build
versions build --prefix v --major 1 --minor 2 --patch 3
versions build --numbers 1,2,3,4
versions build --prefix v --major 1 --suffix -alpha1
```

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_bump` | `version_string`, `bump_type` (major/minor/patch) | bumped version string |
| `version_core` | `version_string` | core version string |
| `version_build` | `prefix?`, `major?`, `minor?`, `patch?`, `suffix?`, `numbers?` | constructed version string |

## Cross-References

- [[version-check]] — check if a version is stable/prerelease before bumping
- [[version-properties]] — inspect version components before mutation
- [[version-parsing]] — parse version strings before mutating

## Important Notes

- **All `With*` and `Bump*` methods are immutable** — they return a new Version; the original is unchanged.
- **`Bump*` clears the suffix**: `1.2.3-beta1.BumpPatch()` returns `1.2.4`, not `1.2.4-beta1`.
- **`With*` preserves the suffix** unless you explicitly change it with `WithSuffix`.
- **CLI `set-suffix` needs `--`**: suffixes starting with `-` require `--` separator: `versions set-suffix 1.2.3 -- -beta1`.
- **`WithNumbers` replaces ALL numbers**, overriding Major/Minor/Patch in the builder.
- **`VersionBuilder.Numbers()` overrides Major/Minor/Patch** — use it for arbitrary-segment versions like `1.2.3.4.5`.
