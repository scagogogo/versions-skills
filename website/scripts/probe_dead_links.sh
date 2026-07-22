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