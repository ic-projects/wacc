package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// GenericError is an interface that errors implement, which allows for elegent
// printing of errors.
type GenericError interface {
	String() string
	Pos() Position
}

// ErrorDeclaration is an interface that extends a GenericError to also give
// a declaration position
type ErrorDeclaration interface {
	GenericError
	PosDeclared() Position
}

/**************** CUSTOM ERROR ****************/

// CustomError is a struct that stores a particular error message.
type CustomError struct {
	pos  Position
	text string
}

// NewCustomError builds a CustomError.
func NewCustomError(pos Position, text string) CustomError {
	return CustomError{
		pos:  pos,
		text: text,
	}
}

// Pos returns the position of this error.
func (e CustomError) Pos() Position {
	return e.pos
}

func (e CustomError) String() string {
	return e.text
}

/**************** TYPE ERROR ****************/

// TypeError is a struct for a TypeError, storing a list of acceptable
// TypeNodes, and the actual (wrong) TypeNode the semantic checker saw.
type TypeError struct {
	pos      Position
	got      TypeNode
	expected []TypeNode
}

// NewTypeError builds a TypeError
func NewTypeError(got TypeNode, expected []TypeNode) TypeError {
	return TypeError{
		got:      got,
		expected: expected,
	}
}

// Pos returns the position of this error.
func (e TypeError) Pos() Position {
	return e.pos
}

func (e TypeError) String() string {
	var b bytes.Buffer
	b.WriteString("Expected type ")
	i := 1
	for _, t := range e.expected {
		// If type mismatch on VOID, then trying to return from global Scope
		if node, ok := t.(*BaseTypeNode); ok {
			if node.T == VOID {
				return "Cannot return from global Scope"
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

/**************** TYPE ERROR DECLARATION ****************/

// TypeErrorDeclaration is a struct that stores a TypeError and where an
// identifier was declared, for more useful error messages.
type TypeErrorDeclaration struct {
	typeError   TypeError
	posDeclared Position
}

// NewTypeErrorDeclaration builds a TypeErrorDeclaration
func NewTypeErrorDeclaration(err TypeError, pos Position) TypeErrorDeclaration {
	return TypeErrorDeclaration{
		typeError:   err,
		posDeclared: pos,
	}
}

// Pos returns the position of this error.
func (e TypeErrorDeclaration) Pos() Position {
	return e.typeError.pos
}

// PosDeclared returns the position where this variable was first declared.
func (e TypeErrorDeclaration) PosDeclared() Position {
	return e.posDeclared
}

func (e TypeErrorDeclaration) String() string {
	return e.typeError.String()
}

/**************** DECLARATION ERROR ****************/

// DeclarationError is a struct for a declaration error, for example, using an
// identifier before it is declared. It implements GenericError.
type DeclarationError struct {
	pos        Position
	isFunction bool
	isDefined  bool
	identifier string
}

// NewDeclarationError builds a DeclarationError
func NewDeclarationError(
	pos Position,
	isFunction bool,
	isDefined bool,
	identifier string,
) DeclarationError {
	return DeclarationError{
		pos:        pos,
		isFunction: isFunction,
		isDefined:  isDefined,
		identifier: identifier,
	}
}

// Pos returns the position of this error.
func (e DeclarationError) Pos() Position {
	return e.pos
}

func (e DeclarationError) String() string {
	var b bytes.Buffer
	if e.isFunction {
		if e.isDefined {
			b.WriteString(fmt.Sprintf(
				"Function \"%s\" is already defined",
				e.identifier,
			))
		} else {
			b.WriteString(fmt.Sprintf(
				"Function \"%s\" is not defined",
				e.identifier,
			))
		}
	} else {
		if e.isDefined {
			b.WriteString(fmt.Sprintf(
				"Variable \"%s\" is already defined in the current Scope",
				e.identifier,
			))
		} else {
			b.WriteString(fmt.Sprintf(
				"Variable \"%s\" is not defined in the current Scope",
				e.identifier,
			))
		}
	}
	return b.String()
}

/**************** PREVIOUSLY DECLARED ****************/

// PreviouslyDeclared is a struct that extends a DeclarationError with a
// position of where the variable was first declared.
type PreviouslyDeclared struct {
	declarationError DeclarationError
	posDeclared      Position
}

// NewPreviouslyDeclared builds a PreviouslyDeclared
func NewPreviouslyDeclared(
	declarationError DeclarationError,
	posDeclared Position,
) PreviouslyDeclared {
	return PreviouslyDeclared{
		declarationError: declarationError,
		posDeclared:      posDeclared,
	}
}

func (e PreviouslyDeclared) String() string {
	return e.declarationError.String()
}

// Pos returns the position of this error.
func (e PreviouslyDeclared) Pos() Position {
	return e.declarationError.pos
}

// PosDeclared returns the position where this variable was first declared.
func (e PreviouslyDeclared) PosDeclared() Position {
	return e.posDeclared
}

/**************** ERROR HELPER FUNCTIONS ****************/

// getLine returns a requested line from the WACC source file.
func getLine(path string, n int) string {
	// Open the WACC source file
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: Unable to open the specified WACC source file")
		os.Exit(100)
	}

	reader := bufio.NewReader(f)
	var line string

	for i := 0; i < n; i++ {
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: Unable to read specified line")
			os.Exit(100)
		}
	}

	return line
}

// PrintErrors pretty prints an error.
func PrintErrors(errors []GenericError, filepath string) {
	maxErrors := 4
	for i, e := range errors {
		if i >= maxErrors {
			fmt.Printf("\nAnd %d other error(s)", len(errors)-maxErrors)
			break
		}

		var b bytes.Buffer
		b.WriteString("\nSemantic Error at ")
		b.WriteString(fmt.Sprintf("%s\n", e.Pos()))

		// Remove leading spaces and tabs
		line := getLine(filepath, e.Pos().LineNumber())
		leadingChars := 0
		for _, c := range line {
			if c == '\t' || c == ' ' {
				leadingChars++
			} else {
				break
			}
		}
		b.WriteString(line[leadingChars:])
		b.WriteString(strings.Repeat(" ", e.Pos().ColNumber()-leadingChars))
		b.WriteString("^\n")
		b.WriteString(fmt.Sprintln(e))
		if typeDeclarationError, ok := e.(ErrorDeclaration); ok {
			b.WriteString("Declared at ")
			b.WriteString(fmt.Sprintf(
				"%s\n",
				typeDeclarationError.PosDeclared(),
			))

			// Remove leading spaces and tabs
			line := getLine(
				filepath,
				typeDeclarationError.PosDeclared().LineNumber(),
			)
			leadingChars := 0
			for _, c := range line {
				if c == '\t' || c == ' ' {
					leadingChars++
				} else {
					break
				}
			}
			b.WriteString(line[leadingChars:])
			b.WriteString(strings.Repeat(
				" ",
				typeDeclarationError.PosDeclared().ColNumber()-leadingChars),
			)
			b.WriteString("^")
		}
		fmt.Println(b.String())
	}
}
