# 网站死链（404）根治实施计划

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 根治 https://scagogogo.github.io/versions-skills/ 上的 404 死链问题——根因是英文 locale（`/en/`）只有一页，VitePress 语言切换器运行时生成 375 个 `/en/<path>` 链接全部指向不存在的页面；同时让 CI 死链检查器能真正捕获此类运行时死链，防止复发。

**Architecture:** 分两层修复。① 根因层：删除 `website/.vitepress/config.ts` 中的 `en` locale 块。实测确认死链根因是 VitePress 语言切换器本身——只要 `en` locale 存在，每个中文页都会被渲染出 `/en/<当前路径>` 链接（对当前页自动加 locale 前缀），而英文站只有 `en/index.md` 一页，故 375 个 `/en/<path>` 全 404。实测还确认：仅改 en nav 项无效（切换器不读 nav），而移除整个 `en` locale 后 `en/index.md` 仍作为普通页面被构建为 `dist/en/index.html`（200），中文页不再渲染任何 `/en/<path>`。同时在中文 locale nav 加 `{ text: 'English', link: '/en/' }` 保留英文入口，并给 `en/index.md` frontmatter 加 `lang: en-US`（移除 locale 后否则继承站点默认 zh-CN）。② 工具层：增强 `website/scripts/check_dead_links.py`——删除"整体跳过 `/en/`"的掩盖逻辑，新增 locale 链接残留检测（扫描 dist HTML 里实际渲染的 `/en/<path>` 链接并验证对应产物存在），主动捕获语言切换器可能残留的死链。用户要求"自己安装 markdown 死链检查工具"——实测 `markdown-link-check` 与 VitePress cleanUrl/无扩展名链接约定不兼容（对 `/sdk/`、`./why` 一律报 400 误报），但 VitePress build 本身已在构建阶段做 markdown 死链检查并令 build 失败，故 markdown 源级检查复用 VitePress 原生能力，盲区（运行时 locale 死链）由增强后的 Python 脚本 + locale 残留检测补齐。

**Tech Stack:** VitePress 1.6.x（已装，含原生 markdown 死链检查），Node.js 22，Python 3（现有 `check_dead_links.py` 增强），GitHub Actions（`.github/workflows/deploy-website.yml`）

**Risks:**
- Task 2 移除 `en` locale 会移除顶部语言切换器下拉 → 缓解：在中文 nav 加 "English" 链接指向 `/en/`，保留英文入口；待英文内容补齐后可重新启用 locale（届时需配套补齐 en 页面）
- Task 2 移除 locale 后 `en/index.html` 的 `<html lang>` 会继承 zh-CN → 缓解：Step 3 在 `en/index.md` frontmatter 加 `lang: en-US`
- 增强后的 Python 脚本若 CI 环境无 Python 3 会失败 → 缓解：现有 workflow 已用 `python3 website/scripts/check_dead_links.py`，环境已具备
- 现有脚本跳过 `/en/` 是"故意的"（注释称英文站 WIP）→ 缓解：策略从"整体跳过"改为"逐条按产物判定 + locale 残留检测"，根因修复（Task 2）后 locale 残留检测报告为 0
- VitePress build 原生死链检查可能误拦合法链接 → 缓解：config 已有 `ignoreDeadLinks` 配置 github.com 等外链，维持现状

---

### Task 1: 全量探明死链清单并固化回归基线

**Depends on:** None
**Files:**
- Create: `website/scripts/probe_dead_links.sh`

- [ ] **Step 1: 创建线上死链探测脚本 — 批量 curl 探测所有 dist 路径的线上状态，产出可复现的 404 清单**

