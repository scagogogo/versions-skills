---
name: installation
description: One-time setup for versions-skills SDK, CLI, and MCP server. Run first before using any version-* skill.
argument-hint: <sdk|cli|mcp>
---

# Installation

> **Run this once** before using any `version-*` skill. All other skills assume the tooling is available.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

## Decision Tree

```
Which access path does the task need?
├─ SDK (Go code)        → Step 1
├─ CLI (shell/scripts)  → Step 2
└─ MCP (AI tool calls)  → Step 3
```

---

## Step 1: SDK (Go library)

```bash
go get github.com/scagogogo/versions-skills
```

```go
import "github.com/scagogogo/versions-skills"
```

Verify:
```go
v := versions.NewVersion("1.2.3")
fmt.Println(v.IsValid()) // true
```

---

## Step 2: CLI binary

**Option A — One-line installer (recommended):**
```bash
curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash
```

**Option B — Manual download:**
Pick the archive for your platform from the [releases page](https://github.com/scagogogo/versions-skills/releases/latest). Replace `{VERSION}` and `{OS}_{ARCH}`:

```bash
curl -sL https://github.com/scagogogo/versions-skills/releases/latest/download/versions_{VERSION}_{OS}_{ARCH}.tar.gz | tar xz
chmod +x versions && sudo mv versions /usr/local/bin/
```

**Option C — Go install:**
```bash
go install github.com/scagogogo/versions-skills/cmd/versions@latest
```

Verify:
```bash
versions --version
versions parse 1.2.3
```

---

## Step 3: MCP server

**Option A — Download binary** (see Step 2, use `versions-mcp` binary name).

**Option B — Go install:**
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

For SSE mode (network-accessible server):
```bash
versions-mcp --transport sse --port 8080
```

Available tools: `version_parse`, `version_validate`, `version_info`, `version_compare`, `version_sort`, `version_group`, `version_constraint_check`, `version_range_query`, `version_filter`, `version_min`, `version_max`, `version_latest_stable`, `version_latest_prerelease`, `version_unique`, `version_set_operation`, `version_visualize`, `version_build`, `version_bump`, `version_core`, `version_read_file`, `version_write_file`.

---

## Important Notes

- **CLI default output is JSON** — use `-f table` or `-f text` for human-readable output.
- **CLI quiet mode** (`-q`) outputs raw data without the envelope — ideal for shell pipelines.
- **MCP tools return JSON** — parse the `content[0].text` field.
- **Platform coverage**: Linux, macOS, Windows, FreeBSD, OpenBSD, NetBSD on amd64, arm64, arm, 386, mips, mips64, mips64le, ppc64, ppc64le, s390x, riscv64. Linux packages: deb, rpm, apk.
