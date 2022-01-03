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

var SingleToken map[rune]TokenType = map[rune]TokenType{
	'=':    Assign,
	'+':    Plus,
	'(':    LeftParen,
	')':    RightParen,
	'{':    LeftBrace,
	'}':    RightBrace,
	',':    Comma,
	';':    Semicolon,
	'\x00': EOF,
}

func GetIdentType(ident string) TokenType {
	keywords := map[string]TokenType{
		"let": Let,
		"fn":  Function,
	}

	tokenType, isKeyword := keywords[ident]

	if isKeyword {
		return tokenType
	}

	return Ident
}
