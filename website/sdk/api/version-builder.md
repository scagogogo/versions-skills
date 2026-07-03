# VersionBuilder

::: info 类型 · 根包
```go
type VersionBuilder struct {
	prefix     string
	numbers    []int
	suffix     string
	metadata   string
	publicTime time.Time
}
```
:::

## 📖 说明

VersionBuilder 提供流式 API 构建版本对象

VersionBuilder 允许通过方法链的方式逐步构建版本对象，
适用于需要程序化生成版本号的场景。


```go
v := versions.NewVersionBuilder().
Prefix("v").
Major(1).
Minor(2).
Patch(3).
Suffix("-beta1").
Build()
// v.Raw == "v1.2.3-beta1"
```


---

::: details 源码位置
定义于 [`version_builder.go`](https://github.com/scagogogo/versions-skills/blob/main/version_builder.go)
:::
