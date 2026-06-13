package mcp

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Server 封装 MCP 服务器
type Server struct {
	server  *server.MCPServer
	version string
}

// NewServer 创建新的 MCP 服务器
func NewServer(version string) (*Server, error) {
	s := &Server{
		version: version,
	}

	mcpServer := server.NewMCPServer(
		"versions",
		version,
		server.WithToolCapabilities(true),
	)

	// 注册所有工具
	s.registerTools(mcpServer)

	s.server = mcpServer
	return s, nil
}

// ServeStdio 以 stdio 模式启动服务器
func (s *Server) ServeStdio() error {
	return server.ServeStdio(s.server)
}

// ServeSSE 以 SSE 模式启动服务器
func (s *Server) ServeSSE(addr string) error {
	sseServer := server.NewSSEServer(s.server)
	log.Printf("MCP SSE 服务器启动在 %s", addr)
	return sseServer.Start(addr)
}

// registerTools 注册所有 MCP 工具
func (s *Server) registerTools(mcpServer *server.MCPServer) {
	tools := []struct {
		tool    mcp.Tool
		handler server.ToolHandlerFunc
	}{
		{s.toolParse(), s.handleParse},
		{s.toolValidate(), s.handleValidate},
		{s.toolInfo(), s.handleInfo},
		{s.toolCompare(), s.handleCompare},
		{s.toolSort(), s.handleSort},
		{s.toolGroup(), s.handleGroup},
		{s.toolConstraintCheck(), s.handleConstraintCheck},
		{s.toolRangeQuery(), s.handleRangeQuery},
		{s.toolFilter(), s.handleFilter},
		{s.toolMin(), s.handleMin},
		{s.toolMax(), s.handleMax},
		{s.toolLatestStable(), s.handleLatestStable},
		{s.toolLatestPrerelease(), s.handleLatestPrerelease},
		{s.toolUnique(), s.handleUnique},
		{s.toolSetOperation(), s.handleSetOperation},
		{s.toolVisualize(), s.handleVisualize},
		{s.toolBuild(), s.handleBuild},
		{s.toolBump(), s.handleBump},
		{s.toolCore(), s.handleCore},
		{s.toolReadFile(), s.handleReadFile},
		{s.toolWriteFile(), s.handleWriteFile},
	}

	for _, t := range tools {
		mcpServer.AddTool(t.tool, t.handler)
	}
}

// getArgs 获取请求参数 map（兼容 mcp-go v0.32 的 any 类型 Arguments）
func getArgs(request mcp.CallToolRequest) map[string]any {
	return request.GetArguments()
}

// 辅助函数：从请求参数中获取字符串
func getStringParam(request mcp.CallToolRequest, key string) (string, error) {
	args := getArgs(request)
	val, ok := args[key]
	if !ok {
		return "", fmt.Errorf("缺少参数: %s", key)
	}
	str, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("参数 %s 必须是字符串", key)
	}
	return str, nil
}

// 辅助函数：从请求参数中获取可选字符串
func getOptionalStringParam(request mcp.CallToolRequest, key string) string {
	args := getArgs(request)
	val, ok := args[key]
	if !ok {
		return ""
	}
	str, ok := val.(string)
	if !ok {
		return ""
	}
	return str
}

// 辅助函数：从请求参数中获取布尔值
func getBoolParam(request mcp.CallToolRequest, key string) bool {
	args := getArgs(request)
	val, ok := args[key]
	if !ok {
		return false
	}
	b, ok := val.(bool)
	if !ok {
		return false
	}
	return b
}

// 辅助函数：从请求参数中获取可选数字
func getOptionalNumberParam(request mcp.CallToolRequest, key string) (int, bool) {
	args := getArgs(request)
	val, ok := args[key]
	if !ok {
		return 0, false
	}
	f, ok := val.(float64)
	if !ok {
		return 0, false
	}
	return int(f), true
}

// 辅助函数：从请求参数中获取版本字符串数组
func getVersionStringsParam(request mcp.CallToolRequest, key string) ([]string, error) {
	args := getArgs(request)
	val, ok := args[key]
	if !ok {
		return nil, fmt.Errorf("缺少参数: %s", key)
	}
	arr, ok := val.([]interface{})
	if !ok {
		return nil, fmt.Errorf("参数 %s 必须是字符串数组", key)
	}
	result := make([]string, len(arr))
	for i, v := range arr {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("参数 %s 的元素必须是字符串", key)
		}
		result[i] = s
	}
	return result, nil
}

// 辅助函数：将错误转为 MCP 结果
func errorResult(err error) *mcp.CallToolResult {
	return mcp.NewToolResultError(err.Error())
}

// 辅助函数：将数据转为 MCP JSON 结果
func jsonResult(data interface{}) *mcp.CallToolResult {
	jsonBytes, err := marshalJSON(data)
	if err != nil {
		return errorResult(fmt.Errorf("JSON 序列化失败: %w", err))
	}
	return mcp.NewToolResultText(string(jsonBytes))
}

// 确保 context 被引用
var _ = context.Background
