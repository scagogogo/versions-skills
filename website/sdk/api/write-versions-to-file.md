# WriteVersionsToFile

::: info 函数 · 根包
```go
func WriteVersionsToFile(versions []*Version, filepath string) error
```
:::

## 📖 说明

WriteVersionsToFile 将版本列表写入文件

每个版本号占一行，写入版本号数字部分拼接而成的字符串。
该函数会先对版本进行排序，确保输出有序。


#### 参数

- `versions`：要写入的版本对象列表
- `filepath`：输出文件路径


#### 返回

- `error`：如果文件写入失败则返回错误


```go
vs := versions.NewVersions("2.0.0", "1.0.0", "1.1.0")
err := versions.WriteVersionsToFile(vs, "./output.txt")
```


---

::: details 源码位置
定义于 [`file.go`](https://github.com/scagogogo/versions-skills/blob/main/file.go)
:::
