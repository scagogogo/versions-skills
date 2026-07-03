# Core

::: info 方法 · `Version`
```go
func (x *Version) Core() *Version
```
:::

## 📖 说明

Core 返回版本的核心部分（去除后缀）

返回一个新 Version 对象，只保留前缀和版本号数字部分，去除所有后缀。
这是获取版本"纯净"数字部分的快捷方式。


#### 返回

- `*Version`：去除后缀后的核心版本对象


```go
v := versions.NewVersion("1.2.3-beta1")
core := v.Core()
fmt.Println(core.RawString()) // 输出: "1.2.3"
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
