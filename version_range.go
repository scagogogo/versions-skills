package versions

import (
	"fmt"
	"regexp"
)

// VersionRange 表示一个版本范围
//
// VersionRange 是一个包含下界和上界的版本区间，支持开区间和闭区间。
// 它可以用于判断版本是否落在某个范围内，以及过滤版本列表。
//
// 使用示例:
//
//	low := versions.NewVersion("1.0.0")
//	high := versions.NewVersion("2.0.0")
//	r := versions.NewVersionRange(low, high, true, true)
//	v := versions.NewVersion("1.5.0")
//	if r.Contains(v) {
//	    fmt.Println("1.5.0 in [1.0.0, 2.0.0]")
//	}
type VersionRange struct {
	// Low 下界版本
	Low *Version

	// High 上界版本
	High *Version

	// LowInclusive 下界是否为闭区间（包含下界）
	LowInclusive bool

	// HighInclusive 上界是否为闭区间（包含上界）
	HighInclusive bool
}

// NewVersionRange 创建一个新的版本范围
//
// 参数:
//   - low: 下界版本，nil 表示无下界
//   - high: 上界版本，nil 表示无上界
//   - lowInclusive: 下界是否为闭区间
//   - highInclusive: 上界是否为闭区间
//
// 返回:
//   - *VersionRange: 新的版本范围对象
//
// 使用示例:
//
//	r := versions.NewVersionRange(
//	    versions.NewVersion("1.0.0"),
//	    versions.NewVersion("2.0.0"),
//	    true,  // [1.0.0
//	    false, // 2.0.0)
//	)
func NewVersionRange(low, high *Version, lowInclusive, highInclusive bool) *VersionRange {
	return &VersionRange{
		Low:           low,
		High:          high,
		LowInclusive:  lowInclusive,
		HighInclusive: highInclusive,
	}
}

// NewClosedRange 创建一个闭区间版本范围 [low, high]
//
// 使用示例:
//
//	r := versions.NewClosedRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
func NewClosedRange(low, high *Version) *VersionRange {
	return NewVersionRange(low, high, true, true)
}

// NewOpenRange 创建一个开区间版本范围 (low, high)
//
// 使用示例:
//
//	r := versions.NewOpenRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
func NewOpenRange(low, high *Version) *VersionRange {
	return NewVersionRange(low, high, false, false)
}

// Contains 判断版本是否在当前范围内
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本在范围内则返回 true
//
// 使用示例:
//
//	r := versions.NewClosedRange(versions.NewVersion("1.0.0"), versions.NewVersion("2.0.0"))
//	v := versions.NewVersion("1.5.0")
//	fmt.Println(r.Contains(v)) // true
func (r *VersionRange) Contains(v *Version) bool {
	if r.Low != nil {
		cmp := v.CompareTo(r.Low)
		if cmp < 0 {
			return false
		}
		if cmp == 0 && !r.LowInclusive {
			return false
		}
	}
	if r.High != nil {
		cmp := v.CompareTo(r.High)
		if cmp > 0 {
			return false
		}
		if cmp == 0 && !r.HighInclusive {
			return false
		}
	}
	return true
}

// String 返回版本范围的字符串表示
//
// 闭区间使用 []，开区间使用 ()。
//
// 返回:
//   - string: 如 "[1.0.0, 2.0.0]"、"(1.0.0, 2.0.0)"等
func (r *VersionRange) String() string {
	var lowStr, highStr string
	if r.Low == nil {
		lowStr = "*"
	} else {
		lowStr = r.Low.Raw
	}
	if r.High == nil {
		highStr = "*"
	} else {
		highStr = r.High.Raw
	}

	left := "("
	if r.LowInclusive {
		left = "["
	}
	right := ")"
	if r.HighInclusive {
		right = "]"
	}

	return fmt.Sprintf("%s%s, %s%s", left, lowStr, highStr, right)
}

// Filter 过滤版本列表，只保留在范围内的版本
//
// 参数:
//   - versions: 版本对象列表
//
// 返回:
//   - []*Version: 范围内的版本列表
func (r *VersionRange) Filter(versions []*Version) []*Version {
	return Filter(versions, func(v *Version) bool {
		return r.Contains(v)
	})
}

