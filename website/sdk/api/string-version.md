# String

::: info 方法 · `Version`
```go
func (x *Version) String() string
```
:::

## 📖 说明

String 返回版本的JSON字符串表示

该方法将Version对象序列化为JSON字符串，便于打印和调试。


#### 返回

- `string`：版本的JSON字符串表示


```go
version := versions.NewVersion("1.2.3")
fmt.Println(version.String()) // 输出JSON格式的版本信息
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
