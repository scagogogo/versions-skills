# MustParse

::: info 函数 · 根包
```go
func MustParse(versionStr string) *Version
```
:::

## 📖 说明

MustParse 解析版本号字符串，如果解析失败则 panic

这是 NewVersionE() 的 panic 变体，适用于初始化时确定版本号合法的场景，
类似于 regexp.MustCompile 的模式。在版本号来自硬编码或测试数据的场景下非常有用。


#### 参数

- `versionStr`：版本号字符串


#### 返回

- `*Version`：解析后的版本对象


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
