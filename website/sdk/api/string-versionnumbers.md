# String

::: info 方法 · `VersionNumbers`
```go
func (x VersionNumbers) String() string
```
:::

## 📖 说明

String 返回版本号数字部分的字符串表示

等价于 BuildGroupID()，提供更符合 Go 惯例的 String() 方法。


#### 返回

- `string`：如 "1.2.3"


---

::: details 源码位置
定义于 [`version_numbers.go`](https://github.com/scagogogo/versions-skills/blob/main/version_numbers.go)
:::
