package parser

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/computerphilosopher/monkey-interpreter/ast"
	"github.com/computerphilosopher/monkey-interpreter/lexer"
	"github.com/computerphilosopher/monkey-interpreter/token"
	"github.com/stretchr/testify/assert"
)

func TestLetStatement(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assert.Equal(0, len(p.Errors()))
		assert.Equal(1, len(program.Statements))

		stmt := program.Statements[0]
		assert.Equal("let", stmt.TokenLiteral())

		letStmt, ok := stmt.(*ast.LetStatement)
		assert.True(ok)

		assert.Equal(tt.expectedIdentifier, letStmt.Name.Value)

		val := stmt.(*ast.LetStatement).Value
		testLiteralExpression(t, val, tt.expectedValue)

	}
}

func TestReturnStatement(t *testing.T) {
	assert := assert.New(t)

	test := []struct {
		input         string
		expectedValue int
	}{
		{"return 5;", 5},
		{"return 10;", 10},
		{"return 993322;", 993322},
	}

	for _, tt := range test {
		l := lexer.NewLexer(tt.input)
		p := New(l)

		program := p.ParseProgram()
		assert.Equal(0, len(p.Errors()))
		assert.Equal(1, len(program.Statements))

		returnStmt, ok := program.Statements[0].(*ast.ReturnStatement)
		assert.True(ok)
		assert.Equal("return", returnStmt.TokenLiteral())
		assert.Equal(strconv.Itoa(tt.expectedValue), returnStmt.ReturnValue.TokenLiteral())

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

	literal, ok := stmt.Expression.(*ast.BooleanLiteral)
	assert.True(ok)
	testLiteralExpression(t, literal, true)
}

func testBooleanLiteral(t *testing.T, boolean ast.Expression, value bool) {
	literal, ok := boolean.(*ast.BooleanLiteral)
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
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
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

func TestIfExpression(t *testing.T) {
	assert := assert.New(t)
	input := `if (x < y) { x }`

	l := lexer.NewLexer(input)
	p := New(l)
	program := p.ParseProgram()

	assert.Equal(0, len(p.Errors()))
	fmt.Printf("%+v\n", p.Errors())
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	exp, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(ok)
	testInfixExpression(t, exp.Condition, "x", "<", "y")

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	testIdentifier(t, consequence.Expression, "x")

	assert.Nil(exp.Alternative)

}

func TestIfElseExpression(t *testing.T) {
	assert := assert.New(t)
	input := `if (x < y) { x } else { y }`

	l := lexer.NewLexer(input)
	p := New(l)
	program := p.ParseProgram()

	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	exp, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(ok)
	testInfixExpression(t, exp.Condition, "x", "<", "y")

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)
	testIdentifier(t, consequence.Expression, "x")

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)
	testIdentifier(t, alternative.Expression, "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{
			input:          "fn() {};",
			expectedParams: []string{},
		},
		{
			input:          "fn(x) {};",
			expectedParams: []string{"x"},
		},
		{
			input:          "fn(x, y, z) {};",
			expectedParams: []string{"x", "y", "z"},
		},
	}

	assert := assert.New(t)

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := New(l)
		program := p.ParseProgram()
		assert.Equal(0, len(p.Errors()))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(ok)
		function := stmt.Expression.(*ast.FunctionLiteral)

		assert.Equal(len(tt.expectedParams), len(function.Parameters))

	}
}
func TestFunctionLiteralParsing(t *testing.T) {
	assert := assert.New(t)
	input := "fn(x, y) { x + y; }"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	assert.True(ok)

	assert.Equal(2, len(function.Parameters))

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	assert.Equal(1, len(function.Body.Statements))

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestCallExpressionParsing(t *testing.T) {
	assert := assert.New(t)
	input := "add(1, 2 * 3, 4 + 5);"

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	assert.Equal(0, len(p.Errors()))
	assert.Equal(1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(ok)

	exp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(ok)

	testIdentifier(t, exp.Function, "add")

	assert.Equal(3, len(exp.Arguments))

	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}
