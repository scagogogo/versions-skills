package cli

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintJSON_Success(t *testing.T) {
	withFormat(t, "json")
	restore, get := captureStdout(t)
	printJSON(Result{Command: "test", Success: true, Data: "x"})
	restore()
	out := get()
	assert.Contains(t, out, `"command": "test"`)
	assert.Contains(t, out, `"success": true`)
}

func TestPrintTable_Success_Map(t *testing.T) {
	withFormat(t, "table")
	restore, get := captureStdout(t)
	printTable(Result{Command: "test", Success: true, Data: map[string]interface{}{"a": 1, "b": "two"}})
	restore()
	out := get()
	assert.Contains(t, out, "a")
	assert.Contains(t, out, "two")
}

func TestPrintTable_Failure(t *testing.T) {
	withFormat(t, "table")
	// printTable 在 failure 时写 stderr，不写 stdout
	restore, get := captureStdout(t)
	printTable(Result{Command: "test", Success: false, Error: "boom"})
	restore()
	assert.Empty(t, get())
}

func TestPrintText_Success_Map(t *testing.T) {
	withFormat(t, "text")
	restore, get := captureStdout(t)
	printText(Result{Command: "test", Success: true, Data: map[string]interface{}{"k": "v"}})
	restore()
	out := get()
	assert.Contains(t, out, "k: v")
}

func TestPrintText_Failure(t *testing.T) {
	withFormat(t, "text")
	restore, get := captureStdout(t)
	printText(Result{Command: "test", Success: false, Error: "boom"})
	restore()
	assert.Empty(t, get())
}

func TestPrintDataAsTable_Map(t *testing.T) {
	var sb strings.Builder
	printDataAsTable(&sb, map[string]interface{}{"k1": "v1", "k2": 42})
	out := sb.String()
	assert.Contains(t, out, "k1")
	assert.Contains(t, out, "v1")
	assert.Contains(t, out, "42")
}

func TestPrintDataAsTable_SliceOfMap(t *testing.T) {
	var sb strings.Builder
	items := []map[string]interface{}{
		{"name": "a", "ver": "1.0.0"},
		{"name": "b", "ver": "2.0.0"},
	}
	printDataAsTable(&sb, items)
	out := sb.String()
	assert.Contains(t, out, "name")
	assert.Contains(t, out, "1.0.0")
	assert.Contains(t, out, "2.0.0")
}

func TestPrintDataAsTable_SliceOfString(t *testing.T) {
	var sb strings.Builder
	printDataAsTable(&sb, []string{"x", "y", "z"})
	out := sb.String()
	assert.Contains(t, out, "x")
	assert.Contains(t, out, "z")
}

func TestPrintDataAsTable_Default(t *testing.T) {
	var sb strings.Builder
	printDataAsTable(&sb, 12345)
	assert.Contains(t, sb.String(), "12345")
}

func TestPrintDataAsTable_MapWithSliceValue(t *testing.T) {
	var sb strings.Builder
	printDataAsTable(&sb, map[string]interface{}{"nums": []interface{}{1, 2, 3}})
	out := sb.String()
	// 切片以 . 连接
	assert.Contains(t, out, "1.2.3")
}

func TestPrintDataAsTable_MapWithMapValue(t *testing.T) {
	var sb strings.Builder
	printDataAsTable(&sb, map[string]interface{}{"nested": map[string]interface{}{"a": 1}})
	out := sb.String()
	assert.Contains(t, out, "nested")
	assert.Contains(t, out, "a")
}

func TestPrintDataAsText_Map(t *testing.T) {
	var sb strings.Builder
	printDataAsText(&sb, map[string]interface{}{"k": "v"})
	assert.Contains(t, sb.String(), "k: v")
}

func TestPrintDataAsText_SliceOfMap(t *testing.T) {
	var sb strings.Builder
	items := []map[string]interface{}{
		{"a": "1"},
		{"b": "2"},
	}
	printDataAsText(&sb, items)
	out := sb.String()
	assert.Contains(t, out, "a: 1")
	assert.Contains(t, out, "b: 2")
}