```bash
#!/usr/bin/env bash
# website/scripts/probe_dead_links.sh
# 探测线上 GitHub Pages 的真实 HTTP 状态，对照本地 dist 产物找出运行时 404。
# 用法：bash website/scripts/probe_dead_links.sh [--locale-only]
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
DIST="$REPO_ROOT/website/.vitepress/dist"
BASE_URL="https://scagogogo.github.io/versions-skills"

if [ ! -d "$DIST" ]; then
  echo "❌ dist 不存在，请先在 website/ 下执行 npm run build" >&2
  exit 1
fi

# 由 dist HTML 产物推出 cleanUrl 路径（去掉 .html，index.html → 目录）
paths_file="$(mktemp)"
dead_file="$(mktemp)"
trap 'rm -f "$paths_file" "$dead_file"' EXIT

if [ "${1:-}" = "--locale-only" ]; then
  # 仅探测 /en/ locale 对应路径（语言切换器会生成的链接）
  find "$DIST" -name "*.html" -not -path "*/en/*" \
    | sed "s|$DIST||" \
    | sed 's|/index\.html$|/|; s|\.html$||' \
    | sed 's|^|/en|; s|/$|/|' \
    | grep -v '/en/404$' \
    | sort -u > "$paths_file"
  echo "🔍 探测 /en/ locale 路径（语言切换器视角）..."
else
  find "$DIST" -name "*.html" \
    | sed "s|$DIST||" \
    | sed 's|/index\.html$|/|; s|\.html$||' \
    | grep -v '/404$' \
    | sort -u > "$paths_file"
  echo "🔍 探测全量路径..."
fi

total=$(wc -l < "$paths_file")
echo "待探测 URL 数: $total"
echo ""

cat "$paths_file" \
  | xargs -P 20 -I{} sh -c 'code=$(curl -s -o /dev/null -w "%{http_code}" "'"${BASE_URL}"'{}"); printf "%s %s\n" "$code" "{}"' \
  | sort > "$dead_file"

dead_count=$(grep -c '^404 ' "$dead_file" || true)
echo "=== 404 死链（共 $dead_count 条）==="
grep '^404 ' "$dead_file" || true
echo ""
non200=$(grep -cv '^200 ' "$dead_file" || true)
echo "非 200 总数: $non200"

# CI 语义：存在 404 即失败
if [ "$dead_count" -gt 0 ]; then
  exit 1
fi
exit 0
```

- [ ] **Step 2: 本地构建并运行探测，固化修复前的死链基线**

Run: `cd website && npm run build && cd .. && bash website/scripts/probe_dead_links.sh --locale-only; echo "EXIT=$?"`
Expected:
  - Exit code: 1（修复前应有 404）
  - Output contains: "404 /en/sdk/" 和 "404 /en/cli/" 等约 375 条
  - 这一步只固化基线，不要求通过

- [ ] **Step 3: 提交**
Run: `git add website/scripts/probe_dead_links.sh && git commit -m "chore(website): add live dead-link probe script for /en/ locale baseline"`

---

### Task 2: 移除 en locale 根因修复 — 消除语言切换器生成的逐页 /en/ 死链

**Depends on:** Task 1
**Files:**
- Modify: `website/.vitepress/config.ts:60-80`（删除整个 `en` locale 块）与 `website/.vitepress/config.ts:32-58`（中文 locale nav 加 English 入口）
- Modify: `website/en/index.md:1-3`（frontmatter 加 `lang: en-US`）

- [ ] **Step 1: 删除 config.ts 中的 en locale 块 — 从源头移除语言切换器对当前页做 locale 本地化的机制**

事实前提（Phase 1 实测确认，非假设）：死链根因不是 `en` locale 的 nav 配置，而是 **VitePress 语言切换器本身**——只要 `en` locale 存在，每个中文页（如 `/sdk/`）的 HTML 都会被渲染出一个"切到英文"的链接 `/en/sdk/`（对当前页路径自动加 locale 前缀）。英文站只有 `en/index.md` 一页，故 375 个 `/en/<path>` 全部 404。实测：仅改 en nav 项**无效**（语言切换器不读 nav，仍渲染 `/en/sdk/`）；实测：移除整个 `en` locale 后，中文 `sdk/index.html` 不再渲染任何 `/en/<path>` 死链，且 `en/index.md` 仍被当作普通页面构建出 `dist/en/index.html`（200）。

文件: `website/.vitepress/config.ts:60-80`（删除 `en` locale 整块，含其上方注释）

