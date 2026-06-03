package versions

import (
	"sort"

	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"github.com/golang-infrastructure/go-tuple"
)

// VersionGroup 表示一个版本组，一个版本组中可能有一个或多个版本
//
// VersionGroup 用于管理具有相同主版本号数字部分的一组版本。它提供了版本的添加、查询、排序和范围查询等功能。
// 版本组是版本管理系统中的重要概念，用于将相似版本聚合在一起，便于进行版本管理和分析。
//
// 每个版本组都有一个唯一的ID，由其主版本号数字部分生成，如 "1.2"。
//
// 使用示例:
//
//	// 创建一个新的版本组
//	group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
//
//	// 添加版本到组中
//	v1 := versions.NewVersion("1.2.0")
//	v2 := versions.NewVersion("1.2.1")
//	group.Add(v1)
//	group.Add(v2)
//
//	// 获取组中的所有版本
//	allVersions := group.Versions()
//
//	// 获取排序后的版本
//	sortedVersions := group.SortVersions()
type VersionGroup struct {

	// GroupVersionNumbers 组的版本号中的数字部分
	// 例如对于版本号 "1.2.x"，GroupVersionNumbers 为 [1,2]
	GroupVersionNumbers VersionNumbers

	// VersionMap 组中包含的所有版本，键为版本的原始字符串，值为版本对象
	VersionMap map[string]*Version
}

var _ compare_anything.Comparable[*VersionGroup] = &VersionGroup{}

// NewVersionGroup 创建一个新的版本组
//
// 该方法根据指定的版本号数字部分创建一个新的版本组。创建时需要传递能够在包范围下唯一区分组的组ID，
// 这个ID选择的是版本中的数字部分。
//
// 参数:
//   - groupVersionNumbers: 版本组的数字部分，如 [1,2] 表示 "1.2" 版本组
//
// 返回:
//   - *VersionGroup: 新创建的版本组对象
//
// 使用示例:
//
//	// 创建表示 "1.2" 版本组的对象
//	group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
func NewVersionGroup(groupVersionNumbers VersionNumbers) *VersionGroup {
	return &VersionGroup{
		GroupVersionNumbers: groupVersionNumbers,
		VersionMap:          make(map[string]*Version, 0),
	}
}

// NewVersionGroupFromVersions 从版本数组创建一个版本组
//
// 该方法基于给定的版本数组创建一个版本组。所有版本将被添加到同一个组中，
// 该组的数字部分取自第一个版本的数字部分。
//
// 参数:
//   - versions: 要添加到组中的版本数组
//
// 返回:
//   - *VersionGroup: 新创建的包含所有指定版本的版本组，如果输入为空则返回 nil
//
// 使用示例:
//
//	versions := versions.NewVersions("1.2.0", "1.2.1", "1.2.2")
//	group := versions.NewVersionGroupFromVersions(versions)
func NewVersionGroupFromVersions(versions []*Version) *VersionGroup {
	if len(versions) == 0 {
		return nil
	}
	group := NewVersionGroup(versions[0].VersionNumbers)
	for _, v := range versions {
		group.Add(v)
	}
	return group
}

// Add 把给定的版本添加到本版本组中
//
// 该方法将指定的版本添加到版本组中。如果版本已存在，则会被覆盖。
//
// 参数:
//   - v: 要添加的版本对象
//
// 返回:
//   - bool: 如果版本之前已存在于组中则返回 true，否则返回 false
//
// 使用示例:
//
//	group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
//	version := versions.NewVersion("1.2.3")
//	exists := group.Add(version)
//	if !exists {
//	    fmt.Println("添加了新版本")
//	}
func (x *VersionGroup) Add(v *Version) bool {
	_, exists := x.VersionMap[v.Raw]
	x.VersionMap[v.Raw] = v
	return exists
}

// Contains 判断本版本组中是否包含给定的版本
//
// 该方法检查指定的版本是否已存在于版本组中。
//
// 参数:
//   - v: 要检查的版本对象
//
// 返回:
//   - bool: 如果版本存在于组中则返回 true，否则返回 false
//
// 使用示例:
//
//	group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
//	version := versions.NewVersion("1.2.3")
//	if !group.Contains(version) {
//	    group.Add(version)
//	}
func (x *VersionGroup) Contains(v *Version) bool {
	_, exists := x.VersionMap[v.Raw]
	return exists
}

// ID 返回组的ID
//
// 该方法返回版本组的唯一标识符，由其数字部分生成。
//
// 返回:
//   - string: 版本组的ID，例如 "1.2"
//
// 使用示例:
//
//	group := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
//	groupID := group.ID() // 返回 "1.2"
func (x *VersionGroup) ID() string {
	return x.GroupVersionNumbers.BuildGroupID()
}

// CompareTo 比较两个版本组的大小
//
// 该方法通过比较版本组的数字部分来确定两个版本组的先后顺序。
//
// 参数:
//   - target: 要比较的目标版本组
//
// 返回:
//   - int: 如果当前版本组小于目标版本组，返回负数；如果相等，返回0；如果大于，返回正数
//
// 使用示例:
//
//	group1 := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 2}))
//	group2 := versions.NewVersionGroup(versions.NewVersionNumbers([]int{1, 3}))
//
//	if group1.CompareTo(group2) < 0 {
//	    fmt.Println("group1 比 group2 旧")
//	}
func (x *VersionGroup) CompareTo(target *VersionGroup) int {
	return x.GroupVersionNumbers.CompareTo(target.GroupVersionNumbers)
}

