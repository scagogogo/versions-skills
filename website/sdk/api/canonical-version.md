# Canonical

::: info 方法 · `Version`
```go
func (x *Version) Canonical() string
```
:::

## 📖 说明

Canonical 返回版本的规范字符串表示

规范格式为：[前缀]主版本号.次版本号.修订版本号[-后缀][+元数据]
始终输出三段版本号，不足的补零。


#### 返回

- `string`：规范化的版本字符串


```go
v := versions.NewVersion("1.2")
fmt.Println(v.Canonical()) // "1.2.0"

v2 := versions.NewVersion("v1.2.3-beta+build.1")
fmt.Println(v2.Canonical()) // "v1.2.3-beta+build.1"
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
