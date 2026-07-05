package verfmt

import (
	"testing"

	"github.com/scagogogo/versions-skills"
	"github.com/stretchr/testify/assert"
)

func TestFormatVersion(t *testing.T) {
	v := versions.NewVersion("v1.2.3-beta.1")
	m := FormatVersion(v)
	assert.Equal(t, "v1.2.3-beta.1", m["raw"])
	assert.Equal(t, true, m["valid"])
	assert.Equal(t, "v", m["prefix"])
	assert.Equal(t, int(1), m["major"])
	assert.Equal(t, int(2), m["minor"])
	assert.Equal(t, int(3), m["patch"])
	assert.NotNil(t, m["version_numbers"])
	assert.NotNil(t, m["suffix"])
	assert.NotNil(t, m["suffix_weight"])
	assert.NotNil(t, m["group_id"])
}

func TestFormatVersion_Nil(t *testing.T) {
	assert.Nil(t, FormatVersion(nil))
	assert.Nil(t, FormatVersionDetailed(nil))
	assert.Nil(t, FormatVersionSimple(nil))
}

func TestFormatVersionSimple(t *testing.T) {
	v := versions.NewVersion("1.0.0")
	m := FormatVersionSimple(v)
	assert.Equal(t, "1.0.0", m["raw"])
	assert.Equal(t, true, m["valid"])
	assert.Equal(t, 2, len(m))
}

func TestFormatVersionDetailed(t *testing.T) {
	v := versions.NewVersion("1.0.0-beta")
	m := FormatVersionDetailed(v)
	assert.Equal(t, "1.0.0-beta", m["raw"])
	assert.Equal(t, true, m["is_prerelease"])
	assert.Equal(t, false, m["is_stable"])
	assert.Equal(t, true, m["is_beta"])
	assert.Equal(t, false, m["is_alpha"])
	// Is* 全字段都应存在
	for _, k := range []string{"sub_version", "metadata", "is_dev", "is_snapshot",
		"is_milestone", "is_nightly", "is_final", "is_ga", "is_pre", "is_release",
		"is_sp", "is_post", "is_zero", "is_rc", "core"} {
		_, ok := m[k]
		assert.True(t, ok, "field %s should exist", k)
	}
}

func TestFormatVersionStrings(t *testing.T) {
	vs := versions.NewVersions("1.0.0", "2.0.0", "3.0.0")
	got := FormatVersionStrings(vs)
	assert.Equal(t, []string{"1.0.0", "2.0.0", "3.0.0"}, got)
	assert.Equal(t, 0, len(FormatVersionStrings(nil)))
}

func TestFormatConstraint(t *testing.T) {
	c, err := versions.ParseConstraint(">=1.0.0")
	assert.NoError(t, err)
	m := FormatConstraint(c)
	assert.Equal(t, ">=", m["operator"])
	assert.Equal(t, "1.0.0", m["version"])
	assert.Equal(t, ">=1.0.0", m["string"])
}

func TestFormatConstraint_Nil(t *testing.T) {
	assert.Nil(t, FormatConstraint(nil))
}

func TestFormatConstraintSet(t *testing.T) {
	cs, err := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
	assert.NoError(t, err)
	m := FormatConstraintSet(cs)
	assert.Equal(t, 2, m["len"])
	assert.Equal(t, ">=1.0.0,<2.0.0", m["string"])
	constraints, ok := m["constraints"].([]map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, 2, len(constraints))
}

func TestFormatConstraintSet_Nil(t *testing.T) {
	assert.Nil(t, FormatConstraintSet(nil))
}

func TestFormatConstraintUnion(t *testing.T) {
	cu, err := versions.ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0,<4.0.0")
	assert.NoError(t, err)
	m := FormatConstraintUnion(cu)
	assert.Equal(t, ">=1.0.0,<2.0.0 || >=3.0.0,<4.0.0", m["string"])
	sets, ok := m["sets"].([]map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, 2, len(sets))
}

func TestFormatConstraintUnion_Nil(t *testing.T) {
	assert.Nil(t, FormatConstraintUnion(nil))
}

func TestFormatVersionGroup(t *testing.T) {
	g := versions.NewVersionGroupFromVersions(versions.NewVersions("1.0.0", "1.1.0-beta", "1.2.0"))
	m := FormatVersionGroup(g)
	assert.Equal(t, 3, m["count"])
	assert.Equal(t, "1.2.0", m["latest"])
	assert.Equal(t, "1.0.0", m["oldest"])
	assert.Equal(t, "1.2.0", m["latest_stable"])
	assert.Equal(t, "1.1.0-beta", m["latest_prerelease"])
	vs, ok := m["versions"].([]string)
	assert.True(t, ok)
	assert.Equal(t, 3, len(vs))
}

func TestFormatVersionGroup_Nil(t *testing.T) {
	assert.Nil(t, FormatVersionGroup(nil))
}

func TestFormatVersionGroupMap(t *testing.T) {
	// 构造一个分组 map（按全部数字段分组，相同数字段的版本同组）
	sg := versions.NewSortedVersionGroups(versions.NewVersions(
		"1.0.0", "1.0.0-beta", "2.0.0", "2.0.0-rc1",
	))
	gm := map[string]*versions.VersionGroup{}
	for _, id := range sg.GroupIDs() {
		gm[id] = sg.Get(id)
	}
	result := FormatVersionGroupMap(gm)
	// 按 ID 排序，应为 2 组
	assert.Equal(t, 2, len(result))
	// 第一组的 ID 应小于第二组
	first, ok := result[0]["id"].(string)
	assert.True(t, ok)
	second, _ := result[1]["id"].(string)
	assert.Less(t, first, second)
}

func TestFormatVersionGroupMap_Empty(t *testing.T) {
	result := FormatVersionGroupMap(map[string]*versions.VersionGroup{})
	assert.Equal(t, 0, len(result))
}

func TestFormatIntSlice(t *testing.T) {
	got := formatIntSlice([]int{1, 2, 3})
	assert.Equal(t, 3, len(got))
	assert.Equal(t, 1, got[0])
	assert.Equal(t, 0, len(formatIntSlice(nil)))
}
