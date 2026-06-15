# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- golangci-lint to CI and `.golangci.yml` configuration.
- Dependabot configuration for Go modules and GitHub Actions.
- Community files: CONTRIBUTING, SECURITY, CODE_OF_CONDUCT.
- Expanded `.gitignore` for build artifacts, coverage, and binaries.

### Changed
- README badge now points to pkg.go.dev instead of the deprecated godoc.org.
- Surfaces the one-line `install.sh` installer in README and skills.

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

[Unreleased]: https://github.com/scagogogo/versions-skills/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/scagogogo/versions-skills/releases/tag/v0.1.0
