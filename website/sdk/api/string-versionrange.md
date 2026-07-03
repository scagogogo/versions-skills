# String

::: info 方法 · `VersionRange`
```go
func (r *VersionRange) String() string
```
:::

## 📖 说明

String 返回版本范围的字符串表示

闭区间使用 []，开区间使用 ()。


#### 返回

- `string`：如 "[1.0.0, 2.0.0]"、"(1.0.0, 2.0.0)"等


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
