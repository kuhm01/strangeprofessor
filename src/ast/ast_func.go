package ast

import "bytes"

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (y *YeoseokStatement) statementNode()       {}
func (y *YeoseokStatement) TokenLiteral() string { return y.Token.Literal }
func (y *YeoseokStatement) String() string       { return y.Token.Literal }