删除前（现状）：
```typescript
    // 英文（/en/ 前缀）
    'en': {
      label: 'English',
      lang: 'en-US',
      link: '/en/',
      themeConfig: {
        nav: [
          { text: 'Home', link: '/en/' },
          {
            text: 'Reference',
            items: [
              { text: 'Go SDK API', link: '/sdk/' },
              { text: 'CLI Commands', link: '/cli/' },
              { text: 'MCP Tools', link: '/mcp/' },
              { text: 'Skills', link: '/skills/' },
            ],
          },
          { text: '中文文档', link: '/' },
        ],
      },
    },
```

删除后（该块整段移除，`locales` 对象只剩 `root`）。`en/index.md` 不再挂在 locale 机制下，作为普通页面存在，不再触发语言切换器。

- [ ] **Step 2: 在中文 locale nav 加 English 入口 — 保留英文首页的可达入口**

文件: `website/.vitepress/config.ts:33-57`（`root` locale 的 `themeConfig.nav`，在 `{ text: 'AI Agent 接入', link: '/ai-agents' }` 之后、`{ text: '参考' }` 之前插入一项）

插入前（现状片段）：
```typescript
          { text: 'AI Agent 接入', link: '/ai-agents' },
          {
            text: '参考',
            items: [
              { text: 'Go SDK API', link: '/sdk/' },
```

插入后（加一项 English 指向英文首页）：
```typescript
          { text: 'AI Agent 接入', link: '/ai-agents' },
          { text: 'English', link: '/en/' },
          {
            text: '参考',
            items: [
              { text: 'Go SDK API', link: '/sdk/' },
```

说明：移除 `en` locale 后顶部不再有语言切换器下拉，故在 nav 显式加一个 "English" 链接指向 `/en/`（英文首页），保留英文入口可达。`/en/` 仍由 `en/index.md` 普通页面构建产出（200）。

- [ ] **Step 3: 给 en/index.md frontmatter 加 lang: en-US — 移除 locale 后保持英文页正确语言声明**

文件: `website/en/index.md:1-3`

事实前提（实测）：移除 `en` locale 后，`dist/en/index.html` 的 `<html lang>` 会继承站点默认 `zh-CN`，导致英文页声明为中文。需在 frontmatter 显式设 `lang: en-US`。

替换前（现状 frontmatter）：
```yaml
---
# https://vitepress.dev/reference/default-theme-home-page
layout: home
```

替换后：
```yaml
---
# https://vitepress.dev/reference/default-theme-home-page
layout: home
lang: en-US
```

- [ ] **Step 4: 在英文首页 en/index.md 的 Reference 区块补充说明文字 — 标注这些链接跳向中文站，防止后续误改成 /en/ 前缀**

文件: `website/en/index.md`（`## Reference` 区块，约文件末尾）

事实前提（实测）：`en/index.md` 正文里的 `[Go SDK API](/sdk/)` 是 markdown 正文链接，VitePress 不会对正文链接做 locale 本地化，指向 `/sdk/` 会正确跳到中文 SDK 页（线上 200，不 404）。本 Step 不改正文链接，只补一句说明。

替换前（现状）：
```markdown
## Reference

The full reference is in [简体中文](/), but the structure is language-neutral:

- [Go SDK API](/sdk/) — every type and function, with runnable examples
- [CLI Commands](/cli/) — all 44 subcommands
- [MCP Tools](/mcp/) — all 21 `version_*` tools
- [Skills](/skills/) — 13 Claude Code slash commands
```

替换后（仅追加一句说明，正文链接保持 `/sdk/` 不变）：
```markdown
## Reference

The full reference is in [简体中文](/), but the structure is language-neutral:

- [Go SDK API](/sdk/) — every type and function, with runnable examples
- [CLI Commands](/cli/) — all 44 subcommands
- [MCP Tools](/mcp/) — all 21 `version_*` tools
- [Skills](/skills/) — 13 Claude Code slash commands

> ⚠️ These links intentionally point to `/sdk/`, `/cli/` (Chinese site), NOT `/en/sdk/`. This English landing page is a standalone page (no VitePress locale), so body links are not localized — they stay on the Chinese site (HTTP 200). Adding an `/en/` prefix here would create 404s.
```

- [ ] **Step 5: 重新构建并验证死链根因消除（本地 dist 检查，不依赖线上）**

