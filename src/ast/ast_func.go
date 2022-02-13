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

func (sv *StaticVarStatement) statementNode()       {}
func (sv *StaticVarStatement) TokenLiteral() string { return sv.Token.Literal }
func (sv *StaticVarStatement) String() string       { return sv.Token.Literal }

func (psv *ProfessorTobuffStatement) statementNode()       {}
func (psv *ProfessorTobuffStatement) TokenLiteral() string { return psv.Token.Literal }
func (psv *ProfessorTobuffStatement) String() string       { return psv.Token.Literal }

func (mf *MajorFlagStatement) statementNode()       {}
func (mf *MajorFlagStatement) TokenLiteral() string { return mf.Token.Literal }
func (mf *MajorFlagStatement) String() string       { return mf.Token.Literal }

func (jf *JumptoFlagStatement) statementNode()       {}
func (jf *JumptoFlagStatement) TokenLiteral() string { return jf.Token.Literal }
func (jf *JumptoFlagStatement) String() string       { return jf.Token.Literal }

func (c *ConstantExpression) expressionNode()      {}
func (c *ConstantExpression) TokenLiteral() string { return c.Token.Literal }
func (c *ConstantExpression) String() string       { return c.Token.Literal }