func TestPrintDataAsText_SliceOfString(t *testing.T) {
	var sb strings.Builder
	printDataAsText(&sb, []string{"a", "b"})
	out := sb.String()
	assert.Contains(t, out, "a")
	assert.Contains(t, out, "b")
}

func TestPrintDataAsText_Default(t *testing.T) {
	var sb strings.Builder
	printDataAsText(&sb, "hello")
	assert.Contains(t, sb.String(), "hello")
}

func TestPrintMapAsText_NestedMap(t *testing.T) {
	var sb strings.Builder
	m := map[string]interface{}{
		"outer": map[string]interface{}{
			"inner": "val",
		},
		"nums":  []interface{}{1, 2},
		"plain": "p",
	}
	printMapAsText(&sb, m, 0)
	out := sb.String()
	assert.Contains(t, out, "outer:")
	assert.Contains(t, out, "inner: val")
	assert.Contains(t, out, "nums: 1.2")
	assert.Contains(t, out, "plain: p")
}

func TestPrintSliceOfMapAsTable_Empty(t *testing.T) {
	var sb strings.Builder
	printSliceOfMapAsTable(&sb, []map[string]interface{}{})
	assert.Empty(t, sb.String())
}

func TestPrintSliceOfMapAsTable_WithSliceValuesAndMissingKeys(t *testing.T) {
	var sb strings.Builder
	items := []map[string]interface{}{
		{"name": "a", "nums": []interface{}{1, 2}},
		{"name": "b"}, // nums 缺失
	}
	printSliceOfMapAsTable(&sb, items)
	out := sb.String()
	assert.Contains(t, out, "name")
	assert.Contains(t, out, "1.2")
}

func TestPrintMapAsTable(t *testing.T) {
	var sb strings.Builder
	m := map[string]interface{}{
		"s":   "str",
		"n":   7,
		"arr": []interface{}{9, 8},
		"mp":  map[string]interface{}{"x": 1},
	}
	printMapAsTable(&sb, m)
	out := sb.String()
	assert.Contains(t, out, "str")
	assert.Contains(t, out, "9.8")
}

// PrintResult 非退出路径

func TestPrintResult_QuietWithNilData(t *testing.T) {
	withQuiet(t, true)
	restore, get := captureStdout(t)
	PrintResult("test", nil, nil)
	restore()
	assert.Empty(t, get())
}

func TestPrintResult_QuietWithData(t *testing.T) {
	withQuiet(t, true)
	restore, get := captureStdout(t)
	PrintResult("test", map[string]interface{}{"k": "v"}, nil)
	restore()
	out := get()
	assert.Contains(t, out, `"k": "v"`)
}

func TestPrintResult_JsonSuccess(t *testing.T) {
	withFormat(t, "json")
	withQuiet(t, false)
	restore, get := captureStdout(t)
	PrintResult("test", map[string]interface{}{"k": "v"}, nil)
	restore()
	out := get()
	assert.Contains(t, out, `"success": true`)
}

func TestPrintResult_TableSuccess(t *testing.T) {
	withFormat(t, "table")
	restore, get := captureStdout(t)
	PrintResult("test", map[string]interface{}{"k": "v"}, nil)
	restore()
	assert.Contains(t, get(), "k")
}

func TestPrintResult_TextSuccess(t *testing.T) {
	withFormat(t, "text")
	restore, get := captureStdout(t)
	PrintResult("test", map[string]interface{}{"k": "v"}, nil)
	restore()
	assert.Contains(t, get(), "k: v")
}

func TestPrintResult_DefaultFormat(t *testing.T) {
	// format 为未知值时走 default 分支 → printJSON
	withFormat(t, "weird")
	restore, get := captureStdout(t)
	PrintResult("test", "x", nil)
	restore()
	assert.Contains(t, get(), `"command": "test"`)
}

func TestPrintResult_TableError_NoExit(t *testing.T) {
	// table 格式且 err != nil → printTable 写 stderr，不 os.Exit（printTable 失败路径不退出）
	// 但 PrintResult 末尾会 os.Exit(1)，所以这个路径用子进程测（见 subprocess_test.go）。
	// 这里只验证 printTable 直接调用（已在 TestPrintTable_Failure 覆盖）。
	t.Log("covered by TestPrintTable_Failure and subprocess tests")
}
