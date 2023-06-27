package interpreter

import (
	"errors"
	"fmt"
	"fsl-parser/parser"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"strconv"
)

const INIT string = "init"

type FSLInterpreter struct {
	variables map[string]float64
	functions map[string][]parser.ICommandContext
	params    map[string]parser.PassedParameterContext // specifically houses parameters for user defined functions
}

// NewFSLInterpreter creates an empty FSLInterpreter
func NewFSLInterpreter() *FSLInterpreter {
	return &FSLInterpreter{
		variables: make(map[string]float64),
		functions: make(map[string][]parser.ICommandContext),
		params:    make(map[string]parser.PassedParameterContext),
	}
}

// VisitProgram entrypoint for walking tree using the interpreter, specifically calls the init function last
func (i *FSLInterpreter) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, stmt := range ctx.AllStatement() {
		err := i.visitStatement(stmt.(*parser.StatementContext))
		if err != nil {
			fmt.Printf("Error interpreting statment: %v\n", err)
			os.Exit(1)
		}
	}
	// call the init function after entire script is interpreted
	err := i.callFunction(INIT)
	if err != nil {
		fmt.Printf("Error interpreting init function: %v\n", err)
		os.Exit(1)
	}
	return nil
}

// VisitStatement entrypoint for a Statement, the possible types are defined in the grammar
func (i *FSLInterpreter) visitStatement(ctx *parser.StatementContext) error {
	if ctx.VariableDeclaration() != nil {
		err := i.visitVariableDeclaration(ctx.VariableDeclaration().(*parser.VariableDeclarationContext))
		if err != nil {
			return err
		}
	} else if ctx.Command() != nil {
		err := i.visitCommand(ctx.Command().(*parser.CommandContext))
		if err != nil {
			return err
		}
	} else if ctx.FunctionDefinition() != nil {
		i.visitFunctionDefinition(ctx.FunctionDefinition().(*parser.FunctionDefinitionContext))
	}
	return nil
}

// visitVariableDeclaration converts a variable to a float64 and stores it in the variables map
func (i *FSLInterpreter) visitVariableDeclaration(ctx *parser.VariableDeclarationContext) error {
	varName := ctx.ID().GetText()
	varValue, err := strconv.ParseFloat(ctx.Value().GetText(), 64)
	if err != nil {
		return err
	}
	i.variables[varName] = varValue
	return nil
}

// visitCommand checks the command type and calls the top level evaluator for that type
func (i *FSLInterpreter) visitCommand(ctx *parser.CommandContext) error {
	if ctx.ArithmeticOperation() != nil {
		err := i.VisitArithmeticOperation(ctx.ArithmeticOperation().(*parser.ArithmeticOperationContext))
		if err != nil {
			return err
		}
	} else if ctx.CrudOperation() != nil {
		err := i.VisitCrudOperation(ctx.CrudOperation().(*parser.CrudOperationContext))
		if err != nil {
			return err
		}
	} else if ctx.PrintOperation() != nil {
		i.visitPrintOperation(ctx.PrintOperation().(*parser.PrintOperationContext))
	} else {
		// user defined function or unknown command context type
		err := i.visitFunctionCall(ctx.FunctionCall().(*parser.FunctionCallContext))
		if err != nil {
			return err
		}
	}
	return nil
}

// visitFunctionDefinition handles storing user defined functions in memory while a script runs
func (i *FSLInterpreter) visitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	functionName := ctx.GetChild(0).(antlr.ParseTree).GetText()

	// check to see if function has been defined before
	_, ok := i.functions[functionName]
	if ok {
		// erase function to allow to overwrite with new function commands
		delete(i.functions, functionName)
	}

	/* functions are structured as X: [ cmd ] <- as seen in the FSL.g4 grammar
	   child 0 -> X, child 1 -> :, child 2 -> [, child 3+ are command functions
	   store all commands until the last child, which is the closing bracket */
	for x := 3; x < ctx.GetChildCount()-1; x++ {
		child := ctx.GetChild(x).(*parser.CommandContext)
		i.functions[functionName] = append(i.functions[functionName], child)
	}
}

// visitFunctionCall calls a function for a given tree context, will store params for the duration of the script
func (i *FSLInterpreter) visitFunctionCall(ctx *parser.FunctionCallContext) error {
	functionName := ctx.GetChild(1).(antlr.ParseTree).GetText()
	/*
		User defined functions may have parameters passed when called
		e.g. cmd:#sum, id:var1, value1:#var1, value2:#var2
		All parameter references are parsed as type ParameterContext,
		but a faster way to parse them is to exploit that are always all at odd indices, starting at 3
		TODO: With a more general language all ParameterContext in order
	*/
	for x := 3; x < len(ctx.GetChildren()); x += 4 {
		id := ctx.GetChild(x).(antlr.ParseTree).GetText()
		value := ctx.GetChild(x + 2).(*parser.PassedParameterContext)
		i.params[id] = *value
	}

	// actually call the function
	err := i.callFunction(functionName)
	if err != nil {
		return err
	}
	return nil
}

// callFunction calls a known function's commands in order
func (i *FSLInterpreter) callFunction(functionName string) error {
	functionCommands, ok := i.functions[functionName]
	if !ok {
		return errors.New(fmt.Sprintf("Undefined function: %s\n", functionName))
	}

	// functions are comprised of 1...M commands
	for _, cmd := range functionCommands {
		err := i.visitCommand(cmd.(*parser.CommandContext))
		if err != nil {
			return err
		}
	}
	return nil
}

// visitPrintOperation prints the given tree context value to the console
func (i *FSLInterpreter) visitPrintOperation(ctx *parser.PrintOperationContext) {
	if ctx.Value() != nil {
		// context is a value
		fmt.Println(ctx.Value().GetText())
	} else if ctx.VariableReference() != nil {
		// context is a reference to a value
		reference, err := i.VisitVariableReference(ctx.VariableReference().(*parser.VariableReferenceContext))
		if err != nil {
			// sample script shows unknown vars printed as undefined in the output
			fmt.Println("undefined")
			return
		}
		fmt.Println(reference)
	}
}

// VisitVariableReference looks up a variable in memory and returns it
func (i *FSLInterpreter) VisitVariableReference(ctx *parser.VariableReferenceContext) (float64, error) {
	refID := ctx.ID().GetText()
	value, ok := i.variables[refID]
	if !ok {
		return 0, errors.New("error: variable reference value not found in map")
	}
	return value, nil
}
