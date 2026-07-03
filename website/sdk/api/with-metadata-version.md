# WithMetadata

::: info 方法 · `Version`
```go
func (x *Version) WithMetadata(metadata string) *Version
```
:::

## 📖 说明

WithMetadata 返回一个修改构建元数据的新版本对象

原版本对象不变，返回一个新对象，其 Metadata 字段被替换为指定值。


#### 参数

- `metadata`：新的构建元数据字符串，如 "build123"


#### 返回

- `*Version`：修改元数据后的新版本对象


```go
v := versions.NewVersion("1.2.3")
newV := v.WithMetadata("build.123")
fmt.Println(newV.Metadata) // "build.123"
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
