#!/usr/bin/env bash
# 验证全项目单元测试覆盖率 100%。
#
# 策略：
#   1. 各包 `go test -coverprofile` 收集父级覆盖率（成功路径 + 非 os.Exit 路径）。
#   2. internal/cli 的 os.Exit 路径 + cmd 的 main() 路径，通过 `go build -cover` 二进制
#      在带 GOCOVERDIR 的子进程中触发 os.Exit 时刷新覆盖率（go test 二进制不刷新）。
#   3. `go tool covdata textfmt` 转换 + 按块取最大值合并。
set -euo pipefail

cd "$(dirname "$0")/.."

export GOTOOLCHAIN=local
GO_BIN=go1.25.7
ROOT=$(pwd)
WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT

echo "[1/6] 各包父级覆盖率"
$GO_BIN test -coverprofile="$WORK/core.out" -coverpkg=github.com/scagogogo/versions-skills .
$GO_BIN test -coverprofile="$WORK/verfmt.out" ./internal/verfmt/
$GO_BIN test -coverprofile="$WORK/mcp.out" ./internal/mcp/
$GO_BIN test -coverprofile="$WORK/cli.out" ./internal/cli/
$GO_BIN test -coverprofile="$WORK/cmdv.out" ./cmd/versions/
$GO_BIN test -coverprofile="$WORK/cmdm.out" ./cmd/versions-mcp/

echo "[2/6] internal/cli os.Exit 路径覆盖率"
$GO_BIN build -cover -covermode=atomic -tags cli_coverage_driver \
    -o "$WORK/cli_driver" ./internal/cli/testdriver
mkdir -p "$WORK/cli_cover"
CLI_DRIVER_GOCOVERDIR="$WORK/cli_cover" "$WORK/cli_driver" >/dev/null 2>&1 || true
$GO_BIN tool covdata textfmt -i="$WORK/cli_cover" -o "$WORK/cli_driver.out"

echo "[3/6] cmd 子进程覆盖率（cmd/versions + cmd/versions-mcp 的 main 路径）"
# cmd 的测试已在父级 TestRun_Subprocess 中用 go build -cover 构建子进程并跑，
# 但 GOCOVERDIR 数据在 t.TempDir() 里，父级 go test 收不到。
# 这里重新构建并跑一遍，收集到固定目录。
mkdir -p "$WORK/cmd_cover"
build_cover_bin() {
    local pkg="$1" out="$2"
    $GO_BIN build -cover -covermode=atomic -o "$out" "$pkg"
}
build_cover_bin ./cmd/versions "$WORK/versions_bin"
build_cover_bin ./cmd/versions-mcp "$WORK/versions_mcp_bin"

run_cover() {
    local bin="$1"; shift
    "$bin" "$@" >/dev/null 2>&1 || true
}
# cmd/versions 场景
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_bin" parse 1.2.3
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_bin" --help
# 未知命令 → cobra 返回 error → cli.Execute 返回 error → run() return 1 → main os.Exit(1)
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_bin" unknown-cmd-xyz
# cmd/versions-mcp 场景
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_mcp_bin" --help
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_mcp_bin" --transport stdio < /dev/null
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_mcp_bin" --transport bad-xyz
# 未知 flag → cobra 返回 error → Execute 返回 error → run() return 1
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_mcp_bin" --unknown-flag-xyz
# sse 端口冲突 → ServeSSE.Start 返回 error → log.Fatalf
GOCOVERDIR="$WORK/cmd_cover" "$WORK/versions_mcp_bin" --transport sse --port 19876 >/dev/null 2>&1 &
SSE_PID=$!
sleep 1
GOCOVERDIR="$WORK/cmd_cover" run_cover "$WORK/versions_mcp_bin" --transport sse --port 19876
kill $SSE_PID 2>/dev/null || true
# stdio + SIGINT → ServeStdio 因 ctx cancel 返回 error → log.Fatalf
# sleep 保持 stdin 打开，否则 stdin EOF 让 ServeStdio 提前返回 nil 退出，SIGINT 无的放矢
sleep 5 | GOCOVERDIR="$WORK/cmd_cover" "$WORK/versions_mcp_bin" --transport stdio >/dev/null 2>&1 &
STDIO_PID=$!
sleep 0.5
kill -INT $STDIO_PID 2>/dev/null || true
wait $STDIO_PID 2>/dev/null || true
$GO_BIN tool covdata textfmt -i="$WORK/cmd_cover" -o "$WORK/cmd_driver.out"

echo "[4/6] 合并所有覆盖率 profile"
# 合并：父级各包 + cli driver + cmd driver，按块取最大值
python3 - "$WORK" <<'PY'
import sys, os, glob
work = sys.argv[1]
def load(p):
    blocks = {}
    mode = None
    if not os.path.exists(p): return blocks, mode
    for line in open(p):
        line = line.rstrip('\n')
        if line.startswith('mode:'):
            mode = line.split(':',1)[1].strip(); continue
        parts = line.rsplit(' ', 1)
        if len(parts) != 2: continue
        try: val = int(parts[1])
        except ValueError: continue
        blocks[parts[0]] = max(blocks.get(parts[0], 0), val)
    return blocks, mode
all_blocks = {}
mode = 'atomic'
for f in ['core.out','verfmt.out','mcp.out','cli.out','cmdv.out','cmdm.out']:
    b, m = load(f'{work}/{f}')
    if m: mode = m
    for k,v in b.items(): all_blocks[k] = max(all_blocks.get(k,0), v)
# driver 文件：cli_driver.out 和 cmd_driver.out
for f in ['cli_driver.out','cmd_driver.out']:
    b, m = load(f'{work}/{f}')
    for k,v in b.items():
        # 只保留本项目代码
        if 'github.com/scagogogo/versions-skills/' in k and '/testdriver/' not in k:
            all_blocks[k] = max(all_blocks.get(k,0), v)
with open(f'{work}/merged.out','w') as o:
    o.write(f'mode: {mode}\n')
    for k in sorted(all_blocks):
        o.write(f'{k} {all_blocks[k]}\n')
print('merged profile written')
PY

echo "[5/6] 覆盖率结果"
$GO_BIN tool cover -func="$WORK/merged.out" | tail -5

echo "[6/6] 复制合并 profile"
cp "$WORK/merged.out" "$ROOT/coverage_merged.out"
echo "已写入 coverage_merged.out"
