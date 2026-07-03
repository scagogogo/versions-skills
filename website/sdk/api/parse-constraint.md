# ParseConstraint

::: info 函数 · 根包
```go
func ParseConstraint(expr string) (*Constraint, error)
```
:::

## 📖 说明

ParseConstraint 解析单个版本约束表达式

支持的操作符: =, !=, >, >=, <, <=, ^, ~
支持的通配符: x, X, * (如 1.x, 1.2.*)


#### 参数

- `expr`：约束表达式，如 ">=1.0.0", "^1.2.3", "~1.2"


#### 返回

- `*Constraint`：解析后的约束对象
- `error`：如果表达式格式错误则返回错误


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
