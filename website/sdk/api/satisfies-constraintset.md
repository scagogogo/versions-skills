# Satisfies

::: info 方法 · `ConstraintSet`
```go
func (cs *ConstraintSet) Satisfies(v *Version) bool
```
:::

## 📖 说明

Satisfies 判断版本是否满足约束集合

这是 Match(v) 的语义化别名，使调用方式更自然：
cs.Satisfies(v) 等价于 cs.Match(v)，与 Version.Satisfies(constraint) 对称。


#### 参数

- `v`：要检查的版本对象


#### 返回

- `bool`：如果版本满足所有约束则返回 true


## 🔗 同类方法

- [`ConstraintSet.Match`](/sdk/api/match-constraintset)
- [`ConstraintSet.String`](/sdk/api/string-constraintset)
- [`ConstraintSet.Len`](/sdk/api/len-constraintset)


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
