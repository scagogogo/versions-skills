# 可运行示例

仓库 `examples/` 目录提供 7 个完整可运行的 Go 示例。每个都可 `cd` 进去直接 `go run main.go`。

## 📚 示例列表

| # | 目录 | 主题 | 涉及 API |
|:--|:--|:--|:--|
| 00 | `00_quick_start` | 快速开始 | `NewVersion` / `CompareTo` |
| 01 | `01_basic_version_parsing` | 基础解析 | `NewVersion` / `IsValid` / `Segments` |
| 02 | `02_reading_versions_from_file` | 从文件读取 | `ReadVersionsFromFile` |
| 03 | `03_version_sorting` | 排序 | `SortVersionSlice` |
| 04 | `04_version_grouping` | 分组 | `Group` / `VersionGroup` |
| 05 | `05_version_range_queries` | 范围查询 | `NewVersionRange` |
| 06 | `06_version_visualization` | 可视化 | `VisualizeVersions` |

## ▶️ 运行

```bash
cd examples/00_quick_start
go run main.go
```

## 📖 对照教程

每个示例都对应一篇教程：

- 00 → [10 分钟入门](/tutorials/getting-started)
- 01 → [解析与检查](/tutorials/parse-and-check)
- 02 → [文件批处理](/tutorials/file-batch)
- 03 → [排序与极值](/tutorials/sort-and-minmax)
- 04 → [分组与聚合](/tutorials/grouping)
- 05 → [范围查询](/tutorials/range-query)
- 06 → [可视化与报告](/tutorials/visualization)
