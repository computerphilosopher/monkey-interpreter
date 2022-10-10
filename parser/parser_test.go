package parser

import (
	"testing"

	"github.com/computerphilosopher/monkey-interpreter/ast"
	"github.com/computerphilosopher/monkey-interpreter/lexer"
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
