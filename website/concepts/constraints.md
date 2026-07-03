# 约束表达式

::: tip 关键
约束表达式用于判断「某版本是否满足条件」，支持 9 种运算符与 AND/OR 组合。
:::

## 🔢 运算符

| 运算符 | 常量 | 语义 | 示例 |
|:--|:--|:--|:--|
| `=` | `ConstraintEqual` | 等于 | `=1.2.3` |
| `!=` | `ConstraintNotEqual` | 不等于 | `!=1.0.0` |
| `>` | `ConstraintGreaterThan` | 大于 | `>1.0.0` |
| `>=` | `ConstraintGreaterThanOrEqual` | 大于等于 | `>=1.0.0` |
| `<` | `ConstraintLessThan` | 小于 | `<2.0.0` |
| `<=` | `ConstraintLessThanOrEqual` | 小于等于 | `<=2.0.0` |
| `^` | `ConstraintCaret` | 兼容（左起首个非零段锁定） | `^1.2.3` ⇒ `>=1.2.3,<2.0.0` |
| `~` | `ConstraintTilde` | 近似（锁定到 minor） | `~1.2.3` ⇒ `>=1.2.3,<1.3.0` |
| `x`/`X`/`*` | `ConstraintWildcard` | 通配 | `1.2.x` ⇒ `>=1.2.0,<1.3.0` |

## 🔗 组合

- **AND（ConstraintSet）**：逗号分隔，所有条件都须满足。`>=1.0.0,<2.0.0`
- **OR（ConstraintUnion）**：`||` 分隔，任一满足即可。`1.x || 2.0.0`

## 🧪 示例

```go
// AND
cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
cs.Match(versions.NewVersion("1.5.0")) // true
cs.Match(versions.NewVersion("2.0.0")) // false

// OR
cu, _ := versions.ParseConstraintUnion("1.x || 3.0.0")
cu.Match(versions.NewVersion("1.9.0")) // true
cu.Match(versions.NewVersion("2.0.0")) // false

// Version 便捷方法
v := versions.NewVersion("1.5.0")
ok, _ := v.Matches(">=1.0.0,<2.0.0") // true
```

## 📚 延伸

- API：[`Constraint`](/sdk/api/constraint) · [`ConstraintSet`](/sdk/api/constraint-set) · [`ConstraintUnion`](/sdk/api/constraint-union) · [`NegateConstraint`](/sdk/api/negate-constraint)
- 概念：[范围与包含策略](/concepts/range-and-policy)
- CLI：[`constraint`](/cli/commands/constraint) · [`satisfies`](/cli/commands/satisfies)
