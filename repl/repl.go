package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/computerphilosopher/monkey-interpreter/lexer"
	"github.com/computerphilosopher/monkey-interpreter/parser"
	"github.com/computerphilosopher/monkey-interpreter/token"

	log "github.com/sirupsen/logrus"
)

const Prompt = ">> "

func scan(scanner *bufio.Scanner) error {
	if scanned := scanner.Scan(); scanned {
		return nil
	}
	return scanner.Err()
}

func lex(line string, out io.Writer) {
	l := lexer.NewLexer(line)

	for {
		t := l.NextToken()
		if t.Type == token.EOF {
			break
		}
		fmt.Fprintf(out, "%+v\n", t)
	}
}

func Start(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Fprintf(writer, Prompt)
		err := scan(scanner)
		if err != nil {
			log.Error(err)
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(writer, p.Errors())
			continue
		}

		io.WriteString(writer, program.String())
		io.WriteString(writer, "\n")
	}
}

func printParserErrors(out io.Writer, errors []error) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg.Error()+"\n")
	}
}
