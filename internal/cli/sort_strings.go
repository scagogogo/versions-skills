package cli

import (
	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var (
	sortStringsDesc    bool
	sortStringsFromFile string
)

var sortStringsCmd = &cobra.Command{
	Use:   "sort-strings [version-strings...]",
	Short: "对版本号字符串列表排序（返回原始字符串，非 Version 对象）",
	Long: `对版本号字符串列表排序，返回原始字符串而非结构化 Version 对象。
适用于仅需排序结果不需要解析详情的场景。

示例:
  versions sort-strings 2.0.0 1.0.0 1.5.0
  versions sort-strings --desc 1.0 2.0 1.5
  cat versions.txt | versions sort-strings`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, sortStringsFromFile)
		if err != nil {
			PrintResult("sort-strings", nil, err)
			return
		}

		// 使用 SortVersionStringSlice
		inputStrings := make([]string, 0, len(vs))
		for _, v := range vs {
			inputStrings = append(inputStrings, v.RawString())
		}

		sorted := versions.SortVersionStringSlice(inputStrings)

		// 如果降序，反转
		if sortStringsDesc {
			ReverseSlice(sorted)
		}

		PrintResult("sort-strings", sorted, nil)
	},
}

func init() {
	sortStringsCmd.Flags().BoolVarP(&sortStringsDesc, "desc", "d", false, "降序排列")
	sortStringsCmd.Flags().StringVar(&sortStringsFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(sortStringsCmd)
}
