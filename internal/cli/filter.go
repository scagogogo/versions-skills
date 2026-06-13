package cli

import (
	"fmt"
	"strconv"

	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var (
	filterStable         bool
	filterPrerelease     bool
	filterMajor          string
	filterMinor          string
	filterPatch          string
	filterPrefix         string
	filterSuffix         string
	filterConstraint     string
	filterConstraintType string
	filterFromFile       string
)

var filterCmd = &cobra.Command{
	Use:   "filter [version-strings...]",
	Short: "按条件过滤版本号列表",
	Long: `按各种条件过滤版本号列表。可以组合多个过滤条件（AND 逻辑）。

支持通过参数、--from-file 或 stdin 提供版本号列表。
--constraint 支持 single（单个约束）、set（默认，逗号分隔 AND）和 union（|| 分隔 OR）三种类型。

示例:
  versions filter --stable 1.0.0 1.0.0-beta 2.0.0
  versions filter --prerelease 1.0.0-alpha 1.0.0 1.0.0-beta
  versions filter --major 1 1.0.0 2.0.0 1.5.0
  versions filter --constraint ">=1.0.0,<2.0.0" 1.0.0 1.5.0 2.0.0
  versions filter --constraint ">=1.0.0" --constraint-type single 1.0.0 1.5.0 2.0.0
  versions filter --constraint ">=1.0.0 || >=3.0.0" --constraint-type union 1.0.0 2.0.0 3.0.0
  cat versions.txt | versions filter --stable`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, filterFromFile)
		if err != nil {
			PrintResult("filter", nil, err)
			return
		}

		result := vs

		// 按稳定版过滤
		if filterStable {
			result = versions.FilterByStable(result)
		}

		// 按预发布版过滤
		if filterPrerelease {
			result = versions.FilterByPrerelease(result)
		}

		// 按 Major 过滤
		if filterMajor != "" {
			major, err := strconv.Atoi(filterMajor)
			if err != nil {
				PrintResult("filter", nil, fmt.Errorf("无效的 major 值: %s", filterMajor))
				return
			}
			result = versions.FilterByMajor(result, major)
		}

		// 按 Minor 过滤
		if filterMinor != "" {
			minor, err := strconv.Atoi(filterMinor)
			if err != nil {
				PrintResult("filter", nil, fmt.Errorf("无效的 minor 值: %s", filterMinor))
				return
			}
			result = versions.FilterByMinor(result, minor)
		}

		// 按 Patch 过滤
		if filterPatch != "" {
			patch, err := strconv.Atoi(filterPatch)
			if err != nil {
				PrintResult("filter", nil, fmt.Errorf("无效的 patch 值: %s", filterPatch))
				return
			}
			result = versions.FilterByPatch(result, patch)
		}

		// 按 Prefix 过滤
		if filterPrefix != "" {
			result = versions.FilterByPrefix(result, filterPrefix)
		}

		// 按 Suffix 过滤
		if filterSuffix != "" {
			result = versions.FilterBySuffix(result, filterSuffix)
		}

		// 按约束表达式过滤
		if filterConstraint != "" {
			switch filterConstraintType {
			case "single":
				c, parseErr := versions.ParseConstraint(filterConstraint)
				if parseErr != nil {
					PrintResult("filter", nil, fmt.Errorf("解析约束表达式失败: %w", parseErr))
					return
				}
				result = versions.FilterByConstraint(result, c)
			case "union":
				cu, parseErr := versions.ParseConstraintUnion(filterConstraint)
				if parseErr != nil {
					PrintResult("filter", nil, fmt.Errorf("解析约束表达式失败: %w", parseErr))
					return
				}
				result = versions.Filter(result, func(v *versions.Version) bool {
					return cu.Satisfies(v)
				})
			default:
				cs, parseErr := versions.ParseConstraintSet(filterConstraint)
				if parseErr != nil {
					PrintResult("filter", nil, fmt.Errorf("解析约束表达式失败: %w", parseErr))
					return
				}
				result = versions.FilterByConstraintSet(result, cs)
			}
		}

		data := verfmt.FormatVersionStrings(result)
		PrintResult("filter", data, nil)
	},
}

func init() {
	filterCmd.Flags().BoolVar(&filterStable, "stable", false, "仅保留稳定版本")
	filterCmd.Flags().BoolVar(&filterPrerelease, "prerelease", false, "仅保留预发布版本")
	filterCmd.Flags().StringVar(&filterMajor, "major", "", "按 Major 版本号过滤")
	filterCmd.Flags().StringVar(&filterMinor, "minor", "", "按 Minor 版本号过滤")
	filterCmd.Flags().StringVar(&filterPatch, "patch", "", "按 Patch 版本号过滤")
	filterCmd.Flags().StringVar(&filterPrefix, "prefix", "", "按前缀过滤")
	filterCmd.Flags().StringVar(&filterSuffix, "suffix", "", "按后缀过滤")
	filterCmd.Flags().StringVar(&filterConstraint, "constraint", "", "按约束表达式过滤")
	filterCmd.Flags().StringVar(&filterConstraintType, "constraint-type", "set", "约束类型: single|set|union")
	filterCmd.Flags().StringVar(&filterFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(filterCmd)
}
