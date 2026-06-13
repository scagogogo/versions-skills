package cli

import (
	"bytes"

	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var (
	visualizeMaxItems int
	visualizeGroups   bool
	visualizeFromFile string
)

var visualizeCmd = &cobra.Command{
	Use:   "visualize [version-strings...]",
	Short: "可视化版本号层级结构",
	Long: `以文本树状图展示版本号的层级结构，便于理解版本间的关系。

示例:
  versions visualize 1.0.0 1.0.0-alpha 1.0.0-beta 2.0.0 2.0.0-rc1
  versions visualize --max-items 5 1.0.0 1.0.1 1.0.2 2.0.0
  versions visualize --groups 1.0.0 2.0.0
  cat versions.txt | versions visualize`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, visualizeFromFile)
		if err != nil {
			PrintResult("visualize", nil, err)
			return
		}

		var buf bytes.Buffer

		if visualizeGroups {
			versions.VisualizeVersionGroups(vs, &buf)
		} else {
			maxItems := visualizeMaxItems
			if maxItems <= 0 {
				maxItems = 0 // 0 表示不限制
			}
			versions.VisualizeVersions(vs, &buf, maxItems)
		}

		data := map[string]interface{}{
			"text":      buf.String(),
			"count":     len(vs),
			"groups":    visualizeGroups,
			"max_items": visualizeMaxItems,
		}

		PrintResult("visualize", data, nil)
	},
}

func init() {
	visualizeCmd.Flags().IntVar(&visualizeMaxItems, "max-items", 0, "每组最多显示的版本数（0=不限）")
	visualizeCmd.Flags().BoolVar(&visualizeGroups, "groups", false, "仅显示分组摘要")
	visualizeCmd.Flags().StringVar(&visualizeFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(visualizeCmd)
}
