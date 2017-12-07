package ast

import (
	"bytes"
	"fmt"
	"strconv"
	"utils"
)

// ExpressionNode is an interface for expression nodes to implement.
type ExpressionNode interface {
	ProgramNode
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
		return ToValue(a.T).(ArrayTypeNode).GetDimElement(len(node.Exprs))
	case *DynamicTypeNode:
		return node.getValue()
	case *StructElementNode:
		return node.StructType.T
	case *IdentifierNode:
		v, _ := s.SearchForIdent(node.Ident)
		return v.T
	case *PointerDereferenceNode:
		v, _ := s.SearchForIdent(node.Ident.Ident)
		return (v.T.(*PointerTypeNode)).T
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
	position utils.Position,
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

// Weight returns the number of registers used to evaluate the given
// ExpressionNode.
func Weight(n ExpressionNode) int {
	switch node := n.(type) {
	case *UnaryOperatorNode:
		return Weight(node.Expr)
	case *BinaryOperatorNode:
		lhsWeight := utils.Max(Weight(node.Expr1), Weight(node.Expr2)+1)
		rhsWeight := utils.Max(Weight(node.Expr1)+1, Weight(node.Expr2))
		return utils.Min(lhsWeight, rhsWeight)
	}
	return 1
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
	Pos utils.Position
	Val int
}

// NewIntegerLiteralNode builds an IntegerLiteralNode
func NewIntegerLiteralNode(pos utils.Position, val int) *IntegerLiteralNode {
	return &IntegerLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node *IntegerLiteralNode) String() string {
	return fmt.Sprintf("- %d", node.Val)
}

func (node *IntegerLiteralNode) walkNode(visitor Visitor) {
}

/**************** BOOLEAN LITERAL NODE ****************/

// BooleanLiteralNode is a struct which stores the position and value of a
// boolean literal.
//
// E.g.
//  false
type BooleanLiteralNode struct {
	Pos utils.Position
	Val bool
}

// NewBooleanLiteralNode builds an BooleanLiteralNode
func NewBooleanLiteralNode(pos utils.Position, val bool) *BooleanLiteralNode {
	return &BooleanLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node *BooleanLiteralNode) String() string {
	return fmt.Sprintf("- %s", strconv.FormatBool(node.Val))
}

func (node *BooleanLiteralNode) walkNode(visitor Visitor) {
}

/**************** CHARACTER LITERAL NODE ****************/

// CharacterLiteralNode is a struct which stores the position and value of a
// character literal.
//
// E.g.
//  'c'
type CharacterLiteralNode struct {
	Pos utils.Position
	Val rune
}

// NewCharacterLiteralNode builds a CharacterLiteralNode
func NewCharacterLiteralNode(
	pos utils.Position,
	val rune,
) *CharacterLiteralNode {
	return &CharacterLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node *CharacterLiteralNode) String() string {
	if node.Val == '\000' {
		return "- '\\0'"
	}
	if node.Val == '"' {
		return "- '\\\"'"
	}
	return fmt.Sprintf("- %q", node.Val)
}

func (node *CharacterLiteralNode) walkNode(visitor Visitor) {
}

/**************** STRING LITERAL NODE ****************/

// StringLiteralNode is a struct which stores the position and value of a string
// literal.
//
// E.g.
//  "Hello World!"
type StringLiteralNode struct {
	Pos utils.Position
	Val string
}

// NewStringLiteralNode builds a StringLiteralNode
func NewStringLiteralNode(pos utils.Position, val string) *StringLiteralNode {
	return &StringLiteralNode{
		Pos: pos,
		Val: val,
	}
}

func (node *StringLiteralNode) String() string {
	return fmt.Sprintf("- \"%s\"", node.Val)
}

func (node *StringLiteralNode) walkNode(visitor Visitor) {
}

/**************** NULL NODE ****************/

// NullNode is a struct which stores the position of a pair literal.
// This does not store the value of the pair literal since the value of a pair
// literal is always null.
type NullNode struct {
	Pos utils.Position
}

// NewNullNode builds a NullNode
func NewNullNode(pos utils.Position) *NullNode {
	return &NullNode{
		Pos: pos,
	}
}

func (node *NullNode) String() string {
	return "- null\n"
}

func (node *NullNode) walkNode(visitor Visitor) {
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
	Pos  utils.Position
	Op   UnaryOperator
	Expr ExpressionNode
}

// NewUnaryOperatorNode builds a UnaryOperatorNode
func NewUnaryOperatorNode(
	pos utils.Position,
	op UnaryOperator,
	expr ExpressionNode,
) *UnaryOperatorNode {
	return &UnaryOperatorNode{
		Pos:  pos,
		Op:   op,
		Expr: expr,
	}
}

func (node *UnaryOperatorNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s\n", node.Op))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.Expr), "  "))

	return buf.String()
}

func (node *UnaryOperatorNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
}

/**************** BINARY OPERATOR NODE ****************/

// BinaryOperatorNode is a struct which stores the position, (binary) operator
// and the two expressions of a binary operation on two expressions.
//
// E.g.
//
//  5 + 2
type BinaryOperatorNode struct {
	Pos   utils.Position
	Op    BinaryOperator
	Expr1 ExpressionNode
	Expr2 ExpressionNode
}

// NewBinaryOperatorNode builds a BinaryOperatorNode
func NewBinaryOperatorNode(
	pos utils.Position,
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

func (node *BinaryOperatorNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s\n", node.Op))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.Expr1), "  "))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.Expr2), "  "))

	return buf.String()
}

func (node *BinaryOperatorNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr1)
	Walk(visitor, node.Expr2)
}
