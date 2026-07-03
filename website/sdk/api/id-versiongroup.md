# ID

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) ID() string
```
:::

## 📖 说明

ID 返回组的ID

该方法返回版本组的唯一标识符，由其数字部分生成。


#### 返回

- `string`：版本组的ID，例如 "1.2"


```go
group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
groupID := group.ID() // 返回 "1.2"
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
