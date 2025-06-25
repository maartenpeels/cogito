#!/bin/bash

# Cogito MCP Server Configuration Generator

SERVER_NAME="cogito"
BINARY_PATH="$(pwd)/cogito"
VSCODE_CONFIG="vscode-mcp-settings.json"

echo "Generating Cogito MCP Server configuration..."

# Generate VS Code settings configuration
cat > "$VSCODE_CONFIG" << EOF
{
  "mcp.servers": {
    "$SERVER_NAME": {
      "command": "$BINARY_PATH",
      "args": [],
      "env": {}
    }
  }
}
EOF

echo "VS Code MCP configuration generated: $VSCODE_CONFIG"
echo ""
echo "To use with VS Code:"
echo "1. Open VS Code Settings (JSON) - Cmd+Shift+P -> 'Preferences: Open Settings (JSON)'"
echo "2. Add the contents of $VSCODE_CONFIG to your settings.json"
echo "3. Restart VS Code"
echo ""
echo "Your server details:"
echo "  Name: $SERVER_NAME"
echo "  Binary: $BINARY_PATH"
echo ""
echo "To test the server directly:"
echo "  echo '{\"jsonrpc\": \"2.0\", \"id\": 1, \"method\": \"initialize\", \"params\": {\"protocolVersion\": \"2024-11-05\", \"capabilities\": {\"tools\": {}}}}' | $BINARY_PATH"
