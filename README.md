# Cogito - AI Tools MCP Server

A modular MCP (Model Context Protocol) server with extensible tool architecture for GitHub Copilot and other MCP clients.

## Project Structure

```
.
├── cmd/
│   └── server/          # Application entry point
│       └── main.go
├── internal/
│   ├── server/          # MCP server wrapper and configuration
│   │   └── server.go
│   └── tools/           # Tool implementations
│       ├── tool.go      # Tool interface and registry
│       └── ...          # Tool implementations
├── go.mod
├── go.sum
└── README.md
```

## Architecture

### Tool Interface

All tools implement the `tools.Tool` interface:

```go
type Tool interface {
    Name() string
    Definition() mcp.Tool
    Execute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}
```

### Adding New Tools

1. Create a new file in `internal/tools/`
2. Implement the `Tool` interface
3. Register the tool using the `tools.Register` function in the `init` function of your tool file.

Example:

```go
// internal/tools/mytool.go
type MyTool struct{}

func init() {
	Register("my-tool", NewMyTool)
}

func NewMyTool() *MyTool {
    return &MyTool{}
}

func (m *MyTool) Name() string {
    return "my-tool"
}

func (m *MyTool) Definition() *mcp.Tool {
    return mcp.NewTool("my-tool",
        mcp.WithDescription("My custom tool"),
        // Add parameters here
    )
}

func (m *MyTool) Execute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    // Implementation here
}
```

## Example Tool (Calculator)

Performs basic arithmetic operations (add, subtract, multiply, divide).

**Parameters:**
- `operation`: The operation to perform (add, subtract, multiply, divide)
- `x`: First number
- `y`: Second number

## Running the Server

```bash
go run cmd/server/main.go
```

## Building

```bash
go build -o cogito cmd/server/main.go
```

## Using with MCP Clients

### GitHub Copilot (VS Code)

1. **Build the Server**:
   ```bash
   go build -o cogito cmd/server/main.go
   ```

2. **Configure MCP Client**: Add this to your VS Code settings or MCP client config:
   ```json
   {
     "mcp": {
       "servers": {
         "cogito": {
           "command": "/path/to/your/cogito",
           "args": [],
           "env": {}
         }
       }
     }
   }
   ```
