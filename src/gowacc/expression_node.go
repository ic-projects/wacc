package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// ExpressionNode is an empty interface for expression nodes to implement.
type ExpressionNode interface {
	fmt.Stringer
}

/**************** EXPRESSION HELPER FUNCTIONS ****************/

// Type returns the correct TypeNode for a given ExpressionNode
func Type(e ExpressionNode, s *SymbolTable) TypeNode {
	switch node := e.(type) {
	case *BinaryOperatorNode:
		switch node.Op {
		case MUL, DIV, MOD, ADD, SUB:
			return NewBaseTypeNode(INT)
		case GT, GEQ, LT, LEQ, EQ, NEQ, AND, OR:
			return NewBaseTypeNode(BOOL)
		}
	case *UnaryOperatorNode:
		switch node.Op {
		case NOT:
			return NewBaseTypeNode(BOOL)
		case NEG, LEN, ORD:
			return NewBaseTypeNode(INT)
		case CHR:
			return NewBaseTypeNode(CHAR)
		}
	case *NullNode:
		return NewNullTypeNode()
	case *PairTypeNode:
		return NewBaseTypeNode(PAIR)
	case *IntegerLiteralNode:
		return NewBaseTypeNode(INT)
	case *BooleanLiteralNode:
		return NewBaseTypeNode(BOOL)
	case *CharacterLiteralNode:
		return NewBaseTypeNode(CHAR)
	case *StringLiteralNode:
		return NewStringArrayTypeNode()
	case *ArrayElementNode:
		a, _ := s.SearchForIdent(node.Ident.Ident)
		arr := a.T.(*ArrayTypeNode)
		dimLeft := arr.Dim - len(node.Exprs)
		if dimLeft == 0 {
			return arr.T
		}
		return NewArrayTypeNode(arr.T, dimLeft)
	case *StructElementNode:
		return node.stuctType.T
	case *IdentifierNode:
		v, _ := s.SearchForIdent(node.Ident)
		return v.T
	}
	return nil
}

// BuildBinOpTree is a function that builds the correct tree of binary operation
// when given the first expression, a list of the remaining binary operators and
// expressions and the position inside the source file (used for error
// messages).
//
// The list of remaining binary operators and expressions is given in the form
// [[space, BinaryOperator, space, Expression], ...]
// where space is ignored.
func BuildBinOpTree(
	first ExpressionNode,
	list []interface{},
	position Position,
) ExpressionNode {
	if len(list) > 1 {
		// Generate the LHS expression node
		var toParse []interface{}
		for i := 0; i < len(list)-1; i++ {
			toParse = append(toParse, list[i])
		}
		lhs := BuildBinOpTree(first, toParse, position)

		// Get the RHS node
		// Note that last is in the form [space, BinaryOperator, space,
		// Expression],So we use last[1] to get the BinaryOperator and last[3]
		// to get the Expression
		last := list[len(list)-1].([]interface{})
		return NewBinaryOperatorNode(
			position,
			last[1].(BinaryOperator),
			lhs,
			last[3].(ExpressionNode),
		)
	}
	return NewBinaryOperatorNode(
		position,
		list[0].([]interface{})[1].(BinaryOperator),
		first,
		list[0].([]interface{})[3].(ExpressionNode),
	)
}

/**************** UNARY OPERATOR ****************/

// UnaryOperator is an enum which defines the different unary operators.
type UnaryOperator int

const (
	// NOT Not (!)
	NOT UnaryOperator = iota
	// NEG Negate (-)
	NEG
	// LEN Length (len)
	LEN
	// ORD Ordinate (ord)
	ORD
	// CHR Character (chr)
	CHR
)

func (unOp UnaryOperator) String() string {
	switch unOp {
	case NOT:
		return "- !"
	case NEG:
		return "- -"
	case LEN:
		return "- len"
	case ORD:
		return "- ord"
	case CHR:
		return "- chr"
	}
	return "ERROR"
}

/**************** BINARY OPERATOR ****************/

// BinaryOperator is an enum which defines the different binary operators.
type BinaryOperator int

const (
	// MUL Multiply (*)
	MUL BinaryOperator = iota
	// DIV Divide (/)
	DIV
	// MOD Modulus (%)
	MOD
	// ADD Add (+)
	ADD
	// SUB Subtract (-)
	SUB
	// GT Greater than (>)
	GT
	// GEQ Greater than or equal to (>=)
	GEQ
	// LT Less than (<)
	LT
	// LEQ Less than or equal to (<=)
	LEQ
	// EQ Equal (==)
	EQ
	// NEQ Not equal (!=)
	NEQ
	// AND And (&&)
	AND
	// OR Or (||)
	OR
)

func (binOp BinaryOperator) String() string {
	switch binOp {
	case OR:
		return "- ||"
	case AND:
		return "- &&"
	case MUL:
		return "- *"
	case DIV:
		return "- /"
	case MOD:
		return "- %"
	case SUB:
		return "- -"
	case ADD:
		return "- +"
	case GEQ:
		return "- >="
	case GT:
		return "- >"
	case LEQ:
		return "- <="
	case LT:
		return "- <"
	case EQ:
		return "- =="
	case NEQ:
		return "- !="
	}
	return "ERROR"
}

