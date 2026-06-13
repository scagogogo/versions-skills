package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var (
	bumpMajor bool
	bumpMinor bool
	bumpPatch bool
)

var bumpCmd = &cobra.Command{
	Use:   "bump <version-string>",
	Short: "递增版本号",
	Long: `递增版本号的指定部分，并清除后缀。

  --major: 递增 Major，Minor 和 Patch 清零，后缀清除 (1.2.3 → 2.0.0)
  --minor: 递增 Minor，Patch 清零，后缀清除 (1.2.3 → 1.3.0)
  --patch: 递增 Patch，后缀清除 (1.2.3 → 1.2.4)

必须指定其中一个递增类型。

示例:
  versions bump 1.2.3 --major    # 2.0.0
  versions bump 1.2.3 --minor    # 1.3.0
  versions bump 1.2.3 --patch    # 1.2.4`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, err := ParseValidVersion(args[0])
		if err != nil {
			PrintResult("bump", nil, err)
			return
		}

		var result *versions.Version

		switch {
		case bumpMajor:
			result = v.BumpMajor()
		case bumpMinor:
			result = v.BumpMinor()
		case bumpPatch:
			result = v.BumpPatch()
		default:
			PrintResult("bump", nil, fmt.Errorf("请指定递增类型: --major, --minor 或 --patch"))
			return
		}

		data := map[string]interface{}{
			"original": v.RawString(),
			"bumped":   result.RawString(),
			"type":     getBumpType(),
		}

		PrintResult("bump", data, nil)
	},
}

func getBumpType() string {
	switch {
	case bumpMajor:
		return "major"
	case bumpMinor:
		return "minor"
	case bumpPatch:
		return "patch"
	default:
		return ""
	}
}

func init() {
	bumpCmd.Flags().BoolVar(&bumpMajor, "major", false, "递增 Major 版本号")
	bumpCmd.Flags().BoolVar(&bumpMinor, "minor", false, "递增 Minor 版本号")
	bumpCmd.Flags().BoolVar(&bumpPatch, "patch", false, "递增 Patch 版本号")
	rootCmd.AddCommand(bumpCmd)
}