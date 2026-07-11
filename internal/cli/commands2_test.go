package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------- constraint ----------

func TestCmd_Constraint_Single(t *testing.T) {
	out, err := runCmd(t, "constraint", "--type", "single", ">=1.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "satisfied")
	assert.Contains(t, out, "true")
}

func TestCmd_Constraint_Set(t *testing.T) {
	out, err := runCmd(t, "constraint", ">=1.0.0,<2.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "satisfied")
}

func TestCmd_Constraint_Union(t *testing.T) {
	out, err := runCmd(t, "constraint", "--type", "union", ">=1.0.0 || >=3.0.0", "3.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "satisfied")
}

func TestCmd_Constraint_InvalidVersion_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "constraint_invalid_version")
	assert.Equal(t, 1, code)
}

func TestCmd_Constraint_SingleParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "constraint_single_fail")
	assert.Equal(t, 1, code)
}

func TestCmd_Constraint_UnionParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "constraint_union_fail")
	assert.Equal(t, 1, code)
}

func TestCmd_Constraint_SetParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "constraint_set_fail")
	assert.Equal(t, 1, code)
}

// ---------- range ----------

func TestCmd_Range(t *testing.T) {
	out, err := runCmd(t, "range", "1.0.0", "3.0.0", "1.0.0", "1.5.0", "2.0.0", "3.0.0", "4.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.5.0")
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_Range_IncludeEnd(t *testing.T) {
	out, err := runCmd(t, "range", "--include-end", "1.0.0", "3.0.0", "1.0.0", "3.0.0", "4.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "3.0.0")
}

func TestCmd_Range_InvalidStart_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "range_invalid_start")
	assert.Equal(t, 1, code)
}

func TestCmd_Range_InvalidEnd_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "range_invalid_end")
	assert.Equal(t, 1, code)
}

// ---------- build ----------

func TestCmd_Build_MajorMinorPatch(t *testing.T) {
	out, err := runCmd(t, "build", "--major", "1", "--minor", "2", "--patch", "3")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3")
}

func TestCmd_Build_PrefixSuffix(t *testing.T) {
	out, err := runCmd(t, "build", "--prefix", "v", "--major", "1", "--minor", "0", "--suffix", "-alpha1")
	require.NoError(t, err)
	assert.Contains(t, out, "v1.0")
}

func TestCmd_Build_Numbers(t *testing.T) {
	out, err := runCmd(t, "build", "--numbers", "1,2,3,4", "--prefix", "v")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3.4")
}

func TestCmd_Build_NumbersInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "build_numbers_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Build_MajorInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "build_major_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Build_MinorInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "build_minor_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Build_PatchInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "build_patch_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Build_NoArgs(t *testing.T) {
	// build 用 NoArgs，无 flag 时构建空版本
	out, err := runCmd(t, "build")
	require.NoError(t, err)
	assert.Contains(t, out, "valid")
}

// ---------- bump ----------

func TestCmd_Bump_Major(t *testing.T) {
	out, err := runCmd(t, "bump", "1.2.3", "--major")
	require.NoError(t, err)
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_Bump_Minor(t *testing.T) {
	out, err := runCmd(t, "bump", "1.2.3", "--minor")
	require.NoError(t, err)
	assert.Contains(t, out, "1.3.0")
}

func TestCmd_Bump_Patch(t *testing.T) {
	out, err := runCmd(t, "bump", "1.2.3", "--patch")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.4")
}

func TestCmd_Bump_InvalidVersion_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "bump_invalid_version")
	assert.Equal(t, 1, code)
}

func TestCmd_Bump_NoType_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "bump_no_type")
	assert.Equal(t, 1, code)
}

func TestGetBumpType(t *testing.T) {
	oldM, oldMi, oldP := bumpMajor, bumpMinor, bumpPatch
	defer func() { bumpMajor, bumpMinor, bumpPatch = oldM, oldMi, oldP }()

	bumpMajor, bumpMinor, bumpPatch = false, false, false
	assert.Equal(t, "", getBumpType())

	bumpMajor = true
	assert.Equal(t, "major", getBumpType())

	bumpMajor, bumpMinor = false, true
	assert.Equal(t, "minor", getBumpType())

	bumpMinor, bumpPatch = false, true
	assert.Equal(t, "patch", getBumpType())
}

// ---------- core ----------

func TestCmd_Core(t *testing.T) {
	out, err := runCmd(t, "core", "v1.2.3-beta1")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3")
}

