# 文件批处理

读取版本号文件，去重、排序、写回——批量维护版本清单的常见流程。

:::mermaid
flowchart LR
  FILE["releases.txt<br/>含重复/乱序"] --> READ["ReadVersionsFromFile"]
  READ --> VS["[]*Version<br/>含重复"]
  VS --> UNIQUE["Unique<br/>去重"]
  UNIQUE --> SORT["SortVersionSlice<br/>排序"]
  SORT --> WRITE["WriteVersionsToFile<br/>（自动排序）"]
  WRITE --> OUT["sorted.txt<br/>有序无重复"]

  style FILE fill:#f8fafc,stroke:#475569
  style READ fill:#eff6ff,stroke:#2563eb
  style UNIQUE fill:#eff6ff,stroke:#2563eb
  style SORT fill:#eff6ff,stroke:#2563eb
  style WRITE fill:#eff6ff,stroke:#2563eb
  style OUT fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
:::

## 📖 读取

```go
// 解析为 Version 对象
vs, err := versions.ReadVersionsFromFile("releases.txt")
if err != nil { log.Fatal(err) }

// 或只读原始字符串（不解析）
strs, err := versions.ReadVersionsStringFromFile("releases.txt")
```

文件格式：每行一个版本，空行忽略。**`#` 不是注释符**——见 [文件格式](/concepts/file-format)。

## 🧹 清洗管道

```go
// 读 → 去重 → 排序 → 写回
vs, _ := versions.ReadVersionsFromFile("releases.txt")
cleaned := versions.SortVersionSlice(versions.Unique(vs))
versions.WriteVersionsToFile(cleaned, "sorted.txt")
```

`WriteVersionsToFile` 写入前会自动排序，但先 `Unique` 能避免重复行。

## 🌊 从任意 Reader 读取

```go
// 从 stdin / HTTP body / 网络流读取
vs, _ := versions.ReadVersionsFromReader(os.Stdin)
```

## 🎯 结合过滤

```go
vs, _ := versions.ReadVersionsFromFile("all.txt")
stable := versions.FilterByStable(vs)
versions.WriteVersionsToFile(stable, "stable.txt")
```

## 🤖 CLI 一行流

```bash
# 读 → 排序 → 写
versions sort --from-file raw.txt | versions write sorted.txt

# 读 → 取最新稳定
versions latest-stable --from-file releases.txt
```

## 🚀 下一步

- [可视化与报告](/tutorials/visualization)
- API：[`ReadVersionsFromFile`](/sdk/api/read-versions-from-file) · [`WriteVersionsToFile`](/sdk/api/write-versions-to-file) · [`Unique`](/sdk/api/unique)
- CLI：[`read`](/cli/commands/read) · [`write`](/cli/commands/write) · [`read-strings`](/cli/commands/read-strings)
