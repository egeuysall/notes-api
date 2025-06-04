#!/bin/sh
set -e

VERSION="1.1.0"
REPO="egeuysall/notes-api"
TMPDIR=$(mktemp -d)

cleanup() {
  rm -rf "$TMPDIR"
}
trap cleanup EXIT

# Detect OS
OS=$(uname | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux|darwin) ;;
  mingw*|msys*|cygwin*)
    OS="windows"
    ;;
  *)
    echo "❌ Unsupported OS: $OS"
    exit 1
    ;;
esac

# Detect Architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)
    echo "❌ Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

FILE="notes-api_${VERSION}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/v${VERSION}/${FILE}"
CHECKSUM_URL="https://github.com/${REPO}/releases/download/v${VERSION}/notes-api_${VERSION}_checksums.txt"

echo "📦 Downloading checksum file..."
if ! curl -fsSL "$CHECKSUM_URL" -o "$TMPDIR/checksums.txt"; then
  echo "❌ Failed to download checksum file"
  exit 1
fi

echo "📦 Downloading $FILE ..."
if ! curl -fsSL "$URL" -o "$TMPDIR/$FILE"; then
  echo "❌ Failed to download $FILE"
  exit 1
fi

# Verify the downloaded file exists
if [ ! -f "$TMPDIR/$FILE" ]; then
  echo "❌ Downloaded file not found: $TMPDIR/$FILE"
  exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then
  CHECKSUM_TOOL="sha256sum"
elif command -v shasum >/dev/null 2>&1; then
  CHECKSUM_TOOL="shasum -a 256"
else
  echo "❌ Error: neither sha256sum nor shasum found."
  exit 1
fi

echo "🔍 Verifying checksum..."

EXPECTED_SUM=$(grep "$FILE" "$TMPDIR/checksums.txt" | awk '{print $1}')
if [ -z "$EXPECTED_SUM" ]; then
  echo "❌ Could not find checksum for $FILE in checksums.txt"
  exit 1
fi

ACTUAL_SUM=$($CHECKSUM_TOOL "$TMPDIR/$FILE" | awk '{print $1}')

if [ "$EXPECTED_SUM" != "$ACTUAL_SUM" ]; then
  echo "❌ Checksum verification failed!"
  echo "Expected: $EXPECTED_SUM"
  echo "Actual:   $ACTUAL_SUM"
  exit 1
fi

echo "✅ Checksum verified."

echo "📂 Extracting archive..."
if ! tar -xzf "$TMPDIR/$FILE" -C "$TMPDIR"; then
  echo "❌ Failed to extract archive"
  exit 1
fi

# Debug: List what was extracted
echo "📋 Archive contents:"
tar -tzf "$TMPDIR/$FILE"

# Get the first directory/file name from the archive
FIRST_ENTRY=$(tar -tzf "$TMPDIR/$FILE" | head -1)
echo "🔍 First entry in archive: $FIRST_ENTRY"

# Check if the archive contains a directory or direct files
if echo "$FIRST_ENTRY" | grep -q "/"; then
  # Archive contains directories
  EXTRACTED_DIR=$(echo "$FIRST_ENTRY" | cut -f1 -d"/")
  BINARY_PATH="$TMPDIR/$EXTRACTED_DIR/notes-api"
else
  # Archive contains files directly
  EXTRACTED_DIR=""
  BINARY_PATH="$TMPDIR/notes-api"
fi

echo "🔍 Looking for binary at: $BINARY_PATH"

# Verify the binary exists before proceeding
if [ ! -f "$BINARY_PATH" ]; then
  echo "❌ Error: Binary not found at expected location: $BINARY_PATH"
  echo "📋 Contents of extraction directory:"
  if [ -n "$EXTRACTED_DIR" ] && [ -d "$TMPDIR/$EXTRACTED_DIR" ]; then
    ls -la "$TMPDIR/$EXTRACTED_DIR"
  else
    ls -la "$TMPDIR"
  fi
  
  # Try to find the binary anywhere in the temp directory
  echo "🔍 Searching for 'notes-api' binary in temp directory..."
  find "$TMPDIR" -name "notes-api" -type f 2>/dev/null || echo "No 'notes-api' binary found"
  exit 1
fi

# Set up installation directory
USER_BIN="$HOME/.local/bin"
mkdir -p "$USER_BIN"

echo "📁 Installing binary to $USER_BIN ..."

# Make binary executable
if ! chmod +x "$BINARY_PATH"; then
  echo "❌ Failed to make binary executable"
  exit 1
fi

# Move binary to installation directory
if ! mv "$BINARY_PATH" "$USER_BIN/notes-api"; then
  echo "❌ Failed to move binary to $USER_BIN"
  exit 1
fi

# Check if the binary is in PATH
case ":$PATH:" in
  *":$USER_BIN:"*) 
    echo "✅ $USER_BIN is already in your PATH"
    ;;
  *)
    echo "⚠️  Note: $USER_BIN is not in your PATH."
    echo "Add it to your PATH by running one of the following:"
    echo ""
    echo "For bash:"
    echo "  echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.bashrc"
    echo "  source ~/.bashrc"
    echo ""
    echo "For zsh:"
    echo "  echo 'export PATH=\"\$HOME/.local/bin:\$PATH\"' >> ~/.zshrc"
    echo "  source ~/.zshrc"
    echo ""
    echo "For fish:"
    echo "  echo 'set -gx PATH \$HOME/.local/bin \$PATH' >> ~/.config/fish/config.fish"
    echo ""
    ;;
esac

echo "✅ Installation complete!"

# Verify installation
if command -v notes-api >/dev/null 2>&1; then
  echo "🎉 Success! 'notes-api' is now available in your PATH"
  echo "📋 Version info:"
  notes-api --version
else
  echo "⚠️  'notes-api' is installed but not in your current PATH"
  echo "You can run it directly with: $USER_BIN/notes-api --version"
  echo "Or add $USER_BIN to your PATH as shown above"
fi