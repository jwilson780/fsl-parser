// Code generated from FSL.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // FSL
import "github.com/antlr4-go/antlr/v4"

// FSLListener is a complete listener for a parse tree produced by FSLParser.
type FSLListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterArithmeticOperation is called when entering the arithmeticOperation production.
	EnterArithmeticOperation(c *ArithmeticOperationContext)

	// EnterCrudOperation is called when entering the crudOperation production.
	EnterCrudOperation(c *CrudOperationContext)

	// EnterAddOperation is called when entering the addOperation production.
	EnterAddOperation(c *AddOperationContext)

	// EnterSubtractOperation is called when entering the subtractOperation production.
	EnterSubtractOperation(c *SubtractOperationContext)

	// EnterMultiplyOperation is called when entering the multiplyOperation production.
	EnterMultiplyOperation(c *MultiplyOperationContext)

	// EnterDivideOperation is called when entering the divideOperation production.
	EnterDivideOperation(c *DivideOperationContext)

	// EnterCreateOperation is called when entering the createOperation production.
	EnterCreateOperation(c *CreateOperationContext)

	// EnterUpdateOperation is called when entering the updateOperation production.
	EnterUpdateOperation(c *UpdateOperationContext)

	// EnterDeleteOperation is called when entering the deleteOperation production.
	EnterDeleteOperation(c *DeleteOperationContext)

	// EnterPrintOperation is called when entering the printOperation production.
	EnterPrintOperation(c *PrintOperationContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterPassedParameter is called when entering the passedParameter production.
	EnterPassedParameter(c *PassedParameterContext)

	// EnterOperationParameter is called when entering the operationParameter production.
	EnterOperationParameter(c *OperationParameterContext)

	// EnterIdentificationParameter is called when entering the identificationParameter production.
	EnterIdentificationParameter(c *IdentificationParameterContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterVariableReference is called when entering the variableReference production.
	EnterVariableReference(c *VariableReferenceContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitArithmeticOperation is called when exiting the arithmeticOperation production.
	ExitArithmeticOperation(c *ArithmeticOperationContext)

	// ExitCrudOperation is called when exiting the crudOperation production.
	ExitCrudOperation(c *CrudOperationContext)

	// ExitAddOperation is called when exiting the addOperation production.
	ExitAddOperation(c *AddOperationContext)

	// ExitSubtractOperation is called when exiting the subtractOperation production.
	ExitSubtractOperation(c *SubtractOperationContext)

	// ExitMultiplyOperation is called when exiting the multiplyOperation production.
	ExitMultiplyOperation(c *MultiplyOperationContext)

	// ExitDivideOperation is called when exiting the divideOperation production.
	ExitDivideOperation(c *DivideOperationContext)

	// ExitCreateOperation is called when exiting the createOperation production.
	ExitCreateOperation(c *CreateOperationContext)

	// ExitUpdateOperation is called when exiting the updateOperation production.
	ExitUpdateOperation(c *UpdateOperationContext)

	// ExitDeleteOperation is called when exiting the deleteOperation production.
	ExitDeleteOperation(c *DeleteOperationContext)

	// ExitPrintOperation is called when exiting the printOperation production.
	ExitPrintOperation(c *PrintOperationContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitPassedParameter is called when exiting the passedParameter production.
	ExitPassedParameter(c *PassedParameterContext)

	// ExitOperationParameter is called when exiting the operationParameter production.
	ExitOperationParameter(c *OperationParameterContext)

	// ExitIdentificationParameter is called when exiting the identificationParameter production.
	ExitIdentificationParameter(c *IdentificationParameterContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitVariableReference is called when exiting the variableReference production.
	ExitVariableReference(c *VariableReferenceContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
