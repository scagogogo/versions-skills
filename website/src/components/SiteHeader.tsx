import { Layout, Menu, Button } from 'antd'
import { GithubOutlined } from '@ant-design/icons'
import { useLocation, useNavigate } from 'react-router-dom'

const { Header } = Layout

const navItems = [
  { key: '/', label: '首页' },
  { key: '/#features', label: '功能' },
  { key: '/#access', label: '接入方式' },
  { key: '/#cases', label: '案例' },
  { key: '/#tutorials', label: '教程' },
  { key: '/#ai-integration', label: 'AI 集成' },
  { key: '/#quickstart', label: '快速开始' },
]

const SiteHeader: React.FC = () => {
  const location = useLocation()
  const navigate = useNavigate()

  const handleMenuClick = (info: { key: string }) => {
    const key = info.key
    if (key.startsWith('/#')) {
      const id = key.slice(2)
      if (location.pathname !== '/') {
        navigate('/')
        setTimeout(() => document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' }), 100)
      } else {
        document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' })
      }
    } else {
      navigate(key)
    }
  }

  return (
    <Header
      style={{
        display: 'flex',
        alignItems: 'center',
        background: '#fff',
        borderBottom: '1px solid #e2e8f0',
        padding: '0 24px',
        position: 'sticky',
        top: 0,
        zIndex: 100,
        height: 56,
      }}
    >
      <div
        style={{
          fontSize: 17,
          fontWeight: 700,
          marginRight: 20,
          cursor: 'pointer',
          color: '#2563eb',
          letterSpacing: '-0.02em',
          flexShrink: 0,
        }}
        onClick={() => navigate('/')}
      >
        Versions-Skills
      </div>
      <div style={{ flex: 1, minWidth: 0, overflowX: 'auto' }}>
        <Menu
          mode="horizontal"
          selectedKeys={[location.pathname]}
          items={navItems}
          onClick={handleMenuClick}
          style={{ border: 'none', fontWeight: 500, fontSize: 14, minWidth: 'max-content' }}
        />
      </div>
      <Button
        type="default"
        href="https://github.com/scagogogo/versions-skills"
        target="_blank"
        icon={<GithubOutlined />}
        style={{ fontWeight: 600, fontSize: 14, borderRadius: 4, flexShrink: 0 }}
      >
        GitHub
      </Button>
    </Header>
  )
}

export default SiteHeader
