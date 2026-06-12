package versions

import "github.com/golang-infrastructure/go-tuple"

// SortedVersionGroups 表示已排序的版本组集合
//
// SortedVersionGroups 封装了已排序的版本组切片和索引映射，
// 用于高效地进行版本组查询和范围检索。它通过预先排序和建立索引，
// 优化了版本组的查找性能。
//
// 结构特点:
// 1. 保持版本组的有序性，便于范围查询
// 2. 维护组ID到数组索引的映射，支持快速定位
// 3. 支持基于版本范围的高效查询
//
// 使用示例:
//
//	// 创建已排序的版本组
//	allVersions := versions.NewVersions("1.0.0", "1.1.0", "2.0.0", "2.1.0")
//	sortedGroups := versions.NewSortedVersionGroups(allVersions)
//
//	// 执行范围查询
//	startVer := versions.NewVersion("1.0.0")
//	endVer := versions.NewVersion("2.0.0")
//	startTuple := tuple.NewTuple2(startVer, versions.ContainsPolicyYes)
//	endTuple := tuple.NewTuple2(endVer, versions.ContainsPolicyYes)
//
//	rangeResult := sortedGroups.QueryRange(startTuple, endTuple)
type SortedVersionGroups struct {

	// groupIdToIndexMap 用于根据组的ID快速定位到这个组在排好序的切片中的位置
	// 键为版本组ID，值为该组在groupSlice中的索引位置
	groupIdToIndexMap map[string]int

	// groupSlice 排好序的版本组切片
	// 按照版本组的大小顺序排列
	groupSlice []*VersionGroup
}

// NewSortedVersionGroups 为版本号创建有序的分组
//
// 该方法接收一个版本对象数组，将其分组并排序，返回一个包含已排序版本组的SortedVersionGroups对象。
// 处理流程:
// 1. 首先将版本按照其数字部分分组
// 2. 然后对所有分组进行排序
// 3. 最后构建组ID到索引的映射，用于快速查找
//
// 参数:
//   - versions: 需要分组和排序的版本对象数组
//
// 返回:
//   - *SortedVersionGroups: 包含已排序版本组的对象
//
// 使用示例:
//
//	// 创建版本对象数组
//	allVersions := versions.NewVersions("1.0.0", "1.1.0", "1.2.0", "2.0.0", "2.1.0")
//
//	// 创建已排序的版本组
//	sortedGroups := versions.NewSortedVersionGroups(allVersions)
func NewSortedVersionGroups(versions []*Version) *SortedVersionGroups {

	// 先对所有的版本进行分组
	groupMap := Group(versions)

	// 对所有的分组排序
	groupSlice := SortVersionGroupMap(groupMap)

	// 然后构造有序Group
	groups := &SortedVersionGroups{
		groupIdToIndexMap: make(map[string]int),
		groupSlice:        groupSlice,
	}
	for i, g := range groupSlice {
		groups.groupIdToIndexMap[g.ID()] = i
	}
	return groups
}

// GroupIDs 返回所有版本组的ID列表
//
// 该方法返回按顺序排列的版本组ID列表。
// 返回的ID列表与内部的版本组切片顺序一致，确保保持排序状态。
//
// 返回:
//   - []string: 含有所有版本组ID的字符串切片
//
// 使用示例:
//
//	sortedGroups := versions.NewSortedVersionGroups(allVersions)
//	groupIDs := sortedGroups.GroupIDs()
//	for _, id := range groupIDs {
//	    fmt.Printf("版本组: %s\n", id)
//	}
func (x *SortedVersionGroups) GroupIDs() []string {
	result := make([]string, len(x.groupSlice))
	for i, group := range x.groupSlice {
		result[i] = group.ID()
	}
	return result
}

// QueryRange 在有序版本组中查询指定范围内的版本
//
// 该方法根据给定的起始和结束版本范围，返回所有符合条件的版本对象数组。
// 方法利用预排序的版本组结构，快速定位并收集符合范围条件的版本。
//
// 参数:
//   - start: 包含起始版本和包含策略的元组
//   - end: 包含结束版本和包含策略的元组
//
// 返回:
//   - []*Version: 符合查询范围条件的版本对象数组
//
// 使用示例:
//
//	// 创建已排序的版本组
//	sortedGroups := versions.NewSortedVersionGroups(allVersions)
//
//	// 定义查询范围
//	startVer := versions.NewVersion("1.0.0")
//	endVer := versions.NewVersion("2.0.0")
//	startTuple := tuple.NewTuple2(startVer, versions.ContainsPolicyYes) // 包含1.0.0
//	endTuple := tuple.NewTuple2(endVer, versions.ContainsPolicyNo)      // 不包含2.0.0
//
//	// 执行范围查询
//	rangeResult := sortedGroups.QueryRange(startTuple, endTuple)
//	fmt.Printf("在范围内的版本数: %d\n", len(rangeResult))
func (x *SortedVersionGroups) QueryRange(start, end *tuple.Tuple2[*Version, ContainsPolicy]) []*Version {

	// 如果要查询的版本组都不存在的话则直接返回空即可
	var i int
	if start.V1.Raw != "0" {
		// 仅当开始的版本不为0的时候才进行跳，否则认为是从最开始遍历
		var exists bool
		i, exists = x.groupIdToIndexMap[start.V1.BuildGroupID()]
		if !exists {
			return nil
		}
	}

	versions := make([]*Version, 0)
	for i < len(x.groupSlice) {
		g := x.groupSlice[i]
		i++

		versions = append(versions, g.QueryRangeVersions(start, end)...)
	}
	return versions
}

// Len 返回版本组的数量
func (x *SortedVersionGroups) Len() int {
	return len(x.groupSlice)
}

// Get 根据组 ID 获取版本组
//
// 如果组 ID 不存在则返回 nil。
//
// 参数:
//   - groupID: 版本组 ID，如 "1.2"
//
// 返回:
//   - *VersionGroup: 对应的版本组，不存在则返回 nil
func (x *SortedVersionGroups) Get(groupID string) *VersionGroup {
	idx, exists := x.groupIdToIndexMap[groupID]
	if !exists {
		return nil
	}
	return x.groupSlice[idx]
}

// At 根据索引获取版本组
//
// 返回排序后指定位置的版本组。如果索引越界则返回 nil。
//
// 参数:
//   - index: 从 0 开始的索引位置
//
// 返回:
//   - *VersionGroup: 对应的版本组，越界则返回 nil
func (x *SortedVersionGroups) At(index int) *VersionGroup {
	if index < 0 || index >= len(x.groupSlice) {
		return nil
	}
	return x.groupSlice[index]
}

// Contains 检查是否包含指定组 ID 的版本组
//
// 参数:
//   - groupID: 版本组 ID
//
// 返回:
//   - bool: 如果存在则返回 true
func (x *SortedVersionGroups) Contains(groupID string) bool {
	_, exists := x.groupIdToIndexMap[groupID]
	return exists
}

// Versions 返回所有版本组中的所有版本
//
// 版本按组排序，组内按版本排序。
//
// 返回:
//   - []*Version: 所有版本的有序列表
func (x *SortedVersionGroups) Versions() []*Version {
	result := make([]*Version, 0)
	for _, g := range x.groupSlice {
		result = append(result, g.SortVersions()...)
	}
	return result
}
