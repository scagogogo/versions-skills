import { Layout } from 'antd'
import { GithubOutlined } from '@ant-design/icons'

const { Footer } = Layout

const SiteFooter: React.FC = () => {
  return (
    <Footer
      style={{
        textAlign: 'center',
        background: '#001529',
        color: 'rgba(255,255,255,0.65)',
        padding: '24px 48px',
      }}
    >
      <div style={{ marginBottom: 8 }}>
        <a
          href="https://github.com/scagogogo/versions-skills"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.65)', marginRight: 16 }}
        >
          <GithubOutlined style={{ fontSize: 18 }} /> GitHub
        </a>
        <span style={{ margin: '0 8px' }}>|</span>
        <a
          href="https://pkg.go.dev/github.com/scagogogo/versions-skills"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.65)', marginRight: 16 }}
        >
          Go Doc
        </a>
        <span style={{ margin: '0 8px' }}>|</span>
        <a
          href="https://github.com/scagogogo/versions-skills/releases/latest"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.65)' }}
        >
          Releases
        </a>
      </div>
      <div>
        Versions-Skills ©{new Date().getFullYear()} — MIT License — Built with ❤️ by{' '}
        <a
          href="https://github.com/scagogogo"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: '#1677ff' }}
        >
          scagogogo
        </a>
      </div>
    </Footer>
  )
}

export default SiteFooter
