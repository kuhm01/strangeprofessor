package evaluate

import (
	"fmt"
	"professorc/src/ast"
	"professorc/src/environment"
)

func evalStaticVarStatement(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	stmt := node.(*ast.StaticVarStatement)
	index := stmt.Index
	if index > 6 {
		fmt.Printf("Not exist %dth Professor.", index)
		return
	}
	level := stmt.Level
	value := stmt.Constant.(*ast.ConstantExpression).Value
	resultvalue := value * int64(level)
	env.Professor.Setter(int(resultvalue), index)
}

func evalBuffFromProfessor(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	stmt := node.(*ast.ProfessorTobuffStatement)
	index := stmt.Index
	if index > 6 {
		fmt.Printf("Not exist %dth Professor.", index)
		return
	}
	jisija := stmt.Type
	n, boolean := env.Professor.Response_Interview(index)
	if !boolean {
		fmt.Printf("Professor don't response your request")
		return
	}

	ppr := &environment.PrintCharacter{Node: n, PrintType: jisija}
	env.Buffer.SetinBuffer(ppr)
}
