package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var (
	groupID       string
	groupFromFile string
)

var groupCmd = &cobra.Command{
	Use:   "group [version-strings...]",
	Short: "按版本号数字部分分组",
	Long: `按版本号数字部分（VersionNumbers）对版本列表进行分组。
数字部分相同的版本归入同一组，如 1.2.3 和 1.2.3-beta 属于同一组。

示例:
  versions group 1.0.0 1.0.0-alpha 2.0.0 1.0.0-beta 2.0.0-rc1
  versions group --id 1.0.0 1.0.0 1.0.0-alpha 2.0.0
  cat versions.txt | versions group`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, groupFromFile)
		if err != nil {
			PrintResult("group", nil, err)
			return
		}

		groupMap := versions.Group(vs)

		// 如果指定了 --id，只输出该组
		if groupID != "" {
			g, ok := groupMap[groupID]
			if !ok {
				PrintResult("group", nil, fmt.Errorf("组 %q 不存在", groupID))
				return
			}
			data := verfmt.FormatVersionGroup(g)
			PrintResult("group", data, nil)
			return
		}

		// 输出所有组
		data := verfmt.FormatVersionGroupMap(groupMap)
		PrintResult("group", data, nil)
	},
}

func init() {
	groupCmd.Flags().StringVar(&groupID, "id", "", "仅显示指定组 ID 的详情")
	groupCmd.Flags().StringVar(&groupFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(groupCmd)
}
