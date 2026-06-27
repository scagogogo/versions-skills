import { Layout, Menu, Button, Space } from 'antd'
import { GithubOutlined } from '@ant-design/icons'
import { useLocation, useNavigate } from 'react-router-dom'

const { Header } = Layout

const navItems = [
  { key: '/', label: '首页' },
  { key: '/#features', label: '功能特性' },
  { key: '/#access', label: '接入方式' },
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
    <Header
      style={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        background: '#fff',
        borderBottom: '1px solid #f0f0f0',
        padding: '0 48px',
        position: 'sticky',
        top: 0,
        zIndex: 100,
        boxShadow: '0 2px 8px rgba(0,0,0,0.06)',
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <div
          style={{
            fontSize: 20,
            fontWeight: 700,
            marginRight: 40,
            cursor: 'pointer',
            color: '#1677ff',
          }}
          onClick={() => navigate('/')}
        >
          🔄 Versions-Skills
        </div>
        <Menu
          mode="horizontal"
          selectedKeys={[location.pathname]}
          items={navItems}
          onClick={handleMenuClick}
          style={{ border: 'none', flex: 1 }}
        />
      </div>
      <Space>
        <Button
          type="primary"
          href="https://github.com/scagogogo/versions-skills"
          target="_blank"
          icon={<GithubOutlined />}
          size="large"
        >
          GitHub
        </Button>
      </Space>
    </Header>
  )
}

export default SiteHeader
