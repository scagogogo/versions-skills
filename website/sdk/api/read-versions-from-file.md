# ReadVersionsFromFile

::: info 函数 · 根包
```go
func ReadVersionsFromFile(filepath string) ([]*Version, error)
```
:::

## 📖 说明

ReadVersionsFromFile 从文件中读取版本号并解析为Version对象

该函数从指定的文件中读取版本号列表，每行一个版本号，并将其解析为Version对象数组。
函数会自动忽略空行和进行行尾空白字符的清理。


#### 参数

- `filepath`：包含版本号列表的文件路径


#### 返回

- `[]*Version`：解析后的Version对象数组
- `error`：如果文件读取失败则返回相应错误

示例文件内容:

1.1.28
1.1.29
1.1.30
1.1.31
1.1.31.sec01
1.1.31.sec04
1.1.31.sec06


```go
// 从版本列表文件中读取版本
versions, err := versions.ReadVersionsFromFile("./versions.txt")
if err != nil {
    log.Fatalf("读取版本文件失败: %v", err)
}

// 打印解析的版本数
fmt.Printf("共读取 %d 个版本\n", len(versions))

// 对版本进行排序
sortedVersions := versions.SortVersionSlice(versions)
```


---

::: details 源码位置
定义于 [`file.go`](https://github.com/scagogogo/versions-skills/blob/main/file.go)
:::