func TestCmd_Core_Invalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "core_invalid")
	assert.Equal(t, 1, code)
}

// ---------- count ----------

func TestCmd_Count_Stable(t *testing.T) {
	out, err := runCmd(t, "count", "--stable", "1.0.0", "1.0.0-beta", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "count")
}

func TestCmd_Count_Major(t *testing.T) {
	out, err := runCmd(t, "count", "--major", "1", "1.0.0", "2.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "count")
}

func TestCmd_Count_Minor(t *testing.T) {
	out, err := runCmd(t, "count", "--minor", "0", "1.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "count")
}

func TestCmd_Count_Patch(t *testing.T) {
	out, err := runCmd(t, "count", "--patch", "0", "1.0.0", "1.0.5")
	require.NoError(t, err)
	assert.Contains(t, out, "count")
}

func TestCmd_Count_Prerelease(t *testing.T) {
	out, err := runCmd(t, "count", "--prerelease", "1.0.0-alpha", "1.0.0", "2.0.0-beta")
	require.NoError(t, err)
	assert.Contains(t, out, "count")
}

func TestCmd_Count_MajorInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "count_major_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Count_MinorInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "count_minor_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Count_PatchInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "count_patch_invalid")
	assert.Equal(t, 1, code)
}

// ---------- fileio ----------

func TestCmd_Read(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "v.txt")
	require.NoError(t, os.WriteFile(p, []byte("1.0.0\n2.0.0\n"), 0o644))
	out, err := runCmd(t, "read", p)
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_Read_NotExist_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "read_notexist")
	assert.Equal(t, 1, code)
}

func TestCmd_Write(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "out.txt")
	out, err := runCmd(t, "write", p, "1.0.0", "2.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "filepath")
	// 验证文件已写入并排序
	content, err := os.ReadFile(p)
	require.NoError(t, err)
	assert.Contains(t, string(content), "1.0.0")
	assert.Contains(t, string(content), "2.0.0")
}

func TestCmd_ReadStrings(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "v.txt")
	require.NoError(t, os.WriteFile(p, []byte("1.0.0\nnot-a-version\n2.0.0\n"), 0o644))
	out, err := runCmd(t, "read-strings", p)
	require.NoError(t, err)
	assert.Contains(t, out, "not-a-version")
}

func TestCmd_ReadStrings_NotExist_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "readstrings_notexist")
	assert.Equal(t, 1, code)
}

// ---------- info ----------

func TestCmd_Info(t *testing.T) {
	out, err := runCmd(t, "info", "v1.2.3-beta1")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3")
}

// ---------- minmax ----------

func TestCmd_Min(t *testing.T) {
	out, err := runCmd(t, "min", "2.0.0", "1.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_Max(t *testing.T) {
	out, err := runCmd(t, "max", "2.0.0", "1.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_LatestStable(t *testing.T) {
	out, err := runCmd(t, "latest-stable", "1.0.0-alpha", "1.0.0", "2.0.0-beta", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_LatestPrerelease(t *testing.T) {
	out, err := runCmd(t, "latest-prerelease", "1.0.0-alpha", "1.0.0-beta", "1.0.0", "2.0.0-rc1")
	require.NoError(t, err)
	assert.Contains(t, out, "rc1")
}

func TestCmd_Min_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "min_noinput")
	assert.Equal(t, 1, code)
}

// ---------- partition ----------

func TestCmd_Partition_Stable(t *testing.T) {
	out, err := runCmd(t, "partition", "--stable", "1.0.0-alpha", "1.0.0", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "matched")
}

func TestCmd_Partition_Prerelease(t *testing.T) {
	out, err := runCmd(t, "partition", "--prerelease", "1.0.0-alpha", "1.0.0", "2.0.0-rc1")
	require.NoError(t, err)
	assert.Contains(t, out, "matched")
}

func TestCmd_Partition_NoCondition_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "partition_no_condition")
	assert.Equal(t, 1, code)
}

// ---------- property ----------

func TestCmd_Segments(t *testing.T) {
	out, err := runCmd(t, "segments", "1.2.3")
	require.NoError(t, err)
	assert.Contains(t, out, "segments")
}

func TestCmd_SubVersion(t *testing.T) {
	out, err := runCmd(t, "sub-version", "1.2.3-beta2")
	require.NoError(t, err)
	assert.Contains(t, out, "sub_version")
}

