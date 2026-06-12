package versions

import (
	"fmt"
	"strconv"
	"strings"
)

// ConstraintOperator 约束操作符类型
type ConstraintOperator string

const (
	// ConstraintEqual 等于 (=)
	ConstraintEqual ConstraintOperator = "="

	// ConstraintNotEqual 不等于 (!=)
	ConstraintNotEqual ConstraintOperator = "!="

	// ConstraintGreaterThan 大于 (>)
	ConstraintGreaterThan ConstraintOperator = ">"

	// ConstraintGreaterThanOrEqual 大于等于 (>=)
	ConstraintGreaterThanOrEqual ConstraintOperator = ">="

	// ConstraintLessThan 小于 (<)
	ConstraintLessThan ConstraintOperator = "<"

	// ConstraintLessThanOrEqual 小于等于 (<=)
	ConstraintLessThanOrEqual ConstraintOperator = "<="

	// ConstraintCaret 兼容版本 (^) — 兼容左起第一个非零版本号
	ConstraintCaret ConstraintOperator = "^"

	// ConstraintTilde 近似版本 (~) — 兼容到次版本号
	ConstraintTilde ConstraintOperator = "~"

	// ConstraintWildcard 通配符 (x/X/*) — 匹配任意子版本
	ConstraintWildcard ConstraintOperator = "x"
)

// Constraint 表示一个版本约束条件
//
// Constraint 用于判断某个版本是否满足指定的约束条件，
// 如 ">=1.0.0", "^1.2.3", "~1.2.3", "1.x" 等。
//
// 使用示例:
//
//	c, err := versions.ParseConstraint(">=1.0.0")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	v := versions.NewVersion("1.5.0")
//	if c.Match(v) {
//	    fmt.Println("1.5.0 satisfies >=1.0.0")
//	}
type Constraint struct {
	// Operator 约束操作符
	Operator ConstraintOperator

	// Version 约束目标版本
	Version *Version
}

// ConstraintSet 表示一组 AND 组合的约束条件
//
// 多个约束条件之间是 AND 关系，所有条件都必须满足。
// 例如 ">=1.0.0,<2.0.0" 表示版本必须同时满足 >=1.0.0 和 <2.0.0。
type ConstraintSet struct {
	Constraints []Constraint
}

// ParseConstraint 解析单个版本约束表达式
//
// 支持的操作符: =, !=, >, >=, <, <=, ^, ~
// 支持的通配符: x, X, * (如 1.x, 1.2.*)
//
// 参数:
//   - expr: 约束表达式，如 ">=1.0.0", "^1.2.3", "~1.2"
//
// 返回:
//   - *Constraint: 解析后的约束对象
//   - error: 如果表达式格式错误则返回错误
func ParseConstraint(expr string) (*Constraint, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, ErrEmptyConstraint
	}

	// 按操作符长度降序匹配，避免 >= 被 > 先匹配
	operators := []struct {
		op  ConstraintOperator
		len int
	}{
		{ConstraintGreaterThanOrEqual, 2},
		{ConstraintLessThanOrEqual, 2},
		{ConstraintNotEqual, 2},
		{ConstraintCaret, 1},
		{ConstraintTilde, 1},
		{ConstraintGreaterThan, 1},
		{ConstraintLessThan, 1},
		{ConstraintEqual, 1},
	}

	for _, o := range operators {
		if strings.HasPrefix(expr, string(o.op)) {
			versionStr := strings.TrimSpace(expr[o.len:])
			if versionStr == "" {
				return nil, ErrMissingVersionInConstraint
			}
			v := NewVersion(versionStr)
			if !v.IsValid() {
				return nil, ErrInvalidVersionInConstraint
			}
			return &Constraint{Operator: o.op, Version: v}, nil
		}
	}

	// 检查通配符
	if isWildcardVersion(expr) {
		return &Constraint{Operator: ConstraintWildcard, Version: NewVersion(replaceWildcardWithZero(expr))}, nil
	}

	// 无操作符前缀，视为等于
	v := NewVersion(expr)
	if !v.IsValid() {
		return nil, ErrInvalidVersionInConstraint
	}
	return &Constraint{Operator: ConstraintEqual, Version: v}, nil
}

// ParseConstraintSet 解析逗号分隔的 AND 组合约束
//
// 支持格式: ">=1.0.0,<2.0.0", "^1.2.3", "~1.2"
//
// 参数:
//   - expr: 逗号分隔的约束表达式
//
// 返回:
//   - *ConstraintSet: 解析后的约束集合
//   - error: 如果任何子表达式格式错误则返回错误
func ParseConstraintSet(expr string) (*ConstraintSet, error) {
	parts := strings.Split(expr, ",")
	cs := &ConstraintSet{Constraints: make([]Constraint, 0, len(parts))}
	for _, part := range parts {
		c, err := ParseConstraint(part)
		if err != nil {
			return nil, fmt.Errorf("parse constraint %q: %w", part, err)
		}
		cs.Constraints = append(cs.Constraints, *c)
	}
	return cs, nil
}

