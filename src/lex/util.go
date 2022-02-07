package lex

func (l *Lexer) skipWhitespace() {
	for l.ch == " " || l.ch == "\t" || l.ch == "\n" || l.ch == "\r" {
		l.readChar()
	}
}
