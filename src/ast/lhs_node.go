package ast

import (
  "bytes"
  "fmt"
)

type LHSNode interface {
}

type LHSNodeStruct struct {
}

func NewLHSNode() LHSNode {
	return LHSNodeStruct{}
}

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
