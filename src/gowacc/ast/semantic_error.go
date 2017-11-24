package ast

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
			if node.T == VOID {
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

type ErrorDeclaration interface {
	String() string
	Pos() Position
	PosDeclared() Position
}

// TypeErrorDeclaration is a struct that stores a TypeError and where an identifier
// was declared, for more useful error messages.
type TypeErrorDeclaration struct {
	typeError   TypeError
	posDeclared Position
}

func NewTypeErrorDeclaration(err TypeError, pos Position) TypeErrorDeclaration {
	return TypeErrorDeclaration{
		typeError:   err,
		posDeclared: pos,
	}
}

func (e TypeErrorDeclaration) Pos() Position {
	return e.typeError.pos
}

func (e TypeErrorDeclaration) PosDeclared() Position {
	return e.posDeclared
}

func (e TypeErrorDeclaration) String() string {
	return e.typeError.String()
}

func (e TypeErrorDeclaration) addPos(pos Position) GenericError {
	return e.addPos(pos)
}

// DeclarationError is a struct for a declaration error, for example, using an
// identifier before it is declared. It implements GenericError.
type DeclarationError struct {
	pos         Position
	isFunction  bool
	isDefined   bool
	identifier  string
	posDeclared Position
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

type PreviouslyDelcared struct {
	declarationError DeclarationError
	posDeclared      Position
}

func NewPreviouslyDeclared(declarationError DeclarationError, posDeclared Position) PreviouslyDelcared {
	return PreviouslyDelcared{
		declarationError: declarationError,
		posDeclared:      posDeclared,
	}
}

func (e PreviouslyDelcared) String() string {
	return e.declarationError.String()
}

func (e PreviouslyDelcared) Pos() Position {
	return e.declarationError.pos
}

func (e PreviouslyDelcared) PosDeclared() Position {
	return e.posDeclared
}

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
			b.WriteString(fmt.Sprintf("%s\n", typeDeclarationError.PosDeclared()))

			// Remove leading spaces and tabs
			line := getLine(filepath, typeDeclarationError.PosDeclared().LineNumber())
			leadingChars := 0
			for _, c := range line {
				if c == '\t' || c == ' ' {
					leadingChars++
				} else {
					break
				}
			}
			b.WriteString(line[leadingChars:])
			b.WriteString(strings.Repeat(" ", typeDeclarationError.PosDeclared().ColNumber()-leadingChars))
			b.WriteString("^")
		}
		fmt.Println(b.String())
	}
}
