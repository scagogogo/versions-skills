# IsEmpty

::: info 方法 · `VersionRange`
```go
func (r *VersionRange) IsEmpty() bool
```
:::

## 📖 说明

IsEmpty 判断范围是否为空（下界大于上界）


#### 返回

- `bool`：如果范围为空则返回 true


## 🔗 同类方法

- [`VersionRange.Contains`](/sdk/api/contains-versionrange)
- [`VersionRange.String`](/sdk/api/string-versionrange)
- [`VersionRange.Filter`](/sdk/api/filter-versionrange)


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
