# IsBetween

::: info 方法 · `Version`
```go
func (x *Version) IsBetween(low, high *Version) bool
```
:::

## 📖 说明

IsBetween 判断当前版本是否在两个版本之间（包含边界）

如果 low <= x <= high 则返回 true。
如果 low 或 high 为 nil，则忽略对应的边界检查。


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
