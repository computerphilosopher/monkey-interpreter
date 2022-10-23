package parser

import "github.com/computerphilosopher/monkey-interpreter/token"

const (
	_ int = iota
	Lowest
	Equals      //==
	LessGreater // < or >
	Sum         // +
	Product     // *
	Prefix      // - or +
	Call        // myFunction(x)
)

func noPrefixParseFnError(t token.TokenType) {
}
