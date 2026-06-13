package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

// readCmd 从文件读取版本号列表
var readCmd = &cobra.Command{
	Use:   "read <filepath>",
	Short: "从文件读取版本号列表",
	Long: `从文件中读取版本号列表，每行一个版本字符串。

示例:
  versions read versions.txt`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := versions.ReadVersionsFromFile(args[0])
		if err != nil {
			PrintResult("read", nil, fmt.Errorf("读取文件失败: %w", err))
			return
		}

		data := map[string]interface{}{
			"filepath": args[0],
			"count":    len(vs),
			"versions": verfmt.FormatVersionStrings(vs),
		}

		PrintResult("read", data, nil)
	},
}

// writeCmd 将版本号列表写入文件
var writeCmd = &cobra.Command{
	Use:   "write <filepath> [version-strings...]",
	Short: "将版本号列表写入文件",
	Long: `将版本号列表写入文件，每行一个版本字符串。版本号会先排序再写入。

示例:
  versions write output.txt 1.0.0 2.0.0 1.5.0
  versions write output.txt --from-file input.txt`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]
		vs, err := ResolveVersions(args[1:], writeFromFile)
		if err != nil {
			PrintResult("write", nil, err)
			return
		}

		// 排序后写入
		sorted := versions.SortVersionSlice(vs)

		if err := versions.WriteVersionsToFile(sorted, filepath); err != nil {
			PrintResult("write", nil, fmt.Errorf("写入文件失败: %w", err))
			return
		}

		data := map[string]interface{}{
			"filepath": filepath,
			"count":    len(sorted),
		}

		PrintResult("write", data, nil)
	},
}

var writeFromFile string

func init() {
	writeCmd.Flags().StringVar(&writeFromFile, "from-file", "", "从文件读取版本号列表（用于写入到目标文件）")
	rootCmd.AddCommand(readCmd, writeCmd, readStringsCmd)
}

// readStringsCmd 从文件读取原始版本字符串列表（不解析为 Version）
var readStringsCmd = &cobra.Command{
	Use:   "read-strings <filepath>",
	Short: "从文件读取原始版本字符串列表（不解析）",
	Long: `从文件中读取版本字符串列表，每行一个字符串，不做任何解析或验证。

与 read 命令不同，此命令仅返回原始字符串列表。

示例:
  versions read-strings versions.txt`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		strs, err := versions.ReadVersionsStringFromFile(args[0])
		if err != nil {
			PrintResult("read-strings", nil, fmt.Errorf("读取文件失败: %w", err))
			return
		}

		data := map[string]interface{}{
			"filepath": args[0],
			"count":    len(strs),
			"strings":  strs,
		}

		PrintResult("read-strings", data, nil)
	},
}