Run: `cd website && npm run build && echo "=== en 产物 ===" && find .vitepress/dist/en -name "*.html" && echo "=== 中文页是否仍渲染 /en/<path> 死链 ===" && python3 -c "import re,glob; hits=set(); [hits.update(re.findall(r'/versions-skills/en/[a-z]', open(f,encoding='utf-8').read())) for f in glob.glob('.vitepress/dist/**/*.html',recursive=True) if '/en/' not in f]; print('  死链链接:', sorted(hits) or '无——根因已消除')"`
Expected:
  - Exit code: 0
  - Output contains: `.vitepress/dist/en/index.html`（英文首页仍生成）
  - Output contains: "死链链接: 无——根因已消除"
  - `find .vitepress/dist/en` 仅列出 `index.html` 一个文件

- [ ] **Step 6: 提交**
Run: `git add website/.vitepress/config.ts website/en/index.md && git commit -m "fix(website): remove en locale to stop language switcher generating 375 /en/<path> 404s"`

---

### Task 3: 验证 VitePress 原生 markdown 死链检查覆盖源级链接 — 确认工具选型正确

**Depends on:** Task 2
**Files:**
- Create: `website/scripts/verify_md_links.sh`

- [ ] **Step 1: 创建 markdown 源级死链验证脚本 — 调用 VitePress build 原生检查并固化契约**

工具选型说明（基于实测，非假设）：用户要求"安装 markdown 死链检查工具"。实测发现 `markdown-link-check` 与 VitePress 链接约定不兼容——VitePress 用 cleanUrl（`/sdk/` 对应目录 `sdk/index.md`）和无扩展名相对链接（`./why` 对应 `why.md`），而 `markdown-link-check` 不懂这两种约定，对 `/sdk/` 和 `./why` 一律报 400 死链，产生海量误报。但实测同时确认：**VitePress build 本身已在 build 阶段做 markdown 死链检查**——注入一个 `/this-page-does-not-exist` 链接后 `npm run build` 直接失败并报 `1 dead link(s) found`。因此正确方案不是引入第三方 markdown 死链工具，而是**依赖 VitePress 原生 build 检查**，并用本脚本把这一契约固化（CI 里 build 失败即拦截）。

```bash
#!/usr/bin/env bash
# website/scripts/verify_md_links.sh
# 验证 VitePress 原生 markdown 死链检查生效：build 时任何指向不存在
# md 的内部链接都会让构建失败。本脚本封装该检查并提供清晰输出。
# 用法：bash website/scripts/verify_md_links.sh
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WEBSITE_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

echo "🔍 验证 VitePress 原生 markdown 死链检查（npm run build）..."
cd "$WEBSITE_DIR"

# build 失败 = 存在 markdown 死链（VitePress 报 "dead link(s) found"）
if npm run build 2>&1 | tee /tmp/vp_build.log; then
  if grep -q "dead link" /tmp/vp_build.log; then
    echo "❌ 检测到 markdown 死链（VitePress 报 dead link）" >&2
    exit 1
  fi
  echo "✅ markdown 源级内部链接全部有效（VitePress build 通过）"
  exit 0
else
  echo "❌ 构建失败——可能是 markdown 死链或构建错误，见上方日志" >&2
  exit 1
fi
```

- [ ] **Step 2: 运行验证脚本 — 确认现状 markdown 源级链接无死链**

Run: `bash website/scripts/verify_md_links.sh; echo "EXIT=$?"`
Expected:
  - Exit code: 0
  - Output contains: "✅ markdown 源级内部链接全部有效"
  - Output does NOT contain: "dead link(s) found"

- [ ] **Step 3: 反向验证 VitePress 能拦截注入的死链 — 临时注入死链确认 build 失败**

Run: `cd website && cp index.md /tmp/index.md.bak && printf '\n[dead-test](/this-page-does-not-exist)\n' >> index.md && bash scripts/verify_md_links.sh; ec=$?; cp /tmp/index.md.bak index.md && rm -rf .vitepress/dist && npm run build >/dev/null 2>&1; echo "EXIT=$ec"`
Expected:
  - 注入死链后 EXIT=1
  - Output contains: "dead link(s) found"
  - 还原 index.md 后重新 build 应通过

- [ ] **Step 4: 提交**
Run: `git add website/scripts/verify_md_links.sh && git commit -m "chore(website): pin VitePress native md dead-link check as source-level gate"`

