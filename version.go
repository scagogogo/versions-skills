package versions

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	compare_anything "github.com/golang-infrastructure/go-compare-anything"
)

var (
	// ErrVersionInvalid 表示版本号格式无效的错误
	//
	// 当尝试解析不符合要求的版本号字符串时返回此错误
	ErrVersionInvalid = errors.New("version invalid")

	// ErrEmptyConstraint 表示约束表达式为空的错误
	ErrEmptyConstraint = errors.New("empty constraint expression")

	// ErrMissingVersionInConstraint 表示约束表达式中缺少版本号的错误
	ErrMissingVersionInConstraint = errors.New("missing version in constraint")

	// ErrInvalidVersionInConstraint 表示约束表达式中版本号无效的错误
	ErrInvalidVersionInConstraint = errors.New("invalid version in constraint")
)

// Version 用于表示一个版本号
//
// Version 结构体封装了版本号的各个组成部分，包括原始字符串、发布时间、数字部分、
// 前缀和后缀。它支持版本号的解析、比较和分组等操作，实现了 Comparable 接口以便
// 进行版本排序。
//
// 一个典型的版本号格式可能为：v1.2.3-beta1，其中：
// - "v" 是前缀
// - "1.2.3" 是数字部分
// - "-beta1" 是后缀
//
// 使用示例:
//
//	// 创建一个版本对象
//	version := versions.NewVersion("v1.2.3-rc1")
//
//	// 检查版本是否有效
//	if version.IsValid() {
//	    fmt.Printf("版本号有效: %s\n", version.Raw)
//	    fmt.Printf("版本号数字部分: %v\n", version.VersionNumbers)
//	}
//
//	// 比较两个版本
//	v1 := versions.NewVersion("1.2.3")
//	v2 := versions.NewVersion("1.3.0")
//	if v1.CompareTo(v2) < 0 {
//	    fmt.Println("v1 比 v2 旧")
//	}
type Version struct {

	// Raw 原始的版本号字符串
	Raw string `json:"raw"`

	// PublicTime 此版本的发布时间
	PublicTime time.Time `json:"public_time"`

	// VersionNumbers 版本号中的数字部分
	// 例如对于版本号 "v1.2.3-beta1"，VersionNumbers 为 [1,2,3]
	VersionNumbers VersionNumbers `json:"version_numbers"`

	// Prefix 版本号数字部分之前的前缀
	// 例如对于版本号 "v1.2.3"，Prefix 为 "v"
	Prefix VersionPrefix `json:"prefix"`

	// Suffix 版本号数字部分之后的后缀
	// 例如对于版本号 "1.2.3-beta1"，Suffix 为 "-beta1"
	Suffix VersionSuffix `json:"suffix"`

	// Metadata semver 构建元数据
	//
	// 在 semver 规范中，构建元数据是版本号中 + 号后面的部分，如 "1.0.0+build123" 中的 "build123"。
	// 根据 semver 规范，构建元数据不参与版本比较。
	Metadata string `json:"metadata,omitempty"`
}

var _ compare_anything.Comparable[*Version] = &Version{}

// NewVersion 从版本字符串创建一个新的 Version 对象
//
// 该方法解析给定的版本字符串，并返回一个填充了相应字段的 Version 对象。
// 即使版本字符串格式不正确，该方法也会返回一个对象，但其 IsValid() 方法可能返回 false。
//
// 参数:
//   - versionStr: 要解析的版本号字符串，如 "1.2.3" 或 "v1.2.3-rc1"
//
// 返回:
//   - *Version: 解析后的 Version 对象
//
// 使用示例:
//
//	version := versions.NewVersion("v1.2.3-beta1")
func NewVersion(versionStr string) *Version {
	return NewVersionStringParser(versionStr).Parse()
}

// NewVersionE 从版本字符串创建一个新的 Version 对象，并返回可能的错误
//
// 与 NewVersion 不同，该方法会在版本字符串格式不正确时返回错误。
//
// 参数:
//   - versionStr: 要解析的版本号字符串
//
// 返回:
//   - *Version: 解析后的 Version 对象，如果解析失败则为 nil
//   - error: 如果版本号无效，则返回 ErrVersionInvalid 错误
//
// 使用示例:
//
//	version, err := versions.NewVersionE("v1.2.3-beta1")
//	if err != nil {
//	    log.Fatalf("无效的版本号: %v", err)
//	}
func NewVersionE(versionStr string) (*Version, error) {
	v := NewVersionStringParser(versionStr).Parse()
	if v.IsValid() {
		return v, nil
	} else {
		return nil, ErrVersionInvalid
	}
}

