---
name: version-parsing
description: Parse, validate, and extract structured components from version strings via SDK, CLI, or MCP. Covers NewVersion/MustParse, all Is* type checks, Segments, Core, Clone, SuffixWeight, serialization, database scanning, and custom parser options.
argument-hint: <version-string-or-task>
---

# Version Parsing Skill

## When to Use

- Parse a version string (e.g. "1.2.3", "v1.2.3-beta1") into structured components (prefix, numbers, suffix, metadata)
- Validate whether a string is a valid version number
- Check version type: prerelease, stable, dev, alpha, beta, RC, snapshot, milestone, nightly, final, GA, pre, release, SP, post
- Extract segments, sub-version, suffix weight, group ID, or core version
- Clone or immutably modify a version (WithPrefix, WithSuffix, WithMajor, etc.)
- Serialize/deserialize versions (JSON, text, SQL)
- Use custom delimiters for non-standard version formats (e.g. underscore-separated)

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

> Replace `{VERSION}` with the latest release tag. See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all available platforms and the current version.

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

## Quick Start

### SDK (Go)

```go
v := versions.NewVersion("v1.2.3-rc1")
fmt.Println(v.Prefix)         // "v"
fmt.Println(v.VersionNumbers) // [1 2 3]
fmt.Println(v.Suffix)         // "-rc1"
fmt.Println(v.IsValid())      // true
```

### CLI

```bash
versions parse v1.2.3-rc1
versions validate 1.2.3
versions info v1.2.3-beta1
```

### MCP

```json
{"tool": "version_parse", "arguments": {"version_string": "v1.2.3-rc1"}}
{"tool": "version_validate", "arguments": {"version_string": "1.2.3"}}
{"tool": "version_info", "arguments": {"version_string": "v1.2.3-beta1"}}
```

## API Reference -- SDK

### Package

```go
import "github.com/scagogogo/versions-skills"
```

### Constructor Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| **NewVersion** | `(versionStr string) *Version` | Parse a version string. Never returns nil; check `IsValid()` for invalid inputs. |
| **NewVersionE** | `(versionStr string) (*Version, error)` | Parse with error return. Returns `ErrVersionInvalid` for invalid strings. |
| **MustParse** | `(versionStr string) *Version` | Parse or panic. Like `regexp.MustCompile`, for hardcoded/test data. |
| **NewVersions** | `(versionStringSlice ...string) []*Version` | Batch-parse multiple version strings. |
| **NewVersionWithOption** | `(versionStr string, option ParserOption) *Version` | Parse with custom options (e.g. custom delimiters). |

### ParserOption

```go
type ParserOption struct {
    Delimiters string // Number separators. Default: "." -- extend to ".-_" for RPM/Python formats.
}
```

```go
// Parse underscore-separated version (Python/RPM style)
v := versions.NewVersionWithOption("1_2_3", versions.ParserOption{Delimiters: ".-_"})
```

Use `versions.DefaultParserOption()` to get the default (Delimiters: ".").

### Version Struct

```go
type Version struct {
    Raw            string         // Original string, e.g. "v1.2.3-beta1"
    PublicTime     time.Time      // Release time
    VersionNumbers VersionNumbers // Number part, e.g. [1,2,3]
    Prefix         VersionPrefix  // Prefix, e.g. "v"
    Suffix         VersionSuffix  // Suffix, e.g. "-beta1"
    Metadata       string         // Semver build metadata (after +), e.g. "build123"
}
```

### Validation Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| **IsValid** | `() bool` | True if VersionNumbers is non-empty (has at least one number). |
| **Validate** | `() error` | Strict validation: numbers must be non-empty and all >= 0. Returns `ErrVersionInvalid` or a negative-number error. |
| **IsZero** | `() bool` | True if the Version is the uninitialized zero-value struct (all fields default). Different from `IsValid()`. |

### Segment Accessors

| Method | Signature | Description |
|--------|-----------|-------------|
| **Major** | `() int` | First number. Returns 0 if empty. |
| **Minor** | `() int` | Second number. Returns 0 if fewer than 2 segments. |
| **Patch** | `() int` | Third number. Returns 0 if fewer than 3 segments. |
| **Segments** | `() []int` | All numbers as `[]int`. Safe copy. |
| **Segments64** | `() []int64` | All numbers as `[]int64`. For large values. |

### Suffix & Type Checks

