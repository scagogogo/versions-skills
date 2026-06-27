import { Layout } from 'antd'
import SiteHeader from '../components/SiteHeader'
import SiteFooter from '../components/SiteFooter'
import HeroSection from '../components/HeroSection'
import FeaturesSection from '../components/FeaturesSection'
import ArchitectureSection from '../components/ArchitectureSection'
import AccessSection from '../components/AccessSection'
import CasesSection from '../components/CasesSection'
import TutorialsSection from '../components/TutorialsSection'
import AiIntegrationSection from '../components/AiIntegrationSection'
import QuickStartSection from '../components/QuickStartSection'

const { Content } = Layout

const HomePage: React.FC = () => {
  return (
    <Layout style={{ minHeight: '100vh' }}>
      <SiteHeader />
      <Content>
        <HeroSection />
        <FeaturesSection />
        <ArchitectureSection />
        <AccessSection />
        <CasesSection />
        <TutorialsSection />
        <AiIntegrationSection />
        <QuickStartSection />
      </Content>
      <SiteFooter />
    </Layout>
  )
}

export default HomePage