// MustParse 解析版本号字符串，如果解析失败则 panic
//
// 这是 NewVersionE() 的 panic 变体，适用于初始化时确定版本号合法的场景，
// 类似于 regexp.MustCompile 的模式。在版本号来自硬编码或测试数据的场景下非常有用。
//
// 参数:
//   - versionStr: 版本号字符串
//
// 返回:
//   - *Version: 解析后的版本对象
func MustParse(versionStr string) *Version {
	v, err := NewVersionE(versionStr)
	if err != nil {
		panic(fmt.Sprintf("MustParse(%q): %v", versionStr, err))
	}
	return v
}

// NewVersions 批量创建多个 Version 对象
//
// 该方法接受多个版本字符串，并返回相应的 Version 对象数组。
//
// 参数:
//   - versionStringSlice: 一个或多个版本号字符串
//
// 返回:
//   - []*Version: 解析后的 Version 对象数组
//
// 使用示例:
//
//	versions := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")
//	for _, v := range versions {
//	    fmt.Println(v.Raw)
//	}
func NewVersions(versionStringSlice ...string) []*Version {
	versions := make([]*Version, len(versionStringSlice))
	for i, versionStr := range versionStringSlice {
		versions[i] = NewVersion(versionStr)
	}
	return versions
}

// IsValid 检查版本号是否有效
//
// 判断依据是版本号中是否包含数字部分。只有当解析到了版本号数字时才认为是有效的版本号。
//
// 返回:
//   - bool: 如果版本号有效则返回 true，否则返回 false
//
// 使用示例:
//
//	version := versions.NewVersion("not-a-version")
//	if !version.IsValid() {
//	    fmt.Println("无效的版本号")
//	}
func (v *Version) IsValid() bool {
	// 版本号数组不为空，则表示为有效版本
	return len(v.VersionNumbers) > 0
}

// BuildGroupID 构造版本所属的组的ID
//
// 该方法根据版本号的数字部分生成一个组ID，用于将相似版本分组。
//
// 返回:
//   - string: 表示版本组的ID字符串
//
// 使用示例:
//
//	version := versions.NewVersion("1.2.3")
//	groupID := version.BuildGroupID()
//	fmt.Printf("版本组ID: %s\n", groupID)
func (x *Version) BuildGroupID() string {
	return x.VersionNumbers.BuildGroupID()
}

// CompareTo 比较两个版本号
//
// 该方法按以下顺序比较两个版本号：
// 1. 首先比较主版本号数字部分
// 2. 其次比较发布时间
// 3. 然后比较后缀
// 4. 最后比较原始版本号字符串
//
// 参数:
//   - target: 要比较的目标版本对象
//
// 返回:
//   - int: 如果当前版本小于目标版本，返回-1；如果相等，返回0；如果大于，返回1
//
// 使用示例:
//
//	v1 := versions.NewVersion("1.0.0")
//	v2 := versions.NewVersion("1.1.0")
//
//	switch v1.CompareTo(v2) {
//	case -1:
//	    fmt.Println("v1 < v2")
//	case 0:
//	    fmt.Println("v1 = v2")
//	case 1:
//	    fmt.Println("v1 > v2")
//	}
func (x *Version) CompareTo(target *Version) int {

	// 1. 先按照主版本号排序，仅当两个的主版本号都存在的时候才会进行比较，它们的长度不必相等，但是不能有为空的
	if len(x.VersionNumbers) != 0 && len(target.VersionNumbers) != 0 {
		r := x.VersionNumbers.CompareTo(target.VersionNumbers)
		if r != 0 {
			return r
		}
	}

	// 2. 然后按照后缀排序，修复：空后缀(release)应大于有后缀(pre-release)
	switch {
	case x.Suffix.IsEmpty() && target.Suffix.IsEmpty():
		// 两者都是正式版，后缀相等，继续比较
	case x.Suffix.IsEmpty() && !target.Suffix.IsEmpty():
		// 当前是正式版，目标是预发布版，当前更大
		return 1
	case !x.Suffix.IsEmpty() && target.Suffix.IsEmpty():
		// 当前是预发布版，目标是正式版，当前更小
		return -1
	default:
		// 两者都有后缀，比较后缀
		r := x.Suffix.CompareTo(target.Suffix)
		if r != 0 {
			return r
		}
	}

	// 3. 然后按照发布时间排序
	if !target.PublicTime.IsZero() && !x.PublicTime.IsZero() {
		r2 := x.PublicTime.UnixMilli() - target.PublicTime.UnixMilli()
		if r2 != 0 {
			// 不做类型转换是为了避免特殊情况下因为类型转换而丢失精度结果错误，而采用比较的方式
			if r2 > 0 {
				return 1
			} else {
				return -1
			}
		}
	}

	// 4. 最后实在不行就是比较原始版本号的字典序吧
	if x.Raw == target.Raw {
		return 0
	} else if x.Raw < target.Raw {
		return -1
	} else if x.Raw > target.Raw {
		return 1
	}

	// unreachable
	return 0
}

