package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set-prefix <version-string> <prefix>",
	Short: "修改版本号的前缀（不可变修改，返回新版本）",
	Long: `修改版本号的前缀，返回新版本字符串。原版本不变。

示例:
  versions set-prefix 1.2.3 v       # v1.2.3
  versions set-prefix v1.2.3 ""     # 1.2.3`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		if !v.IsValid() {
			PrintResult("set-prefix", nil, fmt.Errorf("无效的版本号: %s", args[0]))
			return
		}
		result := v.WithPrefix(args[1])
		data := map[string]interface{}{
			"original":  v.RawString(),
			"modified":  result.RawString(),
			"operation": "set-prefix",
			"value":     args[1],
		}
		PrintResult("set-prefix", data, nil)
	},
}

var setSuffixCmd = &cobra.Command{
	Use:   "set-suffix <version-string> <suffix>",
	Short: "修改版本号的后缀（不可变修改，返回新版本）",
	Long: `修改版本号的后缀，返回新版本字符串。原版本不变。

示例:
  versions set-suffix 1.2.3 -beta1   # 1.2.3-beta1
  versions set-suffix 1.2.3-beta1 "" # 1.2.3`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		if !v.IsValid() {
			PrintResult("set-suffix", nil, fmt.Errorf("无效的版本号: %s", args[0]))
			return
		}
		result := v.WithSuffix(args[1])
		data := map[string]interface{}{
			"original":  v.RawString(),
			"modified":  result.RawString(),
			"operation": "set-suffix",
			"value":     args[1],
		}
		PrintResult("set-suffix", data, nil)
	},
}

var setMajorCmd = &cobra.Command{
	Use:   "set-major <version-string> <major>",
	Short: "修改版本号的 Major 号（不可变修改，返回新版本）",
	Long: `修改版本号的 Major 版本号，返回新版本字符串。原版本不变。

示例:
  versions set-major 1.2.3 2  # 2.2.3`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		if !v.IsValid() {
			PrintResult("set-major", nil, fmt.Errorf("无效的版本号: %s", args[0]))
			return
		}
		major, err := strconv.Atoi(args[1])
		if err != nil {
			PrintResult("set-major", nil, fmt.Errorf("无效的 Major 值: %s", args[1]))
			return
		}
		result := v.WithMajor(major)
		data := map[string]interface{}{
			"original":  v.RawString(),
			"modified":  result.RawString(),
			"operation": "set-major",
			"value":     major,
		}
		PrintResult("set-major", data, nil)
	},
}

var setMinorCmd = &cobra.Command{
	Use:   "set-minor <version-string> <minor>",
	Short: "修改版本号的 Minor 号（不可变修改，返回新版本）",
	Long: `修改版本号的 Minor 版本号，返回新版本字符串。原版本不变。

示例:
  versions set-minor 1.2.3 5  # 1.5.3`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		if !v.IsValid() {
			PrintResult("set-minor", nil, fmt.Errorf("无效的版本号: %s", args[0]))
			return
		}
		minor, err := strconv.Atoi(args[1])
		if err != nil {
			PrintResult("set-minor", nil, fmt.Errorf("无效的 Minor 值: %s", args[1]))
			return
		}
		result := v.WithMinor(minor)
		data := map[string]interface{}{
			"original":  v.RawString(),
			"modified":  result.RawString(),
			"operation": "set-minor",
			"value":     minor,
		}
		PrintResult("set-minor", data, nil)
	},
}

var setPatchCmd = &cobra.Command{
	Use:   "set-patch <version-string> <patch>",
	Short: "修改版本号的 Patch 号（不可变修改，返回新版本）",
	Long: `修改版本号的 Patch 版本号，返回新版本字符串。原版本不变。

示例:
  versions set-patch 1.2.3 9  # 1.2.9`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		if !v.IsValid() {
			PrintResult("set-patch", nil, fmt.Errorf("无效的版本号: %s", args[0]))
			return
		}
		patch, err := strconv.Atoi(args[1])
		if err != nil {
			PrintResult("set-patch", nil, fmt.Errorf("无效的 Patch 值: %s", args[1]))
			return
		}
		result := v.WithPatch(patch)
		data := map[string]interface{}{
			"original":  v.RawString(),
			"modified":  result.RawString(),
			"operation": "set-patch",
			"value":     patch,
		}
		PrintResult("set-patch", data, nil)
	},
}

var setNumbersCmd = &cobra.Command{
	Use:   "set-numbers <version-string> <numbers>",
	Short: "修改版本号的数字部分（不可变修改，返回新版本）",
	Long: `修改版本号的数字部分（VersionNumbers），返回新版本字符串。原版本不变。
numbers 参数以逗号分隔的数字字符串提供。

示例:
  versions set-numbers 1.2.3 4,5,6   # 4.5.6
  versions set-numbers v1.2.3 2,0,0  # v2.0.0`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		v := versions.NewVersion(args[0])
		if !v.IsValid() {
			PrintResult("set-numbers", nil, fmt.Errorf("无效的版本号: %s", args[0]))
			return
		}

		// 解析逗号分隔的数字
		parts := strings.Split(args[1], ",")
		numbers := make([]int, len(parts))
		for i, p := range parts {
			n, err := strconv.Atoi(strings.TrimSpace(p))
			if err != nil {
				PrintResult("set-numbers", nil, fmt.Errorf("无效的数字: %s", p))
				return
			}
			numbers[i] = n
		}

		result := v.WithNumbers(numbers)
		data := map[string]interface{}{
			"original":  v.RawString(),
			"modified":  result.RawString(),
			"operation": "set-numbers",
			"value":     numbers,
		}
		PrintResult("set-numbers", data, nil)
	},
}

func init() {
	rootCmd.AddCommand(setCmd, setSuffixCmd, setMajorCmd, setMinorCmd, setPatchCmd, setNumbersCmd)
}