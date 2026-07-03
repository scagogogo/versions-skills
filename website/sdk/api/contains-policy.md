# ContainsPolicy

::: info 类型 · 根包
```go
type ContainsPolicy int
```
:::

## 📖 说明

ContainsPolicy 用于控制版本查询时的包含/排除策略

该类型定义了在版本过滤或查询操作中，是否应包含或排除特定版本的策略选项。
它作为枚举类型使用，提供了三种可能的状态：未指定、包含和排除。


```go
// 使用包含策略进行版本过滤
filter := &VersionFilter{
Contains: "beta",
ContainsPolicy: ContainsPolicyYes,
}

// 使用排除策略进行版本过滤
filter := &VersionFilter{
Contains: "snapshot",
ContainsPolicy: ContainsPolicyNo,
}
```


---

::: details 源码位置
定义于 [`contains_policy.go`](https://github.com/scagogogo/versions-skills/blob/main/contains_policy.go)
:::
