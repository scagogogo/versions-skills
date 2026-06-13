package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var (
	groupExtraFromFile string
	groupExtraGroupID  string
	groupExtraVersion  string // for group-contains
)

// groupIDsCmd 列出所有分组 ID
var groupIDsCmd = &cobra.Command{
	Use:   "group-ids [version-strings...]",
	Short: "列出所有版本分组的 ID",
	Long: `列出所有版本分组的 ID（即 VersionNumbers 以 . 连接）。

示例:
  versions group-ids 1.0.0 1.0.0-alpha 2.0.0 2.0.0-rc1`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, groupExtraFromFile)
		if err != nil {
			PrintResult("group-ids", nil, err)
			return
		}
		svg := versions.NewSortedVersionGroups(vs)
		PrintResult("group-ids", svg.GroupIDs(), nil)
	},
}

// resolveGroupExtra 辅助函数：解析版本列表并获取指定分组
func resolveGroupExtra(args []string) (*versions.SortedVersionGroups, *versions.VersionGroup, error) {
	vs, err := ResolveVersions(args, groupExtraFromFile)
	if err != nil {
		return nil, nil, err
	}
	svg := versions.NewSortedVersionGroups(vs)
	g := svg.Get(groupExtraGroupID)
	if g == nil {
		return svg, nil, fmt.Errorf("分组 %q 不存在", groupExtraGroupID)
	}
	return svg, g, nil
}

