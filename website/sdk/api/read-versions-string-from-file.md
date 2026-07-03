# ReadVersionsStringFromFile

::: info 函数 · 根包
```go
func ReadVersionsStringFromFile(filepath string) ([]string, error)
```
:::

## 📖 说明

ReadVersionsStringFromFile 从文件中读取版本号字符串

该函数从指定的文件中读取版本号列表，每行一个版本号，并返回字符串数组。
与 ReadVersionsFromFile 不同，此函数不会将版本号解析为Version对象，
而是保留原始字符串形式，适用于不需要解析和比较的场景。


#### 参数

- `filepath`：包含版本号列表的文件路径


#### 返回

- `[]string`：读取的版本号字符串数组
- `error`：如果文件读取失败则返回相应错误


```go
// 从版本列表文件中读取版本字符串
versionStrings, err := versions.ReadVersionsStringFromFile("./versions.txt")
if err != nil {
log.Fatalf("读取版本文件失败: %v", err)
}

// 使用版本字符串
for _, vStr := range versionStrings {
fmt.Printf("发现版本: %s\n", vStr)
}
```


---

::: details 源码位置
定义于 [`file.go`](https://github.com/scagogogo/versions-skills/blob/main/file.go)
:::
