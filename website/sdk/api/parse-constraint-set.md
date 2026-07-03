# ParseConstraintSet

::: info 函数 · 根包
```go
func ParseConstraintSet(expr string) (*ConstraintSet, error)
```
:::

## 📖 说明

ParseConstraintSet 解析逗号分隔的 AND 组合约束

支持格式: ">=1.0.0,<2.0.0", "^1.2.3", "~1.2"


#### 参数

- `expr`：逗号分隔的约束表达式


#### 返回

- `*ConstraintSet`：解析后的约束集合
- `error`：如果任何子表达式格式错误则返回错误


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
