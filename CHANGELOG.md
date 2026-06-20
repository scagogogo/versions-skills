# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- README: AI Agent Integration Architecture diagram (ASCII) showing how Skills,
  MCP Server, CLI, and SDK layer together.
- README: Chinese architecture diagram with translated labels.
- README: AI agent compatibility subtitle in header (English + Chinese).
- Skills: `installation/SKILL.md` now includes Claude Code plugin install as Step 0.
- Skills: `mcp-operations/SKILL.md` now includes per-client configuration table
  (Claude Code, Cursor, Windsurf, VS Code Copilot).

### Changed
- README: improved AI Agent integration docs — correct Claude Code plugin install commands
  (two-step marketplace add → plugin install), added multi-client MCP configuration
  (Claude Code, Cursor, Windsurf, VS Code Copilot), added Plugin vs MCP comparison,
  added "How it works" explanation for Skills, removed duplicate sections.
- README: updated Chinese section with matching AI Agent docs improvements.
- Skills: updated `installation/SKILL.md` with multi-client MCP configuration.
- Skills: updated `mcp-operations/SKILL.md` with multi-client configuration
  and per-client config table.
- CONTRIBUTING: updated Go version requirement from 1.23+ to 1.25+.

## [0.2.0] - 2026-06-19

### Added
- golangci-lint to CI and `.golangci.yml` configuration (v2 format).
- Dependabot configuration for Go modules and GitHub Actions.
- Community files: CHANGELOG, CONTRIBUTING, SECURITY, CODE_OF_CONDUCT.
- Expanded `.gitignore` for build artifacts, coverage, and binaries.
- `skills/installation/SKILL.md` — new skill for installation guidance.

### Changed
- Rewrote all 13 skills as AI-agent operational instructions (clearer prompts,
  better examples, consistent structure).
- README badge now points to pkg.go.dev instead of the deprecated godoc.org.
- Surfaces the one-line `install.sh` installer in README and skills.
- CI workflows opt into Node.js 24 for GitHub Actions runners.
- goreleaser now builds `versions-mcp` for all 6 OSes × 12 architectures.

### Fixed
- Removed a dead `make([]int, 0)` allocation in `parser.go` (found by ineffassign).
- Check the error return of `encoder.Encode` in CLI quiet mode (found by errcheck).
- Handle `f.Close()` and `tw.Flush()` error returns explicitly (errcheck).
- Fix `Version` var comment to follow Go conventions (staticcheck ST1022).
- Rename `groupIdToIndexMap` to `groupIDToIndexMap` (staticcheck ST1003).
- Exclude staticcheck ST1000/ST1016/ST1020 in `.golangci.yml` to match revive exclusions.

## [0.1.0] - 2026-06-13

### Added
- Core version library: parsing, comparison, sorting, grouping, range queries,
  constraint checking, visualization, and file I/O.
- CLI with 40+ subcommands (`versions`).
- MCP server with 22 tools (`versions-mcp`).
- 13 Claude Code Skills.
- Claude Code plugin marketplace support (`.claude-plugin/`).
- goreleaser config producing binaries for 6 OSes × 12 architectures plus
  deb/rpm/apk packages.
- GitHub Actions workflows for tests, CLI/MCP builds, and releases.
- Bilingual README (English + Chinese).

[Unreleased]: https://github.com/scagogogo/versions-skills/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/scagogogo/versions-skills/releases/tag/v0.2.0
[0.1.0]: https://github.com/scagogogo/versions-skills/releases/tag/v0.1.0
