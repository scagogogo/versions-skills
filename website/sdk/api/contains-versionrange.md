# Contains

::: info 方法 · `VersionRange`
```go
func (r *VersionRange) Contains(v *Version) bool
```
:::

## 📖 说明

Contains 判断版本是否在当前范围内


#### 参数

- `v`：要检查的版本对象


#### 返回

- `bool`：如果版本在范围内则返回 true


```go
r := versions.NewClosedRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
v := versions.NewVersion("1.5.0")
fmt.Println(r.Contains(v)) // true
```


## 🔗 同类方法

- [`VersionRange.String`](/sdk/api/string-versionrange)
- [`VersionRange.Filter`](/sdk/api/filter-versionrange)
- [`VersionRange.IsEmpty`](/sdk/api/is-empty-versionrange)


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