// String 返回版本的JSON字符串表示
//
// 该方法将Version对象序列化为JSON字符串，便于打印和调试。
//
// 返回:
//   - string: 版本的JSON字符串表示
//
// 使用示例:
//
//	version := versions.NewVersion("1.2.3")
//	fmt.Println(version.String()) // 输出JSON格式的版本信息
func (x *Version) String() string {
	marshal, _ := json.Marshal(struct {
		Raw            string         `json:"raw"`
		PublicTime     time.Time      `json:"public_time"`
		VersionNumbers VersionNumbers `json:"version_numbers"`
		Prefix         VersionPrefix  `json:"prefix"`
		Suffix         VersionSuffix  `json:"suffix"`
	}{
		Raw:            x.Raw,
		PublicTime:     x.PublicTime,
		VersionNumbers: x.VersionNumbers,
		Prefix:         x.Prefix,
		Suffix:         x.Suffix,
	})
	return string(marshal)
}

// IsPrerelease 判断版本是否为预发布版本
//
// 预发布版本是指带有后缀（如 -alpha, -beta, -rc, -SNAPSHOT 等）的版本。
// 正式版本（无后缀）返回 false。
//
// 返回:
//   - bool: 如果是预发布版本则返回 true
//
// 使用示例:
//
//	v := versions.NewVersion("1.0.0-beta")
//	if v.IsPrerelease() {
//	    fmt.Println("这是预发布版本")
//	}
func (x *Version) IsPrerelease() bool {
	return !x.Suffix.IsEmpty()
}

// IsStable 判断版本是否为正式稳定版本
//
// 正式稳定版本是指不带任何后缀的版本，如 "1.0.0"。
//
// 返回:
//   - bool: 如果是正式稳定版本则返回 true
func (x *Version) IsStable() bool {
	return x.Suffix.IsEmpty()
}

// IsDev 判断版本是否为开发版
//
// 开发版是指后缀包含 dev 标识的版本，如 "1.0.0-dev1"。
func (x *Version) IsDev() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightDev
}

// IsAlpha 判断版本是否为 Alpha 版
func (x *Version) IsAlpha() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightAlpha
}

// IsBeta 判断版本是否为 Beta 版
func (x *Version) IsBeta() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightBeta
}

// IsRC 判断版本是否为候选发布版(RC)
func (x *Version) IsRC() bool {
	w := GetSuffixWeight(string(x.Suffix))
	return w == SuffixWeightRC || w == SuffixWeightCR
}

// IsSnapshot 判断版本是否为快照版
func (x *Version) IsSnapshot() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightSnapshot
}

// IsMilestone 判断版本是否为里程碑版
//
// 里程碑版是指后缀包含 milestone/m 标识的版本，如 "1.0.0-m1"、"1.0.0-milestone1"。
func (x *Version) IsMilestone() bool {
	w := GetSuffixWeight(string(x.Suffix))
	return w == SuffixWeightMilestone
}

