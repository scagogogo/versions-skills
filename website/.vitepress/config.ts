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

  locales: {
    // 默认语言：简体中文（根路径，无前缀）
    root: {
      label: '简体中文',
      lang: 'zh-CN',
      themeConfig: {
        nav: [
          { text: '首页', link: '/' },
          { text: '为什么用', link: '/why' },
          {
            text: '学习',
            items: [
              { text: '核心概念', link: '/concepts/' },
              { text: '教程', link: '/tutorials/' },
              { text: '实用配方', link: '/recipes/' },
              { text: '可运行示例', link: '/examples/' },
              { text: '工作原理', link: '/how-it-works' },
              { text: '算法详解', link: '/algorithms' },
            ],
          },
          { text: 'AI Agent 接入', link: '/ai-agents' },
          { text: 'English', link: '/en/' },
          {
            text: '参考',
            items: [
              { text: 'Go SDK API', link: '/sdk/' },
              { text: 'CLI 命令', link: '/cli/' },
              { text: 'MCP 工具', link: '/mcp/' },
              { text: 'Skills 斜杠命令', link: '/skills/' },
            ],
          },
        ],
      },
    },
    },

  themeConfig: {
    // 站点级导航已移至 locales 中按语言配置
    // 侧边栏：按路径分目录，每部分独立侧边栏
    sidebar: {
      '/sdk/': [
        {
          text: 'SDK 总览',
          items: [
            { text: 'API 索引', link: '/sdk/' },
          ],
        },
        {
          text: '核心类型',
          collapsed: true,
          items: [
            { text: 'Version', link: '/sdk/api/version' },
            { text: 'VersionNumbers', link: '/sdk/api/version-numbers' },
            { text: 'VersionPrefix', link: '/sdk/api/version-prefix' },
            { text: 'VersionSuffix', link: '/sdk/api/version-suffix' },
            { text: 'SuffixWeight', link: '/sdk/api/suffix-weight' },
            { text: 'VersionBuilder', link: '/sdk/api/version-builder' },
            { text: 'VersionSlice', link: '/sdk/api/version-slice' },
            { text: 'VersionRange', link: '/sdk/api/version-range' },
            { text: 'VersionDiff', link: '/sdk/api/version-diff' },
            { text: 'VersionGroup', link: '/sdk/api/version-group' },
            { text: 'SortedVersionGroups', link: '/sdk/api/sorted-version-groups' },
            { text: 'ContainsPolicy', link: '/sdk/api/contains-policy' },
          ],
        },
        {
          text: '约束类型',
          collapsed: true,
          items: [
            { text: 'Constraint', link: '/sdk/api/constraint' },
            { text: 'ConstraintOperator', link: '/sdk/api/constraint-operator' },
            { text: 'ConstraintSet', link: '/sdk/api/constraint-set' },
            { text: 'ConstraintUnion', link: '/sdk/api/constraint-union' },
          ],
        },
        {
          text: '能力域',
          collapsed: true,
          items: [
            { text: '解析构造', link: '/sdk/parsing' },
            { text: '版本属性', link: '/sdk/properties' },
            { text: '类型判断', link: '/sdk/predicates' },
            { text: '比较', link: '/sdk/comparison' },
            { text: '排序', link: '/sdk/sorting' },
            { text: '分组', link: '/sdk/grouping' },
            { text: '过滤', link: '/sdk/filtering' },
            { text: '集合运算', link: '/sdk/set-operations' },
            { text: '约束', link: '/sdk/constraints' },
            { text: '范围查询', link: '/sdk/range-query' },
            { text: '不可变变更', link: '/sdk/mutation' },
            { text: '文件 IO', link: '/sdk/file-io' },
            { text: '可视化', link: '/sdk/visualization' },
            { text: '序列化', link: '/sdk/serialization' },
          ],
        },
      ],

      '/cli/': [
        {
          text: 'CLI 总览',
          items: [{ text: '命令索引', link: '/cli/' }],
        },
        {
          text: '解析与信息',
          collapsed: true,
          items: [
            { text: 'parse', link: '/cli/commands/parse' },
            { text: 'info', link: '/cli/commands/info' },
            { text: 'validate', link: '/cli/commands/validate' },
            { text: 'segments', link: '/cli/commands/segments' },
            { text: 'sub-version', link: '/cli/commands/sub-version' },
            { text: 'suffix-weight', link: '/cli/commands/suffix-weight' },
            { text: 'pure-prefix', link: '/cli/commands/pure-prefix' },
            { text: 'group-id', link: '/cli/commands/group-id' },
            { text: 'core', link: '/cli/commands/core' },
            { text: 'clone', link: '/cli/commands/clone' },
          ],
        },
        {
          text: '比较与检查',
          collapsed: true,
          items: [
            { text: 'compare', link: '/cli/commands/compare' },
            { text: 'check', link: '/cli/commands/check' },
          ],
        },
        {
          text: '排序与极值',
          collapsed: true,
          items: [
            { text: 'sort', link: '/cli/commands/sort' },
            { text: 'sort-strings', link: '/cli/commands/sort-strings' },
            { text: 'min', link: '/cli/commands/min' },
            { text: 'max', link: '/cli/commands/max' },
            { text: 'latest-stable', link: '/cli/commands/latest-stable' },
            { text: 'latest-prerelease', link: '/cli/commands/latest-prerelease' },
          ],
        },
        {
          text: '分组与过滤',
          collapsed: true,
          items: [
            { text: 'group', link: '/cli/commands/group' },
            { text: 'group-ids', link: '/cli/commands/group-ids' },
            { text: 'group-latest', link: '/cli/commands/group-latest' },
            { text: 'group-oldest', link: '/cli/commands/group-oldest' },
            { text: 'group-stable', link: '/cli/commands/group-stable' },
            { text: 'group-prerelease', link: '/cli/commands/group-prerelease' },
            { text: 'group-latest-stable', link: '/cli/commands/group-latest-stable' },
            { text: 'group-latest-prerelease', link: '/cli/commands/group-latest-prerelease' },
            { text: 'group-contains', link: '/cli/commands/group-contains' },
            { text: 'filter', link: '/cli/commands/filter' },
            { text: 'partition', link: '/cli/commands/partition' },
            { text: 'count', link: '/cli/commands/count' },
          ],
        },
        {
          text: '约束与范围',
          collapsed: true,
          items: [
            { text: 'constraint', link: '/cli/commands/constraint' },
            { text: 'satisfies', link: '/cli/commands/satisfies' },
            { text: 'range', link: '/cli/commands/range' },
          ],
        },
        {
          text: '变更与构造',
          collapsed: true,
          items: [
            { text: 'build', link: '/cli/commands/build' },
            { text: 'bump', link: '/cli/commands/bump' },
            { text: 'set-prefix', link: '/cli/commands/set-prefix' },
            { text: 'set-suffix', link: '/cli/commands/set-suffix' },
            { text: 'set-major', link: '/cli/commands/set-major' },
            { text: 'set-minor', link: '/cli/commands/set-minor' },
            { text: 'set-patch', link: '/cli/commands/set-patch' },
            { text: 'set-numbers', link: '/cli/commands/set-numbers' },
          ],
        },
        {
          text: '文件与可视化',
          collapsed: true,
          items: [
            { text: 'read', link: '/cli/commands/read' },
            { text: 'write', link: '/cli/commands/write' },
            { text: 'read-strings', link: '/cli/commands/read-strings' },
            { text: 'visualize', link: '/cli/commands/visualize' },
          ],
        },
      ],

      '/mcp/': [
        {
          text: 'MCP 总览',
          items: [{ text: '工具索引', link: '/mcp/' }],
        },
        {
          text: '解析与信息',
          collapsed: true,
          items: [
            { text: 'version_parse', link: '/mcp/tools/version-parse' },
            { text: 'version_validate', link: '/mcp/tools/version-validate' },
            { text: 'version_info', link: '/mcp/tools/version-info' },
            { text: 'version_core', link: '/mcp/tools/version-core' },
          ],
        },
        {
          text: '比较',
          collapsed: true,
          items: [{ text: 'version_compare', link: '/mcp/tools/version-compare' }],
        },
        {
          text: '排序与极值',
          collapsed: true,
          items: [
            { text: 'version_sort', link: '/mcp/tools/version-sort' },
            { text: 'version_min', link: '/mcp/tools/version-min' },
            { text: 'version_max', link: '/mcp/tools/version-max' },
            { text: 'version_latest_stable', link: '/mcp/tools/version-latest-stable' },
            { text: 'version_latest_prerelease', link: '/mcp/tools/version-latest-prerelease' },
          ],
        },
        {
          text: '分组与过滤',
          collapsed: true,
          items: [
            { text: 'version_group', link: '/mcp/tools/version-group' },
            { text: 'version_filter', link: '/mcp/tools/version-filter' },
            { text: 'version_unique', link: '/mcp/tools/version-unique' },
            { text: 'version_set_operation', link: '/mcp/tools/version-set-operation' },
          ],
        },
        {
          text: '约束与范围',
          collapsed: true,
          items: [
            { text: 'version_constraint_check', link: '/mcp/tools/version-constraint-check' },
            { text: 'version_range_query', link: '/mcp/tools/version-range-query' },
          ],
        },
        {
          text: '变更与构造',
          collapsed: true,
          items: [
            { text: 'version_build', link: '/mcp/tools/version-build' },
            { text: 'version_bump', link: '/mcp/tools/version-bump' },
          ],
        },
        {
          text: '文件与可视化',
          collapsed: true,
          items: [
            { text: 'version_read_file', link: '/mcp/tools/version-read-file' },
            { text: 'version_write_file', link: '/mcp/tools/version-write-file' },
            { text: 'version_visualize', link: '/mcp/tools/version-visualize' },
          ],
        },
      ],

      '/skills/': [
        {
          text: 'Skills 斜杠命令',
          items: [
            { text: 'installation', link: '/skills/installation' },
            { text: 'version-parsing', link: '/skills/version-parsing' },
            { text: 'version-comparison', link: '/skills/version-comparison' },
            { text: 'version-check', link: '/skills/version-check' },
            { text: 'version-constraints', link: '/skills/version-constraints' },
            { text: 'version-sorting', link: '/skills/version-sorting' },
            { text: 'version-grouping', link: '/skills/version-grouping' },
            { text: 'version-range-query', link: '/skills/version-range-query' },
            { text: 'version-mutation', link: '/skills/version-mutation' },
            { text: 'version-properties', link: '/skills/version-properties' },
            { text: 'version-file-operations', link: '/skills/version-file-operations' },
            { text: 'version-visualization', link: '/skills/version-visualization' },
            { text: 'cli-operations', link: '/skills/cli-operations' },
            { text: 'mcp-operations', link: '/skills/mcp-operations' },
          ],
        },
      ],

      '/concepts/': [
        {
          text: '核心概念',
          items: [{ text: '概念索引', link: '/concepts/' }],
        },
        {
          text: '专题',
          collapsed: false,
          items: [
            { text: '版本号结构', link: '/concepts/version-anatomy' },
            { text: '后缀权重', link: '/concepts/suffix-weight' },
            { text: '比较优先级', link: '/concepts/compare-priority' },
            { text: '分组语义', link: '/concepts/grouping' },
            { text: '约束表达式', link: '/concepts/constraints' },
            { text: '范围与包含策略', link: '/concepts/range-and-policy' },
            { text: '不可变性', link: '/concepts/immutability' },
            { text: 'SemVer 规范', link: '/concepts/semver' },
            { text: '文件格式', link: '/concepts/file-format' },
            { text: '序列化', link: '/concepts/serialization' },
            { text: '三层接入', link: '/concepts/three-layers' },
            { text: '零依赖设计', link: '/concepts/zero-deps' },
          ],
        },
      ],

      '/tutorials/': [
        {
          text: '教程',
          items: [{ text: '教程索引', link: '/tutorials/' }],
        },
        {
          text: '入门',
          collapsed: false,
          items: [
            { text: '10 分钟入门', link: '/tutorials/getting-started' },
            { text: '解析与检查', link: '/tutorials/parse-and-check' },
            { text: '排序与极值', link: '/tutorials/sort-and-minmax' },
            { text: '分组与聚合', link: '/tutorials/grouping' },
          ],
        },
        {
          text: '进阶',
          collapsed: false,
          items: [
            { text: '约束表达式实战', link: '/tutorials/constraints-in-practice' },
            { text: '范围查询', link: '/tutorials/range-query' },
            { text: '变更与发布流程', link: '/tutorials/bump-and-release' },
            { text: '文件批处理', link: '/tutorials/file-batch' },
            { text: '可视化与报告', link: '/tutorials/visualization' },
          ],
        },
        {
          text: 'AI 集成',
          collapsed: false,
          items: [
            { text: '让 Claude 管理版本', link: '/tutorials/ai-version-management' },
            { text: 'CI/CD 中的版本判断', link: '/tutorials/ci-cd' },
          ],
        },
      ],

      '/recipes/': [
        {
          text: '实用配方',
          items: [{ text: '配方索引', link: '/recipes/' }],
        },
      ],

      '/examples/': [
        {
          text: '可运行示例',
          items: [{ text: '示例索引', link: '/examples/' }],
        },
      ],

      // 默认（根页面）
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
          text: '学习',
          items: [
            { text: '核心概念', link: '/concepts/' },
            { text: '教程', link: '/tutorials/' },
            { text: '实用配方', link: '/recipes/' },
            { text: '可运行示例', link: '/examples/' },
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
            { text: 'Go SDK API', link: '/sdk/' },
            { text: 'CLI 命令', link: '/cli/' },
            { text: 'MCP 工具', link: '/mcp/' },
            { text: 'Skills 斜杠命令', link: '/skills/' },
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
