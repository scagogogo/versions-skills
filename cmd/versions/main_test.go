package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// goBin 返回当前可用的 go 可执行文件路径。
// 用 exec.LookPath("go") 从 PATH 查找，本机与 CI（setup-go 装的 go 在 PATH）
// 均可用，不依赖 PATH 里的特定版本名（如 go1.25.7）。
func goBin() string {
	bin, err := exec.LookPath("go")
	if err != nil {
		panic("go binary not found in PATH: " + err.Error())
	}
	return bin
}

// TestRun_Help 测 run() 成功路径：--help 让 cli.Execute 返回 nil，run 返回 0。
func TestRun_Help(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"versions", "--help"}
	assert.Equal(t, 0, run())
}

// TestRun_Subprocess 通过子进程覆盖 main() 与 run() 的各路径。
// main() 调 os.Exit 无法在测试进程内直接覆盖，故构建带覆盖率插桩的二进制，
// 在带 GOCOVERDIR 的子进程中执行，由运行时在 os.Exit 时刷新覆盖率。
func TestRun_Subprocess(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	tmpDir := t.TempDir()
	binPath := filepath.Join(tmpDir, "versions")
	gocoverdir := filepath.Join(tmpDir, "cover")
	if err := os.MkdirAll(gocoverdir, 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}

	build := exec.Command(goBin(), "build", "-cover", "-covermode=atomic",
		"-o", binPath, ".")
	build.Env = append(os.Environ(), "GOTOOLCHAIN=local")
	build.Dir = repoRoot(t) + "/cmd/versions"
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}

	runSub := func(args ...string) {
		cmd := exec.Command(binPath, args...)
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+gocoverdir)
		_ = cmd.Run()
	}
	// 成功路径：parse 一个版本（run 返回 0，main 正常 os.Exit(0)）
	runSub("parse", "1.2.3")
	// --help
	runSub("--help")
	// 错误路径：未知命令 → cobra 返回 error → cli.Execute 返回 error → run() return 1
	// → main os.Exit(1)。覆盖 main.go:11-12 的 error 分支。
	runSub("unknown-cmd-xyz")
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatal("go.mod not found")
	return ""
}
