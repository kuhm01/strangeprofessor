package lex

//https://github.com/rycont/umjunsik-lang/blob/master/umjunsik-lang-go/lexer/lexer.go
func (l *Lexer) readChar() {
	Byte1 := l.input[l.readPosition]

	if Byte1 >= 128 {
		l.ch = string([]byte{Byte1, l.input[l.readPosition+1], l.input[l.readPosition+2]})

		l.position = l.readPosition
		l.readPosition += 3
	} else {
		l.ch = string(Byte1)
		l.position = l.readPosition
		l.readPosition += 1
	}
}

//https://github.com/rycont/umjunsik-lang/blob/master/umjunsik-lang-go/lexer/lexer.go
func (l *Lexer) peekchar() string {
	if l.input[l.readPosition] >= 128 {
		return string([]byte{l.input[l.readPosition], l.input[l.readPosition+1], l.input[l.readPosition+2]})
	} else {
		return string(l.input[l.readPosition])
	}
}

func (l *Lexer) readKeyword() string {
	position := l.position
	for l.ch != "," {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
