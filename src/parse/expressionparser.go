package parse

import (
	"fmt"
	"professorc/src/ast"
	"professorc/src/token"
)

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression() ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	return leftExp
}

func (p *Parser) parseConstantExpression() ast.Expression {
	exp := &ast.ConstantExpression{Token: p.curToken}
	sv := p.curToken.Literal
	result := 0
	for i := 0; i < len(sv); i++ {
		switch sv[i] {
		case 'A':
			result += 4
		case 'B':
			result += 3
		case 'C':
			result += 2
		case 'D':
			result += 1
		case 'F':
			result -= 1
		}
	}
	exp.Value = int64(result)
	return exp
}
