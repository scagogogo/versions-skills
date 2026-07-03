# Segments64

::: info 方法 · `Version`
```go
func (x *Version) Segments64() []int64
```
:::

## 📖 说明

Segments64 返回版本号数字段的 int64 数组

与 Segments() 相同，但返回 int64 类型，
适用于需要大数值范围的场景。


#### 返回

- `[]int64`：版本号数字段 int64 数组


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
