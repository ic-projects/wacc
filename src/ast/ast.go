package ast

import (
	"bytes"
	"fmt"
	"strings"
	"utils"
)

// DebugMode is a setting for printing extra debugging information when true.
var DebugMode bool

// ProgramNode is an interface for AST nodes to implement.
type ProgramNode interface {
	fmt.Stringer
	walkNode(Visitor)
}

/**************** PRINTING ****************/

func writeSimpleString(name string, nodes ...fmt.Stringer) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s\n", name))
	for _, n := range nodes {
		buf.WriteString(utils.Indent(n.String(), "  "))
	}
	return buf.String()
}

func writeExpressionsString(name string, exprs []ExpressionNode) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s\n", name))
	for _, e := range exprs {
		buf.WriteString(utils.Indent(e.String(), "  "))
	}
	return buf.String()
}

/**************** PROGRAM ****************/

// Program the struct that encapsulates the entire program and will be the root
// of the AST.
type Program struct {
	// Functions the list of all the Functions in the program in the order they
	// are declared, the last function will be the "main" function.
	Structs   []*StructNode
	Functions []*FunctionNode
}

// NewProgram builds a Program.
func NewProgram(structs []*StructNode, functions []*FunctionNode) *Program {
	return &Program{
		Structs:   structs,
		Functions: functions,
	}
}

func (program *Program) String() string {
	var tempbuf bytes.Buffer
	tempbuf.WriteString(fmt.Sprintln("Program"))
	for _, f := range program.Structs {
		tempbuf.WriteString(utils.Indent(f.String(), "  "))
	}
	for _, f := range program.Functions {
		tempbuf.WriteString(utils.Indent(f.String(), "  "))
	}
	var buf bytes.Buffer
	for i, line := range strings.Split(tempbuf.String(), "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%d\t%s\n", i, line))
		}
	}
	return buf.String()
}

func (program *Program) walkNode(visitor Visitor) {
	for _, f := range program.Structs {
		Walk(visitor, f)
	}
	for _, f := range program.Functions {
		Walk(visitor, f)
	}
}

/**************** STRUCT NODE ****************/

// StructNode holds information about the name, members and size of a WACC
// struct.
type StructNode struct {
	Pos        utils.Position
	Ident      *IdentifierNode
	Types      []*StructInternalNode
	TypesMap   map[string]int
	MemorySize int
}

// NewStructNode builds a StructNode.
func NewStructNode(
	pos utils.Position,
	ident *IdentifierNode,
	types []*StructInternalNode,
) *StructNode {
	structNode := StructNode{
		Pos:      pos,
		Ident:    ident,
		Types:    types,
		TypesMap: make(map[string]int),
	}
	mem := 0
	for i, t := range structNode.Types {
		t.MemoryOffset = mem
		mem += SizeOf(t.T)
		structNode.TypesMap[t.Ident.Ident] = i
	}
	structNode.MemorySize = mem
	return &structNode
}

func (node *StructNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(
		"- STRUCT %s (size: %d)\n",
		node.Ident.String()[2:],
		node.MemorySize,
	))
	for _, p := range node.Types {
		buf.WriteString(fmt.Sprintf("%s\n", p))
	}
	return buf.String()
}

func (node *StructNode) walkNode(visitor Visitor) {
	for _, f := range node.Types {
		Walk(visitor, f)
	}
}

/**************** STRUCT INTERNAL NODE ****************/

// StructInternalNode holds information about each member of a struct.
type StructInternalNode struct {
	Pos          utils.Position
	Ident        *IdentifierNode
	T            TypeNode
	MemoryOffset int
}

// NewStructInternalNode builds a StructInternalNode.
func NewStructInternalNode(
	pos utils.Position,
	ident *IdentifierNode,
	t TypeNode,
) *StructInternalNode {
	return &StructInternalNode{
		Pos:   pos,
		Ident: ident,
		T:     t,
	}
}

func (node *StructInternalNode) String() string {
	return fmt.Sprintf(
		"  %s %s (offset: %d)",
		node.Ident,
		node.T,
		node.MemoryOffset,
	)
}

func (node *StructInternalNode) walkNode(visitor Visitor) {
}

/**************** FUNCTION NODE ****************/

// FunctionNode is the struct that holds information about a function, the
// return type, parameters and internal body.
type FunctionNode struct {
	Pos utils.Position

	// T is the return type of the function.
	T TypeNode

	// Ident is the identifier used to reference the function.
	Ident *IdentifierNode

	// Params is the list of parameters required to call the function.
	Params Parameters

	// Stats is the list of statements contained within the function body.
	Stats Statements
}

// NewFunctionNode builds a FunctionNode.
func NewFunctionNode(
	pos utils.Position, t TypeNode,
	ident *IdentifierNode,
	params []*ParameterNode,
	stats []StatementNode,
) *FunctionNode {
	return &FunctionNode{
		Pos:    pos,
		T:      t,
		Ident:  ident,
		Params: params,
		Stats:  stats,
	}
}

func (node *FunctionNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(
		"- %s %s%s",
		node.T,
		node.Ident.String()[2:],
		node.Params.String(),
	))
	buf.WriteString(utils.Indent(node.Stats.String(), "  "))
	return buf.String()
}

func (node *FunctionNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Params)
	Walk(visitor, node.Stats)
}

/**************** PARAMETER NODE ****************/

// ParameterNode is the struct that holds information about a parameter for a
// function, the type and identifier of the single parameter.
type ParameterNode struct {
	Pos utils.Position

	// T is the type of the parameter.
	T TypeNode

	// Ident is the identifier used for the parameter.
	Ident *IdentifierNode
}

// NewParameterNode builds a ParameterNode.
func NewParameterNode(
	pos utils.Position,
	t TypeNode,
	ident *IdentifierNode,
) *ParameterNode {
	return &ParameterNode{
		Pos:   pos,
		T:     t,
		Ident: ident,
	}
}

func (node *ParameterNode) String() string {
	return fmt.Sprintf("%s %s", node.T, node.Ident.String()[2:])
}

func (node *ParameterNode) walkNode(visitor Visitor) {
}

/**************** PARAMETER NODE SLICE ****************/

// Parameters is a slice of ParameterNodes. We need this type so that visitors
// can easily tell when to change scope.
type Parameters []*ParameterNode

func (params Parameters) String() string {
	var buf bytes.Buffer
	buf.WriteString("(")
	for i, p := range params {
		if i == 0 {
			buf.WriteString(p.String())
		} else {
			buf.WriteString(fmt.Sprintf(", %s", p))
		}
	}
	buf.WriteString(fmt.Sprintln(")"))
	return buf.String()
}

func (params Parameters) walkNode(visitor Visitor) {
	for _, p := range params {
		Walk(visitor, p)
	}
}
