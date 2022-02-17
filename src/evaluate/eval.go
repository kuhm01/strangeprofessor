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

	case *ast.StaticVarStatement:
		evalStaticVarStatement(node, env, pc)

	case *ast.ProfessorTobuffStatement:
		evalBuffFromProfessor(node, env, pc)

	case *ast.PrintBufferStatement:
		evalPrintBuffer(env, pc)

	case *ast.ClassStatement:
		evalClass(node, env, pc)

	case *ast.NewProfessorStatement:
		evalNewProfessor(env)

	case *ast.StudentStatement:
		evalStack(node.Type, env)

	case *ast.ProfessorToStudentStatement:
		evalPtoS(node, env)

	case *ast.StudentToProfessorStatement:
		evalStoP(node, env)

	case *ast.BlockStatement:
		evalBlockStatement(node, env, pc)

	case *ast.ProfessorTominiBufferStatement:
		evalPtomB(node, env)
	}
}

func evalProgram(program *ast.Program, env *environment.Environment, pc *ProgramCounter) {
	for pc.Over() {
		if !pc.isThisPlused() {
			Eval(program.Statements[pc.Index], env, pc)
		}
		pc.Index++
	}
}
