package ast

import (
	"bytes"
	"fmt"
	"strings"
)

// FinalStatIsValid given the last statement from a statement list, this function
// traverses to the last statement checks that statement is a valid end statement,
// such as a return or exit.
func FinalStatIsValid(s StatementNode) bool {
	switch s.(type) {
	case ReturnNode:
		return true
	case ExitNode:
		return true
	case ScopeNode:
		stats := s.(ScopeNode).stats
		finalStat := stats[len(stats)-1]
		return FinalStatIsValid(finalStat)
	case IfNode:
		ifStats := s.(IfNode).ifStats
		ifFinalStat := ifStats[len(ifStats)-1]
		elseStats := s.(IfNode).elseStats
		elseFinalStat := elseStats[len(elseStats)-1]
		return FinalStatIsValid(ifFinalStat) && FinalStatIsValid(elseFinalStat)
	default:
		return false
	}
}

type ProgramNode interface {
}

// indent is a function to indent when printing the AST, given the string s, it indents
// it with all previous indents plus the new indent (sep)
func indent(s string, sep string) string {
	var buf bytes.Buffer
	for _, line := range strings.Split(s, "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%s%s\n", sep, line))
		}
	}
	return buf.String()
}

// Program the struct that encapsulates the entire program and will be the root of
// the AST.
type Program struct {
	// functions the list of all the functions in the program in the order they are
	// declared, the last function will be the "main" function.
	functions []FunctionNode
}

func NewProgram(functions []FunctionNode) Program {
	return Program{
		functions: functions,
	}
}

func (program Program) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("Program"))
	for _, f := range program.functions {
		buf.WriteString(indent(fmt.Sprintf("%s", f), "  "))
	}
	return buf.String()
}

// FunctionNode is the struct that holds information about a function, the return type,
// parameters and internal body.
type FunctionNode struct {
	pos Position

	// t is the return type of the function.
	t TypeNode

	// ident is the identifier used to reference the function.
	ident IdentifierNode

	// params is the list of parameters required to call the function.
	params []ParameterNode

	// stats is the list of statements contained within the function body.
	stats []StatementNode
}

func NewFunctionNode(pos Position, t TypeNode, ident IdentifierNode, params []ParameterNode, stats []StatementNode) FunctionNode {
	return FunctionNode{
		pos:    pos,
		t:      t,
		ident:  ident,
		params: params,
		stats:  stats,
	}
}

func (node FunctionNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s %s(", node.t, node.ident.String()[2:]))
	for i, p := range node.params {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("%s", p))
		} else {
			buf.WriteString(fmt.Sprintf(", %s", p))
		}
	}
	buf.WriteString(fmt.Sprintln(")"))
	for _, s := range node.stats {
		buf.WriteString(indent(fmt.Sprintf("%s", s), "  "))
	}
	return buf.String()
}

// ParameterNode is the struct that holds information about a parameter for a function,
// the type and identifier of the single parameter.
type ParameterNode struct {
	pos Position

	// t is the type of the parameter.
	t TypeNode

	// ident is the identifier used for the parameter.
	ident IdentifierNode
}

func NewParameterNode(pos Position, t TypeNode, ident IdentifierNode) ParameterNode {
	return ParameterNode{
		pos:   pos,
		t:     t,
		ident: ident,
	}
}

func (node ParameterNode) String() string {
	return fmt.Sprintf("%s %s", node.t, node.ident.String()[2:])
}

// Position stores the position of a node within the original code. The linenumber,
// column number and offset from the beginning of the file.
type Position struct {
	lineNumber int
	colNumber  int
	offset     int
}

// LineNumber returns the line number of a Position.
func (p Position) LineNumber() int {
	return p.lineNumber
}

// ColNumber returns the column number of a Position.
func (p Position) ColNumber() int {
	colNum := p.colNumber
	if colNum != 0 {
		colNum--
	}
	return colNum
}

func NewPosition(lineNumber int, colNumber int, offset int) Position {
	return Position{
		lineNumber: lineNumber,
		colNumber:  colNumber,
		offset:     offset,
	}
}

func (p Position) String() string {
	colNum := p.colNumber
	if colNum != 0 {
		colNum--
	}

	if DEBUG_MODE {
		offsetNum := p.offset
		if offsetNum != 0 {
			offsetNum--
		}
		return fmt.Sprintf("line %d, column %d, offset %d", p.lineNumber, colNum, offsetNum)
	}

	return fmt.Sprintf("%d:%d", p.lineNumber, colNum)
}
