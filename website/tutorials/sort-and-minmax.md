# 排序与极值

给版本号列表排序，找出最新/最旧、最新稳定/预发布版本。

## 📊 排序

```go
vs := versions.NewVersions("1.10.0", "1.2.0", "1.1.0", "1.0.0")

asc := versions.SortVersionSlice(vs)                          // 升序
desc := versions.SortVersionSlice(vs)                         // 原地已排
// 降序可用 VersionSlice
slice := versions.VersionSlice(asc)
slice.Sort()
// 反转
for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
    slice[i], slice[j] = slice[j], slice[i]
}
```

字符串版（保留原始书写）：

```go
strs := versions.SortVersionStringSlice([]string{"v1.10.0", "1.2.0"})
// ["1.2.0" "v1.10.0"]
```

## 🏆 极值

```go
vs := versions.NewVersions("1.0.0", "1.0.0-rc1", "1.1.0-beta1", "2.0.0")

versions.Max(vs).Raw               // 2.0.0（最新，含预发布）
versions.Min(vs).Raw               // 1.0.0-rc1（最旧）
versions.LatestStable(vs).Raw      // 2.0.0（最新稳定）
versions.LatestPrerelease(vs).Raw  // 1.1.0-beta1（最新预发布）
```

## 🎚 过滤

```go
stable := versions.FilterByStable(vs)       // 仅稳定版
pre := versions.FilterByPrerelease(vs)      // 仅预发布
maj2 := versions.FilterByMajor(vs, 2)       // 仅 Major=2
```

## 🚀 下一步

- [分组与聚合](/tutorials/grouping)
- API：[`SortVersionSlice`](/sdk/api/sort-version-slice) · [`Max`](/sdk/api/max) · [`LatestStable`](/sdk/api/latest-stable) · [`FilterByStable`](/sdk/api/filter-by-stable)
- CLI：[`sort`](/cli/commands/sort) · [`max`](/cli/commands/max) · [`latest-stable`](/cli/commands/latest-stable)
