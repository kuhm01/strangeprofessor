package lex

import (
	"professorc/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	교수님!!,
	전공, aster_big,
	졸업,
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STATICVAR, "교수님!!"},
		{token.COLON, ","},
		{token.FLAG, "전공"},
		{token.COLON, ","},
		{token.IDENT, "aster_big"},
		{token.COLON, ","},
		{token.EXITPOINT, "졸업"},
		{token.COLON, ","},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype worng. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