// IsEmpty 判断范围是否为空（下界大于上界）
//
// 返回:
//   - bool: 如果范围为空则返回 true
func (r *VersionRange) IsEmpty() bool {
	if r.Low == nil || r.High == nil {
		return false
	}
	cmp := r.Low.CompareTo(r.High)
	if cmp > 0 {
		return true
	}
	if cmp == 0 && (!r.LowInclusive || !r.HighInclusive) {
		return true
	}
	return false
}

// semverRegex 用于校验是否符合 semver 规范
var semverRegex = regexp.MustCompile(
	`^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`,
)

// IsSemver 判断版本字符串是否符合 SemVer 2.0.0 规范
//
// 严格的 semver 格式要求：主版本号.次版本号.修订版本号，
// 可选的预发布标识（以 - 分隔）和构建元数据（以 + 分隔）。
// 不允许前导零（如 01.02.03）。
//
// 返回:
//   - bool: 如果符合 semver 规范则返回 true
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	fmt.Println(v.IsSemver()) // true
//
//	v2 := versions.NewVersion("1.2.3-alpha.1+build.123")
//	fmt.Println(v2.IsSemver()) // true
//
//	v3 := versions.NewVersion("1.2")
//	fmt.Println(v3.IsSemver()) // false（不够3段）
func (x *Version) IsSemver() bool {
	return semverRegex.MatchString(x.Raw)
}

