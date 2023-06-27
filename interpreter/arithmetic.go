package interpreter

import (
	"errors"
	"fsl-parser/parser"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

// VisitArithmeticOperation handles all arithmetic operations defined in the sample script
func (i *FSLInterpreter) VisitArithmeticOperation(ctx *parser.ArithmeticOperationContext) error {
	if ctx.AddOperation() != nil {
		err := i.visitAddOperation(ctx.AddOperation().(*parser.AddOperationContext))
		if err != nil {
			return err
		}
	} else if ctx.SubtractOperation() != nil {
		err := i.visitSubtractOperation(ctx.SubtractOperation().(*parser.SubtractOperationContext))
		if err != nil {
			return err
		}
	} else if ctx.MultiplyOperation() != nil {
		err := i.visitMultiplyOperation(ctx.MultiplyOperation().(*parser.MultiplyOperationContext))
		if err != nil {
			return err
		}
	} else if ctx.DivideOperation() != nil {
		err := i.visitDivideOperation(ctx.DivideOperation().(*parser.DivideOperationContext))
		if err != nil {
			return err
		}
	}
	return nil
}

// visitOperationParameterContext is a helper method to interpret the OperationParameterContext and return a value
func (i *FSLInterpreter) visitOperationParameterContext(ctx *parser.OperationParameterContext) (interface{}, error) {
	if ctx.Value() != nil {
		// basic int or float value
		v := ctx.Value().GetText()
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		return f, nil
	} else if ctx.VariableReference() != nil {
		return i.VisitVariableReference(ctx.VariableReference().(*parser.VariableReferenceContext))
	} else if ctx.FunctionParameter() != nil {
		return i.interpretFunctionParameter(ctx.FunctionParameter().(*parser.FunctionParameterContext))
	}
	return nil, errors.New("no operational parameter typed detected")
}

// visitIdentificationParameterContext is a helper method to interpret the IdentificationParameterContext and return a value
func (i *FSLInterpreter) visitIdentificationParameterContext(ctx *parser.IdentificationParameterContext) (interface{}, error) {
	if ctx.FunctionParameter() != nil {
		return i.interpretFunctionParameter(ctx.FunctionParameter().(*parser.FunctionParameterContext))
	} else if ctx.ID() != nil {
		return ctx.ID().GetText(), nil
	}
	return nil, errors.New("unknown identification parameter context")
}

func (i *FSLInterpreter) interpretFunctionParameter(ctx *parser.FunctionParameterContext) (interface{}, error) {
	v := ctx.GetText()[1:] // params are prefaced with a $
	value, ok := i.params[v]
	if !ok {
		return nil, errors.New("parameter does not exist")
	}
	// if value is a reference, we look it up
	if value.VariableReference() != nil {
		return i.VisitVariableReference(value.VariableReference().(*parser.VariableReferenceContext))
	} else {
		// value is an ID so just grab the text
		return value.GetText(), nil
	}
}

// visitAddOperation Adds two numbers and sets their value in the id parameter
func (i *FSLInterpreter) visitAddOperation(ctx *parser.AddOperationContext) error {
	dest, op1, op2, err := i.parseArithmeticParams(ctx.GetChildren())
	if err != nil {
		return nil
	}
	i.variables[dest.(string)] = op1.(float64) + op2.(float64)
	return nil
}

// visitSubtractOperation Subtract two numbers and sets their value in the id parameter
func (i *FSLInterpreter) visitSubtractOperation(ctx *parser.SubtractOperationContext) error {
	dest, op1, op2, err := i.parseArithmeticParams(ctx.GetChildren())
	if err != nil {
		return nil
	}
	i.variables[dest.(string)] = op1.(float64) - op2.(float64)
	return nil
}

// visitMultiplyOperation Multiply two numbers and sets their value in the id parameter
func (i *FSLInterpreter) visitMultiplyOperation(ctx *parser.MultiplyOperationContext) error {
	dest, op1, op2, err := i.parseArithmeticParams(ctx.GetChildren())
	if err != nil {
		return nil
	}
	i.variables[dest.(string)] = op1.(float64) * op2.(float64)
	return nil
}

// visitDivideOperation Divide two numbers and sets their value in the id parameter
func (i *FSLInterpreter) visitDivideOperation(ctx *parser.DivideOperationContext) error {
	dest, op1, op2, err := i.parseArithmeticParams(ctx.GetChildren())
	if err != nil {
		return nil
	}
	i.variables[dest.(string)] = op2.(float64) / op1.(float64)
	return nil
}

// parseArithmeticParams pulls the parameters from known positions in the child tree
func (i *FSLInterpreter) parseArithmeticParams(children []antlr.Tree) (interface{}, interface{}, interface{}, error) {
	// Gather Parameters
	dest, err := i.visitIdentificationParameterContext(children[4].(*parser.IdentificationParameterContext))
	if err != nil {
		return nil, nil, nil, err
	}

	op1, err := i.visitOperationParameterContext(children[8].(*parser.OperationParameterContext))
	if err != nil {
		return nil, nil, nil, err
	}

	op2, err := i.visitOperationParameterContext(children[12].(*parser.OperationParameterContext))
	if err != nil {
		return nil, nil, nil, err
	}

	return dest, op1, op2, nil
}
