# NewVersion

::: info 函数 · 根包
```go
func NewVersion(versionStr string) *Version
```
:::

## 📖 说明

NewVersion 从版本字符串创建一个新的 Version 对象

该方法解析给定的版本字符串，并返回一个填充了相应字段的 Version 对象。
即使版本字符串格式不正确，该方法也会返回一个对象，但其 IsValid() 方法可能返回 false。


#### 参数

- `versionStr`：要解析的版本号字符串，如 "1.2.3" 或 "v1.2.3-rc1"


#### 返回

- `*Version`：解析后的 Version 对象


```go
version := versions.NewVersion("v1.2.3-beta1")
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
