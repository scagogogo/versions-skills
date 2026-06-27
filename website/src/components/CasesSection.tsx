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
    icon: <CloudServerOutlined style={{ fontSize: 28 }} />,
    color: '#4f46e5',
    bg: '#eef2ff',
    title: '依赖版本管理',
    tags: ['SDK', '约束检查'],
    description: '在 CI/CD 中检查项目依赖是否满足版本约束，自动判断是否需要更新依赖或发出安全告警。',
    code: `// 检查所有依赖是否满足最低版本要求
for _, dep := range dependencies {
    v := versions.NewVersion(dep.Version)
    ok, _ := v.Matches(dep.Constraint)
    if !ok {
        report(dep.Name, "不满足约束", dep.Constraint)
    }
}`,
  },
  {
    icon: <BugOutlined style={{ fontSize: 28 }} />,
    color: '#dc2626',
    bg: '#fef2f2',
    title: '安全漏洞范围排查',
    tags: ['CLI', '范围查询'],
    description: '当 CVE 公告披露受影响版本范围时，快速筛选出项目中受影响的版本列表。',
    code: `# 查找受 CVE-2024-XXXX 影响的版本
versions range 2.0.0 2.3.5 \\
  $(cat installed_versions.txt)

# 输出: 2.1.0, 2.2.3, 2.3.1 (在范围内)`,
  },
  {
    icon: <ContainerOutlined style={{ fontSize: 28 }} />,
    color: '#059669',
    bg: '#ecfdf5',
    title: 'Docker 镜像 Tag 解析',
    tags: ['MCP', '解析'],
    description: '在 AI Agent 辅助的运维场景中，解析 Docker 镜像 Tag 判断是否为稳定版本。',
    code: `# 通过 MCP 工具调用
version_parse("nginx:1.25.3-alpine")
→ {
    "prefix": "",
    "major": 1, "minor": 25, "patch": 3,
    "suffix": "alpine",
    "is_prerelease": false
  }`,
  },
  {
    icon: <BranchesOutlined style={{ fontSize: 28 }} />,
    color: '#7c3aed',
    bg: '#f5f3ff',
    title: 'Release 自动排序',
    tags: ['Skills', '排序'],
    description: '让 Claude Code 帮你自动整理和排序 Git Tag 列表，生成版本发布报告。',
    code: `# 在 Claude Code 中使用 Skills
/version-sorting

> 请帮我排序这些 Git Tag:
> v1.0.0, v1.10.0, v1.2.0, v2.0.0-beta, v1.9.0
>
> 排序结果:
> v1.0.0 → v1.2.0 → v1.9.0 →
> v1.10.0 → v2.0.0-beta`,
  },
  {
    icon: <SafetyCertificateOutlined style={{ fontSize: 28 }} />,
    color: '#d97706',
    bg: '#fffbeb',
    title: 'Go Module 兼容性检查',
    tags: ['SDK', 'Semver'],
    description: '在发布 Go Module 新版本前验证版本号是否符合 Semver 规范，确保 go.mod 兼容性。',
    code: `v := versions.NewVersion("v2.0.0-rc1")
versions.ValidateSemver(v)
// → true

v.IsPrerelease()
// → true (不应作为正式发布)`,
  },
  {
    icon: <DashboardOutlined style={{ fontSize: 28 }} />,
    color: '#0891b2',
    bg: '#ecfeff',
    title: '监控面板版本聚合',
    tags: ['SDK', '分组'],
    description: '将服务实例的版本号按主版本分组，在监控面板中展示各版本的实例分布。',
    code: `instances := loadVersionsFromConsul()
groups := versions.GroupByMajor(instances)
// → {
//     1: [1.0.0×3, 1.1.0×5, 1.2.0×2],
//     2: [2.0.0×8]
//   }`,
  },
]

const CasesSection: React.FC = () => {
  return (
    <div id="cases" style={{ padding: '80px 32px', background: '#f8fafc' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">使用案例</div>
          <p className="section-subtitle">真实场景中 Versions-Skills 如何发挥价值</p>
        </div>

        <Row gutter={[20, 24]}>
          {cases.map((c, index) => (
            <Col xs={24} md={12} key={index}>
              <Card
                hoverable
                className="hover-card"
                style={{ height: '100%', borderRadius: 16, border: '1px solid #e2e8f0' }}
                styles={{ body: { padding: 24 } }}
              >
                <div style={{ display: 'flex', alignItems: 'center', marginBottom: 12 }}>
                  <div
                    style={{
                      width: 48,
                      height: 48,
                      borderRadius: 12,
                      background: c.bg,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      color: c.color,
                      marginRight: 14,
                      flexShrink: 0,
                    }}
                  >
                    {c.icon}
                  </div>
                  <div>
                    <Title level={5} style={{ marginBottom: 4, fontWeight: 700 }}>
                      {c.title}
                    </Title>
                    <div>
                      {c.tags.map((tag, i) => (
                        <Tag key={i} color={c.color} style={{ fontSize: 11, borderRadius: 4, margin: '0 4px 0 0' }}>
                          {tag}
                        </Tag>
                      ))}
                    </div>
                  </div>
                </div>
                <Paragraph style={{ color: '#64748b', marginBottom: 14, fontSize: 14 }}>
                  {c.description}
                </Paragraph>
                <pre
                  style={{
                    background: '#1e1e2e',
                    color: '#cdd6f4',
                    padding: 14,
                    borderRadius: 10,
                    fontSize: 12,
                    lineHeight: 1.5,
                    overflow: 'auto',
                    margin: 0,
                    border: '1px solid #313244',
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
