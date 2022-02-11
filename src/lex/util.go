package lex

func (l *Lexer) skipWhitespace() {
	for l.ch == "\t" || l.ch == "\n" || l.ch == "\r" {
		l.readChar()
	}
}

func isLetter(ch string) bool {
	return "a" <= ch && ch <= "z" || ch == "_"
}

func isDigit(ch string) bool {
	return "0" <= ch && ch <= "9"
}

func isConst(ch string) bool {
	return "A" <= ch && ch <= "D" || ch == "F"
}