// IsNightly 判断版本是否为夜间构建版
//
// 夜间构建版是指后缀包含 nightly 标识的版本，如 "1.0.0-nightly"。
func (x *Version) IsNightly() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightNightly
}

// IsFinal 判断版本是否为 Final 版本
//
// Final 版本是指后缀包含 final 标识的版本（Maven 生态常见），如 "1.0.0-final"。
// 注意：Final 后缀与无后缀的正式版语义相同，但 IsFinal 专用于检测显式的 final 后缀。
func (x *Version) IsFinal() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightFinal
}

// IsGA 判断版本是否为 GA（Generally Available）版本
//
// GA 版本是指后缀包含 ga 标识的版本，如 "1.0.0-ga"。
func (x *Version) IsGA() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightGA
}

// IsPre 判断版本是否为预发布版（pre 标识）
//
// 预发布版是指后缀包含 pre 标识的版本，如 "1.0.0-pre1"。
// 注意：这是显式的 pre 后缀，与 IsPrerelease()（判断是否有任何后缀）不同。
func (x *Version) IsPre() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightPre
}

// IsRelease 判断版本是否带有 release 后缀
//
// 带 release 后缀的版本如 "1.0.0-release"。
// 注意：这与 IsStable()（无后缀）不同，IsRelease 检测的是显式的 "-release" 后缀。
func (x *Version) IsRelease() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightRelease
}

// IsSP 判断版本是否为服务包版本
//
// 服务包版本是指后缀包含 sp 标识的版本，如 "1.0.0-sp1"。
func (x *Version) IsSP() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightSP
}

// IsPost 判断版本是否为 Post 发布版
//
// Post 发布版是指后缀包含 post 标识的版本（PEP 440 规范），如 "1.0.0-post1"。
func (x *Version) IsPost() bool {
	return GetSuffixWeight(string(x.Suffix)) == SuffixWeightPost
}

// Satisfies 判断版本是否满足给定的约束条件
//
// 这是 Constraint.Match(v) 的反向调用方式，语义更自然：
// v.Satisfies(constraint) 等价于 constraint.Match(v)。
//
// 参数:
//   - constraint: 版本约束条件
//
// 返回:
//   - bool: 如果版本满足约束则返回 true
//
// 使用示例:
//
//	c, _ := versions.ParseConstraint(">=1.0.0")
//	v := versions.NewVersion("1.5.0")
//	if v.Satisfies(c) {
//	    fmt.Println("版本满足约束")
//	}
func (x *Version) Satisfies(constraint *Constraint) bool {
	return constraint.Match(x)
}

// Matches 判断版本是否满足约束表达式字符串
//
// 该方法是 ParseConstraintSet + Match 的便捷组合，
// 适用于需要从字符串快速判断版本是否满足约束的场景。
//
// 参数:
//   - expr: 约束表达式字符串，如 ">=1.0.0"
//
// 返回:
//   - bool: 如果版本满足约束则返回 true
//   - error: 如果约束表达式解析失败则返回错误
//
// 使用示例:
//
//	v := versions.NewVersion("1.5.0")
//	ok, err := v.Matches(">=1.0.0,<2.0.0")
//	if ok {
//	    fmt.Println("版本在范围内")
//	}
func (x *Version) Matches(expr string) (bool, error) {
	cs, err := ParseConstraintSet(expr)
	if err != nil {
		return false, err
	}
	return cs.Match(x), nil
}

// IsNewerThan 判断当前版本是否比目标版本更新
//
// 等价于 CompareTo(target) > 0，但语义更清晰。
func (x *Version) IsNewerThan(target *Version) bool {
	return x.CompareTo(target) > 0
}

// IsOlderThan 判断当前版本是否比目标版本更旧
//
// 等价于 CompareTo(target) < 0，但语义更清晰。
func (x *Version) IsOlderThan(target *Version) bool {
	return x.CompareTo(target) < 0
}

// Equals 判断当前版本是否与目标版本相等
//
// 等价于 CompareTo(target) == 0，但语义更清晰。
func (x *Version) Equals(target *Version) bool {
	return x.CompareTo(target) == 0
}

