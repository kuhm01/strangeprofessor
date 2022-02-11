package parse

import (
	"fmt"
	"professorc/src/ast"
	"professorc/src/token"
	"strconv"
)

func (p *Parser) parseYeoseokStatement() *ast.YeoseokStatement {
	stmt := &ast.YeoseokStatement{Token: p.curToken}

	if p.peekTokenIs(token.INT) {
		p.nextToken()
		value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
		if err != nil {
			msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
			p.errors = append(p.errors, msg)
			return nil
		}
		stmt.Value = value
		if !p.expectPeek(token.JARI) {
			return nil
		}
	} else {
		return nil
	}

	return stmt
}

func (p *Parser) parseStaticVarOperatingStatement() *ast.StaticVarStatement {
	return nil
}

func (p *Parser) parseMajorFlag() *ast.MajorFlagStatement {
	stmt := &ast.MajorFlagStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Identifier = p.curToken.Literal

	return stmt
}

func (p *Parser) parseJumptoMajorFlag() *ast.JumptoFlagStatement {
	stmt := &ast.JumptoFlagStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.JumpIdentifier = p.curToken.Literal

	return stmt
}
