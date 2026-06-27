import { Typography, Row, Col, Card } from 'antd'
import {
  SyncOutlined,
  AppstoreOutlined,
  BarChartOutlined,
  FilterOutlined,
  SearchOutlined,
  SafetyCertificateOutlined,
  FileTextOutlined,
  CodeOutlined,
  ToolOutlined,
  LinkOutlined,
  EyeOutlined,
  ThunderboltOutlined,
} from '@ant-design/icons'

const { Title, Paragraph } = Typography

const features = [
  {
    icon: <SyncOutlined style={{ fontSize: 26 }} />,
    color: '#4f46e5',
    bg: '#eef2ff',
    title: '全面的版本号支持',
    description: '标准语义化版本（1.2.3）、带前缀（v1.2.3）、预发布（1.2.3-beta1）及自定义格式',
  },
  {
    icon: <CodeOutlined style={{ fontSize: 26 }} />,
    color: '#7c3aed',
    bg: '#f5f3ff',
    title: '灵活的解析',
    description: '自动识别前缀、数字部分、后缀和元数据，支持自定义分隔符',
  },
  {
    icon: <BarChartOutlined style={{ fontSize: 26 }} />,
    color: '#0891b2',
    bg: '#ecfeff',
    title: '语义化比较',
    description: '基于后缀权重排序：dev < alpha < beta < rc < stable，精确判定版本先后',
  },
  {
    icon: <AppstoreOutlined style={{ fontSize: 26 }} />,
    color: '#db2777',
    bg: '#fdf2f8',
    title: '分组与排序',
    description: '按主/次版本号分组，支持稳定的预发布版本排序',
  },
  {
    icon: <SearchOutlined style={{ fontSize: 26 }} />,
    color: '#d97706',
    bg: '#fffbeb',
    title: '范围查询',
    description: '支持灵活的边界包含/排除策略，O(log n) 二分查找性能',
  },
  {
    icon: <FilterOutlined style={{ fontSize: 26 }} />,
    color: '#059669',
    bg: '#ecfdf5',
    title: '约束表达式',
    description: '完整 npm 风格约束：>=1.0.0、^1.2.3、~1.2、1.x、>=1.0.0,<2.0.0 || >=3.0.0',
  },
  {
    icon: <SafetyCertificateOutlined style={{ fontSize: 26 }} />,
    color: '#2563eb',
    bg: '#eff6ff',
    title: 'Semver 规范',
    description: 'IsSemver()、ValidateSemver() 严格遵循 SemVer 2.0.0 规范验证',
  },
  {
    icon: <FileTextOutlined style={{ fontSize: 26 }} />,
    color: '#ca8a04',
    bg: '#fefce8',
    title: '文件支持',
    description: '从文件读取/写入版本号列表，支持 # 注释和空行忽略',
  },
  {
    icon: <EyeOutlined style={{ fontSize: 26 }} />,
    color: '#dc2626',
    bg: '#fef2f2',
    title: '可视化',
    description: 'Unicode 树形版本层次结构展示，直观理解版本关系',
  },
  {
    icon: <ToolOutlined style={{ fontSize: 26 }} />,
    color: '#9333ea',
    bg: '#faf5ff',
    title: '不可变操作',
    description: 'With* 和 Bump* 方法永不修改原始对象，安全无副作用',
  },
  {
    icon: <LinkOutlined style={{ fontSize: 26 }} />,
    color: '#4f46e5',
    bg: '#eef2ff',
    title: '序列化',
    description: '内置 JSON、Text、SQL Scanner/Valuer 支持，开箱即用',
  },
  {
    icon: <ThunderboltOutlined style={{ fontSize: 26 }} />,
    color: '#0d9488',
    bg: '#f0fdfa',
    title: '零依赖',
    description: '核心库无外部依赖，Go 标准库即可编译运行',
  },
]

const FeaturesSection: React.FC = () => {
  return (
    <div id="features" style={{ padding: '80px 32px', background: '#f8fafc' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">功能特性</div>
          <p className="section-subtitle">12 大核心能力，覆盖版本号操作全场景</p>
        </div>

        <Row gutter={[20, 20]}>
          {features.map((feature, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <Card
                hoverable
                className="hover-card"
                style={{ height: '100%', borderRadius: 16, border: '1px solid #e2e8f0' }}
                styles={{
                  body: { padding: 24 },
                }}
              >
                <div
                  style={{
                    width: 48,
                    height: 48,
                    borderRadius: 12,
                    background: feature.bg,
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    marginBottom: 16,
                    color: feature.color,
                  }}
                >
                  {feature.icon}
                </div>
                <Title level={5} style={{ marginBottom: 8, fontWeight: 700 }}>
                  {feature.title}
                </Title>
                <Paragraph style={{ color: '#64748b', marginBottom: 0, fontSize: 14, lineHeight: 1.6 }}>
                  {feature.description}
                </Paragraph>
              </Card>
            </Col>
          ))}
        </Row>
      </div>
    </div>
  )
}

export default FeaturesSection
