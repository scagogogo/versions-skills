package cli

import (
	"github.com/spf13/cobra"
)

var coreCmd = &cobra.Command{
	Use:   "core <version-string>",
	Short: "获取核心版本号（去除后缀）",
	Long: `获取版本号的核心部分，即去除预发布后缀后的版本。

例如: v1.2.3-beta1 → v1.2.3

示例:
  versions core v1.2.3-beta1
  versions core 2.0.0-rc1`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("core", nil, err)
			return
		}

		coreV := v.Core()

		data := map[string]interface{}{
			"original": v.RawString(),
			"core":     coreV.RawString(),
		}

		PrintResult("core", data, nil)
	},
}

func init() {
	rootCmd.AddCommand(coreCmd)
}
