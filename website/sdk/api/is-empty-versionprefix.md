# IsEmpty

::: info 方法 · `VersionPrefix`
```go
func (x VersionPrefix) IsEmpty() bool
```
:::

## 📖 说明

IsEmpty 返回前缀是否为空

该方法检查版本前缀是否为空，即是否等于 EmptyVersionPrefix。


#### 返回

- `bool`：如果前缀为空则返回 true，否则返回 false


```go
version := versions.NewVersion("1.2.3") // 没有前缀
if version.Prefix.IsEmpty() {
    fmt.Println("版本没有前缀")
}

version2 := versions.NewVersion("v1.2.3") // 有前缀
if !version2.Prefix.IsEmpty() {
    fmt.Printf("版本前缀是: %s\n", version2.Prefix)
}
```


## 🔗 同类方法

- [`VersionPrefix.String`](/sdk/api/string-versionprefix)
- [`VersionPrefix.PurePrefix`](/sdk/api/pure-prefix-versionprefix)


---

::: details 源码位置
定义于 [`version_prefix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_prefix.go)
:::
