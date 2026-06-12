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
