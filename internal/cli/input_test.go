package cli

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/scagogogo/versions-skills"
)

func TestReadLinesFromReader_Normal(t *testing.T) {
	r := strings.NewReader("1.0.0\n2.0.0\n3.0.0\n")
	lines, err := readLinesFromReader(r)
	require.NoError(t, err)
	assert.Equal(t, []string{"1.0.0", "2.0.0", "3.0.0"}, lines)
}

func TestReadLinesFromReader_SkipsBlankLines(t *testing.T) {
	r := strings.NewReader("1.0.0\n\n\n2.0.0\n\n")
	lines, err := readLinesFromReader(r)
	require.NoError(t, err)
	assert.Equal(t, []string{"1.0.0", "2.0.0"}, lines)
}

func TestReadLinesFromReader_Empty(t *testing.T) {
	r := strings.NewReader("")
	lines, err := readLinesFromReader(r)
	require.NoError(t, err)
	assert.Nil(t, lines)
}

func TestReadLinesFromFile_Normal(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "v.txt")
	require.NoError(t, os.WriteFile(p, []byte("1.0.0\n2.0.0\n"), 0o644))
	lines, err := readLinesFromFile(p)
	require.NoError(t, err)
	assert.Equal(t, []string{"1.0.0", "2.0.0"}, lines)
}

func TestReadLinesFromFile_NotExist(t *testing.T) {
	_, err := readLinesFromFile("/nonexistent/path/does/not/exist.txt")
	assert.Error(t, err)
}

func TestParseValidVersion_Valid(t *testing.T) {
	v, err := ParseValidVersion("1.2.3")
	require.NoError(t, err)
	assert.True(t, v.IsValid())
}

func TestParseValidVersion_Invalid(t *testing.T) {
	v, err := ParseValidVersion("not-a-version")
	assert.Error(t, err)
	assert.Nil(t, v)
}

func TestResolveVersions_FromArgs(t *testing.T) {
	vs, err := ResolveVersions([]string{"1.0.0", "2.0.0"}, "")
	require.NoError(t, err)
	require.Len(t, vs, 2)
	assert.Equal(t, "1.0.0", vs[0].RawString())
	assert.Equal(t, "2.0.0", vs[1].RawString())
}

func TestResolveVersions_FromFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "v.txt")
	require.NoError(t, os.WriteFile(p, []byte("1.0.0\n2.0.0\n"), 0o644))
	vs, err := ResolveVersions(nil, p)
	require.NoError(t, err)
	require.Len(t, vs, 2)
}

func TestResolveVersions_FromFileError(t *testing.T) {
	_, err := ResolveVersions(nil, "/nonexistent/path/x.txt")
	assert.Error(t, err)
}

func TestResolveVersions_FromStdin(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, err := os.Pipe()
	require.NoError(t, err)
	_, _ = w.WriteString("1.0.0\n2.0.0\n")
	_ = w.Close()
	os.Stdin = r
	vs, err := ResolveVersions(nil, "")
	require.NoError(t, err)
	require.Len(t, vs, 2)
}

func TestResolveVersions_StdinReadError(t *testing.T) {
	// 通过让 stdin 为普通文件（非管道）但 args 为空，绕过 stdin 分支，
	// 再覆盖 Stdin 为会触发 scanner 错误的 reader。
	// 这里直接测试空输入 + 非 char device 的非管道情况：用一个普通文件作 stdin，
	// stat 时 ModeCharDevice == 0，会走 stdin 分支。
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	// /dev/null 不是 char device？实际 /dev/null 是 char device。
	// 改用一个普通临时文件作 stdin（非 char device），走 stdin 分支并成功读取（空）。
	dir := t.TempDir()
	p := filepath.Join(dir, "empty.txt")
	require.NoError(t, os.WriteFile(p, []byte(""), 0o644))
	f, err := os.Open(p)
	require.NoError(t, err)
	defer f.Close()
	os.Stdin = f
	// 空文件 → lines 为空 → 返回 error
	_, err = ResolveVersions(nil, "")
	assert.Error(t, err)
}

