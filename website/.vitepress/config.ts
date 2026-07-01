import { defineConfig } from 'vitepress'
import { withMermaid } from 'vitepress-plugin-mermaid'

// https://vitepress.dev/reference/site-config
export default withMermaid(
  defineConfig({
  lang: 'zh-CN',
  title: 'Versions-Skills',
  description: '面向 AI 原生的版本号工具集 — Skills + MCP + Go SDK + CLI',
  lastUpdated: true,
  cleanUrls: true,

  // 忽略指向仓库内部目录（非站点页面）的死链检查
  ignoreDeadLinks: [
    /^https?:\/\/github\.com/,
    /\/website-react-legacy/,
  ],

  // GitHub Pages 部署在 /versions-skills/ 子路径下
  base: '/versions-skills/',

  head: [
    ['link', { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' }],
    ['meta', { name: 'theme-color', content: '#2563eb' }],
  ],

  themeConfig: {
    // 站点级导航
    nav: [
      { text: '首页', link: '/' },
      { text: '为什么用', link: '/why' },
      { text: '原理', link: '/how-it-works' },
      { text: '算法详解', link: '/algorithms' },
      { text: 'AI Agent 接入', link: '/ai-agents' },
      {
        text: '参考',
        items: [
          { text: 'Go SDK API', link: '/sdk' },
          { text: 'CLI 命令', link: '/cli' },
          { text: 'MCP 工具', link: '/mcp' },
        ],
      },
    ],

    // 侧边栏：按目录分组
    sidebar: {
      '/': [
        {
          text: '开始',
          items: [
            { text: '首页', link: '/' },
            { text: '为什么需要 versions-skills', link: '/why' },
            { text: '快速开始', link: '/quick-start' },
          ],
        },
        {
          text: '原理',
          items: [
            { text: '工作原理', link: '/how-it-works' },
            { text: '算法详解', link: '/algorithms' },
          ],
        },
        {
          text: 'AI Agent 接入',
          items: [
            { text: '接入指南', link: '/ai-agents' },
            { text: '一键提示词', link: '/prompts' },
          ],
        },
        {
          text: '参考',
          items: [
            { text: 'Go SDK API', link: '/sdk' },
            { text: 'CLI 命令', link: '/cli' },
            { text: 'MCP 工具', link: '/mcp' },
            { text: 'Skills 斜杠命令', link: '/skills' },
          ],
        },
      ],
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/scagogogo/versions-skills' },
    ],

    search: {
      provider: 'local',
    },

    footer: {
      message: '基于 MIT 协议发布',
      copyright: 'Copyright © 2023-2026 scagogogo',
    },

    outline: {
      label: '本页目录',
      level: [2, 3],
    },

    docFooter: {
      prev: '上一页',
      next: '下一页',
    },

    lastUpdated: {
      text: '最后更新于',
    },

    returnToTopLabel: '回到顶部',
    sidebarMenuLabel: '菜单',
  },

  // Mermaid 图表配置
  mermaid: {
    enableMermaid: true,
  },
})
)
