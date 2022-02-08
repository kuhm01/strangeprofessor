package lex

func (l *Lexer) skipWhitespace() {
	for l.ch == " " || l.ch == "\t" || l.ch == "\n" || l.ch == "\r" {
		l.readChar()
	}
}

func isLetter(ch string) bool {
	return "a" <= ch && ch <= "z" || "A" <= ch && ch <= "Z" || ch == "_"
}
