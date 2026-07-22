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