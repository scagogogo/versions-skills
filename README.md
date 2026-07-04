# Versions-Skills

<div align="center">

[![Go Tests](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/versions-skills)](https://goreportcard.com/report/github.com/scagogogo/versions-skills)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/versions-skills.svg)](https://pkg.go.dev/github.com/scagogogo/versions-skills)
[![GitHub Release](https://img.shields.io/github/v/release/scagogogo/versions-skills?include_prereleases)](https://github.com/scagogogo/versions-skills/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**An AI-native toolkit for version numbers — parse, compare, sort, group, and check constraints in Go**

Built for the agentic era: every capability is reachable through 🤖 **Claude Code Skills** and 🔌 **MCP** so that AI agents can call them directly, with a 📦 **Go SDK** and 💻 **CLI** underneath.

Works natively with **Claude Code**, **Codex**, **Cursor**, **Windsurf**, **Cline**, **VS Code Copilot**, and any MCP-compatible AI agent.

[Features](#features) · [AI Agent Integration](#ai-agent-integration-architecture) · [简体中文](./README_CN.md)

</div>

---

## Capability Map

```mermaid
flowchart TB
  subgraph PARSE["解析与验证"]
    P1[NewVersion]
    P2[MustParse]
    P3[Coerce]
    P4[Validate]
  end

  subgraph COMP["比较"]
    C1[CompareTo]
    C2[IsNewerThan]
    C3[Diff]
  end

  subgraph SORT["排序与过滤"]
    S1[SortVersionSlice]
    S2[Filter]
    S3[Unique]
  end

  subgraph GROUP["分组"]
    G1[GroupByMajor]
    G2[GroupByMinor]
    G3[SortedVersionGroups]
  end

  subgraph CONSTR["约束"]
    N1[ParseConstraint]
    N2[Satisfies]
    N3[Matches]
  end

  subgraph RANGE["范围查询"]
    R1[NewClosedRange]
    R2[Contains]
    R3[QueryRange]
  end

  subgraph CHECK["类型判断"]
    K1[IsStable]
    K2[IsPrerelease]
    K3[IsSemver]
  end

  subgraph MUT["变更"]
    M1[BumpMajor]
    M2[WithSuffix]
    M3[VersionBuilder]
  end

  subgraph FILE["文件"]
    F1[ReadFromFile]
    F2[WriteToFile]
  end

  subgraph VIS["可视化"]
    V1[VisualizeVersions]
    V2[VisualizeGroups]
  end

  subgraph SER["序列化"]
    E1[MarshalJSON]
    E2[Scan/Value]
  end

  subgraph SET["集合运算"]
    T1[Min/Max]
    T2[Intersection]
    T3[Partition]
  end
```

*12 functional domains, 3-level hierarchy: Category → Sub-system → Specific API*

---

## Architecture

```mermaid
flowchart TB
  subgraph L1["AI Agent Layer"]
    CC[Claude Code]
    CX[Cursor]
    WS[Windsurf]
    CL[Cline]
    CD[Codex]
  end

  subgraph L2["Access Layer"]
    SK[Skills 13个]
    MCP[MCP Server 21工具]
    SDK[Go SDK]
    CLI[CLI 44命令]
  end

  subgraph L3["Core Library"]
    direction LR
    Parse[Parse]
    Compare[Compare]
    Sort[Sort]
    Group[Group]
    Constrain[Constraint]
    Range[Range]
    Check[TypeCheck]
    Mutate[Mutate]
    FileIO[FileIO]
    Visual[Visualize]
    Serial[Serialize]
    SetOps[SetOps]
  end

  subgraph L4["Foundation"]
    GO[Go Runtime]
    IM[Immutable Design]
    ZD[Zero External Deps]
  end

  L1 --> L2
  L2 --> L3
  L3 --> L4
```

*Four-layer architecture: AI Agent → Interface → Feature → Core Library*

The original ASCII architecture diagram is also available in [中文版架构图](#ai-agent-集成架构) below.

**Two paths for AI agents:**
1. **Skills Plugin** — Claude Code reads `SKILL.md` files as domain knowledge, then calls CLI/MCP/SDK under the hood. Best for guided workflows and one-off tasks.
2. **MCP Server** — Any MCP-compatible client calls `version_*` tools directly. Best for programmatic use, batch operations, and non-Claude agents.

**Use both together** for the best experience — Skills provide the "how-to" knowledge, MCP provides the execution engine.

---

## Data Flow

```mermaid
flowchart LR
  IN[Input<br/>Raw String] --> PARSE[Parse<br/>NewVersion]
  PARSE --> PROC[Process<br/>Compare/Sort/Group]
  PROC --> TRANS[Transform<br/>Filter/Bump/With]
  TRANS --> OUT[Output<br/>Version/Slice/Map]
```

*Version strings flow through Input → Parse → Process → Transform → Output pipeline*

---

## Access Methods

```mermaid
flowchart TB
  CORE[versions-skills<br/>Core Library]

  CORE --> SK[Skills<br/>13 SKILL.md files<br/>Claude Code only]
  CORE --> MCP[MCP Server<br/>21 version_* tools<br/>Any MCP client]
  CORE --> SDK[Go SDK<br/>Type-safe full API<br/>Go programs]
  CORE --> CLI[CLI<br/>44 subcommands<br/>Shell/CI/CD]

  SK --> AI[AI Agents]
  MCP --> AI
  SDK --> APP[Applications]
  CLI --> SCRIPT[Scripts]
```

*Four access methods connecting to 14 core capabilities*

### 🤖 Skills (Claude Code) — Recommended for AI-powered workflows

**One command to install all 13 version skills as slash commands in Claude Code:**

```bash
# Step 1: Add the marketplace (one-time)
claude plugin marketplace add https://github.com/scagogogo/versions-skills

# Step 2: Install the plugin
claude plugin install versions
```

After installation, 13 slash commands are available in any Claude Code session:

| Command | What it does |
|:--------|:-------------|
| `/version-parsing` | Parse, validate, extract version components |
| `/version-comparison` | Compare versions, check ordering |
| `/version-sorting` | Sort version lists ascending/descending |
| `/version-grouping` | Group versions by major/minor numbers |
| `/version-constraints` | Parse and check constraint expressions |
| `/version-range-query` | Query versions within ranges |
| `/version-visualization` | Tree-based version hierarchy display |
| `/version-file-operations` | Read/write version lists from files |
| `/version-check` | Boolean type checks (IsBeta, IsStable, etc.) |
| `/version-mutation` | Bump versions, immutable modifications |
| `/version-properties` | Access segments, suffix weight, prefix |
| `/cli-operations` | Full CLI command reference |
| `/mcp-operations` | MCP server setup and tool reference |

> **How it works:** The plugin ships 13 `SKILL.md` files under `skills/`. Claude Code reads these as domain knowledge — when you type `/version-parsing`, Claude loads the skill's API reference, code examples, and decision tree, then uses the SDK/CLI/MCP to execute your request. No API key or runtime dependency needed; the skill tells Claude how to call the tools you already have installed.

<details>
<summary>📖 Plugin vs MCP Server — which should I use?</summary>

| | **Plugin (Skills)** | **MCP Server** |
|:--|:--|:--|
| **Install** | `claude plugin install versions` | `go install .../versions-mcp@latest` + config |
| **How it works** | Injects domain knowledge as slash commands | Exposes 21 tools as AI-callable functions |
| **Best for** | Guided workflows, learning the API, one-off tasks | Programmatic tool calls, batch operations, other AI agents |
| **Requires** | Claude Code | Any MCP-compatible client |
| **Use both?** | ✅ Yes — they complement each other | ✅ Yes — they complement each other |

</details>

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
# One-line installer (Linux/macOS, auto-detects platform + version)
curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash

# Or download a binary from GitHub Releases:
# https://github.com/scagogogo/versions-skills/releases/latest

# Or install via Go
go install github.com/scagogogo/versions-skills/cmd/versions@latest

# Usage
versions parse v1.2.3-beta1
versions compare 1.0.0 2.0.0
versions sort 3.0.0 1.0.0 2.0.0
```

### 🔌 MCP Server — Recommended for AI tool integration

The MCP server exposes all SDK capabilities as 21 AI-callable tools, compatible with **Claude Code**, **Codex**, **Cursor**, **Windsurf**, **Cline**, **VS Code Copilot**, and any MCP-compatible client.

```bash
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

<details>
<summary>⚙️ Configuration for each AI client</summary>

**Claude Code** — add to `~/.claude/settings.json` (user scope) or `.claude/settings.json` (project scope):

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

**Cursor** — add to `.cursor/mcp.json` in your project root:

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

**Windsurf** — add to `.windsurf/mcp.json` in your project root:

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

**VS Code (Copilot)** — add to `.vscode/mcp.json` in your project root:

```json
{
  "servers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}
```

**Codex (OpenAI CLI)** — register with the `codex mcp` command (writes to `~/.codex/config.toml`):

```bash
codex mcp add versions -- versions-mcp --transport stdio
```

Or edit `~/.codex/config.toml` directly:

```toml
[mcp_servers.versions]
command = "versions-mcp"
args = ["--transport", "stdio"]
```

> In Codex, the server appears as the `versions` MCP server. Start a session with `codex` and the `version_*` tools are available to the agent automatically.

**Cline (VS Code)** — add to `~/.cline/cline_mcp_settings.json` (or via the Cline UI → *MCP Servers → + Add Server*):

```json
{
  "mcpServers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"],
      "disabled": false,
      "autoApprove": []
    }
  }
}
```

**Network mode (SSE)** — for shared/team deployments:

```bash
versions-mcp --transport sse --port 8080
```

Then point any MCP client at `http://localhost:8080/sse` (use the SSE/HTTP transport option in your client's config instead of a `command`).

</details>

> **Don't see your client?** Any tool that speaks MCP can connect — point its stdio command at `versions-mcp --transport stdio`, or its HTTP transport at the SSE endpoint above. See the [AI Agent Integration](#ai-agent-integration-architecture) section for the decision table.

**Available tools:** `version_parse`, `version_validate`, `version_info`, `version_compare`, `version_sort`, `version_filter`, `version_group`, `version_range_query`, `version_constraint_check`, `version_min`, `version_max`, `version_latest_stable`, `version_latest_prerelease`, `version_unique`, `version_set_operation`, `version_build`, `version_bump`, `version_core`, `version_read_file`, `version_write_file`, `version_visualize`

---

## AI Agent Integration Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    AI Agent / IDE                        │
│  (Claude Code · Codex · Cursor · Windsurf · Cline · …) │
├──────────────────────┬──────────────────────────────────┤
│                      │                                  │
│   🤖 Skills Plugin   │   🔌 MCP Server                  │
│   (Claude Code only) │   (any MCP client)               │
│                      │                                  │
│   13 SKILL.md files  │   21 version_* tools             │
│   → slash commands   │   → AI-callable functions         │
│   → domain knowledge │   → structured JSON responses     │
│                      │                                  │
├──────────────────────┴──────────────────────────────────┤
│                                                         │
│   💻 CLI binary          📦 Go SDK                      │
│   (shell/CI/CD)          (Go programs)                  │
│                                                         │
├─────────────────────────────────────────────────────────┤
│              Core Library (Go · zero dependencies)       │
└─────────────────────────────────────────────────────────┘
```

**Two paths for AI agents:**
1. **Skills Plugin** — Claude Code reads `SKILL.md` files as domain knowledge, then calls CLI/MCP/SDK under the hood. Best for guided workflows and one-off tasks.
2. **MCP Server** — Any MCP-compatible client calls `version_*` tools directly. Best for programmatic use, batch operations, and non-Claude agents.

**Use both together** for the best experience — Skills provide the "how-to" knowledge, MCP provides the execution engine.

### Which agent, which path?

| AI Agent / IDE | Skills | MCP (stdio) | MCP (SSE) | Config location |
|:--|:--:|:--:|:--:|:--|
| **Claude Code** | ✅ `claude plugin install versions` | ✅ | ✅ | `~/.claude/settings.json` |
| **Codex** (OpenAI CLI) | — | ✅ | ✅ | `~/.codex/config.toml` (or `codex mcp add`) |
| **Cursor** | — | ✅ | ✅ | `.cursor/mcp.json` |
| **Windsurf** | — | ✅ | ✅ | `.windsurf/mcp.json` |
| **Cline** (VS Code) | — | ✅ | ✅ | `~/.cline/cline_mcp_settings.json` |
| **VS Code Copilot** | — | ✅ | ✅ | `.vscode/mcp.json` |
| *Any MCP client* | — | ✅ | ✅ | client-specific |

> **Recommendation:** On Claude Code, install **both** the Skills plugin and the MCP server — the plugin gives Claude the versioning domain knowledge (when to use constraints vs. ranges, how suffixes sort), while MCP gives it a fast, structured execution channel. On every other agent, the MCP server is the only path you need.

### Quick start by agent

**Claude Code** — get the slash commands *and* the MCP tools:

```bash
# 1. Domain knowledge as 13 slash commands
claude plugin marketplace add https://github.com/scagogogo/versions-skills
claude plugin install versions

# 2. Direct tool execution (optional but recommended)
claude mcp add versions -- versions-mcp --transport stdio
```

Now you can either type `/version-sorting` to be guided through it, or just ask "sort these versions" and Claude will call `version_sort` directly.

**Codex** (OpenAI CLI):

```bash
# Install the binary once
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest

# Register with Codex
codex mcp add versions -- versions-mcp --transport stdio
```

**Cursor / Windsurf / VS Code Copilot / Cline** — install the binary, then add the JSON snippet from the [🔌 MCP Server](#-mcp-server--recommended-for-ai-tool-integration) section to the config file listed in the table above.

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

```mermaid
flowchart LR
  DEV["dev<br/>50"] --> SNAP["snapshot<br/>60"]
  SNAP --> NGHT["nightly<br/>70"]
  NGHT --> A["alpha<br/>100"]
  A --> B["beta<br/>200"]
  B --> M["milestone<br/>300"]
  M --> RC["rc<br/>400"]
  RC --> FINAL["final/release/ga<br/>500"]
  FINAL --> SP["sp<br/>600"]
  SP --> PATCH["patch<br/>700"]
  PATCH --> POST["post<br/>800"]

  style DEV fill:#fef2f2,stroke:#dc2626
  style FINAL fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
  style POST fill:#eff6ff,stroke:#2563eb
```

*Semantic comparison priority: suffix weight determines pre-release ordering*

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

```mermaid
flowchart TD
  U["Union (OR)<br/>分隔符: ||"] -->|任一 Set 命中| S1["Set₁ (AND)<br/>分隔符: ,"]
  U --> S2["Set₂ (AND)"]
  S1 -->|所有 Single 命中| C1["Single<br/>操作符 + 版本"]
  S1 --> C2["Single<br/>操作符 + 版本"]
  S2 --> C3["Single<br/>操作符 + 版本"]
  C1 --> OP["操作符<br/>= != &gt; &lt; &gt;= &lt;= ^ ~ x"]
  OP --> V["目标版本<br/>走 CompareTo"]
  style U fill:#eff6ff,stroke:#2563eb
  style S1 fill:#f0fdf4,stroke:#16a34a
  style S2 fill:#f0fdf4,stroke:#16a34a
```

*Three-level grammar: Union (OR) → Set (AND) → Single Constraint → Operator + Version*

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

## How the Algorithms Work

This section documents the exact semantics an AI agent (or human) needs to predict library behavior without reading source. Every rule below is implemented in the referenced source file.

### 1. Parsing — three segments + metadata (`parser.go`)

Every version string is split into four fields:

```
 v1.2.3-beta1+build.7
 │ └─┬─┘ └──┬──┘ └──┬──┘
 │  │      │       └─ Metadata   (semver build metadata, after "+")
 │  │      └─ Suffix             (everything after the numeric part)
 │  └─ VersionNumbers            (the integer segments, dot-delimited)
 └─ Prefix                       (non-digit lead-in, e.g. "v", "release-")
```

Algorithm, in order:

1. **Trim** whitespace. Empty string → invalid version with empty `VersionNumbers`.
2. **Strip metadata**: find the last `+`. The text after it is metadata **only if it contains no `-`** — so `0.9.0+121-bcc5decc` keeps `121-bcc5decc` as part of the suffix (Scala/Maven style), while `1.2.3+build.7` strips `build.7` into `Metadata`. This disambiguates semver metadata from pre-release identifiers that legitimately contain `-`.
3. **Pure-alpha shortcut**: if the remaining string has no digit at all, the whole thing becomes `Prefix` and `VersionNumbers` is empty (an invalid version, e.g. `"abc"`).
4. **Read prefix**: scan to the first digit; everything before it is the prefix. `v1.2.3` → prefix `"v"`. `.1` → prefix `""`.
5. **Read numbers**: from the first digit, collect digit-runs separated by the configured delimiter (default `.`). Consecutive delimiters collapse; parsing stops at the first non-digit, non-delimiter char. `1.2.48` → `[1,2,48]`.
6. **Read suffix**: everything after the last number-position. A regex cascade handles common patterns (`-snapshot.…`, `-v2.x.x`, `-revN-…`, `+nnn-xxxx`, `-beta1`) before falling back to "remainder of the string."

`Coerce(s)` does the same but first extracts a version-shaped substring out of arbitrary text (`"program-1.2.3-linux-amd64"` → `"1.2.3"`).

> **Delimiters are unifiable**: parsing respects a configurable `Delimiters` option, but `VersionNumbers.BuildGroupID()` always rejoins with `.` — so `"1/2/3"` and `"1.2.3"` produce the same canonical numeric form.

### 2. Comparison — four-level priority (`version.go: CompareTo`)

`a.CompareTo(b)` returns `-1/0/1` by trying these keys in order; the **first that differs wins**:

| # | Key | Rule | Source |
|:--:|:--|:--|:--|
| 1 | **VersionNumbers** | Element-wise `int` compare left→right; if all shared positions tie, the **longer** one is greater (`1.2` < `1.2.0`? — see note). | `version_numbers.go` |
| 2 | **Suffix** | Stable (no suffix) **beats** pre-release (has suffix): `1.0.0` > `1.0.0-rc`. Two suffixes compared by [suffix weight](#3-suffix-weight-ordering-suffix_weightgo). | `version.go` / `version_suffix.go` |
| 3 | **PublicTime** | If both have a non-zero `PublicTime`, later wins. Only consulted when numbers + suffix both tie. | `version.go` |
| 4 | **Raw string** | Final tiebreaker: lexicographic on the original string. | `version.go` |

> **Note on length:** `VersionNumbers.CompareTo` treats `1.2` as **less than** `1.2.0` (shorter is smaller when shared segments tie). In practice this means a 2-segment version sorts below its 3-segment sibling — be explicit with `1.2.0` if you mean the release.

This ordering is why `1.0.0-alpha` < `1.0.0-beta` < `1.0.0-rc1` < `1.0.0` < `1.0.0-post1`, matching real release ladders rather than naive ASCII order.

### 3. Suffix weight ordering (`suffix_weight.go`)

Each suffix is matched (case-insensitive, optional leading `-`/`.`) against an ordered pattern table. The matched **weight** determines its rank; if both suffixes have the same weight, their trailing integer (the "sub-version", e.g. the `1` in `-alpha1`) breaks the tie.

| Weight | Suffix patterns (examples) | Meaning |
|-------:|:--|:--|
| 50 | `dev`, `dev1`, `.dev.2` | Development build |
| 60 | `snapshot`, `snapshot20201012…` | Snapshot |
| 70 | `nightly` | Nightly build |
| 100 | `a`, `alpha`, `alpha1`, `.alpha.2` | Alpha |
| 200 | `b`, `beta`, `beta2` | Beta |
| 300 | `m`, `milestone`, `m1` | Milestone |
| 400 | `rc`, `rc1` | Release candidate |
| 410 | `pre`, `pre1` | Pre-release |
| 420 | `cr`, `cr1` | CR variant of RC |
| 500 | `final`, `release`, `ga` | Release / GA (same as *no suffix*) |
| 600 | `sp`, `sp1` | Service pack |
| 700 | `patch`, `patch1` | Patch |
| 800 | `post`, `post1` | Post-release (PEP 440) |

Rules worth knowing:
- **Unknown suffixes sort after known ones** within the same comparison — a suffix that matches no pattern is treated as less important than any recognized pre-release type but falls back to plain string order among themselves.
- `final`, `release`, `ga` all share weight 500 — identical in sort order to having no suffix at all, but `IsStable()` only returns true when the suffix is *literally empty*. A version `1.0.0-ga` is release-weighted but not "stable" by the empty-suffix test.

### 4. Constraint grammar — three levels (`constraint.go`)

```
Union  (OR)   : ">1.0.0,<2.0.0 || >=3.0.0"     ← split on "||"
  └─ Set (AND) : ">=1.0.0,<2.0.0"              ← split on ","
       └─ Single: ">=1.0.0" | "^1.2.3" | "~1.2" | "1.x" | "1.2.3"
```

`ConstraintUnion.Match(v)` = true if **any** set matches. `ConstraintSet.Match(v)` = true if **all** singles match. A single `Constraint` is `Operator + Version`.

Supported operators and their exact semantics (all comparisons use `Version.CompareTo`, so suffix weights apply):

| Op | Name | `base` | matches `v` when |
|:--:|:--|:--|:--|
| `=` | equal | `1.2.3` | `v == 1.2.3` (bare `1.2.3` means `=`) |
| `!=` | not equal | `1.2.3` | `v != 1.2.3` |
| `>` `<` `>=` `<=` | range comparison | `1.2.3` | straightforward `CompareTo` |
| `^` | caret | `^1.2.3` | `v >= 1.2.3` AND `v < {first-non-zero-segment+1, 0…}` |
| `~` | tilde | `~1.2.3` | `v >= 1.2.3` AND `v < {major, minor+1, 0…}` |
| `x`/`X`/`*` | wildcard | `1.x` | `v >= 1.0.0` AND `v < {last-non-zero+1, 0…}` |

**Caret bounds** (compatibility range — "leftmost non-zero"):
- `^1.2.3` → `>=1.2.3, <2.0.0` (bump the `1`)
- `^0.2.3` → `>=0.2.3, <0.3.0` (first non-zero is the minor)
- `^0.0.3` → `>=0.0.3, <0.0.4` (first non-zero is the patch)

**Tilde bounds** (lock to minor):
- `~1.2.3` → `>=1.2.3, <1.3.0`
- `~1.2`   → `>=1.2.0, <1.3.0` (patch is open)

**Wildcard bounds** (bump the last specified digit):
- `1.x`   → `>=1.0.0, <2.0.0`
- `1.2.x` → `>=1.2.0, <1.3.0`

### 5. Range queries — open/closed boundaries (`version_range.go`)

`VersionRange{Low, High, LowInclusive, HighInclusive}` tests membership with four boundary checks:

- `Low == nil` → no lower bound; `High == nil` → no upper bound.
- On the low side: `v < Low` fails; `v == Low && !LowInclusive` fails (excludes the endpoint of an open interval).
- Symmetric on the high side.

So `NewClosedRange(1.0.0, 2.0.0)` = `[1.0.0, 2.0.0]` (both endpoints in), `NewOpenRange` = `(1.0.0, 2.0.0)` (both out), and you can mix: `NewVersionRange(1.0.0, 2.0.0, true, false)` = `[1.0.0, 2.0.0)`. `IsEmpty()` detects a degenerate range where `Low > High`, or where `Low == High` but at least one side is open.

### 6. Sorting & grouping — two-phase, group-aware (`sort.go`, `version_group.go`)

`SortVersionSlice` is **not** a flat `sort.Slice`. It runs two phases:

1. **Group** all versions by `BuildGroupID()` (the full numeric string, e.g. `1.2.3`).
2. **Sort groups** by their numeric prefix (`VersionGroup.CompareTo`), then **sort within each group** and concatenate.

The payoff: `1.10.0` correctly sorts *after* `1.2.0` (numeric, not string), and versions of the same lineage stay clustered.

**Group granularity — two different APIs, don't confuse them:**

| Function | Groups by | Key type | Example buckets |
|:--|:--|:--|:--|
| `Group(versions)` | **full numeric string** (`BuildGroupID`) | `map[string]*VersionGroup` | `1.2.3` and `1.2.4` land in **different** groups (`"1.2.3"`, `"1.2.4"`) |
| `GroupByMajor(versions)` | **major segment only** (`Major()`) | `map[int][]*Version` | `1.2.3`, `1.2.4`, `1.9.0` all in group `1` |
| `GroupByMinor(versions)` | **major.minor** | `map[string][]*Version` | `1.2.3`, `1.2.4` in group `"1.2"` |

`Group()` is what sorting and range queries use internally; `GroupByMajor`/`GroupByMinor` are convenience buckets for "give me everything in the 1.x line".

### 7. Range queries over a sorted index — `SortedVersionGroups` (`sorted_version_groups.go`)

For repeated range queries over a large version set, build a `SortedVersionGroups` once:

```go
sg := versions.NewSortedVersionGroups(allVersions)   // group + sort + index, O(n log n)
start := tuple.NewTuple2(versions.NewVersion("1.0.0"), versions.ContainsPolicyYes)
end   := tuple.NewTuple2(versions.NewVersion("2.0.0"), versions.ContainsPolicyNo)
hits  := sg.QueryRange(start, end)                    // O(groups) scan with index jump
```

It pre-sorts groups and builds a `groupID → index` map. `QueryRange` jumps straight to the start group via the map (skipping everything below it), then walks groups collecting `QueryRangeVersions` until it passes the end. The `ContainsPolicy` on each tuple decides whether the boundary version itself is included (`Yes`) or excluded (`No`). This is the engine behind `version_range_query` and is far cheaper than re-filtering the whole list per query.

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

## Installation

### Skills (Claude Code Plugin)

```bash
# Add marketplace (one-time)
claude plugin marketplace add https://github.com/scagogogo/versions-skills

# Install the plugin
claude plugin install versions
```

> After installation, 13 slash commands are available in Claude Code. See [🤖 Skills](#-skills-claude-code--recommended-for-ai-powered-workflows) above for the full list.

### MCP Server (for AI Agents)

The MCP server works with **Claude Code**, **Codex**, **Cursor**, **Windsurf**, **Cline**, **VS Code Copilot**, and any MCP-compatible client. See [🔌 MCP Server](#-mcp-server--recommended-for-ai-tool-integration) above for per-client configuration, and the [AI Agent Integration](#ai-agent-integration-architecture) section for a decision table.

```bash
# Download binary from GitHub Releases
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions-mcp_{VERSION}_linux_amd64.tar.gz | tar xz
chmod +x versions-mcp && sudo mv versions-mcp /usr/local/bin/

# Or install via Go
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

> `{VERSION}` is the release tag shown at the top of the [releases page](https://github.com/scagogogo/versions-skills/releases/latest).

### Go SDK

```bash
go get github.com/scagogogo/versions-skills
```

### CLI Binary

Pre-built binaries for **Linux**, **macOS**, **Windows**, **FreeBSD**, **OpenBSD**, and **NetBSD** on **amd64**, **arm64**, **arm**, **386**, **mips**, **mips64**, **mips64le**, **ppc64**, **ppc64le**, **s390x**, and **riscv64** architectures. Linux packages: **deb**, **rpm**, **apk**.

```bash
# One-line installer (auto-detects platform + version):
curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash

# Or download a specific binary manually:
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

> Prefer the one-line `install.sh` above, which resolves `{VERSION}` for you. For manual download, `{VERSION}` is the release tag shown at the top of the [releases page](https://github.com/scagogogo/versions-skills/releases/latest).

---

## Performance

- Version parsing: `O(n)` where n is the version string length
- Version comparison: `O(m)` where m is the max numeric segment count
- Version sorting: `O(n log n)` where n is the list length
- Range queries: `O(log n)` via sorted version groups with binary search

---

## License

[MIT License](./LICENSE) — Copyright © 2023-2026 scagogogo

