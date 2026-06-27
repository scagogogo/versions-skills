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
  { icon: <SyncOutlined />, title: '全面的版本号支持', description: '标准语义化版本（1.2.3）、带前缀（v1.2.3）、预发布（1.2.3-beta1）及自定义格式', color: '#2563eb' },
  { icon: <CodeOutlined />, title: '灵活的解析', description: '自动识别前缀、数字部分、后缀和元数据，支持自定义分隔符', color: '#0ea5e9' },
  { icon: <BarChartOutlined />, title: '语义化比较', description: '基于后缀权重排序：dev < alpha < beta < rc < stable', color: '#16a34a' },
  { icon: <AppstoreOutlined />, title: '分组与排序', description: '按主/次版本号分组，支持稳定的预发布版本排序', color: '#ea580c' },
  { icon: <SearchOutlined />, title: '范围查询', description: '支持灵活的边界策略，O(log n) 二分查找性能', color: '#2563eb' },
  { icon: <FilterOutlined />, title: '约束表达式', description: '完整 npm 风格：>=1.0.0、^1.2.3、~1.2、1.x', color: '#0ea5e9' },
  { icon: <SafetyCertificateOutlined />, title: 'Semver 规范', description: 'IsSemver()、ValidateSemver() 严格遵循 SemVer 2.0.0', color: '#16a34a' },
  { icon: <FileTextOutlined />, title: '文件支持', description: '从文件读取/写入版本号列表，支持 # 注释', color: '#ea580c' },
  { icon: <EyeOutlined />, title: '可视化', description: 'Unicode 树形版本层次结构展示', color: '#2563eb' },
  { icon: <ToolOutlined />, title: '不可变操作', description: 'With* 和 Bump* 方法永不修改原始对象', color: '#0ea5e9' },
  { icon: <LinkOutlined />, title: '序列化', description: '内置 JSON、Text、SQL Scanner/Valuer 支持', color: '#16a34a' },
  { icon: <ThunderboltOutlined />, title: '零依赖', description: '核心库无外部依赖，Go 标准库即可编译运行', color: '#ea580c' },
]

const FeaturesSection: React.FC = () => {
  return (
    <div id="features" style={{ padding: '64px 24px', background: '#f8fafc' }}>
      <div style={{ maxWidth: 1120, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">功能特性</div>
          <p className="section-subtitle">12 大核心能力，覆盖版本号操作全场景</p>
        </div>

        <Row gutter={[16, 16]}>
          {features.map((feature, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <Card
                className="flat-card"
                style={{ height: '100%', borderRadius: 4, border: '1px solid #e2e8f0' }}
                styles={{ body: { padding: 20 } }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 10 }}>
                  <span style={{ color: feature.color, fontSize: 18 }}>{feature.icon}</span>
                  <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>{feature.title}</Title>
                </div>
                <Paragraph style={{ color: '#64748b', marginBottom: 0, fontSize: 13, lineHeight: 1.6 }}>
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
