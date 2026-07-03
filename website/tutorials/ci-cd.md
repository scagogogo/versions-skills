# CI/CD 中的版本判断

在 GitHub Actions 等 CI 环境中用 versions CLI 做版本判断与门禁。

## 📦 在 CI 中安装

```yaml
# .github/workflows/release-check.yml
jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v6
        with:
          go-version: '1.25'
      - run: go install github.com/scagogogo/versions-skills/cmd/versions@latest
```

## 🚦 exit code 门禁

`check` 命令以 exit code 返回布尔结果（0=真/1=假），天然适合 CI 条件判断：

```bash
# 仅当版本是预发布时继续
if versions check 1.2.3-rc1 --prerelease; then
  echo "预发布版本，跳过正式发布流程"
  exit 0
fi
```

## 📋 典型场景

### 1. 拒绝降级

```bash
NEW_VERSION="1.5.0"
LAST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "0.0.0")

if versions check "$NEW_VERSION" --newer "$LAST_TAG"; then
  echo "✅ $NEW_VERSION 新于 $LAST_TAG"
else
  echo "❌ 拒绝降级：$NEW_VERSION 不新于 $LAST_TAG"
  exit 1
fi
```

### 2. 约束兼容性检查

```bash
# 检查依赖版本是否满足 ^1.2.0
if versions constraint "^1.2.0" "$DEP_VERSION"; then
  echo "兼容"
else
  echo "不兼容"
  exit 1
fi
```

### 3. 排序 changelog 中的版本

```bash
versions sort --from-file versions-in-changelog.txt
```

### 4. 找最新稳定版做发布基线

```bash
versions latest-stable --from-file published-versions.txt
```

## 🔗 GitHub Pages 部署

本文档站本身就用 GitHub Actions + GitHub Pages 部署，CI 配置在 `.github/workflows/`，构建 VitePress 后部署到 `/versions-skills/` 子路径。

## 🚀 下一步

- [让 Claude 管理版本](/tutorials/ai-version-management)
- CLI：[`check`](/cli/commands/check) · [`constraint`](/cli/commands/constraint) · [`sort`](/cli/commands/sort)
- 入门：[快速开始](/quick-start)
