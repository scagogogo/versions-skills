#!/usr/bin/env bash
#
# install.sh — Quick installer for the versions CLI
#
# Usage:
#   curl -sL https://raw.githubusercontent.com/scagogogo/versions-skills/main/install.sh | bash
#
# Detects OS and architecture, downloads the latest release from GitHub,
# and installs to /usr/local/bin (may require sudo).

set -euo pipefail

REPO="scagogogo/versions-skills"
BINARY="versions"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# --- Detect OS ---
detect_os() {
    local os
    case "$(uname -s)" in
        Linux)  os="linux" ;;
        Darwin) os="darwin" ;;
        FreeBSD) os="freebsd" ;;
        OpenBSD) os="openbsd" ;;
        NetBSD)  os="netbsd" ;;
        *)
            # Windows via MSYS/Cygwin
            case "$(uname -s)" in
                MINGW*|MSYS*|CYGWIN*) os="windows" ;;
                *) echo "❌ Unsupported OS: $(uname -s)"; exit 1 ;;
            esac
            ;;
    esac
    echo "$os"
}

# --- Detect Architecture ---
detect_arch() {
    local arch
    case "$(uname -m)" in
        x86_64|amd64)   arch="amd64" ;;
        aarch64|arm64)  arch="arm64" ;;
        armv7l|armv6l)  arch="arm" ;;
        i386|i686)      arch="386" ;;
        mips)           arch="mips" ;;
        mips64)         arch="mips64" ;;
        ppc64)          arch="ppc64" ;;
        ppc64le)        arch="ppc64le" ;;
        s390x)          arch="s390x" ;;
        riscv64)        arch="riscv64" ;;
        *)
            echo "❌ Unsupported architecture: $(uname -m)"
            exit 1
            ;;
    esac
    echo "$arch"
}

# --- Get latest version tag ---
get_latest_version() {
    local tag
    tag=$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    if [ -z "$tag" ]; then
        echo "❌ Failed to fetch latest release info"
        exit 1
    fi
    echo "$tag"
}

# --- Main ---
main() {
    echo "🔍 Detecting platform..."
    OS=$(detect_os)
    ARCH=$(detect_arch)
    echo "   OS: $OS, Arch: $ARCH"

    echo "📦 Fetching latest release..."
    VERSION=$(get_latest_version)
    echo "   Version: $VERSION"

    if [ "$OS" = "windows" ]; then
        ARCHIVE="${BINARY}_${VERSION}_${OS}_${ARCH}.zip"
        echo "⚠ Windows detected. Please download manually from:"
        echo "   https://github.com/${REPO}/releases/latest/download/${ARCHIVE}"
        exit 0
    fi

    ARCHIVE="${BINARY}_${VERSION}_${OS}_${ARCH}.tar.gz"
    DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/${ARCHIVE}"

    echo "⬇ Downloading ${ARCHIVE}..."
    TMPDIR=$(mktemp -d)
    curl -sL "$DOWNLOAD_URL" -o "${TMPDIR}/${ARCHIVE}"

    echo "📦 Extracting..."
    tar xzf "${TMPDIR}/${ARCHIVE}" -C "$TMPDIR"

    echo "📋 Installing to ${INSTALL_DIR}..."
    if [ -w "$INSTALL_DIR" ]; then
        mv "${TMPDIR}/${BINARY}" "${INSTALL_DIR}/"
    else
        sudo mv "${TMPDIR}/${BINARY}" "${INSTALL_DIR}/"
    fi

    rm -rf "$TMPDIR"

    echo "✅ ${BINARY} ${VERSION} installed to ${INSTALL_DIR}/${BINARY}"
    echo ""
    echo "Try it: ${BINARY} parse v1.2.3"
}

main "$@"
