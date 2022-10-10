package token

type TokenType int

const (
	Illegal = iota
	EOF
	Ident
	Int
	True
	False
	Bang
	Assign
	Equal
	NotEqual
	Plus
	Minus
	Star
	Slash
	LessThan
	GreaterThan
	Comma
	Semicolon
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	Function
	Let
	Return
	If
	Else
)

type Token struct {
	Type    TokenType
	Literal string
}

var SingleToken map[rune]TokenType = map[rune]TokenType{
	'!':    Bang,
	'=':    Assign,
	'+':    Plus,
	'-':    Minus,
	'*':    Star,
	'/':    Slash,
	'<':    LessThan,
	'>':    GreaterThan,
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
		"let":    Let,
		"fn":     Function,
		"return": Return,
		"if":     If,
		"else":   Else,
		"true":   True,
		"false":  False,
	}

	tokenType, isKeyword := keywords[ident]

	if isKeyword {
		return tokenType
	}

	return Ident
}

var TokenTypeLiteral = map[TokenType]string{
	Illegal:     "Illegal",
	EOF:         "EOF",
	Ident:       "Ident",
	Int:         "Int",
	True:        "True",
	False:       "False",
	Bang:        "Bang",
	Assign:      "Assign",
	Equal:       "Equal",
	NotEqual:    "NotEqual",
	Plus:        "Plus",
	Minus:       "Minus",
	Star:        "Star",
	Slash:       "Slash",
	LessThan:    "LessThan",
	GreaterThan: "GreaterThan",
	Comma:       "Comma",
	Semicolon:   "Semicolon",
	LeftParen:   "LeftParen",
	RightParen:  "RightParen",
	LeftBrace:   "LeftBrace",
	RightBrace:  "RightBrace",
	Function:    "Function",
	Let:         "Let",
	Return:      "Return",
	If:          "If",
	Else:        "Else",
}
