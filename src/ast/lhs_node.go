package ast

import (
	"bytes"
	"fmt"
	"utils"
)

// LHSNode is an interface for LHS nodes to implement.
type LHSNode interface {
	ProgramNode
}

/**************** IDENTIFIER NODE ****************/

// IdentifierNode is a struct that stores the position and string of an
// identifier.
type IdentifierNode struct {
	Pos   utils.Position
	Ident string
}

// NewIdentifierNode builds an IdentifierNode
func NewIdentifierNode(pos utils.Position, ident string) *IdentifierNode {
	return &IdentifierNode{
		Pos:   pos,
		Ident: ident,
	}
}

func (node *IdentifierNode) String() string {
	if node.Ident == "" {
		return "- main"
	}
	return fmt.Sprintf("- %s", node.Ident)
}

func (node *IdentifierNode) walkNode(visitor Visitor) {
}

/**************** PAIR FIRST ELEMENT NODE ****************/

// PairFirstElementNode is a struct that stores the position and expression of
// an access to a pair's first element.
//
// E.g.
//
//  Fst i
type PairFirstElementNode struct {
	Pos     utils.Position
	Expr    ExpressionNode
	Pointer bool
}

// SetPointer sets whether this is accessed using a pointer.
func (node *PairFirstElementNode) SetPointer(p bool) {
	node.Pointer = p
}

// NewPairFirstElementNode builds a PairFirstElementNode
func NewPairFirstElementNode(
	pos utils.Position,
	expr ExpressionNode,
) *PairFirstElementNode {
	return &PairFirstElementNode{
		Pos:     pos,
		Expr:    expr,
		Pointer: false,
	}
}

func (node *PairFirstElementNode) String() string {
	return writeSimpleString("FST", node.Expr)
}

func (node *PairFirstElementNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *PairFirstElementNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
}

/**************** PAIR SECOND ELEMENT NODE ****************/

// PairSecondElementNode is a struct that stores the position and expression of
// an access to a pair's second element.
//
// E.g.
//
//  Snd i
type PairSecondElementNode struct {
	Pos     utils.Position
	Expr    ExpressionNode
	Pointer bool
}

// SetPointer sets whether this is accessed using a pointer.
func (node *PairSecondElementNode) SetPointer(p bool) {
	node.Pointer = p
}

// NewPairSecondElementNode builds a PairSecondElementNode
func NewPairSecondElementNode(
	pos utils.Position,
	expr ExpressionNode,
) *PairSecondElementNode {
	return &PairSecondElementNode{
		Pos:     pos,
		Expr:    expr,
		Pointer: false,
	}
}

func (node *PairSecondElementNode) String() string {
	return writeSimpleString("SND", node.Expr)
}

func (node *PairSecondElementNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *PairSecondElementNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
}

/**************** ARRAY ELEMENT NODE ****************/

// ArrayElementNode is a struct that stores the position, identifier and
// expressions of an access to an array.
//
// E.g.
//  i[4][3+2]
type ArrayElementNode struct {
	Pos     utils.Position
	Ident   *IdentifierNode
	Exprs   []ExpressionNode
	Pointer bool
}

// SetPointer sets whether this is accessed using a pointer.
func (node *ArrayElementNode) SetPointer(p bool) {
	node.Pointer = p
}

// NewArrayElementNode builds an ArrayElementNode
func NewArrayElementNode(
	pos utils.Position,
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

func (node *ArrayElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\n", node.Ident))
	for _, e := range node.Exprs {
		buf.WriteString(fmt.Sprintln("  - []"))
		buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", e), "    "))
	}
	return buf.String()
}

func (node *ArrayElementNode) MapExpressions(m Mapper) {
	for i, e := range node.Exprs {
		node.Exprs[i] = m(e)
	}
}

func (node *ArrayElementNode) walkNode(visitor Visitor) {
	for _, e := range node.Exprs {
		Walk(visitor, e)
	}
}

/**************** STRUCT ELEMENT NODE ****************/

// StructElementNode is a struct that stores the position, identifier and
// member of an access to a struct.
//
// E.g.
//  s.lhs
type StructElementNode struct {
	Pos        utils.Position
	Struct     *IdentifierNode
	Ident      *IdentifierNode
	StructType *StructInternalNode
	Pointer    bool
}

// SetStructType sets the type of a struct element.
func (node *StructElementNode) SetStructType(p *StructInternalNode) {
	node.StructType = p
}

// SetPointer sets whether this is accessed using a pointer.
func (node *StructElementNode) SetPointer(p bool) {
	node.Pointer = p
}

// NewStructElementNode builds an StructElementNode
func NewStructElementNode(
	pos utils.Position,
	struc *IdentifierNode,
	ident *IdentifierNode,
) *StructElementNode {
	return &StructElementNode{
		Pos:     pos,
		Struct:  struc,
		Ident:   ident,
		Pointer: false,
	}
}

func (node *StructElementNode) String() string {
	return fmt.Sprintf(
		"structelem %s.%s (pointer: %t)\n",
		node.Struct,
		node.Ident.String()[2:],
		node.Pointer,
	)
}

func (node *StructElementNode) walkNode(visitor Visitor) {
}

/**************** POINTER NODE ****************/

// PointerNode is a struct that stores the identifier of a pointer.
type PointerNode struct {
	Pos   utils.Position
	Ident *IdentifierNode
}

// NewPointerNode builds a PointerNode.
func NewPointerNode(pos utils.Position, ident *IdentifierNode) *PointerNode {
	return &PointerNode{
		Pos:   pos,
		Ident: ident,
	}
}

func (node *PointerNode) String() string {
	return fmt.Sprintf(
		"POINTER %s\n",
		node.Ident.String()[2:],
	)
}

func (node *PointerNode) walkNode(visitor Visitor) {
}
