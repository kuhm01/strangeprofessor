package evaluate

import (
	"professorc/src/environment"
	"professorc/src/lex"
	"professorc/src/parse"
	"testing"
)

func TestEval(t *testing.T) {
	input := `
	교수님??? AAAAAA 내놔,
	교수님!!! 성적발표,
	교수님??? AA 내놔,
	교수님??? A 주세요,
	교수님??? D 주세요,
	교수님!!! 성적발표,
	교수님??? AB 주세요,
	교수님!!! 성적발표,
	교수님!!! 성적발표,
	교수님??? B 주세요,
	교수님!!! 성적발표,
	교수님? AAAA 줘,
	교수님! 성적발표,
	교수님?? AB 주세요,
	교수님?? AAAAAAAAAA 줘,
	교수님!! 성적발표,
	교수님!!! 성적발표,
	교수님??? B 주세요,
	교수님!!! 성적발표,
	교수님??? FFF 줘,
	교수님!!! 성적발표,
	교수님??? FFFF 줘,
	교수님!!! 성적발표,
	교수님???? AAC 주세요,
	교수님!!!! 성적발표,
	공지,
	졸업
	`

	l := lex.New(input)
	p := parse.New(l)

	program := p.ParseProgram()
	env := environment.New()
	pc := NewPC(len(program.Statements))

	Eval(program, env, pc)
}

func TestClassEval(t *testing.T) {
	input := `
	휴강,
	여석신청 15 자리,
	졸업
	`

	l := lex.New(input)
	p := parse.New(l)

	program := p.ParseProgram()
	env := environment.New()
	pc := NewPC(len(program.Statements))

	Eval(program, env, pc)
}

func TestNewProfessor(t *testing.T) {
	input := `
	교수님?? A 줘,
	교수님??? A 줘,
	신규임용,
	교수님??? A 줘,
	졸업
	`

	l := lex.New(input)
	p := parse.New(l)

	program := p.ParseProgram()
	env := environment.New()
	pc := NewPC(len(program.Statements))

	Eval(program, env, pc)
}
