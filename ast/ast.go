package ast

import "flink_monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST containing all statments
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Identifier
//
// letter = "a" | "b" | ... | "z" | "A" | ... | "Z";
//
// digit = "0" | "1" | "2" | ... | "9";
//
// identifier = letter , { letter | digit | "_" } ;
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func NewIdentifier(token token.Token, value string) Identifier {
	return Identifier{token, value}
}

func (id *Identifier) expressionNode() {
}

func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

// VarStatement ::= var <identifier> = <expression>
type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {
}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}

// ReturnStatement ::= return <expression>;
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
