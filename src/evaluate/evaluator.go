package evaluate

import (
	"professorc/src/ast"
	"professorc/src/environment"
)

func evalStaticVarStatement(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	stmt := node.(*ast.StaticVarStatement)
	index := stmt.Index
	level := stmt.Level
	value := stmt.Constant.(*ast.ConstantExpression).Value
	resultvalue := value * int64(level)
}
