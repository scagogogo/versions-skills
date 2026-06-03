package versions

import (
	"strconv"
	"strings"

	compare_anything "github.com/golang-infrastructure/go-compare-anything"
)

const (
	// DefaultVersionDelimiter 默认的版本数字分隔符
	//
	// 这是用于连接版本号数字部分的标准分隔符。即使原始版本使用了其它分隔符（如"/"），
	// 解析完毕后也会被统一为这个标准分隔符（即"."）。
	DefaultVersionDelimiter = "."
)

// VersionNumbers 表示版本号中的数字部分
//
// VersionNumbers 是一个整数切片类型，用于表示和操作版本号的数字部分。
// 例如对于版本号 "v1.2.3-beta1"，其 VersionNumbers 为 [1,2,3]。
// 它实现了 Comparable 接口，支持版本号数字部分的比较操作。
//
// 使用示例:
//
//	// 创建一个版本号数字部分
//	numbers := versions.NewVersionNumbers([]int{1, 2, 3})
//
//	// 获取版本组ID
//	groupID := numbers.BuildGroupID() // 返回 "1.2.3"
//
//	// 比较两个版本号数字部分
//	other := versions.NewVersionNumbers([]int{1, 3, 0})
//	if numbers.CompareTo(other) < 0 {
//	    fmt.Println("1.2.3 比 1.3.0 旧")
//	}
type VersionNumbers []int

var _ compare_anything.Comparable[[]int] = VersionNumbers{}

// NewVersionNumbers 创建一个新的 VersionNumbers 对象
//
// 该方法将整数切片转换为 VersionNumbers 类型，便于进行版本号相关操作。
//
// 参数:
//   - versionNumbers: 表示版本号的整数切片，如 [1,2,3] 表示版本号 "1.2.3"
//
// 返回:
//   - VersionNumbers: 一个新的 VersionNumbers 对象
//
// 使用示例:
//
//	numbers := versions.NewVersionNumbers([]int{1, 2, 3})
func NewVersionNumbers(versionNumbers []int) VersionNumbers {
	return VersionNumbers(versionNumbers)
}

// CompareTo 比较两个版本号数字部分的大小
//
// 该方法按照版本号比较规则，从左到右逐位比较两个版本号数字部分的大小。
// 对于不同长度的版本号，首先比较共有部分，如果共有部分相等，则较长的版本号更大。
//
// 参数:
//   - target: 要比较的目标版本号数字部分
//
// 返回:
//   - int: 如果当前版本小于目标版本，返回负数；如果相等，返回0；如果大于，返回正数
//
// 使用示例:
//
//	v1 := versions.NewVersionNumbers([]int{1, 0, 0})
//	v2 := versions.NewVersionNumbers([]int{1, 1, 0})
//
//	result := v1.CompareTo(v2)
//	if result < 0 {
//	    fmt.Println("v1 比 v2 旧")
//	} else if result > 0 {
//	    fmt.Println("v1 比 v2 新")
//	} else {
//	    fmt.Println("v1 和 v2 相等")
//	}
func (x VersionNumbers) CompareTo(target []int) int {
	minLen := len(x)
	if len(target) < minLen {
		minLen = len(target)
	}
	for i := 0; i < minLen; i++ {
		if x[i] < target[i] {
			return -1
		} else if x[i] > target[i] {
			return 1
		}
	}
	// 共有部分相等，比较长度
	if len(x) < len(target) {
		return -1
	} else if len(x) > len(target) {
		return 1
	}
	return 0
}

// BuildGroupID 根据版本号数字部分构造版本组的ID
//
// 该方法将版本号数字部分以默认分隔符（"."）连接成字符串，用于表示版本组ID。
// 版本组ID可用于对相同主版本号的版本进行分组和管理。
//
// 返回:
//   - string: 版本组ID字符串，例如 "1.2.3"
//
// 使用示例:
//
//	numbers := versions.NewVersionNumbers([]int{1, 2, 3})
//	groupID := numbers.BuildGroupID() // 返回 "1.2.3"
//
//	// 可用于版本分组
//	versionGroups := make(map[string][]*Version)
//	for _, version := range allVersions {
//	    groupID := version.VersionNumbers.BuildGroupID()
//	    versionGroups[groupID] = append(versionGroups[groupID], version)
//	}
func (x VersionNumbers) BuildGroupID() string {
	s := strings.Builder{}
	for i, v := range x {
		s.WriteString(strconv.Itoa(v))
		if i+1 != len(x) {
			s.WriteString(DefaultVersionDelimiter)
		}
	}
	return s.String()
}

// Len 返回版本号数字部分的长度
func (x VersionNumbers) Len() int {
	return len(x)
}

// At 返回指定索引位置的版本号数字
//
// 如果索引越界则返回 0。
//
// 参数:
//   - index: 从 0 开始的索引
//
// 返回:
//   - int: 对应位置的版本号数字，越界返回 0
func (x VersionNumbers) At(index int) int {
	if index < 0 || index >= len(x) {
		return 0
	}
	return x[index]
}

// String 返回版本号数字部分的字符串表示
//
// 等价于 BuildGroupID()，提供更符合 Go 惯例的 String() 方法。
//
// 返回:
//   - string: 如 "1.2.3"
func (x VersionNumbers) String() string {
	return x.BuildGroupID()
}

// Equals 判断两个版本号数字部分是否相等
//
// 参数:
//   - target: 目标版本号数字部分
//
// 返回:
//   - bool: 如果完全相等则返回 true
func (x VersionNumbers) Equals(target VersionNumbers) bool {
	if len(x) != len(target) {
		return false
	}
	for i := range x {
		if x[i] != target[i] {
			return false
		}
	}
	return true
}
