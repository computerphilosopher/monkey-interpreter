package parser

import (
	"testing"

	"github.com/computerphilosopher/monkey-interpreter/ast"
	"github.com/computerphilosopher/monkey-interpreter/lexer"
	"github.com/computerphilosopher/monkey-interpreter/token"
	"github.com/stretchr/testify/assert"
)

func TestLetStatement(t *testing.T) {
	assert := assert.New(t)

	input := "let x = 5;\n" +
		"let y = 10;\n" +
		"let foobar = 838383;\n"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))

	assert.NotNil(program)

	assert.Equal(3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		assert.Equal("let", stmt.TokenLiteral())

		letStmt, ok := stmt.(*ast.LetStatement)
		assert.True(ok)

		assert.Equal(tt.expectedIdentifier, letStmt.Name.Value)
		assert.Equal(tt.expectedIdentifier, letStmt.Name.TokenLiteral())

	}
}

func TestReturnStatement(t *testing.T) {
	assert := assert.New(t)

	input := "return 5;\n" +
		"return 10;\n" +
		"return 993322;\n"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(3, len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.True(ok)
		assert.Equal("return", returnStmt.TokenLiteral())
	}
}

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{
					Type:    token.Let,
					Literal: "let",
				},
				Name: &ast.Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	assert.Equal(t, "let myVar = anotherVar;", program.String())
}

func TestIdentifierExpression(t *testing.T) {
	assert := assert.New(t)
	input := "foobar;"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	ident, ok := stmt.Expression.(*ast.Identifier)
	assert.True(ok)
	assert.Equal("foobar", ident.Value)
	assert.Equal("foobar", ident.TokenLiteral())
}

func TestIntegerLiteralExpression(t *testing.T) {
	assert := assert.New(t)
	input := "5;"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.True(ok)
	assert.Equal(int64(5), literal.Value)
	assert.Equal("5", literal.TokenLiteral())

}

func testIntegerLiteral(t *testing.T, literal ast.Expression, value int64) {
	integer, ok := literal.(*ast.IntegerLiteral)
	assert.True(t, ok)
	assert.Equal(t, value, integer.Value)
}

func TestParsingPrefixExpression(t *testing.T) {
	assert := assert.New(t)
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", int64(5)},
		{"-15;", "-", int64(15)},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := New(l)

		program := p.ParseProgram()
		assert.Equal(0, len(p.Errors()))
		assert.Equal(1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(ok)

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.True(ok)
		assert.Equal(tt.operator, exp.Operator)
		testIntegerLiteral(t, exp.Right, tt.integerValue)
	}
}
