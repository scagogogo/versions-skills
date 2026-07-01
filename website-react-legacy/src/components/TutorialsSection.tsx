import { Typography, Row, Col, Card, Tag, Timeline } from 'antd'
import {
  BookOutlined,
  CodeOutlined,
  FilterOutlined,
  BarChartOutlined,
  CloudOutlined,
  DatabaseOutlined,
} from '@ant-design/icons'

const { Title, Paragraph, Text } = Typography

const tutorials = [
  {
    icon: <BookOutlined style={{ fontSize: 18, color: '#2563eb' }} />,
    title: '版本号解析完全指南',
    description: '从零开始学习版本号的结构和语义：前缀、数字段、后缀权重、元数据。',
    level: '入门',
    levelColor: '#16a34a',
    steps: [
      '理解版本号各部分：Raw、Prefix、VersionNumbers、Suffix、Metadata',
      'NewVersion 自动解析 vs Coerce 从字符串提取',
      '自定义 Parser 选项：设置分隔符和前缀策略',
    ],
  },
  {
    icon: <CodeOutlined style={{ fontSize: 18, color: '#0ea5e9' }} />,
    title: '约束表达式深入实践',
    description: '掌握 npm 风格的版本约束语法，从简单比较到复杂的 OR/AND 组合。',
    level: '进阶',
    levelColor: '#ea580c',
    steps: [
      '单约束操作符：=、>、<、>=、<=、~、^',
      'AND 组合（逗号分隔）和 OR 组合（|| 分隔）',
      'NegateConstraint 反转约束，构建排除策略',
    ],
  },
  {
    icon: <FilterOutlined style={{ fontSize: 18, color: '#16a34a' }} />,
    title: '排序、分组与过滤',
    description: '高效处理版本列表：语义化排序、按主版本号分组、按约束筛选。',
    level: '入门',
    levelColor: '#16a34a',
    steps: [
      'SortVersionSlice 语义化排序 vs 字典序排序',
      'GroupByMajor / GroupByMinor 分组策略',
      'Filter + Constraint 组合构建版本过滤管道',
    ],
  },
  {
    icon: <BarChartOutlined style={{ fontSize: 18, color: '#ea580c' }} />,
    title: '范围查询与性能优化',
    description: '使用 SortedVersionGroups 构建索引，实现 O(log n) 的版本范围查询。',
    level: '高级',
    levelColor: '#dc2626',
    steps: [
      'NewClosedRange / NewOpenRange 创建范围',
      'SortedVersionGroups 预排序索引构建',
      '大规模数据下的性能基准和最佳实践',
    ],
  },
  {
    icon: <CloudOutlined style={{ fontSize: 18, color: '#2563eb' }} />,
    title: 'MCP Server 部署与集成',
    description: '将 Versions-Skills 部署为 MCP Server，接入各类 AI 客户端。',
    level: '进阶',
    levelColor: '#ea580c',
    steps: [
      '安装和配置 versions-mcp',
      'SSE 模式部署（团队共享）',
      '各客户端配置：Claude Code / Cursor / Windsurf',
    ],
  },
  {
    icon: <DatabaseOutlined style={{ fontSize: 18, color: '#0ea5e9' }} />,
    title: '文件 I/O 与数据管道',
    description: '从版本文件读取、写入、增量更新，构建 CI/CD 版本号管理管道。',
    level: '入门',
    levelColor: '#16a34a',
    steps: [
      '版本文件格式规范：一行一个、# 注释',
      'ReadVersionsFromFile / WriteVersionsToFile',
      '与 Shell 管道配合的 CLI 工作流',
    ],
  },
]

const TutorialsSection: React.FC = () => {
  return (
    <div id="tutorials" style={{ padding: '64px 24px', background: '#f8fafc' }}>
      <div style={{ maxWidth: 1120, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">技术教程</div>
          <p className="section-subtitle">从入门到高级，系统学习版本号操作</p>
        </div>

        <Row gutter={[16, 16]}>
          {tutorials.map((t, index) => (
            <Col xs={24} md={12} lg={8} key={index}>
              <Card
                className="flat-card"
                style={{ height: '100%', borderRadius: 4, border: '1px solid #e2e8f0' }}
                styles={{ body: { padding: 20 } }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 10 }}>
                  {t.icon}
                  <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 14 }}>{t.title}</Title>
                  <Tag style={{ fontSize: 11, borderRadius: 2, color: t.levelColor, borderColor: t.levelColor, background: 'transparent', fontWeight: 600, padding: '0 6px' }}>
                    {t.level}
                  </Tag>
                </div>
                <Paragraph style={{ color: '#64748b', marginBottom: 12, fontSize: 13, lineHeight: 1.5 }}>{t.description}</Paragraph>
                <Timeline
                  items={t.steps.map((step) => ({
                    children: <Text style={{ fontSize: 12, color: '#475569' }}>{step}</Text>,
                  }))}
                  style={{ marginTop: 0, paddingTop: 0 }}
                />
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </div>
  )
}

export default TutorialsSection
