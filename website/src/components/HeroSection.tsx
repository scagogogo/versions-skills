import { Typography, Space, Button, Tag } from 'antd'
import {
  RocketOutlined,
  CodeOutlined,
  ApiOutlined,
  ThunderboltOutlined,
} from '@ant-design/icons'

const { Title, Paragraph } = Typography

const HeroSection: React.FC = () => {
  return (
    <div
      style={{
        background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
        padding: '80px 48px',
        textAlign: 'center',
        color: '#fff',
      }}
    >
      <Space direction="vertical" size="large" style={{ maxWidth: 900, margin: '0 auto' }}>
        <div>
          <Tag
            color="blue"
            style={{
              fontSize: 14,
              padding: '4px 16px',
              borderRadius: 20,
              marginBottom: 16,
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
            fontSize: 48,
            lineHeight: 1.2,
          }}
        >
          Versions-Skills
        </Title>

        <Title
          level={3}
          style={{
            color: 'rgba(255,255,255,0.9)',
            margin: 0,
            fontWeight: 400,
            fontSize: 22,
          }}
        >
          强大的 Go 语言版本号解析、比较、排序、分组和约束检查库
        </Title>

        <Paragraph
          style={{
            color: 'rgba(255,255,255,0.8)',
            fontSize: 16,
            maxWidth: 700,
            margin: '0 auto',
          }}
        >
          通过 Skills · Go SDK · CLI · MCP Server 四种方式接入，
          <br />
          兼容 Claude Code、Cursor、Windsurf、VS Code Copilot 及所有 MCP 兼容的 AI Agent
        </Paragraph>

        <Space size="large" wrap style={{ marginTop: 8 }}>
          <Tag
            icon={<RocketOutlined />}
            color="cyan"
            style={{ fontSize: 14, padding: '6px 14px', borderRadius: 6 }}
          >
            13 Skills
          </Tag>
          <Tag
            icon={<CodeOutlined />}
            color="green"
            style={{ fontSize: 14, padding: '6px 14px', borderRadius: 6 }}
          >
            21 MCP Tools
          </Tag>
          <Tag
            icon={<ApiOutlined />}
            color="orange"
            style={{ fontSize: 14, padding: '6px 14px', borderRadius: 6 }}
          >
            40+ CLI Commands
          </Tag>
          <Tag
            icon={<ThunderboltOutlined />}
            color="magenta"
            style={{ fontSize: 14, padding: '6px 14px', borderRadius: 6 }}
          >
            Zero Dependencies
          </Tag>
        </Space>

        <Space size="middle" style={{ marginTop: 16 }}>
          <Button
            type="primary"
            size="large"
            href="#quickstart"
            style={{
              height: 48,
              paddingInline: 32,
              fontSize: 16,
              background: '#fff',
              color: '#667eea',
              borderColor: '#fff',
            }}
          >
            快速开始
          </Button>
          <Button
            size="large"
            href="https://github.com/scagogogo/versions-skills"
            target="_blank"
            style={{
              height: 48,
              paddingInline: 32,
              fontSize: 16,
              color: '#fff',
              borderColor: 'rgba(255,255,255,0.6)',
            }}
            ghost
          >
            GitHub →
          </Button>
        </Space>
      </Space>
    </div>
  )
}

export default HeroSection
