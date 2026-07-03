# IsVersionNumberDelimiter

::: info 方法 · `VersionStringParser`
```go
func (x *VersionStringParser) IsVersionNumberDelimiter(c rune) bool
```
:::

## 📖 说明

IsVersionNumberDelimiter 判断是否是版本数字的分隔符

该方法检查给定的字符是否为版本号数字部分的分隔符（目前仅支持点号）。


#### 参数

- `c`：要检查的字符


#### 返回

- `bool`：如果是分隔符则返回 true，否则返回 false


---

::: details 源码位置
定义于 [`parser.go`](https://github.com/scagogogo/versions-skills/blob/main/parser.go)
:::
