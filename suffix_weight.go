package versions

import (
	"regexp"
	"strings"
)

// SuffixWeight 表示版本后缀的语义权重
//
// SuffixWeight 用于在版本比较时为不同类型的后缀分配语义化的权重值，
// 使得预发布版本的排序符合实际发布顺序，而非简单的字典序。
//
// 权重规则（从低到高）:
//   - dev/snapshot < alpha/a < beta/b < milestone/m < rc/cr/pre < 正式版(无后缀)
//
// 使用示例:
//
//	weight := GetSuffixWeight("-alpha1")
//	fmt.Println(weight) // 100
type SuffixWeight int

const (
	// SuffixWeightUnknown 未知后缀类型
	SuffixWeightUnknown SuffixWeight = 0

	// SuffixWeightDev 开发版后缀
	SuffixWeightDev SuffixWeight = 50

	// SuffixWeightSnapshot 快照版后缀
	SuffixWeightSnapshot SuffixWeight = 60

	// SuffixWeightNightly 夜间构建版后缀
	SuffixWeightNightly SuffixWeight = 70

	// SuffixWeightAlpha Alpha版后缀
	SuffixWeightAlpha SuffixWeight = 100

	// SuffixWeightBeta Beta版后缀
	SuffixWeightBeta SuffixWeight = 200

	// SuffixWeightMilestone 里程碑版后缀
	SuffixWeightMilestone SuffixWeight = 300

	// SuffixWeightRC 候选发布版后缀
	SuffixWeightRC SuffixWeight = 400

	// SuffixWeightPre 预发布版后缀
	SuffixWeightPre SuffixWeight = 410

	// SuffixWeightCR 候选发布版(CR变体)
	SuffixWeightCR SuffixWeight = 420

	// SuffixWeightFinal Final后缀(Maven生态)
	SuffixWeightFinal SuffixWeight = 500

	// SuffixWeightRelease Release后缀
	SuffixWeightRelease SuffixWeight = 500

	// SuffixWeightGA GA后缀(Generally Available)
	SuffixWeightGA SuffixWeight = 500

	// SuffixWeightSP 服务包后缀
	SuffixWeightSP SuffixWeight = 600

	// SuffixWeightPatch 补丁版后缀
	SuffixWeightPatch SuffixWeight = 700

	// SuffixWeightPost Post发布版后缀(PEP 440)
	SuffixWeightPost SuffixWeight = 800
)

// suffixWeightPatterns 后缀权重匹配规则，按优先级排列
var suffixWeightPatterns = []struct {
	Pattern *regexp.Regexp
	Weight  SuffixWeight
}{
	{regexp.MustCompile(`(?i)^[-.]?dev\.?\d*$`), SuffixWeightDev},
	{regexp.MustCompile(`(?i)^[-.]?snapshot.*$`), SuffixWeightSnapshot},
	{regexp.MustCompile(`(?i)^[-.]?nightly.*$`), SuffixWeightNightly},
	{regexp.MustCompile(`(?i)^[-.]?a(?:lpha)?\.?\d*$`), SuffixWeightAlpha},
	{regexp.MustCompile(`(?i)^[-.]?b(?:eta)?\.?\d*$`), SuffixWeightBeta},
	{regexp.MustCompile(`(?i)^[-.]?m(?:ilestone)?\.?\d*$`), SuffixWeightMilestone},
	{regexp.MustCompile(`(?i)^[-.]?rc\.?\d*$`), SuffixWeightRC},
	{regexp.MustCompile(`(?i)^[-.]?cr\.?\d*$`), SuffixWeightCR},
	{regexp.MustCompile(`(?i)^[-.]?pre\.?\d*$`), SuffixWeightPre},
	{regexp.MustCompile(`(?i)^[-.]?final$`), SuffixWeightFinal},
	{regexp.MustCompile(`(?i)^[-.]?release\.?\d*$`), SuffixWeightRelease},
	{regexp.MustCompile(`(?i)^[-.]?ga$`), SuffixWeightGA},
	{regexp.MustCompile(`(?i)^[-.]?sp\.?\d*$`), SuffixWeightSP},
	{regexp.MustCompile(`(?i)^[-.]?patch\.?\d*$`), SuffixWeightPatch},
	{regexp.MustCompile(`(?i)^[-.]?post\.?\d*$`), SuffixWeightPost},
}

// GetSuffixWeight 获取后缀的语义权重
//
// 根据后缀字符串匹配预定义的权重规则，返回对应的语义权重。
// 如果没有匹配到任何规则，返回 SuffixWeightUnknown。
//
// 参数:
//   - suffix: 版本后缀字符串，如 "-alpha1"、"-beta2"、"-rc1"
//
// 返回:
//   - SuffixWeight: 后缀的语义权重值
func GetSuffixWeight(suffix string) SuffixWeight {
	normalized := strings.ToLower(strings.TrimSpace(suffix))
	for _, rule := range suffixWeightPatterns {
		if rule.Pattern.MatchString(normalized) {
			return rule.Weight
		}
	}
	return SuffixWeightUnknown
}

// String 返回后缀权重的可读名称
//
// 实现 fmt.Stringer 接口。
//
// 返回:
//   - string: 权重的可读名称，如 "dev"、"alpha"、"beta"、"rc"、"unknown" 等
func (w SuffixWeight) String() string {
	switch w {
	case SuffixWeightDev:
		return "dev"
	case SuffixWeightSnapshot:
		return "snapshot"
	case SuffixWeightNightly:
		return "nightly"
	case SuffixWeightAlpha:
		return "alpha"
	case SuffixWeightBeta:
		return "beta"
	case SuffixWeightMilestone:
		return "milestone"
	case SuffixWeightRC:
		return "rc"
	case SuffixWeightPre:
		return "pre"
	case SuffixWeightCR:
		return "cr"
	case SuffixWeightFinal:
		// SuffixWeightRelease and SuffixWeightGA share the same value (500),
		// so they fall through to this case automatically.
		return "release"
	case SuffixWeightSP:
		return "sp"
	case SuffixWeightPatch:
		return "patch"
	case SuffixWeightPost:
		return "post"
	default:
		return "unknown"
	}
}

// extractSubVersion 从后缀中提取子版本号
//
// 例如 "-alpha1" 提取出 1，"-rc2" 提取出 2。
// 如果无法提取，返回 0。
func extractSubVersion(suffix string) int {
	re := regexp.MustCompile(`(\d+)$`)
	matches := re.FindStringSubmatch(suffix)
	if len(matches) > 1 {
		n := 0
		for _, c := range matches[1] {
			n = n*10 + int(c-'0')
		}
		return n
	}
	return 0
}