func TestCmd_SuffixWeight(t *testing.T) {
	out, err := runCmd(t, "suffix-weight", "1.2.3-beta1")
	require.NoError(t, err)
	assert.Contains(t, out, "suffix_weight")
}

func TestCmd_PurePrefix(t *testing.T) {
	out, err := runCmd(t, "pure-prefix", "curl-7.85.0")
	require.NoError(t, err)
	assert.Contains(t, out, "pure_prefix")
}

func TestCmd_GroupID(t *testing.T) {
	out, err := runCmd(t, "group-id", "v1.2.3-beta1")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3")
}

func TestCmd_Satisfies(t *testing.T) {
	out, err := runCmd(t, "satisfies", "1.5.0", ">=1.0.0,<2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "satisfied")
}

func TestCmd_Satisfies_ParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "satisfies_parse_fail")
	assert.Equal(t, 1, code)
}

func TestCmd_Clone(t *testing.T) {
	out, err := runCmd(t, "clone", "v1.2.3-beta1")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3")
}

// property 各 invalid version 路径
func TestCmd_Property_Invalid_Subprocess(t *testing.T) {
	for _, sub := range []string{
		"segments_invalid", "subversion_invalid", "suffixweight_invalid",
		"pureprefix_invalid", "groupid_invalid", "clone_invalid",
	} {
		code, _, _ := runSubprocess(t, sub)
		assert.Equal(t, 1, code, sub)
	}
}

// ---------- set ----------

func TestCmd_SetPrefix(t *testing.T) {
	out, err := runCmd(t, "set-prefix", "1.2.3", "v")
	require.NoError(t, err)
	assert.Contains(t, out, "v1.2.3")
}

func TestCmd_SetSuffix(t *testing.T) {
	out, err := runCmd(t, "set-suffix", "1.2.3", "--", "-beta1")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3-beta1")
}

func TestCmd_SetMajor(t *testing.T) {
	out, err := runCmd(t, "set-major", "1.2.3", "2")
	require.NoError(t, err)
	assert.Contains(t, out, "2.2.3")
}

func TestCmd_SetMinor(t *testing.T) {
	out, err := runCmd(t, "set-minor", "1.2.3", "5")
	require.NoError(t, err)
	assert.Contains(t, out, "1.5.3")
}

func TestCmd_SetPatch(t *testing.T) {
	out, err := runCmd(t, "set-patch", "1.2.3", "9")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.9")
}

func TestCmd_SetNumbers(t *testing.T) {
	out, err := runCmd(t, "set-numbers", "1.2.3", "4,5,6")
	require.NoError(t, err)
	assert.Contains(t, out, "4.5.6")
}

func TestCmd_Set_InvalidVersion_Subprocess(t *testing.T) {
	for _, sub := range []string{
		"setprefix_invalid", "setsuffix_invalid", "setmajor_invalid",
		"setminor_invalid", "setpatch_invalid", "setnumbers_invalid",
	} {
		code, _, _ := runSubprocess(t, sub)
		assert.Equal(t, 1, code, sub)
	}
}

func TestCmd_Set_NumbersParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "setnumbers_parse_fail")
	assert.Equal(t, 1, code)
}

func TestCmd_SetMajor_ValueInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "setmajor_value_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_SetMinor_ValueInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "setminor_value_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_SetPatch_ValueInvalid_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "setpatch_value_invalid")
	assert.Equal(t, 1, code)
}

// ---------- sort-strings ----------

