# Group

::: info 函数 · 根包
```go
func Group(versions []*Version) map[string]*VersionGroup
```
:::

## 📖 说明

Group 对版本号进行分组

该函数将版本对象数组按照其数字部分（主版本号）进行分组，为每个不同的主版本号创建一个版本组。
分组的依据是版本号的数字部分生成的组ID，如 "1.2.3" 和 "1.2.4" 会被分到同一组 "1.2"。

分组可用于：
1. 对特定主版本系列进行管理和查询
2. 分析版本演进路径
3. 查找某主版本下的最新/最旧版本


#### 参数

- `versions`：需要分组的版本对象数组


#### 返回

- `map[string]*VersionGroup`：分组结果，键为版本组ID，值为对应的版本组对象


```go
allVersions := versions.NewVersions("1.0.0", "1.1.0", "1.2.0", "2.0.0", "2.1.0")
groups := versions.Group(allVersions)

// 遍历所有版本组
for groupID, group := range groups {
    fmt.Printf("版本组 %s 包含 %d 个版本\n", groupID, len(group.Versions))

    // 获取组内最新版本
    latestVersion := group.GetLatestVersion()
    fmt.Printf("组内最新版本: %s\n", latestVersion.Raw)
}
```


---

::: details 源码位置
定义于 [`group.go`](https://github.com/scagogogo/versions-skills/blob/main/group.go)
:::
