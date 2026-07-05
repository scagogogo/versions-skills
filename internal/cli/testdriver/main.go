// Package main 是 internal/cli 的覆盖率驱动程序。
//
// 仅用于在 os.Exit 路径下收集覆盖率（go build -cover -covermode=atomic + GOCOVERDIR）。
// go test -cover 的测试二进制在 os.Exit 时不会刷新覆盖率，因此 PrintResult 等触发
// os.Exit 的错误路径必须通过此驱动（go build -cover 构建）来收集。
//
// 工作方式：主进程遍历所有场景，每个场景用子进程执行（子进程设 GOCOVERDIR），
// os.Exit 时运行时钩子刷新覆盖率到 GOCOVERDIR，主进程不受影响。
//
// 受 build tag 保护，不参与正常构建。
//
//go:build cli_coverage_driver
// +build cli_coverage_driver

package main

import (
	"os"
	"os/exec"

	"github.com/scagogogo/versions-skills/internal/cli"
)

// scenarios 是所有触发 os.Exit 的命令场景。
var scenarios = [][]string{
	// ResolveVersions/Strict 错误（无版本输入，stdin 为 char device）
	{"min"}, {"max"}, {"latest-stable"}, {"latest-prerelease"},
	{"filter", "--stable"}, {"count", "--stable"}, {"group"},
	{"partition", "--stable"}, {"visualize"}, {"sort-strings"},
	{"group-ids"}, {"sort"}, {"range", "1.0.0", "3.0.0"},
	{"write", "/tmp/cli_driver_noinput.txt"},
	// LatestStable/LatestPrerelease nil
	{"latest-stable", "1.0.0-alpha", "2.0.0-beta"},
	{"latest-prerelease", "1.0.0", "2.0.0"},
	// compare invalid
	{"compare", "not-a-version", "1.0.0"},
	{"compare", "1.0.0", "not-a-version"},
	// constraint
	{"constraint", ">=1.0.0", "not-a-version"},
	{"constraint", "--type", "single", "notvalid", "1.0.0"},
	{"constraint", "--type", "union", "notvalid", "1.0.0"},
	{"constraint", "--type", "set", "notvalid", "1.0.0"},
	// build
	{"build", "--numbers", "1,abc,3"},
	{"build", "--major", "abc"},
	{"build", "--minor", "abc"},
	{"build", "--patch", "abc"},
	// bump
	{"bump", "not-a-version"},
	{"bump", "1.2.3"},
	// core
	{"core", "not-a-version"},
	// count invalid
	{"count", "--major", "abc", "1.0.0"},
	{"count", "--minor", "abc", "1.0.0"},
	{"count", "--patch", "abc", "1.0.0"},
	// filter invalid
	{"filter", "--major", "abc", "1.0.0"},
	{"filter", "--minor", "abc", "1.0.0"},
	{"filter", "--patch", "abc", "1.0.0"},
	{"filter", "--constraint", "notvalid", "--constraint-type", "single", "1.0.0"},
	{"filter", "--constraint", "notvalid", "--constraint-type", "union", "1.0.0"},
	{"filter", "--constraint", "notvalid", "--constraint-type", "set", "1.0.0"},
	// group --id 不存在
	{"group", "--id", "9.9.9", "1.0.0"},
	// partition 无条件
	{"partition", "1.0.0", "2.0.0"},
	// property invalid
	{"segments", "not-a-version"},
	{"sub-version", "not-a-version"},
	{"suffix-weight", "not-a-version"},
	{"pure-prefix", "not-a-version"},
	{"group-id", "not-a-version"},
	{"clone", "not-a-version"},
	// satisfies: ParseValidVersion 失败 + Matches 失败
	{"satisfies", "not-a-version", ">=1.0.0"},
	{"satisfies", "1.0.0", "notvalid"},
	// set invalid version + invalid value
	{"set-prefix", "not-a-version", "v"},
	{"set-suffix", "not-a-version", "--", "-beta"},
	{"set-major", "not-a-version", "1"},
	{"set-minor", "not-a-version", "1"},
	{"set-patch", "not-a-version", "1"},
	{"set-numbers", "not-a-version", "1,2"},
	{"set-numbers", "1.0.0", "1,abc"},
	{"set-major", "1.0.0", "abc"},
	{"set-minor", "1.0.0", "abc"},
	{"set-patch", "1.0.0", "abc"},
	// validate 无效
	{"validate", "not-a-version"},
	// check: 无效版本 / zero false / 无效目标 / 无类型
	{"check", "--stable", "not-a-version"},
	{"check", "--zero", "0.0.0"},
	{"check", "--newer", "not-a-version", "1.0.0"},
	{"check", "--older", "not-a-version", "1.0.0"},
	{"check", "--equal", "not-a-version", "1.0.0"},
	{"check", "--between-low", "not-a-version", "--between-high", "3.0.0", "2.0.0"},
	{"check", "--between-low", "1.0.0", "--between-high", "not-a-version", "2.0.0"},
	{"check", "1.0.0"},
	// fileio write 失败（写入目录）
	{"write", "/tmp", "1.0.0"},
	// group_extra: PreRunE 缺 --group-id
	{"group-latest", "1.0.0"},
	{"group-oldest", "1.0.0"},
	{"group-stable", "1.0.0"},
	{"group-prerelease", "1.0.0"},
	{"group-latest-stable", "1.0.0"},
	{"group-latest-prerelease", "1.0.0"},
	// group_extra: resolveGroupExtra 错误（有 --group-id 无版本输入）
	{"group-latest", "--group-id", "1.0.0"},
	{"group-oldest", "--group-id", "1.0.0"},
	{"group-stable", "--group-id", "1.0.0"},
	{"group-prerelease", "--group-id", "1.0.0"},
	{"group-latest-stable", "--group-id", "1.0.0"},
	{"group-latest-prerelease", "--group-id", "1.0.0"},
	// group_extra: 分组不存在
	{"group-latest", "--group-id", "9.9.9", "1.0.0"},
	{"group-oldest", "--group-id", "9.9.9", "1.0.0"},
	{"group-stable", "--group-id", "9.9.9", "1.0.0"},
	{"group-prerelease", "--group-id", "9.9.9", "1.0.0"},
	{"group-latest-stable", "--group-id", "9.9.9", "1.0.0"},
	{"group-latest-prerelease", "--group-id", "9.9.9", "1.0.0"},
	// group_extra: LatestStable/LatestPrerelease nil（分组存在但无对应类型）
	{"group-latest-stable", "--group-id", "1.0.0", "1.0.0-alpha"},
	{"group-latest-prerelease", "--group-id", "1.0.0", "1.0.0"},
	// group-contains: 缺 --group-id / ResolveVersions 错误 + 分组不存在 + 缺 --version
	{"group-contains", "1.0.0"},
	{"group-contains", "--group-id", "1.0.0", "--version", "1.0.0-alpha"},
	{"group-contains", "--group-id", "9.9.9", "--version", "1.0.0-alpha", "1.0.0"},
	{"group-contains", "--group-id", "1.0.0", "1.0.0"},
	// read / read-strings 不存在
	{"read", "/nonexistent/cli/x.txt"},
	{"read-strings", "/nonexistent/cli/x.txt"},
}

func main() {
	// 单场景模式：被自身作为子进程调用时执行单个场景
	if len(os.Args) >= 2 && os.Args[1] == "-scenario" {
		args := os.Args[2:]
		os.Args = append([]string{"versions"}, args...)
		if err := cli.Execute(); err != nil {
			os.Exit(1)
		}
		return
	}

	// 主模式：遍历所有场景，每个用子进程执行（带 GOCOVERDIR）
	gocoverdir := os.Getenv("CLI_DRIVER_GOCOVERDIR")
	for _, args := range scenarios {
		cmd := exec.Command(os.Args[0], append([]string{"-scenario"}, args...)...)
		env := os.Environ()
		if gocoverdir != "" {
			env = append(env, "GOCOVERDIR="+gocoverdir)
		}
		cmd.Env = env
		cmd.Stdout = nil
		cmd.Stderr = nil
		_ = cmd.Run()
	}
}
