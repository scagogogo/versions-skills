# ConstraintSet

::: info 类型 · 根包
```go
type ConstraintSet struct {
	Constraints []Constraint
}
```
:::

## 📖 说明

ConstraintSet 表示一组 AND 组合的约束条件

多个约束条件之间是 AND 关系，所有条件都必须满足。
例如 ">=1.0.0,<2.0.0" 表示版本必须同时满足 >=1.0.0 和 <2.0.0。


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
