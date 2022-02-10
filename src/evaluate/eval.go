package evaluate

import (
	"professorc/src/ast"
	"professorc/src/environment"
)

func Eval(node ast.Node, env *environment.Environment) {
	switch node := node.(type) {
	case *ast.Program:
		evalProgram(node, env)

	case *ast.YeoseokStatement:
		env.Student.RenewStack(int(node.Value))
	}
}

func evalProgram(program *ast.Program, env *environment.Environment) {
	for _, statement := range program.Statements {
		Eval(statement, env)
	}
}
