# Constraint

::: info 类型 · 根包
```go
type Constraint struct {
	// Operator 约束操作符
	Operator ConstraintOperator

	// Version 约束目标版本
	Version *Version
}
```
:::

## 📖 说明

Constraint 表示一个版本约束条件

Constraint 用于判断某个版本是否满足指定的约束条件，
如 ">=1.0.0", "^1.2.3", "~1.2.3", "1.x" 等。


```go
c, err := versions.ParseConstraint(">=1.0.0")
if err != nil {
    log.Fatal(err)
}
v := versions.NewVersion("1.5.0")
if c.Match(v) {
    fmt.Println("1.5.0 satisfies >=1.0.0")
}
```


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
