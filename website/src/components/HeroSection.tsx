import { Typography, Space, Button, Tag } from 'antd'
import {
  RocketOutlined,
  CodeOutlined,
  ApiOutlined,
  ThunderboltOutlined,
  ArrowRightOutlined,
} from '@ant-design/icons'

const { Title, Paragraph } = Typography

const HeroSection: React.FC = () => {
  return (
    <div
      style={{
        background: 'linear-gradient(135deg, #312e81 0%, #4f46e5 30%, #7c3aed 60%, #06b6d4 100%)',
        padding: '100px 48px 80px',
        textAlign: 'center',
        color: '#fff',
        position: 'relative',
        overflow: 'hidden',
      }}
    >
      {/* Decorative circles */}
      <div
        style={{
          position: 'absolute',
          top: -120,
          right: -80,
          width: 360,
          height: 360,
          borderRadius: '50%',
          background: 'rgba(255,255,255,0.04)',
          pointerEvents: 'none',
        }}
      />
      <div
        style={{
          position: 'absolute',
          bottom: -60,
          left: -40,
          width: 240,
          height: 240,
          borderRadius: '50%',
          background: 'rgba(255,255,255,0.03)',
          pointerEvents: 'none',
        }}
      />

      <Space direction="vertical" size="large" style={{ maxWidth: 900, margin: '0 auto', position: 'relative', zIndex: 1 }}>
        <div>
          <Tag
            style={{
              fontSize: 14,
              padding: '6px 20px',
              borderRadius: 24,
              marginBottom: 16,
              background: 'rgba(255,255,255,0.15)',
              color: '#fff',
              border: '1px solid rgba(255,255,255,0.25)',
              backdropFilter: 'blur(4px)',
            }}
          >
            🤖 AI Agent Ready
          </Tag>
        </div>

        <Title
          level={1}
          style={{
            color: '#fff',
            margin: 0,
            fontSize: 56,
            lineHeight: 1.15,
            fontWeight: 800,
            letterSpacing: '-0.03em',
          }}
        >
          Versions-Skills
        </Title>

        <Title
          level={3}
          style={{
            color: 'rgba(255,255,255,0.88)',
            margin: 0,
            fontWeight: 400,
            fontSize: 21,
            lineHeight: 1.5,
          }}
        >
          强大的 Go 语言版本号解析、比较、排序、分组和约束检查库
        </Title>

        <Paragraph
          style={{
            color: 'rgba(255,255,255,0.72)',
            fontSize: 16,
            maxWidth: 680,
            margin: '0 auto',
            lineHeight: 1.7,
          }}
        >
          通过 Skills · Go SDK · CLI · MCP Server 四种方式接入
          <br />
          兼容 Claude Code、Cursor、Windsurf、VS Code Copilot 及所有 MCP 兼容的 AI Agent
        </Paragraph>

        <Space size="middle" wrap style={{ marginTop: 8 }}>
          <Tag
            icon={<RocketOutlined />}
            style={{ fontSize: 13, padding: '6px 16px', borderRadius: 8, background: 'rgba(6,182,212,0.25)', color: '#a5f3fc', border: '1px solid rgba(6,182,212,0.4)' }}
          >
            13 Skills
          </Tag>
          <Tag
            icon={<CodeOutlined />}
            style={{ fontSize: 13, padding: '6px 16px', borderRadius: 8, background: 'rgba(52,211,153,0.25)', color: '#a7f3d0', border: '1px solid rgba(52,211,153,0.4)' }}
          >
            21 MCP Tools
          </Tag>
          <Tag
            icon={<ApiOutlined />}
            style={{ fontSize: 13, padding: '6px 16px', borderRadius: 8, background: 'rgba(251,191,36,0.25)', color: '#fde68a', border: '1px solid rgba(251,191,36,0.4)' }}
          >
            40+ CLI Commands
          </Tag>
          <Tag
            icon={<ThunderboltOutlined />}
            style={{ fontSize: 13, padding: '6px 16px', borderRadius: 8, background: 'rgba(244,114,182,0.25)', color: '#fbcfe8', border: '1px solid rgba(244,114,182,0.4)' }}
          >
            Zero Dependencies
          </Tag>
        </Space>

        <Space size="middle" style={{ marginTop: 20 }}>
          <Button
            type="primary"
            size="large"
            href="#quickstart"
            icon={<ArrowRightOutlined />}
            style={{
              height: 52,
              paddingInline: 36,
              fontSize: 16,
              fontWeight: 600,
              background: '#fff',
              color: '#4f46e5',
              borderColor: '#fff',
              borderRadius: 12,
              boxShadow: '0 4px 14px rgba(0,0,0,0.15)',
            }}
          >
            快速开始
          </Button>
          <Button
            size="large"
            href="https://github.com/scagogogo/versions-skills"
            target="_blank"
            icon={<RocketOutlined />}
            style={{
              height: 52,
              paddingInline: 36,
              fontSize: 16,
              fontWeight: 600,
              color: '#fff',
              borderColor: 'rgba(255,255,255,0.4)',
              borderRadius: 12,
              background: 'rgba(255,255,255,0.08)',
            }}
          >
            GitHub
          </Button>
        </Space>
      </Space>
    </div>
  )
}

export default HeroSection
