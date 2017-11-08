package ast

import (
	"bytes"
	"fmt"
)

// LHSNode is an empty interface for lhs nodes to implement.
type LHSNode interface {
}

// IdentifierNode is a struct that stores the position and string of an identifier.
type IdentifierNode struct {
	pos   Position
	ident string
}

func NewIdentifierNode(pos Position, ident string) IdentifierNode {
	return IdentifierNode{
		pos:   pos,
		ident: ident,
	}
}

func (node IdentifierNode) String() string {
	if node.ident == "" {
		return "- main"
	}
	return fmt.Sprintf("- %s", node.ident)
}

// PairFirstElementNode is a struct that stores the position and expression of
// an access to a pair's first element.
//
// E.g.
//
//  fst i
type PairFirstElementNode struct {
	pos  Position
	expr ExpressionNode
}

func NewPairFirstElementNode(pos Position, expr ExpressionNode) PairFirstElementNode {
	return PairFirstElementNode{
		pos:  pos,
		expr: expr,
	}
}

func (node PairFirstElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- FST"))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.expr), "  "))
	return buf.String()
}

// PairSecondElementNode is a struct that stores the position and expression of
// an access to a pair's second element.
//
// E.g.
//  snd i
type PairSecondElementNode struct {
	pos  Position
	expr ExpressionNode
}

func NewPairSecondElementNode(pos Position, expr ExpressionNode) PairSecondElementNode {
	return PairSecondElementNode{
		pos:  pos,
		expr: expr,
	}
}

func (node PairSecondElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- SND"))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.expr), "  "))
	return buf.String()
}

// ArrayElementNode is a struct that stores the position, identifier and expressions of
// an access to an array.
//
// E.g.
//
//  i[4][3+2]
type ArrayElementNode struct {
	pos   Position
	ident IdentifierNode
	exprs []ExpressionNode
}

func NewArrayElementNode(pos Position, ident IdentifierNode, exprs []ExpressionNode) ArrayElementNode {
	return ArrayElementNode{
		pos:   pos,
		ident: ident,
		exprs: exprs,
	}
}

func (node ArrayElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\n", node.ident))
	for _, e := range node.exprs {
		buf.WriteString(fmt.Sprintln("  - []"))
		buf.WriteString(indent(fmt.Sprintf("%s\n", e), "    "))
	}
	return buf.String()
}
