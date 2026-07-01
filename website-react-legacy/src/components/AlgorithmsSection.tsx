import { Typography, Row, Col, Card, Table, Tag } from 'antd'
import {
  ScissorOutlined,
  SwapOutlined,
  SortAscendingOutlined,
  FilterOutlined,
  ApartmentOutlined,
  FieldBinaryOutlined,
} from '@ant-design/icons'

const { Title, Paragraph, Text } = Typography

// 解析：三段式 + 元数据 的字段拆解
const parseFields = [
  { seg: 'Prefix', desc: '数字之前的非数字前导', ex: '"v"、"release-"' },
  { seg: 'VersionNumbers', desc: '整数段，解析后统一以 "." 拼接', ex: '[1, 2, 3]' },
  { seg: 'Suffix', desc: '数字部分之后的全部内容', ex: '"-beta1"' },
  { seg: 'Metadata', desc: '"+" 之后且不含 "-" 的部分（semver 元数据）', ex: '"build.7"' },
]

// 比较优先级
const comparePriority = [
  { key: '1', field: 'VersionNumbers', rule: '从左到右逐位 int 比较；共享位相等时更长者更大', src: 'version_numbers.go' },
  { key: '2', field: 'Suffix', rule: '稳定版（无后缀）> 预发布版；后缀按权重比', src: 'version_suffix.go' },
  { key: '3', field: 'PublicTime', rule: '两者均非零时，晚者胜（仅在前两级相等时生效）', src: 'version.go' },
  { key: '4', field: 'Raw', rule: '最终兜底：原始字符串字典序', src: 'version.go' },
]

// 后缀权重表
const suffixWeights = [
  { w: '50', pat: 'dev / dev1', mean: '开发构建' },
  { w: '60', pat: 'snapshot', mean: '快照' },
  { w: '70', pat: 'nightly', mean: '夜间构建' },
  { w: '100', pat: 'a / alpha / alpha1', mean: 'Alpha' },
  { w: '200', pat: 'b / beta / beta2', mean: 'Beta' },
  { w: '300', pat: 'm / milestone / m1', mean: '里程碑' },
  { w: '400', pat: 'rc / rc1', mean: '候选发布' },
  { w: '410', pat: 'pre / pre1', mean: '预发布' },
  { w: '420', pat: 'cr / cr1', mean: 'RC 的 CR 变体' },
  { w: '500', pat: 'final / release / ga', mean: '正式版（与无后缀等权）' },
  { w: '600', pat: 'sp / sp1', mean: '服务包' },
  { w: '700', pat: 'patch / patch1', mean: '补丁' },
  { w: '800', pat: 'post / post1', mean: '后发布（PEP 440）' },
]

// 约束操作符
const constraintOps = [
  { op: '=', base: '1.2.3', when: 'v == 1.2.3（裸写即等于）' },
  { op: '!=', base: '1.2.3', when: 'v != 1.2.3' },
  { op: '> < >= <=', base: '1.2.3', when: '直接 CompareTo' },
  { op: '^', base: '^1.2.3', when: '>=1.2.3 且 <2.0.0（左起首个非零位进一）' },
  { op: '~', base: '~1.2.3', when: '>=1.2.3 且 <1.3.0（锁定到 minor）' },
  { op: 'x / X / *', base: '1.x', when: '>=1.0.0 且 <2.0.0（末位非零进一）' },
]

const groupGranularity = [
  { fn: 'Group()', by: '完整数字串', key: 'map[string]*VersionGroup', ex: '1.2.3 与 1.2.4 分属不同组' },
  { fn: 'GroupByMajor()', by: '仅 major 段', key: 'map[int][]*Version', ex: '1.2.3、1.9.0 都在组 1' },
  { fn: 'GroupByMinor()', by: 'major.minor', key: 'map[string][]*Version', ex: '1.2.3、1.2.4 在组 "1.2"' },
]

