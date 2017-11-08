package ast

import (
	"bytes"
	"fmt"
)

// RHSNode is an empty interface for lhs nodes to implement.
type RHSNode interface {
}

// ExpressionNode - defined in expression_node.go

// ArrayLiteralNode stores the position and elements of an array literal.
//
// E.g.
//
//  [2, 4]
type ArrayLiteralNode struct {
	pos   Position
	exprs []ExpressionNode
}

func NewArrayLiteralNode(pos Position, exprs []ExpressionNode) ArrayLiteralNode {
	return ArrayLiteralNode{
		pos:   pos,
		exprs: exprs,
	}
}

func (node ArrayLiteralNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- ARRAY LITERAL"))
	for _, e := range node.exprs {
		buf.WriteString(indent(fmt.Sprintf("%s\n", e), "  "))
	}
	return buf.String()
}

// NewPairNode stores the position and elements of a newpair call.
//
// E.g.
//  newpair(4, 2)
type NewPairNode struct {
	pos Position
	fst ExpressionNode
	snd ExpressionNode
}

func NewNewPairNode(pos Position, fst ExpressionNode, snd ExpressionNode) NewPairNode {
	return NewPairNode{
		pos: pos,
		fst: fst,
		snd: snd,
	}
}

func (node NewPairNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- NEW_PAIR"))
	buf.WriteString(indent(fmt.Sprintln("- FST"), "  "))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.fst), "    "))
	buf.WriteString(indent(fmt.Sprintln("- SND"), "  "))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.snd), "    "))
	return buf.String()
}

// PairFirstElementNode - defined in lhs_node.go

// PairSecondElementNode - defined in lhs_node.go

// FunctionCallNode stores the position, identifier and passed in parameters for
// a function call.
//
// e.g.
//
//  call f(true, false)
type FunctionCallNode struct {
	pos   Position
	ident IdentifierNode
	exprs []ExpressionNode
}

func NewFunctionCallNode(pos Position, ident IdentifierNode, exprs []ExpressionNode) FunctionCallNode {
	return FunctionCallNode{
		pos:   pos,
		ident: ident,
		exprs: exprs,
	}
}

func (node FunctionCallNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\n", node.ident))
	for _, e := range node.exprs {
		buf.WriteString(fmt.Sprintf("%s\n", e))
	}
	return buf.String()
}
