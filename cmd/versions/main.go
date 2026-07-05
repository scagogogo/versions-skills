package main

import (
	"os"

	"github.com/scagogogo/versions-skills/internal/cli"
)

// run 是 main 的可测试版本：执行 CLI 并返回退出码。
func run() int {
	if err := cli.Execute(); err != nil {
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
