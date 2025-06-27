#!/usr/bin/env bash

set -e

REPO="smash-hq/scrape-cli"
VERSION="${1:-v2.0.0}"
INSTALL_DIR="$HOME/.local/bin"

echo "📦 Installing scrape-cli $VERSION..."

# Detect OS
OS="$(uname | tr '[:upper:]' '[:lower:]')"
case "$OS" in
  linux) OS="linux" ;;
  darwin) OS="darwin" ;;
  msys* | mingw* | cygwin* | windows*) OS="windows" ;;
  *) echo "❌ Unsupported OS: $OS" && exit 1 ;;
esac

# Detect ARCH
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64 | amd64) ARCH="amd64" ;;
  i386 | i686) ARCH="386" ;;
  arm64 | aarch64) ARCH="arm64" ;;
  *) echo "❌ Unsupported architecture: $ARCH" && exit 1 ;;
esac

FILENAME="scrape-cli_${OS}_${ARCH}"
EXT="tar.gz"
[[ "$OS" == "windows" ]] && EXT="zip"

TARBALL="$FILENAME.$EXT"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$VERSION/$TARBALL"

echo "🔽 Downloading from $DOWNLOAD_URL..."
curl -fsSL -o "$TARBALL" "$DOWNLOAD_URL"

echo "📂 Extracting..."
if [[ "$EXT" == "zip" ]]; then
  unzip -o "$TARBALL"
else
  tar -xzf "$TARBALL"
fi
rm "$TARBALL"

BINARY_NAME="scrape-cli"
[[ "$OS" == "windows" ]] && BINARY_NAME="scrape-cli.exe"

chmod +x "$BINARY_NAME"

mkdir -p "$INSTALL_DIR"
mv -f "$BINARY_NAME" "$INSTALL_DIR/"

echo "✅ Installed at: $INSTALL_DIR/$BINARY_NAME"

# Check if in PATH
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "🔔 Add the following to your shell config (~/.bashrc or ~/.zshrc):"
  echo "    export PATH=\"\$HOME/.local/bin:\$PATH\""
else
  echo "🚀 You can now run: $BINARY_NAME"
fi
