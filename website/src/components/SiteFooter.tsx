import { Layout } from 'antd'
import { GithubOutlined, BookOutlined, SendOutlined } from '@ant-design/icons'

const { Footer } = Layout

const SiteFooter: React.FC = () => {
  return (
    <Footer
      style={{
        textAlign: 'center',
        background: '#0f172a',
        color: 'rgba(255,255,255,0.6)',
        padding: '32px 48px',
      }}
    >
      <div style={{ display: 'flex', justifyContent: 'center', gap: 32, marginBottom: 20, flexWrap: 'wrap' }}>
        <a
          href="https://github.com/scagogogo/versions-skills"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.7)', display: 'flex', alignItems: 'center', gap: 8, transition: 'color 0.2s' }}
          onMouseEnter={(e) => (e.currentTarget.style.color = '#818cf8')}
          onMouseLeave={(e) => (e.currentTarget.style.color = 'rgba(255,255,255,0.7)')}
        >
          <GithubOutlined style={{ fontSize: 18 }} /> GitHub
        </a>
        <a
          href="https://pkg.go.dev/github.com/scagogogo/versions-skills"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.7)', display: 'flex', alignItems: 'center', gap: 8, transition: 'color 0.2s' }}
          onMouseEnter={(e) => (e.currentTarget.style.color = '#818cf8')}
          onMouseLeave={(e) => (e.currentTarget.style.color = 'rgba(255,255,255,0.7)')}
        >
          <BookOutlined style={{ fontSize: 18 }} /> Go Doc
        </a>
        <a
          href="https://github.com/scagogogo/versions-skills/releases/latest"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.7)', display: 'flex', alignItems: 'center', gap: 8, transition: 'color 0.2s' }}
          onMouseEnter={(e) => (e.currentTarget.style.color = '#818cf8')}
          onMouseLeave={(e) => (e.currentTarget.style.color = 'rgba(255,255,255,0.7)')}
        >
          <SendOutlined style={{ fontSize: 18 }} /> Releases
        </a>
      </div>
      <div style={{ fontSize: 14 }}>
        Versions-Skills ©{new Date().getFullYear()} — MIT License — Built with ❤️ by{' '}
        <a
          href="https://github.com/scagogogo"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: '#818cf8', transition: 'color 0.2s' }}
        >
          scagogogo
        </a>
      </div>
    </Footer>
  )
}

export default SiteFooter