| Method | Signature | Description |
|--------|-----------|-------------|
| **IsPrerelease** | `() bool` | Has any suffix (alpha, beta, rc, snapshot, etc.). |
| **IsStable** | `() bool` | No suffix -- a stable/release version. |
| **IsDev** | `() bool` | Suffix matches dev pattern (e.g. "-dev1"). |
| **IsAlpha** | `() bool` | Suffix matches alpha/a pattern (e.g. "-alpha1", "-a2"). |
| **IsBeta** | `() bool` | Suffix matches beta/b pattern (e.g. "-beta2", "-b3"). |
| **IsRC** | `() bool` | Suffix matches rc or cr pattern (e.g. "-rc1", "-cr2"). |
| **IsSnapshot** | `() bool` | Suffix matches snapshot pattern. |
| **IsMilestone** | `() bool` | Suffix matches milestone/m pattern (e.g. "-m1", "-milestone1"). |
| **IsNightly** | `() bool` | Suffix matches nightly pattern. |
| **IsFinal** | `() bool` | Explicit "-final" suffix (Maven ecosystem). |
| **IsGA** | `() bool` | Explicit "-ga" suffix. |
| **IsPre** | `() bool` | Explicit "-pre" suffix. Not the same as IsPrerelease(). |
| **IsRelease** | `() bool` | Explicit "-release" suffix. Not the same as IsStable(). |
| **IsSP** | `() bool` | Service pack suffix (e.g. "-sp1"). |
| **IsPost** | `() bool` | Post-release suffix (PEP 440, e.g. "-post1"). |

### Suffix Analysis

| Method | Signature | Description |
|--------|-----------|-------------|
| **SuffixWeight** | `() SuffixWeight` | Semantic weight of the suffix. Equals `GetSuffixWeight(string(v.Suffix))`. |
| **SubVersion** | `() int` | Numeric part of the suffix. E.g. "-beta2" returns 2, "-rc1" returns 1. Returns 0 if none. |

SuffixWeight constants (ascending): `dev(50)` < `snapshot(60)` < `nightly(70)` < `alpha(100)` < `beta(200)` < `milestone(300)` < `rc(400)` < `pre(410)` < `cr(420)` < `final/release/ga(500)` < `sp(600)` < `patch(700)` < `post(800)`.

### Core & Group

| Method | Signature | Description |
|--------|-----------|-------------|
| **Core** | `() *Version` | New Version with suffix removed. E.g. "1.2.3-beta1" -> "1.2.3". |
| **BuildGroupID** | `() string` | Dot-joined numbers. E.g. "1.2.3". Used for grouping. |
| **RawString** | `() string` | The original input string. E.g. "v1.2.3-beta1". Different from `String()` which returns JSON. |

### Clone & Immutable Modification

| Method | Signature | Description |
|--------|-----------|-------------|
| **Clone** | `() *Version` | Deep copy. Modifications to the clone do not affect the original. |
| **WithPrefix** | `(prefix string) *Version` | New Version with replaced prefix. |
| **WithSuffix** | `(suffix string) *Version` | New Version with replaced suffix. |
| **WithMajor** | `(major int) *Version` | New Version with replaced major number. |
| **WithMinor** | `(minor int) *Version` | New Version with replaced minor number. |
| **WithPatch** | `(patch int) *Version` | New Version with replaced patch number. |
| **WithNumbers** | `(numbers []int) *Version` | New Version with replaced number array. |
| **WithPublicTime** | `(t time.Time) *Version` | New Version with replaced release time. |

All `With*` methods return a new `*Version` -- the original is never modified. The `Metadata` field is preserved.

### Comparison Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| **CompareTo** | `(target *Version) int` | -1 (older), 0 (equal), 1 (newer). Compares numbers, then suffix, then time, then raw string. |
| **IsNewerThan** | `(target *Version) bool` | `CompareTo(target) > 0`. |
| **IsOlderThan** | `(target *Version) bool` | `CompareTo(target) < 0`. |
| **Equals** | `(target *Version) bool` | `CompareTo(target) == 0`. |
| **IsBetween** | `(low, high *Version) bool` | `low <= v <= high`. Nil bounds are ignored. |

### Constraint Methods

| Method | Signature | Description |
|--------|-----------|-------------|
| **Satisfies** | `(constraint *Constraint) bool` | Version-centric: `v.Satisfies(c)` == `c.Match(v)`. |
| **Matches** | `(expr string) (bool, error)` | Parse expression + match in one call. E.g. `v.Matches(">=1.0.0,<2.0.0")`. |

### Serialization Interfaces

| Method | Interface | Behavior |
|--------|-----------|----------|
| **MarshalJSON** | `json.Marshaler` | Encodes as JSON string: `"v1.2.3-beta1"`. |
| **UnmarshalJSON** | `json.Unmarshaler` | Decodes from JSON string. Returns `ErrVersionInvalid` for invalid strings. |
| **MarshalText** | `encoding.TextMarshaler` | Encodes as raw bytes of the version string. Works with toml/yaml. |
| **UnmarshalText** | `encoding.TextUnmarshaler` | Decodes from bytes. Returns `ErrVersionInvalid` for invalid strings. |
| **Scan** | `sql.Scanner` | Reads from database (string or []byte). Returns `ErrVersionInvalid` for invalid values. |
| **Value** | `driver.Valuer` | Returns raw string for database storage. |
| **String** | `fmt.Stringer` | Returns full JSON representation of the struct. |

