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

	input := "=!!=+==-*/(){},;"
	expected := []token.Token{
		{token.Assign, "="},
		{token.Bang, "!"},
		{token.NotEqual, "!="},
		{token.Plus, "+"},
		{token.Equal, "=="},
		{token.Minus, "-"},
		{token.Star, "*"},
		{token.Slash, "/"},
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
		"return x + y;\n" +
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
		{token.Return, "return"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
	}

	Helper(t, funcDeclare, expected)

	conditionalFunc := "if (5 < 10) {\n" +
		"return true\n" +
		"} else if (10 > 5){\n" +
		"return true\n}" +
		"else { return false\n" +
		"}"

	expected = []token.Token{
		{token.If, "if"},
		{token.LeftParen, "("},
		{token.Int, "5"},
		{token.LessThan, "<"},
		{token.Int, "10"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.RightBrace, "}"},
		{token.Else, "else"},
		{token.If, "if"},
		{token.LeftParen, "("},
		{token.Int, "10"},
		{token.GreaterThan, ">"},
		{token.Int, "5"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.RightBrace, "}"},
		{token.Else, "else"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.RightBrace, "}"},
	}

	Helper(t, conditionalFunc, expected)
}
