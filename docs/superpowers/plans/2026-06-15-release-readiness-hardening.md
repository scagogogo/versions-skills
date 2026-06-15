# Release-Readiness Hardening Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** Take versions-skills from "already released as v0.1.0" to a polished, officially-public project by closing standard OSS hygiene gaps, raising the quality bar with a real linter, and fixing documentation warts — then cutting a clean v0.2.0 that includes the CI/CD improvements from PR #1.

**Architecture:** No runtime code changes. Data flow is: new community/health files (CHANGELOG, CONTRIBUTING, SECURITY, CODE_OF_CONDUCT) + tooling config (.golangci.yml, dependabot.yml, fixed .gitignore) are added to the repo root and `.github/`; CI gains a golangci-lint step; README + 13 SKILL.md files get doc corrections; finally PR #1 is merged and a v0.2.0 tag triggers the existing goreleaser pipeline. Reuses existing goreleaser config and the 3 existing workflows — no new pipeline.

**Tech Stack:** Go 1.23, cobra v1.10, mcp-go v0.32, testify v1.9, goreleaser v2, GitHub Actions, golangci-lint v1.x, Keep a Changelog, Contributor Covenant 2.1.

**Risks:**
- Task 3 introduces golangci-lint, which may surface pre-existing findings (the 2 dated TODO notes in `parser.go`, possible `ineffassign`/`errcheck` hits) → mitigation: configure `.golangci.yml` to exclude the known parser TODOs as documented future-work, and fix any real findings rather than silencing.
- Task 5 edits 13 SKILL.md files that share an identical `{VERSION}` pattern → mitigation: README gets the full rewrite; the shared misleading `e.g. 0.2.0` example is corrected with a single verified `sed` so all 13 stay consistent.
- Task 6 (merge PR #1 + tag v0.2.0) is outward-facing and hard to reverse → mitigation: it is a checkpoint task; do NOT execute until the user confirms the version number and timing.

---

## Pre-Planning Analysis (assessment of current state)

**Feature:** Release-readiness hardening
**Scope:** Single repo, no runtime code changes (docs + config + CI only)

**What is ALREADY done and solid:**
- Core library: ~7,700 LOC across 30+ source files (parse, compare, sort, group, range, constraint, visualize, file I/O, builder, clone, suffix weights). Zero core dependencies.
- Tests: all pass with `-race`; 22 test files; ~88% core coverage.
- CLI: 17 command files → 40+ subcommands (parse, compare, sort, group, info, validate, check, constraint, range, filter, min/max/latest-stable/latest-prerelease, build, bump, set*, partition, count, core, segments, suffixWeight, read/write files, visualize). `--version` is wired (cobra `Version` field + goreleaser ldflags injection).
- MCP server: 22 tools.
- Skills: 13 SKILL.md files, all present, all with SDK/CLI/MCP paths.
- Plugin marketplace: `.claude-plugin/{plugin,marketplace}.json` — `claude plugin validate . --strict` PASSES; installable via `claude marketplace add versions https://github.com/scagogogo/versions-skills`.
- CI/CD: 3 workflows (go-test, go-build-cli, go-release), all green, Node 24 opted in.
- goreleaser: full multi-platform (6 OSes × 12 archs), deb/rpm/apk packages for both `versions` and `versions-mcp`.
- README: 428 lines, bilingual (English default + Chinese), 5 badges.
- **Already publicly released**: `v0.1.0` exists on GitHub as a full goreleaser release (checksums + all platform assets), dated 2026-06-13.

**Gaps (none are blockers — all are polish):**
1. Missing standard OSS files: `CHANGELOG.md`, `CONTRIBUTING.md`, `SECURITY.md`, `CODE_OF_CONDUCT.md`.
2. `.gitignore` is near-empty (`test`, `test2`) — does NOT ignore `dist/`, `coverage.out`, built binaries (`versions`, `versions-mcp`), `*.exe`. Risk of accidentally committing artifacts.
3. No `.github/dependabot.yml` → cobra/mcp-go/testify go stale silently.
4. CI only runs `go vet` + gofmt; no `golangci-lint` → misses staticcheck/ineffassign/errcheck classes.
5. README uses the deprecated `godoc.org` badge instead of `pkg.go.dev`.
6. `install.sh` exists and auto-detects platform/version, but is surfaced in only 1 of 13 skills and NOT in README.
7. `{VERSION}` download placeholder: works when substituted, but the README + skills show `releases/latest/download/...{VERSION}...` with a misleading example `e.g. 0.2.0` (latest is `0.1.0`). Users must manually look up the version to use a "latest" link.
8. PR #1 (`ci/cicd-enhancements`) is unmerged → `main` and the next release lack the latest CI/CD + goreleaser improvements. Two dated TODO design notes in `parser.go` (future parsing enhancements, not bugs) remain.

**Verdict:** Releasable publicly — YES, and it already is (`v0.1.0`). The work below elevates it to a polished, officially-supported project and produces a clean `v0.2.0` that bundles PR #1.

**Files Create:** `CHANGELOG.md`, `CONTRIBUTING.md`, `SECURITY.md`, `CODE_OF_CONDUCT.md`, `.github/dependabot.yml`, `.golangci.yml`
**Files Modify:** `.gitignore`, `.github/workflows/go-test.yml` (lint job), `README.md` (badge + install.sh + version example), `skills/*/SKILL.md` (misleading version example)
**Tasks:** 6
**Order:** 1 → 2 → 3 → 4 → 5 → 6 (6 is the checkpoint)

---

### Task 1: Repository hygiene files

**Depends on:** None
**Files:**
- Create: `CHANGELOG.md`
- Create: `CONTRIBUTING.md`
- Create: `SECURITY.md`
- Create: `CODE_OF_CONDUCT.md`

- [ ] **Step 1: 创建 CHANGELOG.md — 采用 Keep a Changelog 格式，记录已发布的 v0.1.0**

```markdown
# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- golangci-lint to CI and `.golangci.yml` configuration.
- Dependabot configuration for Go modules and GitHub Actions.
- Community files: CONTRIBUTING, SECURITY, CODE_OF_CONDUCT.

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
```

- [ ] **Step 2: 创建 CONTRIBUTING.md — 说明开发流程、测试、提交规范、插件校验**

```markdown
# Contributing to versions-skills

Thanks for your interest in contributing! This project is a Go library + CLI +
MCP server + Claude Code Skills bundle for version number operations.

## Development setup

Requires Go 1.23+.

```bash
git clone https://github.com/scagogogo/versions-skills.git
cd versions-skills
go test ./...
```

## Before opening a pull request

1. **Format and vet your code.**
   ```bash
   gofmt -w .
   go vet ./...
   ```
2. **Run the linter.**
   ```bash
   golangci-lint run ./...
   ```
3. **Run all tests with the race detector.**
   ```bash
   go test -race ./...
   ```
4. **Validate the plugin manifests** (if you touched `.claude-plugin/` or skills).
   ```bash
   claude plugin validate . --strict
   ```

## Commit messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):
`feat(scope): ...`, `fix(scope): ...`, `docs: ...`, `test: ...`, `refactor: ...`,
`ci: ...`, `chore: ...`.

## Adding or modifying skills

Skills live in `skills/*/SKILL.md`. Each skill documents three access paths
(SDK, CLI, MCP) with code examples and an "Important Notes" section. Match the
existing structure and depth. See `CLAUDE.md` for details.

## Reporting issues

Open a [GitHub issue](https://github.com/scagogogo/versions-skills/issues).
For security vulnerabilities, see [SECURITY.md](SECURITY.md) instead of a
public issue.
```

- [ ] **Step 3: 创建 SECURITY.md — 漏洞报告策略**

```markdown
# Security Policy

## Supported versions

Only the latest release line receives security fixes.

## Reporting a vulnerability

Please **do not** open a public GitHub issue for security vulnerabilities.

Report vulnerabilities privately by emailing the maintainer at
**scagogogo@gmail.com**. Include:

- A description of the issue and its impact.
- Steps to reproduce or a proof of concept.
- Affected versions, if known.

You should receive an acknowledgment within 72 hours. Please allow reasonable
time for a fix to be developed and published before any public disclosure.
```

- [ ] **Step 4: 创建 CODE_OF_CONDUCT.md — Contributor Covenant 2.1**

```markdown
# Contributor Covenant Code of Conduct

## Our Pledge

We as members, contributors, and leaders pledge to make participation in our
community a harassment-free experience for everyone, regardless of age, body
size, visible or invisible disability, ethnicity, sex characteristics, gender
identity and expression, level of experience, education, socio-economic status,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

We pledge to act and interact in ways that contribute to an open, welcoming,
diverse, inclusive, and healthy community.

## Our Standards

Examples of behavior that contributes to a positive environment include:

- Demonstrating empathy and kindness toward other people.
- Being respectful of differing opinions, viewpoints, and experiences.
- Giving and gracefully accepting constructive feedback.
- Accepting responsibility and apologizing to those affected by our mistakes,
  and learning from the experience.
- Focusing on what is best for the overall community.

Examples of unacceptable behavior include:

- The use of sexualized language or imagery, and sexual attention or advances.
- Trolling, insulting or derogatory comments, and personal or political attacks.
- Public or private harassment.
- Publishing others' private information without explicit permission.
- Other conduct which could reasonably be considered inappropriate in a
  professional setting.

## Enforcement Responsibilities

Community leaders are responsible for clarifying and enforcing our standards of
acceptable behavior and will take appropriate and fair corrective action in
response to any behavior that they deem inappropriate, threatening, offensive,
or harmful.

## Scope

This Code of Conduct applies within all community spaces, and also applies when
an individual is officially representing the community in public spaces.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported to the community leaders responsible for enforcement at
**scagogogo@gmail.com**. All complaints will be reviewed and investigated
promptly and fairly.

See the [Contributor Covenant](https://www.contributor-covenant.org/version/2/1/code_of_conduct/)
for the full text, including enforcement guidelines and attribution.
```

- [ ] **Step 5: 验证文件存在且非空**
Run: `test -s CHANGELOG.md && test -s CONTRIBUTING.md && test -s SECURITY.md && test -s CODE_OF_CONDUCT.md && echo OK`
Expected:
  - Exit code: 0
  - Output contains: "OK"

- [ ] **Step 6: 提交**
Run: `git add CHANGELOG.md CONTRIBUTING.md SECURITY.md CODE_OF_CONDUCT.md && git commit -m "docs: add CHANGELOG, CONTRIBUTING, SECURITY, and CODE_OF_CONDUCT"`

---

### Task 2: Fix .gitignore and add Dependabot

**Depends on:** None
**Files:**
- Modify: `.gitignore` (full replacement)
- Create: `.github/dependabot.yml`

- [ ] **Step 1: 替换 .gitignore — 忽略构建产物、覆盖率、编辑器文件**

```text
# Binaries / build outputs
/versions
/versions-mcp
*.exe
*.exe~
*.dll
*.so
*.dylib
/dist/

# Test / coverage artifacts
coverage.out
coverage.html
*.prof

# Go workspace
go.work
go.work.sum

# Editor / OS
.DS_Store
.idea/
.vscode/
*.swp
*~

# Legacy local dirs
test
test2
```

- [ ] **Step 2: 创建 .github/dependabot.yml — 每周检查 Go 模块与 GitHub Actions 更新**

```yaml
version: 2
updates:
  - package-ecosystem: gomod
    directory: /
    schedule:
      interval: weekly
    open-pull-requests-limit: 5
  - package-ecosystem: github-actions
    directory: /
    schedule:
      interval: weekly
    open-pull-requests-limit: 5
```

- [ ] **Step 3: 验证 .gitignore 现在忽略 dist 与二进制**
Run: `printf 'dist/coverage.out\n' | git check-ignore dist/ coverage.out /versions && echo OK`
Expected:
  - Exit code: 0
  - Output contains: "OK"

- [ ] **Step 4: 提交**
Run: `git add .gitignore .github/dependabot.yml && git commit -m "chore: expand .gitignore and add Dependabot config"`

---

### Task 3: Add golangci-lint config and wire into CI

**Depends on:** None
**Files:**
- Create: `.golangci.yml`
- Modify: `.github/workflows/go-test.yml` (the `lint` job)

- [ ] **Step 1: 创建 .golangci.yml — 启用 vet/errcheck/staticcheck/ineffassign/unused/gofmt/revive**

```yaml
run:
  timeout: 5m
  go: "1.23"

linters:
  enable:
    - govet
    - errcheck
    - staticcheck
    - ineffassign
    - unused
    - gofmt
    - revive
    - gocritic
    - misspell

linters-settings:
  govet:
    enable-all: false
  revive:
    severity: warning

issues:
  exclude-rules:
    # parser.go contains two dated design notes about future multi-version
    # parsing enhancements; they are intentional future-work, not TODO debt.
    - path: parser\.go
      linters:
        - gocritic
      text: "TODO"
  max-issues-per-linter: 0
  max-same-issues: 0
```

- [ ] **Step 2: 修改 go-test.yml 的 lint 作业 — 在 go vet/gofmt 之后运行 golangci-lint**
文件: `.github/workflows/go-test.yml:52-78`（替换整个 `lint` job）

```yaml
  lint:
    name: Code Quality
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.23'
          check-latest: true

      - name: Run go vet
        run: go vet ./...

      - name: Run gofmt check
        run: |
          output=$(gofmt -l .)
          if [ -n "$output" ]; then
            echo "❌ The following files are not properly formatted:"
            echo "$output"
            echo ""
            echo "Run 'gofmt -w .' to fix formatting."
            exit 1
          fi
          echo "✅ All files are properly formatted."

      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v6
        with:
          version: latest
          args: --timeout 5m
```

- [ ] **Step 3: 验证本地 golangci-lint 通过**
Run: `golangci-lint run --timeout 5m ./... 2>&1 | tail -20`
Expected:
  - Exit code: 0
  - Output does NOT contain: "error" / "Failed" (warnings about unused config are acceptable; fix any real findings before committing)

- [ ] **Step 4: 提交**
Run: `git add .golangci.yml .github/workflows/go-test.yml && git commit -m "ci: add golangci-lint with .golangci.yml and wire into go-test workflow"`

---

### Task 4: Polish README version docs and surface install.sh

**Depends on:** None
**Files:**
- Modify: `README.md` (badge line ~7, CLI install section ~47-60, Installation section ~317-356)

- [ ] **Step 1: 修改 GoDoc 徽章 — 从已弃用的 godoc.org 换成 pkg.go.dev**
文件: `README.md:7`

```markdown
[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/versions-skills.svg)](https://pkg.go.dev/github.com/scagogogo/versions-skills)
```

- [ ] **Step 2: 在 README 的 CLI 接入区突出 install.sh 一键安装**
文件: `README.md:47-60`（替换 CLI 接入区块开头）

```markdown
### 💻 CLI — Recommended for scripts and CI/CD

One-line install (Linux/macOS, auto-detects platform and version):

```bash
curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash
```

# Or download a binary from GitHub Releases
# https://github.com/scagogogo/versions-skills/releases/latest
```

- [ ] **Step 3: 修正 Installation 区里误导的 {VERSION} 示例**
文件: `README.md:343` 与 `README.md:356`（`e.g. 0.2.0` → 当前正确表述）

将：
```markdown
> Replace `{VERSION}` with the latest release tag (e.g. `0.2.0`). See the [releases page](https://github.com/scagogogo/versions-skills/releases/latest) for all available platforms and the current version.
```
改为：
```markdown
> `{VERSION}` is the release tag shown at the top of the [releases page](https://github.com/scagogogo/versions-skills/releases/latest). Prefer the one-line `install.sh` above, which resolves it automatically.
```

- [ ] **Step 4: 验证 README 无残留 0.2.0 误导字样**
Run: `! grep -n "e.g. .0\.2\.0" README.md && echo OK`
Expected:
  - Exit code: 0
  - Output contains: "OK"

- [ ] **Step 5: 提交**
Run: `git add README.md && git commit -m "docs: use pkg.go.dev badge, surface install.sh, fix version-download guidance"`

---

### Task 5: Fix misleading version example across all 13 skills

**Depends on:** None
**Files:**
- Modify: `skills/*/SKILL.md` (the `(e.g. 0.2.0)` example and install guidance)

- [ ] **Step 1: 统一修正 13 个 skill 里的误导版本示例**
Run: `grep -rln "(e.g. .0\.2\.0)" skills/ | xargs sed -i 's/ (e.g. `0\.2\.0`)//; s/(e.g. 0\.2\.0)//'`
Expected:
  - Exit code: 0
  - `grep -rn "0.2.0" skills/` returns nothing

- [ ] **Step 2: 在每个 skill 的 CLI 安装段补一行 install.sh 提示（已存在于 cli-operations，跳过）**
Run: `for f in skills/*/SKILL.md; do grep -q "install.sh" "$f" || sed -i '/releases\/latest/a\> One-line install: \`curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash\` (resolves version automatically).' "$f"; done`
Expected:
  - Exit code: 0
  - Every `skills/*/SKILL.md` contains "install.sh"

- [ ] **Step 3: 验证插件 manifest 与 skill 仍有效**
Run: `claude plugin validate . --strict`
Expected:
  - Output contains: "Validation passed"

- [ ] **Step 4: 提交**
Run: `git add skills/ && git commit -m "docs(skills): fix misleading version example and surface install.sh"`

---

### Task 6: Merge PR #1 and cut v0.2.0 — CHECKPOINT (needs user confirmation)

**Depends on:** Task 1, Task 2, Task 3, Task 4, Task 5
**Files:** none (release process)

> ⚠️ This task is outward-facing and hard to reverse (public release tag + main merge). Execute ONLY after the user confirms the version number (v0.2.0 recommended) and timing. The user's question was an assessment; do not publish without explicit go-ahead.

- [ ] **Step 1: 推送 polish 提交到 PR #1 分支**
Run: `git push origin ci/cicd-enhancements`
Expected:
  - Exit code: 0
  - CI on PR #1 runs green

- [ ] **Step 2: 确认 PR #1 全部检查通过**
Run: `gh pr checks 1 --watch`
Expected:
  - All checks: pass/success

- [ ] **Step 3: 合并 PR #1 到 main（squash 或 merge 由用户偏好决定）**
Run: `gh pr merge 1 --squash --delete-branch`
Expected:
  - Exit code: 0

- [ ] **Step 4: 打 v0.2.0 tag 并推送，触发 goreleaser 生产发布**
Run: `git checkout main && git pull && git tag v0.2.0 && git push origin v0.2.0`
Expected:
  - `go-release.yml` triggers and produces the GitHub Release with multi-platform assets

- [ ] **Step 5: 监控发布完成**
Run: `gh run watch $(gh run list --workflow=go-release.yml --limit 1 --json databaseId -q '.[0].databaseId')`
Expected:
  - Release job: success
  - `gh release view v0.2.0` shows checksums + platform archives + deb/rpm/apk

---
