package cli

import (
	"strconv"

	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var (
	countStable     bool
	countPrerelease bool
	countMajor      string
	countMinor      string
	countPatch      string
	countFromFile   string
)

var countCmd = &cobra.Command{
	Use:   "count [version-strings...]",
	Short: "统计满足条件的版本号数量",
	Long: `统计版本号列表中满足指定条件的版本数量。

多个条件之间为 AND 逻辑。

示例:
  versions count --stable 1.0.0 1.0.0-beta 2.0.0
  versions count --prerelease 1.0.0-alpha 1.0.0 2.0.0-beta
  versions count --major 1 1.0.0 2.0.0 1.5.0
  versions count --stable --major 2 1.0.0 2.0.0 2.5.0-beta 3.0.0
  cat versions.txt | versions count --stable`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, countFromFile)
		if err != nil {
			PrintResult("count", nil, err)
			return
		}

		// 解析 --major/--minor/--patch
		var majorVal, minorVal, patchVal int
		var hasMajor, hasMinor, hasPatch bool

		if countMajor != "" {
			majorVal, err = strconv.Atoi(countMajor)
			if err != nil {
				PrintResult("count", nil, err)
				return
			}
			hasMajor = true
		}
		if countMinor != "" {
			minorVal, err = strconv.Atoi(countMinor)
			if err != nil {
				PrintResult("count", nil, err)
				return
			}
			hasMinor = true
		}
		if countPatch != "" {
			patchVal, err = strconv.Atoi(countPatch)
			if err != nil {
				PrintResult("count", nil, err)
				return
			}
			hasPatch = true
		}

		// 构造过滤谓词
		predicate := func(v *versions.Version) bool {
			if countStable && !v.IsStable() {
				return false
			}
			if countPrerelease && !v.IsPrerelease() {
				return false
			}
			if hasMajor && v.Major() != majorVal {
				return false
			}
			if hasMinor && v.Minor() != minorVal {
				return false
			}
			if hasPatch && v.Patch() != patchVal {
				return false
			}
			return true
		}

		result := versions.Count(vs, predicate)

		data := map[string]interface{}{
			"total": len(vs),
			"count": result,
		}

		PrintResult("count", data, nil)
	},
}

func init() {
	countCmd.Flags().BoolVar(&countStable, "stable", false, "统计稳定版本数量")
	countCmd.Flags().BoolVar(&countPrerelease, "prerelease", false, "统计预发布版本数量")
	countCmd.Flags().StringVar(&countMajor, "major", "", "按 Major 版本号统计")
	countCmd.Flags().StringVar(&countMinor, "minor", "", "按 Minor 版本号统计")
	countCmd.Flags().StringVar(&countPatch, "patch", "", "按 Patch 版本号统计")
	countCmd.Flags().StringVar(&countFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(countCmd)
}
