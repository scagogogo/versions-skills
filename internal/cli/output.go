package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

// Result 是所有命令输出的统一信封格式
type Result struct {
	Command string      `json:"command"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PrintResult 按指定格式输出结果
func PrintResult(cmdName string, data interface{}, err error) {
	// 静默模式：仅输出数据部分，不含信封
	if isQuiet() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		if data != nil {
			encoder := json.NewEncoder(os.Stdout)
			encoder.SetIndent("", "  ")
			encoder.Encode(data)
		}
		return
	}

	result := Result{
		Command: cmdName,
		Success: err == nil,
	}
	if err != nil {
		result.Error = err.Error()
	} else {
		result.Data = data
	}

	switch getFormat() {
	case "json":
		printJSON(result)
	case "table":
		printTable(result)
	case "text":
		printText(result)
	default:
		printJSON(result)
	}

	// 如果命令失败，设置退出码
	if err != nil {
		os.Exit(1)
	}
}

// printJSON 以 JSON 格式输出
func printJSON(result Result) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		fmt.Fprintf(os.Stderr, "JSON 编码失败: %v\n", err)
		os.Exit(1)
	}
}

// printTable 以表格格式输出
func printTable(result Result) {
	if !result.Success {
		fmt.Fprintf(os.Stderr, "错误: %s\n", result.Error)
		return
	}
	printDataAsTable(os.Stdout, result.Data)
}

// printText 以纯文本格式输出
func printText(result Result) {
	if !result.Success {
		fmt.Fprintf(os.Stderr, "错误: %s\n", result.Error)
		return
	}
	printDataAsText(os.Stdout, result.Data)
}

// printDataAsTable 将数据以表格形式输出
func printDataAsTable(w io.Writer, data interface{}) {
	switch d := data.(type) {
	case map[string]interface{}:
		printMapAsTable(w, d)
	case []map[string]interface{}:
		printSliceOfMapAsTable(w, d)
	case []string:
		for _, s := range d {
			fmt.Fprintln(w, s)
		}
	default:
		fmt.Fprintln(w, d)
	}
}

// printDataAsText 将数据以纯文本形式输出
func printDataAsText(w io.Writer, data interface{}) {
	switch d := data.(type) {
	case map[string]interface{}:
		printMapAsText(w, d, 0)
	case []map[string]interface{}:
		for i, item := range d {
			if i > 0 {
				fmt.Fprintln(w)
			}
			printMapAsText(w, item, 0)
		}
	case []string:
		for _, s := range d {
			fmt.Fprintln(w, s)
		}
	default:
		fmt.Fprintln(w, d)
	}
}

// printMapAsTable 将 map 以 key-value 表格形式输出
func printMapAsTable(w io.Writer, m map[string]interface{}) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	for k, v := range m {
		switch val := v.(type) {
		case []interface{}:
			items := make([]string, len(val))
			for i, item := range val {
				items[i] = fmt.Sprintf("%v", item)
			}
			fmt.Fprintf(tw, "%s\t%s\n", k, strings.Join(items, "."))
		case map[string]interface{}:
			fmt.Fprintf(tw, "%s\t%v\n", k, val)
		default:
			fmt.Fprintf(tw, "%s\t%v\n", k, v)
		}
	}
	tw.Flush()
}

// printSliceOfMapAsTable 将 []map[string]interface{} 以对齐表格形式输出
func printSliceOfMapAsTable(w io.Writer, items []map[string]interface{}) {
	if len(items) == 0 {
		return
	}

	// 收集所有 key（保持顺序）
	keySet := make([]string, 0)
	keySeen := make(map[string]bool)
	for _, item := range items {
		for k := range item {
			if !keySeen[k] {
				keySeen[k] = true
				keySet = append(keySet, k)
			}
		}
	}

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	// 表头
	fmt.Fprintf(tw, "%s", strings.Join(keySet, "\t"))
	fmt.Fprintln(tw)

	// 数据行
	for _, item := range items {
		vals := make([]string, len(keySet))
		for i, k := range keySet {
			v, ok := item[k]
			if !ok {
				vals[i] = ""
				continue
			}
			switch val := v.(type) {
			case []interface{}:
				items := make([]string, len(val))
				for j, item := range val {
					items[j] = fmt.Sprintf("%v", item)
				}
				vals[i] = strings.Join(items, ".")
			default:
				vals[i] = fmt.Sprintf("%v", val)
			}
		}
		fmt.Fprintf(tw, "%s", strings.Join(vals, "\t"))
		fmt.Fprintln(tw)
	}
	tw.Flush()
}

// printMapAsText 将 map 以缩进文本形式输出
func printMapAsText(w io.Writer, m map[string]interface{}, indent int) {
	prefix := strings.Repeat("  ", indent)
	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			fmt.Fprintf(w, "%s%s:\n", prefix, k)
			printMapAsText(w, val, indent+1)
		case []interface{}:
			items := make([]string, len(val))
			for i, item := range val {
				items[i] = fmt.Sprintf("%v", item)
			}
			fmt.Fprintf(w, "%s%s: %s\n", prefix, k, strings.Join(items, "."))
		default:
			fmt.Fprintf(w, "%s%s: %v\n", prefix, k, v)
		}
	}
}
