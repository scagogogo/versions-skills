package main

import (
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

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

// TestRun_Help 测 run() 的成功路径：--help 进入 cobra help，不触发 Run 闭包，
// Execute 返回 nil，run 返回 0。覆盖命令组装、flag 绑定、Execute 成功路径。
func TestRun_Help(t *testing.T) {
	defer resetFlags()
	exitCode := run([]string{"--help"})
	assert.Equal(t, 0, exitCode)
}

func resetFlags() {
	transportFlag = "stdio"
	portFlag = 8080
	versionFlag = "0.0.0-dev"
}

// TestRun_Subprocess 通过子进程覆盖 Run 闭包内的路径（ServeStdio/ServeSSE/bad transport/os.Exit）。
// Run 闭包内调 os.Exit / log.Fatalf，无法在测试进程内直接覆盖，故用带覆盖率插桩的子进程。
func TestRun_Subprocess(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping subprocess test in short mode")
	}

	tmpDir := t.TempDir()
	binPath := filepath.Join(tmpDir, "versions-mcp")
	gocoverdir := filepath.Join(tmpDir, "cover")
	if err := os.MkdirAll(gocoverdir, 0o755); err != nil {
		t.Fatal(err)
	}

	// 构建带覆盖率插桩的二进制
	build := exec.Command(goBin(), "build", "-cover", "-covermode=atomic",
		"-o", binPath, ".")
	build.Env = append(os.Environ(), "GOTOOLCHAIN=local")
	build.Dir = repoRoot(t)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}

	runSub := func(name string, args []string, stdin string) {
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command(binPath, args...)
			cmd.Env = append(os.Environ(), "GOCOVERDIR="+gocoverdir)
			cmd.Stdin = strings.NewReader(stdin)
			_ = cmd.Run() // 退出码不重要，只要收集覆盖率
		})
	}

	// stdio 模式：stdin 立即 EOF → ServeStdio 返回 nil（成功路径，不触发 log.Fatalf）
	runSub("stdio", []string{"--transport", "stdio"}, "")

	// 未知 flag → cobra 返回 error → Execute 返回 error → run() return 1
	// 覆盖 main.go:52-53 的 error 分支（不进入 Run 闭包）。
	runSub("unknown-flag", []string{"--unknown-flag-xyz"}, "")

	// stdio 模式 + SIGINT：ServeStdio 监听信号，收到 SIGINT 后 ctx cancel，
	// processInputStream 因 ctx.Err() 返回 context.Canceled → ServeStdio 返回 error
	// → 触发 main.go:32-34 的 log.Fatalf 路径。
	// 用 StdinPipe 保持 stdin 打开（不写入不关闭 → 子进程读阻塞，不 EOF），
	// 否则 stdin EOF 让 ServeStdio 提前返回 nil 退出，SIGINT 无的放矢。
	{
		cmd := exec.Command(binPath, "--transport", "stdio")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+gocoverdir)
		stdinPipe, err := cmd.StdinPipe()
		if err != nil {
			t.Fatalf("StdinPipe: %v", err)
		}
		if err := cmd.Start(); err == nil {
			time.Sleep(200 * time.Millisecond)
			_ = cmd.Process.Signal(os.Interrupt)
			_, _ = cmd.Process.Wait()
		}
		_ = stdinPipe.Close()
	}
	// 不支持的 transport → os.Exit(1)
	runSub("bad-transport", []string{"--transport", "bad-xyz"}, "")

	// sse 端口冲突：先占住端口，再跑第二个会 log.Fatalf
	port := "19876"
	blocker := exec.Command(binPath, "--transport", "sse", "--port", port)
	blocker.Env = append(os.Environ(), "GOCOVERDIR="+gocoverdir)
	blocker.Stdin = strings.NewReader("")
	if err := blocker.Start(); err == nil {
		waitForPort(t, port)
		runSub("sse-port-conflict", []string{"--transport", "sse", "--port", port}, "")
		_ = blocker.Process.Kill()
	}
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

func waitForPort(t *testing.T, port string) {
	t.Helper()
	for i := 0; i < 50; i++ {
		c, err := net.Dial("tcp", "127.0.0.1:"+port)
		if err == nil {
			_ = c.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
}
