import { Typography, Tabs, Card, Row, Col } from 'antd'
import { CodeOutlined, CodeSandboxOutlined } from '@ant-design/icons'

const { Text } = Typography

const sdkExamples = [
  {
    title: '解析与比较',
    code: `import "github.com/scagogogo/versions-skills"

v1 := versions.NewVersion("1.2.3")
v2 := versions.NewVersion("v1.3.0-beta")

v1.IsOlderThan(v2)      // true
v2.IsPrerelease()       // true
v2.PreReleaseType()     // "beta"
v1.Diff(v2).IsUpgrade() // true`,
  },
  {
    title: '排序与分组',
    code: `list := versions.NewVersions("2.0.0", "1.0.0", "1.10.0", "1.5.0-beta")

// Sort
sorted := versions.SortVersionSlice(list)
// → [1.0.0, 1.5.0-beta, 1.10.0, 2.0.0]

// Group by major version
groups := versions.GroupByMajor(list)
// → {1: [1.0.0, 1.5.0-beta, 1.10.0], 2: [2.0.0]}`,
  },
  {
    title: '约束检查',
    code: `v := versions.NewVersion("1.5.0")

// Check single constraint
c, _ := versions.ParseConstraint(">=1.0.0")
v.Satisfies(c)  // true

// Check constraint expression
ok, _ := v.Matches(">=1.0.0,<2.0.0")  // true

// Negate a constraint
neg := versions.NegateConstraint(c)  // <1.0.0`,
  },
  {
    title: '范围查询',
    code: `low := versions.NewVersion("1.0.0")
high := versions.NewVersion("2.0.0")

r := versions.NewClosedRange(low, high)
r.Contains(versions.NewVersion("1.5.0"))  // true
r.Contains(versions.NewVersion("2.1.0"))  // false`,
  },
]

const cliExamples = [
  {
    title: '解析与验证',
    code: `versions parse v1.2.3-rc1
versions validate 1.2.3
versions info v1.2.3-beta1`,
  },
  {
    title: '比较与检查',
    code: `versions compare 1.0.0 2.0.0
versions check --stable 1.2.3
versions check --beta 1.2.3-beta1
versions check --newer 1.0.0 1.5.0`,
  },
  {
    title: '排序与过滤',
    code: `versions sort 3.0.0 1.0.0 2.0.0
versions sort --desc 3.0.0 1.0.0 2.0.0
versions filter --stable 1.0.0-alpha 1.0.0 2.0.0-beta 2.0.0
versions filter --constraint ">=1.0.0,<2.0.0" 0.5.0 1.0.0 1.5.0 2.0.0`,
  },
  {
    title: '分组与范围',
    code: `versions group 1.0.0 1.1.0 2.0.0
versions range 1.0.0 2.0.0 1.0.0 1.5.0 2.0.0 3.0.0
versions satisfies 1.5.0 ">=1.0.0,<2.0.0"`,
  },
]

const CodeBlock: React.FC<{ code: string }> = ({ code }) => (
  <pre
    style={{
      background: '#1e1e2e',
      color: '#cdd6f4',
      padding: 20,
      borderRadius: 12,
      fontSize: 13,
      lineHeight: 1.65,
      overflow: 'auto',
      margin: 0,
      border: '1px solid #313244',
    }}
  >
    <code>{code}</code>
  </pre>
)

const QuickStartSection: React.FC = () => {
  return (
    <div id="quickstart" style={{ padding: '80px 32px', background: '#ffffff' }}>
      <div style={{ maxWidth: 1200, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">快速开始</div>
          <p className="section-subtitle">几行代码即可上手，覆盖版本号操作全场景</p>
        </div>

        <Tabs
          defaultActiveKey="sdk"
          centered
          size="large"
          items={[
            {
              key: 'sdk',
              label: (
                <span style={{ fontWeight: 600 }}>
                  <CodeOutlined /> Go SDK
                </span>
              ),
              children: (
                <Row gutter={[20, 20]}>
                  {sdkExamples.map((example, index) => (
                    <Col xs={24} md={12} key={index}>
                      <Card
                        title={example.title}
                        size="small"
                        style={{ borderRadius: 12, border: '1px solid #e2e8f0' }}
                        styles={{ header: { fontWeight: 700, fontSize: 15 } }}
                      >
                        <CodeBlock code={example.code} />
                      </Card>
                    </Col>
                  ))}
                </Row>
              ),
            },
            {
              key: 'cli',
              label: (
                <span style={{ fontWeight: 600 }}>
                  <CodeSandboxOutlined /> CLI
                </span>
              ),
              children: (
                <Row gutter={[20, 20]}>
                  {cliExamples.map((example, index) => (
                    <Col xs={24} md={12} key={index}>
                      <Card
                        title={example.title}
                        size="small"
                        style={{ borderRadius: 12, border: '1px solid #e2e8f0' }}
                        styles={{ header: { fontWeight: 700, fontSize: 15 } }}
                      >
                        <CodeBlock code={example.code} />
                      </Card>
                    </Col>
                  ))}
                </Row>
              ),
            },
          ]}
        />

        <div style={{ textAlign: 'center', marginTop: 32 }}>
          <Text type="secondary" style={{ fontSize: 15 }}>
            更多用法请参考{' '}
            <a
              href="https://github.com/scagogogo/versions-skills#readme"
              target="_blank"
              rel="noopener noreferrer"
              style={{ color: '#4f46e5', fontWeight: 600 }}
            >
              GitHub README
            </a>{' '}
            和{' '}
            <a
              href="https://pkg.go.dev/github.com/scagogogo/versions-skills"
              target="_blank"
              rel="noopener noreferrer"
              style={{ color: '#4f46e5', fontWeight: 600 }}
            >
              Go Doc
            </a>
          </Text>
        </div>
      </div>
    </div>
  )
}

export default QuickStartSection
