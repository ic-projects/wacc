package ast

import (
  "bytes"
	"strconv"
  "fmt"
)

type ExpressionNode interface {
}

type UnaryOperator int

const (
	NOT UnaryOperator = iota
	NEG
	LEN
	ORD
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

type BinaryOperator int

const (
	MUL BinaryOperator = iota
	DIV
	MOD
	ADD
	SUB
	GT
	GEQ
	LT
	LEQ
	EQ
	NEQ
	AND
	OR
)

func (binOp BinaryOperator) String() string {
	switch binOp {
	case MUL:
		return "- *"
	case DIV:
		return "- /"
	case MOD:
		return "- %"
	case ADD:
		return "- +"
	case SUB:
		return "- -"
	case GT:
		return "- >"
	case LT:
		return "- <"
	case LEQ:
		return "- <="
	case GEQ:
		return "- >="
	case EQ:
		return "- =="
	case NEQ:
		return "- !="
	case AND:
		return "- &&"
	case OR:
		return "- ||"
	}
	return "ERROR"
}

type IntegerLiteralNode struct {
	pos Position
	val int
}

func NewIntegerLiteralNode(pos Position, val int) IntegerLiteralNode {
	return IntegerLiteralNode{
		pos: pos,
		val: val,
	}
}

func (node IntegerLiteralNode) String() string {
	return fmt.Sprintf("- %d", node.val)
}

type BooleanLiteralNode struct {
	pos Position
	val bool
}

func NewBooleanLiteralNode(pos Position, val bool) BooleanLiteralNode {
	return BooleanLiteralNode{
		pos: pos,
		val: val,
	}
}

func (node BooleanLiteralNode) String() string {
	return fmt.Sprintf("- %s", strconv.FormatBool(node.val))
}

type CharacterLiteralNode struct {
	pos Position
	val rune
}

func NewCharacterLiteralNode(pos Position, val rune) CharacterLiteralNode {
	return CharacterLiteralNode{
		pos: pos,
		val: val,
	}
}

func (node CharacterLiteralNode) String() string {
	if node.val == '\000' {
		return "- '\\0'"
	}
	if node.val == '"' {
		return "- '\\\"'"
	}
	return fmt.Sprintf("- %q", node.val)
}

type StringLiteralNode struct {
	pos Position
	val string
}

func NewStringLiteralNode(pos Position, val string) StringLiteralNode {
	return StringLiteralNode{
		pos: pos,
		val: val,
	}
}

func (node StringLiteralNode) String() string {
	return fmt.Sprintf("- %s", node.val)
}

type PairLiteralNode struct {
	pos Position
}

func NewPairLiteralNode(pos Position) PairLiteralNode {
	return PairLiteralNode{
		pos: pos,
	}
}

func (node PairLiteralNode) String() string {
	return "- null\n"
}

// IdentifierNode

// ArrayElementNode

type UnaryOperatorNode struct {
	pos  Position
	op   UnaryOperator
	expr ExpressionNode
}

func NewUnaryOperatorNode(pos Position, op UnaryOperator, expr ExpressionNode) UnaryOperatorNode {
	return UnaryOperatorNode{
		pos:  pos,
		op:   op,
		expr: expr,
	}
}

func (node UnaryOperatorNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s\n", node.op))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.expr), "  "))

	return buf.String()
}

type BinaryOperatorNode struct {
	pos   Position
	op    BinaryOperator
	expr1 ExpressionNode
	expr2 ExpressionNode
}

func NewBinaryOperatorNode(pos Position, op BinaryOperator, expr1 ExpressionNode, expr2 ExpressionNode) BinaryOperatorNode {
	return BinaryOperatorNode{
		pos:   pos,
		op:    op,
		expr1: expr1,
		expr2: expr2,
	}
}

func (node BinaryOperatorNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s\n", node.op))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.expr1), "  "))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.expr2), "  "))

	return buf.String()
}

func BuildBinOpTree(first ExpressionNode, list []interface{}, position Position) ExpressionNode {
	if len(list) > 1 {
		var toparse []interface{}
		for i := 0; i < len(list)-1; i++ {
			toparse = append(toparse, list[i])
		}
		rest := BuildBinOpTree(first, toparse, position)
		last := list[len(list)-1].([]interface{})
		return NewBinaryOperatorNode(position, last[1].(BinaryOperator), rest, last[3])
	} else {
		return NewBinaryOperatorNode(position, list[0].([]interface{})[1].(BinaryOperator), first, list[0].([]interface{})[3])
	}
}
