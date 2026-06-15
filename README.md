# Versions-Skills

<div align="center">

[![Go Tests](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/versions-skills)](https://goreportcard.com/report/github.com/scagogogo/versions-skills)
[![GoDoc](https://godoc.org/github.com/scagogogo/versions-skills?status.svg)](https://godoc.org/github.com/scagogogo/versions-skills)
[![GitHub Release](https://img.shields.io/github/v/release/scagogogo/versions-skills?include_prereleases)](https://github.com/scagogogo/versions-skills/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**A powerful version number parsing, comparison, sorting, grouping, and constraint checking library for Go**

Accessible via 🤖 **Skills** · 📦 **Go SDK** · 💻 **CLI** · 🔌 **MCP Server**

[English](#features) · [简体中文](#功能特性)

</div>

---

## Access Methods

### 🤖 Skills (Claude Code) — Recommended for AI-powered workflows

One-click install for [Claude Code](https://claude.ai/code):

```bash
claude marketplace add versions https://github.com/scagogogo/versions-skills
```

Then use slash commands in Claude Code: `/version-parsing`, `/version-comparison`, `/version-sorting`, `/version-grouping`, `/version-constraints`, `/version-range-query`, `/version-visualization`, `/version-file-operations`, `/version-check`, `/version-mutation`, `/version-properties`, `/cli-operations`, `/mcp-operations`

### 📦 Go SDK — Recommended for Go developers

```bash
go get github.com/scagogogo/versions-skills
```

```go
import "github.com/scagogogo/versions-skills"

v := versions.NewVersion("v1.2.3-beta1")
fmt.Println(v.Major())    // 1
fmt.Println(v.IsValid())  // true
```

### 💻 CLI — Recommended for scripts and CI/CD

```bash
# Download from GitHub Releases (recommended)
# https://github.com/scagogogo/versions-skills/releases/latest

# Or install via Go
go install github.com/scagogogo/versions-skills/cmd/versions@latest

# Usage
versions parse v1.2.3-beta1
versions compare 1.0.0 2.0.0
versions sort 3.0.0 1.0.0 2.0.0
```

### 🔌 MCP Server — Recommended for AI tool integration

```bash
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

Configure in Claude Code `settings.json`:

```json
{
  "mcpServers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}
```

---

## Features

- 🔄 **Comprehensive version support** — Standard semver (`1.2.3`), prefixed (`v1.2.3`), pre-release (`1.2.3-beta1`), and custom formats
- 🧩 **Flexible parsing** — Auto-detect prefix, numbers, suffix, and metadata with customizable delimiters
- 📊 **Version comparison** — Semantic-aware comparison with suffix weight ordering (dev < alpha < beta < rc < stable)
- 📦 **Grouping & sorting** — Group by major/minor version, sort with stable pre-release ordering
- 🔍 **Range queries** — Query versions within ranges with flexible boundary policies
- 📋 **Constraint expressions** — Full npm-style constraints: `>=1.0.0`, `^1.2.3`, `~1.2`, `1.x`, `>=1.0.0,<2.0.0 || >=3.0.0`
- 🏷️ **Semver compliance** — `IsSemver()`, `ValidateSemver()` for strict SemVer 2.0.0 validation
- 📁 **File I/O** — Read/write version lists from files with comment support
- 🌳 **Visualization** — Unicode tree-based version hierarchy display
- 🔧 **Immutable mutations** — `With*` methods and `Bump*` operations that never modify the original
- 🔗 **Serialization** — JSON, Text, SQL Scanner/Valuer out of the box
- 🚀 **Zero dependencies** — Core library has no external dependencies

---

## Quick Start

### Parse & Compare

```go
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("v1.3.0-beta")

v1.IsOlderThan(v2)      // true
v2.IsPrerelease()       // true
v2.PreReleaseType()     // "beta"
v1.Diff(v2).IsUpgrade() // true
```

### Sort & Group

```go
list := versions.NewVersions("2.0.0", "1.0.0", "1.10.0", "1.5.0-beta")

// Sort
sorted := versions.SortVersionSlice(list)
// → [1.0.0, 1.5.0-beta, 1.10.0, 2.0.0]

// Group by major version
groups := versions.GroupByMajor(list)
// → {1: [1.0.0, 1.5.0-beta, 1.10.0], 2: [2.0.0]}
```

### Constraints

```go
v := versions.NewVersion("1.5.0")

// Check single constraint
c, _ := versions.ParseConstraint(">=1.0.0")
v.Satisfies(c)  // true

// Check constraint expression
ok, _ := v.Matches(">=1.0.0,<2.0.0")  // true

// Negate a constraint
neg := versions.NegateConstraint(c)  // <1.0.0
```

### Range Queries

```go
low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")

r := versions.NewClosedRange(low, high)
r.Contains(versions.NewVersion("1.5.0"))  // true
r.Contains(versions.NewVersion("2.1.0"))  // false
```

### Extract from Strings

```go
v := versions.Coerce("program-1.2.3-linux-amd64")
fmt.Println(v.Raw)  // "1.2.3"
```

---

## API Overview

### Core Types

| Type | Description |
|:-----|:------------|
| `Version` | Represents a version with Raw, PublicTime, VersionNumbers, Prefix, Suffix, Metadata |
| `VersionNumbers` | `[]int` — the numeric segments of a version |
| `VersionPrefix` | `string` — the prefix before numbers (e.g. `"v"`) |
| `VersionSuffix` | `string` — the suffix after numbers (e.g. `"-beta1"`) |
| `VersionRange` | First-class version range with open/closed boundary support |
| `VersionDiff` | Structured difference between two versions |
| `VersionGroup` | Groups versions sharing the same numeric prefix |
| `SortedVersionGroups` | Pre-sorted version group collection for efficient range queries |
| `Constraint` | Single version constraint (operator + target version) |
| `ConstraintSet` | AND-combined constraints (e.g. `>=1.0.0,<2.0.0`) |
| `ConstraintUnion` | OR-combined constraint sets (e.g. `>=1.0.0 || >=3.0.0`) |
| `VersionBuilder` | Fluent builder for constructing Version objects |
| `VersionSlice` | `[]*Version` implementing `sort.Interface` with utility methods |
| `SuffixWeight` | Semantic weight enum for suffix ordering |

### Key Functions

| Category | Functions |
|:---------|:----------|
| **Parse** | `NewVersion`, `NewVersionE`, `MustParse`, `NewVersions`, `Coerce`, `CoerceE` |
| **Compare** | `CompareTo`, `IsNewerThan`, `IsOlderThan`, `Equals`, `IsBetween`, `Diff` |
| **Sort** | `SortVersionSlice`, `SortVersionStringSlice`, `VersionSlice.Sort()` |
| **Group** | `Group`, `GroupByMajor`, `GroupByMinor`, `NewSortedVersionGroups` |
| **Filter** | `Filter`, `FilterByConstraint`, `FilterByStable`, `FilterByMajor`, `Unique` |
| **Constraint** | `ParseConstraint`, `ParseConstraintSet`, `ParseConstraintUnion`, `NegateConstraint` |
| **Range** | `NewClosedRange`, `NewOpenRange`, `VersionRange.Contains`, `VersionRange.Filter` |
| **Check** | `IsPrerelease`, `IsStable`, `IsSemver`, `ValidateSemver`, `PreReleaseType` |
| **Mutate** | `BumpMajor`, `BumpMinor`, `BumpPatch`, `WithPrefix`, `WithSuffix`, `WithMajor`, `Increment` |
| **Utils** | `Min`, `Max`, `LatestStable`, `ContainsVersion`, `IndexOf`, `Difference`, `Intersection`, `Union`, `Partition` |
| **File** | `ReadVersionsFromFile`, `WriteVersionsToFile`, `ReadVersionsFromReader` |
| **Visualize** | `VisualizeVersions`, `VisualizeVersionGroups` |
| **Serialize** | `MarshalJSON`, `UnmarshalJSON`, `MarshalText`, `UnmarshalText`, `Scan`, `Value` |

### Version Methods (full list)

```
IsValid, IsZero, IsPrerelease, IsStable, IsDev, IsAlpha, IsBeta, IsRC,
IsSnapshot, IsMilestone, IsNightly, IsFinal, IsGA, IsPre, IsRelease,
IsSP, IsPost, IsSemver, IsNewerThan, IsOlderThan, Equals, IsBetween,
Satisfies, Matches, CompareTo, Major, Minor, Patch, SubVersion,
SuffixWeight, PreReleaseType, BuildGroupID, Segments, Segments64,
Core, Clone, Validate, ValidateSemver, Diff, Hash, Canonical, Format,
Increment, RawString, String, BumpMajor, BumpMinor, BumpPatch,
WithPrefix, WithSuffix, WithMajor, WithMinor, WithPatch,
WithNumbers, WithPublicTime, WithMetadata,
MarshalText, UnmarshalText, MarshalJSON, UnmarshalJSON, Scan, Value
```

---

## CLI Reference

```bash
# Parsing & Validation
versions parse v1.2.3-rc1
versions validate 1.2.3
versions info v1.2.3-beta1

# Comparison & Checks
versions compare 1.0.0 2.0.0
versions check --stable 1.2.3
versions check --beta 1.2.3-beta1
versions check --newer 1.0.0 1.5.0

# Sorting & Filtering
versions sort 3.0.0 1.0.0 2.0.0
versions sort --desc 3.0.0 1.0.0 2.0.0
versions filter --stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0
versions filter --constraint ">=1.0.0,<2.0.0" 0.5.0 1.0.0 1.5.0 2.0.0

# Grouping & Range
versions group 1.0.0 1.1.0 2.0.0
versions range 1.0.0 2.0.0 1.0.0 1.5.0 2.0.0 3.0.0

# Constraints
versions satisfies 1.5.0 ">=1.0.0,<2.0.0"

# Min/Max
versions min 3.0.0 1.0.0 2.0.0
versions max 3.0.0 1.0.0 2.0.0
versions latest-stable 1.0.0-alpha 1.0.0 2.0.0

# Construction & Mutation
versions build --prefix v --major 1 --minor 2 --patch 3
versions bump 1.2.3 --patch
versions core 1.2.3-beta

# File I/O
versions read versions.txt
versions write output.txt 1.0.0 2.0.0 3.0.0

# Visualization
versions visualize 1.0.0 1.1.0 2.0.0 --groups
```

---

## MCP Server

The MCP server provides all SDK capabilities as AI-callable tools:

```bash
# Start MCP server
versions-mcp --transport stdio     # for Claude Code integration
versions-mcp --transport sse --port 8080  # for network access
```

**Available tools:** `version_parse`, `version_validate`, `version_info`, `version_compare`, `version_sort`, `version_filter`, `version_group`, `version_range_query`, `version_constraint_check`, `version_min`, `version_max`, `version_latest_stable`, `version_latest_prerelease`, `version_unique`, `version_set_operation`, `version_build`, `version_bump`, `version_core`, `version_read_file`, `version_write_file`, `version_visualize`

---

## Claude Code Skills

13 specialized skills for version operations in Claude Code:

| Skill | Command | Description |
|:------|:--------|:------------|
| Version Parsing | `/version-parsing` | Parse, validate, extract version components |
| Version Comparison | `/version-comparison` | Compare versions, check ordering |
| Version Sorting | `/version-sorting` | Sort version lists ascending/descending |
| Version Grouping | `/version-grouping` | Group versions by major/minor numbers |
| Version Constraints | `/version-constraints` | Parse and check constraint expressions |
| Version Range Query | `/version-range-query` | Query versions within ranges |
| Version Visualization | `/version-visualization` | Tree-based version hierarchy display |
| Version File Operations | `/version-file-operations` | Read/write version lists from files |
| Version Check | `/version-check` | Boolean type checks (IsBeta, IsStable, etc.) |
| Version Mutation | `/version-mutation` | Bump versions, immutable modifications |
| Version Properties | `/version-properties` | Access segments, suffix weight, prefix |
| CLI Operations | `/cli-operations` | Full CLI command reference |
| MCP Operations | `/mcp-operations` | MCP server setup and tool reference |

---

## Installation

### Skills (Claude Code)

```bash
claude marketplace add versions https://github.com/scagogogo/versions-skills
```

### Go SDK

```bash
go get github.com/scagogogo/versions-skills
```

### CLI Binary

Pre-built binaries for **Linux**, **macOS**, **Windows**, **FreeBSD**, **OpenBSD**, and **NetBSD** on **amd64**, **arm64**, **arm**, **386**, **mips**, **mips64**, **mips64le**, **ppc64**, **ppc64le**, **s390x**, and **riscv64** architectures. Linux packages: **deb**, **rpm**, **apk**.

```bash
# Linux (amd64)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_linux_amd64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# macOS arm64 (Apple Silicon)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_darwin_arm64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# macOS amd64 (Intel)
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_darwin_amd64.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/

# Or install via package manager (Linux only):
# Debian/Ubuntu: dpkg -i versions_{VERSION}_linux_amd64.deb
# RHEL/Fedora:   rpm -i versions_{VERSION}_linux_amd64.rpm
# Alpine:        apk add versions_{VERSION}_linux_amd64.apk

# Or install via Go
go install github.com/scagogogo/versions-skills/cmd/versions@latest
```

> Replace `{VERSION}` with the latest release tag (e.g. `0.2.0`). See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all available platforms and the current version.

### MCP Server

```bash
# Download binary from GitHub Releases
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions-mcp_{VERSION}_linux_amd64.tar.gz | tar xz
chmod +x versions-mcp && sudo mv versions-mcp /usr/local/bin/

# Or install via Go
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

> Replace `{VERSION}` with the latest release tag. See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all platforms.

---

## Performance

- Version parsing: `O(n)` where n is the version string length
- Version comparison: `O(m)` where m is the max numeric segment count
- Version sorting: `O(n log n)` where n is the list length
- Range queries: `O(log n)` via sorted version groups with binary search

---

## License

[MIT License](./LICENSE) — Copyright © 2023-2026 scagogogo

---

## 功能特性

<div align="center">

**一个强大的 Go 语言版本号解析、比较、排序、分组和约束检查库**

通过 🤖 **Skills** · 📦 **Go SDK** · 💻 **CLI** · 🔌 **MCP Server** 接入

</div>

- 🔄 **全面的版本号支持** — 标准语义化版本（`1.2.3`）、带前缀（`v1.2.3`）、预发布（`1.2.3-beta1`）及自定义格式
- 🧩 **灵活的解析** — 自动识别前缀、数字部分、后缀和元数据，支持自定义分隔符
- 📊 **语义化比较** — 基于后缀权重排序（dev < alpha < beta < rc < stable）
- 📦 **分组与排序** — 按主/次版本号分组，支持稳定的预发布版本排序
- 🔍 **范围查询** — 支持灵活的边界包含/排除策略
- 📋 **约束表达式** — 完整的 npm 风格约束：`>=1.0.0`、`^1.2.3`、`~1.2`、`1.x`、`>=1.0.0,<2.0.0 || >=3.0.0`
- 🏷️ **Semver 规范** — `IsSemver()`、`ValidateSemver()` 严格遵循 SemVer 2.0.0
- 📁 **文件支持** — 从文件读取/写入版本号列表，支持注释
- 🌳 **可视化** — Unicode 树形版本层次结构展示
- 🔧 **不可变操作** — `With*` 和 `Bump*` 方法永不修改原始对象
- 🔗 **序列化** — 内置 JSON、Text、SQL Scanner/Valuer 支持
- 🚀 **零依赖** — 核心库无外部依赖

### 接入方式

#### 🤖 Skills（Claude Code）— AI 工作流推荐

一键安装：

```bash
claude marketplace add versions https://github.com/scagogogo/versions-skills
```

安装后在 Claude Code 中使用斜杠命令：`/version-parsing`、`/version-comparison`、`/version-sorting` 等

#### 📦 Go SDK — Go 开发者推荐

```bash
go get github.com/scagogogo/versions-skills
```

#### 💻 CLI — 脚本和 CI/CD 推荐

从 [GitHub Releases](https://github.com/scagogogo/versions-skills/releases/latest) 下载，或：

```bash
go install github.com/scagogogo/versions-skills/cmd/versions@latest
```

#### 🔌 MCP Server — AI 工具集成推荐

```bash
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```
