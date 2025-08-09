#!/bin/bash

echo "=== Testing MCP Server ==="
echo ""

echo "1. Initialize:"
echo '{"jsonrpc":"2.0","method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{"roots":false,"sampling":false},"clientInfo":{"name":"test-client","version":"1.0.0"}},"id":1}' | ./bin/mcp-stdio 2>/dev/null
echo ""

echo "2. List Tools:"
echo '{"jsonrpc":"2.0","method":"tools/list","params":{},"id":2}' | ./bin/mcp-stdio 2>/dev/null
echo ""

echo "3. Call Echo Tool:"
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"echo","arguments":{"text":"Hello from MCP"}},"id":3}' | ./bin/mcp-stdio 2>/dev/null
echo ""

echo "4. Call Add Tool:"
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"add","arguments":{"a":42,"b":58}},"id":4}' | ./bin/mcp-stdio 2>/dev/null
echo ""

echo "=== MCP Server Test Complete ==="