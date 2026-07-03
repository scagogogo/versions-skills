# IsPrerelease

::: info 方法 · `Version`
```go
func (x *Version) IsPrerelease() bool
```
:::

## 📖 说明

IsPrerelease 判断版本是否为预发布版本

预发布版本是指带有后缀（如 -alpha, -beta, -rc, -SNAPSHOT 等）的版本。
正式版本（无后缀）返回 false。


#### 返回

- `bool`：如果是预发布版本则返回 true


```go
v := versions.NewVersion("1.0.0-beta")
if v.IsPrerelease() {
    fmt.Println("这是预发布版本")
}
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
