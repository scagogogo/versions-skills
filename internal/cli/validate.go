package cli

import (
	"os"

	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate <version-string>",
	Short: "严格验证版本字符串是否有效",
	Long: `严格验证版本字符串是否有效。有效版本返回 exit code 0，无效返回 exit code 1。

验证规则:
  - 版本必须包含数字部分
  - 数字部分不能包含负数

示例:
  versions validate 1.2.3         # 有效 (exit 0)
  versions validate not-a-version # 无效 (exit 1)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		isValid := v.IsValid()

		data := map[string]interface{}{
			"raw":   args[0],
			"valid": isValid,
		}

		PrintResult("validate", data, nil)

		if !isValid {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
