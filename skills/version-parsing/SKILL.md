---
name: version-parsing
description: Parse, validate, and extract structured components from version strings via SDK, CLI, or MCP. Covers NewVersion/MustParse, all Is* type checks, Segments, Core, Clone, SuffixWeight, custom parser options.
argument-hint: <version-string-or-task>
---

# Version Parsing

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

## When to Use

- Parse a version string (`"v1.2.3-rc1"`) into structured components (prefix, numbers, suffix)
- Validate whether a string is a valid version number
- Determine version type: prerelease, stable, alpha, beta, RC, dev, snapshot, nightly, etc.
- Extract segments, sub-version, suffix weight, group ID, or core version
- Use custom delimiters for non-standard formats (e.g. underscore-separated)

## Decision Tree

```
Raw version string → what do you need?
├─ Structured breakdown?          → NewVersion() / version_parse / versions parse
├─ Yes/no validity?               → IsValid() / version_validate / versions validate
├─ All type flags at once?        → version_info / versions info
├─ Specific type check?           → IsBeta(), IsStable(), IsRC() etc.
├─ Segments (major/minor/patch)?  → Major()/Minor()/Patch()/Segments()
├─ Core version (no suffix)?      → Core()
└─ Suffix weight for ordering?    → SuffixWeight()
```

## Task Patterns

### Parse & inspect

**Goal:** Break `"v1.2.3-rc1"` into prefix `"v"`, numbers `[1,2,3]`, suffix `"-rc1"`.

| Layer | Approach |
|-------|----------|
| SDK | `v := versions.NewVersion("v1.2.3-rc1"); v.IsValid()` |
| CLI | `versions parse v1.2.3-rc1` |
| MCP | `{"tool": "version_parse", "arguments": {"version_string": "v1.2.3-rc1"}}` |

### Validate

**Goal:** Confirm `"1.2.3"` is valid, `"not-a-version"` is not.

| Layer | Approach |
|-------|----------|
| SDK | `v.Validate() != nil` → invalid |
| CLI | `versions validate 1.2.3` (exit 0 = valid, exit 1 = invalid) |
| MCP | `{"tool": "version_validate", "arguments": {"version_string": "1.2.3"}}` |

### Check type

**Goal:** Determine if `"1.0.0-beta2"` is a beta prerelease.

| Layer | Approach |
|-------|----------|
| SDK | `v.IsBeta()`, `v.IsPrerelease()`, `v.IsStable()`, `v.SubVersion()` |
| CLI | `versions check --beta 1.0.0-beta2` (exit 0 = match) |
| MCP | `{"tool": "version_info", ...}` returns all Is* flags |

### Extract segments

**Goal:** Get `[1, 2, 3]` from `"1.2.3"`.

| Layer | Approach |
|-------|----------|
| SDK | `v.Major()`, `v.Minor()`, `v.Patch()`, `v.Segments()` |
| CLI | `versions segments 1.2.3` |
| MCP | `{"tool": "version_parse", ...}` → `segments` field |

### Get core (strip suffix)

**Goal:** `"v1.2.3-beta1"` → `"v1.2.3"`.

| Layer | Approach |
|-------|----------|
| SDK | `v.Core().RawString()` |
| CLI | `versions core v1.2.3-beta1` |
| MCP | `{"tool": "version_core", "arguments": {"version_string": "v1.2.3-beta1"}}` |

### Custom delimiters

**Goal:** Parse `"curl-7_85_0"` with `-` and `_` delimiters.

| Layer | Approach |
|-------|----------|
| SDK | `versions.NewVersionWithOption(s, versions.ParserOption{Delimiters: ".-_"})` |
| CLI | `versions parse --delimiters "_-" curl-7_85_0` |
| MCP | `{"tool": "version_parse", "arguments": {"version_string": "curl-7_85_0", "delimiters": ".-_"}}` |

## API Reference

### SDK — Constructor Functions

```go
versions.NewVersion(versionStr string) *Version           // never nil, check IsValid()
versions.NewVersionE(versionStr string) (*Version, error)  // returns error for invalid
versions.MustParse(versionStr string) *Version             // panics on invalid — test data only
versions.NewVersions(strings ...string) []*Version         // batch parse
versions.NewVersionWithOption(s string, opt ParserOption) *Version  // custom delimiters
versions.NewVersionStringParser(s string) *VersionStringParser     // low-level parser
```

### SDK — Version Struct Fields

