package versions

import compare_anything "github.com/golang-infrastructure/go-compare-anything"

// VersionSuffix 表示版本号的后缀，版本号中数字后面的部分
//
// VersionSuffix 是一个字符串类型，用于表示和操作版本号的后缀部分。
// 在版本号格式中，后缀是位于数字部分之后的字符串，如 "1.2.3-beta1" 中的 "-beta1"。
// 后缀通常用于表示预发布版本、构建元数据或其他版本标识信息。
//
// 后缀可以为空，表示版本号仅包含数字部分，如 "1.2.3"。
//
// 使用示例:
//
//	// 检查后缀是否为空
//	suffix := versions.VersionSuffix("-beta1")
//	if !suffix.IsEmpty() {
//	    fmt.Printf("版本后缀: %s\n", suffix)
//	}
//
//	// 比较后缀的优先级
//	suffix1 := versions.VersionSuffix("-alpha")
//	suffix2 := versions.VersionSuffix("-beta")
//	if suffix1.CompareTo(suffix2) < 0 {
//	    fmt.Println("alpha 后缀的优先级低于 beta 后缀")
//	}
type VersionSuffix string

// EmptyVersionSuffix 空的后缀
//
// 该常量用于表示版本号没有后缀的情况，如纯数字版本号 "1.2.3"。
// 它也用于检查版本后缀是否为空。
const EmptyVersionSuffix VersionSuffix = ""

var _ compare_anything.Comparable[VersionSuffix] = VersionSuffix("")

// IsEmpty 判断版本后缀是否为空
//
// 该方法检查版本后缀是否为空，即是否等于 EmptyVersionSuffix。
//
// 返回:
//   - bool: 如果后缀为空则返回 true，否则返回 false
//
// 使用示例:
//
//	version := versions.NewVersion("1.2.3") // 没有后缀
//	if version.Suffix.IsEmpty() {
//	    fmt.Println("版本没有后缀")
//	}
//
//	version2 := versions.NewVersion("1.2.3-rc1") // 有后缀
//	if !version2.Suffix.IsEmpty() {
//	    fmt.Printf("版本后缀是: %s\n", version2.Suffix)
//	}
func (x VersionSuffix) IsEmpty() bool {
	return x == EmptyVersionSuffix
}

// CompareTo 比较两个版本后缀的优先级
//
// 该方法使用后缀语义权重系统比较两个版本后缀的优先级。
// 对于已知后缀类型（alpha、beta、rc等），按语义权重排序；
// 对于未知后缀，退化为字典序比较。
//
// 参数:
//   - target: 要比较的目标版本后缀
//
// 返回:
//   - int: 如果当前后缀优先级低于目标后缀，返回-1；如果相等，返回0；如果高于，返回1
//
// 使用示例:
//
//	suffix1 := versions.VersionSuffix("-alpha1")
//	suffix2 := versions.VersionSuffix("-beta1")
//
//	result := suffix1.CompareTo(suffix2)
//	if result < 0 {
//	    fmt.Println("alpha1 后缀的优先级低于 beta1 后缀")
//	}
func (x VersionSuffix) CompareTo(target VersionSuffix) int {
	weightX := GetSuffixWeight(string(x))
	weightTarget := GetSuffixWeight(string(target))

	// 如果两个后缀都能匹配到权重规则，则按权重比较
	if weightX != SuffixWeightUnknown && weightTarget != SuffixWeightUnknown {
		if weightX < weightTarget {
			return -1
		} else if weightX > weightTarget {
			return 1
		}
		// 权重相同，比较子版本号（如 alpha1 vs alpha2）
		subX := extractSubVersion(string(x))
		subTarget := extractSubVersion(string(target))
		if subX < subTarget {
			return -1
		} else if subX > subTarget {
			return 1
		}
		return 0
	}

	// 如果其中一个匹配权重规则，另一个不匹配
	if weightX != SuffixWeightUnknown && weightTarget == SuffixWeightUnknown {
		return -1 // 未知后缀排在已知后缀之后
	}
	if weightX == SuffixWeightUnknown && weightTarget != SuffixWeightUnknown {
		return 1
	}

	// 都无法匹配权重规则，退化为字典序比较
	if x > target {
		return 1
	} else if x < target {
		return -1
	}
	return 0
}
