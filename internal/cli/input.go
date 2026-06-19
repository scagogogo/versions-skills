package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/scagogogo/versions-skills"
)

// ResolveVersions 从多个输入源解析版本列表
// 优先级：fromFile > stdin（当 args 为空时）> args
func ResolveVersions(args []string, fromFile string) ([]*versions.Version, error) {
	var lines []string

	// 1. 从文件读取
	if fromFile != "" {
		fileLines, err := readLinesFromFile(fromFile)
		if err != nil {
			return nil, fmt.Errorf("读取文件 %q 失败: %w", fromFile, err)
		}
		lines = fileLines
	} else if len(args) == 0 {
		// 2. 从 stdin 读取（当没有参数时）
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			stdinLines, err := readLinesFromReader(os.Stdin)
			if err != nil {
				return nil, fmt.Errorf("从 stdin 读取失败: %w", err)
			}
			lines = stdinLines
		}
	} else {
		// 3. 从命令行参数
		lines = args
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("没有提供版本号，请通过参数、--from-file 或 stdin 提供")
	}

	// 解析版本
	result := make([]*versions.Version, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		v := versions.NewVersion(line)
		result = append(result, v)
	}

	return result, nil
}

// ResolveVersionsStrict 从多个输入源解析版本列表，严格模式（过滤无效版本）
func ResolveVersionsStrict(args []string, fromFile string) ([]*versions.Version, error) {
	all, err := ResolveVersions(args, fromFile)
	if err != nil {
		return nil, err
	}

	valid := make([]*versions.Version, 0, len(all))
	for _, v := range all {
		if v.IsValid() {
			valid = append(valid, v)
		}
	}

	if len(valid) == 0 {
		return nil, fmt.Errorf("没有有效的版本号")
	}

	return valid, nil
}

// ParseValidVersion 解析并验证版本号，返回有效版本或错误
func ParseValidVersion(versionStr string) (*versions.Version, error) {
	v := versions.NewVersion(versionStr)
	if !v.IsValid() {
		return nil, fmt.Errorf("无效的版本号: %s", versionStr)
	}
	return v, nil
}

// readLinesFromFile 从文件逐行读取
func readLinesFromFile(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	return readLinesFromReader(f)
}

// readLinesFromReader 从 io.Reader 逐行读取
func readLinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

// ReverseSlice 反转任意切片（通过交换实现）
func ReverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
