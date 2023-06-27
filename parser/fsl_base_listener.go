// Code generated from FSL.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // FSL
import "github.com/antlr4-go/antlr/v4"

// BaseFSLListener is a complete listener for a parse tree produced by FSLParser.
type BaseFSLListener struct{}

var _ FSLListener = &BaseFSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseFSLListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseFSLListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseFSLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseFSLListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseFSLListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseFSLListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseFSLListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseFSLListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseFSLListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseFSLListener) ExitCommand(ctx *CommandContext) {}

// EnterArithmeticOperation is called when production arithmeticOperation is entered.
func (s *BaseFSLListener) EnterArithmeticOperation(ctx *ArithmeticOperationContext) {}

// ExitArithmeticOperation is called when production arithmeticOperation is exited.
func (s *BaseFSLListener) ExitArithmeticOperation(ctx *ArithmeticOperationContext) {}

// EnterCrudOperation is called when production crudOperation is entered.
func (s *BaseFSLListener) EnterCrudOperation(ctx *CrudOperationContext) {}

// ExitCrudOperation is called when production crudOperation is exited.
func (s *BaseFSLListener) ExitCrudOperation(ctx *CrudOperationContext) {}

// EnterAddOperation is called when production addOperation is entered.
func (s *BaseFSLListener) EnterAddOperation(ctx *AddOperationContext) {}

// ExitAddOperation is called when production addOperation is exited.
func (s *BaseFSLListener) ExitAddOperation(ctx *AddOperationContext) {}

// EnterSubtractOperation is called when production subtractOperation is entered.
func (s *BaseFSLListener) EnterSubtractOperation(ctx *SubtractOperationContext) {}

// ExitSubtractOperation is called when production subtractOperation is exited.
func (s *BaseFSLListener) ExitSubtractOperation(ctx *SubtractOperationContext) {}

// EnterMultiplyOperation is called when production multiplyOperation is entered.
func (s *BaseFSLListener) EnterMultiplyOperation(ctx *MultiplyOperationContext) {}

// ExitMultiplyOperation is called when production multiplyOperation is exited.
func (s *BaseFSLListener) ExitMultiplyOperation(ctx *MultiplyOperationContext) {}

// EnterDivideOperation is called when production divideOperation is entered.
func (s *BaseFSLListener) EnterDivideOperation(ctx *DivideOperationContext) {}

// ExitDivideOperation is called when production divideOperation is exited.
func (s *BaseFSLListener) ExitDivideOperation(ctx *DivideOperationContext) {}

// EnterCreateOperation is called when production createOperation is entered.
func (s *BaseFSLListener) EnterCreateOperation(ctx *CreateOperationContext) {}

// ExitCreateOperation is called when production createOperation is exited.
func (s *BaseFSLListener) ExitCreateOperation(ctx *CreateOperationContext) {}

// EnterUpdateOperation is called when production updateOperation is entered.
func (s *BaseFSLListener) EnterUpdateOperation(ctx *UpdateOperationContext) {}

// ExitUpdateOperation is called when production updateOperation is exited.
func (s *BaseFSLListener) ExitUpdateOperation(ctx *UpdateOperationContext) {}

// EnterDeleteOperation is called when production deleteOperation is entered.
func (s *BaseFSLListener) EnterDeleteOperation(ctx *DeleteOperationContext) {}

// ExitDeleteOperation is called when production deleteOperation is exited.
func (s *BaseFSLListener) ExitDeleteOperation(ctx *DeleteOperationContext) {}

// EnterPrintOperation is called when production printOperation is entered.
func (s *BaseFSLListener) EnterPrintOperation(ctx *PrintOperationContext) {}

// ExitPrintOperation is called when production printOperation is exited.
func (s *BaseFSLListener) ExitPrintOperation(ctx *PrintOperationContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseFSLListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseFSLListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterPassedParameter is called when production passedParameter is entered.
func (s *BaseFSLListener) EnterPassedParameter(ctx *PassedParameterContext) {}

// ExitPassedParameter is called when production passedParameter is exited.
func (s *BaseFSLListener) ExitPassedParameter(ctx *PassedParameterContext) {}

// EnterOperationParameter is called when production operationParameter is entered.
func (s *BaseFSLListener) EnterOperationParameter(ctx *OperationParameterContext) {}

// ExitOperationParameter is called when production operationParameter is exited.
func (s *BaseFSLListener) ExitOperationParameter(ctx *OperationParameterContext) {}

// EnterIdentificationParameter is called when production identificationParameter is entered.
func (s *BaseFSLListener) EnterIdentificationParameter(ctx *IdentificationParameterContext) {}

// ExitIdentificationParameter is called when production identificationParameter is exited.
func (s *BaseFSLListener) ExitIdentificationParameter(ctx *IdentificationParameterContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BaseFSLListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BaseFSLListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterVariableReference is called when production variableReference is entered.
func (s *BaseFSLListener) EnterVariableReference(ctx *VariableReferenceContext) {}

// ExitVariableReference is called when production variableReference is exited.
func (s *BaseFSLListener) ExitVariableReference(ctx *VariableReferenceContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFSLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFSLListener) ExitValue(ctx *ValueContext) {}
