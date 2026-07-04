# 10 分钟入门

本教程带你从零跑通 versions-skills：解析、比较、排序一个版本列表。

:::mermaid
flowchart LR
  INSTALL["📦 安装<br/>go get"] --> PARSE["👋 解析版本<br/>NewVersion"]
  PARSE --> COMPARE["⚖️ 比较<br/>CompareTo"]
  COMPARE --> SORT["📊 排序+极值<br/>Sort / Max / Min"]
  SORT --> NEXT["🚀 下一步<br/>约束/分组/范围"]

  style INSTALL fill:#eff6ff,stroke:#2563eb
  style PARSE fill:#eff6ff,stroke:#2563eb
  style COMPARE fill:#eff6ff,stroke:#2563eb
  style SORT fill:#eff6ff,stroke:#2563eb
  style NEXT fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
:::

## 📦 安装

```bash
go get github.com/scagogogo/versions-skills
```

## 👋 第一个版本

```go
package main

import (
	"fmt"
	"github.com/scagogogo/versions-skills"
)

func main() {
	v := versions.NewVersion("v1.2.3-rc1")

	fmt.Println(v.Raw)            // v1.2.3-rc1
	fmt.Println(v.IsValid())      // true
	fmt.Println(v.Major())        // 1
	fmt.Println(v.IsPrerelease()) // true
	fmt.Println(v.IsRC())         // true
}
```

::: tip 永不为 nil
`NewVersion` 即便传入非法字符串也返回非 nil 指针，用 `IsValid()` 判定是否合法。
:::

## ⚖️ 比较两个版本

```go
a := versions.NewVersion("1.0.0")
b := versions.NewVersion("1.0.0-rc1")

fmt.Println(a.IsNewerThan(b)) // true — 稳定版 > 预发布版
fmt.Println(a.CompareTo(b))   // 1（a 更新）
```

比较按 [四级优先级](/concepts/compare-priority)：数字段 → 后缀权重 → 发布时间 → 原始字符串。

## 📊 排序与找最新

```go
vs := versions.NewVersions("1.9.0", "1.10.0", "1.0.0", "v1.2.3-rc1")

sorted := versions.SortVersionSlice(vs)
fmt.Println(sorted[0].Raw)          // 1.0.0（最旧）
fmt.Println(sorted[len(sorted)-1].Raw) // 1.10.0（最新！不是 1.9.0）

fmt.Println(versions.Max(vs).Raw)   // 1.10.0
```

::: warning 语义排序
`1.10.0 > 1.9.0`——按数字比较，不是字典序。这是 versions-skills 解决的[核心问题](/why)。
:::

## 🚀 下一步

- [解析与检查](/tutorials/parse-and-check) — 深入拆解版本号
- [排序与极值](/tutorials/sort-and-minmax) — 更多排序技巧
- [快速开始](/quick-start) — 三层接入速查
