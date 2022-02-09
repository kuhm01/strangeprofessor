package lex

import (
	"professorc/src/token"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           string
	len          int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.len = utf8.RuneCountInString(input)
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case "!":
		tok = newToken(token.BANG, l.ch)
	case "?":
		tok = newToken(token.QUESTION, l.ch)
	case ",":
		tok = newToken(token.COLON, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal, tok.Type = l.readIdentifier()
			return tok
		} else {
			tok.Literal = l.readKeyword()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
	}

	l.readChar()
	return tok
}
