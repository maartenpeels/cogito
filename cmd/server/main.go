package main

import (
	"fmt"
	"log"

	"github.com/maartenpeels/ai-tools/internal/server"
	// Import tools to trigger their init() functions for registration
	_ "github.com/maartenpeels/ai-tools/internal/tools"
)

func main() {
	// Create server configuration
	config := server.Config{
		Name:    "Cogito",
		Version: "1.0.0",
	}

	// Create the MCP server with registered tools
	mcpServer, err := server.NewMCPServer(config)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// List registered tools for debugging
	tools := mcpServer.ListTools()
	fmt.Printf("Registered %d tools:\n", len(tools))
	for _, tool := range tools {
		fmt.Printf("  - %s\n", tool.Name())
	}

	// Start the server
	fmt.Printf("Starting %s v%s...\n", config.Name, config.Version)
	if err := mcpServer.Serve(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
