package parse

import (
	"fmt"
	"professorc/src/ast"
	"professorc/src/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.YEOSEOK:
		return p.parseYeoseokStatement()

	case token.STATICVAR:
		ps := p.curToken.Literal

		switch ps[9] {
		case '!':
			return p.parseInBufferFromStaticVarStatement()
		case '?':
			return p.parseStaticVarOperatingStatement()
		case '.':
			return p.parseTokStackStatement()
		}

	case token.FLAG:
		return p.parseMajorFlag()

	case token.JUMPPOINT:
		return p.parseJumptoMajorFlag()

	case token.PRINTER:
		return &ast.PrintBufferStatement{Token: p.curToken}

	case token.PLUSCLASS:
		return &ast.ClassStatement{Token: p.curToken, Type: "보강"}

	case token.MINUSCLASS:
		return &ast.ClassStatement{Token: p.curToken, Type: "휴강"}

	case token.NEWEMPLOY:
		return &ast.NewProfessorStatement{Token: p.curToken}

	case token.DYNAMICVAR:
		return p.parseStudentStatement()

	case token.SURYO:
		return p.parseTheEnd()

	case token.ILLEGAL:
		msg := fmt.Sprintf("%s is ILLEGAL\n", p.curToken.Literal)
		p.errors = append(p.errors, msg)
	}

	return nil
}
