package main

import (
	"os"

	"github.com/computerphilosopher/monkey-interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
