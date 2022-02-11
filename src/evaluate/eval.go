package evaluate

import (
	"professorc/src/ast"
	"professorc/src/environment"
)

func Eval(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	switch node := node.(type) {
	case *ast.Program:
		evalProgram(node, env, pc)

	case *ast.YeoseokStatement:
		env.Student.RenewStack(int(node.Value))

	case *ast.MajorFlagStatement:
		pc.NewTag(node.Identifier)

	case *ast.JumptoFlagStatement:
		pc.SetCounter(node.JumpIdentifier)
	}
}

func evalProgram(program *ast.Program, env *environment.Environment, pc *ProgramCounter) {
	for pc.Over() {
		Eval(program.Statements[pc.Index], env, pc)
		pc.Index++
	}
}