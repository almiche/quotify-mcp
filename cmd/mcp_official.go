package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var httpAddr = flag.String("http", "", "if set, use streamable HTTP at this address, instead of stdin/stdout")

type EchoArgs struct {
	Text string `json:"text" jsonschema:"the text to echo back"`
}

type AddArgs struct {
	A float64 `json:"a" jsonschema:"first number to add"`
	B float64 `json:"b" jsonschema:"second number to add"`
}

type GreetingArgs struct {
	Name string `json:"name" jsonschema:"the name to greet"`
}

func EchoTool(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[EchoArgs]) (*mcp.CallToolResultFor[struct{}], error) {
	log.Printf("Echo tool called with text: %s", params.Arguments.Text)
	return &mcp.CallToolResultFor[struct{}]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Echo: " + params.Arguments.Text},
		},
	}, nil
}

func AddTool(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[AddArgs]) (*mcp.CallToolResultFor[struct{}], error) {
	log.Printf("Add tool called with: %.2f + %.2f", params.Arguments.A, params.Arguments.B)
	result := params.Arguments.A + params.Arguments.B
	return &mcp.CallToolResultFor[struct{}]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Result: %.2f + %.2f = %.2f", params.Arguments.A, params.Arguments.B, result)},
		},
	}, nil
}

func GreetingPrompt(ctx context.Context, ss *mcp.ServerSession, params *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	name, ok := params.Arguments["name"]
	if !ok {
		name = "World"
	}
	log.Printf("Greeting prompt called with name: %s", name)
	
	return &mcp.GetPromptResult{
		Description: "A friendly greeting message",
		Messages: []*mcp.PromptMessage{
			{Role: "user", Content: &mcp.TextContent{Text: "Hello, " + name + "! How are you doing today?"}},
		},
	}, nil
}

func main() {
	// Log to stderr so it doesn't interfere with MCP stdio
	log.SetOutput(os.Stderr)
	log.Printf("Starting MCP server with official SDK...")

	flag.Parse()

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-testing-server",
		Version: "1.0.0",
	}, nil)

	// Add echo tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "echo",
		Description: "Echo back the input text",
	}, EchoTool)

	// Add add tool  
	mcp.AddTool(server, &mcp.Tool{
		Name:        "add", 
		Description: "Add two numbers together",
	}, AddTool)

	// Add greeting prompt
	server.AddPrompt(&mcp.Prompt{
		Name:        "greeting",
		Description: "Generate a friendly greeting",
	}, GreetingPrompt)

	log.Printf("MCP server ready, starting to serve...")
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Printf("Server error: %v", err)
		panic(err)
	}
}