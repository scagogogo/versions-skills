package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// check 各类型 result=true 路径（在主进程中运行，不触发 os.Exit，覆盖率正常记录）

func TestCmd_Check_Prerelease(t *testing.T) {
	out, err := runCmd(t, "check", "--prerelease", "1.0.0-beta")
	assert.NoError(t, err)
	assert.Contains(t, out, "prerelease")
}

func TestCmd_Check_Stable_True(t *testing.T) {
	out, err := runCmd(t, "check", "--stable", "1.0.0")
	assert.NoError(t, err)
	assert.Contains(t, out, "stable")
}

func TestCmd_Check_Dev(t *testing.T) {
	out, err := runCmd(t, "check", "--dev", "1.0.0-dev1")
	assert.NoError(t, err)
	assert.Contains(t, out, "dev")
}

func TestCmd_Check_Alpha(t *testing.T) {
	out, err := runCmd(t, "check", "--alpha", "1.0.0-alpha1")
	assert.NoError(t, err)
	assert.Contains(t, out, "alpha")
}

func TestCmd_Check_Beta(t *testing.T) {
	out, err := runCmd(t, "check", "--beta", "1.0.0-beta1")
	assert.NoError(t, err)
	assert.Contains(t, out, "beta")
}

func TestCmd_Check_RC(t *testing.T) {
	out, err := runCmd(t, "check", "--rc", "1.0.0-rc1")
	assert.NoError(t, err)
	assert.Contains(t, out, "rc")
}

func TestCmd_Check_Snapshot(t *testing.T) {
	out, err := runCmd(t, "check", "--snapshot", "1.0.0-snapshot")
	assert.NoError(t, err)
	assert.Contains(t, out, "snapshot")
}

func TestCmd_Check_Milestone(t *testing.T) {
	out, err := runCmd(t, "check", "--milestone", "1.0.0-m1")
	assert.NoError(t, err)
	assert.Contains(t, out, "milestone")
}

func TestCmd_Check_Nightly(t *testing.T) {
	out, err := runCmd(t, "check", "--nightly", "1.0.0-nightly")
	assert.NoError(t, err)
	assert.Contains(t, out, "nightly")
}

func TestCmd_Check_Final(t *testing.T) {
	out, err := runCmd(t, "check", "--final", "1.0.0-final")
	assert.NoError(t, err)
	assert.Contains(t, out, "final")
}

func TestCmd_Check_GA(t *testing.T) {
	out, err := runCmd(t, "check", "--ga", "1.0.0-ga")
	assert.NoError(t, err)
	assert.Contains(t, out, "ga")
}

func TestCmd_Check_Pre(t *testing.T) {
	out, err := runCmd(t, "check", "--pre", "1.0.0-pre1")
	assert.NoError(t, err)
	assert.Contains(t, out, "pre")
}

func TestCmd_Check_Release(t *testing.T) {
	out, err := runCmd(t, "check", "--release", "1.0.0-release")
	assert.NoError(t, err)
	assert.Contains(t, out, "release")
}

func TestCmd_Check_SP(t *testing.T) {
	out, err := runCmd(t, "check", "--sp", "1.0.0-sp1")
	assert.NoError(t, err)
	assert.Contains(t, out, "sp")
}

func TestCmd_Check_Post(t *testing.T) {
	out, err := runCmd(t, "check", "--post", "1.0.0-post1")
	assert.NoError(t, err)
	assert.Contains(t, out, "post")
}

func TestCmd_Check_Zero(t *testing.T) {
	// IsZero 仅对未初始化 Version{} 为 true，任何解析后的版本都为 false → os.Exit(1)
	// 用子进程测以避免杀死主测试进程
	code, _, _ := runSubprocess(t, "check_zero_false")
	assert.Equal(t, 1, code)
}

// 比较检查 result=true
func TestCmd_Check_Newer(t *testing.T) {
	out, err := runCmd(t, "check", "--newer", "1.0.0", "2.0.0")
	assert.NoError(t, err)
	assert.Contains(t, out, "newer")
}

func TestCmd_Check_Older(t *testing.T) {
	out, err := runCmd(t, "check", "--older", "2.0.0", "1.0.0")
	assert.NoError(t, err)
	assert.Contains(t, out, "older")
}

func TestCmd_Check_Equal(t *testing.T) {
	out, err := runCmd(t, "check", "--equal", "1.0.0", "1.0.0")
	assert.NoError(t, err)
	assert.Contains(t, out, "equal")
}

func TestCmd_Check_Between(t *testing.T) {
	out, err := runCmd(t, "check", "--between-low", "1.0.0", "--between-high", "3.0.0", "2.0.0")
	assert.NoError(t, err)
	assert.Contains(t, out, "between")
}

// check 比较检查的无效目标版本 → error → os.Exit（子进程）
func TestCmd_Check_Newer_InvalidTarget_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_newer_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Check_Older_InvalidTarget_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_older_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Check_Equal_InvalidTarget_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_equal_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Check_Between_InvalidLow_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_between_low_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Check_Between_InvalidHigh_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_between_high_invalid")
	assert.Equal(t, 1, code)
}

func TestCmd_Check_InvalidVersion_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_invalid_version")
	assert.Equal(t, 1, code)
}

func TestCmd_Check_NoType_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_no_type")
	assert.Equal(t, 1, code)
}

// compare v2 invalid
func TestCmd_Compare_InvalidV2_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "compare_invalid_v2")
	assert.Equal(t, 1, code)
}

// 各命令 ResolveVersions/ResolveVersionsStrict 错误路径（无版本输入）
func TestCmd_Filter_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "filter_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_Count_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "count_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_Partition_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "partition_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_Group_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "group_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_Max_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "max_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_LatestStable_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "lateststable_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_LatestPrerelease_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "latestprerelease_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_Range_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "range_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_Visualize_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "visualize_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_SortStrings_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "sortstrings_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_GroupIDs_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "groupids_noinput")
	assert.Equal(t, 1, code)
}

func TestCmd_Write_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "write_noinput")
	assert.Equal(t, 1, code)
}

// minmax result nil 路径（理论上 ResolveVersionsStrict 过滤后不会 nil，但保险起见）
// LatestStable/LatestPrerelease 无匹配 → PrintResult err → os.Exit（子进程）
func TestCmd_LatestStable_NoneFound_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "lateststable_none")
	assert.Equal(t, 1, code)
}

func TestCmd_LatestPrerelease_NoneFound_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "latestprerelease_none")
	assert.Equal(t, 1, code)
}

// group-latest-stable 无稳定版本 / group-latest-prerelease 无预发布
func TestCmd_GroupLatestStable_None_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "grouplateststable_none")
	assert.Equal(t, 1, code)
}

func TestCmd_GroupLatestPrerelease_None_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "grouplatestprerelease_none")
	assert.Equal(t, 1, code)
}

// group-latest 分组为空（GetLatest 返回 nil）— 此路径不可达
func TestCmd_GroupLatest_EmptyGroup_Unreachable(t *testing.T) {
	t.Skip("GetLatest 对输入构造的分组必非 nil，nil 分支为不可达死代码")
}

// group-oldest 分组为空（GetOldest 返回 nil）— 此路径不可达，从输入构造的分组必非空
func TestCmd_GroupOldest_EmptyGroup_Unreachable(t *testing.T) {
	t.Skip("GetOldest 对输入构造的分组必非 nil，nil 分支为不可达死代码")
}

// group-contains ResolveVersions error
func TestCmd_GroupContains_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "groupcontains_noinput")
	assert.Equal(t, 1, code)
}
