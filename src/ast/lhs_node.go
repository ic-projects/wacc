package ast

import (
	"bytes"
	"fmt"
	"utils"
)

// LHSNode is an empty interface for Lhs nodes to implement.
type LHSNode interface {
}

/******************** IDENTIFIER NODE ********************/

// IdentifierNode is a struct that stores the position and string of an
// identifier.
type IdentifierNode struct {
	Pos   Position
	Ident string
}

func NewIdentifierNode(pos Position, ident string) *IdentifierNode {
	return &IdentifierNode{
		Pos:   pos,
		Ident: ident,
	}
}

func (node IdentifierNode) String() string {
	if node.Ident == "" {
		return "- main"
	}
	return fmt.Sprintf("- %s", node.Ident)
}

/******************** PAIR FIRST ELEMENT NODE ********************/

// PairFirstElementNode is a struct that stores the position and expression of
// an access to a pair's first element.
//
// E.g.
//
//  Fst i
type PairFirstElementNode struct {
	Pos     Position
	Expr    ExpressionNode
	Pointer bool
}

func (fst *PairFirstElementNode) SetPointer(p bool) {
	fst.Pointer = p
}

func NewPairFirstElementNode(
	pos Position,
	expr ExpressionNode,
) *PairFirstElementNode {
	return &PairFirstElementNode{
		Pos:     pos,
		Expr:    expr,
		Pointer: false,
	}
}

func (node PairFirstElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- FST"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.Expr), "  "))
	return buf.String()
}

/******************** PAIR SECOND ELEMENT NODE ********************/

// PairSecondElementNode is a struct that stores the position and expression of
// an access to a pair's second element.
//
// E.g.
//
//  Snd i
type PairSecondElementNode struct {
	Pos     Position
	Expr    ExpressionNode
	Pointer bool
}

func (snd *PairSecondElementNode) SetPointer(p bool) {
	snd.Pointer = p
}

func NewPairSecondElementNode(
	pos Position,
	expr ExpressionNode,
) *PairSecondElementNode {
	return &PairSecondElementNode{
		Pos:     pos,
		Expr:    expr,
		Pointer: false,
	}
}

func (node PairSecondElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- SND"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.Expr), "  "))
	return buf.String()
}

/******************** ARRAY ELEMENT NODE ********************/

// ArrayElementNode is a struct that stores the position, identifier and
// expressions of an access to an array.
//
// E.g.
//  i[4][3+2]
type ArrayElementNode struct {
	Pos     Position
	Ident   *IdentifierNode
	Exprs   []ExpressionNode
	Pointer bool
}

func (arr *ArrayElementNode) SetPointer(p bool) {
	arr.Pointer = p
}

func NewArrayElementNode(
	pos Position,
	ident *IdentifierNode,
	exprs []ExpressionNode,
) *ArrayElementNode {
	return &ArrayElementNode{
		Pos:     pos,
		Ident:   ident,
		Exprs:   exprs,
		Pointer: false,
	}
}

func (node ArrayElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\n", node.Ident))
	for _, e := range node.Exprs {
		buf.WriteString(fmt.Sprintln("  - []"))
		buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", e), "    "))
	}
	return buf.String()
}
