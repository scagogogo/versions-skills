# CLAUDE.md

This file provides guidance to Claude Code when working in this repository.

## Project Overview

**versions-skills** is a Go library + CLI + MCP server + Claude Code Skills bundle for version number operations. It provides comprehensive version number parsing, comparison, sorting, grouping, range queries, constraint checking, visualization, and file I/O.

## Repository Structure

```
.
├── .claude-plugin/       # Claude Code plugin & marketplace manifests
├── cmd/
│   ├── versions/         # CLI binary entrypoint
│   └── versions-mcp/     # MCP server binary entrypoint
├── docs/superpowers/plans/ # Historical development plans
├── examples/             # Runnable Go examples (00–06)
├── internal/
│   ├── cli/              # CLI command implementations (cobra)
│   ├── mcp/              # MCP server tool implementations
│   └── verfmt/           # Version formatting utilities
├── skills/               # 13 Claude Code Skills (SKILL.md files)
├── test_data/            # Version list fixtures for testing
├── *.go                  # Core library (root package)
└── *.go                  # Tests (sibling _test.go files)
```

## Build & Test

```bash
# Build everything
go build ./...

# Run all tests
go test ./...

# Build CLI binary
go build -o versions ./cmd/versions

# Build MCP server binary
go build -o versions-mcp ./cmd/versions-mcp

# Run specific tests
go test -v -run TestParse ./...
```

## Plugin Validation

```bash
# Validate plugin manifests
claude plugin validate .

# Strict validation (for CI)
claude plugin validate . --strict

# Dry-run release tag creation
claude plugin tag . --dry-run
```

## Code Style

- Standard Go conventions: `gofmt`, `go vet`
- Tests use `github.com/stretchr/testify` for assertions
- Test files are named `*_test.go` alongside the source files
- Doc comments follow Go standard format

## Key Design Decisions

- **Zero dependencies in core**: The root package only depends on `go-tuple`, `go-shuffle`, `go-compare-anything` (all from golang-infrastructure)
- **Immutable operations**: All `With*` and `Bump*` methods return new Version objects
- **Never nil from NewVersion**: Always returns a Version struct; check `IsValid()` for invalid inputs
- **Suffix weight ordering**: dev(50) < snapshot(60) < nightly(70) < alpha(100) < beta(200) < milestone(300) < rc(400) < final/release/ga(500) < sp(600) < patch(700) < post(800)
- **CompareTo priority order**: VersionNumbers → Suffix → PublicTime → Raw string
- **File format**: One version per line, `#` comments, blank lines ignored

## Release Process

1. Ensure all tests pass: `go test ./...`
2. Ensure plugin validation passes: `claude plugin validate . --strict`
3. Tag release: `claude plugin tag --push`
4. goreleaser creates binaries, packages, and GitHub Release via CI

## Skills

The 13 skills in `skills/*/SKILL.md` follow this structure:
- YAML frontmatter with `name`, `description`, `argument-hint`
- Three access paths per skill: SDK (Go), CLI, MCP
- Each skill has: When to Use, Installation, Quick Start, API Reference, Code Examples, Important Notes

When adding or modifying skills:
- Match the existing format and depth
- Cover all three access paths (SDK, CLI, MCP)
- Include code examples for each
- Add an "Important Notes" section with gotchas
