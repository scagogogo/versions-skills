package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var (
	buildPrefix  string
	buildMajor   string
	buildMinor   string
	buildPatch   string
	buildSuffix  string
	buildNumbers string
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "构建版本号字符串",
	Long: `通过指定各组成部分构建版本号字符串。

可使用 --numbers 直接指定完整的数字部分（逗号分隔），或使用 --major/--minor/--patch 分别指定。
--numbers 与 --major/--minor/--patch 同时指定时，--numbers 优先生效。

示例:
  versions build --major 1 --minor 2 --patch 3
  versions build --prefix v --major 1 --minor 0 --suffix -alpha1
  versions build --numbers 1,2,3,4 --prefix v`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		builder := versions.NewVersionBuilder()

		if buildPrefix != "" {
			builder.Prefix(buildPrefix)
		}

		// --numbers 优先于 --major/--minor/--patch
		if buildNumbers != "" {
			parts := strings.Split(buildNumbers, ",")
			numbers := make([]int, len(parts))
			for i, p := range parts {
				n, err := strconv.Atoi(strings.TrimSpace(p))
				if err != nil {
					PrintResult("build", nil, fmt.Errorf("无效的数字: %s", p))
					return
				}
				numbers[i] = n
			}
			builder.Numbers(numbers)
		} else {
			if buildMajor != "" {
				major, err := strconv.Atoi(buildMajor)
				if err != nil {
					PrintResult("build", nil, fmt.Errorf("无效的 major 值: %s", buildMajor))
					return
				}
				builder.Major(major)
			}

			if buildMinor != "" {
				minor, err := strconv.Atoi(buildMinor)
				if err != nil {
					PrintResult("build", nil, fmt.Errorf("无效的 minor 值: %s", buildMinor))
					return
				}
				builder.Minor(minor)
			}

			if buildPatch != "" {
				patch, err := strconv.Atoi(buildPatch)
				if err != nil {
					PrintResult("build", nil, fmt.Errorf("无效的 patch 值: %s", buildPatch))
					return
				}
				builder.Patch(patch)
			}
		}

		if buildSuffix != "" {
			builder.Suffix(buildSuffix)
		}

		v := builder.Build()

		data := map[string]interface{}{
			"raw":   v.RawString(),
			"valid": v.IsValid(),
		}

		PrintResult("build", data, nil)
	},
}

func init() {
	buildCmd.Flags().StringVar(&buildPrefix, "prefix", "", "版本前缀（如 'v'）")
	buildCmd.Flags().StringVar(&buildMajor, "major", "", "Major 版本号")
	buildCmd.Flags().StringVar(&buildMinor, "minor", "", "Minor 版本号")
	buildCmd.Flags().StringVar(&buildPatch, "patch", "", "Patch 版本号")
	buildCmd.Flags().StringVar(&buildSuffix, "suffix", "", "版本后缀（如 '-beta1'）")
	buildCmd.Flags().StringVar(&buildNumbers, "numbers", "", "完整数字部分，逗号分隔（如 '1,2,3,4'）")
	rootCmd.AddCommand(buildCmd)
}
