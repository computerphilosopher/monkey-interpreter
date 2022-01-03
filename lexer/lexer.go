package lexer

import (
	"unicode"

	"github.com/computerphilosopher/monkey-interpreter/token"
)

type Lexer struct {
	input        []rune
	ch           rune
	position     int
	readPosition int
}

func NewLexer(input string) *Lexer {
	ret := &Lexer{
		input: []rune(input),
	}
	ret.stepForward()

	return ret
}

func (lexer *Lexer) stepForward() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = '\000'
		return
	}

	lexer.ch = lexer.input[lexer.readPosition]
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func runeToString(ch rune) string {
	if ch == '\000' {
		return ""
	}
	return string(ch)
}

func isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		ch == '_'
}

func (lexer *Lexer) readIdent() token.Token {
	if !isLetter(lexer.ch) && !unicode.IsDigit(lexer.ch) {
		return token.Token{
			Type: token.Illegal,
		}
	}

	keepCondition := func(ch rune) bool {
		if isLetter(ch) {
			return isLetter(ch)
		}
		return unicode.IsDigit(ch)
	}

	begin := lexer.position
	for keepCondition(lexer.ch) {
		lexer.stepForward()
	}

	end := lexer.position
	ident := string(lexer.input[begin:end])

	return token.Token{
		Type:    token.GetIdentType(ident),
		Literal: ident,
	}
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' {
		lexer.stepForward()
	}
}

func (lexer *Lexer) NextToken() token.Token {

	lexer.skipWhitespace()
	ret := func() token.Token {
		if tokenType, isSingletoken :=
			token.SingleToken[lexer.ch]; isSingletoken {
			return token.Token{
				Type:    tokenType,
				Literal: runeToString(lexer.ch),
			}
		}
		return lexer.readIdent()
	}()

	lexer.stepForward()
	return ret
}
