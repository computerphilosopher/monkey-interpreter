package main

import (
	"testing"

	"github.com/computerphilosopher/monkey/token"
)

func TestNextToken(t *testing.T) {

	input := "=+(){},;"
	tests := []struct {
		Type token.TokenType
		Literal
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
	for _, expected := tests {
		tok := lexer.NextToken()
		assert.Eqaul(t, tok.Type, expected.Type)
		assert.Eqaul(t, tok.Literal, expected.Literal)
	}

}