const AlgorithmsSection: React.FC = () => {
  return (
    <div id="algorithms" style={{ padding: '64px 24px', background: '#fff' }}>
      <div style={{ maxWidth: 1120, margin: '0 auto' }}>
        <div style={{ textAlign: 'center' }}>
          <div className="section-title">算法详解</div>
          <p className="section-subtitle">不读源码也能预测行为 —— 每条规则都可在对应源文件中验证</p>
        </div>

        {/* 1. 解析 */}
        <Card className="flat-card" style={{ marginBottom: 16, borderRadius: 4, border: '1px solid #e2e8f0' }} styles={{ body: { padding: 20 } }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <ScissorOutlined style={{ color: '#2563eb', fontSize: 18 }} />
            <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>1. 解析 —— 三段式 + 元数据</Title>
            <Text code style={{ fontSize: 11, color: '#64748b' }}>parser.go</Text>
          </div>
          <pre style={{ background: '#1e293b', color: '#e2e8f0', padding: 12, borderRadius: 4, fontSize: 12, lineHeight: 1.6, overflow: 'auto', margin: '0 0 12px' }}>
            <code>{' v1.2.3-beta1+build.7\n │ └─┬─┘ └──┬──┘ └──┬──┘\n │  │      │       └─ Metadata   (+ 之后)\n │  │      └─ Suffix             (数字部分之后)\n │  └─ VersionNumbers            (整数段)\n └─ Prefix                       (非数字前导)'}</code>
          </pre>
          <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 8 }}>
            关键规则："+" 之后的内容<strong>仅当不含 "-"</strong>时才视为元数据 —— 因此 <Text code style={{ fontSize: 12 }}>0.9.0+121-bcc5decc</Text> 把 <Text code style={{ fontSize: 12 }}>121-bcc5decc</Text> 留在后缀里（Scala/Maven 风格），区分 semver 元数据与含 "-" 的预发布标识。
          </Paragraph>
          <Table
            size="small"
            pagination={false}
            dataSource={parseFields}
            rowKey="seg"
            columns={[
              { title: '字段', dataIndex: 'seg', key: 'seg', width: 140, render: (t: string) => <Text code style={{ fontSize: 12, color: '#2563eb' }}>{t}</Text> },
              { title: '说明', dataIndex: 'desc', key: 'desc' },
              { title: '示例', dataIndex: 'ex', key: 'ex', render: (t: string) => <Text code style={{ fontSize: 12 }}>{t}</Text> },
            ]}
          />
        </Card>

        {/* 2. 比较 */}
        <Card className="flat-card" style={{ marginBottom: 16, borderRadius: 4, border: '1px solid #e2e8f0' }} styles={{ body: { padding: 20 } }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <SwapOutlined style={{ color: '#0ea5e9', fontSize: 18 }} />
            <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>2. 比较 —— 四级优先级</Title>
            <Text code style={{ fontSize: 11, color: '#64748b' }}>version.go · CompareTo</Text>
          </div>
          <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 8 }}>
            按顺序尝试以下键，<strong>第一个不同的胜出</strong>。正是这套排序让 <Text code style={{ fontSize: 12 }}>1.0.0-alpha &lt; 1.0.0-beta &lt; 1.0.0-rc1 &lt; 1.0.0</Text>，贴合真实发布阶梯而非 ASCII 序。
          </Paragraph>
          <Table
            size="small"
            pagination={false}
            dataSource={comparePriority}
            rowKey="key"
            columns={[
              { title: '#', dataIndex: 'key', key: 'key', width: 48, align: 'center' },
              { title: '键', dataIndex: 'field', key: 'field', width: 140, render: (t: string) => <Text strong style={{ fontSize: 13 }}>{t}</Text> },
              { title: '规则', dataIndex: 'rule', key: 'rule' },
              { title: '源', dataIndex: 'src', key: 'src', width: 140, render: (t: string) => <Text code style={{ fontSize: 11 }}>{t}</Text> },
            ]}
          />
        </Card>

        {/* 3. 后缀权重 */}
        <Card className="flat-card" style={{ marginBottom: 16, borderRadius: 4, border: '1px solid #e2e8f0' }} styles={{ body: { padding: 20 } }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <SortAscendingOutlined style={{ color: '#16a34a', fontSize: 18 }} />
            <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>3. 后缀权重排序</Title>
            <Text code style={{ fontSize: 11, color: '#64748b' }}>suffix_weight.go</Text>
          </div>
          <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 8 }}>
            每个后缀与有序模式表匹配，命中权重决定排名；权重相同再用尾部整数（子版本号，如 <Text code style={{ fontSize: 12 }}>-alpha1</Text> 的 <Text code style={{ fontSize: 12 }}>1</Text>）破平。未知后缀排在已知后缀之后。
          </Paragraph>
          <Row gutter={[12, 12]}>
            {suffixWeights.map((s) => (
              <Col xs={12} sm={8} md={6} key={s.w}>
                <div style={{ border: '1px solid #e2e8f0', borderRadius: 4, padding: '10px 12px', background: '#f8fafc' }}>
                  <Text strong style={{ color: '#16a34a', fontSize: 16 }}>{s.w}</Text>
                  <div style={{ marginTop: 2 }}><Text code style={{ fontSize: 11, color: '#475569' }}>{s.pat}</Text></div>
                  <Text type="secondary" style={{ fontSize: 11 }}>{s.mean}</Text>
                </div>
              </Col>
            ))}
          </Row>
        </Card>

        {/* 4. 约束语法 */}
        <Card className="flat-card" style={{ marginBottom: 16, borderRadius: 4, border: '1px solid #e2e8f0' }} styles={{ body: { padding: 20 } }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <FilterOutlined style={{ color: '#ea580c', fontSize: 18 }} />
            <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>4. 约束语法 —— 三层</Title>
            <Text code style={{ fontSize: 11, color: '#64748b' }}>constraint.go</Text>
          </div>
          <pre style={{ background: '#1e293b', color: '#e2e8f0', padding: 12, borderRadius: 4, fontSize: 12, lineHeight: 1.6, overflow: 'auto', margin: '0 0 12px' }}>
            <code>{'Union  (OR) : ">=1.0.0,<2.0.0 || >=3.0.0"   ← 以 "||" 切分\n  └─ Set(AND): ">=1.0.0,<2.0.0"          ← 以 ","  切分\n       └─ Single: ">=1.0.0" | "^1.2.3" | "~1.2" | "1.x"'}</code>
          </pre>
          <Table
            size="small"
            pagination={false}
            dataSource={constraintOps}
            rowKey="op"
            columns={[
              { title: '操作符', dataIndex: 'op', key: 'op', width: 110, render: (t: string) => <Text code style={{ fontSize: 12, color: '#2563eb', fontWeight: 600 }}>{t}</Text> },
              { title: 'base', dataIndex: 'base', key: 'base', width: 120, render: (t: string) => <Text code style={{ fontSize: 12 }}>{t}</Text> },
              { title: 'v 命中条件', dataIndex: 'when', key: 'when' },
            ]}
          />
          <div style={{ marginTop: 12, display: 'flex', gap: 8, flexWrap: 'wrap' }}>
            <Tag style={{ fontSize: 12, borderRadius: 2, color: '#2563eb', borderColor: '#bfdbfe', background: '#eff6ff' }}>{'^0.2.3 → >=0.2.3, <0.3.0'}</Tag>
            <Tag style={{ fontSize: 12, borderRadius: 2, color: '#0ea5e9', borderColor: '#bae6fd', background: '#f0f9ff' }}>{'~1.2 → >=1.2.0, <1.3.0'}</Tag>
            <Tag style={{ fontSize: 12, borderRadius: 2, color: '#16a34a', borderColor: '#bbf7d0', background: '#f0fdf4' }}>{'1.2.x → >=1.2.0, <1.3.0'}</Tag>
          </div>
        </Card>

        {/* 5. 范围查询 */}
        <Card className="flat-card" style={{ marginBottom: 16, borderRadius: 4, border: '1px solid #e2e8f0' }} styles={{ body: { padding: 20 } }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <FieldBinaryOutlined style={{ color: '#2563eb', fontSize: 18 }} />
            <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>5. 范围查询 —— 开/闭区间</Title>
            <Text code style={{ fontSize: 11, color: '#64748b' }}>version_range.go</Text>
          </div>
          <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 8 }}>
            <Text code style={{ fontSize: 12 }}>NewClosedRange(1.0.0, 2.0.0)</Text> = <Text code style={{ fontSize: 12 }}>[1.0.0, 2.0.0]</Text> 两端含；<Text code style={{ fontSize: 12 }}>NewOpenRange</Text> 两端不含；可混用 <Text code style={{ fontSize: 12 }}>[1.0.0, 2.0.0)</Text>。nil 边界表示无界。
          </Paragraph>
        </Card>

        {/* 6. 排序与分组 */}
        <Card className="flat-card" style={{ marginBottom: 16, borderRadius: 4, border: '1px solid #e2e8f0' }} styles={{ body: { padding: 20 } }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <ApartmentOutlined style={{ color: '#0ea5e9', fontSize: 18 }} />
            <Title level={5} style={{ margin: 0, fontWeight: 600, fontSize: 15 }}>6. 排序与分组 —— 两阶段、组感知</Title>
            <Text code style={{ fontSize: 11, color: '#64748b' }}>sort.go · version_group.go</Text>
          </div>
          <Paragraph style={{ color: '#64748b', fontSize: 13, marginBottom: 8 }}>
            <Text code style={{ fontSize: 12 }}>SortVersionSlice</Text> 先按 <Text code style={{ fontSize: 12 }}>BuildGroupID()</Text> 分组、组间排序、组内再排序后拼接。收益：<Text code style={{ fontSize: 12 }}>1.10.0</Text> 正确排在 <Text code style={{ fontSize: 12 }}>1.2.0</Text> 之后（数值而非字符串序），同族聚拢。
          </Paragraph>
          <Table
            size="small"
            pagination={false}
            dataSource={groupGranularity}
            rowKey="fn"
            columns={[
              { title: '函数', dataIndex: 'fn', key: 'fn', width: 160, render: (t: string) => <Text code style={{ fontSize: 12, color: '#2563eb', fontWeight: 600 }}>{t}</Text> },
              { title: '分组依据', dataIndex: 'by', key: 'by', width: 130 },
              { title: '键类型', dataIndex: 'key', key: 'key', render: (t: string) => <Text code style={{ fontSize: 11 }}>{t}</Text> },
              { title: '示例', dataIndex: 'ex', key: 'ex' },
            ]}
          />
        </Card>
      </div>
    </div>
  )
}

export default AlgorithmsSection
