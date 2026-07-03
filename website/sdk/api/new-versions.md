# NewVersions

::: info 函数 · 根包
```go
func NewVersions(versionStringSlice ...string) []*Version
```
:::

## 📖 说明

NewVersions 批量创建多个 Version 对象

该方法接受多个版本字符串，并返回相应的 Version 对象数组。


#### 参数

- `versionStringSlice`：一个或多个版本号字符串


#### 返回

- `[]*Version`：解析后的 Version 对象数组


```go
versions := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")
for _, v := range versions {
    fmt.Println(v.Raw)
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
