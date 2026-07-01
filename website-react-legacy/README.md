# Versions-Skills Website

The marketing & documentation site for [versions-skills](https://github.com/scagogogo/versions-skills) — an AI-native version-number toolkit (Skills + MCP + Go SDK + CLI).

Deployed at **https://scagogogo.github.io/versions-skills/** via GitHub Pages (build artifacts committed under `dist/` and published by a GitHub Action).

## Stack

- **React 19** + **TypeScript 6** + **Vite 8** (Rolldown bundler)
- **Ant Design 6** component library
- **react-router-dom 7** (single-page, `basename="/versions-skills"`)
- **oxlint** for linting
- Flat in-house design system (blue-gray palette) — see `src/index.css`

## Project Structure

```
website/
├── src/
│   ├── App.tsx                 # Router, single HomePage route
│   ├── main.tsx                # Entry point
│   ├── index.css               # Design tokens + shared classes (.section-title, .flat-card …)
│   ├── pages/
│   │   └── HomePage.tsx        # Section composition
│   └── components/
│       ├── SiteHeader.tsx      # Sticky nav + GitHub button
│       ├── HeroSection.tsx     # Hero banner
│       ├── FeaturesSection.tsx # 12 core capabilities grid
│       ├── AlgorithmsSection.tsx # ★ Algorithm deep-dive (parse / compare / suffix weight / constraints / range / sort+group)
│       ├── ArchitectureSection.tsx # 4-layer architecture diagram + 13 skills
│       ├── AccessSection.tsx   # 4 access methods (Skills / SDK / CLI / MCP)
│       ├── CasesSection.tsx    # Use cases
│       ├── TutorialsSection.tsx
│       ├── AiIntegrationSection.tsx # Per-agent MCP config (Claude Code / Codex / Cursor / Windsurf / Cline / VS Code)
│       ├── QuickStartSection.tsx
│       └── SiteFooter.tsx
├── public/
│   └── favicon.svg
├── index.html
└── vite.config.ts              # base: '/versions-skills/'
```

The page is composed of vertical sections, each with an `id` used as a hash anchor by the header nav: `#features`, `#algorithms`, `#access`, `#cases`, `#tutorials`, `#ai-integration`, `#quickstart`.

## Development

```bash
# from the website/ directory
npm install        # first time only
npm run dev        # Vite dev server with HMR → http://localhost:5173
npm run build      # type-check (tsc -b) + production build → dist/
npm run preview    # serve the production build locally
npm run lint       # oxlint
```

> The dev server runs at a root path, but the production build uses `base: '/versions-skills/'` so assets resolve under the GitHub Pages subpath. Internal links go through `react-router` with `basename="/versions-skills"`.

## Editing Content

All copy lives in the section components under `src/components/` — there is no CMS. To change wording, edit the relevant `*Section.tsx` and its data arrays (e.g. `features`, `aiClients`, `suffixWeights`). The algorithm tables in `AlgorithmsSection.tsx` mirror the Go source in the repo root (`parser.go`, `version.go`, `suffix_weight.go`, `constraint.go`, `version_range.go`, `sort.go`) — keep them in sync when the Go semantics change.

## Design Notes

- **Flat design** — no shadows, no gradients; `1px solid #e2e8f0` borders, `4px` radius.
- Palette is CSS custom properties in `src/index.css` (`--c-primary` … `--c-gray-900`).
- Code blocks use the dark slate theme (`background: #1e293b; color: #e2e8f0`).
- Shared classes: `.section-title` (32px / 700), `.section-subtitle` (16px / gray-500), `.flat-card` (hover lifts border to gray-300).

## CI / Deployment

Pushing to `main` triggers the Pages workflow, which runs `npm run build` and uploads `dist/`. See `.github/workflows/` in the repo root.
