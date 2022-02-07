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
	default:

	}

	l.readChar()
	return tok
}
