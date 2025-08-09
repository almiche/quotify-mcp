package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/example/mcp-testing/internal/server"
)

// JSON-RPC 2.0 message structure
type JSONRPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      interface{} `json:"id"`
}

type JSONRPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// MCP protocol types for stdio transport
type InitializeParams struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    map[string]interface{} `json:"capabilities"`
	ClientInfo      map[string]string      `json:"clientInfo"`
}

type InitializeResult struct {
	ProtocolVersion string                 `json:"protocolVersion"`
	Capabilities    map[string]interface{} `json:"capabilities"`
	ServerInfo      map[string]string      `json:"serverInfo"`
}

type ListToolsResult struct {
	Tools []map[string]interface{} `json:"tools"`
}

type CallToolParams struct {
	Name      string            `json:"name"`
	Arguments map[string]string `json:"arguments"`
}

type CallToolResult struct {
	Content []map[string]string `json:"content"`
	IsError bool                `json:"isError"`
}

func main() {
	// Disable regular logging to avoid interfering with stdio
	log.SetOutput(os.Stderr)
	
	mcpServer := server.NewMCPServer()
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		
		var req JSONRPCRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			sendError(-32700, "Parse error", req.ID)
			continue
		}
		
		handleRequest(mcpServer, &req)
	}
}

func handleRequest(mcpServer *server.MCPServer, req *JSONRPCRequest) {
	ctx := context.Background()
	
	switch req.Method {
	case "initialize":
		var params InitializeParams
		if err := mapToStruct(req.Params, &params); err != nil {
			sendError(-32602, "Invalid params", req.ID)
			return
		}
		
		// Convert to gRPC request
		grpcReq := &mcpProto.InitializeRequest{
			ProtocolVersion: params.ProtocolVersion,
			Capabilities: &mcpProto.ClientCapabilities{
				Roots:    getBool(params.Capabilities, "roots"),
				Sampling: getBool(params.Capabilities, "sampling"),
			},
			ClientInfo: &mcpProto.ClientInfo{
				Name:    params.ClientInfo["name"],
				Version: params.ClientInfo["version"],
			},
		}
		
		resp, err := mcpServer.Initialize(ctx, grpcReq)
		if err != nil {
			sendError(-32603, "Internal error", req.ID)
			return
		}
		
		result := InitializeResult{
			ProtocolVersion: resp.ProtocolVersion,
			Capabilities: map[string]interface{}{
				"logging":   resp.Capabilities.Logging,
				"prompts":   resp.Capabilities.Prompts,
				"resources": resp.Capabilities.Resources,
				"tools":     resp.Capabilities.Tools,
			},
			ServerInfo: map[string]string{
				"name":    resp.ServerInfo.Name,
				"version": resp.ServerInfo.Version,
			},
		}
		
		sendResult(result, req.ID)
		
	case "tools/list":
		grpcReq := &mcpProto.ListToolsRequest{}
		resp, err := mcpServer.ListTools(ctx, grpcReq)
		if err != nil {
			sendError(-32603, "Internal error", req.ID)
			return
		}
		
		var tools []map[string]interface{}
		for _, tool := range resp.Tools {
			tools = append(tools, map[string]interface{}{
				"name":        tool.Name,
				"description": tool.Description,
				"inputSchema": tool.InputSchema,
			})
		}
		
		sendResult(ListToolsResult{Tools: tools}, req.ID)
		
	case "tools/call":
		var params CallToolParams
		if err := mapToStruct(req.Params, &params); err != nil {
			sendError(-32602, "Invalid params", req.ID)
			return
		}
		
		grpcReq := &mcpProto.CallToolRequest{
			Name:      params.Name,
			Arguments: params.Arguments,
		}
		
		resp, err := mcpServer.CallTool(ctx, grpcReq)
		if err != nil {
			sendError(-32603, "Internal error", req.ID)
			return
		}
		
		var content []map[string]string
		for _, c := range resp.Content {
			content = append(content, map[string]string{
				"type": c.Type,
				"text": c.Text,
			})
		}
		
		sendResult(CallToolResult{
			Content: content,
			IsError: resp.IsError,
		}, req.ID)
		
	default:
		sendError(-32601, "Method not found", req.ID)
	}
}

func sendResult(result interface{}, id interface{}) {
	resp := JSONRPCResponse{
		Jsonrpc: "2.0",
		Result:  result,
		ID:      id,
	}
	
	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}

func sendError(code int, message string, id interface{}) {
	resp := JSONRPCResponse{
		Jsonrpc: "2.0",
		Error: &RPCError{
			Code:    code,
			Message: message,
		},
		ID: id,
	}
	
	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
}

func mapToStruct(input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, output)
}

func getBool(m map[string]interface{}, key string) bool {
	if val, ok := m[key]; ok {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}