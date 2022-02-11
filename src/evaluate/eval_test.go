package evaluate

import (
	"professorc/src/environment"
	"professorc/src/lex"
	"professorc/src/parse"
	"testing"
)

func TestEval(t *testing.T) {
	input := `전공 civa,
	여석신청 12 자리,
	여석신청 15 자리,
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
