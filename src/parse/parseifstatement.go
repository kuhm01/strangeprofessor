package parse

import (
	"professorc/src/ast"
	"professorc/src/token"
)

func (p *Parser) parseIfStatement() *ast.LiberalStatement {
	stmt := &ast.LiberalStatement{Token: p.curToken}

	p.nextToken()

	if !p.expectPeek(token.IFFLAG) {
		return nil
	}

	stmt.Mandatory = p.parseBlockStatement()
	stmt.Selection = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.ELSEFLAG) && !p.curTokenIs(token.IFFLAGEXIT) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}
