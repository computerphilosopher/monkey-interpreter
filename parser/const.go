package parser

import "github.com/computerphilosopher/monkey-interpreter/token"

const (
	_ int = iota
	Lowest
	Equals      //== or !=
	LessGreater // < or >
	Sum         // - or +
	Product     // * or /
	Prefix      // - or +
	Call        // myFunction(x)
)

func noPrefixParseFnError(t token.TokenType) {
}