### VersionPrefix Type

```go
type VersionPrefix string
```

Methods: `IsEmpty()`, `String()`, `PurePrefix()` (strips trailing delimiters like `-` or `.`).

### VersionSuffix Type

```go
type VersionSuffix string
```

Methods: `IsEmpty()`, `String()`, `CompareTo(VersionSuffix) int`.

## CLI Commands

### parse

```bash
versions parse v1.2.3-beta1
versions parse --delimiters "_-" curl-7_85_0
```

Shows prefix, numbers, suffix, suffix weight, metadata. The `--delimiters` flag sets custom number separators (default: `.`).

### validate

```bash
versions validate 1.2.3         # exit 0 (valid)
versions validate not-a-version # exit 1 (invalid)
```

Runs strict validation (IsValid + Validate). Exit code 0 = valid, 1 = invalid.

### info

```bash
versions info v1.2.3-beta1
```

Displays all version details including every Is* type check result (IsPrerelease, IsStable, IsDev, IsAlpha, IsBeta, IsRC, IsSnapshot, IsMilestone, IsNightly, IsFinal, IsGA, IsPre, IsRelease, IsSP, IsPost, IsZero).

### segments

```bash
versions segments 1.2.3       # [1, 2, 3]
versions segments v1.2.3.4   # [1, 2, 3, 4]
```

### sub-version

```bash
versions sub-version 1.2.3-beta2  # 2
versions sub-version 1.2.3        # 0
```

### suffix-weight

```bash
versions suffix-weight 1.2.3-beta1  # beta (200)
versions suffix-weight 1.2.3        # unknown (0)
```

### pure-prefix

```bash
versions pure-prefix v1.2.3       # v
versions pure-prefix curl-7.85.0  # curl
```

### group-id

```bash
versions group-id v1.2.3-beta1  # 1.2.3
versions group-id 2.0.0         # 2.0.0
```

### core

```bash
versions core v1.2.3-beta1  # v1.2.3
versions core 2.0.0-rc1     # 2.0.0
```

### check

```bash
versions check --prerelease 1.0.0-beta
versions check --stable 1.2.3
versions check --beta v1.2.3-beta1
versions check --dev 1.0.0-dev1
versions check --alpha 1.0.0-a2
versions check --rc 1.0.0-rc1
versions check --snapshot 1.0.0-SNAPSHOT
versions check --milestone 1.0.0-m1
versions check --nightly 1.0.0-nightly
versions check --final 1.0.0-final
versions check --ga 1.0.0-ga
versions check --pre 1.0.0-pre1
versions check --release 1.0.0-release
versions check --sp 1.0.0-sp1
versions check --post 1.0.0-post1
versions check --zero ""
versions check --newer 1.0.0 2.0.0
versions check --older 2.0.0 1.0.0
versions check --equal 1.0.0 1.0.0
versions check --between-low 1.0.0 --between-high 3.0.0 2.0.0
```

Returns JSON with `result` (bool) and `description`. Exit code 0 = true, 1 = false.

### clone

```bash
versions clone v1.2.3-beta1
```

Returns a deep copy of the version.

## MCP Tools

### version_parse

Parse a version string into structured components.

```json
{"tool": "version_parse", "arguments": {"version_string": "v1.2.3-beta1"}}
```

Returns: `raw`, `prefix`, `numbers`, `suffix`, `suffix_weight`, `metadata`.

### version_validate

Validate a version string (strict: IsValid + Validate).

```json
{"tool": "version_validate", "arguments": {"version_string": "1.2.3"}}
```

Returns: `raw`, `valid`, and `error` if validation fails.

### version_info

Get complete version information including all type checks.

```json
{"tool": "version_info", "arguments": {"version_string": "v1.2.3-beta1"}}
```

Returns: all structured fields plus every Is* boolean (IsPrerelease, IsStable, IsDev, IsAlpha, IsBeta, IsRC, IsSnapshot, IsMilestone, IsNightly, IsFinal, IsGA, IsPre, IsRelease, IsSP, IsPost, IsZero).

## Code Examples (SDK)

### Basic Parsing

```go
v := versions.NewVersion("v1.2.3-rc1")
fmt.Println(v.Prefix)         // "v"
fmt.Println(v.VersionNumbers) // [1 2 3]
fmt.Println(v.Suffix)         // "-rc1"
fmt.Println(v.Metadata)       // "" (no build metadata)
```

### Error-Checking Parse

