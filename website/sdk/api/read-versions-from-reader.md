# ReadVersionsFromReader

::: info 函数 · 根包
```go
func ReadVersionsFromReader(reader io.Reader) ([]*Version, error)
```
:::

## 📖 说明

ReadVersionsFromReader 从 io.Reader 读取版本号并解析

该函数从任意的 io.Reader 中读取版本号列表，每行一个版本号，
并将其解析为 Version 对象数组。适用于从网络连接、字符串缓冲区等读取版本。


#### 参数

- `reader`：实现 io.Reader 接口的读取器


#### 返回

- `[]*Version`：解析后的 Version 对象数组
- `error`：如果读取失败则返回错误


```go
data := strings.NewReader("1.0.0\n1.1.0\n2.0.0\n")
versions, err := versions.ReadVersionsFromReader(data)
```


---

::: details 源码位置
定义于 [`file.go`](https://github.com/scagogogo/versions-skills/blob/main/file.go)
:::
