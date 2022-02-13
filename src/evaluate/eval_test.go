package evaluate

import (
	"professorc/src/environment"
	"professorc/src/lex"
	"professorc/src/parse"
	"testing"
)

func TestEval(t *testing.T) {
	input := `
	교수님??? AAA 줘,
	재수강 civa,
	졸업
	`

	l := lex.New(input)
	p := parse.New(l)

	program := p.ParseProgram()
	env := environment.New()
	pc := NewPC(len(program.Statements))

	Eval(program, env, pc)
}