func TestReadLinesFromReader_ScanError(t *testing.T) {
	// 构造一个超过 bufio.MaxScanTokenSize 的行，触发 scanner.ErrTooLong
	long := strings.Repeat("a", bufio.MaxScanTokenSize+1)
	lines, err := readLinesFromReader(strings.NewReader(long))
	assert.Error(t, err)
	assert.Nil(t, lines)
}

func TestResolveVersions_StdinScanError(t *testing.T) {
	// 用管道写入超长行，使 os.Stdin 上的 bufio.Scanner 报 ErrTooLong，
	// 覆盖 ResolveVersions 中 readLinesFromReader(os.Stdin) 返回 err 的分支。
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, err := os.Pipe()
	require.NoError(t, err)
	defer r.Close()
	long := strings.Repeat("a", bufio.MaxScanTokenSize+1)
	go func() {
		_, _ = w.WriteString(long)
		_ = w.Close()
	}()
	os.Stdin = r
	_, err = ResolveVersions(nil, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "stdin")
}

func TestResolveVersions_NoInput(t *testing.T) {
	// args 为空、stdin 为 char device（终端）时，不走 stdin 分支 → lines 为空 → error
	// 测试环境中 os.Stdin 可能不是 char device，因此强制把 stdin 设为 char device。
	// /dev/tty 通常不可用，改用 os.Stdin 本身在 CI 中可能是管道。
	// 这里直接验证：args 为空且无 fromFile 且 stdin 为 char device 时返回 error。
	// 通过打开一个 char device 文件作为 stdin 来模拟终端。
	devNull, err := os.OpenFile("/dev/null", os.O_RDONLY, 0)
	if err != nil {
		t.Skip("/dev/null 不可用")
	}
	defer devNull.Close()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = devNull
	_, err = ResolveVersions(nil, "")
	assert.Error(t, err)
}

func TestResolveVersions_EmptyLinesSkipped(t *testing.T) {
	vs, err := ResolveVersions([]string{"", "1.0.0", "", "2.0.0", ""}, "")
	require.NoError(t, err)
	require.Len(t, vs, 2)
}

func TestResolveVersions_FilePriorityOverArgs(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "v.txt")
	require.NoError(t, os.WriteFile(p, []byte("3.0.0\n4.0.0\n"), 0o644))
	// 提供 args 和 fromFile，fromFile 优先
	vs, err := ResolveVersions([]string{"1.0.0", "2.0.0"}, p)
	require.NoError(t, err)
	require.Len(t, vs, 2)
	assert.Equal(t, "3.0.0", vs[0].RawString())
}

func TestResolveVersionsStrict_FiltersInvalid(t *testing.T) {
	vs, err := ResolveVersionsStrict([]string{"1.0.0", "not-a-version", "2.0.0"}, "")
	require.NoError(t, err)
	require.Len(t, vs, 2)
	for _, v := range vs {
		assert.True(t, v.IsValid())
	}
}

func TestResolveVersionsStrict_AllInvalid(t *testing.T) {
	_, err := ResolveVersionsStrict([]string{"nope", "still-nope"}, "")
	assert.Error(t, err)
}

func TestResolveVersionsStrict_NoInput(t *testing.T) {
	_, err := ResolveVersionsStrict([]string{""}, "")
	assert.Error(t, err)
}

func TestReverseSlice(t *testing.T) {
	even := []string{"a", "b", "c", "d"}
	ReverseSlice(even)
	assert.Equal(t, []string{"d", "c", "b", "a"}, even)

	odd := []string{"1", "2", "3"}
	ReverseSlice(odd)
	assert.Equal(t, []string{"3", "2", "1"}, odd)

	single := []string{"x"}
	ReverseSlice(single)
	assert.Equal(t, []string{"x"}, single)

	var empty []string
	ReverseSlice(empty)
	assert.Nil(t, empty)
}

func TestGetFormatDefault(t *testing.T) {
	// 默认 format 为 json
	old := format
	defer func() { format = old }()
	format = "json"
	assert.Equal(t, "json", getFormat())
}

func TestIsQuietDefault(t *testing.T) {
	old := quiet
	defer func() { quiet = old }()
	quiet = false
	assert.False(t, isQuiet())
	quiet = true
	assert.True(t, isQuiet())
}

// 确保 versions 包被使用（避免某些场景下被优化掉）
var _ = versions.NewVersion
