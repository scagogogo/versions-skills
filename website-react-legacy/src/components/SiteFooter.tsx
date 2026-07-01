import { Layout } from 'antd'
import { GithubOutlined, BookOutlined, SendOutlined } from '@ant-design/icons'

const { Footer } = Layout

const SiteFooter: React.FC = () => {
  return (
    <Footer
      style={{
        textAlign: 'center',
        background: '#0f172a',
        color: 'rgba(255,255,255,0.5)',
        padding: '28px 24px',
        fontSize: 13,
      }}
    >
      <div style={{ display: 'flex', justifyContent: 'center', gap: 24, marginBottom: 16, flexWrap: 'wrap' }}>
        <a href="https://github.com/scagogogo/versions-skills" target="_blank" rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.6)', display: 'flex', alignItems: 'center', gap: 6, fontSize: 13 }}>
          <GithubOutlined style={{ fontSize: 15 }} /> GitHub
        </a>
        <a href="https://pkg.go.dev/github.com/scagogogo/versions-skills" target="_blank" rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.6)', display: 'flex', alignItems: 'center', gap: 6, fontSize: 13 }}>
          <BookOutlined style={{ fontSize: 15 }} /> Go Doc
        </a>
        <a href="https://github.com/scagogogo/versions-skills/releases/latest" target="_blank" rel="noopener noreferrer"
          style={{ color: 'rgba(255,255,255,0.6)', display: 'flex', alignItems: 'center', gap: 6, fontSize: 13 }}>
          <SendOutlined style={{ fontSize: 15 }} /> Releases
        </a>
      </div>
      <div>
        Versions-Skills ©{new Date().getFullYear()} — MIT — by{' '}
        <a href="https://github.com/scagogogo" target="_blank" rel="noopener noreferrer" style={{ color: '#60a5fa' }}>
          scagogogo
        </a>
      </div>
    </Footer>
  )
}

export default SiteFooter
