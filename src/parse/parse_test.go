package parse

import (
	"fmt"
	"professorc/src/lex"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestParse(t *testing.T) {
	input := `
	교수님!! ABBCDA 주세요,
	전공 aster_big, 몰루,
	여석신청 12 자리,
	졸업
	`

	l := lex.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	for _, tt := range program.Statements {
		fmt.Printf("%s\n", tt.String())
	}
}

func TestIllegal(t *testing.T) {
	input := `
	몰루,
	전공 교수,
	졸업,
	`

	l := lex.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	for _, tt := range program.Statements {
		fmt.Printf("%s\n", tt.String())
	}
}

func TestIfflag(t *testing.T) {
	input := `
	교양,
	교양필수 여석신청 2 자리,
	전공 ast,
	교수님?? A 줘, 교수님.. 입학했습니다,
	재학생 전출,
	교양선택,
	전공 ca, 교수님? A 주세요, 교수님. 입학했습니다,
	재학생 전출,
	교양끝,
	졸업,
	`

	l := lex.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	for _, tt := range program.Statements {
		fmt.Printf("%s\n", tt.String())
	}
}
