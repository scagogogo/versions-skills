# 文件格式

::: tip 关键
**每行一个版本号字符串**，空行忽略。`#` **不是**注释符——`#`-开头的行会被当作（非法）版本号解析。
:::

## 📄 格式规范

```
1.0.0
1.1.0
v1.2.3-rc1

2.0.0
```

- 每行一个版本
- **空行被忽略**
- 前后空白会被 `TrimSpace`
- **无注释语法**：`# 这是一行` 会被解析为名为 `# 这是一行` 的非法版本

::: warning 重要
若你的输入文件含 `#` 注释，须**在调用前自行剥离**。versions-skills 不会把 `#` 当注释。
:::

## 📝 读写 API

| 函数 | 说明 |
|:--|:--|
| `ReadVersionsFromFile(path)` | 读文件，返回 `[]*Version`（解析） |
| `ReadVersionsStringFromFile(path)` | 读文件，返回 `[]string`（不解析，原始行） |
| `WriteVersionsToFile(versions, path)` | **排序后**写入文件 |
| `ReadVersionsFromReader(r)` | 从任意 `io.Reader` 读取 |

```go
vs, err := versions.ReadVersionsFromFile("releases.txt")
// 顺便去重排序
versions.WriteVersionsToFile(versions.Unique(versions.SortVersionSlice(vs)), "sorted.txt")
```

## ✍ 写入行为

`WriteVersionsToFile` 在写入前会**先排序**——输出文件总是有序的。若想保留原始顺序，需自行用 `ReadVersionsStringFromFile` + 原始字符串写入。

## 📚 延伸

- API：[`ReadVersionsFromFile`](/sdk/api/read-versions-from-file) · [`WriteVersionsToFile`](/sdk/api/write-versions-to-file) · [`ReadVersionsFromReader`](/sdk/api/read-versions-from-reader)
- CLI：[`read`](/cli/commands/read) · [`write`](/cli/commands/write) · [`read-strings`](/cli/commands/read-strings)
- MCP：[`version_read_file`](/mcp/tools/version-read-file) · [`version_write_file`](/mcp/tools/version-write-file)
