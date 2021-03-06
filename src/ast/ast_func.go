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

func (pb *PrintBufferStatement) statementNode()       {}
func (pb *PrintBufferStatement) TokenLiteral() string { return pb.Token.Literal }
func (pb *PrintBufferStatement) String() string       { return pb.Token.Literal }

func (c *ClassStatement) statementNode()       {}
func (c *ClassStatement) TokenLiteral() string { return c.Token.Literal }
func (c *ClassStatement) String() string       { return c.Token.Literal }

func (np *NewProfessorStatement) statementNode()       {}
func (np *NewProfessorStatement) TokenLiteral() string { return np.Token.Literal }
func (np *NewProfessorStatement) String() string       { return np.Token.Literal }

func (s *StudentStatement) statementNode()       {}
func (s *StudentStatement) TokenLiteral() string { return s.Token.Literal }
func (s *StudentStatement) String() string       { return s.Token.Literal }

func (pts *ProfessorToStudentStatement) statementNode()       {}
func (pts *ProfessorToStudentStatement) TokenLiteral() string { return pts.Token.Literal }
func (pts *ProfessorToStudentStatement) String() string       { return pts.Token.Literal }

func (stp *StudentToProfessorStatement) statementNode()       {}
func (stp *StudentToProfessorStatement) TokenLiteral() string { return stp.Token.Literal }
func (stp *StudentToProfessorStatement) String() string       { return stp.Token.Literal }

func (ptmb *ProfessorTominiBufferStatement) statementNode()       {}
func (ptmb *ProfessorTominiBufferStatement) TokenLiteral() string { return ptmb.Token.Literal }
func (ptmb *ProfessorTominiBufferStatement) String() string       { return ptmb.Token.Literal }

func (mbo *MiniBufferOperationStatement) statementNode()       {}
func (mbo *MiniBufferOperationStatement) TokenLiteral() string { return mbo.Token.Literal }
func (mbo *MiniBufferOperationStatement) String() string       { return mbo.Token.Literal }

func (c *ConstantExpression) expressionNode()      {}
func (c *ConstantExpression) TokenLiteral() string { return c.Token.Literal }
func (c *ConstantExpression) String() string       { return c.Token.Literal }

func (l *LiberalExpression) expressionNode()      {}
func (l *LiberalExpression) TokenLiteral() string { return l.Token.Literal }
func (l *LiberalExpression) String() string       { return l.Token.Literal }

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String() + "\n")
	}

	return out.String()
}

func (l *LiberalStatement) statementNode()       {}
func (l *LiberalStatement) TokenLiteral() string { return l.Token.Literal }
func (l *LiberalStatement) String() string {
	var out bytes.Buffer

	out.WriteString(l.Token.Literal + "\n")
	out.WriteString(l.Mandatory.Token.Literal + l.Mandatory.String() + "\n")
	out.WriteString(l.Selection.Token.Literal + l.Selection.String() + "\n")

	return out.String()
}
