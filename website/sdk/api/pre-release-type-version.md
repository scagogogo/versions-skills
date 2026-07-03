# PreReleaseType

::: info 方法 · `Version`
```go
func (x *Version) PreReleaseType() string
```
:::

## 📖 说明

PreReleaseType 返回预发布版本的类型标识

返回后缀的语义权重类型名称字符串，如 "alpha"、"beta"、"rc" 等。
如果不是预发布版本则返回空字符串。


#### 返回

- `string`：预发布类型名称


```go
v := versions.NewVersion("1.0.0-beta2")
fmt.Println(v.PreReleaseType()) // "beta"

s := versions.NewVersion("1.0.0")
fmt.Println(s.PreReleaseType()) // ""
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
