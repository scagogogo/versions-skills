package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	checkPrerelease  bool
	checkStable      bool
	checkDev         bool
	checkAlpha       bool
	checkBeta        bool
	checkRC          bool
	checkSnapshot    bool
	checkMilestone   bool
	checkNightly     bool
	checkFinal       bool
	checkGA          bool
	checkPre         bool
	checkRelease     bool
	checkSP          bool
	checkPost        bool
	checkZero        bool
	checkNewer       string
	checkOlder       string
	checkEqual       string
	checkBetweenLow  string
	checkBetweenHigh string
)

var checkCmd = &cobra.Command{
	Use:   "check <version-string>",
	Short: "检查版本号属性（返回布尔结果，exit code 0=真/1=假）",
	Long: `检查版本号的各类属性，返回布尔结果。

使用对应的 flag 检查特定属性，命令返回 JSON 结果，同时以 exit code 表示真假：
  exit code 0 = 属性为真
  exit code 1 = 属性为假

支持的属性检查:
  --prerelease    是否为预发布版本
  --stable        是否为稳定版本
  --dev           是否为开发版
  --alpha         是否为 Alpha 版
  --beta          是否为 Beta 版
  --rc            是否为 RC 版
  --snapshot      是否为快照版
  --milestone     是否为里程碑版
  --nightly       是否为夜间构建版
  --final         是否为 Final 版
  --ga            是否为 GA 版
  --pre           是否为 Pre 版
  --release       是否为 Release 版
  --sp            是否为 SP 版
  --post          是否为 Post 版
  --zero          是否为零值版本
  --newer <v>     是否比指定版本新
  --older <v>     是否比指定版本旧
  --equal <v>     是否与指定版本相等
  --between-low/--between-high  是否在指定范围内

示例:
  versions check --beta v1.2.3-beta1
  versions check --stable 1.2.3
  versions check --stable 1.2.3-beta
  versions check --newer 1.0.0 2.0.0
  versions check --between-low 1.0.0 --between-high 3.0.0 2.0.0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("check", nil, err)
			return
		}

		var result bool
		var description string
		var checkType string

		// Is* 检查
		switch {
		case checkPrerelease:
			result = v.IsPrerelease()
			description = fmt.Sprintf("%s 是预发布版本", v.RawString())
			checkType = "prerelease"
		case checkStable:
			result = v.IsStable()
			description = fmt.Sprintf("%s 是稳定版本", v.RawString())
			checkType = "stable"
		case checkDev:
			result = v.IsDev()
			description = fmt.Sprintf("%s 是开发版", v.RawString())
			checkType = "dev"
		case checkAlpha:
			result = v.IsAlpha()
			description = fmt.Sprintf("%s 是 Alpha 版", v.RawString())
			checkType = "alpha"
		case checkBeta:
			result = v.IsBeta()
			description = fmt.Sprintf("%s 是 Beta 版", v.RawString())
			checkType = "beta"
		case checkRC:
			result = v.IsRC()
			description = fmt.Sprintf("%s 是 RC 版", v.RawString())
			checkType = "rc"
		case checkSnapshot:
			result = v.IsSnapshot()
			description = fmt.Sprintf("%s 是快照版", v.RawString())
			checkType = "snapshot"
		case checkMilestone:
			result = v.IsMilestone()
			description = fmt.Sprintf("%s 是里程碑版", v.RawString())
			checkType = "milestone"
		case checkNightly:
			result = v.IsNightly()
			description = fmt.Sprintf("%s 是夜间构建版", v.RawString())
			checkType = "nightly"
		case checkFinal:
			result = v.IsFinal()
			description = fmt.Sprintf("%s 是 Final 版", v.RawString())
			checkType = "final"
		case checkGA:
			result = v.IsGA()
			description = fmt.Sprintf("%s 是 GA 版", v.RawString())
			checkType = "ga"
		case checkPre:
			result = v.IsPre()
			description = fmt.Sprintf("%s 是 Pre 版", v.RawString())
			checkType = "pre"
		case checkRelease:
			result = v.IsRelease()
			description = fmt.Sprintf("%s 是 Release 版", v.RawString())
			checkType = "release"
		case checkSP:
			result = v.IsSP()
			description = fmt.Sprintf("%s 是 SP 版", v.RawString())
			checkType = "sp"
		case checkPost:
			result = v.IsPost()
			description = fmt.Sprintf("%s 是 Post 版", v.RawString())
			checkType = "post"
		case checkZero:
			result = v.IsZero()
			description = fmt.Sprintf("%s 是零值版本", v.RawString())
			checkType = "zero"

		// 比较检查
		case checkNewer != "":
			target, targetErr := ParseValidVersion(checkNewer)
			if targetErr != nil {
				PrintResult("check", nil, fmt.Errorf("无效的目标版本号: %s", checkNewer))
				return
			}
			result = v.IsNewerThan(target)
			description = fmt.Sprintf("%s 是否比 %s 新", v.RawString(), target.RawString())
			checkType = "newer"
		case checkOlder != "":
			target, targetErr := ParseValidVersion(checkOlder)
			if targetErr != nil {
				PrintResult("check", nil, fmt.Errorf("无效的目标版本号: %s", checkOlder))
				return
			}
			result = v.IsOlderThan(target)
			description = fmt.Sprintf("%s 是否比 %s 旧", v.RawString(), target.RawString())
			checkType = "older"
		case checkEqual != "":
			target, targetErr := ParseValidVersion(checkEqual)
			if targetErr != nil {
				PrintResult("check", nil, fmt.Errorf("无效的目标版本号: %s", checkEqual))
				return
			}
			result = v.Equals(target)
			description = fmt.Sprintf("%s 是否等于 %s", v.RawString(), target.RawString())
			checkType = "equal"

		// 范围检查
		case checkBetweenLow != "" && checkBetweenHigh != "":
			low, lowErr := ParseValidVersion(checkBetweenLow)
			if lowErr != nil {
				PrintResult("check", nil, fmt.Errorf("无效的最低版本号: %s", checkBetweenLow))
				return
			}
			high, highErr := ParseValidVersion(checkBetweenHigh)
			if highErr != nil {
				PrintResult("check", nil, fmt.Errorf("无效的最高版本号: %s", checkBetweenHigh))
				return
			}
			result = v.IsBetween(low, high)
			description = fmt.Sprintf("%s 是否在 %s 和 %s 之间", v.RawString(), low.RawString(), high.RawString())
			checkType = "between"

		default:
			PrintResult("check", nil, fmt.Errorf("请指定检查类型，使用 --prerelease/--stable/--dev/--alpha/--beta/--rc/--snapshot/--milestone/--nightly/--final/--ga/--pre/--release/--sp/--post/--zero/--newer/--older/--equal/--between-low/--between-high"))
			return
		}

		data := map[string]interface{}{
			"version":     v.RawString(),
			"check_type":  checkType,
			"result":      result,
			"description": description,
		}

		PrintResult("check", data, nil)

		if !result {
			os.Exit(1)
		}
	},
}

func init() {
	checkCmd.Flags().BoolVar(&checkPrerelease, "prerelease", false, "是否为预发布版本")
	checkCmd.Flags().BoolVar(&checkStable, "stable", false, "是否为稳定版本")
	checkCmd.Flags().BoolVar(&checkDev, "dev", false, "是否为开发版")
	checkCmd.Flags().BoolVar(&checkAlpha, "alpha", false, "是否为 Alpha 版")
	checkCmd.Flags().BoolVar(&checkBeta, "beta", false, "是否为 Beta 版")
	checkCmd.Flags().BoolVar(&checkRC, "rc", false, "是否为 RC 版")
	checkCmd.Flags().BoolVar(&checkSnapshot, "snapshot", false, "是否为快照版")
	checkCmd.Flags().BoolVar(&checkMilestone, "milestone", false, "是否为里程碑版")
	checkCmd.Flags().BoolVar(&checkNightly, "nightly", false, "是否为夜间构建版")
	checkCmd.Flags().BoolVar(&checkFinal, "final", false, "是否为 Final 版")
	checkCmd.Flags().BoolVar(&checkGA, "ga", false, "是否为 GA 版")
	checkCmd.Flags().BoolVar(&checkPre, "pre", false, "是否为 Pre 版")
	checkCmd.Flags().BoolVar(&checkRelease, "release", false, "是否为 Release 版")
	checkCmd.Flags().BoolVar(&checkSP, "sp", false, "是否为 SP 版")
	checkCmd.Flags().BoolVar(&checkPost, "post", false, "是否为 Post 版")
	checkCmd.Flags().BoolVar(&checkZero, "zero", false, "是否为零值版本")
	checkCmd.Flags().StringVar(&checkNewer, "newer", "", "是否比指定版本新（提供目标版本号）")
	checkCmd.Flags().StringVar(&checkOlder, "older", "", "是否比指定版本旧（提供目标版本号）")
	checkCmd.Flags().StringVar(&checkEqual, "equal", "", "是否与指定版本相等（提供目标版本号）")
	checkCmd.Flags().StringVar(&checkBetweenLow, "between-low", "", "范围检查的最低版本")
	checkCmd.Flags().StringVar(&checkBetweenHigh, "between-high", "", "范围检查的最高版本")
	rootCmd.AddCommand(checkCmd)
}
