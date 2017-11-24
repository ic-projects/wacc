package ast

import (
	"bytes"
	"fmt"
	"strings"
	"utils"
)

var DEBUG_MODE bool

type ProgramNode interface {
}

// Program the struct that encapsulates the entire program and will be the root of
// the AST.
type Program struct {
	// Functions the list of all the Functions in the program in the order they are
	// declared, the last function will be the "main" function.
	Functions []FunctionNode
}

func NewProgram(functions []FunctionNode) Program {
	return Program{
		Functions: functions,
	}
}

func (program Program) String() string {
	var tempbuf bytes.Buffer
	tempbuf.WriteString(fmt.Sprintln("Program"))
	for _, f := range program.Functions {
		tempbuf.WriteString(utils.Indent(fmt.Sprintf("%s", f), "  "))
	}
	var buf bytes.Buffer
	for i, line := range strings.Split(tempbuf.String(), "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%d\t%s\n", i, line))
		}
	}
	return buf.String()
}

// FunctionNode is the struct that holds information about a function, the return type,
// parameters and internal body.
type FunctionNode struct {
	Pos Position

	// T is the return type of the function.
	T TypeNode

	// Ident is the identifier used to reference the function.
	Ident IdentifierNode

	// Params is the list of parameters required to call the function.
	Params []ParameterNode

	// Stats is the list of statements contained within the function body.
	Stats []StatementNode
}

func NewFunctionNode(pos Position, t TypeNode, ident IdentifierNode, params []ParameterNode, stats []StatementNode) FunctionNode {
	return FunctionNode{
		Pos:    pos,
		T:      t,
		Ident:  ident,
		Params: params,
		Stats:  stats,
	}
}

func (node FunctionNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s %s(", node.T, node.Ident.String()[2:]))
	for i, p := range node.Params {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("%s", p))
		} else {
			buf.WriteString(fmt.Sprintf(", %s", p))
		}
	}
	buf.WriteString(fmt.Sprintln(")"))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "  "))
	}
	return buf.String()
}

// ParameterNode is the struct that holds information about a parameter for a function,
// the type and identifier of the single parameter.
type ParameterNode struct {
	Pos Position

	// T is the type of the parameter.
	T TypeNode

	// Ident is the identifier used for the parameter.
	Ident IdentifierNode
}

func NewParameterNode(pos Position, t TypeNode, ident IdentifierNode) ParameterNode {
	return ParameterNode{
		Pos:   pos,
		T:     t,
		Ident: ident,
	}
}

func (node ParameterNode) String() string {
	return fmt.Sprintf("%s %s", node.T, node.Ident.String()[2:])
}
