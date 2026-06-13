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

func main() {
	rootCmd := &cobra.Command{
		Use:   "versions-mcp",
		Short: "Versions MCP 服务器",
		Long:  `Versions MCP 服务器，提供版本号操作的 MCP 协议接口。`,
		Run: func(cmd *cobra.Command, args []string) {
			server, err := mcpserver.NewServer(versionFlag)
			if err != nil {
				fmt.Fprintf(os.Stderr, "创建 MCP 服务器失败: %v\n", err)
				os.Exit(1)
			}

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

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
