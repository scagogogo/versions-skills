# Equals

::: info 方法 · `VersionNumbers`
```go
func (x VersionNumbers) Equals(target VersionNumbers) bool
```
:::

## 📖 说明

Equals 判断两个版本号数字部分是否相等


#### 参数

- `target`：目标版本号数字部分


#### 返回

- `bool`：如果完全相等则返回 true


## 🔗 同类方法

- [`VersionNumbers.CompareTo`](/sdk/api/compare-to-versionnumbers)
- [`VersionNumbers.BuildGroupID`](/sdk/api/build-group-id-versionnumbers)
- [`VersionNumbers.Len`](/sdk/api/len-versionnumbers)
- [`VersionNumbers.At`](/sdk/api/at-versionnumbers)
- [`VersionNumbers.String`](/sdk/api/string-versionnumbers)


---

::: details 源码位置
定义于 [`version_numbers.go`](https://github.com/scagogogo/versions-skills/blob/main/version_numbers.go)
:::
