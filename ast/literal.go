package ast

import "github.com/computerphilosopher/monkey-interpreter/token"

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

type Boolean struct {
	Token token.Token
	Value bool
}

func (boolean *Boolean) expressionNode() {}
func (boolean *Boolean) TokenLiteral() string {
	return boolean.Token.Literal
}
func (boolean *Boolean) String() string {
	return boolean.Token.Literal
}
