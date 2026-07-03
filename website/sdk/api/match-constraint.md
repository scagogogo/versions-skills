# Match

::: info 方法 · `Constraint`
```go
func (c *Constraint) Match(v *Version) bool
```
:::

## 📖 说明

Match 判断版本是否满足约束条件


#### 参数

- `v`：要检查的版本对象


#### 返回

- `bool`：如果版本满足约束则返回 true


---

::: details 源码位置
定义于 [`constraint.go`](https://github.com/scagogogo/versions-skills/blob/main/constraint.go)
:::
