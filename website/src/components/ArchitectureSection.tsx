import { Typography, Card, Row, Col, Tag } from 'antd'
import {
  RobotOutlined,
  ApiOutlined,
  CodeSandboxOutlined,
  CodeOutlined,
} from '@ant-design/icons'

const { Text } = Typography

const skills = [
  { name: '/version-parsing', desc: '解析、验证、提取版本号组件', color: '#4f46e5' },
  { name: '/version-comparison', desc: '比较版本号，检查排序关系', color: '#7c3aed' },
  { name: '/version-sorting', desc: '升序/降序排序版本号列表', color: '#0891b2' },
  { name: '/version-grouping', desc: '按主/次版本号分组', color: '#db2777' },
  { name: '/version-constraints', desc: '解析和检查约束表达式', color: '#d97706' },
  { name: '/version-range-query', desc: '查询范围内的版本号', color: '#059669' },
  { name: '/version-visualization', desc: '树形版本层次结构展示', color: '#dc2626' },
  { name: '/version-file-operations', desc: '读写版本号列表文件', color: '#ca8a04' },
  { name: '/version-check', desc: '布尔类型检查（IsBeta、IsStable 等）', color: '#2563eb' },
  { name: '/version-mutation', desc: '版本号 Bump，不可变修改', color: '#9333ea' },
  { name: '/version-properties', desc: '访问版本号段落、后缀权重、前缀', color: '#4f46e5' },
  { name: '/cli-operations', desc: '完整 CLI 命令参考', color: '#0d9488' },
  { name: '/mcp-operations', desc: 'MCP 服务器设置与工具参考', color: '#06b6d4' },
]

const ArchitectureSection: React.FC = () => {
  return (
    <div id="architecture" style={{ padding: '80px 32px', background: '#ffffff' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">架构设计</div>
          <p className="section-subtitle">四层架构：AI Agent → 接口层 → 功能层 → 核心库</p>
        </div>

        {/* Architecture Diagram */}
        <Card
          style={{
            maxWidth: 800,
            margin: '0 auto 48px',
            borderRadius: 20,
            border: '1px solid #e2e8f0',
            overflow: 'hidden',
          }}
          styles={{ body: { padding: 0 } }}
        >
          <div style={{ textAlign: 'center' }}>
            {/* Layer 1: AI Agent */}
            <div
              style={{
                background: 'linear-gradient(135deg, #4f46e5 0%, #7c3aed 100%)',
                color: '#fff',
                padding: '20px 24px',
                fontSize: 17,
                fontWeight: 700,
                letterSpacing: '-0.01em',
              }}
            >
              🤖 AI Agent / IDE
              <br />
              <Text style={{ color: 'rgba(255,255,255,0.75)', fontSize: 13, fontWeight: 400 }}>
                Claude Code · Cursor · Windsurf · VS Code Copilot
              </Text>
            </div>

            {/* Layer 2: Interface */}
            <Row>
              <Col span={12}
                style={{
                  padding: '24px 16px',
                  borderRight: '1px solid #e2e8f0',
                  borderBottom: '1px solid #e2e8f0',
                  background: '#faf5ff',
                }}
              >
                <RobotOutlined style={{ fontSize: 24, color: '#7c3aed' }} />
                <div style={{ fontWeight: 700, marginTop: 6, color: '#7c3aed' }}>Skills Plugin</div>
                <Text type="secondary" style={{ fontSize: 12 }}>13 SKILL.md → 斜杠命令</Text>
              </Col>
              <Col span={12} style={{ padding: '24px 16px', borderBottom: '1px solid #e2e8f0', background: '#eff6ff' }}>
                <ApiOutlined style={{ fontSize: 24, color: '#2563eb' }} />
                <div style={{ fontWeight: 700, marginTop: 6, color: '#2563eb' }}>MCP Server</div>
                <Text type="secondary" style={{ fontSize: 12 }}>21 version_* 工具</Text>
              </Col>
            </Row>

            {/* Layer 3: Feature */}
            <Row>
              <Col span={12}
                style={{
                  padding: '20px 16px',
                  borderRight: '1px solid #e2e8f0',
                  borderBottom: '1px solid #e2e8f0',
                  background: '#fffbeb',
                }}
              >
                <CodeOutlined style={{ fontSize: 20, color: '#d97706' }} />
                <div style={{ fontWeight: 700, color: '#d97706' }}>CLI Binary</div>
                <Text type="secondary" style={{ fontSize: 12 }}>Shell / CI/CD</Text>
              </Col>
              <Col span={12} style={{ padding: '20px 16px', borderBottom: '1px solid #e2e8f0', background: '#ecfdf5' }}>
                <CodeSandboxOutlined style={{ fontSize: 20, color: '#059669' }} />
                <div style={{ fontWeight: 700, color: '#059669' }}>Go SDK</div>
                <Text type="secondary" style={{ fontSize: 12 }}>Go 程序</Text>
              </Col>
            </Row>

            {/* Layer 4: Core */}
            <div
              style={{
                background: '#f0fdfa',
                padding: '16px',
                fontWeight: 700,
                color: '#0d9488',
              }}
            >
              🏗️ Core Library (Go · Zero Dependencies)
            </div>
          </div>
        </Card>

        {/* Skills Grid */}
        <div style={{ textAlign: 'center', marginBottom: 32 }}>
          <Tag color="purple" style={{ fontSize: 15, padding: '6px 20px', borderRadius: 10, fontWeight: 600 }}>
            🤖 13 Skills 斜杠命令
          </Tag>
        </div>

        <Row gutter={[16, 16]}>
          {skills.map((skill, index) => (
            <Col xs={24} sm={12} md={8} lg={6} key={index}>
              <Card
                size="small"
                hoverable
                className="hover-card"
                style={{ height: '100%', borderRadius: 12, border: '1px solid #e2e8f0' }}
              >
                <Text code style={{ fontSize: 13, color: skill.color, fontWeight: 600 }}>
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
