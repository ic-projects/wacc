package ast

import (
    "bytes"
    "fmt"
)

/*  WACC Abstract Syntax Tree

    For similar implementations, see:
     *  Pigeon itself:
        PEG: https://github.com/mna/pigeon/blob/master/grammar/pigeon.peg
        AST: https://github.com/mna/pigeon/blob/master/ast/ast.go

     *  Pigeon 'indentation' example:
        PEG: https://github.com/mna/pigeon/blob/master/examples/indentation/indentation.peg
        AST: https://github.com/mna/pigeon/blob/master/examples/indentation/indentation_ast.go

     *  Logstash example:
        PEG: https://github.com/breml/logstash-config/blob/master/logstash_config.peg
        AST: https://github.com/breml/logstash-config/blob/master/ast/ast.go

*/

type Program struct {
    functions []FunctionNode
}

func NewProgram(functions []FunctionNode) Program {
    return Program {
        functions: functions,
    }
}

func (program Program) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("Program"))
    for _, f := range program.functions {
        buf.WriteString(fmt.Sprintf("%v\n", f))
    }

    return buf.String()
}

type FunctionNode struct {
    pos Position
    t TypeNode
    ident string
    params []ParameterNode
    stats []StatementNode
}

func NewFunctionNode(pos Position, t TypeNode, ident string, params []ParameterNode, stats []StatementNode) FunctionNode {
    return FunctionNode {
        pos: pos,
        t: t,
        ident: ident,
        params: params,
        stats: stats,
    }
}

func (node FunctionNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintf("- %s %s(", node.t, node.ident))
    for i, p := range node.params {
        if i == 0 {
            buf.WriteString(fmt.Sprintf("%s", p))
        } else {
            buf.WriteString(fmt.Sprintf(", %s", p))
        }
    }
    buf.WriteString(fmt.Sprintln(")"));
    for _, s := range node.stats {
        buf.WriteString(fmt.Sprintf("%s\n", s))
    }

    return buf.String()
}

type ParameterNode struct {
    pos Position
    t TypeNode
    ident string
}

func NewParameterNode(pos Position, t TypeNode, ident string) ParameterNode {
    return ParameterNode {
        pos: pos,
        t: t,
        ident: ident,
    }
}

type Position struct {
    lineNumber int
    colNumber int
    offset int
}

func NewPosition(lineNumber int, colNumber int, offset int) Position {
    return Position {
        lineNumber: lineNumber,
        colNumber: colNumber,
        offset: offset,
    }
}

/**** StatementNodes ****/

type StatementNode interface {

}

type SkipNode struct {
    pos Position
}

func NewSkipNode(pos Position) SkipNode {
    return SkipNode {
        pos: pos,
    }
}

func (node SkipNode) String() string {
    return "- SKIP"
}

type DeclareNode struct {
    pos Position
    t TypeNode
    ident string
    rhs RHSNode
}

func NewDeclareNode(pos Position, t TypeNode, ident string, rhs RHSNode) DeclareNode {
    return DeclareNode {
        pos: pos,
        t: t,
        ident: ident,
        rhs: rhs,
    }
}

func (node DeclareNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- DECLARE"))
    buf.WriteString(fmt.Sprintln("- TYPE"))
    buf.WriteString(fmt.Sprintf("%s\n", node.t))
    buf.WriteString(fmt.Sprintln("- LHS"))
    buf.WriteString(fmt.Sprintf("%s\n", node.ident))
    buf.WriteString(fmt.Sprintln("- RHS"))
    buf.WriteString(fmt.Sprintf("%s\n", node.rhs))

    return buf.String()
}

type AssignNode struct {
    pos Position
    lhs LHSNode
    rhs RHSNode
}

func NewAssignNode(pos Position, lhs LHSNode, rhs RHSNode) AssignNode {
    return AssignNode {
        pos: pos,
        lhs: lhs,
        rhs: rhs,
    }
}

func (node AssignNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- ASSIGNMENT"))
    buf.WriteString(fmt.Sprintln("- LHS"))
    buf.WriteString(fmt.Sprintf("%s\n", node.lhs))
    buf.WriteString(fmt.Sprintln("- RHS"))
    buf.WriteString(fmt.Sprintf("%s\n", node.rhs))

    return buf.String()
}

