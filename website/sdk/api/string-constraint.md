# String

::: info 方法 · `Constraint`
```go
func (c *Constraint) String() string
```
:::

## 📖 说明

String 返回约束条件的字符串表示

将约束条件序列化为可解析的字符串格式，如 ">=1.0.0"、"^1.2.3"、"~1.2"。
对于通配符约束（1.x），返回原始版本字符串形式。


#### 返回

- `string`：约束条件的字符串表示


```go
c, _ := versions.ParseConstraint(">=1.0.0")
fmt.Println(c.String()) // 输出: ">=1.0.0"
```


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
