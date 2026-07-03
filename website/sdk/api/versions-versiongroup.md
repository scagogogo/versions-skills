# Versions

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) Versions() []*Version
```
:::

## 📖 说明

Versions 返回组下的所有版本

该方法返回版本组中包含的所有版本的数组，结果不保证有序。


#### 返回

- `[]*Version`：版本组中所有版本的数组


```go
group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.0", "1.2.1"))
allVersions := group.Versions()
fmt.Printf("组中包含 %d 个版本\n", len(allVersions))
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
