package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// CalculatorTool implements basic arithmetic operations
type CalculatorTool struct{}

// Register the calculator tool
func init() {
	Register("calculate", NewCalculatorTool)
}

// NewCalculatorTool creates a new calculator tool instance
func NewCalculatorTool() Tool {
	return &CalculatorTool{}
}

// Name returns the tool name
func (c *CalculatorTool) Name() string {
	return "calculate"
}

// Definition returns the MCP tool definition
func (c *CalculatorTool) Definition() mcp.Tool {
	return mcp.NewTool("calculate",
		mcp.WithDescription("Perform basic arithmetic operations"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
			mcp.Enum("add", "subtract", "multiply", "divide"),
		),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("First number"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Second number"),
		),
	)
}

// Execute handles the calculator tool execution
func (c *CalculatorTool) Execute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract and validate operation parameter
	operation, err := request.RequireString("operation")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("invalid operation parameter: %v", err)), nil
	}

	// Extract and validate x parameter
	x, err := request.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("invalid x parameter: %v", err)), nil
	}

	// Extract and validate y parameter
	y, err := request.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("invalid y parameter: %v", err)), nil
	}

	// Perform calculation
	result, err := c.calculate(operation, x, y)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
}

// calculate performs the actual arithmetic operation
func (c *CalculatorTool) calculate(operation string, x, y float64) (float64, error) {
	switch operation {
	case "add":
		return x + y, nil
	case "subtract":
		return x - y, nil
	case "multiply":
		return x * y, nil
	case "divide":
		if y == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", operation)
	}
}
