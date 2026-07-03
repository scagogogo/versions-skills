# PurePrefix

::: info 方法 · `VersionPrefix`
```go
func (x VersionPrefix) PurePrefix() string
```
:::

## 📖 说明

PurePrefix 返回去除分隔符后的纯净前缀

纯净前缀是去除了末尾分隔符（如 "-"、"."、"_"）的前缀部分。
例如 "curl-" 的纯净前缀为 "curl"，"v" 的纯净前缀仍为 "v"。


#### 返回

- `string`：去除末尾分隔符后的前缀


## 🔗 同类方法

- [`VersionPrefix.IsEmpty`](/sdk/api/is-empty-versionprefix)
- [`VersionPrefix.String`](/sdk/api/string-versionprefix)


---

::: details 源码位置
定义于 [`version_prefix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_prefix.go)
:::