/**************** INTEGER LITERAL NODE ****************/

// IntegerLiteralNode is a struct which stores the position and value of an
// integer literal.
//
// E.g.
//  7
type IntegerLiteralNode struct {
	Pos Position
	Val int
}

// NewIntegerLiteralNode builds an IntegerLiteralNode
func NewIntegerLiteralNode(pos Position, val int) *IntegerLiteralNode {
	return &IntegerLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node IntegerLiteralNode) String() string {
	return fmt.Sprintf("- %d", node.Val)
}

/**************** BOOLEAN LITERAL NODE ****************/

// BooleanLiteralNode is a struct which stores the position and value of a
// boolean literal.
//
// E.g.
//  false
type BooleanLiteralNode struct {
	Pos Position
	Val bool
}

// NewBooleanLiteralNode builds an BooleanLiteralNode
func NewBooleanLiteralNode(pos Position, val bool) *BooleanLiteralNode {
	return &BooleanLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node BooleanLiteralNode) String() string {
	return fmt.Sprintf("- %s", strconv.FormatBool(node.Val))
}

/**************** CHARACTER LITERAL NODE ****************/

// CharacterLiteralNode is a struct which stores the position and value of a
// character literal.
//
// E.g.
//  'c'
type CharacterLiteralNode struct {
	Pos Position
	Val rune
}

// NewCharacterLiteralNode builds a CharacterLiteralNode
func NewCharacterLiteralNode(pos Position, val rune) *CharacterLiteralNode {
	return &CharacterLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node CharacterLiteralNode) String() string {
	if node.Val == '\000' {
		return "- '\\0'"
	}
	if node.Val == '"' {
		return "- '\\\"'"
	}
	return fmt.Sprintf("- %q", node.Val)
}

/**************** STRING LITERAL NODE ****************/

// StringLiteralNode is a struct which stores the position and value of a string
// literal.
//
// E.g.
//  "Hello World!"
type StringLiteralNode struct {
	Pos Position
	Val string
}

// NewStringLiteralNode builds a StringLiteralNode
func NewStringLiteralNode(pos Position, val string) *StringLiteralNode {
	return &StringLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node StringLiteralNode) String() string {
	return fmt.Sprintf("- \"%s\"", node.Val)
}

/**************** NULL NODE ****************/

// NullNode is a struct which stores the position of a pair literal.
// This does not store the value of the pair literal since the value of a pair
// literal is always null.
type NullNode struct {
	Pos Position
}

// NewNullNode builds a NullNode
func NewNullNode(pos Position) *NullNode {
	return &NullNode{
		Pos: pos,
	}
}

func (node NullNode) String() string {
	return "- null\n"
}

/**************** IDENTIFIER NODE ****************/

// IdentifierNode - defined in lhs_node.go

/**************** ARRAY ELEMENT NODE ****************/

// ArrayElementNode - defined in lhs_node.go

/**************** UNARY OPERATOR NODE ****************/

// UnaryOperatorNode is a struct which stores the position, (unary) operator and
// expression of a unary operator operation on an expression.
//
// E.g.
//
//  !true
type UnaryOperatorNode struct {
	Pos  Position
	Op   UnaryOperator
	Expr ExpressionNode
}

// NewUnaryOperatorNode builds a UnaryOperatorNode
func NewUnaryOperatorNode(
	pos Position,
	op UnaryOperator,
	expr ExpressionNode,
) *UnaryOperatorNode {
	return &UnaryOperatorNode{
		Pos:  pos,
		Op:   op,
		Expr: expr,
	}
}

func (node UnaryOperatorNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s\n", node.Op))
	buf.WriteString(Indent(fmt.Sprintf("%s\n", node.Expr), "  "))

	return buf.String()
}

/**************** BINARY OPERATOR NODE ****************/

// BinaryOperatorNode is a struct which stores the position, (binary) operator
// and the two expressions of a binary operation on two expressions.
//
// E.g.
//
//  5 + 2
type BinaryOperatorNode struct {
	Pos   Position
	Op    BinaryOperator
	Expr1 ExpressionNode
	Expr2 ExpressionNode
}

// NewBinaryOperatorNode builds a BinaryOperatorNode
func NewBinaryOperatorNode(
	pos Position,
	op BinaryOperator,
	expr1 ExpressionNode,
	expr2 ExpressionNode,
) *BinaryOperatorNode {
	return &BinaryOperatorNode{
		Pos:   pos,
		Op:    op,
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (node BinaryOperatorNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s\n", node.Op))
	buf.WriteString(Indent(fmt.Sprintf("%s\n", node.Expr1), "  "))
	buf.WriteString(Indent(fmt.Sprintf("%s\n", node.Expr2), "  "))

	return buf.String()
}

// Weight returns the number of registers used to evaluate the given
// ExpressionNode.
func Weight(n ExpressionNode) int {
	switch node := n.(type) {
	case *UnaryOperatorNode:
		return Weight(node.Expr)
	case *BinaryOperatorNode:
		lhsWeight := Max(Weight(node.Expr1), Weight(node.Expr2)+1)
		rhsWeight := Max(Weight(node.Expr1)+1, Weight(node.Expr2))
		return Min(lhsWeight, rhsWeight)
	}
	return 1
}
