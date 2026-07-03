# Filter

::: info 方法 · `VersionRange`
```go
func (r *VersionRange) Filter(versions []*Version) []*Version
```
:::

## 📖 说明

Filter 过滤版本列表，只保留在范围内的版本


#### 参数

- `versions`：版本对象列表


#### 返回

- `[]*Version`：范围内的版本列表


## 🔗 同类方法

- [`VersionRange.Contains`](/sdk/api/contains-versionrange)
- [`VersionRange.String`](/sdk/api/string-versionrange)
- [`VersionRange.IsEmpty`](/sdk/api/is-empty-versionrange)


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
