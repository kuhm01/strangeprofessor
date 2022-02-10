package evaluate

import (
	"io"
	"os"
	"professorc/src/ast"
)

func Eval(node ast.Node) {
	switch node := node.(type) {
	case *ast.Program:
		evalProgram(node)

	case *ast.YeoseokStatement:
		io.WriteString(os.Stdout, "여석신청이요!")
	}
}

func evalProgram(program *ast.Program) {
	for _, statement := range program.Statements {
		Eval(statement)
	}
}
