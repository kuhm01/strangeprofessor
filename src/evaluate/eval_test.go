package evaluate

import (
	"professorc/src/lex"
	"professorc/src/parse"
	"testing"
)

func TestEval(t *testing.T) {
	input := `여석신청 12 자리, 졸업
	`

	l := lex.New(input)
	p := parse.New(l)
	program := p.ParseProgram()
	Eval(program)
}
