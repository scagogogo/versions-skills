import { Typography, Row, Col, Card, Tag } from 'antd'
import {
  RobotOutlined,
  CodeOutlined,
  SettingOutlined,
  CheckCircleOutlined,
  ThunderboltOutlined,
} from '@ant-design/icons'

const { Title, Paragraph, Text } = Typography

const aiClients = [
  {
    name: 'Claude Code',
    icon: <RobotOutlined style={{ fontSize: 22, color: '#7c3aed' }} />,
    skills: true,
    mcp: true,
    configPath: '~/.claude/settings.json',
    config: `{
  "mcpServers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}`,
    note: '同时支持 Skills Plugin 和 MCP Server，推荐两者都安装',
  },
  {
    name: 'Cursor',
    icon: <CodeOutlined style={{ fontSize: 22, color: '#2563eb' }} />,
    skills: false,
    mcp: true,
    configPath: '.cursor/mcp.json',
    config: `{
  "mcpServers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}`,
    note: '仅支持 MCP Server，项目级配置',
  },
  {
    name: 'Windsurf',
    icon: <ThunderboltOutlined style={{ fontSize: 22, color: '#0891b2' }} />,
    skills: false,
    mcp: true,
    configPath: '.windsurf/mcp.json',
    config: `{
  "mcpServers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}`,
    note: '仅支持 MCP Server，项目级配置',
  },
  {
    name: 'VS Code Copilot',
    icon: <SettingOutlined style={{ fontSize: 22, color: '#059669' }} />,
    skills: false,
    mcp: true,
    configPath: '.vscode/mcp.json',
    config: `{
  "servers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}`,
    note: '仅支持 MCP Server，注意字段名是 "servers"',
  },
]

const AiIntegrationSection: React.FC = () => {
  return (
    <div id="ai-integration" style={{ padding: '80px 32px', background: '#faf5ff' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">AI Agent 集成指南</div>
          <p className="section-subtitle">一键接入主流 AI 编程工具，让版本号操作成为 AI 的原生能力</p>
        </div>

        {/* Overview */}
        <Card
          style={{
            marginBottom: 32,
            borderRadius: 16,
            background: 'linear-gradient(135deg, #eef2ff 0%, #faf5ff 100%)',
            border: '1px solid #c7d2fe',
          }}
          styles={{ body: { padding: 24 } }}
        >
          <Row gutter={24} align="middle">
            <Col xs={24} md={8}>
              <div style={{ textAlign: 'center' }}>
                <RobotOutlined style={{ fontSize: 48, color: '#7c3aed' }} />
                <Title level={4} style={{ marginTop: 12, fontWeight: 700 }}>
                  双路径接入
                </Title>
              </div>
            </Col>
            <Col xs={24} md={16}>
              <Row gutter={16}>
                <Col span={12}>
                  <Card
                    size="small"
                    style={{ borderRadius: 12, background: '#fff', border: '1px solid #e2e8f0' }}
                    styles={{ body: { padding: 16 } }}
                  >
                    <Tag color="purple" style={{ fontSize: 13, marginBottom: 8, fontWeight: 600 }}>
                      Skills 插件
                    </Tag>
                    <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 0 }}>
                      Claude Code 读取 13 个 SKILL.md 文件作为领域知识，通过斜杠命令直接调用。适合引导式工作流和一次性任务。
                    </Paragraph>
                  </Card>
                </Col>
                <Col span={12}>
                  <Card
                    size="small"
                    style={{ borderRadius: 12, background: '#fff', border: '1px solid #e2e8f0' }}
                    styles={{ body: { padding: 16 } }}
                  >
                    <Tag color="blue" style={{ fontSize: 13, marginBottom: 8, fontWeight: 600 }}>
                      MCP Server
                    </Tag>
                    <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 0 }}>
                      任何 MCP 兼容客户端调用 21 个 version_* 工具。适合编程式调用、批量操作、非 Claude Agent。
                    </Paragraph>
                  </Card>
                </Col>
              </Row>
            </Col>
          </Row>
        </Card>

        {/* Client Configuration Cards */}
        <Row gutter={[20, 24]}>
          {aiClients.map((client, index) => (
            <Col xs={24} sm={12} key={index}>
              <Card
                hoverable
                className="hover-card"
                style={{ height: '100%', borderRadius: 16, border: '1px solid #e2e8f0' }}
                styles={{ body: { padding: 24 } }}
              >
                <div style={{ display: 'flex', alignItems: 'center', marginBottom: 14 }}>
                  {client.icon}
                  <Title level={5} style={{ margin: '0 0 0 10', fontWeight: 700 }}>
                    {client.name}
                  </Title>
                  {client.skills && (
                    <Tag color="purple" style={{ marginLeft: 8, fontSize: 11, borderRadius: 4 }}>
                      Skills
                    </Tag>
                  )}
                  {client.mcp && (
                    <Tag color="blue" style={{ marginLeft: 4, fontSize: 11, borderRadius: 4 }}>
                      MCP
                    </Tag>
                  )}
                </div>

                <Text type="secondary" style={{ fontSize: 12, marginBottom: 8 }}>
                  配置文件：<Text code style={{ fontSize: 12 }}>{client.configPath}</Text>
                </Text>

                <pre
                  style={{
                    background: '#1e1e2e',
                    color: '#cdd6f4',
                    padding: 14,
                    borderRadius: 10,
                    fontSize: 12,
                    lineHeight: 1.5,
                    overflow: 'auto',
                    margin: '8px 0 12',
                    border: '1px solid #313244',
                  }}
                >
                  <code>{client.config}</code>
                </pre>

                <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
                  <CheckCircleOutlined style={{ color: '#059669', fontSize: 14 }} />
                  <Text style={{ fontSize: 12, color: '#64748b' }}>{client.note}</Text>
                </div>
              </Card>
            </Col>
          ))}
        </Row>

        {/* SSE Network Mode */}
        <Card
          style={{ marginTop: 24, borderRadius: 16, background: '#fff', border: '1px solid #e2e8f0' }}
          styles={{ body: { padding: 24 } }}
        >
          <Title level={5} style={{ fontWeight: 700, marginBottom: 12 }}>
            🌐 SSE 网络模式（团队共享部署）
          </Title>
          <Paragraph style={{ color: '#64748b', fontSize: 14, marginBottom: 16 }}>
            对于团队共享场景，可以将 MCP Server 部署为 SSE 服务，所有团队成员通过网络访问同一个版本号工具实例。
          </Paragraph>
          <pre
            style={{
              background: '#1e1e2e',
              color: '#cdd6f4',
              padding: 14,
              borderRadius: 10,
              fontSize: 13,
              overflow: 'auto',
              margin: 0,
              border: '1px solid #313244',
            }}
          >
            <code>versions-mcp --transport sse --port 8080{'\n'}{'\n'}# 然后在客户端配置中指向: http://localhost:8080/sse</code>
          </pre>
        </Card>
      </div>
    </div>
  )
}

export default AiIntegrationSection
