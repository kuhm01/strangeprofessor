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
	교수님!!,
	전공, aster_big, 몰루,
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
