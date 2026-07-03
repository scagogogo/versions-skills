# String

::: info 方法 · `VersionSuffix`
```go
func (x VersionSuffix) String() string
```
:::

## 📖 说明

String 返回后缀的字符串表示

实现 `fmt.Stringer` 接口，返回版本后缀的原始字符串形式。

#### 返回

- `string`：后缀字符串，如 `"-beta1"`、`"-rc1"`；空后缀返回 `""`


```go
suffix := versions.VersionSuffix("-beta1")
fmt.Println(suffix.String()) // -beta1

empty := versions.EmptyVersionSuffix
fmt.Println(empty.String() == "") // true
```


## 🔗 同类方法

- [`VersionSuffix.IsEmpty`](/sdk/api/is-empty-versionsuffix)
- [`VersionSuffix.CompareTo`](/sdk/api/compare-to-versionsuffix)


---

::: details 源码位置
定义于 [`version_suffix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_suffix.go)
:::
