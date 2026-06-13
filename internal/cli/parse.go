package cli

import (
	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var parseDelimiters string

var parseCmd = &cobra.Command{
	Use:   "parse <version-string>",
	Short: "解析版本字符串，显示各组成部分",
	Long: `解析版本字符串并显示其结构化组成部分，包括前缀、数字部分、后缀、后缀权重等。

可通过 --delimiters 指定自定义分隔符（默认为 "."），用于解析非标准格式的版本号。

示例:
  versions parse v1.2.3-beta1
  versions parse 2.0.0
  versions parse curl-7_85_0
  versions parse --delimiters "_-" curl-7_85_0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var v *versions.Version
		if parseDelimiters != "" {
			option := versions.ParserOption{Delimiters: parseDelimiters}
			v = versions.NewVersionWithOption(args[0], option)
		} else {
			v = versions.NewVersion(args[0])
		}
		data := verfmt.FormatVersion(v)
		if parseDelimiters != "" {
			data["delimiters"] = parseDelimiters
		}
		PrintResult("parse", data, nil)
	},
}

func init() {
	parseCmd.Flags().StringVar(&parseDelimiters, "delimiters", "", "自定义版本号分隔符（默认为 \".\"）")
	rootCmd.AddCommand(parseCmd)
}
