# String

::: info 方法 · `ConstraintSet`
```go
func (cs *ConstraintSet) String() string
```
:::

## 📖 说明

String 返回约束集合的字符串表示

将约束集合序列化为逗号分隔的字符串格式，如 ">=1.0.0,<2.0.0"。


#### 返回

- `string`：约束集合的字符串表示


```go
cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
fmt.Println(cs.String()) // 输出: ">=1.0.0,<2.0.0"
```


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
