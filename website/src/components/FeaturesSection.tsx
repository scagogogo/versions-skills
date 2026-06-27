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
    icon: <SyncOutlined style={{ fontSize: 28, color: '#1677ff' }} />,
    title: '全面的版本号支持',
    description: '标准语义化版本（1.2.3）、带前缀（v1.2.3）、预发布（1.2.3-beta1）及自定义格式',
  },
  {
    icon: <CodeOutlined style={{ fontSize: 28, color: '#722ed1' }} />,
    title: '灵活的解析',
    description: '自动识别前缀、数字部分、后缀和元数据，支持自定义分隔符',
  },
  {
    icon: <BarChartOutlined style={{ fontSize: 28, color: '#13c2c2' }} />,
    title: '语义化比较',
    description: '基于后缀权重排序：dev < alpha < beta < rc < stable，精确判定版本先后',
  },
  {
    icon: <AppstoreOutlined style={{ fontSize: 28, color: '#eb2f96' }} />,
    title: '分组与排序',
    description: '按主/次版本号分组，支持稳定的预发布版本排序',
  },
  {
    icon: <SearchOutlined style={{ fontSize: 28, color: '#fa8c16' }} />,
    title: '范围查询',
    description: '支持灵活的边界包含/排除策略，O(log n) 二分查找性能',
  },
  {
    icon: <FilterOutlined style={{ fontSize: 28, color: '#52c41a' }} />,
    title: '约束表达式',
    description: '完整 npm 风格约束：>=1.0.0、^1.2.3、~1.2、1.x、>=1.0.0,<2.0.0 || >=3.0.0',
  },
  {
    icon: <SafetyCertificateOutlined style={{ fontSize: 28, color: '#2f54eb' }} />,
    title: 'Semver 规范',
    description: 'IsSemver()、ValidateSemver() 严格遵循 SemVer 2.0.0 规范验证',
  },
  {
    icon: <FileTextOutlined style={{ fontSize: 28, color: '#faad14' }} />,
    title: '文件支持',
    description: '从文件读取/写入版本号列表，支持 # 注释和空行忽略',
  },
  {
    icon: <EyeOutlined style={{ fontSize: 28, color: '#f5222d' }} />,
    title: '可视化',
    description: 'Unicode 树形版本层次结构展示，直观理解版本关系',
  },
  {
    icon: <ToolOutlined style={{ fontSize: 28, color: '#9254de' }} />,
    title: '不可变操作',
    description: 'With* 和 Bump* 方法永不修改原始对象，安全无副作用',
  },
  {
    icon: <LinkOutlined style={{ fontSize: 28, color: '#597ef7' }} />,
    title: '序列化',
    description: '内置 JSON、Text、SQL Scanner/Valuer 支持，开箱即用',
  },
  {
    icon: <ThunderboltOutlined style={{ fontSize: 28, color: '#36cfc9' }} />,
    title: '零依赖',
    description: '核心库无外部依赖，Go 标准库即可编译运行',
  },
]

const FeaturesSection: React.FC = () => {
  return (
    <div id="features" style={{ padding: '80px 48px', background: '#fafafa' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <Title level={2} style={{ textAlign: 'center', marginBottom: 8 }}>
          功能特性
        </Title>
        <Paragraph
          style={{ textAlign: 'center', color: '#666', fontSize: 16, marginBottom: 48 }}
        >
          12 大核心能力，覆盖版本号操作全场景
        </Paragraph>

        <Row gutter={[24, 24]}>
          {features.map((feature, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <Card
                hoverable
                style={{ height: '100%' }}
                styles={{
                  body: { padding: 24 },
                }}
              >
                <div style={{ marginBottom: 12 }}>{feature.icon}</div>
                <Title level={5} style={{ marginBottom: 8 }}>
                  {feature.title}
                </Title>
                <Paragraph style={{ color: '#666', marginBottom: 0 }}>
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
