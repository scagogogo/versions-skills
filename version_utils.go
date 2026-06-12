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

// FilterByMinor 过滤指定次版本号的版本
//
// 返回所有次版本号等于指定值的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - minor: 目标次版本号
//
// 返回:
//   - []*Version: 满足条件的版本列表
func FilterByMinor(versions []*Version, minor int) []*Version {
	return Filter(versions, func(v *Version) bool {
		return v.Minor() == minor
	})
}

// FilterByPrefix 过滤指定前缀的版本
//
// 返回所有前缀等于指定值的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - prefix: 目标前缀字符串，如 "v"
//
// 返回:
//   - []*Version: 满足条件的版本列表
func FilterByPrefix(versions []*Version, prefix string) []*Version {
	return Filter(versions, func(v *Version) bool {
		return string(v.Prefix) == prefix
	})
}

// Difference 返回在 a 中但不在 b 中的版本（差集）
//
// 根据 Raw 字段判断版本是否相同。返回的版本保持 a 中的原始顺序。
//
// 参数:
//   - a: 版本对象列表
//   - b: 要排除的版本对象列表
//
// 返回:
//   - []*Version: 差集版本列表
func Difference(a, b []*Version) []*Version {
	bSet := make(map[string]bool, len(b))
	for _, v := range b {
		bSet[v.Raw] = true
	}
	return Filter(a, func(v *Version) bool {
		return !bSet[v.Raw]
	})
}

// Intersection 返回同时存在于 a 和 b 中的版本（交集）
//
// 根据 Raw 字段判断版本是否相同。返回的版本保持 a 中的原始顺序。
//
// 参数:
//   - a: 版本对象列表
//   - b: 版本对象列表
//
// 返回:
//   - []*Version: 交集版本列表
func Intersection(a, b []*Version) []*Version {
	bSet := make(map[string]bool, len(b))
	for _, v := range b {
		bSet[v.Raw] = true
	}
	return Filter(a, func(v *Version) bool {
		return bSet[v.Raw]
	})
}

// Union 返回 a 和 b 中所有唯一版本的并集
//
// 根据 Raw 字段去重，保持 a 中元素的原始顺序，b 中不重复的元素追加到末尾。
//
// 参数:
//   - a: 版本对象列表
//   - b: 版本对象列表
//
// 返回:
//   - []*Version: 并集版本列表
func Union(a, b []*Version) []*Version {
	result := Unique(a)
	aSet := make(map[string]bool, len(a))
	for _, v := range a {
		aSet[v.Raw] = true
	}
	for _, v := range b {
		if !aSet[v.Raw] {
			result = append(result, v)
		}
	}
	return result
}

// Partition 根据谓词将版本列表分为两组
//
// 返回两个切片：第一个包含满足谓词的版本，第二个包含不满足谓词的版本。
// 保持原始顺序。
//
// 参数:
//   - versions: 版本对象列表
//   - predicate: 分区谓词函数
//
// 返回:
//   - []*Version: 满足谓词的版本列表
//   - []*Version: 不满足谓词的版本列表
func Partition(versions []*Version, predicate func(*Version) bool) ([]*Version, []*Version) {
	matched := make([]*Version, 0)
	unmatched := make([]*Version, 0)
	for _, v := range versions {
		if predicate(v) {
			matched = append(matched, v)
		} else {
			unmatched = append(unmatched, v)
		}
	}
	return matched, unmatched
}

// FilterByPatch 过滤指定修订版本号的版本
//
// 返回所有修订版本号等于指定值的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - patch: 目标修订版本号
//
// 返回:
//   - []*Version: 满足条件的版本列表
func FilterByPatch(versions []*Version, patch int) []*Version {
	return Filter(versions, func(v *Version) bool {
		return v.Patch() == patch
	})
}

// FilterBySuffix 过滤指定后缀的版本
//
// 返回所有后缀字符串等于指定值的版本。
//
// 参数:
//   - versions: 版本对象列表
//   - suffix: 目标后缀字符串，如 "-beta"
//
// 返回:
//   - []*Version: 满足条件的版本列表
func FilterBySuffix(versions []*Version, suffix string) []*Version {
	return Filter(versions, func(v *Version) bool {
		return string(v.Suffix) == suffix
	})
}

// FilterByStable 返回所有稳定版本
//
// 稳定版本是指不带任何后缀的版本。等价于 Filter(versions, Version.IsStable)。
//
// 参数:
//   - versions: 版本对象列表
//
// 返回:
//   - []*Version: 所有稳定版本
func FilterByStable(versions []*Version) []*Version {
	return Filter(versions, func(v *Version) bool {
		return v.IsStable()
	})
}

// FilterByPrerelease 返回所有预发布版本
//
// 预发布版本是指带有后缀的版本。等价于 Filter(versions, Version.IsPrerelease)。
//
// 参数:
//   - versions: 版本对象列表
//
// 返回:
//   - []*Version: 所有预发布版本
func FilterByPrerelease(versions []*Version) []*Version {
	return Filter(versions, func(v *Version) bool {
		return v.IsPrerelease()
	})
}
