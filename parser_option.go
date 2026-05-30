package versions

// ParserOption 配置版本号解析器的行为
//
// ParserOption 允许调用者自定义解析器支持的数字分隔符，
// 以适应不同语言生态系统的版本号格式差异。
//
// 使用示例:
//
//	// 支持 underscore 分隔（Python/RPM 生态）
//	v := versions.NewVersionWithOption("1_2_3",
//	    versions.ParserOption{Delimiters: ".-_"})
type ParserOption struct {
	// Delimiters 版本号数字部分的分隔符集合
	// 默认值: "." (仅点号)
	// 常见扩展: ".-_" (支持 RPM/Debian 的连字符和 Python 的下划线)
	Delimiters string
}

// DefaultParserOption 返回默认的解析器选项
func DefaultParserOption() ParserOption {
	return ParserOption{
		Delimiters: ".",
	}
}

// NewVersionWithOption 使用指定选项创建版本对象
func NewVersionWithOption(versionStr string, option ParserOption) *Version {
	return NewVersionStringParserWithOptions(versionStr, option).Parse()
}

// NewVersionStringParserWithOptions 创建带选项的版本号解析器
func NewVersionStringParserWithOptions(versionStr string, option ParserOption) *VersionStringParser {
	p := NewVersionStringParser(versionStr)
	p.option = option
	return p
}
