package evaluate

import (
	"professorc/src/ast"
	"professorc/src/environment"
)

func evalLiberalStatement(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	n := node.(*ast.LiberalStatement)
	var n1v, n2v int
	n1, b := env.Professor.Response_Interview(1)
	if !b {
		n1v = 0
	} else {
		n1v = n1.Value
	}

	n2, b := env.Professor.Response_Interview(2)
	if !b {
		n2v = 0
	} else {
		n2v = n2.Value
	}

	if n1v < n2v {
		evalBlockStatement(n.Mandatory, env, pc)
	} else {
		evalBlockStatement(n.Selection, env, pc)
	}
}
