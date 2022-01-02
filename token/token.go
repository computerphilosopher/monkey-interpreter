package main

type Token int

const (
	Illegal = iota
	Eof
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
