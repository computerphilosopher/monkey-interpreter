package token_test

import (
	"testing"

	"github.com/computerphilosopher/monkey-interpreter/token"
	"github.com/stretchr/testify/assert"
)

func TestTokenType(t *testing.T) {
	tokenType, err := token.GetTokenType("=")
	assert.Equal(t, err)
	assert.Equal(t, token.Assign, tokenType)
}