// ValidateSemver 按照 SemVer 2.0.0 规范严格校验版本号
//
// 与 Validate()（仅做基本校验）不同，ValidateSemver 要求：
// 1. 必须是三段版本号（major.minor.patch）
// 2. 每段数字不允许前导零
// 3. 预发布标识和构建元数据必须符合规范字符集
//
// 返回:
//   - error: 如果不符合 semver 规范则返回错误
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	if err := v.ValidateSemver(); err != nil {
//	    fmt.Println("不符合 semver 规范:", err)
//	}
func (x *Version) ValidateSemver() error {
	if !x.IsValid() {
		return ErrVersionInvalid
	}
	if len(x.VersionNumbers) != 3 {
		return fmt.Errorf("semver requires exactly 3 segments, got %d", len(x.VersionNumbers))
	}
	if !semverRegex.MatchString(x.Raw) {
		return fmt.Errorf("version %q does not conform to semver 2.0.0 specification", x.Raw)
	}
	return nil
}

// PreReleaseType 返回预发布版本的类型标识
//
// 返回后缀的语义权重类型名称字符串，如 "alpha"、"beta"、"rc" 等。
// 如果不是预发布版本则返回空字符串。
//
// 返回:
//   - string: 预发布类型名称
//
// 使用示例:
//
//	v := versions.NewVersion("1.0.0-beta2")
//	fmt.Println(v.PreReleaseType()) // "beta"
//
//	s := versions.NewVersion("1.0.0")
//	fmt.Println(s.PreReleaseType()) // ""
func (x *Version) PreReleaseType() string {
	if x.IsStable() {
		return ""
	}
	w := GetSuffixWeight(string(x.Suffix))
	if w == SuffixWeightUnknown {
		return "unknown"
	}
	return w.String()
}

// Diff 计算两个版本之间的差异
//
// 返回一个 VersionDiff 结构体，包含各版本号段的差异。
// 如果目标版本为 nil，返回 nil。
//
// 参数:
//   - target: 目标版本
//
// 返回:
//   - *VersionDiff: 版本差异对象
//
// 使用示例:
//
//	v1 := versions.NewVersion("1.2.3")
//	v2 := versions.NewVersion("2.0.0")
//	d := v1.Diff(v2)
//	fmt.Printf("major diff: %d\n", d.Major) // 1
func (x *Version) Diff(target *Version) *VersionDiff {
	if target == nil {
		return nil
	}
	return &VersionDiff{
		Major:   target.Major() - x.Major(),
		Minor:   target.Minor() - x.Minor(),
		Patch:   target.Patch() - x.Patch(),
		RawFrom: x.Raw,
		RawTo:   target.Raw,
	}
}

// VersionDiff 表示两个版本之间的差异
//
// VersionDiff 包含各版本号段的差值，以及原始版本字符串。
type VersionDiff struct {
	// Major 主版本号差值（target - source）
	Major int

	// Minor 次版本号差值（target - source）
	Minor int

	// Patch 修订版本号差值（target - source）
	Patch int

	// RawFrom 源版本原始字符串
	RawFrom string

	// RawTo 目标版本原始字符串
	RawTo string
}

// String 返回差异的字符串表示
func (d *VersionDiff) String() string {
	return fmt.Sprintf("%s → %s (major:%+d, minor:%+d, patch:%+d)", d.RawFrom, d.RawTo, d.Major, d.Minor, d.Patch)
}

// IsUpgrade 判断差异是否为升级（主、次或修订版本号至少有一项增加）
func (d *VersionDiff) IsUpgrade() bool {
	return d.Major > 0 || d.Minor > 0 || d.Patch > 0
}

// IsDowngrade 判断差异是否为降级（主、次或修订版本号至少有一项减少）
func (d *VersionDiff) IsDowngrade() bool {
	return d.Major < 0 || d.Minor < 0 || d.Patch < 0
}

// IsMajorChange 判断差异是否涉及主版本号变化
func (d *VersionDiff) IsMajorChange() bool {
	return d.Major != 0
}

// IsMinorChange 判断差异是否仅涉及次版本号变化（主版本号不变）
func (d *VersionDiff) IsMinorChange() bool {
	return d.Major == 0 && d.Minor != 0
}

// IsPatchChange 判断差异是否仅涉及修订版本号变化（主、次版本号不变）
func (d *VersionDiff) IsPatchChange() bool {
	return d.Major == 0 && d.Minor == 0 && d.Patch != 0
}

// Coerce 从任意字符串中提取版本号
//
// Coerce 尝试在输入字符串中查找第一个符合版本号模式的子串。
// 如果找不到则返回一个无效的 Version 对象。
//
// 参数:
//   - s: 可能包含版本号的字符串
//
// 返回:
//   - *Version: 提取到的版本对象
//
// 使用示例:
//
//	v := versions.Coerce("program-1.2.3-linux-amd64")
//	fmt.Println(v.Raw) // "1.2.3"
//
//	v2 := versions.Coerce("download/v2.0.0-beta.tar.gz")
//	fmt.Println(v2.Raw) // "2.0.0-beta"
func Coerce(s string) *Version {
	// 尝试匹配带前缀和后缀的版本号
	re := regexp.MustCompile(`v?\d+(?:\.\d+)+(?:[-._]?(?:alpha|beta|rc|pre|dev|snapshot|nightly|milestone|m|a|b|cr|final|release|ga|sp|patch|post)[-._]?\d*)*`)
	matches := re.FindString(s)
	if matches == "" {
		// 降级：尝试匹配更简单的模式
		re2 := regexp.MustCompile(`\d+(?:\.\d+)+`)
		matches = re2.FindString(s)
	}
	if matches == "" {
		return &Version{}
	}
	return NewVersion(matches)
}

// CoerceE 从任意字符串中提取版本号，找不到则返回错误
//
// 参数:
//   - s: 可能包含版本号的字符串
//
// 返回:
//   - *Version: 提取到的版本对象
//   - error: 如果找不到版本号则返回错误
func CoerceE(s string) (*Version, error) {
	v := Coerce(s)
	if !v.IsValid() {
		return nil, ErrVersionInvalid
	}
	return v, nil
}

// WithMetadata 返回一个修改构建元数据的新版本对象
//
// 原版本对象不变，返回一个新对象，其 Metadata 字段被替换为指定值。
//
// 参数:
//   - metadata: 新的构建元数据字符串，如 "build123"
//
// 返回:
//   - *Version: 修改元数据后的新版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	newV := v.WithMetadata("build.123")
//	fmt.Println(newV.Metadata) // "build.123"
func (x *Version) WithMetadata(metadata string) *Version {
	cloned := x.Clone()
	cloned.Metadata = metadata
	return cloned
}

// Canonical 返回版本的规范字符串表示
//
// 规范格式为：[前缀]主版本号.次版本号.修订版本号[-后缀][+元数据]
// 始终输出三段版本号，不足的补零。
//
// 返回:
//   - string: 规范化的版本字符串
//
// 使用示例:
//
//	v := versions.NewVersion("1.2")
//	fmt.Println(v.Canonical()) // "1.2.0"
//
//	v2 := versions.NewVersion("v1.2.3-beta+build.1")
//	fmt.Println(v2.Canonical()) // "v1.2.3-beta+build.1"
func (x *Version) Canonical() string {
	numbers := x.Segments()
	// 确保至少有3段
	for len(numbers) < 3 {
		numbers = append(numbers, 0)
	}

	result := string(x.Prefix)
	for i, n := range numbers[:3] {
		if i > 0 {
			result += DefaultVersionDelimiter
		}
		result += fmt.Sprintf("%d", n)
	}
	if !x.Suffix.IsEmpty() {
		result += string(x.Suffix)
	}
	if x.Metadata != "" {
		result += "+" + x.Metadata
	}
	return result
}

// Format 按照模板格式化版本号
//
// 支持的占位符:
//   - %M  主版本号
//   - %m  次版本号
//   - %p  修订版本号
//   - %P  前缀
//   - %s  后缀
//   - %r  原始字符串
//   - %c  规范字符串
//   - %%  百分号
//
// 参数:
//   - template: 格式化模板字符串
//
// 返回:
//   - string: 格式化后的字符串
//
// 使用示例:
//
//	v := versions.NewVersion("v1.2.3-beta")
//	fmt.Println(v.Format("version %M.%m.%p"))       // "version 1.2.3"
//	fmt.Println(v.Format("prefix=%P major=%M"))      // "prefix=v major=1"
func (x *Version) Format(template string) string {
	result := make([]byte, 0, len(template)*2)
	for i := 0; i < len(template); i++ {
		if template[i] == '%' && i+1 < len(template) {
			switch template[i+1] {
			case 'M':
				result = append(result, fmt.Sprintf("%d", x.Major())...)
			case 'm':
				result = append(result, fmt.Sprintf("%d", x.Minor())...)
			case 'p':
				result = append(result, fmt.Sprintf("%d", x.Patch())...)
			case 'P':
				result = append(result, x.Prefix.String()...)
			case 's':
				result = append(result, x.Suffix.String()...)
			case 'r':
				result = append(result, x.Raw...)
			case 'c':
				result = append(result, x.Canonical()...)
			case '%':
				result = append(result, '%')
			default:
				result = append(result, template[i:i+2]...)
			}
			i++
		} else {
			result = append(result, template[i])
		}
	}
	return string(result)
}

// Increment 按位置递增版本号数字段
//
// 与 BumpMajor/BumpMinor/BumpPatch 不同，Increment 可以递增任意位置的版本号段，
// 并且将更高位置的段重置为零。
//
// 参数:
//   - segment: 版本号段的位置索引（0=主版本号，1=次版本号，2=修订版本号，...）
//
// 返回:
//   - *Version: 递增后的新版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3.4")
//	newV := v.Increment(2)  // 递增修订版本号
//	fmt.Println(newV.Raw)   // "1.2.4.0"
func (x *Version) Increment(segment int) *Version {
	if segment < 0 || len(x.VersionNumbers) == 0 {
		return x.Clone()
	}
	numbers := x.Segments()
	// 确保数组足够长
	for len(numbers) <= segment {
		numbers = append(numbers, 0)
	}
	numbers[segment]++
	// 更高位置重置为零
	for i := segment + 1; i < len(numbers); i++ {
		numbers[i] = 0
	}
	return NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Build()
}

// ContainsVersion 判断版本列表中是否包含指定版本
//
// 根据 Raw 字段判断版本是否相同。
//
// 参数:
//   - versions: 版本对象列表
//   - target: 目标版本对象
//
// 返回:
//   - bool: 如果列表中包含目标版本则返回 true
//
// 使用示例:
//
//	list := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")
//	v := versions.NewVersion("1.1.0")
//	fmt.Println(versions.ContainsVersion(list, v)) // true
func ContainsVersion(versions []*Version, target *Version) bool {
	for _, v := range versions {
		if v.Raw == target.Raw {
			return true
		}
	}
	return false
}

// IndexOf 查找版本在列表中的位置
//
// 根据 Raw 字段匹配，返回第一次出现的索引。如果未找到则返回 -1。
//
// 参数:
//   - versions: 版本对象列表
//   - target: 目标版本对象
//
// 返回:
//   - int: 版本在列表中的索引，未找到返回 -1
//
// 使用示例:
//
//	list := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")
//	idx := versions.IndexOf(list, versions.NewVersion("1.1.0"))
//	fmt.Println(idx) // 1
func IndexOf(versions []*Version, target *Version) int {
	for i, v := range versions {
		if v.Raw == target.Raw {
			return i
		}
	}
	return -1
}

// GroupByMajor 按主版本号分组
//
// 返回按主版本号分组的映射，键为主版本号，值为对应的版本组。
//
// 参数:
//   - versions: 版本对象列表
//
// 返回:
//   - map[int][]*Version: 按主版本号分组的结果
//
// 使用示例:
//
//	list := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
//	groups := versions.GroupByMajor(list)
//	for major, vs := range groups {
//	    fmt.Printf("Major %d: %d versions\n", major, len(vs))
//	}
func GroupByMajor(versions []*Version) map[int][]*Version {
	result := make(map[int][]*Version)
	for _, v := range versions {
		result[v.Major()] = append(result[v.Major()], v)
	}
	return result
}

// GroupByMinor 按次版本号分组
//
// 返回按"主版本号.次版本号"分组的映射。
//
// 参数:
//   - versions: 版本对象列表
//
// 返回:
//   - map[string][]*Version: 按主.次版本号分组的结果，键如 "1.2"
//
// 使用示例:
//
//	list := versions.NewVersions("1.0.0", "1.0.1", "1.1.0", "2.0.0")
//	groups := versions.GroupByMinor(list)
//	for key, vs := range groups {
//	    fmt.Printf("Minor %s: %d versions\n", key, len(vs))
//	}
func GroupByMinor(versions []*Version) map[string][]*Version {
	result := make(map[string][]*Version)
	for _, v := range versions {
		key := fmt.Sprintf("%d.%d", v.Major(), v.Minor())
		result[key] = append(result[key], v)
	}
	return result
}

// NegateConstraint 返回约束条件的否定形式
//
// 例如 >=1.0.0 的否定为 <1.0.0，=1.0.0 的否定为 !=1.0.0。
//
// 参数:
//   - c: 要否定的约束条件
//
// 返回:
//   - *Constraint: 否定后的约束条件
//
// 使用示例:
//
//	c, _ := versions.ParseConstraint(">=1.0.0")
//	neg := versions.NegateConstraint(c)
//	fmt.Println(neg.String()) // "<1.0.0"
func NegateConstraint(c *Constraint) *Constraint {
	var negOp ConstraintOperator
	switch c.Operator {
	case ConstraintEqual:
		negOp = ConstraintNotEqual
	case ConstraintNotEqual:
		negOp = ConstraintEqual
	case ConstraintGreaterThan:
		negOp = ConstraintLessThanOrEqual
	case ConstraintGreaterThanOrEqual:
		negOp = ConstraintLessThan
	case ConstraintLessThan:
		negOp = ConstraintGreaterThanOrEqual
	case ConstraintLessThanOrEqual:
		negOp = ConstraintGreaterThan
	default:
		// 对于 ^, ~, x 等复杂操作符，无法简单取反，返回原约束的 != 形式
		negOp = ConstraintNotEqual
	}
	return &Constraint{Operator: negOp, Version: c.Version}
}

// Hash 返回版本的哈希键值
//
// 返回原始版本字符串，可用于 map 的键。
//
// 返回:
//   - string: 可用作哈希键的字符串
func (x *Version) Hash() string {
	return x.Raw
}
