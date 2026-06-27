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
    icon: <RobotOutlined style={{ fontSize: 36, color: '#7c3aed' }} />,
    tag: '推荐',
    tagColor: 'purple',
    title: '🤖 Skills（Claude Code）',
    description: 'Claude Code 读取 SKILL.md 文件作为领域知识，通过斜杠命令直接调用。适合引导式工作流和一次性任务。',
    install: 'claude plugin install versions',
    detail: '安装后获得 13 个斜杠命令：/version-parsing、/version-comparison 等',
  },
  {
    icon: <ApiOutlined style={{ fontSize: 36, color: '#2563eb' }} />,
    tag: '通用',
    tagColor: 'blue',
    title: '🔌 MCP Server',
    description: '任何 MCP 兼容客户端直接调用 version_* 工具。适合编程式调用、批量操作和非 Claude Agent。',
    install: 'go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest',
    detail: '暴露 21 个 AI 可调用工具，兼容 Claude Code / Cursor / Windsurf / VS Code Copilot',
  },
  {
    icon: <CodeSandboxOutlined style={{ fontSize: 36, color: '#059669' }} />,
    tag: 'Go 开发者',
    tagColor: 'green',
    title: '📦 Go SDK',
    description: '直接在 Go 代码中使用完整的版本操作 API。适合集成到 Go 项目中。',
    install: 'go get github.com/scagogogo/versions-skills',
    detail: '零依赖核心库，完整的 Version 类型和方法链',
  },
  {
    icon: <CodeOutlined style={{ fontSize: 36, color: '#d97706' }} />,
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
    <div id="access" style={{ padding: '80px 32px', background: '#f8fafc' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">四种接入方式</div>
          <p className="section-subtitle">根据你的场景选择最合适的接入方式，也可以组合使用获得最佳体验</p>
        </div>

        <Row gutter={[20, 20]}>
          {accessMethods.map((method, index) => (
            <Col xs={24} sm={12} key={index}>
              <Card
                hoverable
                className="hover-card"
                style={{ height: '100%', borderRadius: 16, border: '1px solid #e2e8f0' }}
                styles={{
                  body: { padding: 28 },
                }}
              >
                <div style={{ display: 'flex', alignItems: 'center', marginBottom: 16 }}>
                  {method.icon}
                  <Tag
                    color={method.tagColor}
                    style={{ marginLeft: 12, fontSize: 12, padding: '2px 10px', borderRadius: 6, fontWeight: 600 }}
                  >
                    {method.tag}
                  </Tag>
                </div>
                <Title level={4} style={{ marginBottom: 8, fontWeight: 700 }}>
                  {method.title}
                </Title>
                <Paragraph style={{ color: '#64748b', marginBottom: 16, fontSize: 14 }}>
                  {method.description}
                </Paragraph>
                <div
                  style={{
                    background: '#f1f5f9',
                    borderRadius: 10,
                    padding: '12px 16px',
                    fontFamily: "'JetBrains Mono', 'SF Mono', monospace",
                    fontSize: 13,
                    marginBottom: 12,
                    overflow: 'auto',
                    border: '1px solid #e2e8f0',
                  }}
                >
                  <Text code style={{ border: 'none', background: 'transparent', color: '#4f46e5' }}>
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

        <div style={{ marginTop: 40 }}>
          <Card
            style={{
              background: 'linear-gradient(135deg, #eef2ff 0%, #faf5ff 50%, #ecfeff 100%)',
              border: '1px solid #c7d2fe',
              borderRadius: 16,
            }}
            styles={{ body: { padding: 32 } }}
          >
            <Title level={4} style={{ marginBottom: 24, fontWeight: 700 }}>
              💡 Skills + MCP 组合使用效果最佳
            </Title>
            <Steps
              direction="horizontal"
              current={-1}
              items={[
                {
                  title: 'Skills 提供知识',
                  description: '告诉 AI "如何做"：API 参考、代码示例、决策树',
                  icon: <CheckCircleOutlined style={{ color: '#7c3aed', fontSize: 20 }} />,
                },
                {
                  title: 'MCP 提供引擎',
                  description: '让 AI "执行"：21 个结构化 JSON 工具调用',
                  icon: <CheckCircleOutlined style={{ color: '#2563eb', fontSize: 20 }} />,
                },
                {
                  title: '两者配合',
                  description: '知识 + 执行 = 最佳 AI 版本操作体验',
                  icon: <CheckCircleOutlined style={{ color: '#059669', fontSize: 20 }} />,
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
