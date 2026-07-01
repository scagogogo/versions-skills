import { Typography, Row, Col, Card, Tag } from 'antd'
import {
  CloudServerOutlined,
  BugOutlined,
  ContainerOutlined,
  BranchesOutlined,
  SafetyCertificateOutlined,
  DashboardOutlined,
} from '@ant-design/icons'

const { Title, Paragraph } = Typography

const cases = [
  {
    icon: <CloudServerOutlined />,
    color: '#2563eb',
    title: '依赖版本管理',
    tags: ['SDK', '约束检查'],
    description: '在 CI/CD 中检查项目依赖是否满足版本约束，自动判断是否需要更新依赖或发出安全告警。',
    code: `for _, dep := range dependencies {
    v := versions.NewVersion(dep.Version)
    ok, _ := v.Matches(dep.Constraint)
    if !ok {
        report(dep.Name, "不满足约束")
    }
}`,
  },
  {
    icon: <BugOutlined />,
    color: '#dc2626',
    title: '安全漏洞范围排查',
    tags: ['CLI', '范围查询'],
    description: '当 CVE 公告披露受影响版本范围时，快速筛选出项目中受影响的版本列表。',
    code: `versions range 2.0.0 2.3.5 \\
  $(cat installed_versions.txt)`,
  },
  {
    icon: <ContainerOutlined />,
    color: '#16a34a',
    title: 'Docker 镜像 Tag 解析',
    tags: ['MCP', '解析'],
    description: '在 AI Agent 辅助的运维场景中，解析 Docker 镜像 Tag 判断是否为稳定版本。',
    code: `version_parse("nginx:1.25.3-alpine")
→ { major: 1, minor: 25, patch: 3,
    suffix: "alpine", is_prerelease: false }`,
  },
  {
    icon: <BranchesOutlined />,
    color: '#0ea5e9',
    title: 'Release 自动排序',
    tags: ['Skills', '排序'],
    description: '让 Claude Code 帮你自动整理和排序 Git Tag 列表，生成版本发布报告。',
    code: `/version-sorting

> v1.0.0, v1.10.0, v1.2.0, v2.0.0-beta
> 排序: v1.0.0 → v1.2.0 → v1.10.0 → v2.0.0-beta`,
  },
  {
    icon: <SafetyCertificateOutlined />,
    color: '#ea580c',
    title: 'Go Module 兼容性检查',
    tags: ['SDK', 'Semver'],
    description: '在发布新版本前验证版本号是否符合 Semver 规范，确保 go.mod 兼容性。',
    code: `v := versions.NewVersion("v2.0.0-rc1")
versions.ValidateSemver(v) // true
v.IsPrerelease()           // true`,
  },
  {
    icon: <DashboardOutlined />,
    color: '#2563eb',
    title: '监控面板版本聚合',
    tags: ['SDK', '分组'],
    description: '将服务实例的版本号按主版本分组，在监控面板中展示各版本的实例分布。',
    code: `groups := versions.GroupByMajor(instances)
// → {1: [1.0.0×3, 1.1.0×5],
//     2: [2.0.0×8]}`,
  },
]

const CasesSection: React.FC = () => {
  return (
    <div id="cases" style={{ padding: '64px 24px', background: '#fff' }}>
      <div style={{ maxWidth: 1120, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">使用案例</div>
          <p className="section-subtitle">真实场景中 Versions-Skills 如何发挥价值</p>
        </div>

        <Row gutter={[16, 16]}>
          {cases.map((c, index) => (
            <Col xs={24} md={12} key={index}>
              <Card
                className="flat-card"
                style={{ height: '100%', borderRadius: 4, border: '1px solid #e2e8f0' }}
                styles={{ body: { padding: 20 } }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 10 }}>
                  <span style={{ color: c.color, fontSize: 18 }}>{c.icon}</span>
                  <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>{c.title}</Title>
                  {c.tags.map((tag, i) => (
                    <Tag key={i} style={{ fontSize: 11, borderRadius: 2, margin: 0, color: c.color, borderColor: c.color, background: 'transparent' }}>
                      {tag}
                    </Tag>
                  ))}
                </div>
                <Paragraph style={{ color: '#64748b', marginBottom: 10, fontSize: 13 }}>{c.description}</Paragraph>
                <pre
                  style={{
                    background: '#1e293b',
                    color: '#e2e8f0',
                    padding: 12,
                    borderRadius: 4,
                    fontSize: 12,
                    lineHeight: 1.5,
                    overflow: 'auto',
                    margin: 0,
                  }}
                >
                  <code>{c.code}</code>
                </pre>
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </div>
  )
}

export default CasesSection
