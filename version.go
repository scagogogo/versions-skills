package versions

import (
	"encoding/json"
	"errors"
	"time"

	compare_anything "github.com/golang-infrastructure/go-compare-anything"
)

var (
	// ErrVersionInvalid 表示版本号格式无效的错误
	//
	// 当尝试解析不符合要求的版本号字符串时返回此错误
	ErrVersionInvalid = errors.New("version invalid")
)

// Version 用于表示一个版本号
//
// Version 结构体封装了版本号的各个组成部分，包括原始字符串、发布时间、数字部分、
// 前缀和后缀。它支持版本号的解析、比较和分组等操作，实现了 Comparable 接口以便
// 进行版本排序。
//
// 一个典型的版本号格式可能为：v1.2.3-beta1，其中：
// - "v" 是前缀
// - "1.2.3" 是数字部分
// - "-beta1" 是后缀
//
// 使用示例:
//
//	// 创建一个版本对象
//	version := versions.NewVersion("v1.2.3-rc1")
//
//	// 检查版本是否有效
//	if version.IsValid() {
//	    fmt.Printf("版本号有效: %s\n", version.Raw)
//	    fmt.Printf("版本号数字部分: %v\n", version.VersionNumbers)
//	}
//
//	// 比较两个版本
//	v1 := versions.NewVersion("1.2.3")
//	v2 := versions.NewVersion("1.3.0")
//	if v1.CompareTo(v2) < 0 {
//	    fmt.Println("v1 比 v2 旧")
//	}
type Version struct {

	// Raw 原始的版本号字符串
	Raw string `json:"raw"`

	// PublicTime 此版本的发布时间
	PublicTime time.Time `json:"public_time"`

	// VersionNumbers 版本号中的数字部分
	// 例如对于版本号 "v1.2.3-beta1"，VersionNumbers 为 [1,2,3]
	VersionNumbers VersionNumbers `json:"version_numbers"`

	// Prefix 版本号数字部分之前的前缀
	// 例如对于版本号 "v1.2.3"，Prefix 为 "v"
	Prefix VersionPrefix `json:"prefix"`

	// Suffix 版本号数字部分之后的后缀
	// 例如对于版本号 "1.2.3-beta1"，Suffix 为 "-beta1"
	Suffix VersionSuffix `json:"suffix"`
}

var _ compare_anything.Comparable[*Version] = &Version{}

// NewVersion 从版本字符串创建一个新的 Version 对象
//
// 该方法解析给定的版本字符串，并返回一个填充了相应字段的 Version 对象。
// 即使版本字符串格式不正确，该方法也会返回一个对象，但其 IsValid() 方法可能返回 false。
//
// 参数:
//   - versionStr: 要解析的版本号字符串，如 "1.2.3" 或 "v1.2.3-rc1"
//
// 返回:
//   - *Version: 解析后的 Version 对象
//
// 使用示例:
//
//	version := versions.NewVersion("v1.2.3-beta1")
func NewVersion(versionStr string) *Version {
	return NewVersionStringParser(versionStr).Parse()
}

// NewVersionE 从版本字符串创建一个新的 Version 对象，并返回可能的错误
//
// 与 NewVersion 不同，该方法会在版本字符串格式不正确时返回错误。
//
// 参数:
//   - versionStr: 要解析的版本号字符串
//
// 返回:
//   - *Version: 解析后的 Version 对象，如果解析失败则为 nil
//   - error: 如果版本号无效，则返回 ErrVersionInvalid 错误
//
// 使用示例:
//
//	version, err := versions.NewVersionE("v1.2.3-beta1")
//	if err != nil {
//	    log.Fatalf("无效的版本号: %v", err)
//	}
func NewVersionE(versionStr string) (*Version, error) {
	v := NewVersionStringParser(versionStr).Parse()
	if v.IsValid() {
		return v, nil
	} else {
		return nil, ErrVersionInvalid
	}
}

// NewVersions 批量创建多个 Version 对象
//
// 该方法接受多个版本字符串，并返回相应的 Version 对象数组。
//
// 参数:
//   - versionStringSlice: 一个或多个版本号字符串
//
// 返回:
//   - []*Version: 解析后的 Version 对象数组
//
// 使用示例:
//
//	versions := versions.NewVersions("1.0.0", "1.1.0", "2.0.0")
//	for _, v := range versions {
//	    fmt.Println(v.Raw)
//	}
func NewVersions(versionStringSlice ...string) []*Version {
	versions := make([]*Version, len(versionStringSlice))
	for i, versionStr := range versionStringSlice {
		versions[i] = NewVersion(versionStr)
	}
	return versions
}

