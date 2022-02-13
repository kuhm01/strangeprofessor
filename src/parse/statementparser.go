package parse

import (
	"fmt"
	"professorc/src/ast"
	"professorc/src/environment"
	"professorc/src/token"
	"strconv"
	"strings"
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
	stmt := &ast.StaticVarStatement{Token: p.curToken}

	ps := p.curToken.Literal

	iv := strings.Count(ps, "?")

	stmt.Index = iv

	if !p.expectPeek(token.CONSTINT) {
		return nil
	}

	stmt.Constant = p.parseExpression()

	if !p.expectPeek(token.GIVEME) {
		return nil
	}

	s := p.curToken.Literal
	switch s {
	case "주세요":
		stmt.Level = 1
	case "줘":
		stmt.Level = 2
	case "내놔":
		stmt.Level = 3
	default:
		return nil
	}

	return stmt
}

func (p *Parser) parseInBufferFromStaticVarStatement() *ast.ProfessorTobuffStatement {
	stmt := &ast.ProfessorTobuffStatement{Token: p.curToken}

	ps := p.curToken.Literal

	iv := strings.Count(ps, "!")
	stmt.Index = iv

	if !p.expectPeek(token.P2PRINTER) {
		return nil
	}

	s := p.curToken.Literal
	switch s {
	case "점수발표":
		stmt.Type = environment.Int
	case "성적발표":
		stmt.Type = environment.Char
	}

	return stmt
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
