package lexer_test

import (
	"testing"

	"github.com/computerphilosopher/monkey/lexer"
	"github.com/computerphilosopher/monkey/token"
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
