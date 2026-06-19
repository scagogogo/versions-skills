---
name: version-mutation
description: Bump version numbers, modify version components, strip suffixes, or construct versions from parts.
argument-hint: <mutation-task>
---

# Version Mutation

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

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

**SDK approach:**
```go
v := versions.NewVersion("1.2.3-beta1")
next := v.BumpPatch() // "1.2.4" — suffix cleared
```

**CLI approach:**
```bash
versions bump 1.2.3-beta1 --patch   # 1.2.4
```

**MCP approach:**
```
version_bump(version_string="1.2.3-beta1", bump_type="patch")
```

### Change one component immutably

**Goal:** Modify a single field (prefix, suffix, major, minor, patch) without affecting others.

**SDK approach:**
```go
v := versions.NewVersion("1.2.3-beta1")
v2 := v.WithPrefix("v")  // "v1.2.3-beta1"
v3 := v.WithMajor(2)     // "2.2.3-beta1"
```

**CLI approach:**
```bash
versions set-prefix 1.2.3 v           # v1.2.3
versions set-suffix 1.2.3 -- -beta1   # 1.2.3-beta1
versions set-major 1.2.3 2            # 2.2.3
```

**MCP approach:**
```
version_build(prefix="v", major=2, minor=0, patch=0)
```

### Strip suffix to get core version

**Goal:** Remove the prerelease suffix, keeping prefix and numbers.

**SDK approach:**
```go
v := versions.NewVersion("v1.2.3-beta1")
core := v.Core() // "v1.2.3"
```

**CLI approach:**
```bash
versions core v1.2.3-beta1   # v1.2.3
```

**MCP approach:**
```
version_core(version_string="v1.2.3-beta1")
```

### Build a version from parts

**Goal:** Construct a version string from individual components.

**SDK approach:**
```go
v := versions.NewVersionBuilder().
    Prefix("v").Major(1).Minor(2).Patch(3).
    Suffix("-alpha1").Build()
```

**CLI approach:**
```bash
versions build --prefix v --major 1 --minor 2 --patch 3 --suffix -alpha1
versions build --numbers 1,2,3,4   # 1.2.3.4
```

**MCP approach:**
```
version_build(prefix="v", major=1, minor=2, patch=3, suffix="-alpha1")
```

### Replace all version numbers

**Goal:** Change all numeric segments while keeping prefix and suffix.

**SDK approach:**
```go
v := versions.NewVersion("v1.2.3-beta1")
v2 := v.WithNumbers([]int{4, 5, 6}) // "v4.5.6-beta1"
```

**CLI approach:**
```bash
versions set-numbers v1.2.3 4,5,6   # v4.5.6
```

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
