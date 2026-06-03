package versions

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadVersionsFromFile 测试从文件读取版本号功能
//
// 这个测试函数验证 ReadVersionsFromFile 函数能否正确地从文件中读取版本号列表。
// 测试使用 test_data 目录下的 fast_json_versions.txt 文件作为测试数据源。
//
// 测试步骤:
// 1. 调用 ReadVersionsFromFile 读取测试文件中的版本号
// 2. 验证读取过程中没有发生错误（error 为 nil）
// 3. 验证成功读取到版本号（返回的数组长度大于0）
//
// 预期结果:
// - 函数应当成功读取文件并解析其中的版本号
// - 返回的版本号数组不应为空
func TestReadVersionsFromFile(t *testing.T) {
	versions, err := ReadVersionsFromFile("test_data/fast_json_versions.txt")
	assert.Nil(t, err)
	assert.True(t, len(versions) > 0)
}

func TestWriteVersionsToFile(t *testing.T) {
	versions := NewVersions("2.0.0", "1.0.0", "1.1.0")
	tmpFile := filepath.Join(t.TempDir(), "versions.txt")
	err := WriteVersionsToFile(versions, tmpFile)
	if err != nil {
		t.Fatalf("WriteVersionsToFile error: %v", err)
	}
	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "1.0.0") {
		t.Error("Output should contain 1.0.0")
	}
	if !strings.Contains(content, "1.1.0") {
		t.Error("Output should contain 1.1.0")
	}
	if !strings.Contains(content, "2.0.0") {
		t.Error("Output should contain 2.0.0")
	}
}

func TestReadVersionsFromReader(t *testing.T) {
	data := strings.NewReader("1.0.0\n1.1.0\n2.0.0\n")
	versions, err := ReadVersionsFromReader(data)
	if err != nil {
		t.Fatalf("ReadVersionsFromReader error: %v", err)
	}
	if len(versions) != 3 {
		t.Errorf("len(versions) = %d, want 3", len(versions))
	}
	if versions[0].Raw != "1.0.0" {
		t.Errorf("versions[0].Raw = %q, want %q", versions[0].Raw, "1.0.0")
	}
}
