package object

import (
	"fmt"
	"strings"

	"github.com/computerphilosopher/monkey-interpreter/ast"
)

type ObjectType string

const (
	IntegerObject     = "Integer"
	BooleanObject     = "Boolean"
	NullObject        = "Null"
	ReturnValueObject = "ReturnValue"
	ErrorObject       = "ErrorObject"
	FunctionObject    = "Function"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (integer *Integer) Type() ObjectType {
	return IntegerObject
}

func (integer *Integer) Inspect() string {
	return fmt.Sprintf("%d", integer.Value)
}

type Boolean struct {
	Value bool
}

func (boolean *Boolean) Type() ObjectType {
	return BooleanObject
}

func (boolean *Boolean) Inspect() string {
	return fmt.Sprintf("%t", boolean.Value)
}

type Null struct {
}

func (null *Null) Type() ObjectType {
	return NullObject
}

func (null *Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return ReturnValueObject
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ErrorObject
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FunctionObject
}

func (f *Function) Inspect() string {
	out := strings.Builder{}
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n")

	return out.String()
}
