# IsRelease

::: info 方法 · `Version`
```go
func (x *Version) IsRelease() bool
```
:::

## 📖 说明

IsRelease 判断版本是否带有 release 后缀

带 release 后缀的版本如 "1.0.0-release"。
注意：这与 IsStable()（无后缀）不同，IsRelease 检测的是显式的 "-release" 后缀。


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
