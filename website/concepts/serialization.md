# 序列化

::: tip 关键
`Version` 实现了 **JSON / Text / SQL** 三类序列化接口，可直接用于配置、API 传输与数据库。
:::

## 📦 JSON

```go
v := versions.NewVersion("v1.2.3-beta1")

data, _ := json.Marshal(v)
// {"raw":"v1.2.3-beta1","public_time":"0001-01-01T00:00:00Z","version_numbers":[1,2,3],"prefix":"v","suffix":"-beta1"}

var v2 versions.Version
json.Unmarshal(data, &v2) // 还原
```

字段 tag：`json:"raw"` / `json:"version_numbers"` / `json:"prefix"` / `json:"suffix"` / `json:"metadata,omitempty"` / `json:"public_time"`。

## 📝 Text（encoding.TextMarshaler）

```go
data, _ := v.MarshalText()   // []byte("v1.2.3-beta1")
v2 := &versions.Version{}
v2.UnmarshalText(data)
```

适合 YAML / TOML / 环境变量场景。

## 🗄 SQL（database/sql）

`Version` 同时实现 `driver.Valuer` 与 `sql.Scanner`，可直接作为字段类型存入数据库：

```go
type Release struct {
    Version versions.Version
    Name    string
}

// 插入：Version.Value() 返回 Raw 字符串
db.Exec("INSERT INTO releases(version, name) VALUES(?, ?)", r.Version, r.Name)

// 查询：Version.Scan() 从 string/[]byte 还原
var rel Release
db.QueryRow("SELECT version, name FROM releases WHERE id=?", id).Scan(&rel.Version, &rel.Name)
```

::: warning 注意
`PublicTime` 字段是 `time.Time`，JSON 序列化为 RFC3339。若未设置，输出零值时间。
:::

## 📚 延伸

- API：[`MarshalJSON`](/sdk/api/marshal-json-version) · [`UnmarshalJSON`](/sdk/api/unmarshal-json-version) · [`MarshalText`](/sdk/api/marshal-text-version) · [`Scan`](/sdk/api/scan-version) · [`Value`](/sdk/api/value-version)
- 概念：[版本号结构](/concepts/version-anatomy)
