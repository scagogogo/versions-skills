package versions

import "time"

// Clone 创建版本的深拷贝
//
// 返回一个与原版本完全相同的新 Version 对象，修改拷贝不会影响原版本。
// 对于不可变的 Version 对象，Clone 主要用于与 With* 方法配合使用。
//
// 返回:
//   - *Version: 版本的深拷贝
//
// 使用示例:
//
//	v1 := versions.NewVersion("1.2.3")
//	v2 := v1.Clone()
//	v2.Raw = "modified"
//	fmt.Println(v1.Raw) // 仍然是 "1.2.3"
func (x *Version) Clone() *Version {
	if x == nil {
		return nil
	}
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	return &Version{
		Raw:            x.Raw,
		PublicTime:     x.PublicTime,
		VersionNumbers: numbers,
		Prefix:         x.Prefix,
		Suffix:         x.Suffix,
		Metadata:       x.Metadata,
	}
}

// WithPrefix 返回一个修改前缀的新版本对象
//
// 原版本对象不变，返回一个新对象，其前缀被替换为指定值。
//
// 参数:
//   - prefix: 新的前缀字符串
//
// 返回:
//   - *Version: 修改前缀后的新版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	newV := v.WithPrefix("v")
//	// newV.Raw == "v1.2.3"
func (x *Version) WithPrefix(prefix string) *Version {
	v := NewVersionBuilder().
		Prefix(prefix).
		Numbers(x.VersionNumbers).
		Suffix(string(x.Suffix)).
		Build()
	v.Metadata = x.Metadata
	return v
}

// WithSuffix 返回一个修改后缀的新版本对象
//
// 原版本对象不变，返回一个新对象，其后缀被替换为指定值。
//
// 参数:
//   - suffix: 新的后缀字符串，如 "-beta1"
//
// 返回:
//   - *Version: 修改后缀后的新版本对象
//
// 使用示例:
//
//	v := versions.NewVersion("1.2.3")
//	newV := v.WithSuffix("-rc1")
//	// newV.Raw == "1.2.3-rc1"
func (x *Version) WithSuffix(suffix string) *Version {
	v := NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(x.VersionNumbers).
		Suffix(suffix).
		Build()
	v.Metadata = x.Metadata
	return v
}

// WithMajor 返回一个修改主版本号的新版本对象
//
// 原版本对象不变，返回一个新对象，其主版本号被替换为指定值。
// 后缀和前缀保持不变。
//
// 参数:
//   - major: 新的主版本号
//
// 返回:
//   - *Version: 修改主版本号后的新版本对象
func (x *Version) WithMajor(major int) *Version {
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	if len(numbers) == 0 {
		numbers = []int{major}
	} else {
		numbers[0] = major
	}
	v := NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Suffix(string(x.Suffix)).
		Build()
	v.Metadata = x.Metadata
	return v
}

// WithMinor 返回一个修改次版本号的新版本对象
//
// 原版本对象不变，返回一个新对象，其次版本号被替换为指定值。
// 前缀和后缀保持不变。
//
// 参数:
//   - minor: 新的次版本号
//
// 返回:
//   - *Version: 修改次版本号后的新版本对象
func (x *Version) WithMinor(minor int) *Version {
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	for len(numbers) < 2 {
		numbers = append(numbers, 0)
	}
	numbers[1] = minor
	v := NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Suffix(string(x.Suffix)).
		Build()
	v.Metadata = x.Metadata
	return v
}

// WithPatch 返回一个修改修订版本号的新版本对象
//
// 原版本对象不变，返回一个新对象，其修订版本号被替换为指定值。
// 前缀和后缀保持不变。
//
// 参数:
//   - patch: 新的修订版本号
//
// 返回:
//   - *Version: 修改修订版本号后的新版本对象
func (x *Version) WithPatch(patch int) *Version {
	numbers := make([]int, len(x.VersionNumbers))
	copy(numbers, x.VersionNumbers)
	for len(numbers) < 3 {
		numbers = append(numbers, 0)
	}
	numbers[2] = patch
	v := NewVersionBuilder().
		Prefix(string(x.Prefix)).
		Numbers(numbers).
		Suffix(string(x.Suffix)).
		Build()
	v.Metadata = x.Metadata
	return v
}

// WithPublicTime 返回一个修改发布时间的新版本对象
//
// 原版本对象不变，返回一个新对象，其发布时间被替换为指定值。
//
// 参数:
//   - t: 新的发布时间
//
// 返回:
//   - *Version: 修改发布时间后的新版本对象
func (x *Version) WithPublicTime(t time.Time) *Version {
	cloned := x.Clone()
	cloned.PublicTime = t
	return cloned
}
