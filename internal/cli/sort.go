package cli

import (
	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var (
	sortDesc     bool
	sortFromFile string
)

var sortCmd = &cobra.Command{
	Use:   "sort [version-strings...]",
	Short: "对版本号列表进行排序",
	Long: `对版本号列表进行排序，默认升序排列。

支持通过参数、--from-file 或 stdin 提供版本号列表。

示例:
  versions sort 2.0.0 1.0.0 1.5.0
  versions sort --desc 1.0 2.0 1.5
  cat versions.txt | versions sort
  versions sort --from-file versions.txt`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, sortFromFile)
		if err != nil {
			PrintResult("sort", nil, err)
			return
		}

		sorted := versions.SortVersionSlice(vs)

		// 如果降序，反转
		if sortDesc {
			for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}

		data := verfmt.FormatVersionStrings(sorted)
		PrintResult("sort", data, nil)
	},
}

func init() {
	sortCmd.Flags().BoolVarP(&sortDesc, "desc", "d", false, "降序排列")
	sortCmd.Flags().StringVar(&sortFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(sortCmd)
}
