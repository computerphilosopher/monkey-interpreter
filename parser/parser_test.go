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
