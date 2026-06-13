package cli

import (
	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <version-string>",
	Short: "显示版本号的完整信息",
	Long: `显示版本号的所有详细信息，包括结构化组成部分和所有类型判断（Is* 方法）。

示例:
  versions info v1.2.3-beta1
  versions info 2.0.0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		data := verfmt.FormatVersionDetailed(v)
		PrintResult("info", data, nil)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
