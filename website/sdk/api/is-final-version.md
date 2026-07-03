# IsFinal

::: info 方法 · `Version`
```go
func (x *Version) IsFinal() bool
```
:::

## 📖 说明

IsFinal 判断版本是否为 Final 版本

Final 版本是指后缀包含 final 标识的版本（Maven 生态常见），如 "1.0.0-final"。
注意：Final 后缀与无后缀的正式版语义相同，但 IsFinal 专用于检测显式的 final 后缀。


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
