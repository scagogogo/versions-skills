# ContainsVersion

::: info 函数 · 根包
```go
func ContainsVersion(versions []*Version, target *Version) bool
```
:::

## 📖 说明

ContainsVersion 判断版本列表中是否包含指定版本

根据 Raw 字段判断版本是否相同。


#### 参数

- `versions`：版本对象列表
- `target`：目标版本对象


#### 返回

- `bool`：如果列表中包含目标版本则返回 true


```go
list := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")
v := versions.NewVersion("1.1.0")
fmt.Println(versions.ContainsVersion(list, v)) // true
```


---

::: details 源码位置
定义于 [`version_range.go`](https://github.com/scagogogo/versions-skills/blob/main/version_range.go)
:::
