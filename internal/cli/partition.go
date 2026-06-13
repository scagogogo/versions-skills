package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
	"github.com/scagogogo/versions-skills/internal/verfmt"
	"github.com/spf13/cobra"
)

var (
	partitionStable     bool
	partitionPrerelease bool
	partitionFromFile   string
)

var partitionCmd = &cobra.Command{
	Use:   "partition [version-strings...]",
	Short: "将版本号列表分为两组（满足条件和不满足条件）",
	Long: `将版本号列表按条件分成两组：满足条件和不满足条件的版本。

示例:
  versions partition --stable 1.0.0-alpha 1.0.0 1.0.0-beta 2.0.0
  versions partition --prerelease 1.0.0-alpha 1.0.0 2.0.0-rc1`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vs, err := ResolveVersions(args, partitionFromFile)
		if err != nil {
			PrintResult("partition", nil, err)
			return
		}

		var predicate func(*versions.Version) bool

		switch {
		case partitionStable:
			predicate = func(v *versions.Version) bool { return v.IsStable() }
		case partitionPrerelease:
			predicate = func(v *versions.Version) bool { return v.IsPrerelease() }
		default:
			PrintResult("partition", nil, fmt.Errorf("请指定分区条件: --stable 或 --prerelease"))
			return
		}

		matched, unmatched := versions.Partition(vs, predicate)

		data := map[string]interface{}{
			"matched":         verfmt.FormatVersionStrings(matched),
			"unmatched":       verfmt.FormatVersionStrings(unmatched),
			"matched_count":   len(matched),
			"unmatched_count": len(unmatched),
		}

		PrintResult("partition", data, nil)
	},
}

func init() {
	partitionCmd.Flags().BoolVar(&partitionStable, "stable", false, "按稳定/不稳定分区")
	partitionCmd.Flags().BoolVar(&partitionPrerelease, "prerelease", false, "按预发布/非预发布分区")
	partitionCmd.Flags().StringVar(&partitionFromFile, "from-file", "", "从文件读取版本号列表")
	rootCmd.AddCommand(partitionCmd)
}