func TestCmd_SortStrings(t *testing.T) {
	out, err := runCmd(t, "sort-strings", "2.0.0", "1.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_SortStrings_Desc(t *testing.T) {
	out, err := runCmd(t, "sort-strings", "--desc", "1.0.0", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "2.0.0")
}

// ---------- visualize ----------

func TestCmd_Visualize(t *testing.T) {
	out, err := runCmd(t, "visualize", "1.0.0", "1.0.0-alpha", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "text")
}

func TestCmd_Visualize_Groups(t *testing.T) {
	out, err := runCmd(t, "visualize", "--groups", "1.0.0", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "groups")
}

func TestCmd_Visualize_MaxItems(t *testing.T) {
	out, err := runCmd(t, "visualize", "--max-items", "2", "1.0.0", "1.0.1", "1.0.2", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "text")
}

// ---------- group_extra ----------

func TestResolveGroupExtra_Normal(t *testing.T) {
	oldID, oldFile := groupExtraGroupID, groupExtraFromFile
	defer func() { groupExtraGroupID, groupExtraFromFile = oldID, oldFile }()
	groupExtraGroupID = "1.0.0"
	groupExtraFromFile = ""
	svg, g, err := resolveGroupExtra([]string{"1.0.0", "1.0.0-alpha", "2.0.0"})
	require.NoError(t, err)
	require.NotNil(t, svg)
	require.NotNil(t, g)
}

func TestResolveGroupExtra_NoInput(t *testing.T) {
	oldID, oldFile := groupExtraGroupID, groupExtraFromFile
	defer func() { groupExtraGroupID, groupExtraFromFile = oldID, oldFile }()
	groupExtraGroupID = "1.0.0"
	groupExtraFromFile = ""
	// 让 ResolveVersions 报错：args 为空 + stdin 为 char device
	devNull, err := os.OpenFile("/dev/null", os.O_RDONLY, 0)
	require.NoError(t, err)
	defer func() { _ = devNull.Close() }()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = devNull
	_, _, err = resolveGroupExtra(nil)
	assert.Error(t, err)
}

func TestResolveGroupExtra_GroupNotFound(t *testing.T) {
	oldID, oldFile := groupExtraGroupID, groupExtraFromFile
	defer func() { groupExtraGroupID, groupExtraFromFile = oldID, oldFile }()
	groupExtraGroupID = "9.9.9"
	svg, g, err := resolveGroupExtra([]string{"1.0.0", "2.0.0"})
	require.Error(t, err)
	require.NotNil(t, svg)
	assert.Nil(t, g)
}

func TestCmd_GroupIDs(t *testing.T) {
	out, err := runCmd(t, "group-ids", "1.0.0", "1.0.0-alpha", "2.0.0", "2.0.0-rc1")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_GroupLatest(t *testing.T) {
	out, err := runCmd(t, "group-latest", "--group-id", "1.0.0", "1.0.0-alpha", "1.0.0", "1.0.0-beta", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_GroupOldest(t *testing.T) {
	out, err := runCmd(t, "group-oldest", "--group-id", "1.0.0", "1.0.0-alpha", "1.0.0", "1.0.0-beta")
	require.NoError(t, err)
	assert.Contains(t, out, "alpha")
}

func TestCmd_GroupStable(t *testing.T) {
	out, err := runCmd(t, "group-stable", "--group-id", "1.0.0", "1.0.0-alpha", "1.0.0", "1.0.0-beta")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_GroupPrerelease(t *testing.T) {
	out, err := runCmd(t, "group-prerelease", "--group-id", "1.0.0", "1.0.0-alpha", "1.0.0", "1.0.0-beta")
	require.NoError(t, err)
	assert.Contains(t, out, "alpha")
}

func TestCmd_GroupLatestStable(t *testing.T) {
	out, err := runCmd(t, "group-latest-stable", "--group-id", "1.0.0", "1.0.0-alpha", "1.0.0", "1.0.0-beta", "1.0.1")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0")
}

func TestCmd_GroupLatestPrerelease(t *testing.T) {
	out, err := runCmd(t, "group-latest-prerelease", "--group-id", "1.0.0", "1.0.0-alpha", "1.0.0", "1.0.0-beta")
	require.NoError(t, err)
	assert.Contains(t, out, "beta")
}

func TestCmd_GroupContains(t *testing.T) {
	out, err := runCmd(t, "group-contains", "--group-id", "1.0.0", "--version", "1.0.0-alpha", "1.0.0-alpha", "1.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "contains")
}

func TestCmd_GroupContains_GroupNotFound(t *testing.T) {
	out, err := runCmd(t, "group-contains", "--group-id", "9.9.9", "--version", "1.0.0-alpha", "1.0.0-alpha")
	require.NoError(t, err)
	assert.Contains(t, out, "false")
}

// group-extra PreRunE error: 缺 --group-id
func TestCmd_GroupLatest_NoGroupID_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "grouplatest_no_id")
	assert.NotEqual(t, 0, code)
}

// group-latest 分组不存在 → error → os.Exit
func TestCmd_GroupLatest_GroupNotFound_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "grouplatest_notfound")
	assert.Equal(t, 1, code)
}

// group-contains 缺 --version
func TestCmd_GroupContains_NoVersion_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "groupcontains_no_version")
	assert.NotEqual(t, 0, code)
}
