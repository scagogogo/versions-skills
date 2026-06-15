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
`feat(scope): ...`, `fix(scope): ...`, `docs: ...`, `test: ...`,
`refactor: ...`, `ci: ...`, `chore: ...`.

## Adding or modifying skills

Skills live in `skills/*/SKILL.md`. Each skill documents three access paths
(SDK, CLI, MCP) with code examples and an "Important Notes" section. Match the
existing structure and depth. See `CLAUDE.md` for details.

## Reporting issues

Open a [GitHub issue](https://github.com/scagogogo/versions-skills/issues).
For security vulnerabilities, see [SECURITY.md](SECURITY.md) instead of a
public issue.
