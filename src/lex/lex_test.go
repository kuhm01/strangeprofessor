package lex

import (
	"professorc/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	교수님!!,
	교수님???,
	전공 aster_big 몰루,
	ABDCCDF
	123
	교수님?!,
	졸업
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STATICVAR, "교수님!!"},
		{token.COLON, ","},
		{token.STATICVAR, "교수님???"},
		{token.COLON, ","},
		{token.FLAG, "전공"},
		{token.IDENT, "aster_big"},
		{token.ILLEGAL, "몰루"},
		{token.COLON, ","},
		{token.CONSTINT, "ABDCCDF"},
		{token.INT, "123"},
		{token.STATICVAR, "교수님?!"},
		{token.COLON, ","},
		{token.EXITPOINT, "졸업"},
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
