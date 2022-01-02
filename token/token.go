package token

type TokenType int

const (
	Illegal = iota
	EOF
	Ident
	Int
	Assign
	Plus
	Comma
	Semicolon
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	Function
	Let
)

type Token struct {
	Type    TokenType
	Literal string
}
