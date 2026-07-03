# ValidateSemver

::: info 方法 · `Version`
```go
func (x *Version) ValidateSemver() error
```
:::

## 📖 说明

ValidateSemver 按照 SemVer 2.0.0 规范严格校验版本号

与 Validate()（仅做基本校验）不同，ValidateSemver 要求：
1. 必须是三段版本号（major.minor.patch）
2. 每段数字不允许前导零
3. 预发布标识和构建元数据必须符合规范字符集


#### 返回

- `error`：如果不符合 semver 规范则返回错误


```go
v := versions.NewVersion("1.2.3")
if err := v.ValidateSemver(); err != nil {
    fmt.Println("不符合 semver 规范:", err)
}
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
