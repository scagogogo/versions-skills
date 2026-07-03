# QueryRangeVersions

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version
```
:::

## 📖 说明

QueryRangeVersions 获取组内指定区间内的版本

该方法根据给定的起始和结束版本范围，返回组内符合条件的版本数组。
版本的包含性由 ContainsPolicy 参数控制。


#### 参数

- `start`：包含起始版本和包含策略的元组
- `end`：包含结束版本和包含策略的元组


#### 返回

- `[]*Version`：符合区间条件的版本数组


```go
group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.0", "1.2.1", "1.2.2", "1.2.3"))

// 查询 1.2.0（包含）到 1.2.2（包含）的版本
startTuple := tuple.NewTuple2(versions.NewVersion("1.2.0"), versions.ContainsPolicyYes)
endTuple := tuple.NewTuple2(versions.NewVersion("1.2.2"), versions.ContainsPolicyYes)

rangeVersions := group.QueryRangeVersions(startTuple, endTuple)
// 结果包含: ["1.2.0", "1.2.1", "1.2.2"]
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
