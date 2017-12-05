package main

import (
	"bytes"
	"fmt"
)

// RHSNode is an empty interface for Lhs nodes to implement.
type RHSNode interface {
	fmt.Stringer
}

/**************** EXPRESSION NODE ****************/

// ExpressionNode - defined in expression_node.go

/**************** ARRAY LITERAL NODE ****************/

// ArrayLiteralNode stores the position and elements of an array literal.
//
// E.g.
//
//  [2, 4]
type ArrayLiteralNode struct {
	Pos   Position
	Exprs []ExpressionNode
}

// NewArrayLiteralNode builds an ArrayLiteralNode.
func NewArrayLiteralNode(
	pos Position,
	exprs []ExpressionNode,
) *ArrayLiteralNode {
	return &ArrayLiteralNode{
		Pos:   pos,
		Exprs: exprs,
	}
}

func (node ArrayLiteralNode) String() string {
	return writeExpressionsString("ARRAY LITERAL", node.Exprs)
}

/**************** NEW PAIR NODE ****************/

// NewPairNode stores the position and elements of a newpair call.
//
// E.g.
//  newpair(4, 2)
type NewPairNode struct {
	Pos Position
	Fst ExpressionNode
	Snd ExpressionNode
}

// NewNewPairNode builds a NewPairNode.
func NewNewPairNode(
	pos Position,
	fst ExpressionNode,
	snd ExpressionNode,
) *NewPairNode {
	return &NewPairNode{
		Pos: pos,
		Fst: fst,
		Snd: snd,
	}
}

func (node NewPairNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- NEW_PAIR"))
	buf.WriteString(Indent(fmt.Sprintln("- FST"), "  "))
	buf.WriteString(Indent(fmt.Sprintf("%s\n", node.Fst), "    "))
	buf.WriteString(Indent(fmt.Sprintln("- SND"), "  "))
	buf.WriteString(Indent(fmt.Sprintf("%s\n", node.Snd), "    "))
	return buf.String()
}

/**************** PAIR FIRST ELEMENT NODE ****************/

// PairFirstElementNode - defined in lhs_node.go

/**************** PAIR SECOND ELEMENT NODE ****************/

// PairSecondElementNode - defined in lhs_node.go

/**************** FUNCTION CALL NODE ****************/

// FunctionCallNode stores the position, identifier and passed in parameters for
// a function call.
//
// E.g.
//  call f(true, false)
type FunctionCallNode struct {
	Pos   Position
	Ident *IdentifierNode
	Exprs []ExpressionNode
}

// NewFunctionCallNode builds a FunctionCallNode.
func NewFunctionCallNode(
	pos Position,
	ident *IdentifierNode,
	exprs []ExpressionNode,
) *FunctionCallNode {
	return &FunctionCallNode{
		Pos:   pos,
		Ident: ident,
		Exprs: exprs,
	}
}

func (node FunctionCallNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln(node.Ident))
	for _, e := range node.Exprs {
		buf.WriteString(fmt.Sprintln(e))
	}
	return buf.String()
}

/**************** STRUCT NEW NODE ****************/

// StructNewNode stores the position, type, and members of an initialised
// struct.
type StructNewNode struct {
	Pos        Position
	T          *StructTypeNode
	Exprs      []ExpressionNode
	structNode *StructNode
}

// SetStructType replaces a StructNewNode's StructNode.
func (node *StructNewNode) SetStructType(t *StructNode) {
	node.structNode = t
}

// NewStructNewNode builds a StructNewNode.
func NewStructNewNode(
	pos Position,
	ident *IdentifierNode,
	exprs []ExpressionNode,
) *StructNewNode {
	return &StructNewNode{
		Pos:   pos,
		T:     NewStructTypeNode(ident),
		Exprs: exprs,
	}
}

func (node StructNewNode) String() string {
	return writeExpressionsString(fmt.Sprintf("NEW %s\n", node.T), node.Exprs)
}
