# VersionGroup

::: info 类型 · 根包
```go
type VersionGroup struct {

	// GroupVersionNumbers 组的版本号中的数字部分
	// 例如对于版本号 "1.2.x"，GroupVersionNumbers 为 [1,2]
	GroupVersionNumbers VersionNumbers

	// VersionMap 组中包含的所有版本，键为版本的原始字符串，值为版本对象
	VersionMap map[string]*Version
}
```
:::

## 📖 说明

VersionGroup 表示一个版本组，一个版本组中可能有一个或多个版本

VersionGroup 用于管理具有相同主版本号数字部分的一组版本。它提供了版本的添加、查询、排序和范围查询等功能。
版本组是版本管理系统中的重要概念，用于将相似版本聚合在一起，便于进行版本管理和分析。

每个版本组都有一个唯一的ID，由其主版本号数字部分生成，如 "1.2"。


```go
// 创建一个新的版本组
group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))

// 添加版本到组中
v1 := versions.NewVersion("1.2.0")
v2 := versions.NewVersion("1.2.1")
group.Add(v1)
group.Add(v2)

// 获取组中的所有版本
allVersions := group.Versions()

// 获取排序后的版本
sortedVersions := group.SortVersions()
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
