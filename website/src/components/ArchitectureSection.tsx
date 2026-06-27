import { Typography, Card, Row, Col, Tag } from 'antd'
import {
  RobotOutlined,
  ApiOutlined,
  CodeSandboxOutlined,
  CodeOutlined,
} from '@ant-design/icons'

const { Text } = Typography

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
    <div id="architecture" style={{ padding: '64px 24px', background: '#f8fafc' }}>
      <div style={{ maxWidth: 1120, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">架构设计</div>
          <p className="section-subtitle">四层架构：AI Agent → 接口层 → 功能层 → 核心库</p>
        </div>

        {/* Architecture Diagram - flat style */}
        <Card
          style={{ maxWidth: 720, margin: '0 auto 32px', borderRadius: 4, border: '1px solid #e2e8f0', overflow: 'hidden' }}
          styles={{ body: { padding: 0 } }}
        >
          <div style={{ textAlign: 'center' }}>
            <div style={{ background: '#0f172a', color: '#fff', padding: '14px 20px', fontSize: 15, fontWeight: 600 }}>
              🤖 AI Agent / IDE
              <br />
              <Text style={{ color: 'rgba(255,255,255,0.6)', fontSize: 12, fontWeight: 400 }}>
                Claude Code · Cursor · Windsurf · VS Code Copilot
              </Text>
            </div>
            <Row>
              <Col span={12} style={{ padding: '16px 12px', borderRight: '1px solid #e2e8f0', borderBottom: '1px solid #e2e8f0', background: '#fff' }}>
                <RobotOutlined style={{ fontSize: 20, color: '#2563eb' }} />
                <div style={{ fontWeight: 600, marginTop: 4, fontSize: 13, color: '#2563eb' }}>Skills Plugin</div>
                <Text type="secondary" style={{ fontSize: 11 }}>13 SKILL.md → 斜杠命令</Text>
              </Col>
              <Col span={12} style={{ padding: '16px 12px', borderBottom: '1px solid #e2e8f0', background: '#fff' }}>
                <ApiOutlined style={{ fontSize: 20, color: '#0ea5e9' }} />
                <div style={{ fontWeight: 600, marginTop: 4, fontSize: 13, color: '#0ea5e9' }}>MCP Server</div>
                <Text type="secondary" style={{ fontSize: 11 }}>21 version_* 工具</Text>
              </Col>
            </Row>
            <Row>
              <Col span={12} style={{ padding: '14px 12px', borderRight: '1px solid #e2e8f0', borderBottom: '1px solid #e2e8f0', background: '#fff' }}>
                <CodeOutlined style={{ fontSize: 18, color: '#ea580c' }} />
                <div style={{ fontWeight: 600, fontSize: 13, color: '#ea580c' }}>CLI Binary</div>
                <Text type="secondary" style={{ fontSize: 11 }}>Shell / CI/CD</Text>
              </Col>
              <Col span={12} style={{ padding: '14px 12px', borderBottom: '1px solid #e2e8f0', background: '#fff' }}>
                <CodeSandboxOutlined style={{ fontSize: 18, color: '#16a34a' }} />
                <div style={{ fontWeight: 600, fontSize: 13, color: '#16a34a' }}>Go SDK</div>
                <Text type="secondary" style={{ fontSize: 11 }}>Go 程序</Text>
              </Col>
            </Row>
            <div style={{ background: '#f8fafc', padding: '12px', fontWeight: 600, color: '#475569', fontSize: 13 }}>
              🏗️ Core Library (Go · Zero Dependencies)
            </div>
          </div>
        </Card>

        <div style={{ textAlign: 'center', marginBottom: 20 }}>
          <Tag style={{ fontSize: 13, padding: '4px 14px', borderRadius: 2, fontWeight: 600, color: '#2563eb', borderColor: '#2563eb', background: 'transparent' }}>
            13 Skills 斜杠命令
          </Tag>
        </div>

        <Row gutter={[12, 12]}>
          {skills.map((skill, index) => (
            <Col xs={24} sm={12} md={8} lg={6} key={index}>
              <Card size="small" className="flat-card" style={{ height: '100%', borderRadius: 4, border: '1px solid #e2e8f0' }}>
                <Text code style={{ fontSize: 12, color: '#2563eb', fontWeight: 600 }}>{skill.name}</Text>
                <div style={{ marginTop: 4 }}>
                  <Text type="secondary" style={{ fontSize: 11 }}>{skill.desc}</Text>
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
