package ast

import (
	"bytes"
	"fmt"
)

// GenericError is an interface that errors implement, which allows for elegent
// printing of errors.
type GenericError interface {
	String() string
	Pos() Position
}

// CustomError is a struct that stores a particular error message.
type CustomError struct {
	pos  Position
	text string
}

func NewCustomError(pos Position, text string) CustomError {
	return CustomError{
		pos:  pos,
		text: text,
	}
}

func (e CustomError) Pos() Position {
	return e.pos
}

func (e CustomError) String() string {
	return e.text
}

// TypeError is a struct for a TypeError, storing a list of acceptable TypeNodes,
// and the actual (wrong) TypeNode the semantic checker saw.
type TypeError struct {
	pos      Position
	got      TypeNode
	expected map[TypeNode]bool
}

func NewTypeError(got TypeNode, expected map[TypeNode]bool) TypeError {
	return TypeError{
		got:      got,
		expected: expected,
	}
}

func (e TypeError) Pos() Position {
	return e.pos
}

func (e TypeError) String() string {
	var b bytes.Buffer
	b.WriteString("Expected type ")
	i := 1
	for t := range e.expected {
		// If type mismatch on VOID, then trying to return from global scope
		if node, ok := t.(BaseTypeNode); ok {
			if node.t == VOID {
				return "Cannot return from global scope"
			}
		}

		if i == len(e.expected) {
			b.WriteString(fmt.Sprintf("\"%s\"", t))
		} else {
			b.WriteString(fmt.Sprintf("\"%s\" or ", t))
		}
		i++
	}

	b.WriteString(fmt.Sprintf(" but got type \"%s\"", e.got))
	return b.String()
}

func (e TypeError) addPos(pos Position) GenericError {
	if e.got == nil {
		return nil
	}
	e.pos = pos
	return e
}

// DeclarationError is a struct for a declaration error, for example, using an
// identifier before it is declared. It implements GenericError.
type DeclarationError struct {
	pos        Position
	isFunction bool
	isDefined  bool
	identifier string
}

func NewDeclarationError(pos Position, isFunction bool, isDefined bool, identifier string) DeclarationError {
	return DeclarationError{
		pos:        pos,
		isFunction: isFunction,
		isDefined:  isDefined,
		identifier: identifier,
	}
}

func (e DeclarationError) Pos() Position {
	return e.pos
}

func (e DeclarationError) String() string {
	var b bytes.Buffer
	if e.isFunction {
		if e.isDefined {
			b.WriteString(fmt.Sprintf("Function \"%s\" is already defined", e.identifier))
		} else {
			b.WriteString(fmt.Sprintf("Function \"%s\" is not defined", e.identifier))
		}
	} else {
		if e.isDefined {
			b.WriteString(fmt.Sprintf("Variable \"%s\" is already defined in the current scope", e.identifier))
		} else {
			b.WriteString(fmt.Sprintf("Variable \"%s\" is not defined in the current scope", e.identifier))
		}
	}
	return b.String()
}
