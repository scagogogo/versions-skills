# 约束表达式实战

用约束表达式判断「某版本是否兼容 / 在允许范围内」，这是依赖解析的核心。

## 🔢 基本运算符

```go
v := versions.NewVersion("1.5.0")

c, _ := versions.ParseConstraint(">=1.0.0")
c.Match(v) // true

// Version 便捷方法
ok, _ := v.Matches(">=1.0.0,<2.0.0") // true
```

9 种运算符见 [约束表达式概念](/concepts/constraints)。

## 🤝 caret / tilde（兼容版本）

```go
// ^ 锁定左起首个非零段
cs, _ := versions.ParseConstraintSet("^1.2.3")
cs.Match(versions.NewVersion("1.3.0")) // true  (>=1.2.3,<2.0.0)
cs.Match(versions.NewVersion("2.0.0")) // false

// ~ 锁定到 minor
cs2, _ := versions.ParseConstraintSet("~1.2.3")
cs2.Match(versions.NewVersion("1.2.9")) // true  (>=1.2.3,<1.3.0)
cs2.Match(versions.NewVersion("1.3.0")) // false
```

## 🎚 通配符

```go
cs, _ := versions.ParseConstraintSet("1.2.x")
cs.Match(versions.NewVersion("1.2.0"))  // true
cs.Match(versions.NewVersion("1.2.99")) // true
cs.Match(versions.NewVersion("1.3.0"))  // false
```

## 🔗 AND / OR

```go
// AND（逗号）
set, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0,!=1.5.0")
set.Match(versions.NewVersion("1.5.0")) // false（被排除）
set.Match(versions.NewVersion("1.6.0")) // true

// OR（||）
union, _ := versions.ParseConstraintUnion("1.x || 3.0.0")
union.Match(versions.NewVersion("1.9.0")) // true
union.Match(versions.NewVersion("2.0.0")) // false
union.Match(versions.NewVersion("3.0.0")) // true
```

## 🎯 批量过滤

```go
vs := versions.NewVersions("0.9.0", "1.0.0", "1.5.0", "2.0.0", "2.1.0")
cs, _ := versions.ParseConstraintSet("^1.0.0")
compatible := versions.FilterByConstraintSet(vs, cs) // [1.0.0 1.5.0]
```

:::mermaid
flowchart LR
  IN["全部版本<br/>0.9.0 / 1.0.0 / 1.5.0 / 2.0.0 / 2.1.0"]
  IN --> F["FilterByConstraintSet<br/>约束: ^1.0.0<br/>= ≥1.0.0,&lt;2.0.0"]
  F --> CHK{"逐个 Match"}
  CHK -->|"0.9.0 → ❌<1.0.0"| OUT1["排除"]
  CHK -->|"1.0.0 → ✅"| KEEP["保留"]
  CHK -->|"1.5.0 → ✅"| KEEP
  CHK -->|"2.0.0 → ❌≥2.0.0"| OUT2["排除"]
  CHK -->|"2.1.0 → ❌"| OUT2
  KEEP --> RESULT["兼容版本<br/>[1.0.0, 1.5.0]"]

  style IN fill:#f8fafc,stroke:#475569
  style F fill:#eff6ff,stroke:#2563eb
  style KEEP fill:#f0fdf4,stroke:#16a34a
  style OUT1 fill:#fef2f2,stroke:#dc2626
  style OUT2 fill:#fef2f2,stroke:#dc2626
  style RESULT fill:#f0fdf4,stroke:#16a34a,stroke-width:3px
:::

## ❌ 取反

```go
c, _ := versions.ParseConstraint(">=2.0.0")
negated := versions.NegateConstraint(c) // <2.0.0
negated.Match(versions.NewVersion("1.0.0")) // true
```

## 🚀 下一步

- [范围查询](/tutorials/range-query)
- 概念：[约束表达式](/concepts/constraints)
- API：[`ParseConstraintSet`](/sdk/api/parse-constraint-set) · [`ConstraintUnion`](/sdk/api/constraint-union) · [`NegateConstraint`](/sdk/api/negate-constraint) · [`FilterByConstraintSet`](/sdk/api/filter-by-constraint-set)
