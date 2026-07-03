# Satisfies

::: info 方法 · `ConstraintUnion`
```go
func (cu *ConstraintUnion) Satisfies(v *Version) bool
```
:::

## 📖 说明

Satisfies 判断版本是否满足约束联合

这是 Match(v) 的语义化别名，与 Version.Satisfies() 对称。


#### 参数

- `v`：要检查的版本对象


#### 返回

- `bool`：如果版本满足任意约束集则返回 true


## 🔗 同类方法

- [`ConstraintUnion.Match`](/sdk/api/match-constraintunion)
- [`ConstraintUnion.String`](/sdk/api/string-constraintunion)


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
