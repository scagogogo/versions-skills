package cli

import (
	"flag"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 子进程测试模式：主测试用 exec.Command 调起 helper test，helper 触发 os.Exit，
// 主测试检查退出码与输出。
//
// 使用 -run 配合环境变量 CLI_SUBTEST 选择 helper。

var subtestFlag = flag.String("cli-subtest", "", "internal: run a subtest that triggers os.Exit")

// runSubprocess 在子进程中执行指定的 helper test，返回 (exitCode, stdout, stderr)。
// 子进程通过 -test.gocoverdir 写入覆盖率数据到共享目录，便于主测试合并。
//
// 注意：为了让 os.Exit 路径也能收集覆盖率，子进程必须使用 `go test -c -cover`
// 构建的二进制（环境变量 CLI_TEST_BINARY 指定）。否则 os.Exit 会跳过覆盖率刷新。
func runSubprocess(t *testing.T, subtest string) (int, string, string) {
	t.Helper()
	bin := os.Getenv("CLI_TEST_BINARY")
	if bin == "" {
		bin = os.Args[0]
	}
	args := []string{"-test.run=TestSubprocessMain", "-cli-subtest=" + subtest}
	// 若设置了覆盖率目录，让子进程把覆盖率中间数据写入其中
	if gocoverdir := os.Getenv("CLI_TEST_GOCOVERDIR"); gocoverdir != "" {
		args = append(args, "-test.gocoverdir="+gocoverdir)
	}
	cmd := exec.Command(bin, args...)
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()
	err := cmd.Run()
	code := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		code = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("failed to run subprocess: %v", err)
	}
	return code, stdout.String(), stderr.String()
}

// runCmdSubprocess 在子进程中执行 rootCmd 指定 args，若 Execute 返回 error 则 os.Exit(1)。
// 用于 PreRunE / Args 校验等不触发 os.Exit 的错误路径。
func runCmdSubprocess(args ...string) {
	rootCmd.SetArgs(args)
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}

