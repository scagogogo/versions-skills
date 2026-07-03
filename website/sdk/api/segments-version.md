# Segments

::: info 方法 · `Version`
```go
func (x *Version) Segments() []int
```
:::

## 📖 说明

Segments 返回版本号数字段的整数数组

等价于直接访问 VersionNumbers，但返回 []int 类型，
便于与不使用 VersionNumbers 类型的代码交互。


#### 返回

- `[]int`：版本号数字段数组


```go
v := versions.NewVersion("1.2.3.4")
segments := v.Segments()
// segments == []int{1, 2, 3, 4}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
