# NegateConstraint

::: info 函数 · 根包
```go
func NegateConstraint(c *Constraint) *Constraint
```
:::

## 📖 说明

NegateConstraint 返回约束条件的否定形式

例如 >=1.0.0 的否定为 <1.0.0，=1.0.0 的否定为 !=1.0.0。


#### 参数

- `c`：要否定的约束条件


#### 返回

- `*Constraint`：否定后的约束条件


```go
c, _ := versions.ParseConstraint(">=1.0.0")
neg := versions.NegateConstraint(c)
fmt.Println(neg.String()) // "<1.0.0"
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
