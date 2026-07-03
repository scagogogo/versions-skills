# IsDigit

::: info 方法 · `VersionStringParser`
```go
func (x *VersionStringParser) IsDigit(c rune) bool
```
:::

## 📖 说明

IsDigit 判断是否是数字

该方法检查给定的字符是否为数字字符（0-9）。


#### 参数

- `c`：要检查的字符


#### 返回

- `bool`：如果是数字则返回 true，否则返回 false


## 🔗 同类方法

- [`VersionStringParser.Parse`](/sdk/api/parse-versionstringparser)
- [`VersionStringParser.IsVersionNumberDelimiter`](/sdk/api/is-version-number-delimiter-versionstringparser)


---

::: details 源码位置
定义于 [`parser.go`](https://github.com/scagogogo/versions-skills/blob/main/parser.go)
:::
