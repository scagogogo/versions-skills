# Validate

::: info 方法 · `Version`
```go
func (x *Version) Validate() error
```
:::

## 📖 说明

Validate 严格校验版本号格式

与 IsValid()（仅检查是否有版本号数字）不同，Validate 执行更严格的校验：
1. 版本号数字部分不能为空
2. 每个版本号数字必须 >= 0


#### 返回

- `error`：如果版本号不符合严格格式要求则返回错误


---

::: details 源码位置
定义于 [`version.go`](https://github.com/scagogogo/versions-skills/blob/main/version.go)
:::
