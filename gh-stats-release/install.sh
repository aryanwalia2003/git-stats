#!/bin/bash

# Define the installation directory
INSTALL_DIR="/usr/local/bin"
EXECUTABLE_NAME="gh-stats"

# Detect OS and architecture
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

# Select the correct binary based on OS and architecture
if [ "$OS" = "darwin" ]; then
	    # For macOS, we provide the intel binary which runs on all modern Macs via Rosetta 2
    BINARY="./gh-stats-darwin-amd64"
elif [ "$OS" = "linux" ]; then
    BINARY="./gh-stats-linux-amd64"
else
    echo "❌ Unsupported OS: $OS. Please install manually."
    exit 1
fi

# Check if the binary exists in the downloaded folder
if [ ! -f "$BINARY" ]; then
    echo "❌ Could not find $BINARY. Please run this script from the unzipped directory."
    exit 1
fi

# Give execution permissions
chmod +x "$BINARY"

# Move binary to the global path /usr/local/bin
echo "🚀 Installing $EXECUTABLE_NAME to $INSTALL_DIR..."
sudo cp "$BINARY" "$INSTALL_DIR/$EXECUTABLE_NAME"

if [ $? -eq 0 ]; then
    echo "✅ Success! You can now run '$EXECUTABLE_NAME' in any git repository."
else
    echo "❌ Installation failed. Do you have sudo permissions?"
fi
