package parser

import (
	"fmt"

	"github.com/computerphilosopher/monkey-interpreter/ast"
	"github.com/computerphilosopher/monkey-interpreter/lexer"
	"github.com/computerphilosopher/monkey-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []error
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []error{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}
	if !p.expectPeek(token.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.Assign) {
		return nil
	}

	p.getNextTokenUntilMeet(token.Semicolon)

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	p.getNextTokenUntilMeet(token.Semicolon)

	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type != t {
		p.peekError(t)
		return false
	}
	p.nextToken()
	return true
}

func (p *Parser) peekError(t token.TokenType) {
	err := fmt.Errorf("expected next token to be %s, got %s instead",
		token.TokenTypeLiteral[t], token.TokenTypeLiteral[p.peekToken.Type])
	p.errors = append(p.errors, err)
}

func (p *Parser) getNextTokenUntilMeet(t token.TokenType) {
	for p.curToken.Type != t {
		p.nextToken()
	}
}