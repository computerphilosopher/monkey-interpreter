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
	varDeclare := "let five = 5;"

	expected := []token.Token{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
	}

	Helper(t, varDeclare, expected)

	funcDeclare := "let add = fn(x, y) {\n" +
		"x + y;\n" +
		"};"

	expected = []token.Token{
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LeftParen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
	}

	Helper(t, funcDeclare, expected)
}
