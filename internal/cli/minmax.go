package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var minmaxFromFile string

// minCmd 查找最小版本
var minCmd = &cobra.Command{
	Use:   "min [version-strings...]",
	Short: "查找最小（最旧）版本",
	Long: `从版本号列表中查找最小（最旧）的版本。

示例:
  versions min 2.0.0 1.0.0 1.5.0
  cat versions.txt | versions min`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersionsStrict(args, minmaxFromFile)
		if err != nil {
			PrintResult("min", nil, err)
			return
		}
		result := versions.Min(vs)
		if result == nil {
			PrintResult("min", nil, fmt.Errorf("未找到有效版本"))
			return
		}
		PrintResult("min", verfmt.FormatVersion(result), nil)
	},
}

// maxCmd 查找最大版本
var maxCmd = &cobra.Command{
	Use:   "max [version-strings...]",
	Short: "查找最大（最新）版本",
	Long: `从版本号列表中查找最大（最新）的版本。

示例:
  versions max 2.0.0 1.0.0 1.5.0
  cat versions.txt | versions max`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersionsStrict(args, minmaxFromFile)
		if err != nil {
			PrintResult("max", nil, err)
			return
		}
		result := versions.Max(vs)
		if result == nil {
			PrintResult("max", nil, fmt.Errorf("未找到有效版本"))
			return
		}
		PrintResult("max", verfmt.FormatVersion(result), nil)
	},
}

// latestStableCmd 查找最新稳定版本
var latestStableCmd = &cobra.Command{
	Use:   "latest-stable [version-strings...]",
	Short: "查找最新稳定版本",
	Long: `从版本号列表中查找最新的稳定版本（不含预发布标识）。

示例:
  versions latest-stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersionsStrict(args, minmaxFromFile)
		if err != nil {
			PrintResult("latest-stable", nil, err)
			return
		}
		result := versions.LatestStable(vs)
		if result == nil {
			PrintResult("latest-stable", nil, fmt.Errorf("未找到稳定版本"))
			return
		}
		PrintResult("latest-stable", verfmt.FormatVersion(result), nil)
	},
}

// latestPrereleaseCmd 查找最新预发布版本
var latestPrereleaseCmd = &cobra.Command{
	Use:   "latest-prerelease [version-strings...]",
	Short: "查找最新预发布版本",
	Long: `从版本号列表中查找最新的预发布版本（含 alpha/beta/rc 等标识）。

示例:
  versions latest-prerelease 1.0.0-alpha 1.0.0-beta 1.0.0 2.0.0-rc1`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersionsStrict(args, minmaxFromFile)
		if err != nil {
			PrintResult("latest-prerelease", nil, err)
			return
		}
		result := versions.LatestPrerelease(vs)
		if result == nil {
			PrintResult("latest-prerelease", nil, fmt.Errorf("未找到预发布版本"))
			return
		}
		PrintResult("latest-prerelease", verfmt.FormatVersion(result), nil)
	},
}

func init() {
	for _, c := range []*cobra.Command{minCmd, maxCmd, latestStableCmd, latestPrereleaseCmd} {
		c.Flags().StringVar(&minmaxFromFile, "from-file", "", "从文件读取版本号列表")
	}
	rootCmd.AddCommand(minCmd, maxCmd, latestStableCmd, latestPrereleaseCmd)
}
