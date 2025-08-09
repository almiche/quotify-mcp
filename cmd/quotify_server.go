package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/example/mcp-testing/pkg/quotify"
)

var httpAddr = flag.String("http", "", "if set, use streamable HTTP at this address, instead of stdin/stdout")

type QuotifyArgs struct {
	Format string `json:"format,omitempty" jsonschema:"format for the quote output: 'json' or 'text' (default: text)"`
}

func QuotifyTool(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[QuotifyArgs]) (*mcp.CallToolResultFor[struct{}], error) {
	log.Printf("Quotify tool called with format: %s", params.Arguments.Format)
	
	q := quotify.New()
	quote := q.Generate()
	
	var response string
	
	switch params.Arguments.Format {
	case "json":
		jsonData, err := json.MarshalIndent(quote, "", "  ")
		if err != nil {
			log.Printf("Error marshaling quote to JSON: %v", err)
			return &mcp.CallToolResultFor[struct{}]{
				Content: []mcp.Content{
					&mcp.TextContent{Text: "Error generating JSON quote"},
				},
			}, nil
		}
		response = string(jsonData)
	default:
		response = q.GenerateString()
	}
	
	return &mcp.CallToolResultFor[struct{}]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: response},
		},
	}, nil
}

func main() {
	// Log to stderr so it doesn't interfere with MCP stdio
	log.SetOutput(os.Stderr)
	log.Printf("Starting Quotify MCP server...")

	flag.Parse()

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "quotify-server",
		Version: "1.0.0",
	}, nil)

	// Add quotify tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "quotify",
		Description: "Generate a random quote with a random author attribution in the style of the original quotify Ruby gem",
	}, QuotifyTool)

	log.Printf("Quotify MCP server ready, starting to serve...")
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Printf("Server error: %v", err)
		panic(err)
	}
}