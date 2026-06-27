import { Typography, Space, Button } from 'antd'
import { ArrowRightOutlined } from '@ant-design/icons'

const { Title, Paragraph } = Typography

const HeroSection: React.FC = () => {
  return (
    <div
      style={{
        background: '#0f172a',
        padding: '72px 24px 60px',
        textAlign: 'center',
        color: '#fff',
      }}
    >
      <Space direction="vertical" size={16} style={{ maxWidth: 800, margin: '0 auto' }}>
        <Title
          level={1}
          style={{
            color: '#fff',
            margin: 0,
            fontSize: 44,
            lineHeight: 1.15,
            fontWeight: 700,
            letterSpacing: '-0.03em',
          }}
        >
          Versions-Skills
        </Title>

        <Paragraph
          style={{
            color: 'rgba(255,255,255,0.7)',
            fontSize: 17,
            maxWidth: 600,
            margin: '0 auto',
            lineHeight: 1.6,
          }}
        >
          Go 语言版本号解析、比较、排序、分组和约束检查库<br />
          Skills · Go SDK · CLI · MCP Server 四种接入
        </Paragraph>

        <div style={{ display: 'flex', justifyContent: 'center', gap: 12, flexWrap: 'wrap' }}>
          <span style={{ background: 'rgba(37,99,235,0.3)', color: '#93c5fd', padding: '4px 12px', borderRadius: 4, fontSize: 13, fontWeight: 600 }}>
            13 Skills
          </span>
          <span style={{ background: 'rgba(14,165,233,0.3)', color: '#7dd3fc', padding: '4px 12px', borderRadius: 4, fontSize: 13, fontWeight: 600 }}>
            21 MCP Tools
          </span>
          <span style={{ background: 'rgba(16,185,129,0.3)', color: '#6ee7b7', padding: '4px 12px', borderRadius: 4, fontSize: 13, fontWeight: 600 }}>
            40+ CLI
          </span>
          <span style={{ background: 'rgba(234,179,8,0.3)', color: '#fde68a', padding: '4px 12px', borderRadius: 4, fontSize: 13, fontWeight: 600 }}>
            Zero Deps
          </span>
        </div>

        <Space size="middle" style={{ marginTop: 12 }}>
          <Button
            type="primary"
            size="large"
            href="#quickstart"
            icon={<ArrowRightOutlined />}
            style={{ height: 44, paddingInline: 28, fontSize: 15, fontWeight: 600, borderRadius: 4 }}
          >
            快速开始
          </Button>
          <Button
            size="large"
            href="https://github.com/scagogogo/versions-skills"
            target="_blank"
            style={{ height: 44, paddingInline: 28, fontSize: 15, fontWeight: 600, borderRadius: 4, color: '#fff', borderColor: 'rgba(255,255,255,0.3)' }}
            ghost
          >
            GitHub
          </Button>
        </Space>

        <Paragraph style={{ color: 'rgba(255,255,255,0.45)', fontSize: 13, marginTop: 4 }}>
          兼容 Claude Code · Cursor · Windsurf · VS Code Copilot
        </Paragraph>
      </Space>
    </div>
  )
}

export default HeroSection
