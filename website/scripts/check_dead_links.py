#!/usr/bin/env python3
# coding: utf-8
"""
Scan VitePress dist directory for internal dead links.
Run after npm run build to verify all internal hrefs resolve correctly.

Usage:
    python scripts/check_dead_links.py [--dist DIR] [--base PATH]

Exit codes:
    0 - No dead links found
    1 - Dead links detected (CI will fail)
"""

import os
import re
import html
import argparse
import sys
from pathlib import Path


def normalize_href(href, base_path):
    """Convert href to relative path(s) within dist.
    - Remove base path prefix
    - Handle cleanUrls: /foo → /foo.html or /foo/index.html
    """
    if not href or href.startswith('#') or href.startswith('mailto:'):
        return None

    # External links (http/https) - not checked
    if re.match(r'^https?://', href):
        return None

    # GitHub repository links - not checked (valid external)
    if href.startswith('https://github.com/scagogogo/versions-skills'):
        return None

    # English locale links (/en/...): 不再整体跳过。英文 locale 的死链由 dist
    # 产物逐条判定：/en/<page> 若 dist 里没有对应 HTML 就是真死链。运行时语言
    # 切换器生成的 /en/ 链接虽不在静态 <a href> 中（由 scan_locale_dead_links
    # 单独检测），但源/nav 里出现的 /en/ 链接必须判活。
    # 历史原因：之前整块跳过 /en/ 掩盖了 375 个语言切换器生成的死链。详见:
    # docs/superpowers/plans/2026-07-13-website-deadlinks-fix.md

    # Remove base path
    if base_path and href.startswith(base_path):
        href = href[len(base_path):]

    if not href.startswith('/'):
        href = '/' + href

    # cleanUrls: /foo → /foo.html or /foo/index.html
    if not re.search(r'\.\w+$', href):  # No file extension
        html_path = href + '.html'
        index_path = href.rstrip('/') + '/index.html'
        return [html_path.lstrip('/'), index_path.lstrip('/')]

    return [href.lstrip('/')]


def check_file(relpaths, dist_dir):
    """Check if any of relpaths exists as a file."""
    for rp in relpaths:
        fp = os.path.join(dist_dir, rp)
        if os.path.isfile(fp):
            return True
    return False


def scan_dead_links(dist_dir, base_path):
    """Scan all HTML files in dist_dir for internal dead links."""
    dead_links = []
    checked = set()

    if not os.path.isdir(dist_dir):
        print(f"❌ dist directory not found: {dist_dir}")
        return dead_links

    for html_file in Path(dist_dir).rglob('*.html'):
        rel_html = os.path.relpath(html_file, dist_dir)
        try:
            src = open(html_file, encoding='utf-8').read()
        except Exception as e:
            print(f"⚠️ Failed to read {rel_html}: {e}")
            continue

        # Extract all <a href="...">
        for m in re.finditer(r'<a[^>]*href=["\']([^"\']+)["\']', src):
            href = html.unescape(m.group(1))
            key = (rel_html, href)
            if key in checked:
                continue
            checked.add(key)

            relpaths = normalize_href(href, base_path)
            if relpaths is None:
                continue  # External/anchor/empty

            if not check_file(relpaths, dist_dir):
                dead_links.append((rel_html, href, relpaths))

    return dead_links


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


def main():
    parser = argparse.ArgumentParser(description='Check internal dead links in VitePress dist')
    parser.add_argument('--dist', default='website/.vitepress/dist', help='dist directory path')
    parser.add_argument('--base', default='/versions-skills', help='GitHub Pages base path')
    parser.add_argument('--ci', action='store_true', help='CI mode: exit 1 on dead links')
    args = parser.parse_args()

    # Resolve dist path relative to repo root
    script_dir = Path(__file__).parent
    repo_root = script_dir.parent.parent
    dist_dir = (repo_root / args.dist).resolve()

    print(f"🔍 Scanning dist: {dist_dir}")
    print(f"📍 Base path: {args.base}")

    dead_links = scan_dead_links(str(dist_dir), args.base)

    # locale 链接残留检测（语言切换器生成的 /en/<path>）
    locale_dead = scan_locale_dead_links(str(dist_dir), args.base)
    dead_links.extend(locale_dead)

    if not dead_links:
        print("✅ No internal dead links found")
        return 0

    locale_count = len(locale_dead)
    regular_count = len(dead_links) - locale_count

    print(f"\n⚠️ Found {len(dead_links)} internal dead links:\n")
    if regular_count:
        print(f"📄 常规死链 {regular_count} 条：")
        shown = 0
        for src, href, targets in dead_links:
            if 'locale-switch' not in href and not href.startswith('/en/'):
                if shown >= 50:
                    print(f"  ... and {regular_count - shown} more")
                    break
                print(f"  {src} → {href}")
                print(f"    Targets not found: {targets}")
                shown += 1

    if locale_count:
        print(f"\n🌐 其中 {locale_count} 条是 locale 残留死链（页面渲染了 /en/<path> 但英文站无该页）：")
        print("   修复方式：确认 config.ts 未启用 en locale，或补齐 en/<path>.md 英文页。")
        for src, href, targets in dead_links:
            if 'locale-switch' in href or href.startswith('/en/'):
                print(f"  {src} → {href}")

    if args.ci:
        return 1

    return 0


if __name__ == '__main__':
    sys.exit(main())