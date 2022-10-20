package parser

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
