package lexer_test

import (
	"testing"

	"github.com/computerphilosopher/monkey-interpreter/lexer"
	"github.com/computerphilosopher/monkey-interpreter/token"
	"github.com/stretchr/testify/assert"
)

func Helper(t *testing.T, input string, expectedTokens []token.Token) {

	lexer := lexer.NewLexer(input)
	for _, expected := range expectedTokens {
		tok := lexer.NextToken()
		assert.Equal(t, expected.Type, tok.Type)
		assert.Equal(t, expected.Literal, tok.Literal)
	}
}

func TestSingleToken(t *testing.T) {

	input := "=+(){},;"
	expected := []token.Token{
		{token.Assign, "="},
		{token.Plus, "+"},
		{token.LeftParen, "("},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.RightBrace, "}"},
		{token.Comma, ","},
		{token.Semicolon, ";"},
		{token.EOF, ""},
	}

	Helper(t, input, expected)
}

func TestLetStatement(t *testing.T) {
	varDeclare := "let five = 5;\n" +
		"let ten = 10;"

	expected := []token.Token{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Ident, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
	}

	Helper(t, varDeclare, expected)

	/*
		funcDeclare := "let add = fn(x, y) {\n" +
			"x + y;\n" +
			"};"
	*/

}
