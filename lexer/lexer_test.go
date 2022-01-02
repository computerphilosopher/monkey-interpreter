package main

import (
	"testing"

	"github.com/computerphilosopher/monkey/token"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {

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

	lexer := NewLexer(input)
	for _, expected := range tests {
		tok := lexer.NextToken()
		assert.Equal(t, tok.Type, expected.Type)
		assert.Equal(t, tok.Literal, expected.Literal)
	}

}
