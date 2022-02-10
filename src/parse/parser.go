package parse

import (
	"professorc/src/ast"
	"professorc/src/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.YEOSEOK:
		return p.parseYeoseokStatement()
	case token.STATICVAR:
		return p.parseStaticVarOperatingStatement()
	}

	return nil
}
