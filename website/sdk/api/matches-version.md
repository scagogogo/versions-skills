# Matches

::: info 方法 · `Version`
```go
func (x *Version) Matches(expr string) (bool, error)
```
:::

## 📖 说明

Matches 判断版本是否满足约束表达式字符串

该方法是 ParseConstraintSet + Match 的便捷组合，
适用于需要从字符串快速判断版本是否满足约束的场景。


#### 参数

- `expr`：约束表达式字符串，如 ">=1.0.0"


#### 返回

- `bool`：如果版本满足约束则返回 true
- `error`：如果约束表达式解析失败则返回错误


```go
v := versions.NewVersion("1.5.0")
ok, err := v.Matches(">=1.0.0,<2.0.0")
if ok {
fmt.Println("版本在范围内")
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
