package cli

import (
	"fmt"

	"github.com/golang-infrastructure/go-tuple"
	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var (
	rangeIncludeStart bool
	rangeIncludeEnd   bool
	rangeFromFile     string
)

var rangeCmd = &cobra.Command{
	Use:   "range <start> <end> [version-strings...]",
	Short: "查询指定范围内的版本号",
	Long: `查询指定范围内的版本号列表。

边界包含策略:
  --include-start (默认 true): 包含起始版本
  --include-end (默认 false): 包含结束版本

示例:
  versions range 1.0.0 3.0.0 1.0.0 1.5.0 2.0.0 3.0.0 4.0.0
  versions range 1.0.0 3.0.0 --include-end 1.0.0 3.0.0 4.0.0
  cat versions.txt | versions range 1.0.0 2.0.0`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		startV := versions.NewVersion(args[0])
		endV := versions.NewVersion(args[1])

		if !startV.IsValid() {
			PrintResult("range", nil, fmt.Errorf("无效的起始版本: %s", args[0]))
			return
		}
		if !endV.IsValid() {
			PrintResult("range", nil, fmt.Errorf("无效的结束版本: %s", args[1]))
			return
		}

		// 剩余参数作为版本列表
		vs, err := ResolveVersions(args[2:], rangeFromFile)
		if err != nil {
			PrintResult("range", nil, err)
			return
		}

		// 构建 ContainsPolicy
		startPolicy := versions.ContainsPolicyNo
		if rangeIncludeStart {
			startPolicy = versions.ContainsPolicyYes
		}
		endPolicy := versions.ContainsPolicyNo
		if rangeIncludeEnd {
			endPolicy = versions.ContainsPolicyYes
		}

		// 使用 SortedVersionGroups 进行范围查询
		svg := versions.NewSortedVersionGroups(vs)
		result := svg.QueryRange(
			tuple.New2(startV, startPolicy),
			tuple.New2(endV, endPolicy),
		)

		data := map[string]interface{}{
			"start":         startV.RawString(),
			"end":           endV.RawString(),
			"include_start": rangeIncludeStart,
			"include_end":   rangeIncludeEnd,
			"versions":      verfmt.FormatVersionStrings(result),
			"count":         len(result),
		}

		PrintResult("range", data, nil)
	},
}

func init() {
	rangeCmd.Flags().BoolVar(&rangeIncludeStart, "include-start", true, "包含起始版本")
	rangeCmd.Flags().BoolVar(&rangeIncludeEnd, "include-end", false, "包含结束版本")
	rangeCmd.Flags().StringVar(&rangeFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(rangeCmd)
}
