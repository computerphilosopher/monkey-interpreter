package parser

import (
	"testing"

	"github.com/computerphilosopher/monkey-interpreter/ast"
	"github.com/computerphilosopher/monkey-interpreter/lexer"
	"github.com/computerphilosopher/monkey-interpreter/token"
	"github.com/stretchr/testify/assert"
)

func TestLetStatement(t *testing.T) {
	assert := assert.New(t)

	input := "let x = 5;\n" +
		"let y = 10;\n" +
		"let foobar = 838383;\n"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))

	assert.NotNil(program)

	assert.Equal(3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		assert.Equal("let", stmt.TokenLiteral())

		letStmt, ok := stmt.(*ast.LetStatement)
		assert.True(ok)

		assert.Equal(tt.expectedIdentifier, letStmt.Name.Value)
		assert.Equal(tt.expectedIdentifier, letStmt.Name.TokenLiteral())

	}
}

func TestReturnStatement(t *testing.T) {
	assert := assert.New(t)

	input := "return 5;\n" +
		"return 10;\n" +
		"return 993322;\n"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(3, len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.True(ok)
		assert.Equal("return", returnStmt.TokenLiteral())
	}
}

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{
					Type:    token.Let,
					Literal: "let",
				},
				Name: &ast.Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{
						Type:    token.Ident,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	assert.Equal(t, "let myVar = anotherVar;", program.String())
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) {
	assert := assert.New(t)
	ident, ok := exp.(*ast.Identifier)
	assert.True(ok)
	assert.Equal(value, ident.Value)
	assert.Equal(value, ident.TokenLiteral())
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) {

	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, exp, int64(v))
		return
	case int64:
		testIntegerLiteral(t, exp, v)
		return
	case string:
		testIdentifier(t, exp, v)
		return
	case bool:
		testBooleanLiteral(t, exp, v)
		return
	}
	assert.False(t, true)
}

func TestIdentifierExpression(t *testing.T) {
	assert := assert.New(t)
	input := "foobar;"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)
	testIdentifier(t, stmt.Expression, "foobar")
}

func TestIntegerLiteralExpression(t *testing.T) {
	assert := assert.New(t)
	input := "5;"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.True(ok)
	testLiteralExpression(t, literal, 5)
}

func testIntegerLiteral(t *testing.T, literal ast.Expression, value int64) {
	integer, ok := literal.(*ast.IntegerLiteral)
	assert.True(t, ok)
	assert.Equal(t, value, integer.Value)
}

func TestBooleanExpression(t *testing.T) {
	assert := assert.New(t)
	input := "true;"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	literal, ok := stmt.Expression.(*ast.Boolean)
	assert.True(ok)
	testLiteralExpression(t, literal, true)
}

func testBooleanLiteral(t *testing.T, boolean ast.Expression, value bool) {
	literal, ok := boolean.(*ast.Boolean)
	assert.True(t, ok)
	assert.Equal(t, value, literal.Value)
}

func testInfixExpression(t *testing.T, exp ast.Expression,
	left interface{}, operator string, right interface{}) {

	opExp, ok := exp.(*ast.InfixExpression)
	assert.True(t, ok)

	testLiteralExpression(t, opExp.Left, left)
	assert.Equal(t, operator, opExp.Operator)
}

func TestParsingPrefixExpression(t *testing.T) {
	assert := assert.New(t)
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", int64(5)},
		{"-15;", "-", int64(15)},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := New(l)

		program := p.ParseProgram()
		assert.Equal(0, len(p.Errors()))
		assert.Equal(1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(ok)

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.True(ok)
		assert.Equal(tt.operator, exp.Operator)
		testIntegerLiteral(t, exp.Right, tt.integerValue)
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	assert := assert.New(t)
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assert.Equal(0, len(p.Errors()))
		assert.Equal(1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(ok)
		testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}
	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)

		program := p.ParseProgram()
		assert.Equal(0, len(p.Errors()))

		actual := program.String()
		assert.Equal(tt.expected, actual)
	}
}
