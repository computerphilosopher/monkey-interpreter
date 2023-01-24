package ast

import (
	"bytes"
	"strings"

	"github.com/computerphilosopher/monkey-interpreter/token"
)

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

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (boolean *BooleanLiteral) expressionNode() {}
func (boolean *BooleanLiteral) TokenLiteral() string {
	return boolean.Token.Literal
}
func (boolean *BooleanLiteral) String() string {
	return boolean.Token.Literal
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (function *FunctionLiteral) expressionNode() {}
func (function *FunctionLiteral) TokenLiteral() string {
	return function.Token.Literal
}
func (function *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range function.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(function.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(function.Body.String())

	return out.String()
}