// IsValid 检查版本号是否有效
//
// 判断依据是版本号中是否包含数字部分。只有当解析到了版本号数字时才认为是有效的版本号。
//
// 返回:
//   - bool: 如果版本号有效则返回 true，否则返回 false
//
// 使用示例:
//
//	version := versions.NewVersion("not-a-version")
//	if !version.IsValid() {
//	    fmt.Println("无效的版本号")
//	}
func (v *Version) IsValid() bool {
	// 版本号数组不为空，则表示为有效版本
	return len(v.VersionNumbers) > 0
}

// BuildGroupID 构造版本所属的组的ID
//
// 该方法根据版本号的数字部分生成一个组ID，用于将相似版本分组。
//
// 返回:
//   - string: 表示版本组的ID字符串
//
// 使用示例:
//
//	version := versions.NewVersion("1.2.3")
//	groupID := version.BuildGroupID()
//	fmt.Printf("版本组ID: %s\n", groupID)
func (x *Version) BuildGroupID() string {
	return x.VersionNumbers.BuildGroupID()
}

// CompareTo 比较两个版本号
//
// 该方法按以下顺序比较两个版本号：
// 1. 首先比较主版本号数字部分
// 2. 其次比较发布时间
// 3. 然后比较后缀
// 4. 最后比较原始版本号字符串
//
// 参数:
//   - target: 要比较的目标版本对象
//
// 返回:
//   - int: 如果当前版本小于目标版本，返回-1；如果相等，返回0；如果大于，返回1
//
// 使用示例:
//
//	v1 := versions.NewVersion("1.0.0")
//	v2 := versions.NewVersion("1.1.0")
//
//	switch v1.CompareTo(v2) {
//	case -1:
//	    fmt.Println("v1 < v2")
//	case 0:
//	    fmt.Println("v1 = v2")
//	case 1:
//	    fmt.Println("v1 > v2")
//	}
func (x *Version) CompareTo(target *Version) int {

	// 1. 先按照主版本号排序，仅当两个的主版本号都存在的时候才会进行比较，它们的长度不必相等，但是不能有为空的
	if len(x.VersionNumbers) != 0 && len(target.VersionNumbers) != 0 {
		r := x.VersionNumbers.CompareTo(target.VersionNumbers)
		if r != 0 {
			return r
		}
	}

	// 2. 然后按照后缀排序，修复：空后缀(release)应大于有后缀(pre-release)
	switch {
	case x.Suffix.IsEmpty() && target.Suffix.IsEmpty():
		// 两者都是正式版，后缀相等，继续比较
	case x.Suffix.IsEmpty() && !target.Suffix.IsEmpty():
		// 当前是正式版，目标是预发布版，当前更大
		return 1
	case !x.Suffix.IsEmpty() && target.Suffix.IsEmpty():
		// 当前是预发布版，目标是正式版，当前更小
		return -1
	default:
		// 两者都有后缀，比较后缀
		r := x.Suffix.CompareTo(target.Suffix)
		if r != 0 {
			return r
		}
	}

	// 3. 然后按照发布时间排序
	if !target.PublicTime.IsZero() && !x.PublicTime.IsZero() {
		r2 := x.PublicTime.UnixMilli() - target.PublicTime.UnixMilli()
		if r2 != 0 {
			// 不做类型转换是为了避免特殊情况下因为类型转换而丢失精度结果错误，而采用比较的方式
			if r2 > 0 {
				return 1
			} else {
				return -1
			}
		}
	}

	// 4. 最后实在不行就是比较原始版本号的字典序吧
	if x.Raw == target.Raw {
		return 0
	} else if x.Raw < target.Raw {
		return -1
	} else if x.Raw > target.Raw {
		return 1
	}

	// unreachable
	return 0
}

// String 返回版本的JSON字符串表示
//
// 该方法将Version对象序列化为JSON字符串，便于打印和调试。
//
// 返回:
//   - string: 版本的JSON字符串表示
//
// 使用示例:
//
//	version := versions.NewVersion("1.2.3")
//	fmt.Println(version.String()) // 输出JSON格式的版本信息
func (x *Version) String() string {
	marshal, _ := json.Marshal(x)
	return string(marshal)
}
