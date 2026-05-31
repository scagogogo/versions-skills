package versions

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestVisualizeVersions 测试版本号可视化函数
func TestVisualizeVersions(t *testing.T) {
	// 准备测试数据
	versions := []*Version{
		{
			Raw:            "1.0.0",
			VersionNumbers: []int{1, 0, 0},
			PublicTime:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Raw:            "1.1.0",
			VersionNumbers: []int{1, 1, 0},
			PublicTime:     time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Raw:            "1.2.0",
			VersionNumbers: []int{1, 2, 0},
			PublicTime:     time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Raw:            "2.0.0",
			VersionNumbers: []int{2, 0, 0},
			PublicTime:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Raw:            "2.1.0",
			VersionNumbers: []int{2, 1, 0},
			PublicTime:     time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	// 使用缓冲区捕获输出
	buf := &bytes.Buffer{}

	// 测试默认显示（不限制数量）
	VisualizeVersions(versions, buf, 0)
	output := buf.String()

	// 验证输出包含所有信息
	assert.Contains(t, output, "版本总数: 5")
	assert.Contains(t, output, "版本组数: 5")
	assert.Contains(t, output, "版本组: 1")
	assert.Contains(t, output, "版本组: 2")
	assert.Contains(t, output, "1.0.0 (发布时间: 2020-01-01)")
	assert.Contains(t, output, "2.0.0 (发布时间: 2021-01-01)")

	// 测试限制显示项目数
	buf.Reset()
	VisualizeVersions(versions, buf, 1)
	limitedOutput := buf.String()

	// 验证限制输出
	assert.Contains(t, limitedOutput, "版本组: 1")
}

// TestVisualizeVersionGroups 测试版本组可视化函数
func TestVisualizeVersionGroups(t *testing.T) {
	// 准备测试数据 - 创建更复杂的版本结构
	versions := []*Version{
		{
			Raw:            "1.0.0",
			VersionNumbers: []int{1, 0, 0},
		},
		{
			Raw:            "1.1.0",
			VersionNumbers: []int{1, 1, 0},
		},
		{
			Raw:            "1.1.1",
			VersionNumbers: []int{1, 1, 1},
		},
		{
			Raw:            "2.0.0",
			VersionNumbers: []int{2, 0, 0},
		},
		{
			Raw:            "2.1.0",
			VersionNumbers: []int{2, 1, 0},
		},
		{
			Raw:            "2.1.1",
			VersionNumbers: []int{2, 1, 1},
		},
		{
			Raw:            "10",
			VersionNumbers: []int{10},
		},
	}

	// 使用缓冲区捕获输出
	buf := &bytes.Buffer{}
	VisualizeVersionGroups(versions, buf)
	output := buf.String()

	// 验证输出包含基本信息
	assert.Contains(t, output, "版本总数: 7")
	assert.Contains(t, output, "版本组数: 7")

	// 验证树形结构
	lines := strings.Split(output, "\n")
	treeLines := 0
	for _, line := range lines {
		if strings.Contains(line, "├─") || strings.Contains(line, "└─") {
			treeLines++
		}
	}

	// 验证树形结构包含所有组
	assert.GreaterOrEqual(t, treeLines, 5, "树形结构应包含所有版本组")

	// 验证根节点存在
	rootFound := false
	for _, line := range lines {
		if strings.Contains(line, "10 (") {
			rootFound = true
			break
		}
	}
	assert.True(t, rootFound, "应该显示单数字版本组作为根节点")
}
