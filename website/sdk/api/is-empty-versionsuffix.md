# IsEmpty

::: info 方法 · `VersionSuffix`
```go
func (x VersionSuffix) IsEmpty() bool
```
:::

## 📖 说明

IsEmpty 判断版本后缀是否为空

该方法检查版本后缀是否为空，即是否等于 EmptyVersionSuffix。


#### 返回

- `bool`：如果后缀为空则返回 true，否则返回 false


```go
version := versions.NewVersion("1.2.3") // 没有后缀
if version.Suffix.IsEmpty() {
fmt.Println("版本没有后缀")
}

version2 := versions.NewVersion("1.2.3-rc1") // 有后缀
if !version2.Suffix.IsEmpty() {
fmt.Printf("版本后缀是: %s\n", version2.Suffix)
}
```


---

::: details 源码位置
定义于 [`version_suffix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_suffix.go)
:::
