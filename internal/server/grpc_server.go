package server

import (
	"context"
	"log"

	"github.com/example/mcp-testing/pkg/github.com/example/mcp-testing/pkg/mcp"
)

type MCPServer struct {
	mcp.UnimplementedMCPServiceServer
}

func NewMCPServer() *MCPServer {
	return &MCPServer{}
}

func (s *MCPServer) Initialize(ctx context.Context, req *mcp.InitializeRequest) (*mcp.InitializeResponse, error) {
	log.Printf("Initialize called with protocol version: %s", req.ProtocolVersion)
	
	return &mcp.InitializeResponse{
		ProtocolVersion: "2024-11-05",
		Capabilities: &mcp.ServerCapabilities{
			Logging:   true,
			Prompts:   true,
			Resources: true,
			Tools:     true,
		},
		ServerInfo: &mcp.ServerInfo{
			Name:    "MCP Reference Server",
			Version: "1.0.0",
		},
	}, nil
}

func (s *MCPServer) ListTools(ctx context.Context, req *mcp.ListToolsRequest) (*mcp.ListToolsResponse, error) {
	log.Printf("ListTools called")
	
	tools := []*mcp.Tool{
		{
			Name:        "echo",
			Description: "Echo back the input text",
			InputSchema: map[string]string{
				"type": "object",
				"properties": `{"text": {"type": "string", "description": "Text to echo back"}}`,
				"required": `["text"]`,
			},
		},
		{
			Name:        "add",
			Description: "Add two numbers together",
			InputSchema: map[string]string{
				"type": "object",
				"properties": `{"a": {"type": "number"}, "b": {"type": "number"}}`,
				"required": `["a", "b"]`,
			},
		},
	}
	
	return &mcp.ListToolsResponse{
		Tools: tools,
	}, nil
}

func (s *MCPServer) CallTool(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResponse, error) {
	log.Printf("CallTool called with name: %s", req.Name)
	
	switch req.Name {
	case "echo":
		text, ok := req.Arguments["text"]
		if !ok {
			return &mcp.CallToolResponse{
				Content: []*mcp.ToolResult{
					{
						Type: "text",
						Text: "Error: missing 'text' argument",
					},
				},
				IsError: true,
			}, nil
		}
		
		return &mcp.CallToolResponse{
			Content: []*mcp.ToolResult{
				{
					Type: "text",
					Text: text,
				},
			},
			IsError: false,
		}, nil
		
	case "add":
		aStr, aOk := req.Arguments["a"]
		bStr, bOk := req.Arguments["b"]
		if !aOk || !bOk {
			return &mcp.CallToolResponse{
				Content: []*mcp.ToolResult{
					{
						Type: "text",
						Text: "Error: missing 'a' or 'b' argument",
					},
				},
				IsError: true,
			}, nil
		}
		
		// Simple string concatenation for demo - in real implementation would parse numbers
		result := "Result: " + aStr + " + " + bStr
		
		return &mcp.CallToolResponse{
			Content: []*mcp.ToolResult{
				{
					Type: "text",
					Text: result,
				},
			},
			IsError: false,
		}, nil
		
	default:
		return &mcp.CallToolResponse{
			Content: []*mcp.ToolResult{
				{
					Type: "text",
					Text: "Error: unknown tool '" + req.Name + "'",
				},
			},
			IsError: true,
		}, nil
	}
}

func (s *MCPServer) ListPrompts(ctx context.Context, req *mcp.ListPromptsRequest) (*mcp.ListPromptsResponse, error) {
	log.Printf("ListPrompts called")
	
	prompts := []*mcp.Prompt{
		{
			Name:        "greeting",
			Description: "Generate a greeting message",
			Arguments: []*mcp.PromptArgument{
				{
					Name:        "name",
					Description: "Name of the person to greet",
					Required:    true,
				},
			},
		},
	}
	
	return &mcp.ListPromptsResponse{
		Prompts: prompts,
	}, nil
}

func (s *MCPServer) GetPrompt(ctx context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResponse, error) {
	log.Printf("GetPrompt called with name: %s", req.Name)
	
	switch req.Name {
	case "greeting":
		name, ok := req.Arguments["name"]
		if !ok {
			name = "World"
		}
		
		return &mcp.GetPromptResponse{
			Description: "A friendly greeting",
			Messages: []*mcp.PromptMessage{
				{
					Role:    "user",
					Content: "Hello, " + name + "! How are you today?",
				},
			},
		}, nil
		
	default:
		return &mcp.GetPromptResponse{
			Description: "Unknown prompt",
			Messages: []*mcp.PromptMessage{
				{
					Role:    "system",
					Content: "Error: unknown prompt '" + req.Name + "'",
				},
			},
		}, nil
	}
}

func (s *MCPServer) ListResources(ctx context.Context, req *mcp.ListResourcesRequest) (*mcp.ListResourcesResponse, error) {
	log.Printf("ListResources called")
	
	resources := []*mcp.Resource{
		{
			Uri:         "file://README.md",
			Name:        "README",
			Description: "Project README file",
			MimeType:    "text/markdown",
		},
		{
			Uri:         "file://config.json",
			Name:        "Configuration",
			Description: "Application configuration",
			MimeType:    "application/json",
		},
	}
	
	return &mcp.ListResourcesResponse{
		Resources: resources,
	}, nil
}

func (s *MCPServer) ReadResource(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResponse, error) {
	log.Printf("ReadResource called with URI: %s", req.Uri)
	
	switch req.Uri {
	case "file://README.md":
		return &mcp.ReadResourceResponse{
			Contents: []*mcp.ResourceContent{
				{
					Uri:      req.Uri,
					MimeType: "text/markdown",
					Text:     "# MCP Reference Server\n\nThis is a reference implementation of an MCP server using Go and gRPC.",
				},
			},
		}, nil
		
	case "file://config.json":
		return &mcp.ReadResourceResponse{
			Contents: []*mcp.ResourceContent{
				{
					Uri:      req.Uri,
					MimeType: "application/json",
					Text:     `{"name": "mcp-server", "version": "1.0.0", "debug": true}`,
				},
			},
		}, nil
		
	default:
		return &mcp.ReadResourceResponse{
			Contents: []*mcp.ResourceContent{
				{
					Uri:      req.Uri,
					MimeType: "text/plain",
					Text:     "Error: resource not found",
				},
			},
		}, nil
	}
}