type ReadNode struct {
    pos Position
    expr ExpressionNode
}

func NewReadNode(pos Position, expr ExpressionNode) ReadNode {
    return ReadNode {
        pos: pos,
        expr: expr,
    }
}

func (node ReadNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- READ"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))

    return buf.String()
}

type FreeNode struct {
    pos Position
    expr ExpressionNode
}

func NewFreeNode(pos Position, expr ExpressionNode) FreeNode {
    return FreeNode {
        pos: pos,
        expr: expr,
    }
}

func (node FreeNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- FREE"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))

    return buf.String()
}

type ReturnNode struct {
    pos Position
    expr ExpressionNode
}

func NewReturnNode(pos Position, expr ExpressionNode) ReturnNode {
    return ReturnNode {
        pos: pos,
        expr: expr,
    }
}

func (node ReturnNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- RETURN"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))

    return buf.String()
}

type ExitNode struct {
    pos Position
    expr ExpressionNode
}

func NewExitNode(pos Position, expr ExpressionNode) ExitNode {
    return ExitNode {
        pos: pos,
        expr: expr,
    }
}

func (node ExitNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- EXIT"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))

    return buf.String()
}

type PrintNode struct {
    pos Position
    expr ExpressionNode
}

func NewPrintNode(pos Position, expr ExpressionNode) PrintNode {
    return PrintNode {
        pos: pos,
        expr: expr,
    }
}

func (node PrintNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- PRINT"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))

    return buf.String()
}

type PrintlnNode struct {
    pos Position
    expr ExpressionNode
}

func NewPrintlnNode(pos Position, expr ExpressionNode) PrintlnNode {
    return PrintlnNode {
        pos: pos,
        expr: expr,
    }
}

func (node PrintlnNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- PRINTLN"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))

    return buf.String()
}

type IfNode struct {
    pos Position
    expr ExpressionNode
    ifStats []StatementNode
    elseStats []StatementNode
}

func NewIfNode(pos Position, expr ExpressionNode, ifStats []StatementNode, elseStats []StatementNode) IfNode {
    return IfNode {
        pos: pos,
        expr: expr,
        ifStats: ifStats,
        elseStats: elseStats,
    }
}

func (node IfNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- CONDITION"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))
    buf.WriteString(fmt.Sprintln("- THEN"))
    for _, s := range node.ifStats {
        buf.WriteString(fmt.Sprintf("- %s\n", s))
    }
    buf.WriteString(fmt.Sprintln("- ELSE"))
    buf.WriteString(fmt.Sprintf("%s\n", node.elseStats))
    for _, s := range node.elseStats {
        buf.WriteString(fmt.Sprintf("%s\n", s))
    }

    return buf.String()
}

type LoopNode struct {
    pos Position
    expr ExpressionNode
    stats []StatementNode
}

func NewLoopNode(pos Position, expr ExpressionNode, ifStats []StatementNode, elseStats []StatementNode) LoopNode {
    return LoopNode {
        pos: pos,
        expr: expr,
        stats: ifStats,
    }
}

func (node LoopNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- CONDITION"))
    buf.WriteString(fmt.Sprintf("%s\n", node.expr))
    buf.WriteString(fmt.Sprintln("- DO"))
    for _, s := range node.stats {
        buf.WriteString(fmt.Sprintf("%s\n", s))
    }

    return buf.String()
}

type ScopeNode struct {
    pos Position
    stats []StatementNode
}

func NewScopeNode(pos Position, stats []StatementNode) ScopeNode {
    return ScopeNode {
        pos: pos,
        stats: stats,
    }
}

func (node ScopeNode) String() string {
    var buf bytes.Buffer

    buf.WriteString(fmt.Sprintln("- SCOPE"))
    for _, s := range node.stats {
        buf.WriteString(fmt.Sprintf("%s\n", s))
    }

    return buf.String()
}

/**** LHSNodes ****/

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

/**** RHSNodes ****/

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

/**** TypeNodes ****/

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

func NewBaseTypeNode(t BaseType) BaseTypeNode {
    return BaseTypeNode{
        t: t,
    }
}

type ArrayTypeNode struct {
    dim int
    t BaseType
}

type PairTypeNode struct {
    t1 TypeNode
    t2 TypeNode
}

/**** ExpressionNodes ****/

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