---

### Task 4: 增强现有 check_dead_links.py — 移除 /en/ 整体跳过掩盖逻辑，新增 locale 链接残留检测

**Depends on:** Task 3
**Files:**
- Modify: `website/scripts/check_dead_links.py:46-56`（`normalize_href` 中跳过 `/en/` 的逻辑）与 `scan_dead_links`（新增 locale 残留检测）

- [ ] **Step 1: 移除 normalize_href 中"整体跳过 /en/"的掩盖逻辑 — 改为逐条按 dist 产物判定**

文件: `website/scripts/check_dead_links.py`（`normalize_href` 函数，删除下述整块跳过逻辑）

删除这段（现状）：
```python
    # English locale links (/en/...): the English site is intentionally partial
    # (only the homepage exists; other pages fall back to Chinese). Language-
    # switcher links generated by VitePress point to /en/<page> which may not
    # exist yet — these are expected, not dead links.
    if '/en/' in href or href.endswith('/en'):
        return None
```

替换为（不跳过，让正常逻辑判定死活）：
```python
    # 不再整体跳过 /en/。英文 locale 的死链由 dist 产物逐条判定：
    # /en/<page> 若 dist 里没有对应 HTML 就是真死链。语言切换器生成的
    # 运行时 /en/ 链接虽不在静态 <a href> 中（由 Task 4 Step 2 的
    # locale 残留检测单独校验），但源/nav 里出现的 /en/ 链接必须判活。
```

- [ ] **Step 2: 新增 locale 链接残留检测 — 扫描 dist HTML 里实际渲染的 /en/<path> 链接并验证对应产物存在**

设计说明：Task 2 已移除 `en` locale，语言切换器不再生成 `/en/<path>` 链接。本检测**不假设性地构造所有中文页的 /en/ 对应物**（那样会对每个中文页报一个误报），而是**扫描 dist HTML 里实际渲染出的 `/en/<非首页>` 链接**——只有当某个页面真的渲染了 `/en/sdk/` 这类链接（即 locale 被重新启用或被误加）时才报告。这样 Task 2 修复后该检测为 0（HTML 里无 `/en/<path>`），未来若误加 locale 则能立即抓到。

文件: `website/scripts/check_dead_links.py`（在 `scan_dead_links` 之后新增函数，并在 `main` 中调用）

在 `scan_dead_links` 函数定义之后、`main` 之前插入：

```python
def scan_locale_dead_links(dist_dir, base_path):
    """扫描 dist HTML 里实际渲染的 /en/<非首页> 链接，检查对应产物是否存在。
    VitePress 语言切换器对当前页做 locale 本地化会生成 /en/<path> 链接；
    若英文站无对应页面，运行时点击即 404。本函数只检查 HTML 中实际出现的
    这类链接（而非假设性构造），避免对未渲染的路径误报。"""
    dead = []
    root = Path(dist_dir)
    seen = set()
    # 匹配形如 /en/sdk、/en/cli/commands/parse 的链接（含 base 前缀或无）
    pattern = re.compile(r'href=["\'](?:' + re.escape(base_path) + r')?(/en/[a-z][a-z0-9/-]*)["\']')
    for html_file in root.rglob('*.html'):
        rel = os.path.relpath(html_file, root)
        try:
            src = open(html_file, encoding='utf-8').read()
        except Exception:
            continue
        for m in pattern.finditer(src):
            en_path = m.group(1)  # 形如 /en/sdk 或 /en/cli/commands/parse
            # 排除英文首页本身 /en/ 或 /en
            stripped = en_path.rstrip('/')
            if stripped == '/en':
                continue
            key = (rel, en_path)
            if key in seen:
                continue
            seen.add(key)
            # 构造 dist 内对应文件路径
            rel_part = en_path[len('/en/'):].rstrip('/')
            cand_dir = os.path.join(root, 'en', rel_part, 'index.html')
            cand_file = os.path.join(root, 'en', rel_part + '.html')
            if not (os.path.isfile(cand_dir) or os.path.isfile(cand_file)):
                dead.append((rel, en_path, ['en/' + rel_part + '/index.html',
                                             'en/' + rel_part + '.html']))
    return dead
```

