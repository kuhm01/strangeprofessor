package parse

import (
	"fmt"
	"professorc/src/token"
)

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	if p.curToken.Type == token.EXITPOINT {
		p.peekToken.Type = ""
		p.peekToken.Literal = ""
	} else {
		p.peekToken = p.l.NextToken()
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got=%s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
