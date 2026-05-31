package versions

// Min 从版本列表中找到最小的版本
//
// 如果列表为空则返回 nil。
//
// 参数:
//   - versions: 版本对象列表
//
// 返回:
//   - *Version: 最小的版本对象，列表为空时返回 nil
func Min(versions []*Version) *Version {
	if len(versions) == 0 {
		return nil
	}
	min := versions[0]
	for _, v := range versions[1:] {
		if v.CompareTo(min) < 0 {
			min = v
		}
	}
	return min
}

// Max 从版本列表中找到最大的版本
//
// 如果列表为空则返回 nil。
func Max(versions []*Version) *Version {
	if len(versions) == 0 {
		return nil
	}
	max := versions[0]
	for _, v := range versions[1:] {
		if v.CompareTo(max) > 0 {
			max = v
		}
	}
	return max
}

// LatestStable 从版本列表中找到最新的稳定版本
//
// 稳定版本是指不带后缀的版本。如果不存在稳定版本则返回 nil。
func LatestStable(versions []*Version) *Version {
	var latest *Version
	for _, v := range versions {
		if v.IsStable() {
			if latest == nil || v.CompareTo(latest) > 0 {
				latest = v
			}
		}
	}
	return latest
}

// LatestPrerelease 从版本列表中找到最新的预发布版本
//
// 如果不存在预发布版本则返回 nil。
func LatestPrerelease(versions []*Version) *Version {
	var latest *Version
	for _, v := range versions {
		if v.IsPrerelease() {
			if latest == nil || v.CompareTo(latest) > 0 {
				latest = v
			}
		}
	}
	return latest
}

// Filter 根据谓词函数过滤版本列表
//
// 返回所有满足谓词条件的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - predicate: 过滤谓词函数，返回 true 表示保留该版本
//
// 返回:
//   - []*Version: 满足条件的版本列表
func Filter(versions []*Version, predicate func(*Version) bool) []*Version {
	result := make([]*Version, 0)
	for _, v := range versions {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterByConstraint 根据约束条件过滤版本列表
//
// 返回所有满足约束条件的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - constraint: 版本约束条件
//
// 返回:
//   - []*Version: 满足约束的版本列表
func FilterByConstraint(versions []*Version, constraint *Constraint) []*Version {
	return Filter(versions, func(v *Version) bool {
		return constraint.Match(v)
	})
}

// FilterByConstraintSet 根据约束集合过滤版本列表
//
// 返回所有满足约束集合中所有条件的版本。
func FilterByConstraintSet(versions []*Version, cs *ConstraintSet) []*Version {
	return Filter(versions, func(v *Version) bool {
		return cs.Match(v)
	})
}

// Unique 去除版本列表中的重复版本
//
// 根据 Raw 字段去重，保留第一次出现的版本。
func Unique(versions []*Version) []*Version {
	seen := make(map[string]bool)
	result := make([]*Version, 0)
	for _, v := range versions {
		if !seen[v.Raw] {
			seen[v.Raw] = true
			result = append(result, v)
		}
	}
	return result
}

// FilterByMajor 过滤指定主版本号的版本
func FilterByMajor(versions []*Version, major int) []*Version {
	return Filter(versions, func(v *Version) bool {
		return v.Major() == major
	})
}

// Count 统计版本列表中满足谓词的版本数量
func Count(versions []*Version, predicate func(*Version) bool) int {
	n := 0
	for _, v := range versions {
		if predicate(v) {
			n++
		}
	}
	return n
}
