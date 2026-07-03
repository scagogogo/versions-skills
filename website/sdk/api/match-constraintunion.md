# Match

::: info 方法 · `ConstraintUnion`
```go
func (cu *ConstraintUnion) Match(v *Version) bool
```
:::

## 📖 说明

Match 判断版本是否满足约束联合（OR 逻辑）

只要版本满足任意一个 ConstraintSet 即返回 true。


#### 参数

- `v`：要检查的版本对象


#### 返回

- `bool`：如果版本满足任意约束集则返回 true


## 🔗 同类方法

- [`ConstraintUnion.Satisfies`](/sdk/api/satisfies-constraintunion)
- [`ConstraintUnion.String`](/sdk/api/string-constraintunion)


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
