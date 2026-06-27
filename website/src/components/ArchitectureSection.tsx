import { Typography, Card, Row, Col, Tag, Divider } from 'antd'
import {
  RobotOutlined,
  ApiOutlined,
  CodeSandboxOutlined,
  CodeOutlined,
} from '@ant-design/icons'

const { Title, Paragraph, Text } = Typography

const skills = [
  { name: '/version-parsing', desc: '解析、验证、提取版本号组件' },
  { name: '/version-comparison', desc: '比较版本号，检查排序关系' },
  { name: '/version-sorting', desc: '升序/降序排序版本号列表' },
  { name: '/version-grouping', desc: '按主/次版本号分组' },
  { name: '/version-constraints', desc: '解析和检查约束表达式' },
  { name: '/version-range-query', desc: '查询范围内的版本号' },
  { name: '/version-visualization', desc: '树形版本层次结构展示' },
  { name: '/version-file-operations', desc: '读写版本号列表文件' },
  { name: '/version-check', desc: '布尔类型检查（IsBeta、IsStable 等）' },
  { name: '/version-mutation', desc: '版本号 Bump，不可变修改' },
  { name: '/version-properties', desc: '访问版本号段落、后缀权重、前缀' },
  { name: '/cli-operations', desc: '完整 CLI 命令参考' },
  { name: '/mcp-operations', desc: 'MCP 服务器设置与工具参考' },
]

const ArchitectureSection: React.FC = () => {
  return (
    <div id="architecture" style={{ padding: '80px 48px', background: '#fff' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <Title level={2} style={{ textAlign: 'center', marginBottom: 8 }}>
          架构设计
        </Title>
        <Paragraph
          style={{ textAlign: 'center', color: '#666', fontSize: 16, marginBottom: 48 }}
        >
          四层架构：AI Agent → 接口层 → 功能层 → 核心库
        </Paragraph>

        {/* Architecture Diagram */}
        <Card
          style={{
            maxWidth: 800,
            margin: '0 auto 48px',
            background: '#fafafa',
          }}
        >
          <div style={{ textAlign: 'center' }}>
            {/* Layer 1: AI Agent */}
            <div
              style={{
                background: 'linear-gradient(135deg, #722ed1 0%, #1677ff 100%)',
                color: '#fff',
                padding: '16px 24px',
                borderRadius: '8px 8px 0 0',
                fontSize: 16,
                fontWeight: 600,
              }}
            >
              🤖 AI Agent / IDE
              <br />
              <Text style={{ color: 'rgba(255,255,255,0.8)', fontSize: 13 }}>
                Claude Code · Cursor · Windsurf · VS Code Copilot
              </Text>
            </div>

            {/* Layer 2: Interface */}
            <Row style={{ background: '#e6f7ff' }}>
              <Col span={12}
                style={{
                  padding: '20px 16px',
                  borderRight: '1px solid #91d5ff',
                }}
              >
                <RobotOutlined style={{ fontSize: 24, color: '#722ed1' }} />
                <div style={{ fontWeight: 600, marginTop: 4 }}>Skills Plugin</div>
                <Text type="secondary" style={{ fontSize: 12 }}>13 SKILL.md → 斜杠命令</Text>
              </Col>
              <Col span={12} style={{ padding: '20px 16px' }}>
                <ApiOutlined style={{ fontSize: 24, color: '#1677ff' }} />
                <div style={{ fontWeight: 600, marginTop: 4 }}>MCP Server</div>
                <Text type="secondary" style={{ fontSize: 12 }}>21 version_* 工具</Text>
              </Col>
            </Row>

            {/* Layer 3: Feature */}
            <Row style={{ background: '#f6ffed' }}>
              <Col span={12}
                style={{
                  padding: '16px',
                  borderRight: '1px solid #b7eb8f',
                }}
              >
                <CodeOutlined style={{ fontSize: 20, color: '#fa8c16' }} />
                <div style={{ fontWeight: 600 }}>CLI Binary</div>
                <Text type="secondary" style={{ fontSize: 12 }}>Shell / CI/CD</Text>
              </Col>
              <Col span={12} style={{ padding: '16px' }}>
                <CodeSandboxOutlined style={{ fontSize: 20, color: '#52c41a' }} />
                <div style={{ fontWeight: 600 }}>Go SDK</div>
                <Text type="secondary" style={{ fontSize: 12 }}>Go 程序</Text>
              </Col>
            </Row>

            {/* Layer 4: Core */}
            <div
              style={{
                background: '#fff7e6',
                padding: '16px',
                borderRadius: '0 0 8px 8px',
                fontWeight: 600,
              }}
            >
              🏗️ Core Library (Go · Zero Dependencies)
            </div>
          </div>
        </Card>

        {/* Skills Grid */}
        <Divider>
          <Tag color="purple" style={{ fontSize: 14, padding: '4px 16px' }}>
            🤖 13 Skills 斜杠命令
          </Tag>
        </Divider>

        <Row gutter={[16, 16]} style={{ marginTop: 24 }}>
          {skills.map((skill, index) => (
            <Col xs={24} sm={12} md={8} lg={6} key={index}>
              <Card
                size="small"
                hoverable
                style={{ height: '100%' }}
              >
                <Text code style={{ fontSize: 13, color: '#722ed1' }}>
                  {skill.name}
                </Text>
                <div style={{ marginTop: 6 }}>
                  <Text type="secondary" style={{ fontSize: 12 }}>
                    {skill.desc}
                  </Text>
                </div>
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </div>
  )
}

export default ArchitectureSection
