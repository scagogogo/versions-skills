# IsSemver

::: info 方法 · `Version`
```go
func (x *Version) IsSemver() bool
```
:::

## 📖 说明

IsSemver 判断版本字符串是否符合 SemVer 2.0.0 规范

严格的 semver 格式要求：主版本号.次版本号.修订版本号，
可选的预发布标识（以 - 分隔）和构建元数据（以 + 分隔）。
不允许前导零（如 01.02.03）。


#### 返回

- `bool`：如果符合 semver 规范则返回 true


```go
v := versions.NewVersion("1.2.3")
fmt.Println(v.IsSemver()) // true

v2 := versions.NewVersion("1.2.3-alpha.1+build.123")
fmt.Println(v2.IsSemver()) // true

v3 := versions.NewVersion("1.2")
fmt.Println(v3.IsSemver()) // false（不够3段）
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
