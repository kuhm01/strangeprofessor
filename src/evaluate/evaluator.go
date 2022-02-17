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

func evalPrintBuffer(env *environment.Environment, pc *ProgramCounter) {
	for {
		e := env.Buffer.GetinBuffer()
		if e == nil {
			break
		}
		v := e.Node.Value
		typer := e.PrintType
		switch typer {
		case environment.Char:
			fmt.Printf("%c", v)
		case environment.Int:
			fmt.Printf("%d", v)
		}
	}
}

func evalClass(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	n := node.(*ast.ClassStatement)
	switch n.Type {
	case "보강":
		evalPlusClass(env, pc)
	case "휴강":
		evalMinusClass(env, pc)
	}
}

func evalMinusClass(env *environment.Environment, pc *ProgramCounter) {
	pc.Index++
}

func evalPlusClass(env *environment.Environment, pc *ProgramCounter) {
	pc.ThisisPlused()
	pc.Index--
	pc.Index--
}

func evalNewProfessor(env *environment.Environment) {
	for _, pp := range env.Professor.Professor {
		if pp == nil {
			continue
		}
		pp.Value = 0
	}
}

func evalStack(Type string, env *environment.Environment) {
	switch Type {
	case "전출":
		n, b := env.Student.Pop()
		if !b {
			fmt.Printf("Student is not exist\n")
			return
		}

		env.Changed.Set(n)

	case "전입":
		n, b := env.Changed.Get()
		if !b {
			fmt.Printf("There is No changed student\n")
			return
		}

		env.Student.Push(n)
	}
}

func evalPtoS(node ast.Node, env *environment.Environment) {
	n := node.(*ast.ProfessorToStudentStatement)
	index := n.Index - 1

	if index > 6 {
		fmt.Printf("Not exist %dth Professor.", index)
		return
	}

	newNode := env.Professor.Professor[index]
	newCopiedNode := &environment.Node{Value: newNode.Value}

	env.Student.Push(newCopiedNode)
}

func evalStoP(node ast.Node, env *environment.Environment) {
	n := node.(*ast.StudentToProfessorStatement)
	index := n.Index - 1

	if index > 6 {
		fmt.Printf("Not exist %dth Professor.", index)
		return
	}

	newNode, b := env.Student.Pop()
	if !b {
		fmt.Printf("Student is not exist\n")
	}

	newCopiedNode := &environment.Node{Value: newNode.Value}

	env.Professor.Professor[index] = newCopiedNode
}

func evalPtomB(node ast.Node, env *environment.Environment) {
	n := node.(*ast.ProfessorTominiBufferStatement)
	index := n.Index

	if index > 6 {
		fmt.Printf("Not exist %dth Professor.", index)
		return
	}

	newNode, b := env.Professor.Response_Interview(index)
	if !b {
		fmt.Printf("Haven't value. %dth Professor.", index)
		return
	}

	env.MiniBuffer.Insert(newNode)
}
