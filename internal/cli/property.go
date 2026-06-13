package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

// segmentsCmd 获取版本号的数字段
var segmentsCmd = &cobra.Command{
	Use:   "segments <version-string>",
	Short: "获取版本号的数字段列表",
	Long: `获取版本号的数字段（Segments），返回 int 数组。

示例:
  versions segments 1.2.3       # [1, 2, 3]
  versions segments v1.2.3.4   # [1, 2, 3, 4]`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("segments", nil, err)
			return
		}
		data := map[string]interface{}{
			"raw":      v.RawString(),
			"segments": v.Segments(),
		}
		PrintResult("segments", data, nil)
	},
}

// subVersionCmd 获取后缀中的子版本号
var subVersionCmd = &cobra.Command{
	Use:   "sub-version <version-string>",
	Short: "获取后缀中的子版本号",
	Long: `获取版本号后缀中的子版本号数字部分。

例如 1.2.3-beta2 中的子版本号为 2。

示例:
  versions sub-version 1.2.3-beta2  # 2
  versions sub-version 1.2.3        # 0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("sub-version", nil, err)
			return
		}
		data := map[string]interface{}{
			"raw":         v.RawString(),
			"sub_version": v.SubVersion(),
		}
		PrintResult("sub-version", data, nil)
	},
}

// suffixWeightCmd 获取后缀的语义权重
var suffixWeightCmd = &cobra.Command{
	Use:   "suffix-weight <version-string>",
	Short: "获取版本号后缀的语义权重",
	Long: `获取版本号后缀的语义权重（SuffixWeight）。

权重排序: dev(50) < snapshot(60) < nightly(70) < alpha(100) < beta(200) < milestone(300) < rc(400) < final/release/ga(500) < sp(600) < patch(700) < post(800)

示例:
  versions suffix-weight 1.2.3-beta1  # beta (200)
  versions suffix-weight 1.2.3        # unknown (0)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("suffix-weight", nil, err)
			return
		}
		w := v.SuffixWeight()
		data := map[string]interface{}{
			"raw":           v.RawString(),
			"suffix":        v.Suffix.String(),
			"suffix_weight": w.String(),
			"weight_value":  int(w),
		}
		PrintResult("suffix-weight", data, nil)
	},
}

// purePrefixCmd 获取纯净前缀（去除尾部分隔符）
var purePrefixCmd = &cobra.Command{
	Use:   "pure-prefix <version-string>",
	Short: "获取版本号的纯净前缀（去除尾部分隔符）",
	Long: `获取版本号前缀的纯净形式，即去除尾部的 - 或 . 等分隔符。

例如 v1.2.3 的 Prefix 为 "v"，PurePrefix 也为 "v"。
例如 curl-7.85.0 的 Prefix 为 "curl-"，PurePrefix 为 "curl"。

示例:
  versions pure-prefix v1.2.3       # v
  versions pure-prefix curl-7.85.0  # curl`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("pure-prefix", nil, err)
			return
		}
		data := map[string]interface{}{
			"raw":         v.RawString(),
			"prefix":      v.Prefix.String(),
			"pure_prefix": v.Prefix.PurePrefix(),
		}
		PrintResult("pure-prefix", data, nil)
	},
}

// groupIDCmd 获取版本号的分组 ID
var groupIDCmd = &cobra.Command{
	Use:   "group-id <version-string>",
	Short: "获取版本号的分组 ID",
	Long: `获取版本号的分组 ID（BuildGroupID），即版本号数字部分以 . 连接。

例如 v1.2.3-beta1 的分组 ID 为 "1.2.3"。

示例:
  versions group-id v1.2.3-beta1  # 1.2.3
  versions group-id 2.0.0         # 2.0.0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("group-id", nil, err)
			return
		}
		data := map[string]interface{}{
			"raw":      v.RawString(),
			"group_id": v.BuildGroupID(),
		}
		PrintResult("group-id", data, nil)
	},
}

// satisfiesCmd 版本为中心的约束检查
var satisfiesCmd = &cobra.Command{
	Use:   "satisfies <version-string> <constraint-expression>",
	Short: "检查版本号是否满足约束表达式（版本为中心）",
	Long: `检查版本号是否满足给定的约束表达式。

与 constraint 命令相反：constraint 是约束为中心，satisfies 是版本为中心。
satisfies 使用 Version.Matches() 方法，自动解析表达式。

示例:
  versions satisfies 1.5.0 ">=1.0.0,<2.0.0"
  versions satisfies 2.0.0 "^1.0.0"`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("satisfies", nil, err)
			return
		}

		satisfied, matchErr := v.Matches(args[1])
		if matchErr != nil {
			PrintResult("satisfies", nil, fmt.Errorf("解析约束表达式失败: %w", matchErr))
			return
		}

		data := map[string]interface{}{
			"version":    v.RawString(),
			"expression": args[1],
			"satisfied":  satisfied,
		}

		PrintResult("satisfies", data, nil)
	},
}

// cloneCmd 克隆版本号
var cloneCmd = &cobra.Command{
	Use:   "clone <version-string>",
	Short: "深拷贝一个版本号",
	Long: `深拷贝一个版本号，返回与原始版本完全独立的新副本。

示例:
  versions clone v1.2.3-beta1`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("clone", nil, err)
			return
		}
		cloned := v.Clone()
		data := verfmt.FormatVersion(cloned)
		data["operation"] = "clone"
		data["original"] = v.RawString()
		PrintResult("clone", data, nil)
	},
}

func init() {
	rootCmd.AddCommand(segmentsCmd, subVersionCmd, suffixWeightCmd, purePrefixCmd, groupIDCmd, satisfiesCmd, cloneCmd)
}
