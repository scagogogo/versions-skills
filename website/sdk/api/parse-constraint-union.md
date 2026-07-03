# ParseConstraintUnion

::: info 函数 · 根包
```go
func ParseConstraintUnion(expr string) (*ConstraintUnion, error)
```
:::

## 📖 说明

ParseConstraintUnion 解析包含 OR 逻辑的约束表达式

支持格式: ">=1.0.0,<2.0.0 || >=3.0.0"，其中逗号分隔为 AND，|| 分隔为 OR。
也支持不包含 || 的简单表达式，此时等价于 ParseConstraintSet。


#### 参数

- `expr`：约束表达式


#### 返回

- `*ConstraintUnion`：解析后的约束联合
- `error`：如果表达式格式错误则返回错误


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
