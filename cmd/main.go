package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/example/mcp-testing/internal/server"
	"github.com/example/mcp-testing/pkg/github.com/example/mcp-testing/pkg/mcp"
)

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Create and register the MCP server
	mcpServer := server.NewMCPServer()
	mcp.RegisterMCPServiceServer(s, mcpServer)

	log.Printf("MCP gRPC server listening on :50051")

	// Start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}