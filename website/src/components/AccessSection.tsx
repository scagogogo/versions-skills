import { Typography, Row, Col, Card, Tag, Steps } from 'antd'
import {
  RobotOutlined,
  ApiOutlined,
  CodeSandboxOutlined,
  CodeOutlined,
  CheckCircleOutlined,
} from '@ant-design/icons'

const { Title, Paragraph, Text } = Typography

const accessMethods = [
  {
    icon: <RobotOutlined style={{ fontSize: 36, color: '#722ed1' }} />,
    tag: '推荐',
    tagColor: 'purple',
    title: '🤖 Skills（Claude Code）',
    description: 'Claude Code 读取 SKILL.md 文件作为领域知识，通过斜杠命令直接调用。适合引导式工作流和一次性任务。',
    install: 'claude plugin install versions',
    detail: '安装后获得 13 个斜杠命令：/version-parsing、/version-comparison 等',
  },
  {
    icon: <ApiOutlined style={{ fontSize: 36, color: '#1677ff' }} />,
    tag: '通用',
    tagColor: 'blue',
    title: '🔌 MCP Server',
    description: '任何 MCP 兼容客户端直接调用 version_* 工具。适合编程式调用、批量操作和非 Claude Agent。',
    install: 'go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest',
    detail: '暴露 21 个 AI 可调用工具，兼容 Claude Code / Cursor / Windsurf / VS Code Copilot',
  },
  {
    icon: <CodeSandboxOutlined style={{ fontSize: 36, color: '#52c41a' }} />,
    tag: 'Go 开发者',
    tagColor: 'green',
    title: '📦 Go SDK',
    description: '直接在 Go 代码中使用完整的版本操作 API。适合集成到 Go 项目中。',
    install: 'go get github.com/scagogogo/versions-skills',
    detail: '零依赖核心库，完整的 Version 类型和方法链',
  },
  {
    icon: <CodeOutlined style={{ fontSize: 36, color: '#fa8c16' }} />,
    tag: '脚本/CI',
    tagColor: 'orange',
    title: '💻 CLI',
    description: '命令行工具，40+ 子命令覆盖全部功能。适合 Shell 脚本和 CI/CD 管道。',
    install: 'curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash',
    detail: '支持 6 种 OS × 12 种架构，提供 deb/rpm/apk 包',
  },
]

const AccessSection: React.FC = () => {
  return (
    <div id="access" style={{ padding: '80px 48px', background: '#fff' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <Title level={2} style={{ textAlign: 'center', marginBottom: 8 }}>
          四种接入方式
        </Title>
        <Paragraph
          style={{ textAlign: 'center', color: '#666', fontSize: 16, marginBottom: 48 }}
        >
          根据你的场景选择最合适的接入方式，也可以组合使用获得最佳体验
        </Paragraph>

        <Row gutter={[24, 24]}>
          {accessMethods.map((method, index) => (
            <Col xs={24} sm={12} key={index}>
              <Card
                hoverable
                style={{ height: '100%' }}
                styles={{
                  body: { padding: 28 },
                }}
              >
                <div style={{ display: 'flex', alignItems: 'center', marginBottom: 16 }}>
                  {method.icon}
                  <Tag
                    color={method.tagColor}
                    style={{ marginLeft: 12, fontSize: 13, padding: '2px 10px' }}
                  >
                    {method.tag}
                  </Tag>
                </div>
                <Title level={4} style={{ marginBottom: 8 }}>
                  {method.title}
                </Title>
                <Paragraph style={{ color: '#555', marginBottom: 16 }}>
                  {method.description}
                </Paragraph>
                <div
                  style={{
                    background: '#f6f8fa',
                    borderRadius: 6,
                    padding: '10px 14px',
                    fontFamily: 'monospace',
                    fontSize: 13,
                    marginBottom: 12,
                    overflow: 'auto',
                  }}
                >
                  <Text code style={{ border: 'none', background: 'transparent' }}>
                    {method.install}
                  </Text>
                </div>
                <Text type="secondary" style={{ fontSize: 13 }}>
                  {method.detail}
                </Text>
              </Card>
            </Col>
          ))}
        </Row>

        <div style={{ marginTop: 48 }}>
          <Card
            style={{
              background: 'linear-gradient(135deg, #f0f5ff 0%, #f9f0ff 100%)',
              border: '1px solid #d6e4ff',
            }}
          >
            <Title level={4} style={{ marginBottom: 16 }}>
              💡 Skills + MCP 组合使用效果最佳
            </Title>
            <Steps
              direction="horizontal"
              current={-1}
              items={[
                {
                  title: 'Skills 提供知识',
                  description: '告诉 AI "如何做"：API 参考、代码示例、决策树',
                  icon: <CheckCircleOutlined style={{ color: '#722ed1' }} />,
                },
                {
                  title: 'MCP 提供引擎',
                  description: '让 AI "执行"：21 个结构化 JSON 工具调用',
                  icon: <CheckCircleOutlined style={{ color: '#1677ff' }} />,
                },
                {
                  title: '两者配合',
                  description: '知识 + 执行 = 最佳 AI 版本操作体验',
                  icon: <CheckCircleOutlined style={{ color: '#52c41a' }} />,
                },
              ]}
            />
          </Card>
        </div>
      </div>
    </div>
  )
}

export default AccessSection
