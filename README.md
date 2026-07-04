# Versions-Skills

<div align="center">

[![Go Tests](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/versions-skills)](https://goreportcard.com/report/github.com/scagogogo/versions-skills)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/versions-skills.svg)](https://pkg.go.dev/github.com/scagogogo/versions-skills)
[![GitHub Release](https://img.shields.io/github/v/release/scagogogo/versions-skills?include_prereleases)](https://github.com/scagogogo/versions-skills/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**An AI-native toolkit for version numbers — parse, compare, sort, group, and check constraints in Go**

📚 **[Full documentation →](https://scagogogo.github.io/versions-skills/)** · [简体中文](./README_CN.md)

</div>

---

Built for the agentic era: every capability is reachable through 🤖 **Claude Code Skills** and 🔌 **MCP** so that AI agents can call them directly, with a 📦 **Go SDK** and 💻 **CLI** underneath. Works natively with **Claude Code**, **Codex**, **Cursor**, **Windsurf**, **Cline**, and **VS Code Copilot**.

## Features

- 🔄 **Wide version support** — standard semver (`1.2.3`), prefixed (`v1.2.3`), prerelease (`1.2.3-beta1`), and custom formats
- 🧩 **Flexible parsing** — auto-detects prefix, numeric segments, suffix, and metadata; configurable delimiters
- 📊 **Semantic comparison** — suffix-weight ordering (`dev < alpha < beta < rc < stable`), not lexicographic
- 📋 **npm-style constraints** — `>=`, `^`, `~`, `1.x`, `||` combinations, three-level grammar
- 🌳 **Grouping & visualization** — group by major/minor, Unicode tree rendering of version hierarchies
- 🚀 **Zero dependencies** — core library has no external dependencies; immutable operations

## Quick Start

```bash
go get github.com/scagogogo/versions-skills
```

```go
v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("v1.3.0-beta")

v1.IsOlderThan(v2)      // true
v2.IsPrerelease()       // true
v2.PreReleaseType()     // "beta"
v1.Diff(v2).IsUpgrade() // true
```

## Documentation

**👉 [https://scagogogo.github.io/versions-skills/](https://scagogogo.github.io/versions-skills/)**

The website covers everything: concepts, tutorials, recipes, the algorithm reference, and full API docs for the Go SDK / CLI / MCP / Skills.

## License

[MIT License](./LICENSE) — Copyright © 2023-2026 scagogogo
