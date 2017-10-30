package ast

import (
    "bytes"
)

type Program struct {
    Functions []FunctionNode
}

func NewProgram(functions []FunctionNode) Program {
    return Program{
        Functions: functions,
    }
}

func (p Program) String() string {
    var buf bytes.Buffer

    buf.WriteString("Program")

    return buf.String()
}

type FunctionNode struct {
    pos Position
    t TypeNode
    ident string
    stats []StatementNode
}

type Position struct {
    filename string
    lineNumber int
    colNumber int
}

// ---- StatementNodes ----

type StatementNode interface {

}

type DeclareNode struct {
    pos Position
    t TypeNode
    ident string
    rhs RHSNode
}

type AssignNode struct {
    pos Position
    lhs LHSNode
    rhs RHSNode
}

type ReadNode struct {
    pos Position
    expr ExpressionNode
}

type FreeNode struct {
    pos Position
    expr ExpressionNode
}

type ReturnNode struct {
    pos Position
    expr ExpressionNode
}

type ExitNode struct {
    pos Position
    expr ExpressionNode
}

type PrintNode struct {
    pos Position
    expr ExpressionNode
}

type PrintlnNode struct {
    pos Position
    expr ExpressionNode
}

type IfNode struct {
    pos Position
    expr ExpressionNode
    ifStats []StatementNode
    elseStats []StatementNode
}

type LoopNode struct {
    pos Position
    expr ExpressionNode
    stats []StatementNode
}

type ScopeNode struct {
    pos Position
    stats []StatementNode
}

// LHSNodes

type LHSNode interface {

}

type IdentifierNode struct {
    pos Position
    ident string
}

type PairFirstElementNode struct {
    pos Position
    ident string
    expr ExpressionNode
}

type PairSecondElementNode struct {
    pos Position
    ident string
    expr ExpressionNode
}

type ArrayElementNode struct {
    pos Position
    ident string
    expr ExpressionNode
}

// ---- RHSNodes ----

type RHSNode interface {

}

// ExpressionNode

type ArrayLiteralNode struct {
    pos Position
    exprs []ExpressionNode
}

type PairLiteralNode struct {
    pos Position
    expr1 ExpressionNode
    expr2 ExpressionNode
}

// PairFirstElementNode

// PairSecondElementNode

type FunctionCallNode struct {
    pos Position
    ident string
    exprs []ExpressionNode
}

// ---- TypeNodes ----

type TypeNode interface {
}

type BaseType int
const (
    INT BaseType = iota
    BOOL
    CHAR
    STRING
    PAIR
)

type BaseTypeNode struct {
    t BaseType
}

type ArrayTypeNode struct {
    dim int
    t BaseType
}

type PairTypeNode struct {
    t1 TypeNode
    t2 TypeNode
}

// ---- ExpressionNodes ----

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

type IntegerLiteralNode struct {
    pos Position
    val int
}

type BooleanLiteralNode struct {
    pos Position
    val bool
}

type CharacterLiteralNode struct {
    pos Position
    val rune
}

type StringLiteralNode struct {
    pos Position
    val string
}

// PairLiteralNode

// IdentifierNode

// ArrayElementNode

type UnaryOperatorNode struct {
    pos Position
    op UnaryOperator
    expr ExpressionNode
}

type BinaryOperatorNode struct {
    pos Position
    op BinaryOperator
    expr1 ExpressionNode
    expr2 ExpressionNode
}