// IsBetween 判断当前版本是否在两个版本之间（包含边界）
//
// 如果 low <= x <= high 则返回 true。
// 如果 low 或 high 为 nil，则忽略对应的边界检查。
func (x *Version) IsBetween(low, high *Version) bool {
	if low != nil && x.CompareTo(low) < 0 {
		return false
	}
	if high != nil && x.CompareTo(high) > 0 {
		return false
	}
	return true
}

// Major 返回主版本号
//
// 如果版本号数字部分为空则返回 0。
func (x *Version) Major() int {
	if len(x.VersionNumbers) > 0 {
		return x.VersionNumbers[0]
	}
	return 0
}

// Minor 返回次版本号
//
// 如果版本号数字部分少于2位则返回 0。
func (x *Version) Minor() int {
	if len(x.VersionNumbers) > 1 {
		return x.VersionNumbers[1]
	}
	return 0
}

// Patch 返回修订版本号
//
// 如果版本号数字部分少于3位则返回 0。
func (x *Version) Patch() int {
	if len(x.VersionNumbers) > 2 {
		return x.VersionNumbers[2]
	}
	return 0
}

// RawString 返回版本的原始字符串表示
//
// 与 String()（返回 JSON 格式）不同，RawString() 返回解析前的原始版本字符串，
// 如 "v1.2.3-beta1"。这是获取版本字符串最直接的方式。
//
// 返回:
//   - string: 原始版本字符串
//
// 使用示例:
//
//	v := versions.NewVersion("v1.2.3-beta1")
//	fmt.Println(v.RawString()) // 输出: v1.2.3-beta1
func (x *Version) RawString() string {
	return x.Raw
}

// WithNumbers 返回一个修改版本号数字部分的新版本对象
//
// 原版本对象不变，返回一个新对象，其版本号数字部分被替换为指定值。
// 前缀和后缀保持不变。
//
// 参数:
//   - numbers: 新的版本号数字部分
//
// 返回:
//   - *Version: 修改版本号后的新版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	newV := v.WithNumbers([]int{2, 0, 0})
//	// newV.Raw == "2.0.0"
func (x *Version) WithNumbers(numbers []int) *Version {
	v := NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Suffix(string(x.Suffix)).
		Build()
	v.Metadata = x.Metadata
	return v
}

// SubVersion 返回后缀中的子版本号
//
// 例如 "-beta2" 返回 2，"-rc1" 返回 1。如果后缀中没有数字则返回 0。
// 该方法将内部使用的 extractSubVersion 函数暴露为公开 API。
//
// 返回:
//   - int: 后缀中的子版本号数字
//
// 使用示例:
//
//	v := versions.NewVersion("1.0.0-beta2")
//	fmt.Println(v.SubVersion()) // 输出: 2
func (x *Version) SubVersion() int {
	return extractSubVersion(string(x.Suffix))
}

// SuffixWeight 返回版本后缀的语义权重
//
// 等价于 GetSuffixWeight(string(x.Suffix))，但作为方法调用更方便。
//
// 返回:
//   - SuffixWeight: 后缀的语义权重值
//
// 使用示例:
//
//	v := versions.NewVersion("1.0.0-beta")
//	w := v.SuffixWeight()
//	fmt.Println(w == versions.SuffixWeightBeta) // true
func (x *Version) SuffixWeight() SuffixWeight {
	return GetSuffixWeight(string(x.Suffix))
}

// IsZero 判断版本是否为零值
//
// 零值版本是未初始化的 Version{} 结构体，其所有字段都是默认值。
// 与 IsValid()（检查是否有版本号数字）不同，IsZero 检查是否完全没有被设置。
//
// 返回:
//   - bool: 如果是零值则返回 true
func (x *Version) IsZero() bool {
	return x.Raw == "" && len(x.VersionNumbers) == 0 && x.Prefix.IsEmpty() && x.Suffix.IsEmpty() && x.PublicTime.IsZero()
}

// MarshalText 实现 encoding.TextMarshaler 接口
//
// 将版本序列化为原始版本字符串的字节切片。
// 这使得 Version 可以被 encoding/json、toml、yaml 等序列化格式自动处理。
//
// 返回:
//   - []byte: 原始版本字符串的字节切片
//   - error: 始终为 nil
func (x Version) MarshalText() ([]byte, error) {
	return []byte(x.Raw), nil
}

