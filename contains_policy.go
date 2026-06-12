package versions

// ContainsPolicy 用于控制版本查询时的包含/排除策略
//
// 该类型定义了在版本过滤或查询操作中，是否应包含或排除特定版本的策略选项。
// 它作为枚举类型使用，提供了三种可能的状态：未指定、包含和排除。
//
// 使用示例:
//
//	// 使用包含策略进行版本过滤
//	filter := &VersionFilter{
//	    Contains: "beta",
//	    ContainsPolicy: ContainsPolicyYes,
//	}
//
//	// 使用排除策略进行版本过滤
//	filter := &VersionFilter{
//	    Contains: "snapshot",
//	    ContainsPolicy: ContainsPolicyNo,
//	}
type ContainsPolicy int

const (
	// ContainsPolicyNone 表示未指定包含策略
	//
	// 当设置为此值时，版本查询不会基于包含条件进行过滤
	ContainsPolicyNone ContainsPolicy = iota

	// ContainsPolicyYes 表示包含匹配的版本
	//
	// 当设置为此值时，只有包含指定字符串的版本才会被包含在结果中
	ContainsPolicyYes

	// ContainsPolicyNo 表示排除匹配的版本
	//
	// 当设置为此值时，包含指定字符串的版本将被排除在结果之外
	ContainsPolicyNo
)

// String 返回包含策略的可读名称
//
// 实现 fmt.Stringer 接口。
//
// 返回:
//   - string: 策略名称，如 "none"、"yes"、"no"
func (p ContainsPolicy) String() string {
	switch p {
	case ContainsPolicyYes:
		return "yes"
	case ContainsPolicyNo:
		return "no"
	default:
		return "none"
	}
}
