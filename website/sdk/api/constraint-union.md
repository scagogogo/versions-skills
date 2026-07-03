# ConstraintUnion

::: info 类型 · 根包
```go
type ConstraintUnion struct {
	// Sets AND 约束集合列表，之间是 OR 关系
	Sets []*ConstraintSet
}
```
:::

## 📖 说明

ConstraintUnion 表示一组 OR 组合的约束集合

每个 ConstraintSet 内部是 AND 逻辑，多个 ConstraintSet 之间是 OR 逻辑。
例如 ">=1.0.0,<2.0.0 || >=3.0.0" 表示版本必须满足 (>=1.0.0 AND <2.0.0) OR (>=3.0.0)。


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
