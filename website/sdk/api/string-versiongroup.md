# String

::: info 方法 · `VersionGroup`
```go
func (x *VersionGroup) String() string
```
:::

## 📖 说明

String 返回版本组的字符串表示

实现 fmt.Stringer 接口。
格式为 "组ID (版本数)"，如 "1.2 (5个版本)"。


#### 返回

- `string`：版本组的字符串表示


---

::: details 源码位置
定义于 [`version_group.go`](https://github.com/scagogogo/versions-skills/blob/main/version_group.go)
:::
