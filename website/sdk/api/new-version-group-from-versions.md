# NewVersionGroupFromVersions

::: info 函数 · 根包
```go
func NewVersionGroupFromVersions(versions []*Version) *VersionGroup
```
:::

## 📖 说明

NewVersionGroupFromVersions 从版本数组创建一个版本组

该方法基于给定的版本数组创建一个版本组。所有版本将被添加到同一个组中，
该组的数字部分取自第一个版本的数字部分。


#### 参数

- `versions`：要添加到组中的版本数组


#### 返回

- `*VersionGroup`：新创建的包含所有指定版本的版本组，如果输入为空则返回 nil


```go
vs := versions.NewVersions("1.2.0", "1.2.1", "1.2.2")
group := versions.NewVersionGroupFromVersions(vs)
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
