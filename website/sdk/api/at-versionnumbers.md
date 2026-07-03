# At

::: info 方法 · `VersionNumbers`
```go
func (x VersionNumbers) At(index int) int
```
:::

## 📖 说明

At 返回指定索引位置的版本号数字

如果索引越界则返回 0。


#### 参数

- `index`：从 0 开始的索引


#### 返回

- `int`：对应位置的版本号数字，越界返回 0


---

::: details 源码位置
定义于 [`version_numbers.go`](https://github.com/scagogogo/versions-skills/blob/main/version_numbers.go)
:::
