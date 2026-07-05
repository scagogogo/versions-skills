package main

import (
	"fmt"
	"log"
	"os"

	mcpserver "github.com/scagogogo/versions-skills/internal/mcp"
	"github.com/spf13/cobra"
)

var (
	transportFlag = "stdio"
	portFlag      = 8080
	versionFlag   = "0.0.0-dev"
)

// run 是 main 的可测试版本：执行命令并返回退出码。
// 抽出为函数以便单元测试覆盖，main 仅做 os.Exit 中转。
func run(args []string) int {
	rootCmd := &cobra.Command{
		Use:   "versions-mcp",
		Short: "Versions MCP 服务器",
		Long:  `Versions MCP 服务器，提供版本号操作的 MCP 协议接口。`,
		Run: func(cmd *cobra.Command, args []string) {
			// NewServer 当前实现恒返回 nil error（仅组装 mcpServer + 注册工具），
			// 故此处不再检查 error；若未来 NewServer 可能失败，需在此补错误处理。
			server, _ := mcpserver.NewServer(versionFlag)

			switch transportFlag {
			case "stdio":
				if err := server.ServeStdio(); err != nil {
					log.Fatalf("MCP stdio 服务器错误: %v", err)
				}
			case "sse":
				addr := fmt.Sprintf(":%d", portFlag)
				if err := server.ServeSSE(addr); err != nil {
					log.Fatalf("MCP SSE 服务器错误: %v", err)
				}
			default:
				fmt.Fprintf(os.Stderr, "不支持的传输方式: %s\n", transportFlag)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringVar(&transportFlag, "transport", "stdio", "传输方式: stdio|sse")
	rootCmd.Flags().IntVar(&portFlag, "port", 8080, "SSE 模式监听端口")
	rootCmd.Flags().StringVar(&versionFlag, "version", "0.0.0-dev", "服务器版本号")

	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		return 1
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