```go
type Version struct {
    Prefix         string    // e.g. "v"
    VersionNumbers []int     // e.g. [1, 2, 3]
    Suffix         string    // e.g. "-rc1"
    Metadata       string    // semver build metadata (after +)
    PublicTime     time.Time // optional release timestamp
    Raw            string    // original input string
}
```

### SDK — Type Checks

```go
v.IsValid() bool       // has VersionNumbers
v.IsStable() bool      // no suffix
v.IsPrerelease() bool  // has any suffix
v.IsAlpha() bool       // suffix contains "alpha"
v.IsBeta() bool        // suffix contains "beta"
v.IsRC() bool          // suffix contains "rc"
v.IsDev() bool         // suffix contains "dev"
v.IsSnapshot() bool    // suffix contains "snapshot"
v.IsNightly() bool     // suffix contains "nightly"
v.IsMilestone() bool   // suffix contains "milestone" or "m"
v.IsFinal() bool       // suffix contains "final"
v.IsGA() bool          // suffix contains "ga"
v.IsPre() bool         // suffix contains "-pre"
v.IsRelease() bool     // suffix contains "-release"
v.IsSP() bool          // suffix contains "sp"
v.IsPost() bool        // suffix contains "post"
```

### SDK — Segment Accessors

```go
v.Major() int           // VersionNumbers[0] or 0
v.Minor() int           // VersionNumbers[1] or 0
v.Patch() int           // VersionNumbers[2] or 0
v.Segments() []int      // all VersionNumbers
v.Segments64() []int64  // as int64
v.SubVersion() int      // numeric from suffix (e.g. "beta2" → 2)
v.SuffixWeight() SuffixWeight  // semantic ordering weight
```

### SDK — Core, Clone, Serialization

```go
v.Core() *Version                    // strip suffix
v.Clone() *Version                   // deep copy
v.BuildGroupID() string              // e.g. "1.2.3"
v.WithPrefix(p string) *Version      // immutable — returns new Version
v.WithSuffix(s string) *Version
v.WithMajor(n int) *Version
v.WithMinor(n int) *Version
v.WithPatch(n int) *Version
v.WithNumbers(ns []int) *Version
v.WithPublicTime(t time.Time) *Version
v.WithMetadata(m string) *Version
// JSON/Text/SQL serialization via MarshalText/UnmarshalText/MarshalJSON/UnmarshalJSON/Scan/Value
```

### SDK — Parser Options

```go
type ParserOption struct {
    Delimiters string  // custom delimiter set, e.g. ".-_" (default: ".-")
}
```

### CLI Commands

```bash
versions parse <version>              # structured JSON output
versions parse --delimiters "_-" <v>  # custom delimiters
versions validate <version>           # exit 0 = valid, 1 = invalid
versions info <version>               # all Is* flags + segments
versions segments <version>           # [major, minor, patch, ...]
versions sub-version <version>        # numeric suffix index
versions suffix-weight <version>      # semantic weight int
versions pure-prefix <version>        # prefix without trailing delimiters
versions group-id <version>           # e.g. "v1.2.3-beta" → "1.2.3"
versions core <version>               # strip suffix
versions clone <version>              # deep copy as JSON
versions check --<type> <version>     # --alpha, --beta, --rc, --stable, etc.
versions check --prerelease <v>       # exit 0 if prerelease
versions check --is-valid <v>         # exit 0 if valid
```

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_parse` | `version_string`, `delimiters?` | prefix, numbers, suffix, metadata |
| `version_validate` | `version_string` | `{valid: bool, error: string?}` |
| `version_info` | `version_string` | all Is* flags, segments, suffix info |
| `version_core` | `version_string` | core version string |

## Cross-References

- [[version-check]] — boolean type checks, comparison predicates
- [[version-comparison]] — CompareTo, IsNewerThan, IsOlderThan
- [[version-sorting]] — sorting parsed versions
- [[version-constraints]] — constraint expression matching
- [[version-mutation]] — bumping, building, modifying versions

## Important Notes

- `NewVersion()` **never returns nil** — always check `IsValid()` for invalid input
- `IsValid()` checks for non-empty VersionNumbers; `Validate()` is stricter (rejects negative numbers)
- `IsPrerelease()` means "has any suffix"; `IsPre()` means "has explicit `-pre` suffix"
- `IsStable()` means "no suffix"; `IsRelease()` means "has explicit `-release` suffix"
- Semver build metadata (`+` part) is in `Metadata` field, **not** Suffix
- All `With*` methods return **new** Version objects — the original is never modified
- `MustParse` **panics** on invalid input — use only for hardcoded/test data
