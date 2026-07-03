# String

::: info 方法 · `VersionPrefix`
```go
func (x VersionPrefix) String() string
```
:::

## 📖 说明

String 返回前缀的字符串表示

实现 fmt.Stringer 接口。


#### 返回

- `string`：前缀字符串，如 "v"


## 🔗 同类方法

- [`VersionPrefix.IsEmpty`](/sdk/api/is-empty-versionprefix)
- [`VersionPrefix.PurePrefix`](/sdk/api/pure-prefix-versionprefix)


---

::: details 源码位置
定义于 [`version_prefix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_prefix.go)
:::
