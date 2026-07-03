# CompareTo

::: info 方法 · `VersionSuffix`
```go
func (x VersionSuffix) CompareTo(target VersionSuffix) int
```
:::

## 📖 说明

CompareTo 比较两个版本后缀的优先级

该方法使用后缀语义权重系统比较两个版本后缀的优先级。
对于已知后缀类型（alpha、beta、rc等），按语义权重排序；
对于未知后缀，退化为字典序比较。

::: tip 权重表
后缀权重值见 [`SuffixWeight` 类型页](/sdk/api/suffix-weight) 的完整权重表。注意 `sp`/`patch`/`post` 权重高于正式版。
:::

#### 参数

- `target`：要比较的目标版本后缀

#### 返回

- `int`：如果当前后缀优先级低于目标后缀，返回 -1；如果相等，返回 0；如果高于，返回 1


```go
suffix1 := versions.VersionSuffix("-alpha1")
suffix2 := versions.VersionSuffix("-beta1")

result := suffix1.CompareTo(suffix2)
if result < 0 {
    fmt.Println("alpha1 后缀的优先级低于 beta1 后缀")
}
```


## 🔗 同类方法

- [`VersionSuffix.IsEmpty`](/sdk/api/is-empty-versionsuffix)
- [`VersionSuffix.String`](/sdk/api/string-versionsuffix)


---

::: details 源码位置
定义于 [`version_suffix.go`](https://github.com/scagogogo/versions-skills/blob/main/version_suffix.go)
:::
