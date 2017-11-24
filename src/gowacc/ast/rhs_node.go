package ast

import (
	"bytes"
	"fmt"
)

// RHSNode is an empty interface for Lhs nodes to implement.
type RHSNode interface {
}

// ExpressionNode - defined in expression_node.go

// ArrayLiteralNode stores the position and elements of an array literal.
//
// E.g.
//
//  [2, 4]
type ArrayLiteralNode struct {
	Pos   Position
	Exprs []ExpressionNode
}

func NewArrayLiteralNode(pos Position, exprs []ExpressionNode) ArrayLiteralNode {
	return ArrayLiteralNode{
		Pos:   pos,
		Exprs: exprs,
	}
}

func (node ArrayLiteralNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- ARRAY LITERAL"))
	for _, e := range node.Exprs {
		buf.WriteString(indent(fmt.Sprintf("%s\n", e), "  "))
	}
	return buf.String()
}

// NewPairNode stores the position and elements of a newpair call.
//
// E.g.
//  newpair(4, 2)
type NewPairNode struct {
	Pos Position
	Fst ExpressionNode
	Snd ExpressionNode
}

func NewNewPairNode(pos Position, fst ExpressionNode, snd ExpressionNode) NewPairNode {
	return NewPairNode{
		Pos: pos,
		Fst: fst,
		Snd: snd,
	}
}

func (node NewPairNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- NEW_PAIR"))
	buf.WriteString(indent(fmt.Sprintln("- FST"), "  "))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.Fst), "    "))
	buf.WriteString(indent(fmt.Sprintln("- SND"), "  "))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.Snd), "    "))
	return buf.String()
}

// PairFirstElementNode - defined in lhs_node.go

// PairSecondElementNode - defined in lhs_node.go

// FunctionCallNode stores the position, identifier and passed in parameters for
// a function call.
//
// E.g.
//  call f(true, false)
type FunctionCallNode struct {
	Pos   Position
	Ident IdentifierNode
	Exprs []ExpressionNode
}

func NewFunctionCallNode(pos Position, ident IdentifierNode, exprs []ExpressionNode) FunctionCallNode {
	return FunctionCallNode{
		Pos:   pos,
		Ident: ident,
		Exprs: exprs,
	}
}

func (node FunctionCallNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\n", node.Ident))
	for _, e := range node.Exprs {
		buf.WriteString(fmt.Sprintf("%s\n", e))
	}
	return buf.String()
}