// Versions 返回组下的所有版本
//
// 该方法返回版本组中包含的所有版本的数组，结果不保证有序。
//
// 返回:
//   - []*Version: 版本组中所有版本的数组
//
// 使用示例:
//
//	group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.0", "1.2.1"))
//	allVersions := group.Versions()
//	fmt.Printf("组中包含 %d 个版本\n", len(allVersions))
func (x *VersionGroup) Versions() []*Version {
	slice := make([]*Version, 0)
	for _, v := range x.VersionMap {
		slice = append(slice, v)
	}
	return slice
}

// SortVersions 对组下的所有版本进行排序返回
//
// 该方法返回版本组中所有版本的有序数组，排序遵循版本号的自然排序规则。
//
// 返回:
//   - []*Version: 排序后的版本数组
//
// 使用示例:
//
//	group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.2", "1.2.0", "1.2.1"))
//	sortedVersions := group.SortVersions()
//	// 结果顺序: ["1.2.0", "1.2.1", "1.2.2"]
func (x *VersionGroup) SortVersions() []*Version {
	versions := x.Versions()
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].CompareTo(versions[j]) < 0
	})
	return versions
}

// GetLatest 获取版本组中最新的版本
//
// 返回按排序后最新的版本，如果组为空则返回 nil。
func (x *VersionGroup) GetLatest() *Version {
	sorted := x.SortVersions()
	if len(sorted) == 0 {
		return nil
	}
	return sorted[len(sorted)-1]
}

// GetOldest 获取版本组中最旧的版本
//
// 返回按排序后最旧的版本，如果组为空则返回 nil。
func (x *VersionGroup) GetOldest() *Version {
	sorted := x.SortVersions()
	if len(sorted) == 0 {
		return nil
	}
	return sorted[0]
}

// Count 返回版本组中的版本数量
func (x *VersionGroup) Count() int {
	return len(x.VersionMap)
}

// StableVersions 返回版本组中所有稳定版本
func (x *VersionGroup) StableVersions() []*Version {
	return Filter(x.Versions(), func(v *Version) bool {
		return v.IsStable()
	})
}

// PrereleaseVersions 返回版本组中所有预发布版本
func (x *VersionGroup) PrereleaseVersions() []*Version {
	return Filter(x.Versions(), func(v *Version) bool {
		return v.IsPrerelease()
	})
}

// LatestStable 获取版本组中最新稳定版本
func (x *VersionGroup) LatestStable() *Version {
	return LatestStable(x.Versions())
}

// Remove 从版本组中移除指定的版本
//
// 如果版本存在于组中，则删除并返回 true；否则返回 false。
//
// 参数:
//   - v: 要移除的版本对象
//
// 返回:
//   - bool: 如果版本存在并被移除则返回 true
func (x *VersionGroup) Remove(v *Version) bool {
	if _, exists := x.VersionMap[v.Raw]; exists {
		delete(x.VersionMap, v.Raw)
		return true
	}
	return false
}

// LatestPrerelease 获取版本组中最新预发布版本
//
// 返回按排序后最新的预发布版本，如果组中无预发布版本则返回 nil。
func (x *VersionGroup) LatestPrerelease() *Version {
	return LatestPrerelease(x.Versions())
}

// QueryRangeVersions 获取组内指定区间内的版本
//
// 该方法根据给定的起始和结束版本范围，返回组内符合条件的版本数组。
// 版本的包含性由 ContainsPolicy 参数控制。
//
// 参数:
//   - start: 包含起始版本和包含策略的元组
//   - end: 包含结束版本和包含策略的元组
//
// 返回:
//   - []*Version: 符合区间条件的版本数组
//
// 使用示例:
//
//	group := versions.NewVersionGroupFromVersions(versions.NewVersions("1.2.0", "1.2.1", "1.2.2", "1.2.3"))
//
//	// 查询 1.2.0（包含）到 1.2.2（包含）的版本
//	startTuple := tuple.NewTuple2(versions.NewVersion("1.2.0"), versions.ContainsPolicyYes)
//	endTuple := tuple.NewTuple2(versions.NewVersion("1.2.2"), versions.ContainsPolicyYes)
//
//	rangeVersions := group.QueryRangeVersions(startTuple, endTuple)
//	// 结果包含: ["1.2.0", "1.2.1", "1.2.2"]
func (x *VersionGroup) QueryRangeVersions(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version {

	// 因为这里认为同一个版本组中不会有特别多的版本，所以就不再做索引表直接跳了，如果后面有发现特殊情况再来做优化
	sortedVersions := x.SortVersions()
	versions := make([]*Version, 0)
	for _, v := range sortedVersions {

		// 获取完了则结束
		if v.CompareTo(end.V1) > 0 {
			break
		}

		// 开始区间是否符合条件
		switch start.V2 {
		case ContainsPolicyNone, ContainsPolicyYes:
			if v.CompareTo(start.V1) < 0 {
				continue
			}
		case ContainsPolicyNo:
			if v.CompareTo(start.V1) <= 0 {
				continue
			}
		}

		// 结束区间是否符合条件
		switch end.V2 {
		case ContainsPolicyNone, ContainsPolicyYes:
			if v.CompareTo(end.V1) > 0 {
				continue
			}
		case ContainsPolicyNo:
			if v.CompareTo(end.V1) >= 0 {
				continue
			}
		}

		versions = append(versions, v)
	}
	return versions
}
