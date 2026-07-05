#!/usr/bin/env bash
# 验证 internal/cli 包 100% 覆盖率。
#
# 策略：
#   1. 父级 `go test -coverprofile` 覆盖成功路径 + 非 os.Exit 路径。
#   2. 驱动二进制（go build -cover -covermode=atomic + GOCOVERDIR）覆盖 os.Exit 错误路径
#      ——go test 的测试二进制在 os.Exit 时不刷新覆盖率，因此 PrintResult 触发 os.Exit
#      的错误路径必须由独立构建的覆盖二进制（带运行时钩子）来收集。
#   3. `go tool covdata textfmt` 转换为文本格式，按块取最大值合并。
set -euo pipefail

cd "$(dirname "$0")/../../.."

GO=${GO:-go}
export GOTOOLCHAIN=local
GO_BIN=go1.25.7
ROOT=$(pwd)
WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT

echo "[1/5] 父级测试覆盖率"
$GO_BIN test -coverprofile="$WORK/parent.out" ./internal/cli/

echo "[2/5] 构建覆盖驱动二进制（含 os.Exit 路径钩子）"
$GO_BIN build -cover -covermode=atomic -tags cli_coverage_driver \
    -o "$WORK/driver" ./internal/cli/testdriver

echo "[3/5] 运行驱动（每个场景在带 GOCOVERDIR 的子进程中触发 os.Exit）"
mkdir -p "$WORK/gocoverdir"
CLI_DRIVER_GOCOVERDIR="$WORK/gocoverdir" "$WORK/driver" >/dev/null 2>&1 || true

echo "[4/5] 转换驱动覆盖率 + 合并"
$GO_BIN tool covdata textfmt -i="$WORK/gocoverdir" -o "$WORK/driver.out"
# 仅保留 internal/cli，排除 testdriver 自身
grep '^github.com/scagogogo/versions-skills/internal/cli/' "$WORK/parent.out" \
    | grep -v '/testdriver/' > "$WORK/parent_cli.out"
grep '^github.com/scagogogo/versions-skills/internal/cli/' "$WORK/driver.out" \
    | grep -v '/testdriver/' > "$WORK/driver_cli.out"

python3 - "$WORK/parent_cli.out" "$WORK/driver_cli.out" > "$WORK/merged.out" <<'PY'
import sys
def load(p):
    blocks={}; mode=None
    for line in open(p):
        line=line.rstrip('\n')
        if line.startswith('mode:'):
            mode=line.split(':',1)[1].strip(); continue
        parts=line.rsplit(' ',1)
        if len(parts)!=2: continue
        try: val=int(parts[1])
        except ValueError: continue
        blocks[parts[0]]=max(blocks.get(parts[0],0),val)
    return blocks,mode
a,ma=load(sys.argv[1]); b,mb=load(sys.argv[2])
out={}
for k in set(a)|set(b): out[k]=max(a.get(k,0),b.get(k,0))
print('mode:', ma or mb or 'atomic')
for k in sorted(out): print(f'{k} {out[k]}')
PY

echo "[5/5] 覆盖率结果"
$GO_BIN tool cover -func="$WORK/merged.out" | tail -3
cp "$WORK/merged.out" "$ROOT/cli_coverage_merged.out"
echo "合并覆盖率已写入 cli_coverage_merged.out"