// Match 判断版本是否满足约束条件
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本满足约束则返回 true
func (c *Constraint) Match(v *Version) bool {
	switch c.Operator {
	case ConstraintEqual:
		return v.CompareTo(c.Version) == 0
	case ConstraintNotEqual:
		return v.CompareTo(c.Version) != 0
	case ConstraintGreaterThan:
		return v.CompareTo(c.Version) > 0
	case ConstraintGreaterThanOrEqual:
		return v.CompareTo(c.Version) >= 0
	case ConstraintLessThan:
		return v.CompareTo(c.Version) < 0
	case ConstraintLessThanOrEqual:
		return v.CompareTo(c.Version) <= 0
	case ConstraintCaret:
		return matchCaret(c.Version, v)
	case ConstraintTilde:
		return matchTilde(c.Version, v)
	case ConstraintWildcard:
		return matchWildcard(c.Version, v)
	default:
		return false
	}
}

// String 返回约束条件的字符串表示
//
// 将约束条件序列化为可解析的字符串格式，如 ">=1.0.0"、"^1.2.3"、"~1.2"。
// 对于通配符约束（1.x），返回原始版本字符串形式。
//
// 返回:
//   - string: 约束条件的字符串表示
//
// 使用示例:
//
//	c, _ := versions.ParseConstraint(">=1.0.0")
//	fmt.Println(c.String()) // 输出: ">=1.0.0"
func (c *Constraint) String() string {
	switch c.Operator {
	case ConstraintWildcard:
		// 通配符需要特殊处理：将版本号数字中对应位置的 0 还原为 x
		parts := make([]string, len(c.Version.VersionNumbers))
		for i, n := range c.Version.VersionNumbers {
			parts[i] = strconv.Itoa(n)
		}
		// 最后一个 0 替换为 x（通配符位置）
		for i := len(parts) - 1; i >= 0; i-- {
			if parts[i] == "0" {
				parts[i] = "x"
				break
			}
		}
		return strings.Join(parts, ".")
	default:
		return string(c.Operator) + c.Version.Raw
	}
}

// Match 判断版本是否满足所有约束（AND 逻辑）
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本满足所有约束则返回 true
func (cs *ConstraintSet) Match(v *Version) bool {
	for _, c := range cs.Constraints {
		if !c.Match(v) {
			return false
		}
	}
	return true
}

// String 返回约束集合的字符串表示
//
// 将约束集合序列化为逗号分隔的字符串格式，如 ">=1.0.0,<2.0.0"。
//
// 返回:
//   - string: 约束集合的字符串表示
//
// 使用示例:
//
//	cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
//	fmt.Println(cs.String()) // 输出: ">=1.0.0,<2.0.0"
func (cs *ConstraintSet) String() string {
	parts := make([]string, len(cs.Constraints))
	for i, c := range cs.Constraints {
		parts[i] = c.String()
	}
	return strings.Join(parts, ",")
}

// Satisfies 判断版本是否满足约束集合
//
// 这是 Match(v) 的语义化别名，使调用方式更自然：
// cs.Satisfies(v) 等价于 cs.Match(v)，与 Version.Satisfies(constraint) 对称。
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本满足所有约束则返回 true
func (cs *ConstraintSet) Satisfies(v *Version) bool {
	return cs.Match(v)
}

// Len 返回约束集合中约束条件的数量
func (cs *ConstraintSet) Len() int {
	return len(cs.Constraints)
}

// ConstraintUnion 表示一组 OR 组合的约束集合
//
// 每个 ConstraintSet 内部是 AND 逻辑，多个 ConstraintSet 之间是 OR 逻辑。
// 例如 ">=1.0.0,<2.0.0 || >=3.0.0" 表示版本必须满足 (>=1.0.0 AND <2.0.0) OR (>=3.0.0)。
type ConstraintUnion struct {
	// Sets AND 约束集合列表，之间是 OR 关系
	Sets []*ConstraintSet
}

