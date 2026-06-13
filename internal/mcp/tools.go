package mcp

import (
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

// marshalJSON 安全地序列化为缩进 JSON
func marshalJSON(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// === 工具定义 ===

func (s *Server) toolParse() mcp.Tool {
	return mcp.NewTool("version_parse",
		mcp.WithDescription("解析版本字符串为结构化组件（前缀、数字、后缀等）。支持标准 semver 和多种变体如 'v1.2.3-beta1'。"),
		mcp.WithString("version_string",
			mcp.Description("要解析的版本字符串，如 '1.2.3', 'v1.2.3-beta1'"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolValidate() mcp.Tool {
	return mcp.NewTool("version_validate",
		mcp.WithDescription("验证版本字符串是否有效。有效版本须包含数字部分且数字非负。"),
		mcp.WithString("version_string",
			mcp.Description("要验证的版本字符串"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolInfo() mcp.Tool {
	return mcp.NewTool("version_info",
		mcp.WithDescription("获取版本号的完整信息，包括所有类型判断（IsPrerelease/IsStable/IsBeta 等）。"),
		mcp.WithString("version_string",
			mcp.Description("版本字符串"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolCompare() mcp.Tool {
	return mcp.NewTool("version_compare",
		mcp.WithDescription("比较两个版本号的大小关系。返回 -1(旧于)/0(相等)/1(新于)。"),
		mcp.WithString("version1",
			mcp.Description("第一个版本字符串"),
			mcp.Required(),
		),
		mcp.WithString("version2",
			mcp.Description("第二个版本字符串"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolSort() mcp.Tool {
	return mcp.NewTool("version_sort",
		mcp.WithDescription("对版本号列表进行排序。默认升序，可设置降序。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
		mcp.WithBoolean("descending",
			mcp.Description("是否降序排列（默认 false）"),
		),
	)
}

func (s *Server) toolGroup() mcp.Tool {
	return mcp.NewTool("version_group",
		mcp.WithDescription("按版本号数字部分分组。数字部分相同的版本归入同一组。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolConstraintCheck() mcp.Tool {
	return mcp.NewTool("version_constraint_check",
		mcp.WithDescription("检查版本号是否满足约束表达式。支持运算符: =, !=, >, >=, <, <=, ^(caret), ~(tilde), x/X/*(通配符)。"),
		mcp.WithString("expression",
			mcp.Description("约束表达式，如 '>=1.0.0,<2.0.0'"),
			mcp.Required(),
		),
		mcp.WithString("version",
			mcp.Description("要检查的版本字符串"),
			mcp.Required(),
		),
		mcp.WithString("type",
			mcp.Description("约束类型: set(逗号分隔AND,默认) 或 union(||分隔OR)"),
		),
	)
}

func (s *Server) toolRangeQuery() mcp.Tool {
	return mcp.NewTool("version_range_query",
		mcp.WithDescription("查询指定范围内的版本号列表。"),
		mcp.WithString("start",
			mcp.Description("范围起始版本"),
			mcp.Required(),
		),
		mcp.WithString("end",
			mcp.Description("范围结束版本"),
			mcp.Required(),
		),
		mcp.WithArray("versions",
			mcp.Description("待查询的版本字符串数组"),
			mcp.Required(),
		),
		mcp.WithBoolean("include_start",
			mcp.Description("是否包含起始版本（默认 true）"),
		),
		mcp.WithBoolean("include_end",
			mcp.Description("是否包含结束版本（默认 false）"),
		),
	)
}

func (s *Server) toolFilter() mcp.Tool {
	return mcp.NewTool("version_filter",
		mcp.WithDescription("按条件过滤版本号列表。多个条件为 AND 逻辑。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
		mcp.WithBoolean("stable",
			mcp.Description("仅保留稳定版本"),
		),
		mcp.WithBoolean("prerelease",
			mcp.Description("仅保留预发布版本"),
		),
		mcp.WithNumber("major",
			mcp.Description("按 Major 版本号过滤"),
		),
		mcp.WithNumber("minor",
			mcp.Description("按 Minor 版本号过滤"),
		),
		mcp.WithNumber("patch",
			mcp.Description("按 Patch 版本号过滤"),
		),
		mcp.WithString("prefix",
			mcp.Description("按前缀过滤"),
		),
		mcp.WithString("suffix",
			mcp.Description("按后缀过滤"),
		),
		mcp.WithString("constraint",
			mcp.Description("按约束表达式过滤（ConstraintSet 语法）"),
		),
	)
}

func (s *Server) toolMin() mcp.Tool {
	return mcp.NewTool("version_min",
		mcp.WithDescription("从版本列表中查找最小（最旧）版本。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolMax() mcp.Tool {
	return mcp.NewTool("version_max",
		mcp.WithDescription("从版本列表中查找最大（最新）版本。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolLatestStable() mcp.Tool {
	return mcp.NewTool("version_latest_stable",
		mcp.WithDescription("从版本列表中查找最新的稳定版本。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolLatestPrerelease() mcp.Tool {
	return mcp.NewTool("version_latest_prerelease",
		mcp.WithDescription("从版本列表中查找最新的预发布版本。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolUnique() mcp.Tool {
	return mcp.NewTool("version_unique",
		mcp.WithDescription("去除版本列表中的重复项。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolSetOperation() mcp.Tool {
	return mcp.NewTool("version_set_operation",
		mcp.WithDescription("对两个版本集合执行集合运算：差集(difference)、交集(intersection)、并集(union)。"),
		mcp.WithString("operation",
			mcp.Description("集合运算类型: difference|intersection|union"),
			mcp.Required(),
		),
		mcp.WithArray("set_a",
			mcp.Description("集合 A 的版本字符串数组"),
			mcp.Required(),
		),
		mcp.WithArray("set_b",
			mcp.Description("集合 B 的版本字符串数组"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolVisualize() mcp.Tool {
	return mcp.NewTool("version_visualize",
		mcp.WithDescription("生成版本层级的文本可视化树状图。"),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
		mcp.WithNumber("max_items_per_group",
			mcp.Description("每组最多显示的版本数（0=不限，默认 0）"),
		),
		mcp.WithBoolean("groups_only",
			mcp.Description("仅显示分组摘要"),
		),
	)
}

func (s *Server) toolBuild() mcp.Tool {
	return mcp.NewTool("version_build",
		mcp.WithDescription("从各组成部分构建版本字符串。"),
		mcp.WithString("prefix",
			mcp.Description("版本前缀（如 'v'）"),
		),
		mcp.WithNumber("major",
			mcp.Description("Major 版本号"),
		),
		mcp.WithNumber("minor",
			mcp.Description("Minor 版本号"),
		),
		mcp.WithNumber("patch",
			mcp.Description("Patch 版本号"),
		),
		mcp.WithString("suffix",
			mcp.Description("版本后缀（如 '-beta1'）"),
		),
	)
}

func (s *Server) toolBump() mcp.Tool {
	return mcp.NewTool("version_bump",
		mcp.WithDescription("递增版本号的指定部分并清除后缀。major: 1.2.3→2.0.0; minor: 1.2.3→1.3.0; patch: 1.2.3→1.2.4。"),
		mcp.WithString("version_string",
			mcp.Description("原始版本字符串"),
			mcp.Required(),
		),
		mcp.WithString("bump_type",
			mcp.Description("递增类型: major|minor|patch"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolCore() mcp.Tool {
	return mcp.NewTool("version_core",
		mcp.WithDescription("获取版本号的核心部分（去除预发布后缀）。如 v1.2.3-beta1 → v1.2.3。"),
		mcp.WithString("version_string",
			mcp.Description("版本字符串"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolReadFile() mcp.Tool {
	return mcp.NewTool("version_read_file",
		mcp.WithDescription("从文件中读取版本号列表（每行一个版本字符串）。"),
		mcp.WithString("filepath",
			mcp.Description("文件路径"),
			mcp.Required(),
		),
	)
}

func (s *Server) toolWriteFile() mcp.Tool {
	return mcp.NewTool("version_write_file",
		mcp.WithDescription("将版本号列表排序后写入文件（每行一个版本字符串）。"),
		mcp.WithString("filepath",
			mcp.Description("文件路径"),
			mcp.Required(),
		),
		mcp.WithArray("versions",
			mcp.Description("版本字符串数组"),
			mcp.Required(),
		),
	)
}
