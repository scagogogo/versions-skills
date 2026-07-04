# Versions-Skills Website

The official documentation site for [versions-skills](https://github.com/scagogogo/versions-skills), built with [VitePress](https://vitepress.dev/).

Deployed at **https://scagogogo.github.io/versions-skills/** via GitHub Pages.

> The previous React site is preserved under `website-react-legacy/` for reference. This directory is now the VitePress implementation.

## Stack

- **VitePress 1.x** — Vue-powered static site generator, Markdown-first.
- No framework runtime in the browser beyond VitePress's own.

## Structure

```
website/
├── .vitepress/
│   ├── config.ts          # Nav, sidebar, theme, base path
│   └── dist/              # Build output (gitignored), uploaded by CI
├── public/
│   └── favicon.svg
├── index.md               # Home page (Hero + copy-paste AI prompts)
├── why.md                 # What problem it solves
├── how-it-works.md        # Four-layer architecture
├── algorithms.md          # Algorithm deep-dive
├── ai-agents.md           # Per-agent integration guide
├── prompts.md             # ★ One-click copy-paste prompts for Claude Code / Codex
├── quick-start.md
├── sdk.md                 # Go SDK API
├── cli.md                 # CLI commands
├── mcp.md                 # MCP tools
├── skills.md              # Claude Code slash commands
└── package.json
```

## Development

```bash
# from the website/ directory
npm install
npm run dev        # dev server → http://localhost:5173/versions-skills/
npm run build      # production build → .vitepress/dist/
npm run preview    # preview the build
```

## Editing Content

All content is Markdown — no CMS, no components to write. Edit the `.md` files directly. The `algorithms.md` tables mirror the Go source in the repo root (`parser.go`, `version.go`, `suffix_weight.go`, `constraint.go`, `version_range.go`, `sort.go`) — keep them in sync when Go semantics change.

Nav and sidebar are configured in `.vitepress/config.ts`. The `base: '/versions-skills/'` setting makes assets resolve under the GitHub Pages subpath.

## CI / Deployment

[`deploy-website.yml`](https://github.com/scagogogo/versions-skills/blob/main/.github/workflows/deploy-website.yml) runs on every push to `main` that touches `website/**`:

1. `npm ci` + `npm run build`
2. Uploads `website/.vitepress/dist` as a Pages artifact
3. Deploys to GitHub Pages
