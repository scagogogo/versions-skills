package versions

// VersionSlice 是 []*Version 的有序集合类型，实现了 sort.Interface
//
// VersionSlice 提供了 Go 标准库 sort.Sort() 的直接支持，
// 允许使用 sort.Sort(slice) 而不是 sort.Slice() 配合闭包。
//
// 使用示例:
//
//	slice := versions.VersionSlice(versions.NewVersions("2.0.0", "1.0.0", "1.5.0"))
//	sort.Sort(slice)
//	// slice 现在按版本号排序
type VersionSlice []*Version

// Len 实现 sort.Interface
func (s VersionSlice) Len() int {
	return len(s)
}

// Less 实现 sort.Interface
func (s VersionSlice) Less(i, j int) bool {
	return s[i].CompareTo(s[j]) < 0
}

// Swap 实现 sort.Interface
func (s VersionSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Min 返回切片中最小的版本
//
// 如果切片为空则返回 nil。
func (s VersionSlice) Min() *Version {
	return Min(s)
}

// Max 返回切片中最大的版本
//
// 如果切片为空则返回 nil。
func (s VersionSlice) Max() *Version {
	return Max(s)
}

// Filter 根据谓词函数过滤版本切片
//
// 返回所有满足谓词条件的版本。
func (s VersionSlice) Filter(predicate func(*Version) bool) VersionSlice {
	return VersionSlice(Filter(s, predicate))
}

// Contains 判断切片中是否包含指定版本
//
// 根据 Raw 字段判断版本是否相同。
func (s VersionSlice) Contains(target *Version) bool {
	return ContainsVersion(s, target)
}

// IndexOf 查找版本在切片中的位置
//
// 如果未找到则返回 -1。
func (s VersionSlice) IndexOf(target *Version) int {
	return IndexOf(s, target)
}

// Unique 去除切片中的重复版本
func (s VersionSlice) Unique() VersionSlice {
	return VersionSlice(Unique(s))
}

// Sort 对切片进行原地排序
func (s VersionSlice) Sort() {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i].CompareTo(s[j]) > 0 {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

// Sorted 返回排序后的新切片
func (s VersionSlice) Sorted() VersionSlice {
	result := make(VersionSlice, len(s))
	copy(result, s)
	result.Sort()
	return result
}
