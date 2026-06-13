package cli

import (
	"fmt"

	"github.com/scagogogo/versions-skills"
	"github.com/spf13/cobra"
)

var constraintType string

var constraintCmd = &cobra.Command{
	Use:   "constraint <expression> <version>",
	Short: "检查版本号是否满足约束表达式",
	Long: `检查版本号是否满足给定的约束表达式。

支持三种约束类型:
  - single: 单个 Constraint，如 ">=1.0.0"
  - set (默认): ConstraintSet，逗号分隔，AND 逻辑
  - union: ConstraintUnion，|| 分隔，OR 逻辑

支持的运算符:
  =, !=, >, >=, <, <=, ^ (caret), ~ (tilde), x/X/* (wildcard)

示例:
  versions constraint ">=1.0.0" 1.5.0
  versions constraint ">=1.0.0" 1.5.0 --type single
  versions constraint ">=1.0.0,<2.0.0" 1.5.0
  versions constraint ">=1.0.0 || >=3.0.0" 3.5.0 --type union`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		expr := args[0]
		versionStr := args[1]

		v := versions.NewVersion(versionStr)
		if !v.IsValid() {
			PrintResult("constraint", nil, fmt.Errorf("无效的版本号: %s", versionStr))
			return
		}

		var satisfied bool

		switch constraintType {
		case "single":
			c, err := versions.ParseConstraint(expr)
			if err != nil {
				PrintResult("constraint", nil, fmt.Errorf("解析约束表达式失败: %w", err))
				return
			}
			satisfied = c.Match(v)

		case "union":
			cu, err := versions.ParseConstraintUnion(expr)
			if err != nil {
				PrintResult("constraint", nil, fmt.Errorf("解析约束表达式失败: %w", err))
				return
			}
			satisfied = cu.Satisfies(v)

		case "set":
			fallthrough

		default:
			cs, err := versions.ParseConstraintSet(expr)
			if err != nil {
				PrintResult("constraint", nil, fmt.Errorf("解析约束表达式失败: %w", err))
				return
			}
			satisfied = cs.Satisfies(v)
		}

		data := map[string]interface{}{
			"expression": expr,
			"version":    v.RawString(),
			"satisfied":  satisfied,
			"type":       constraintType,
		}

		PrintResult("constraint", data, nil)
	},
}

func init() {
	constraintCmd.Flags().StringVar(&constraintType, "type", "set", "约束类型: single|set|union")
	rootCmd.AddCommand(constraintCmd)
}
