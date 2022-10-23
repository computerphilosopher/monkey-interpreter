package ast

import (
	"strings"

	"github.com/computerphilosopher/monkey-interpreter/token"
)

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (literal *IntegerLiteral) expressionNode() {}
func (literal *IntegerLiteral) TokenLiteral() string {
	return literal.Token.Literal
}
func (literal *IntegerLiteral) String() string {
	return literal.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (exp *PrefixExpression) expressionNode() {}

func (exp *PrefixExpression) TokenLiteral() string {
	return exp.Token.Literal
}

func (exp *PrefixExpression) String() string {
	var out strings.Builder
	out.WriteString("(")
	out.WriteString(exp.Operator)
	out.WriteString(exp.Right.String())
	out.WriteString(")")
	return out.String()
}
