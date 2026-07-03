# 可视化与报告

把版本号列表渲染成文本树状图，便于理解层级关系。

## 🌳 版本列表可视化

```go
vs := versions.NewVersions("1.0.0", "1.0.1", "1.1.0", "2.0.0-rc1", "2.0.0")
versions.VisualizeVersions(vs, os.Stdout, 0)
```

输出层级树状图（默认 `maxItems=0` 不限每组数量）。

## 🗂 分组可视化

```go
versions.VisualizeVersionGroups(vs, os.Stdout)
```

按分组渲染摘要树，每组显示其下的版本。

## 🎚 限制每组数量

```go
// 每组最多显示 3 个版本
versions.VisualizeVersions(vs, os.Stdout, 3)
```

## 🤖 CLI

```bash
versions visualize 1.0.0 1.0.1 1.1.0 2.0.0-rc1 2.0.0
versions visualize --groups --from-file releases.txt
versions visualize --max-items 5 --from-file releases.txt
```

## 📊 Mermaid 图表

文档站内置 Mermaid 支持。除了文本树，你还可以在 Markdown 里画版本演进图：

```mermaid
graph LR
  A[1.0.0] --> B[1.0.1]
  B --> C[1.1.0]
  C --> D[2.0.0-rc1]
  D --> E[2.0.0]
```

## 🚀 下一步

- [让 Claude 管理版本](/tutorials/ai-version-management)
- API：[`VisualizeVersions`](/sdk/api/visualize-versions) · [`VisualizeVersionGroups`](/sdk/api/visualize-version-groups)
- CLI：[`visualize`](/cli/commands/visualize)