在 `main` 函数中，`dead_links = scan_dead_links(...)` 之后新增调用与汇总：

```python
    dead_links = scan_dead_links(str(dist_dir), args.base)

    # locale 链接残留检测（语言切换器生成的 /en/<path>）
    locale_dead = scan_locale_dead_links(str(dist_dir), args.base)
    dead_links.extend(locale_dead)
```

并在输出区追加对 locale 死链的显式分类提示（在打印 dead_links 之前）：

```python
    locale_count = len(locale_dead)
    if locale_count:
        print(f"\n🌐 其中 {locale_count} 条是 locale 残留死链（页面渲染了 /en/<path> 但英文站无该页）：")
        print("   修复方式：确认 config.ts 未启用 en locale，或补齐 en/<path>.md 英文页。")
```

- [ ] **Step 3: 重新构建并运行增强后的脚本 — 验证 locale 残留检测在修复后为 0**

Run: `cd website && npm run build && cd .. && python3 website/scripts/check_dead_links.py --ci; echo "EXIT=$?"`
Expected:
  - Exit code: 0
  - Output contains: "✅ No internal dead links found"
  - Output does NOT contain: "locale 残留死链"（Task 2 已移除 en locale，HTML 不再渲染 /en/<path>，应为 0）

- [ ] **Step 4: 反向验证脚本能抓到人为注入的死链 — 临时改坏 config 再还原**

Run: `cd website && cp .vitepress/config.ts /tmp/config.ts.bak && sed -i "/{ text: 'English', link: '\\/en\\/' }/a\\          { text: 'Bad', link: '/en/sdk/' }," .vitepress/config.ts && npm run build >/dev/null 2>&1 && cd .. && python3 website/scripts/check_dead_links.py --ci; ec=$?; cp /tmp/config.ts.bak website/.vitepress/config.ts && echo "EXIT=$ec"`
Expected:
  - 临时注入 `/en/sdk/` nav 项后，EXIT=1
  - Output contains: "/en/sdk/" 相关死链报告
  - 还原配置后再次构建应为绿色

- [ ] **Step 5: 提交**
Run: `git add website/scripts/check_dead_links.py && git commit -m "fix(website): remove /en/ blanket-skip mask; add locale link residual detection"`

---

### Task 5: 将新增检查接入 CI — deploy-website.yml 增加双工具门禁

**Depends on:** Task 4
**Files:**
- Modify: `.github/workflows/deploy-website.yml:40-44`（"Check dead links" step 附近）

- [ ] **Step 1: 修改 deploy-website.yml — 强化现有死链检查 step 的语义标注（locale 残留检测已并入该脚本）**

文件: `.github/workflows/deploy-website.yml`（定位 `Check dead links` step，约第 40-44 行）

工具覆盖关系说明（基于实测）：
- **markdown 源级死链**：由 workflow 已有的 `Build website` step（`npm run build`）隐式覆盖——VitePress build 检测到 markdown 死链会直接失败（Task 3 已验证）。无需额外 step。
- **dist HTML 静态死链 + locale 残留检测**：由 Task 4 增强后的 `check_dead_links.py` 覆盖（移除了 `/en/` 跳过掩盖、新增 locale 残留检测）。
- 因此 CI 层只需把现有 step 的语义标注更新为"含 locale 残留检测"，不新增 step，避免重复 build。

现状：
```yaml
      - name: Check dead links
        run: python3 website/scripts/check_dead_links.py --ci
```

替换为：
```yaml
      - name: Check dead links (dist HTML + en-locale residual detection)
        run: python3 website/scripts/check_dead_links.py --ci
```

- [ ] **Step 2: 本地模拟 CI 流程验证门禁通过**

Run: `cd website && npm run build && cd .. && python3 website/scripts/check_dead_links.py --ci; echo "EXIT=$?"`
Expected:
  - Exit code: 0
  - Output contains: "✅ No internal dead links found"
  - Output does NOT contain: "locale 残留死链"（Task 2 修复后应为 0）
  - build 阶段不报 "dead link(s) found"

- [ ] **Step 3: 提交**
Run: `git add .github/workflows/deploy-website.yml && git commit -m "ci(website): clarify dead-link step covers en-locale residual detection"`

---
