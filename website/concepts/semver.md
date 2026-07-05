# SemVer 规范

::: tip 关键
versions-skills 支持 [SemVer 2.0.0](https://semver.org/) 严格校验，同时兼容更宽松的「版本号」概念。
:::

## 📐 严格 SemVer

`IsSemver()` 用正则严格判定，要求：

- **三段**：`MAJOR.MINOR.PATCH`
- **无前导零**：`01.02.03` 非法
- 可选 `v` 前缀
- 可选预发布段：`-alpha.1`、`-rc.2`
- 可选构建元数据：`+build.7`

```go
versions.NewVersion("1.2.3").IsSemver()               // true
versions.NewVersion("1.2.3-alpha.1+build.7").IsSemver() // true
versions.NewVersion("1.2").IsSemver()                 // false（不够3段）
versions.NewVersion("01.2.3").IsSemver()              // false（前导零）
```

## 🔍 Validate vs ValidateSemver

| 方法 | 严格度 | 要求 |
|:--|:--|:--|
| `Validate()` | 基本校验 | 有数字部分且非负 |
| `ValidateSemver()` | 严格 SemVer | 三段 + 无前导零 + 字符符合规 |

## 🆚 宽松版本号

versions-skills 的 `Version` 比 SemVer 宽松——`1.2`、`v1.2.3.4`、`1.0.0-beta1` 都能解析。只有需要 SemVer 合规性时才用 `IsSemver` / `ValidateSemver`。

```mermaid
flowchart TB
  INPUT["任意字符串<br/>1.2.3 / v1.2 / 1.0.0-beta1 / 01.2.3 / 1.2.3.4"]

  INPUT --> PARSE["NewVersion（宽松解析）"]
  PARSE --> V["Version 对象<br/>总能解析（非法则 IsValid=false）"]

  V -->|"Validate()"| BASIC["基本校验<br/>有数字段且非负"]
  V -->|"IsSemver() / ValidateSemver()"| STRICT["严格 SemVer<br/>三段 + 无前导零 + 合规字符"]

  BASIC --> B1["1.2 → ✅<br/>v1.2.3.4 → ✅"]
  STRICT --> S1["1.2 → ❌ 不够3段<br/>01.2.3 → ❌ 前导零<br/>1.2.3.4 → ❌ 超三段"]

  style INPUT fill:#f8fafc,stroke:#475569
  style PARSE fill:#eff6ff,stroke:#2563eb
  style V fill:#eff6ff,stroke:#2563eb,stroke-width:3px
  style BASIC fill:#fff7ed,stroke:#ea580c
  style STRICT fill:#fff7ed,stroke:#ea580c
  style B1 fill:#f0fdf4,stroke:#16a34a
  style S1 fill:#fef2f2,stroke:#dc2626
```

## 📚 延伸

- API：[`IsSemver`](/sdk/api/is-semver-version) · [`ValidateSemver`](/sdk/api/validate-semver-version) · [SDK 索引](/sdk/)
- 概念：[版本号结构](/concepts/version-anatomy) · [后缀权重](/concepts/suffix-weight)
- CLI：[`check --semver`](/cli/commands/check) · [`validate`](/cli/commands/validate)
