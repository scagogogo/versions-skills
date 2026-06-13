package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare <version1> <version2>",
	Short: "比较两个版本号",
	Long: `比较两个版本号的大小关系。

返回结果:
  -1 表示 version1 旧于 version2
   0 表示两个版本相等
   1 表示 version1 新于 version2

示例:
  versions compare 1.2.3 2.0.0     # 1.2.3 旧于 2.0.0
  versions compare v1.0 v1.0.0     # 相等
  versions compare 2.0 1.0         # 2.0 新于 1.0`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v1, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("compare", nil, err)
			return
		}
		v2, err := ParseValidVersion(args[1])
		if err != nil {
			PrintResult("compare", nil, err)
			return
		}

		result := v1.CompareTo(v2)

		var description string
		switch {
		case result < 0:
			description = fmt.Sprintf("%s 旧于 %s", v1.RawString(), v2.RawString())
		case result > 0:
			description = fmt.Sprintf("%s 新于 %s", v1.RawString(), v2.RawString())
		default:
			description = fmt.Sprintf("%s 等于 %s", v1.RawString(), v2.RawString())
		}

		data := map[string]interface{}{
			"v1":          v1.RawString(),
			"v2":          v2.RawString(),
			"result":      result,
			"description": description,
		}

		PrintResult("compare", data, nil)
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}
