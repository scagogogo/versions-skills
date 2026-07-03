# FilterByConstraint

::: info 函数 · 根包
```go
func FilterByConstraint(versions []*Version, constraint *Constraint) []*Version
```
:::

## 📖 说明

FilterByConstraint 根据约束条件过滤版本列表

返回所有满足约束条件的版本。


#### 参数

- `versions`：版本对象列表
- `constraint`：版本约束条件


#### 返回

- `[]*Version`：满足约束的版本列表


---

::: details 源码位置
定义于 [`version_utils.go`](https://github.com/scagogogo/versions-skills/blob/main/version_utils.go)
:::
