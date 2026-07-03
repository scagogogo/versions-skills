# IsStable

::: info 方法 · `Version`
```go
func (x *Version) IsStable() bool
```
:::

## 📖 说明

IsStable 判断版本是否为正式稳定版本

正式稳定版本是指不带任何后缀的版本，如 "1.0.0"。


#### 返回

- `bool`：如果是正式稳定版本则返回 true


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
