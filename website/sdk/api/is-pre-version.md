# IsPre

::: info 方法 · `Version`
```go
func (x *Version) IsPre() bool
```
:::

## 📖 说明

IsPre 判断版本是否为预发布版（pre 标识）

预发布版是指后缀包含 pre 标识的版本，如 "1.0.0-pre1"。
注意：这是显式的 pre 后缀，与 IsPrerelease()（判断是否有任何后缀）不同。


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