// UnmarshalText 实现 encoding.TextUnmarshaler 接口
//
// 从字节切片反序列化版本对象。解析给定的版本字符串，
// 如果版本无效则返回 ErrVersionInvalid。
//
// 参数:
//   - text: 版本字符串的字节切片
//
// 返回:
//   - error: 如果版本无效则返回 ErrVersionInvalid
func (x *Version) UnmarshalText(text []byte) error {
	v := NewVersion(string(text))
	if !v.IsValid() {
		return ErrVersionInvalid
	}
	*x = *v
	return nil
}

// Core 返回版本的核心部分（去除后缀）
//
// 返回一个新 Version 对象，只保留前缀和版本号数字部分，去除所有后缀。
// 这是获取版本"纯净"数字部分的快捷方式。
//
// 返回:
//   - *Version: 去除后缀后的核心版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3-beta1")
//	core := v.Core()
//	fmt.Println(core.RawString()) // 输出: "1.2.3"
func (x *Version) Core() *Version {
	return x.WithSuffix("")
}

// Validate 严格校验版本号格式
//
// 与 IsValid()（仅检查是否有版本号数字）不同，Validate 执行更严格的校验：
// 1. 版本号数字部分不能为空
// 2. 每个版本号数字必须 >= 0
//
// 返回:
//   - error: 如果版本号不符合严格格式要求则返回错误
func (x *Version) Validate() error {
	if len(x.VersionNumbers) == 0 {
		return ErrVersionInvalid
	}
	for _, n := range x.VersionNumbers {
		if n < 0 {
			return fmt.Errorf("version number %d is negative", n)
		}
	}
	return nil
}

// Segments 返回版本号数字段的整数数组
//
// 等价于直接访问 VersionNumbers，但返回 []int 类型，
// 便于与不使用 VersionNumbers 类型的代码交互。
//
// 返回:
//   - []int: 版本号数字段数组
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3.4")
//	segments := v.Segments()
//	// segments == []int{1, 2, 3, 4}
func (x *Version) Segments() []int {
	result := make([]int, len(x.VersionNumbers))
	copy(result, x.VersionNumbers)
	return result
}

// Segments64 返回版本号数字段的 int64 数组
//
// 与 Segments() 相同，但返回 int64 类型，
// 适用于需要大数值范围的场景。
//
// 返回:
//   - []int64: 版本号数字段 int64 数组
func (x *Version) Segments64() []int64 {
	result := make([]int64, len(x.VersionNumbers))
	for i, n := range x.VersionNumbers {
		result[i] = int64(n)
	}
	return result
}

// MarshalJSON 实现 json.Marshaler 接口
//
// 将版本序列化为 JSON 字符串（双引号包裹的原始版本字符串），
// 而非默认的结构体 JSON。这使得版本在 JSON 上下文中表现为简单字符串。
//
// 返回:
//   - []byte: JSON 编码的版本字符串
//   - error: 始终为 nil
func (x Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Raw)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
//
// 从 JSON 字符串反序列化版本对象。
//
// 参数:
//   - data: JSON 编码的版本字符串
//
// 返回:
//   - error: 如果版本无效则返回错误
func (x *Version) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v := NewVersion(s)
	if !v.IsValid() {
		return ErrVersionInvalid
	}
	*x = *v
	return nil
}

// Scan 实现 sql.Scanner 接口
//
// 从数据库扫描版本值。支持 string 和 []byte 类型。
//
// 参数:
//   - src: 数据库值
//
// 返回:
//   - error: 如果值类型不支持或版本无效则返回错误
func (x *Version) Scan(src interface{}) error {
	var s string
	switch v := src.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("cannot scan %T into Version", src)
	}
	parsed := NewVersion(s)
	if !parsed.IsValid() {
		return ErrVersionInvalid
	}
	*x = *parsed
	return nil
}

// Value 实现 driver.Valuer 接口
//
// 返回版本字符串用于数据库存储。
//
// 返回:
//   - driver.Value: 版本原始字符串
//   - error: 始终为 nil
func (x Version) Value() (driver.Value, error) {
	return x.Raw, nil
}
