# BuildGroupID

::: info 方法 · `Version`
```go
func (x *Version) BuildGroupID() string
```
:::

## 📖 说明

BuildGroupID 构造版本所属的组的ID

该方法根据版本号的数字部分生成一个组ID，用于将相似版本分组。


#### 返回

- `string`：表示版本组的ID字符串


```go
version := versions.NewVersion("1.2.3")
groupID := version.BuildGroupID()
fmt.Printf("版本组ID: %s\n", groupID)
```


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