// groupLatestCmd 获取指定分组的最新版本
var groupLatestCmd = &cobra.Command{
	Use:   "group-latest [version-strings...]",
	Short: "获取指定分组的最新版本",
	Long: `获取指定分组 ID 的最新版本。通过 --group-id 指定分组 ID。

示例:
  versions group-latest --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if groupExtraGroupID == "" {
			return fmt.Errorf("请指定 --group-id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, g, err := resolveGroupExtra(args)
		if err != nil {
			PrintResult("group-latest", nil, err)
			return
		}
		latest := g.GetLatest()
		if latest == nil {
			PrintResult("group-latest", nil, fmt.Errorf("分组 %q 为空", groupExtraGroupID))
			return
		}
		PrintResult("group-latest", verfmt.FormatVersion(latest), nil)
	},
}

// groupOldestCmd 获取指定分组的最旧版本
var groupOldestCmd = &cobra.Command{
	Use:   "group-oldest [version-strings...]",
	Short: "获取指定分组的最旧版本",
	Long: `获取指定分组 ID 的最旧版本。通过 --group-id 指定分组 ID。

示例:
  versions group-oldest --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if groupExtraGroupID == "" {
			return fmt.Errorf("请指定 --group-id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, g, err := resolveGroupExtra(args)
		if err != nil {
			PrintResult("group-oldest", nil, err)
			return
		}
		oldest := g.GetOldest()
		if oldest == nil {
			PrintResult("group-oldest", nil, fmt.Errorf("分组 %q 为空", groupExtraGroupID))
			return
		}
		PrintResult("group-oldest", verfmt.FormatVersion(oldest), nil)
	},
}

// groupStableCmd 获取指定分组的稳定版本列表
var groupStableCmd = &cobra.Command{
	Use:   "group-stable [version-strings...]",
	Short: "获取指定分组的稳定版本列表",
	Long: `获取指定分组 ID 的所有稳定版本。通过 --group-id 指定分组 ID。

示例:
  versions group-stable --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if groupExtraGroupID == "" {
			return fmt.Errorf("请指定 --group-id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, g, err := resolveGroupExtra(args)
		if err != nil {
			PrintResult("group-stable", nil, err)
			return
		}
		stable := g.StableVersions()
		PrintResult("group-stable", verfmt.FormatVersionStrings(stable), nil)
	},
}

// groupPrereleaseCmd 获取指定分组的预发布版本列表
var groupPrereleaseCmd = &cobra.Command{
	Use:   "group-prerelease [version-strings...]",
	Short: "获取指定分组的预发布版本列表",
	Long: `获取指定分组 ID 的所有预发布版本。通过 --group-id 指定分组 ID。

示例:
  versions group-prerelease --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if groupExtraGroupID == "" {
			return fmt.Errorf("请指定 --group-id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, g, err := resolveGroupExtra(args)
		if err != nil {
			PrintResult("group-prerelease", nil, err)
			return
		}
		prerelease := g.PrereleaseVersions()
		PrintResult("group-prerelease", verfmt.FormatVersionStrings(prerelease), nil)
	},
}

// groupLatestStableCmd 获取指定分组的最新稳定版本
var groupLatestStableCmd = &cobra.Command{
	Use:   "group-latest-stable [version-strings...]",
	Short: "获取指定分组的最新稳定版本",
	Long: `获取指定分组 ID 的最新稳定版本。通过 --group-id 指定分组 ID。

示例:
  versions group-latest-stable --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if groupExtraGroupID == "" {
			return fmt.Errorf("请指定 --group-id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, g, err := resolveGroupExtra(args)
		if err != nil {
			PrintResult("group-latest-stable", nil, err)
			return
		}
		latest := g.LatestStable()
		if latest == nil {
			PrintResult("group-latest-stable", nil, fmt.Errorf("分组 %q 无稳定版本", groupExtraGroupID))
			return
		}
		PrintResult("group-latest-stable", verfmt.FormatVersion(latest), nil)
	},
}

// groupLatestPrereleaseCmd 获取指定分组的最新预发布版本
var groupLatestPrereleaseCmd = &cobra.Command{
	Use:   "group-latest-prerelease [version-strings...]",
	Short: "获取指定分组的最新预发布版本",
	Long: `获取指定分组 ID 的最新预发布版本。通过 --group-id 指定分组 ID。

示例:
  versions group-latest-prerelease --group-id 1.0.0 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if groupExtraGroupID == "" {
			return fmt.Errorf("请指定 --group-id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, g, err := resolveGroupExtra(args)
		if err != nil {
			PrintResult("group-latest-prerelease", nil, err)
			return
		}
		latest := g.LatestPrerelease()
		if latest == nil {
			PrintResult("group-latest-prerelease", nil, fmt.Errorf("分组 %q 无预发布版本", groupExtraGroupID))
			return
		}
		PrintResult("group-latest-prerelease", verfmt.FormatVersion(latest), nil)
	},
}

// groupContainsCmd 检查分组是否包含指定版本
var groupContainsCmd = &cobra.Command{
	Use:   "group-contains [version-strings...]",
	Short: "检查指定分组是否包含某个版本",
	Long: `检查指定分组 ID 是否包含某个版本号。通过 --group-id 和 --version 指定分组和目标版本。

示例:
  versions group-contains --group-id 1.0.0 --version 1.0.0-alpha 1.0.0-alpha 1.0.0 2.0.0`,
	Args: cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if groupExtraGroupID == "" {
			return fmt.Errorf("请指定 --group-id")
		}
		if groupExtraVersion == "" {
			return fmt.Errorf("请指定 --version")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		targetV := versions.NewVersion(groupExtraVersion)
		vs, err := ResolveVersions(args, groupExtraFromFile)
		if err != nil {
			PrintResult("group-contains", nil, err)
			return
		}
		svg := versions.NewSortedVersionGroups(vs)
		g := svg.Get(groupExtraGroupID)
		if g == nil {
			data := map[string]interface{}{
				"group_id": groupExtraGroupID,
				"version":  targetV.RawString(),
				"contains": false,
			}
			PrintResult("group-contains", data, nil)
			return
		}
		contains := g.Contains(targetV)
		data := map[string]interface{}{
			"group_id": groupExtraGroupID,
			"version":  targetV.RawString(),
			"contains": contains,
		}
		PrintResult("group-contains", data, nil)
	},
}

func init() {
	for _, c := range []*cobra.Command{groupIDsCmd, groupLatestCmd, groupOldestCmd, groupStableCmd, groupPrereleaseCmd, groupLatestStableCmd, groupLatestPrereleaseCmd, groupContainsCmd} {
		c.Flags().StringVar(&groupExtraFromFile, "from-file", "", "从文件读取版本号列表")
	}
	// 所有需要 group-id 的命令统一注册 --group-id flag
	for _, c := range []*cobra.Command{groupLatestCmd, groupOldestCmd, groupStableCmd, groupPrereleaseCmd, groupLatestStableCmd, groupLatestPrereleaseCmd, groupContainsCmd} {
		c.Flags().StringVar(&groupExtraGroupID, "group-id", "", "分组 ID（版本号数字部分，如 '1.0.0'）")
	}
	// group-contains 专用 --version flag
	groupContainsCmd.Flags().StringVar(&groupExtraVersion, "version", "", "要检查的目标版本号")

	rootCmd.AddCommand(groupIDsCmd, groupLatestCmd, groupOldestCmd, groupStableCmd, groupPrereleaseCmd, groupLatestStableCmd, groupLatestPrereleaseCmd, groupContainsCmd)
}