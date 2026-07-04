# NewVersionGroup

::: info 函数 · 根包
```go
func NewVersionGroup(groupVersionNumbers VersionNumbers) *VersionGroup
```
:::

## 📖 说明

NewVersionGroup 创建一个新的版本组

该方法根据指定的版本号数字部分创建一个新的版本组。创建时需要传递能够在包范围下唯一区分组的组ID，
这个ID选择的是版本中的数字部分。


#### 参数

- `groupVersionNumbers`：版本组的数字部分，如 [1,2] 表示 "1.2" 版本组


#### 返回

- `*VersionGroup`：新创建的版本组对象


```go
// 创建表示 "1.2" 版本组的对象
group := versions.NewVersionGroup(vs.NewVersionNumbers([]int{1, 2}))
```


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
