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
    icon: <RobotOutlined style={{ fontSize: 28, color: '#2563eb' }} />,
    tag: '推荐',
    tagColor: '#2563eb',
    title: 'Skills（Claude Code）',
    description: 'Claude Code 读取 SKILL.md 作为领域知识，通过斜杠命令直接调用。适合引导式工作流和一次性任务。',
    install: 'claude plugin install versions',
    detail: '13 个斜杠命令：/version-parsing、/version-comparison 等',
  },
  {
    icon: <ApiOutlined style={{ fontSize: 28, color: '#0ea5e9' }} />,
    tag: '通用',
    tagColor: '#0ea5e9',
    title: 'MCP Server',
    description: '任何 MCP 兼容客户端直接调用 version_* 工具。适合编程式调用、批量操作和非 Claude Agent。',
    install: 'go install github.com/scagogogo/versions-skills/cmd/versions-mcp@latest',
    detail: '21 个 AI 可调用工具，兼容 Claude Code / Cursor / Windsurf / VS Code Copilot',
  },
  {
    icon: <CodeSandboxOutlined style={{ fontSize: 28, color: '#16a34a' }} />,
    tag: 'Go',
    tagColor: '#16a34a',
    title: 'Go SDK',
    description: '直接在 Go 代码中使用完整的版本操作 API。适合集成到 Go 项目中。',
    install: 'go get github.com/scagogogo/versions-skills',
    detail: '零依赖核心库，完整的 Version 类型和方法链',
  },
  {
    icon: <CodeOutlined style={{ fontSize: 28, color: '#ea580c' }} />,
    tag: 'CLI',
    tagColor: '#ea580c',
    title: 'CLI',
    description: '命令行工具，40+ 子命令覆盖全部功能。适合 Shell 脚本和 CI/CD 管道。',
    install: 'curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash',
    detail: '6 种 OS × 12 种架构，提供 deb/rpm/apk 包',
  },
]

const AccessSection: React.FC = () => {
  return (
    <div id="access" style={{ padding: '64px 24px', background: '#fff' }}>
      <div style={{ maxWidth: 1120, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">四种接入方式</div>
          <p className="section-subtitle">根据场景选择最合适的方式，也可以组合使用获得最佳体验</p>
        </div>

        <Row gutter={[16, 16]}>
          {accessMethods.map((method, index) => (
            <Col xs={24} sm={12} key={index}>
              <Card
                className="flat-card"
                style={{ height: '100%', borderRadius: 4, border: '1px solid #e2e8f0' }}
                styles={{ body: { padding: 20 } }}
              >
                <div style={{ display: 'flex', alignItems: 'center', marginBottom: 12 }}>
                  {method.icon}
                  <Tag
                    style={{ marginLeft: 10, fontSize: 12, padding: '1px 8px', borderRadius: 2, fontWeight: 600, color: method.tagColor, borderColor: method.tagColor, background: 'transparent' }}
                  >
                    {method.tag}
                  </Tag>
                </div>
                <Title level={5} style={{ marginBottom: 6, fontWeight: 600 }}>{method.title}</Title>
                <Paragraph style={{ color: '#64748b', marginBottom: 12, fontSize: 13 }}>{method.description}</Paragraph>
                <div
                  style={{
                    background: '#f8fafc',
                    borderRadius: 4,
                    padding: '8px 12px',
                    fontFamily: "'JetBrains Mono', monospace",
                    fontSize: 12,
                    marginBottom: 10,
                    overflow: 'auto',
                    border: '1px solid #e2e8f0',
                  }}
                >
                  <Text code style={{ border: 'none', background: 'transparent', color: '#2563eb' }}>
                    {method.install}
                  </Text>
                </div>
                <Text type="secondary" style={{ fontSize: 12 }}>{method.detail}</Text>
              </Card>
            </Col>
          ))}
        </Row>

        <div style={{ marginTop: 32 }}>
          <Card style={{ background: '#f8fafc', border: '1px solid #e2e8f0', borderRadius: 4 }} styles={{ body: { padding: 20 } }}>
            <Title level={5} style={{ marginBottom: 16, fontWeight: 600 }}>Skills + MCP 组合使用效果最佳</Title>
            <Steps
              direction="horizontal"
              current={-1}
              items={[
                { title: 'Skills 提供知识', description: '告诉 AI "如何做"', icon: <CheckCircleOutlined style={{ color: '#2563eb' }} /> },
                { title: 'MCP 提供引擎', description: '让 AI "执行"操作', icon: <CheckCircleOutlined style={{ color: '#0ea5e9' }} /> },
                { title: '两者配合', description: '知识 + 执行', icon: <CheckCircleOutlined style={{ color: '#16a34a' }} /> },
              ]}
            />
          </Card>
        </div>
      </div>
    </div>
  )
}

export default AccessSection
