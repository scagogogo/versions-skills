---
name: version-parsing
description: Parse, validate, and extract structured components from version strings. Use when you need to break a version string into its parts (prefix, numbers, suffix, metadata), validate a version, or check what type of version something is (alpha, beta, RC, stable, etc.).
argument-hint: <version-string-or-task>
---

# Version Parsing

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- You have a raw version string (e.g. `"v1.2.3-rc1"`) and need its structured components
- You need to validate whether a string is a valid version number
- You need to determine the version type: prerelease, stable, alpha, beta, RC, dev, snapshot, etc.
- You need to extract segments (major/minor/patch), suffix weight, or core version
- You need to serialize/deserialize versions for JSON, SQL, or text formats

## Decision Tree

```
Raw version string in hand?
├─ Need structured breakdown?         → version_parse / versions parse / NewVersion()
├─ Need yes/no validity check?        → version_validate / versions validate / IsValid()
├─ Need all type flags at once?       → version_info / versions info / Is* methods
├─ Need specific type check?          → Call the specific Is* method or CLI check --<type>
├─ Need segments (major/minor/patch)? → versions segments / Segments(), Major(), Minor(), Patch()
├─ Need the core (suffix stripped)?   → versions core / Core()
└─ Need suffix weight for ordering?   → versions suffix-weight / SuffixWeight()
```

## Task Patterns

### Parse a version string into components

**Goal:** Break `"v1.2.3-rc1"` into prefix=`"v"`, numbers=`[1,2,3]`, suffix=`"-rc1"`.

**SDK approach:**
```go
v := versions.NewVersion("v1.2.3-rc1")
if !v.IsValid() {
    // handle invalid input
}
fmt.Println(v.Prefix, v.VersionNumbers, v.Suffix)
```

**CLI approach:**
```bash
versions parse v1.2.3-rc1
```

**MCP approach:**
```json
{"tool": "version_parse", "arguments": {"version_string": "v1.2.3-rc1"}}
```

### Validate a version string

**Goal:** Confirm `"1.2.3"` is valid, `"not-a-version"` is not.

**SDK approach:**
```go
v := versions.NewVersion("1.2.3")
if err := v.Validate(); err != nil {
    // invalid
}
```

**CLI approach:**
```bash
versions validate 1.2.3    # exit 0 = valid
versions validate notaversion  # exit 1 = invalid
```

**MCP approach:**
```json
{"tool": "version_validate", "arguments": {"version_string": "1.2.3"}}
```

### Check version type (alpha, beta, RC, stable, etc.)

**Goal:** Determine if `"1.0.0-beta2"` is a beta prerelease.

**SDK approach:**
```go
v := versions.NewVersion("1.0.0-beta2")
isBeta := v.IsBeta()        // true
isPrerelease := v.IsPrerelease() // true
isStable := v.IsStable()    // false
subVersion := v.SubVersion() // 2
```

**CLI approach:**
```bash
versions check --beta 1.0.0-beta2     # exit 0
versions check --stable 1.0.0-beta2   # exit 1
versions info v1.2.3-beta1            # all Is* flags at once
```

**MCP approach:**
```json
{"tool": "version_info", "arguments": {"version_string": "v1.2.3-beta1"}}
```

### Extract version segments

**Goal:** Get `[1, 2, 3]` from `"1.2.3"`, or just the major version.

**SDK approach:**
```go
v := versions.NewVersion("1.2.3.4")
major := v.Major()   // 1
minor := v.Minor()   // 2
patch := v.Patch()   // 3
all := v.Segments()  // [1 2 3 4]
```

**CLI approach:**
```bash
versions segments 1.2.3       # [1, 2, 3]
versions segments v1.2.3.4    # [1, 2, 3, 4]
```

**MCP approach:**
```json
{"tool": "version_parse", "arguments": {"version_string": "1.2.3"}}
```

### Get core version (strip suffix)

**Goal:** Convert `"v1.2.3-beta1"` to `"v1.2.3"`.

**SDK approach:**
```go
v := versions.NewVersion("v1.2.3-beta1")
core := v.Core()
fmt.Println(core.RawString()) // "v1.2.3"
```

**CLI approach:**
```bash
versions core v1.2.3-beta1    # v1.2.3
```

**MCP approach:**
```json
{"tool": "version_core", "arguments": {"version_string": "v1.2.3-beta1"}}
```

### Parse with custom delimiters

**Goal:** Parse underscore-separated versions like `"1_2_3"`.

**SDK approach:**
```go
v := versions.NewVersionWithOption("1_2_3", versions.ParserOption{Delimiters: ".-_"})
fmt.Println(v.Segments()) // [1 2 3]
```

**CLI approach:**
```bash
versions parse --delimiters "_-" curl-7_85_0
```

**MCP approach:**
```json
{"tool": "version_parse", "arguments": {"version_string": "1_2_3", "delimiters": ".-_"}}
```

## Cross-References

- [[version-check]] — for boolean type checks and comparison predicates
- [[version-comparison]] — for CompareTo, IsNewerThan, IsOlderThan
- [[version-sorting]] — for sorting parsed versions
- [[version-constraints]] — for constraint expression matching

## Important Notes

- `NewVersion()` never returns nil — always check `IsValid()` for invalid inputs
- `IsValid()` checks for non-empty VersionNumbers; `Validate()` is stricter (also rejects negative numbers)
- `IsPrerelease()` means "has any suffix"; `IsPre()` means "has explicit `-pre` suffix" — they are different
- `IsStable()` means "no suffix"; `IsRelease()` means "has explicit `-release` suffix" — they are different
- Semver build metadata (after `+`) is stored in the `Metadata` field, not in Suffix
- All `With*` methods (WithPrefix, WithSuffix, WithMajor, etc.) return new Version objects — the original is never modified
- `MustParse` panics on invalid input — use only for hardcoded/test data
