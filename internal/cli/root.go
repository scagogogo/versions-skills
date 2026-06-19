package cli

import (
	"github.com/spf13/cobra"
)

// Version is the build-time version string, injected via -ldflags.
var Version = "0.0.0-dev"

// 全局 flag 变量
var (
	format string
	quiet  bool
)

// rootCmd 是 CLI 的根命令
var rootCmd = &cobra.Command{
	Use:   "versions",
	Short: "版本号解析、比较、排序、分组、约束检查和可视化的命令行工具",
	Long: `versions 是一个版本号操作命令行工具，基于 github.com/scagogogo/versions-skills SDK。

支持版本号的解析、验证、比较、排序、分组、范围查询、约束检查、
集合运算、可视化和文件操作。默认输出 JSON 格式，便于 AI 和脚本集成。`,
	Version:       Version,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "json", "输出格式: json|table|text")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "静默模式，仅输出数据")
}

// Execute 执行根命令
func Execute() error {
	return rootCmd.Execute()
}

// getFormat 获取全局格式设置
func getFormat() string {
	return format
}

// isQuiet 获取静默模式设置
func isQuiet() bool {
	return quiet
}
