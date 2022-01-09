package lexer

import (
	"unicode"

	"github.com/computerphilosopher/monkey-interpreter/token"
	"github.com/sirupsen/logrus"
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

	lexer.position = lexer.readPosition

	if lexer.readPosition >= len(lexer.input) {
		logrus.Debug("input: ", string(lexer.input))
		logrus.Debug("last char: ", string(lexer.ch))
		logrus.Debug("postion: ", lexer.position,
			" readPoistion: ", lexer.readPosition)
		lexer.ch = '\000'
		return
	}

	lexer.ch = lexer.input[lexer.readPosition]
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

func tokenTypeFunc(start rune) func(string) token.TokenType {
	if isLetter(start) {
		return token.GetIdentType
	}
	if unicode.IsDigit(start) {
		return func(_ string) token.TokenType {
			return token.Int
		}
	}
	return func(_ string) token.TokenType {
		return token.Illegal
	}
}

func keepGoingFunc(start rune) func(rune) bool {
	if isLetter(start) {
		return isLetter
	}
	return unicode.IsDigit
}

func (lexer *Lexer) readStringToken(keepGoing func(rune) bool,
	getTokenType func(string) token.TokenType) token.Token {
	if !keepGoing(lexer.ch) {
		return token.Token{
			Type: token.Illegal,
		}
	}

	begin := lexer.position
	for keepGoing(lexer.ch) {
		lexer.stepForward()
	}
	end := lexer.position

	literal := string(lexer.input[begin:end])

	lexer.position -= 1
	lexer.readPosition -= 1

	return token.Token{
		Type:    getTokenType(literal),
		Literal: literal,
	}
}

func (lexer *Lexer) skipWhitespace() {
	isWhiteSpace := func(ch rune) bool {
		return ch == ' ' || ch == '\n' || ch == '\t'
	}
	for isWhiteSpace(lexer.ch) {
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
		return lexer.readStringToken(keepGoingFunc(lexer.ch),
			tokenTypeFunc(lexer.ch))
	}()

	lexer.stepForward()
	return ret
}