// ParseConstraintUnion 解析包含 OR 逻辑的约束表达式
//
// 支持格式: ">=1.0.0,<2.0.0 || >=3.0.0"，其中逗号分隔为 AND，|| 分隔为 OR。
// 也支持不包含 || 的简单表达式，此时等价于 ParseConstraintSet。
//
// 参数:
//   - expr: 约束表达式
//
// 返回:
//   - *ConstraintUnion: 解析后的约束联合
//   - error: 如果表达式格式错误则返回错误
func ParseConstraintUnion(expr string) (*ConstraintUnion, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, ErrEmptyConstraint
	}
	parts := strings.Split(expr, "||")
	union := &ConstraintUnion{
		Sets: make([]*ConstraintSet, 0, len(parts)),
	}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		cs, err := ParseConstraintSet(part)
		if err != nil {
			return nil, err
		}
		union.Sets = append(union.Sets, cs)
	}
	if len(union.Sets) == 0 {
		return nil, ErrEmptyConstraint
	}
	return union, nil
}

// Match 判断版本是否满足约束联合（OR 逻辑）
//
// 只要版本满足任意一个 ConstraintSet 即返回 true。
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本满足任意约束集则返回 true
func (cu *ConstraintUnion) Match(v *Version) bool {
	for _, cs := range cu.Sets {
		if cs.Match(v) {
			return true
		}
	}
	return false
}

// Satisfies 判断版本是否满足约束联合
//
// 这是 Match(v) 的语义化别名，与 Version.Satisfies() 对称。
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本满足任意约束集则返回 true
func (cu *ConstraintUnion) Satisfies(v *Version) bool {
	return cu.Match(v)
}

// String 返回约束联合的字符串表示
//
// 将约束联合序列化为 || 分隔的字符串格式。
//
// 返回:
//   - string: 约束联合的字符串表示
func (cu *ConstraintUnion) String() string {
	parts := make([]string, len(cu.Sets))
	for i, cs := range cu.Sets {
		parts[i] = cs.String()
	}
	return strings.Join(parts, " || ")
}

// matchCaret 实现 ^ 操作符：兼容左起第一个非零版本号
//
// ^1.2.3 := >=1.2.3, <2.0.0
// ^0.2.3 := >=0.2.3, <0.3.0
// ^0.0.3 := >=0.0.3, <0.0.4
func matchCaret(base, v *Version) bool {
	if v.CompareTo(base) < 0 {
		return false
	}
	if len(base.VersionNumbers) == 0 {
		return true
	}
	// 找到第一个非零位
	firstNonZero := -1
	for i, n := range base.VersionNumbers {
		if n != 0 {
			firstNonZero = i
			break
		}
	}
	if firstNonZero == -1 {
		// 全零，如 ^0.0.0，匹配任何版本
		return true
	}
	// 上界：第一个非零位+1，后面全0
	upper := make([]int, len(base.VersionNumbers))
	upper[firstNonZero] = base.VersionNumbers[firstNonZero] + 1
	return v.VersionNumbers.CompareTo(upper) < 0
}

// matchTilde 实现 ~ 操作符：兼容到次版本号
//
// ~1.2.3 := >=1.2.3, <1.3.0
// ~1.2   := >=1.2.0, <1.3.0
func matchTilde(base, v *Version) bool {
	if v.CompareTo(base) < 0 {
		return false
	}
	if len(base.VersionNumbers) < 2 {
		return true
	}
	upper := make([]int, len(base.VersionNumbers))
	copy(upper, base.VersionNumbers)
	upper[0] = base.VersionNumbers[0]
	upper[1] = base.VersionNumbers[1] + 1
	for i := 2; i < len(upper); i++ {
		upper[i] = 0
	}
	return v.VersionNumbers.CompareTo(upper) < 0
}

// matchWildcard 实现 x/X/* 通配符
//
// 1.x := >=1.0.0, <2.0.0
// 1.2.x := >=1.2.0, <1.3.0
func matchWildcard(base, v *Version) bool {
	if v.CompareTo(base) < 0 {
		return false
	}
	// 最后一个有效数字位+1
	lastNonZero := -1
	for i, n := range base.VersionNumbers {
		if n != 0 {
			lastNonZero = i
		}
	}
	if lastNonZero == -1 {
		return true
	}
	upper := make([]int, len(base.VersionNumbers))
	copy(upper, base.VersionNumbers)
	upper[lastNonZero] = base.VersionNumbers[lastNonZero] + 1
	for i := lastNonZero + 1; i < len(upper); i++ {
		upper[i] = 0
	}
	return v.VersionNumbers.CompareTo(upper) < 0
}

// isWildcardVersion 检查版本字符串是否包含通配符
func isWildcardVersion(s string) bool {
	return strings.ContainsAny(s, "xX*")
}

// replaceWildcardWithZero 将通配符替换为0
func replaceWildcardWithZero(s string) string {
	s = strings.ReplaceAll(s, "x", "0")
	s = strings.ReplaceAll(s, "X", "0")
	s = strings.ReplaceAll(s, "*", "0")
	return s
}
