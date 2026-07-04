# Versions-Skills

<div align="center">

[![Go Tests](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml/badge.svg)](https://github.com/scagogogo/versions-skills/actions/workflows/go-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/versions-skills)](https://goreportcard.com/report/github.com/scagogogo/versions-skills)
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/versions-skills.svg)](https://pkg.go.dev/github.com/scagogogo/versions-skills)
[![GitHub Release](https://img.shields.io/github/v/release/scagogogo/versions-skills?include_prereleases)](https://github.com/scagogogo/versions-skills/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**面向 AI 原生的版本号工具集 — 用 Go 解析、比较、排序、分组和约束检查版本号**

📚 **[完整文档 →](https://scagogogo.github.io/versions-skills/)** · [English](./README.md)

</div>

---

为 Agent 时代而生：所有能力都可通过 🤖 **Claude Code Skills** 和 🔌 **MCP** 直接被 AI Agent 调用，底层另附 📦 **Go SDK** 和 💻 **CLI**。原生兼容 **Claude Code**、**Codex**、**Cursor**、**Windsurf**、**Cline**、**VS Code Copilot** 及所有 MCP 兼容的 AI Agent。

## 功能特性

- 🔄 **全面的版本号支持** — 标准语义化版本（`1.2.3`）、带前缀（`v1.2.3`）、预发布（`1.2.3-beta1`）及自定义格式
- 🧩 **灵活的解析** — 自动识别前缀、数字部分、后缀和元数据，支持自定义分隔符
- 📊 **语义化比较** — 基于后缀权重排序（`dev < alpha < beta < rc < stable`），非字典序
- 📋 **npm 风格约束** — `>=`、`^`、`~`、`1.x`、`||` 组合约束，三层语法
- 🌳 **分组与可视化** — 按主/次版本号分组，Unicode 树形展示版本层次
- 🚀 **零依赖** — 核心库无外部依赖；不可变操作

## 快速开始

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

## 文档

**👉 [https://scagogogo.github.io/versions-skills/](https://scagogogo.github.io/versions-skills/)**

官网涵盖全部内容：核心概念、教程、实用配方、算法详解，以及 Go SDK / CLI / MCP / Skills 的完整 API 文档。

## License

[MIT License](./LICENSE) — Copyright © 2023-2026 scagogogo
