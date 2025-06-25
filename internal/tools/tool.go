package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// ToolFactory is a function that creates a new instance of a tool
type ToolFactory func() Tool

// Global tool registry
var globalRegistry = NewRegistry()

// Register registers a tool factory with the global registry
// This should be called from init() functions in tool implementations
func Register(name string, factory ToolFactory) {
	tool := factory()
	if err := globalRegistry.Register(tool); err != nil {
		panic("failed to register tool '" + name + "': " + err.Error())
	}
}

// GetGlobalRegistry returns the global tool registry
func GetGlobalRegistry() *Registry {
	return globalRegistry
}

// Tool represents a MCP tool that can be registered with the server
type Tool interface {
	// Name returns the unique name of the tool
	Name() string

	// Definition returns the MCP tool definition
	Definition() mcp.Tool

	// Execute handles the tool execution
	Execute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Registry manages all available tools
type Registry struct {
	tools map[string]Tool
}

// NewRegistry creates a new tool registry
func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

// Register adds a tool to the registry
func (r *Registry) Register(tool Tool) error {
	name := tool.Name()
	if _, exists := r.tools[name]; exists {
		return &ToolRegistrationError{Name: name, Reason: "tool already registered"}
	}
	r.tools[name] = tool
	return nil
}

// Get retrieves a tool by name
func (r *Registry) Get(name string) (Tool, bool) {
	tool, exists := r.tools[name]
	return tool, exists
}

// List returns all registered tools
func (r *Registry) List() []Tool {
	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// ToolRegistrationError represents an error during tool registration
type ToolRegistrationError struct {
	Name   string
	Reason string
}

func (e *ToolRegistrationError) Error() string {
	return "failed to register tool '" + e.Name + "': " + e.Reason
}
