package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// runCmd 执行 rootCmd 捕获 stdout，返回输出与 err。
// 每个 test 自动还原 rootCmd args 与命令级 flag。
func runCmd(t *testing.T, args ...string) (string, error) {
	t.Helper()
	resetCmdFlags()
	resetRootArgs(t)
	rootCmd.SetArgs(args)
	restore, get := captureStdout(t)
	err := Execute()
	restore()
	return get(), err
}

// ---------- parse ----------

func TestCmd_Parse_Default(t *testing.T) {
	out, err := runCmd(t, "parse", "1.2.3")
	require.NoError(t, err)
	assert.Contains(t, out, "1.2.3")
	assert.Contains(t, out, "raw")
}

func TestCmd_Parse_WithDelimiters(t *testing.T) {
	out, err := runCmd(t, "parse", "--delimiters", "_-", "curl-7_85_0")
	require.NoError(t, err)
	assert.Contains(t, out, "curl")
	assert.Contains(t, out, "delimiters")
}

func TestCmd_Parse_MissingArg(t *testing.T) {
	_, err := runCmd(t, "parse")
	assert.Error(t, err)
}

// ---------- compare ----------

func TestCmd_Compare(t *testing.T) {
	out, err := runCmd(t, "compare", "1.2.3", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "result")
	assert.Contains(t, out, "旧于")
}

func TestCmd_Compare_Equal(t *testing.T) {
	out, err := runCmd(t, "compare", "1.0.0", "1.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "等于")
}

func TestCmd_Compare_Newer(t *testing.T) {
	out, err := runCmd(t, "compare", "2.0.0", "1.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "新于")
}

func TestCmd_Compare_InvalidV1(t *testing.T) {
	// PrintResult err 路径会 os.Exit(1)，必须用子进程测
	code, _, _ := runSubprocess(t, "compare_invalid_v1")
	assert.Equal(t, 1, code)
}

// ---------- sort ----------

func TestCmd_Sort(t *testing.T) {
	out, err := runCmd(t, "sort", "2.0.0", "1.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_Sort_Desc(t *testing.T) {
	out, err := runCmd(t, "sort", "--desc", "1.0.0", "2.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_Sort_FromFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "v.txt")
	require.NoError(t, os.WriteFile(p, []byte("2.0.0\n1.0.0\n"), 0o644))
	out, err := runCmd(t, "sort", "--from-file", p)
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

// sort 无版本输入 → error → os.Exit。用子进程。
func TestCmd_Sort_NoInput_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "sort_noinput")
	assert.Equal(t, 1, code)
}

// ---------- group ----------

func TestCmd_Group(t *testing.T) {
	out, err := runCmd(t, "group", "1.0.0", "1.0.0-alpha", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_Group_WithID(t *testing.T) {
	out, err := runCmd(t, "group", "--id", "1.0.0", "1.0.0", "1.0.0-alpha", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

// group --id 不存在 → error → os.Exit。子进程。
func TestCmd_Group_IDNotFound_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "group_idnotfound")
	assert.Equal(t, 1, code)
}

// ---------- filter ----------

func TestCmd_Filter_Stable(t *testing.T) {
	out, err := runCmd(t, "filter", "--stable", "1.0.0", "1.0.0-beta", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_Filter_Prerelease(t *testing.T) {
	out, err := runCmd(t, "filter", "--prerelease", "1.0.0-alpha", "1.0.0", "1.0.0-beta")
	require.NoError(t, err)
	assert.Contains(t, out, "alpha")
}

func TestCmd_Filter_Major(t *testing.T) {
	out, err := runCmd(t, "filter", "--major", "1", "1.0.0", "2.0.0", "1.5.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "1.5.0")
}

func TestCmd_Filter_Minor(t *testing.T) {
	out, err := runCmd(t, "filter", "--minor", "0", "1.0.0", "1.5.0", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_Filter_Patch(t *testing.T) {
	out, err := runCmd(t, "filter", "--patch", "0", "1.0.0", "1.0.5", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "2.0.0")
}

func TestCmd_Filter_Prefix(t *testing.T) {
	out, err := runCmd(t, "filter", "--prefix", "v", "v1.0.0", "1.0.0", "v2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "v1.0.0")
}

func TestCmd_Filter_Suffix(t *testing.T) {
	out, err := runCmd(t, "filter", "--suffix=-beta", "1.0.0-beta", "1.0.0", "2.0.0-beta")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0-beta")
}

func TestCmd_Filter_ConstraintSingle(t *testing.T) {
	out, err := runCmd(t, "filter", "--constraint", ">=1.0.0", "--constraint-type", "single", "1.0.0", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
}

func TestCmd_Filter_ConstraintSet(t *testing.T) {
	out, err := runCmd(t, "filter", "--constraint", ">=1.0.0,<2.0.0", "1.0.0", "1.5.0", "2.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "1.5.0")
}

func TestCmd_Filter_ConstraintUnion(t *testing.T) {
	out, err := runCmd(t, "filter", "--constraint", ">=1.0.0 || >=3.0.0", "--constraint-type", "union", "1.0.0", "2.0.0", "3.0.0")
	require.NoError(t, err)
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "3.0.0")
}

// filter 各 error 路径用子进程
func TestCmd_Filter_InvalidMajor_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "filter_invalid_major")
	assert.Equal(t, 1, code)
}

func TestCmd_Filter_InvalidMinor_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "filter_invalid_minor")
	assert.Equal(t, 1, code)
}

func TestCmd_Filter_InvalidPatch_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "filter_invalid_patch")
	assert.Equal(t, 1, code)
}

func TestCmd_Filter_ConstraintSingleParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "filter_constraint_single_fail")
	assert.Equal(t, 1, code)
}

func TestCmd_Filter_ConstraintUnionParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "filter_constraint_union_fail")
	assert.Equal(t, 1, code)
}

func TestCmd_Filter_ConstraintSetParseFail_Subprocess(t *testing.T) {
	code, _, _ := runSubprocess(t, "filter_constraint_set_fail")
	assert.Equal(t, 1, code)
}
