package server

import (
	"context"
	"fmt"

	"github.com/maartenpeels/ai-tools/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// MCPServer wraps the MCP server with tool management
type MCPServer struct {
	server   *server.MCPServer
	registry *tools.Registry
}

// Config holds server configuration
type Config struct {
	Name    string
	Version string
}

// NewMCPServer creates a new MCP server with default configuration and registers all tools
// from the global tool registry.
func NewMCPServer(config Config) (*MCPServer, error) {
	s := server.NewMCPServer(
		config.Name,
		config.Version,
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	mcpServer := &MCPServer{
		server:   s,
		registry: tools.NewRegistry(),
	}

	// Register all tools from the global registry
	if err := mcpServer.RegisterAllFromGlobal(); err != nil {
		return nil, fmt.Errorf("failed to register tools from global registry: %w", err)
	}

	return mcpServer, nil
}

// RegisterAllFromGlobal registers all tools from the global registry
func (s *MCPServer) RegisterAllFromGlobal() error {
	globalRegistry := tools.GetGlobalRegistry()
	for _, tool := range globalRegistry.List() {
		if err := s.RegisterTool(tool); err != nil {
			return fmt.Errorf("failed to register tool '%s': %w", tool.Name(), err)
		}
	}
	return nil
}

// RegisterTool registers a tool with the server
func (s *MCPServer) RegisterTool(tool tools.Tool) error {
	// Register with our registry
	if err := s.registry.Register(tool); err != nil {
		return fmt.Errorf("failed to register tool with registry: %w", err)
	}

	// Add tool to MCP server with a wrapper handler
	s.server.AddTool(tool.Definition(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tool.Execute(ctx, request)
	})

	return nil
}

// GetTool retrieves a registered tool by name
func (s *MCPServer) GetTool(name string) (tools.Tool, bool) {
	return s.registry.Get(name)
}

// ListTools returns all registered tools
func (s *MCPServer) ListTools() []tools.Tool {
	return s.registry.List()
}

// Serve starts the server using stdio transport
func (s *MCPServer) Serve() error {
	return server.ServeStdio(s.server)
}
