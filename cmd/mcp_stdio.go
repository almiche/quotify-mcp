package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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

// MCP protocol types
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

type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"inputSchema"`
}

type CallToolParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

type CallToolResult struct {
	Content []map[string]string `json:"content"`
	IsError bool                `json:"isError"`
}

func main() {
	// Log to stderr to avoid interfering with stdio communication
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	log.Printf("MCP stdio server starting...")
	
	// Keep reading from stdin indefinitely using ReadLine
	reader := bufio.NewReader(os.Stdin)
	
	for {
		log.Printf("Waiting for input...")
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			log.Printf("ReadLine error: %v", err)
			if err.Error() == "EOF" {
				log.Printf("Got EOF, but continuing to wait...")
				// Don't break on EOF, keep waiting
				continue
			}
			break
		}
		
		// Handle partial lines
		if isPrefix {
			log.Printf("Got partial line, continuing to read...")
			continue
		}
		
		lineStr := strings.TrimSpace(string(line))
		if lineStr == "" {
			continue
		}
		
		log.Printf("Received: %s", lineStr)
		
		var req JSONRPCRequest
		if err := json.Unmarshal([]byte(lineStr), &req); err != nil {
			log.Printf("Parse error: %v", err)
			sendError(-32700, "Parse error", nil)
			continue
		}
		
		// Check for shutdown notification
		if req.Method == "notifications/cancelled" || req.Method == "exit" {
			log.Printf("Shutdown requested, exiting")
			break
		}
		
		handleRequest(&req)
		log.Printf("Finished handling %s, waiting for next message...", req.Method)
	}
	
	log.Printf("MCP stdio server shutting down")
}

func handleRequest(req *JSONRPCRequest) {
	log.Printf("Handling method: %s", req.Method)
	
	switch req.Method {
	case "initialize":
		var params InitializeParams
		if err := mapToStruct(req.Params, &params); err != nil {
			log.Printf("Invalid params: %v", err)
			sendError(-32602, "Invalid params", req.ID)
			return
		}
		
		log.Printf("Initialize with protocol version: %s", params.ProtocolVersion)
		
		result := InitializeResult{
			ProtocolVersion: params.ProtocolVersion, // Match client's protocol version
			Capabilities: map[string]interface{}{
				"logging":   true,
				"prompts":   true,
				"resources": true,
				"tools":     true,
			},
			ServerInfo: map[string]string{
				"name":    "MCP Reference Server",
				"version": "1.0.0",
			},
		}
		
		sendResult(result, req.ID)
		
	case "tools/list":
		log.Printf("Listing tools")
		
		tools := []Tool{
			{
				Name:        "echo",
				Description: "Echo back the input text",
				InputSchema: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"text": map[string]interface{}{
							"type":        "string",
							"description": "Text to echo back",
						},
					},
					"required": []string{"text"},
				},
			},
			{
				Name:        "add",
				Description: "Add two numbers together",
				InputSchema: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"a": map[string]interface{}{
							"type": "number",
						},
						"b": map[string]interface{}{
							"type": "number",
						},
					},
					"required": []string{"a", "b"},
				},
			},
		}
		
		sendResult(map[string]interface{}{"tools": tools}, req.ID)
		
	case "tools/call":
		var params CallToolParams
		if err := mapToStruct(req.Params, &params); err != nil {
			log.Printf("Invalid params: %v", err)
			sendError(-32602, "Invalid params", req.ID)
			return
		}
		
		log.Printf("Calling tool: %s with args: %v", params.Name, params.Arguments)
		
		switch params.Name {
		case "echo":
			text, ok := params.Arguments["text"]
			if !ok {
				sendResult(CallToolResult{
					Content: []map[string]string{
						{"type": "text", "text": "Error: missing 'text' argument"},
					},
					IsError: true,
				}, req.ID)
				return
			}
			
			sendResult(CallToolResult{
				Content: []map[string]string{
					{"type": "text", "text": fmt.Sprintf("Echo: %v", text)},
				},
				IsError: false,
			}, req.ID)
			
		case "add":
			a, aOk := params.Arguments["a"]
			b, bOk := params.Arguments["b"]
			if !aOk || !bOk {
				sendResult(CallToolResult{
					Content: []map[string]string{
						{"type": "text", "text": "Error: missing 'a' or 'b' argument"},
					},
					IsError: true,
				}, req.ID)
				return
			}
			
			sendResult(CallToolResult{
				Content: []map[string]string{
					{"type": "text", "text": fmt.Sprintf("Result: %v + %v", a, b)},
				},
				IsError: false,
			}, req.ID)
			
		default:
			sendResult(CallToolResult{
				Content: []map[string]string{
					{"type": "text", "text": fmt.Sprintf("Error: unknown tool '%s'", params.Name)},
				},
				IsError: true,
			}, req.ID)
		}
		
	default:
		log.Printf("Method not found: %s", req.Method)
		sendError(-32601, "Method not found", req.ID)
	}
}

func sendResult(result interface{}, id interface{}) {
	resp := JSONRPCResponse{
		Jsonrpc: "2.0",
		Result:  result,
		ID:      id,
	}
	
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Marshal error: %v", err)
		return
	}
	
	log.Printf("Sending: %s", string(data))
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
	
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Marshal error: %v", err)
		return
	}
	
	log.Printf("Sending error: %s", string(data))
	fmt.Println(string(data))
}

func mapToStruct(input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, output)
}