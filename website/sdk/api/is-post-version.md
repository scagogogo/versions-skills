# IsPost

::: info 方法 · `Version`
```go
func (x *Version) IsPost() bool
```
:::

## 📖 说明

IsPost 判断版本是否为 Post 发布版

Post 发布版是指后缀包含 post 标识的版本（PEP 440 规范），如 "1.0.0-post1"。


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
