package interpreter

import (
	"fsl-parser/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariableDeclaration(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("var: 1 var2: 0")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, float64(1), fslInterpreter.variables["var"])
	assert.Equal(t, float64(0), fslInterpreter.variables["var2"])
}

func TestUpdateCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("var1: 1 var2: 0 " +
		"cmd:update, id:var1, value:3.5")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, 3.5, fslInterpreter.variables["var1"])
	assert.Equal(t, float64(0), fslInterpreter.variables["var2"])
}

func TestDeleteCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("var1: 1 var2: 0" +
		"cmd:delete, id:var1")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, float64(0), fslInterpreter.variables["var1"])
	assert.Equal(t, float64(0), fslInterpreter.variables["var2"])
}

func TestCreateCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("cmd:create, id:var3, value:5")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, float64(5), fslInterpreter.variables["var3"])
}

func TestAddCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("var1:0 cmd:add, id:var1, operand1:10, operand2:5")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, float64(15), fslInterpreter.variables["var1"])
}

func TestSubtractCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("var1:0 cmd:subtract, id:var1, operand1:10, operand2:5")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, float64(5), fslInterpreter.variables["var1"])
}

func TestMultiplyCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("var1:0 cmd:multiply, id:var1, operand1:10, operand2:5")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, float64(50), fslInterpreter.variables["var1"])
}

func TestDivideCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("var1:0 cmd:divide, id:var1, operand1:10, operand2:5")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, 0.5, fslInterpreter.variables["var1"])
}

func TestUserDefinedFunctionCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("sum: [\n\tcmd:add, id: $id, operand1: $value1, operand2: $value2\n]")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, 1, len(fslInterpreter.functions["sum"]))
}

func TestUserDefinedFunctionMultipleCommand(t *testing.T) {
	// Setup
	fslInterpreter := NewFSLInterpreter()
	// Create test input
	inputStream := antlr.NewInputStream("setup: [\n\tcmd:update, id: var1, value:3.5" +
		" cmd:print, value: #var1\n\tcmd:#sum, id: var1, value1:#var1, value2:#var2" +
		" cmd:print, value:#var1\n\tcmd:create, id:var3, value:5" +
		" cmd:delete, id:var1\n\tcmd:#printAll\n]")

	// Create the lexer
	lexer := parser.NewFSLLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewFSLParser(tokenStream)

	// Parse the input string
	tree := p.Program()

	// Execute
	for _, s := range tree.AllStatement() {
		fslInterpreter.visitStatement(s.(*parser.StatementContext))
	}

	// Assert
	assert.Equal(t, 7, len(fslInterpreter.functions["setup"]))
}
