package ast

import (
	"bytes"
	"fmt"
	"strings"
)

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

func indent(s string, sep string) string {
	var buf bytes.Buffer
	for _, line := range strings.Split(s, "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%s%s\n", sep, line))
		}
	}
	return buf.String()
}

type Program struct {
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

type FunctionNode struct {
	pos    Position
	t      TypeNode
	ident  IdentifierNode
	params []ParameterNode
	stats  []StatementNode
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

type ParameterNode struct {
	pos   Position
	t     TypeNode
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

type Position struct {
	lineNumber int
	colNumber  int
	offset     int
}

func (p Position) LineNumber() int {
	return p.lineNumber
}

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
