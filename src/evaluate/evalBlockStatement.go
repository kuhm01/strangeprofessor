package evaluate

import (
	"professorc/src/ast"
	"professorc/src/environment"
)

func evalBlockStatement(block *ast.BlockStatement, env *environment.Environment, outpc *ProgramCounter) {
	inpc := NewPC(len(block.Statements))
	inpc.Outer = outpc

	for inpc.Over() {
		if !inpc.isThisPlused() {
			Eval(block.Statements[inpc.Index], env, inpc)
		}
		inpc.Index++
	}
}