// TestSubprocessMain 是子进程入口，根据 -cli-subtest 派发到具体 helper。
func TestSubprocessMain(t *testing.T) {
	switch *subtestFlag {
	case "":
		t.Skip("no subtest selected")
	case "printresult_error":
		// PrintResult err != nil 非 quiet → os.Exit(1)
		withQuiet(t, false)
		PrintResult("test", nil, simpleError("boom"))
	case "printresult_quiet_error":
		// quiet + err → 写 stderr 后 os.Exit(1)
		withQuiet(t, true)
		PrintResult("test", nil, simpleError("boom"))
	case "printjson_encode_fail":
		// printJSON 编码失败：传一个含 channel 的 data（无法 JSON 编码）
		printJSON(Result{Command: "test", Success: true, Data: make(chan int)})
	case "printresult_quiet_encode_fail":
		// quiet + data 编码失败
		withQuiet(t, true)
		PrintResult("test", make(chan int), nil)
	case "check_false":
		// check 命令 result=false → os.Exit(1)
		rootCmd.SetArgs([]string{"check", "--stable", "1.0.0-beta"})
		_ = Execute()
	case "check_true":
		// check 命令 result=true → 不退出（exit 0）
		rootCmd.SetArgs([]string{"check", "--stable", "1.0.0"})
		_ = Execute()
	case "validate_invalid":
		// validate 无效版本 → os.Exit(1)
		rootCmd.SetArgs([]string{"validate", "not-a-version"})
		_ = Execute()
	case "validate_valid":
		// validate 有效版本 → exit 0
		rootCmd.SetArgs([]string{"validate", "1.2.3"})
		_ = Execute()
	// --- 命令 error 路径（Run 内 PrintResult err → os.Exit） ---
	case "sort_noinput":
		rootCmd.SetArgs([]string{"sort"})
		_ = Execute()
	case "group_idnotfound":
		rootCmd.SetArgs([]string{"group", "--id", "9.9.9", "1.0.0"})
		_ = Execute()
	case "filter_invalid_major":
		rootCmd.SetArgs([]string{"filter", "--major", "abc", "1.0.0"})
		_ = Execute()
	case "filter_invalid_minor":
		rootCmd.SetArgs([]string{"filter", "--minor", "abc", "1.0.0"})
		_ = Execute()
	case "filter_invalid_patch":
		rootCmd.SetArgs([]string{"filter", "--patch", "abc", "1.0.0"})
		_ = Execute()
	case "filter_constraint_single_fail":
		rootCmd.SetArgs([]string{"filter", "--constraint", "notvalid", "--constraint-type", "single", "1.0.0"})
		_ = Execute()
	case "filter_constraint_union_fail":
		rootCmd.SetArgs([]string{"filter", "--constraint", "notvalid", "--constraint-type", "union", "1.0.0"})
		_ = Execute()
	case "filter_constraint_set_fail":
		rootCmd.SetArgs([]string{"filter", "--constraint", "notvalid", "--constraint-type", "set", "1.0.0"})
		_ = Execute()
	// --- compare (Run 内 PrintResult err → os.Exit) ---
	case "compare_invalid_v1":
		rootCmd.SetArgs([]string{"compare", "not-a-version", "1.0.0"})
		_ = Execute()
	case "compare_invalid_v2":
		rootCmd.SetArgs([]string{"compare", "1.0.0", "not-a-version"})
		_ = Execute()
	// --- constraint (Run 内 PrintResult err → os.Exit) ---
	case "constraint_invalid_version":
		rootCmd.SetArgs([]string{"constraint", ">=1.0.0", "not-a-version"})
		_ = Execute()
	case "constraint_single_fail":
		rootCmd.SetArgs([]string{"constraint", "--type", "single", "notvalid", "1.0.0"})
		_ = Execute()
	case "constraint_union_fail":
		rootCmd.SetArgs([]string{"constraint", "--type", "union", "notvalid", "1.0.0"})
		_ = Execute()
	case "constraint_set_fail":
		rootCmd.SetArgs([]string{"constraint", "--type", "set", "notvalid", "1.0.0"})
		_ = Execute()
	// --- range (Run 内 PrintResult err → os.Exit) ---
	case "range_invalid_start":
		rootCmd.SetArgs([]string{"range", "not-a-version", "3.0.0", "1.0.0", "2.0.0"})
		_ = Execute()
	case "range_invalid_end":
		rootCmd.SetArgs([]string{"range", "1.0.0", "not-a-version", "1.0.0", "2.0.0"})
		_ = Execute()
	// --- build (Run 内 PrintResult err → os.Exit) ---
	case "build_numbers_invalid":
		rootCmd.SetArgs([]string{"build", "--numbers", "1,abc,3"})
		_ = Execute()
	case "build_major_invalid":
		rootCmd.SetArgs([]string{"build", "--major", "abc"})
		_ = Execute()
	case "build_minor_invalid":
		rootCmd.SetArgs([]string{"build", "--minor", "abc"})
		_ = Execute()
	case "build_patch_invalid":
		rootCmd.SetArgs([]string{"build", "--patch", "abc"})
		_ = Execute()
	// --- bump (Run 内 PrintResult err → os.Exit) ---
	case "bump_invalid_version":
		rootCmd.SetArgs([]string{"bump", "not-a-version"})
		_ = Execute()
	case "bump_no_type":
		rootCmd.SetArgs([]string{"bump", "1.2.3"})
		_ = Execute()
	// --- core ---
	case "core_invalid":
		rootCmd.SetArgs([]string{"core", "not-a-version"})
		_ = Execute()
	// --- count ---
	case "count_major_invalid":
		rootCmd.SetArgs([]string{"count", "--major", "abc", "1.0.0"})
		_ = Execute()
	case "count_minor_invalid":
		rootCmd.SetArgs([]string{"count", "--minor", "abc", "1.0.0"})
		_ = Execute()
	case "count_patch_invalid":
		rootCmd.SetArgs([]string{"count", "--patch", "abc", "1.0.0"})
		_ = Execute()
	// --- fileio ---
	case "read_notexist":
		rootCmd.SetArgs([]string{"read", "/nonexistent/cli/x.txt"})
		_ = Execute()
	case "readstrings_notexist":
		rootCmd.SetArgs([]string{"read-strings", "/nonexistent/cli/x.txt"})
		_ = Execute()
	// --- minmax ---
	case "min_noinput":
		rootCmd.SetArgs([]string{"min"})
		_ = Execute()
	// --- partition ---
	case "partition_no_condition":
		rootCmd.SetArgs([]string{"partition", "1.0.0", "2.0.0"})
		_ = Execute()
	// --- property ---
	case "segments_invalid":
		rootCmd.SetArgs([]string{"segments", "not-a-version"})
		_ = Execute()
	case "subversion_invalid":
		rootCmd.SetArgs([]string{"sub-version", "not-a-version"})
		_ = Execute()
	case "suffixweight_invalid":
		rootCmd.SetArgs([]string{"suffix-weight", "not-a-version"})
		_ = Execute()
	case "pureprefix_invalid":
		rootCmd.SetArgs([]string{"pure-prefix", "not-a-version"})
		_ = Execute()
	case "groupid_invalid":
		rootCmd.SetArgs([]string{"group-id", "not-a-version"})
		_ = Execute()
	case "clone_invalid":
		rootCmd.SetArgs([]string{"clone", "not-a-version"})
		_ = Execute()
	case "satisfies_parse_fail":
		rootCmd.SetArgs([]string{"satisfies", "1.0.0", "notvalid"})
		_ = Execute()
	// --- set (Run 内 PrintResult err → os.Exit) ---
	case "setprefix_invalid":
		rootCmd.SetArgs([]string{"set-prefix", "not-a-version", "v"})
		_ = Execute()
	case "setsuffix_invalid":
		rootCmd.SetArgs([]string{"set-suffix", "not-a-version", "--", "-beta"})
		_ = Execute()
	case "setmajor_invalid":
		rootCmd.SetArgs([]string{"set-major", "not-a-version", "1"})
		_ = Execute()
	case "setminor_invalid":
		rootCmd.SetArgs([]string{"set-minor", "not-a-version", "1"})
		_ = Execute()
	case "setpatch_invalid":
		rootCmd.SetArgs([]string{"set-patch", "not-a-version", "1"})
		_ = Execute()
	case "setnumbers_invalid":
		rootCmd.SetArgs([]string{"set-numbers", "not-a-version", "1,2"})
		_ = Execute()
	case "setnumbers_parse_fail":
		rootCmd.SetArgs([]string{"set-numbers", "1.0.0", "1,abc"})
		_ = Execute()
	case "setmajor_value_invalid":
		rootCmd.SetArgs([]string{"set-major", "1.0.0", "abc"})
		_ = Execute()
	case "setminor_value_invalid":
		rootCmd.SetArgs([]string{"set-minor", "1.0.0", "abc"})
		_ = Execute()
	case "setpatch_value_invalid":
		rootCmd.SetArgs([]string{"set-patch", "1.0.0", "abc"})
		_ = Execute()
	// --- group_extra: PreRunE / 分组不存在 ---
	case "grouplatest_no_id":
		// PreRunE 返回 error → Execute 返回 error（不 os.Exit），子进程需手动退出
		runCmdSubprocess("group-latest", "1.0.0")
	case "grouplatest_notfound":
		// Run 内 PrintResult err → os.Exit
		rootCmd.SetArgs([]string{"group-latest", "--group-id", "9.9.9", "1.0.0"})
		_ = Execute()
	case "groupcontains_no_version":
		// PreRunE 返回 error
		runCmdSubprocess("group-contains", "--group-id", "1.0.0", "1.0.0")
	// --- check: 无效目标版本 / 无类型 / zero false ---
	case "check_newer_invalid":
		rootCmd.SetArgs([]string{"check", "--newer", "not-a-version", "1.0.0"})
		_ = Execute()
	case "check_older_invalid":
		rootCmd.SetArgs([]string{"check", "--older", "not-a-version", "1.0.0"})
		_ = Execute()
	case "check_equal_invalid":
		rootCmd.SetArgs([]string{"check", "--equal", "not-a-version", "1.0.0"})
		_ = Execute()
	case "check_between_low_invalid":
		rootCmd.SetArgs([]string{"check", "--between-low", "not-a-version", "--between-high", "3.0.0", "2.0.0"})
		_ = Execute()
	case "check_between_high_invalid":
		rootCmd.SetArgs([]string{"check", "--between-low", "1.0.0", "--between-high", "not-a-version", "2.0.0"})
		_ = Execute()
	case "check_invalid_version":
		rootCmd.SetArgs([]string{"check", "--stable", "not-a-version"})
		_ = Execute()
	case "check_no_type":
		rootCmd.SetArgs([]string{"check", "1.0.0"})
		_ = Execute()
	case "check_zero_false":
		rootCmd.SetArgs([]string{"check", "--zero", "0.0.0"})
		_ = Execute()
	// --- 各命令无版本输入（ResolveVersions error） ---
	case "filter_noinput":
		rootCmd.SetArgs([]string{"filter", "--stable"})
		_ = Execute()
	case "count_noinput":
		rootCmd.SetArgs([]string{"count", "--stable"})
		_ = Execute()
	case "partition_noinput":
		rootCmd.SetArgs([]string{"partition", "--stable"})
		_ = Execute()
	case "group_noinput":
		rootCmd.SetArgs([]string{"group"})
		_ = Execute()
	case "max_noinput":
		rootCmd.SetArgs([]string{"max"})
		_ = Execute()
	case "lateststable_noinput":
		rootCmd.SetArgs([]string{"latest-stable"})
		_ = Execute()
	case "latestprerelease_noinput":
		rootCmd.SetArgs([]string{"latest-prerelease"})
		_ = Execute()
	case "range_noinput":
		// range 需 2 个版本参数（start/end），第三个起为列表；无列表 → ResolveVersions error
		rootCmd.SetArgs([]string{"range", "1.0.0", "3.0.0"})
		_ = Execute()
	case "visualize_noinput":
		rootCmd.SetArgs([]string{"visualize"})
		_ = Execute()
	case "sortstrings_noinput":
		rootCmd.SetArgs([]string{"sort-strings"})
		_ = Execute()
	case "groupids_noinput":
		rootCmd.SetArgs([]string{"group-ids"})
		_ = Execute()
	case "write_noinput":
		// write 需要 filepath，第二个参数起为版本列表；无列表 → error
		rootCmd.SetArgs([]string{"write", "/tmp/cli_write_noinput.txt"})
		_ = Execute()
	// --- minmax result nil ---
	case "lateststable_none":
		// 全是预发布 → LatestStable 返回 nil
		rootCmd.SetArgs([]string{"latest-stable", "1.0.0-alpha", "2.0.0-beta"})
		_ = Execute()
	case "latestprerelease_none":
		// 全是稳定版 → LatestPrerelease 返回 nil
		rootCmd.SetArgs([]string{"latest-prerelease", "1.0.0", "2.0.0"})
		_ = Execute()
	// --- group_extra nil 路径 ---
	case "grouplateststable_none":
		// 分组只有预发布 → LatestStable nil
		rootCmd.SetArgs([]string{"group-latest-stable", "--group-id", "1.0.0", "1.0.0-alpha"})
		_ = Execute()
	case "grouplatestprerelease_none":
		// 分组只有稳定版 → LatestPrerelease nil
		rootCmd.SetArgs([]string{"group-latest-prerelease", "--group-id", "1.0.0", "1.0.0"})
		_ = Execute()
	case "groupcontains_noinput":
		// group-contains 缺版本输入
		rootCmd.SetArgs([]string{"group-contains", "--group-id", "1.0.0", "--version", "1.0.0-alpha"})
		_ = Execute()
	default:
		t.Fatalf("unknown subtest: %s", *subtestFlag)
	}
}

