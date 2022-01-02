package lexer_test

import (
	"testing"

	"github.com/computerphilosopher/monkey/lexer"
	"github.com/computerphilosopher/monkey/token"
	"github.com/stretchr/testify/assert"
)

func TestSingleToken(t *testing.T) {

	input := "=+(){},;"
	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
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

	lexer := lexer.NewLexer(input)
	for _, expected := range tests {
		tok := lexer.NextToken()
		assert.Equal(t, expected.Type, tok.Type)
		assert.Equal(t, expected.Literal, tok.Literal)
	}
}