```go
v, err := versions.NewVersionE("not-a-version")
if err != nil {
    fmt.Println("Error:", err) // version invalid
}
```

### MustParse for Known-Good Values

```go
v := versions.MustParse("1.2.3") // panics on invalid input
fmt.Println(v.Major(), v.Minor(), v.Patch()) // 1 2 3
```

### Strict Validation

```go
v := versions.NewVersion("1.2.3")
if err := v.Validate(); err != nil {
    fmt.Println("Invalid:", err)
}
```

### Segment Access

```go
v := versions.NewVersion("1.2.3.4")
fmt.Println(v.Segments())   // [1 2 3 4]
fmt.Println(v.Segments64()) // [1 2 3 4] as int64
fmt.Println(v.Major())      // 1
fmt.Println(v.Minor())      // 2
fmt.Println(v.Patch())      // 3
```

### Type Checks

```go
v := versions.NewVersion("1.0.0-beta2")
fmt.Println(v.IsPrerelease()) // true
fmt.Println(v.IsStable())     // false
fmt.Println(v.IsBeta())       // true
fmt.Println(v.IsRC())         // false
fmt.Println(v.SubVersion())   // 2
```

### SuffixWeight

```go
v := versions.NewVersion("1.0.0-alpha1")
w := v.SuffixWeight()
fmt.Println(w)                         // 100
fmt.Println(w == versions.SuffixWeightAlpha) // true
fmt.Println(w.String())                // "alpha"
```

### Core and GroupID

```go
v := versions.NewVersion("v1.2.3-beta1")
core := v.Core()
fmt.Println(core.RawString()) // "v1.2.3"
fmt.Println(v.BuildGroupID()) // "1.2.3"
```

### Clone and Immutable Modification

```go
v := versions.NewVersion("1.2.3")
cloned := v.Clone()
newV := v.WithPrefix("v")
fmt.Println(v.RawString())     // "1.2.3"  (unchanged)
fmt.Println(cloned.RawString()) // "1.2.3"
fmt.Println(newV.RawString())   // "v1.2.3"

patched := v.WithPatch(5)
fmt.Println(patched.RawString()) // "1.2.5"
```

### Custom Delimiters

```go
v := versions.NewVersionWithOption("1_2_3", versions.ParserOption{Delimiters: ".-_"})
fmt.Println(v.Segments()) // [1 2 3]
```

### Semver Build Metadata

```go
v := versions.NewVersion("1.0.0+build123")
fmt.Println(v.Metadata)       // "build123"
fmt.Println(v.Segments())     // [1 0 0]
fmt.Println(v.Suffix.IsEmpty()) // true (metadata is not a suffix)
```

### JSON and Database Serialization

```go
// JSON
v := versions.MustParse("1.2.3")
data, _ := json.Marshal(v)          // "1.2.3" (string, not object)
var v2 versions.Version
json.Unmarshal(data, &v2)           // v2.RawString() == "1.2.3"

// SQL
var v3 versions.Version
rows.Scan(&v3)                       // reads from DB string/[]byte
value, _ := v3.Value()               // returns "1.2.3" for storage
```

### Text Serialization (toml/yaml)

```go
v := versions.MustParse("1.2.3")
data, _ := v.MarshalText()          // []byte("1.2.3")
var v4 versions.Version
v4.UnmarshalText(data)              // v4.RawString() == "1.2.3"
```

### Constraint Matching

```go
v := versions.MustParse("1.5.0")
ok, _ := v.Matches(">=1.0.0,<2.0.0")
fmt.Println(ok) // true
```

## Important Notes

- `NewVersion` never returns nil -- always check `IsValid()` for invalid inputs
- `IsValid()` checks for non-empty VersionNumbers; `Validate()` is stricter (also rejects negative numbers)
- `IsZero()` checks if the struct is completely uninitialized -- different from `IsValid()`
- `IsPrerelease()` means "has any suffix"; `IsPre()` means "has explicit `-pre` suffix" -- they are different
- `IsStable()` means "no suffix"; `IsRelease()` means "has explicit `-release` suffix" -- they are different
- Supports arbitrary number of segments: "1.2.3.4.5" works
- Leading zeros are stripped: "1.02.003" becomes [1,2,3]
- Pure alphabetic strings like "abc" return an invalid Version (empty VersionNumbers)
- VersionPrefix is a string type -- call `IsEmpty()` to check for no prefix, `PurePrefix()` to strip trailing delimiters
- VersionSuffix is a string type -- call `IsEmpty()` to check for no suffix
- Semver build metadata (after `+`) is stored in the `Metadata` field, not in Suffix
- `String()` returns JSON; `RawString()` returns the original input string
- All `With*` methods are immutable -- they return new Version objects, the original is never modified
- `MustParse` panics on invalid input -- use only for hardcoded/test data