type simpleError string

func (e simpleError) Error() string { return string(e) }

// 主测试：验证各 os.Exit 路径的退出码

func TestSubprocess_PrintResult_Error_Exits1(t *testing.T) {
	code, stdout, _ := runSubprocess(t, "printresult_error")
	assert.Equal(t, 1, code)
	assert.Contains(t, stdout, "boom")
}

func TestSubprocess_PrintResult_QuietError_Exits1(t *testing.T) {
	code, _, stderr := runSubprocess(t, "printresult_quiet_error")
	assert.Equal(t, 1, code)
	assert.Contains(t, stderr, "boom")
}

func TestSubprocess_PrintJSON_EncodeFail_Exits1(t *testing.T) {
	code, _, stderr := runSubprocess(t, "printjson_encode_fail")
	assert.Equal(t, 1, code)
	assert.Contains(t, stderr, "JSON 编码失败")
}

func TestSubprocess_PrintResult_QuietEncodeFail_Exits1(t *testing.T) {
	code, _, stderr := runSubprocess(t, "printresult_quiet_encode_fail")
	assert.Equal(t, 1, code)
	assert.Contains(t, stderr, "JSON 编码失败")
}

func TestSubprocess_Check_False_Exits1(t *testing.T) {
	code, stdout, _ := runSubprocess(t, "check_false")
	assert.Equal(t, 1, code)
	assert.Contains(t, stdout, "stable")
}

func TestSubprocess_Check_True_Exits0(t *testing.T) {
	code, _, _ := runSubprocess(t, "check_true")
	assert.Equal(t, 0, code)
}

func TestSubprocess_Validate_Invalid_Exits1(t *testing.T) {
	code, stdout, _ := runSubprocess(t, "validate_invalid")
	assert.Equal(t, 1, code)
	assert.Contains(t, stdout, "valid")
}

func TestSubprocess_Validate_Valid_Exits0(t *testing.T) {
	code, _, _ := runSubprocess(t, "validate_valid")
	assert.Equal(t, 0, code)
}

// 确保 os/exec 被使用
var _ = exec.Command
