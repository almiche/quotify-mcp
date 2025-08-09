package main

import (
	"fmt"
	"log"
	"os"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

// EchoArguments defines the input for the echo tool
type EchoArguments struct {
	Text string `json:"text" jsonschema:"required,description=Text to echo back"`
}

// AddArguments defines the input for the add tool
type AddArguments struct {
	A float64 `json:"a" jsonschema:"required,description=First number"`
	B float64 `json:"b" jsonschema:"required,description=Second number"`
}

// GreetingArguments defines the input for the greeting prompt
type GreetingArguments struct {
	Name string `json:"name" jsonschema:"required,description=Name of the person to greet"`
}

func main() {
	// Log to stderr so it doesn't interfere with MCP stdio
	log.SetOutput(os.Stderr)
	log.Printf("Starting MCP server...")

	// Create server with stdio transport
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// Register echo tool
	err := server.RegisterTool("echo", "Echo back the input text", func(arguments EchoArguments) (*mcp_golang.ToolResponse, error) {
		log.Printf("Echo tool called with text: %s", arguments.Text)
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Echo: %s", arguments.Text))), nil
	})
	if err != nil {
		panic(err)
	}

	// Register add tool
	err = server.RegisterTool("add", "Add two numbers together", func(arguments AddArguments) (*mcp_golang.ToolResponse, error) {
		log.Printf("Add tool called with: %.2f + %.2f", arguments.A, arguments.B)
		result := arguments.A + arguments.B
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Result: %.2f + %.2f = %.2f", arguments.A, arguments.B, result))), nil
	})
	if err != nil {
		panic(err)
	}

	// Register greeting prompt
	err = server.RegisterPrompt("greeting", "Generate a friendly greeting", func(arguments GreetingArguments) (*mcp_golang.PromptResponse, error) {
		message := fmt.Sprintf("Hello, %s! How are you doing today?", arguments.Name)
		return mcp_golang.NewPromptResponse(
			"A friendly greeting message",
			mcp_golang.NewPromptMessage(
				mcp_golang.NewTextContent(message),
				mcp_golang.RoleUser,
			),
		), nil
	})
	if err != nil {
		panic(err)
	}

	// Register resources
	err = server.RegisterResource("file://README.md", "README", "Project README file", "text/markdown", func() (*mcp_golang.ResourceResponse, error) {
		content := "# MCP Reference Server\n\nThis is a reference implementation of an MCP server using Go and the mcp-golang library."
		return mcp_golang.NewResourceResponse(
			mcp_golang.NewTextEmbeddedResource(
				"file://README.md",
				content,
				"text/markdown",
			),
		), nil
	})
	if err != nil {
		panic(err)
	}

	err = server.RegisterResource("file://config.json", "Configuration", "Application configuration", "application/json", func() (*mcp_golang.ResourceResponse, error) {
		content := `{"name": "mcp-server", "version": "1.0.0", "debug": true}`
		return mcp_golang.NewResourceResponse(
			mcp_golang.NewTextEmbeddedResource(
				"file://config.json",
				content,
				"application/json",
			),
		), nil
	})
	if err != nil {
		panic(err)
	}

	// Start the server
	log.Printf("MCP server ready, starting to serve...")
	err = server.Serve()
	if err != nil {
		log.Printf("Server error: %v", err)
		panic(err)
	}
}