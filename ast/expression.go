package ast

import (
	"bytes"
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

type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (exp *InfixExpression) expressionNode() {}

func (exp *InfixExpression) TokenLiteral() string {
	return exp.Token.Literal
}

func (exp *InfixExpression) String() string {
	var out strings.Builder
	out.WriteString("(")
	out.WriteString(exp.Left.String())
	out.WriteString(" " + exp.Operator + " ")
	out.WriteString(exp.Right.String())
	out.WriteString(")")
	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (exp *IfExpression) expressionNode() {
}

func (exp *IfExpression) TokenLiteral() string {
	return exp.Token.Literal
}

func (exp IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(exp.Condition.String())
	out.WriteString(" ")
	out.WriteString(exp.Consequence.String())

	if exp.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(exp.Alternative.String())
	}

	return out.String()
}
