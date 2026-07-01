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
    icon: <RobotOutlined style={{ fontSize: 18, color: '#2563eb' }} />,
    skills: true,
    mcp: true,
    configPath: '~/.claude/settings.json',
    config: `# 1. Skills 插件（13 个斜杠命令）：
claude plugin marketplace add https://github.com/scagogogo/versions-skills
claude plugin install versions

# 2. MCP Server（直连工具执行）：
claude mcp add versions -- versions-mcp --transport stdio

# 或写入 ~/.claude/settings.json：
# { "mcpServers": { "versions": {
#     "command": "versions-mcp",
#     "args": ["--transport", "stdio"] } } }`,
    note: '同时支持 Skills Plugin 和 MCP Server，推荐两者都装',
  },
  {
    name: 'Cursor',
    icon: <CodeOutlined style={{ fontSize: 18, color: '#0ea5e9' }} />,
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
    icon: <ThunderboltOutlined style={{ fontSize: 18, color: '#16a34a' }} />,
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
    icon: <SettingOutlined style={{ fontSize: 18, color: '#ea580c' }} />,
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
    note: '注意字段名是 "servers"',
  },
  {
    name: 'Codex',
    icon: <RobotOutlined style={{ fontSize: 18, color: '#475569' }} />,
    skills: false,
    mcp: true,
    configPath: '~/.codex/config.toml',
    config: `# 命令式注册（推荐）：
codex mcp add versions -- versions-mcp --transport stdio

# 或直接编辑 ~/.codex/config.toml：
[mcp_servers.versions]
command = "versions-mcp"
args = ["--transport", "stdio"]`,
    note: 'OpenAI Codex CLI，TOML 配置格式',
  },
  {
    name: 'Cline',
    icon: <CodeOutlined style={{ fontSize: 18, color: '#6750a4' }} />,
    skills: false,
    mcp: true,
    configPath: '~/.cline/cline_mcp_settings.json',
    config: `{
  "mcpServers": {
    "versions": {
      "command": "versions-mcp",
      "args": ["--transport", "stdio"],
      "disabled": false,
      "autoApprove": []
    }
  }
}`,
    note: 'VS Code 上的 Cline 插件，含 disabled/autoApprove 字段',
  },
]

const AiIntegrationSection: React.FC = () => {
  return (
    <div id="ai-integration" style={{ padding: '64px 24px', background: '#fff' }}>
      <div style={{ maxWidth: 1120, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">AI Agent 集成指南</div>
          <p className="section-subtitle">一键接入主流 AI 编程工具，让版本号操作成为 AI 的原生能力</p>
        </div>

        <Card className="flat-card" style={{ marginBottom: 16, borderRadius: 4, border: '1px solid #bfdbfe', background: '#eff6ff' }} styles={{ body: { padding: 16 } }}>
          <div style={{ display: 'flex', alignItems: 'flex-start', gap: 10 }}>
            <CheckCircleOutlined style={{ color: '#2563eb', fontSize: 16, marginTop: 2 }} />
            <Text style={{ fontSize: 13, color: '#1e3a8a', lineHeight: 1.7 }}>
              <strong>选型建议：</strong>在 Claude Code 上同时安装 Skills 插件与 MCP Server —— 插件提供版本号领域知识（何时用约束、后缀如何排序），MCP 提供快速结构化执行通道；其它 Agent 仅需 MCP Server 一条路径即可。下方为各客户端的配置方式。
            </Text>
          </div>
        </Card>

        <Row gutter={[16, 16]}>
          {aiClients.map((client, index) => (
            <Col xs={24} sm={12} key={index}>
              <Card
                className="flat-card"
                style={{ height: '100%', borderRadius: 4, border: '1px solid #e2e8f0' }}
                styles={{ body: { padding: 20 } }}
              >
                <div style={{ display: 'flex', alignItems: 'center', marginBottom: 10 }}>
                  {client.icon}
                  <Title level={5} style={{ margin: '0 0 0 8', fontWeight: 600, fontSize: 15 }}>{client.name}</Title>
                  {client.skills && <Tag color="blue" style={{ marginLeft: 6, fontSize: 11, borderRadius: 2 }}>Skills</Tag>}
                  {client.mcp && <Tag color="cyan" style={{ marginLeft: 4, fontSize: 11, borderRadius: 2 }}>MCP</Tag>}
                </div>

                <Text type="secondary" style={{ fontSize: 12, marginBottom: 6 }}>
                  配置文件：<Text code style={{ fontSize: 11 }}>{client.configPath}</Text>
                </Text>

                <pre
                  style={{
                    background: '#1e293b',
                    color: '#e2e8f0',
                    padding: 12,
                    borderRadius: 4,
                    fontSize: 12,
                    lineHeight: 1.5,
                    overflow: 'auto',
                    margin: '6px 0 8',
                  }}
                >
                  <code>{client.config}</code>
                </pre>

                <div style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
                  <CheckCircleOutlined style={{ color: '#16a34a', fontSize: 13 }} />
                  <Text style={{ fontSize: 12, color: '#64748b' }}>{client.note}</Text>
                </div>
              </Card>
            </Col>
          ))}
        </Row>

        <Card style={{ marginTop: 16, borderRadius: 4, border: '1px solid #e2e8f0' }} styles={{ body: { padding: 16 } }}>
          <Title level={5} style={{ fontWeight: 600, marginBottom: 8, fontSize: 14 }}>SSE 网络模式（团队共享部署）</Title>
          <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 8 }}>
            将 MCP Server 部署为 SSE 服务，团队成员通过网络访问同一个实例。
          </Paragraph>
          <pre style={{ background: '#1e293b', color: '#e2e8f0', padding: 10, borderRadius: 4, fontSize: 12, overflow: 'auto', margin: 0 }}>
            <code>versions-mcp --transport sse --port 8080{'\n'}# 客户端指向: http://localhost:8080/sse</code>
          </pre>
        </Card>
      </div>
    </div>
  )
}

export default AiIntegrationSection
