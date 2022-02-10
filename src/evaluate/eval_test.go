package evaluate

import (
	"professorc/src/environment"
	"professorc/src/lex"
	"professorc/src/parse"
	"testing"
)

func TestEval(t *testing.T) {
	input := `여석신청 12 자리,
	여석신청 15 자리,
	교수님!!,
	졸업
	`

	l := lex.New(input)
	p := parse.New(l)
	program := p.ParseProgram()
	env := environment.New()
	Eval(program, env)
}
