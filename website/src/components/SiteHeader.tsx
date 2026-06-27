import { Layout, Menu, Button, Space, Drawer } from 'antd'
import { GithubOutlined, MenuOutlined } from '@ant-design/icons'
import { useLocation, useNavigate } from 'react-router-dom'
import { useState } from 'react'

const { Header } = Layout

const navItems = [
  { key: '/', label: '首页' },
  { key: '/#features', label: '功能特性' },
  { key: '/#architecture', label: '架构' },
  { key: '/#access', label: '接入方式' },
  { key: '/#cases', label: '使用案例' },
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
        setTimeout(() => {
          document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' })
        }, 100)
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
          background: 'rgba(255,255,255,0.92)',
          backdropFilter: 'blur(12px)',
          borderBottom: '1px solid #e2e8f0',
          padding: '0 32px',
          position: 'sticky',
          top: 0,
          zIndex: 100,
          height: 64,
        }}
      >
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <div
            style={{
              fontSize: 20,
              fontWeight: 800,
              marginRight: 32,
              cursor: 'pointer',
              color: '#4f46e5',
              letterSpacing: '-0.02em',
              flexShrink: 0,
            }}
            onClick={() => navigate('/')}
          >
            🔄 Versions
          </div>
          <Menu
            mode="horizontal"
            selectedKeys={[location.pathname]}
            items={navItems}
            onClick={handleMenuClick}
            style={{
              border: 'none',
              flex: 1,
              fontWeight: 500,
              display: 'none',
            }}
            className="desktop-menu"
          />
        </div>
        <Space>
          <Button
            type="primary"
            href="https://github.com/scagogogo/versions-skills"
            target="_blank"
            icon={<GithubOutlined />}
            style={{
              fontWeight: 600,
              borderRadius: 10,
              height: 40,
            }}
          >
            GitHub
          </Button>
          <Button
            type="text"
            icon={<MenuOutlined />}
            onClick={() => setDrawerOpen(true)}
            className="mobile-menu-btn"
            style={{ display: 'none' }}
          />
        </Space>
      </Header>

      <Drawer
        title="导航"
        placement="right"
        onClose={() => setDrawerOpen(false)}
        open={drawerOpen}
        width={260}
      >
        <Menu
          mode="vertical"
          selectedKeys={[location.pathname]}
          items={navItems}
          onClick={handleMenuClick}
          style={{ border: 'none' }}
        />
      </Drawer>

      <style>{`
        @media (min-width: 900px) {
          .desktop-menu { display: flex !important; }
        }
        @media (max-width: 899px) {
          .mobile-menu-btn { display: inline-flex !important; }
        }
      `}</style>
    </>
  )
}

export default SiteHeader
