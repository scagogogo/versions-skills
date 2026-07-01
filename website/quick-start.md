# 快速开始

三种使用方式，按场景选。

## AI Agent（最省事）

直接用 [一键提示词](./prompts)，把提示词粘到 Claude Code / Codex，AI 自动装好。接入后即可对话式使用：

```
你：帮我排序这些版本并标出预发布版
    2.0.0, 1.0.0, 1.10.0, 1.2.0-beta, 1.2.0

AI：[调用 version_sort]
    1.0.0
    1.2.0-beta   ← 预发布
    1.2.0
    1.10.0
    2.0.0
```

## Go SDK

```bash
go get github.com/scagogogo/versions-skills
```

```go
package main

import (
	"fmt"
	"github.com/scagogogo/versions-skills"
)

func main() {
	// 解析
	v := versions.NewVersion("v1.2.3-beta1")
	fmt.Println(v.Major(), v.Minor(), v.Patch()) // 1 2 3
	fmt.Println(v.IsValid())                      // true
	fmt.Println(v.PreReleaseType())               // beta

	// 比较
	a := versions.NewVersion("1.0.0")
	b := versions.NewVersion("1.0.0-rc1")
	fmt.Println(a.IsNewerThan(b)) // true（正式版 > 预发布版）

	// 排序
	list := versions.NewVersions("2.0.0", "1.0.0", "1.10.0", "1.5.0-beta")
	sorted := versions.SortVersionSlice(list)
	for _, v := range sorted {
		fmt.Println(v.Raw)
	}
	// 1.0.0 / 1.5.0-beta / 1.10.0 / 2.0.0

	// 约束
	ok, _ := versions.NewVersion("1.5.0").Matches(">=1.0.0,<2.0.0 || >=3.0.0")
	fmt.Println(ok) // true
}
```

## CLI

```bash
# 安装
go install github.com/scagogogo/versions-skills/cmd/versions@latest
# 或一键脚本
curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash

# 用法
versions parse v1.2.3-rc1
versions compare 1.0.0 2.0.0
versions sort 3.0.0 1.0.0 2.0.0
versions filter --constraint ">=1.0.0,<2.0.0" 0.5.0 1.0.0 1.5.0 2.0.0
versions latest-stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0
versions visualize 1.0.0 1.1.0 2.0.0 --groups
```

## MCP Server

```bash
go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest
```

配置见 [AI Agent 接入指南](./ai-agents)。

---

→ 深入了解能力：[Go SDK API](./sdk)、[CLI 命令](./cli)、[MCP 工具](./mcp)。
→ 理解原理：[算法详解](./algorithms)。
