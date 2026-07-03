# Satisfies

::: info 方法 · `Version`
```go
func (x *Version) Satisfies(constraint *Constraint) bool
```
:::

## 📖 说明

Satisfies 判断版本是否满足给定的约束条件

这是 Constraint.Match(v) 的反向调用方式，语义更自然：
v.Satisfies(constraint) 等价于 constraint.Match(v)。


#### 参数

- `constraint`：版本约束条件


#### 返回

- `bool`：如果版本满足约束则返回 true


```go
c, _ := versions.ParseConstraint(">=1.0.0")
v := versions.NewVersion("1.5.0")
if v.Satisfies(c) {
fmt.Println("版本满足约束")
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
