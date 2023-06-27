package interpreter

import (
	"fsl-parser/parser"
	"strconv"
)

// VisitCrudOperation handles crud based operations for a given parent context
func (i *FSLInterpreter) VisitCrudOperation(ctx *parser.CrudOperationContext) error {
	if ctx.CreateOperation() != nil {
		// create: make a new var from the tree and set it in the map
		op := ctx.CreateOperation().(*parser.CreateOperationContext)
		id := op.ID().GetText()
		value := op.Value().GetText()
		v, err := strconv.ParseFloat(value, 64) // will always be a new float or int per the grammar
		if err != nil {
			return err
		}
		i.variables[id] = v
	} else if ctx.UpdateOperation() != nil {
		// update: get var from the tree and update it in the map
		op := ctx.UpdateOperation().(*parser.UpdateOperationContext)
		id := op.ID().GetText()
		if op.Value() != nil {
			// new int or float value so just set in map
			value := op.Value().GetText()
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			i.variables[id] = v
		} else if op.VariableReference() != nil {
			// lookup known value and set on variable
			reference, err := i.VisitVariableReference(op.VariableReference().(*parser.VariableReferenceContext))
			if err != nil {
				return err
			}
			i.variables[id] = reference
		}
	} else if ctx.DeleteOperation() != nil {
		op := ctx.DeleteOperation().(*parser.DeleteOperationContext)
		id := op.ID().GetText()
		delete(i.variables, id)
	}
	return nil
}
