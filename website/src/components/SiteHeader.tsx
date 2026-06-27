import { Layout, Menu, Button, Drawer } from 'antd'
import { GithubOutlined, MenuOutlined } from '@ant-design/icons'
import { useLocation, useNavigate } from 'react-router-dom'
import { useState } from 'react'

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
  const [drawerOpen, setDrawerOpen] = useState(false)

  const handleMenuClick = (info: { key: string }) => {
    const key = info.key
    setDrawerOpen(false)
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
    <>
      <Header
        style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          background: '#fff',
          borderBottom: '1px solid #e2e8f0',
          padding: '0 24px',
          position: 'sticky',
          top: 0,
          zIndex: 100,
          height: 56,
        }}
      >
        <div style={{ display: 'flex', alignItems: 'center', flex: 1, minWidth: 0 }}>
          <div
            style={{
              fontSize: 17,
              fontWeight: 700,
              marginRight: 24,
              cursor: 'pointer',
              color: '#2563eb',
              letterSpacing: '-0.02em',
              flexShrink: 0,
            }}
            onClick={() => navigate('/')}
          >
            Versions-Skills
          </div>
          <Menu
            mode="horizontal"
            selectedKeys={[location.pathname]}
            items={navItems}
            onClick={handleMenuClick}
            style={{ border: 'none', fontWeight: 500, fontSize: 14 }}
            className="desktop-nav"
          />
        </div>
        <Button
          type="default"
          href="https://github.com/scagogogo/versions-skills"
          target="_blank"
          icon={<GithubOutlined />}
          style={{ fontWeight: 600, fontSize: 14, borderRadius: 4 }}
        >
          GitHub
        </Button>
        <Button
          type="text"
          icon={<MenuOutlined />}
          onClick={() => setDrawerOpen(true)}
          style={{ display: 'none', marginLeft: 8 }}
          className="mobile-nav-btn"
        />
      </Header>

      <Drawer
        title="导航"
        placement="right"
        onClose={() => setDrawerOpen(false)}
        open={drawerOpen}
        width={260}
      >
        <Menu mode="vertical" selectedKeys={[location.pathname]} items={navItems} onClick={handleMenuClick} style={{ border: 'none' }} />
      </Drawer>

      <style>{`
        @media (min-width: 900px) { .desktop-nav { display: flex !important; } }
        @media (max-width: 899px) { .mobile-nav-btn { display: inline-flex !important; } .desktop-nav { display: none !important; } }
      `}</style>
    </>
  )
}

export default SiteHeader
