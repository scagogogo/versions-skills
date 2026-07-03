# BuildGroupID

::: info 方法 · `VersionNumbers`
```go
func (x VersionNumbers) BuildGroupID() string
```
:::

## 📖 说明

BuildGroupID 根据版本号数字部分构造版本组的ID

该方法将版本号数字部分以默认分隔符（"."）连接成字符串，用于表示版本组ID。
版本组ID可用于对相同主版本号的版本进行分组和管理。


#### 返回

- `string`：版本组ID字符串，例如 "1.2.3"


```go
numbers := versions.NewVersionNumbers([]int{1, 2, 3})
groupID := numbers.BuildGroupID() // 返回 "1.2.3"

// 可用于版本分组
versionGroups := make(map[string][]*Version)
for _, version := range allVersions {
    groupID := version.VersionNumbers.BuildGroupID()
    versionGroups[groupID] = append(versionGroups[groupID], version)
}
```


## 🔗 同类方法

- [`VersionNumbers.CompareTo`](/sdk/api/compare-to-versionnumbers)
- [`VersionNumbers.Len`](/sdk/api/len-versionnumbers)
- [`VersionNumbers.At`](/sdk/api/at-versionnumbers)
- [`VersionNumbers.String`](/sdk/api/string-versionnumbers)
- [`VersionNumbers.Equals`](/sdk/api/equals-versionnumbers)


---

::: details 源码位置
定义于 [`version_numbers.go`](https://github.com/scagogogo/versions-skills/blob/main/version_numbers.go)
:::
