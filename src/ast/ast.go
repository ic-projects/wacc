package ast

import (
    "bytes"
    "fmt"
    "strconv"
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
	return Program{
		functions: functions,
	}
}

func (program Program) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("Program"))
	for _, f := range program.functions {
		buf.WriteString(fmt.Sprintf("%s\n", f))
	}
	return buf.String()
}

type FunctionNode struct {
	pos    Position
	t      TypeNode
	ident  string
	params []ParameterNode
	stats  []StatementNode
}

func NewFunctionNode(pos Position, t TypeNode, ident string, params []ParameterNode, stats []StatementNode) FunctionNode {
	return FunctionNode{
		pos:    pos,
		t:      t,
		ident:  ident,
		params: params,
		stats:  stats,
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
	buf.WriteString(fmt.Sprintln(")"))
	for _, s := range node.stats {
		buf.WriteString(fmt.Sprintf("%s\n", s))
	}
	return buf.String()
}

type ParameterNode struct {
	pos   Position
	t     TypeNode
	ident string
}

func NewParameterNode(pos Position, t TypeNode, ident string) ParameterNode {
	return ParameterNode{
		pos:   pos,
		t:     t,
		ident: ident,
	}
}

func (node ParameterNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s %s", node.t, node.ident))

	return buf.String()
}

type Position struct {
	lineNumber int
	colNumber  int
	offset     int
}

func NewPosition(lineNumber int, colNumber int, offset int) Position {
	return Position{
		lineNumber: lineNumber,
		colNumber:  colNumber,
		offset:     offset,
	}
}

/**** StatementNodes ****/

type StatementNode interface {
}

type SkipNode struct {
	pos Position
}

func NewSkipNode(pos Position) SkipNode {
	return SkipNode{
		pos: pos,
	}
}

func (node SkipNode) String() string {
	return "- SKIP\n"
}

type DeclareNode struct {
	pos   Position
	t     TypeNode
	ident string
	rhs   RHSNode
}

func NewDeclareNode(pos Position, t TypeNode, ident string, rhs RHSNode) DeclareNode {
	return DeclareNode{
		pos:   pos,
		t:     t,
		ident: ident,
		rhs:   rhs,
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
	return AssignNode{
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
	pos  Position
	expr ExpressionNode
}

func NewReadNode(pos Position, expr ExpressionNode) ReadNode {
	return ReadNode{
		pos:  pos,
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
	pos  Position
	expr ExpressionNode
}

func NewFreeNode(pos Position, expr ExpressionNode) FreeNode {
	return FreeNode{
		pos:  pos,
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
	pos  Position
	expr ExpressionNode
}

func NewReturnNode(pos Position, expr ExpressionNode) ReturnNode {
	return ReturnNode{
		pos:  pos,
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
	pos  Position
	expr ExpressionNode
}

func NewExitNode(pos Position, expr ExpressionNode) ExitNode {
	return ExitNode{
		pos:  pos,
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
	pos  Position
	expr ExpressionNode
}

func NewPrintNode(pos Position, expr ExpressionNode) PrintNode {
	return PrintNode{
		pos:  pos,
		expr: expr,
	}
}

func (node PrintNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- PRINT"))
	buf.WriteString(fmt.Sprintf("- %s\n", node.expr))
	return buf.String()
}

type PrintlnNode struct {
	pos  Position
	expr ExpressionNode
}

func NewPrintlnNode(pos Position, expr ExpressionNode) PrintlnNode {
	return PrintlnNode{
		pos:  pos,
		expr: expr,
	}
}

func (node PrintlnNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- PRINTLN"))
	buf.WriteString(fmt.Sprintf("- %s\n", node.expr))
	return buf.String()
}

type IfNode struct {
	pos       Position
	expr      ExpressionNode
	ifStats   []StatementNode
	elseStats []StatementNode
}

func NewIfNode(pos Position, expr ExpressionNode, ifStats []StatementNode, elseStats []StatementNode) IfNode {
	return IfNode{
		pos:       pos,
		expr:      expr,
		ifStats:   ifStats,
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
	pos   Position
	expr  ExpressionNode
	stats []StatementNode
}

func NewLoopNode(pos Position, expr ExpressionNode, stats []StatementNode) LoopNode {
	return LoopNode{
		pos:   pos,
		expr:  expr,
		stats: stats,
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
	pos   Position
	stats []StatementNode
}

func NewScopeNode(pos Position, stats []StatementNode) ScopeNode {
	return ScopeNode{
		pos:   pos,
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
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s\n", node.ident))
	return buf.String()
}

type PairFirstElementNode struct {
	pos   Position
	expr  ExpressionNode
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
	buf.WriteString(fmt.Sprintf("%s\n", node.expr))
	return buf.String()
}

type PairSecondElementNode struct {
	pos   Position
	expr  ExpressionNode
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
	buf.WriteString(fmt.Sprintf("%s\n", node.expr))
	return buf.String()
}

type ArrayElementNode struct {
	pos   Position
	ident string
	exprs  []ExpressionNode
}

func NewArrayElementNode(pos Position, ident string, exprs []ExpressionNode) ArrayElementNode {
	return ArrayElementNode{
		pos:   pos,
		ident: ident,
		exprs:  exprs,
	}
}

func (node ArrayElementNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s\n", node.ident))
  for _, e := range node.exprs {
	  buf.WriteString(fmt.Sprintln("  - []"))
	  buf.WriteString(fmt.Sprintf("  %s\n", e))
  }
	return buf.String()
}

/**** RHSNodes ****/

type RHSNode interface {
}

// ExpressionNode

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
		buf.WriteString(fmt.Sprintf("%s\n", e))
	}
	return buf.String()
}

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
    buf.WriteString(fmt.Sprintln("  - FST"))
    buf.WriteString(fmt.Sprintf("  %s\n", node.fst))
    buf.WriteString(fmt.Sprintln("  - SND"))
    buf.WriteString(fmt.Sprintf("  %s\n", node.snd))
    return buf.String()
}

// PairFirstElementNode

// PairSecondElementNode

type FunctionCallNode struct {
	pos   Position
	ident string
	exprs []ExpressionNode
}

func NewFunctionCallNode(pos Position, ident string, exprs []ExpressionNode) FunctionCallNode {
	return FunctionCallNode{
		pos:   pos,
		ident: ident,
		exprs: exprs,
	}
}

func (node FunctionCallNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s\n", node.ident))
	for _, e := range node.exprs {
		buf.WriteString(fmt.Sprintf("  %s\n", e))
	}
	return buf.String()
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

func (t BaseType) String() string {
	switch t {
	case INT:
		return "int"
	case BOOL:
		return "bool"
	case CHAR:
		return "char"
	case STRING:
		return "string"
	case PAIR:
		return "pair"
	}
	return ""
}

type BaseTypeNode struct {
	t BaseType
}

func NewBaseTypeNode(t BaseType) BaseTypeNode {
	return BaseTypeNode{
		t: t,
	}
}

func (node BaseTypeNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s\n", node.t))
	return buf.String()
}

type ArrayTypeNode struct {
	t   TypeNode
	dim int
}

func NewArrayTypeNode(t TypeNode, dim int) ArrayTypeNode {
	return ArrayTypeNode{
		t:   t,
		dim: dim,
	}
}

func (node ArrayTypeNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s", node.t))
	for i := 0; i < node.dim; i++ {
		buf.WriteString("[]")
	}
	buf.WriteString("\n")
	return buf.String()
}

type PairTypeNode struct {
	t1 TypeNode
	t2 TypeNode
}

func NewPairTypeNode(t1 TypeNode, t2 TypeNode) PairTypeNode {
	return PairTypeNode{
		t1: t1,
		t2: t2,
	}
}

func (node PairTypeNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("- %s%s", node.t1, node.t2))
	return buf.String()
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

func (unOp UnaryOperator) String() string {
  switch unOp {
  case NOT:
    return "- !"
  case NEG:
    return "- -"
  case LEN:
    return "- LEN"
  case ORD:
    return "- ORD"
  case CHR:
    return "- CHR"
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
  return IntegerLiteralNode {
    pos: pos,
    val: val,
  }
}

func (node IntegerLiteralNode) String() string {
  var buf bytes.Buffer

  buf.WriteString(fmt.Sprintf("- %d", node.val))

  return buf.String();
}

type BooleanLiteralNode struct {
	pos Position
	val bool
}

func NewBooleanLiteralNode(pos Position, val bool) BooleanLiteralNode {
  return BooleanLiteralNode {
    pos: pos,
    val: val,
  }
}

func (node BooleanLiteralNode) String() string {
  var buf bytes.Buffer

  buf.WriteString(fmt.Sprintf("- %s", strconv.FormatBool(node.val)))

  return buf.String()
}

type CharacterLiteralNode struct {
	pos Position
	val rune
}

func NewCharacterLiteralNode(pos Position, val rune) CharacterLiteralNode {
  return CharacterLiteralNode {
    pos: pos,
    val: val,
  }
}

func (node CharacterLiteralNode) String() string {
  var buf bytes.Buffer

  buf.WriteString(fmt.Sprintf("- '%c'", node.val))

  return buf.String()
}

type StringLiteralNode struct {
	pos Position
	val string
}

func NewStringLiteralNode(pos Position, val string) StringLiteralNode {
  return StringLiteralNode {
    pos: pos,
    val: val,
  }
}

func (node StringLiteralNode) String() string {
  var buf bytes.Buffer

  buf.WriteString(fmt.Sprintf("- %s", node.val))

  return buf.String()
}

type PairLiteralNode struct {
	pos Position
}

func NewPairLiteralNode(pos Position, fst ExpressionNode, snd ExpressionNode) PairLiteralNode {
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
  return UnaryOperatorNode {
    pos: pos,
    op: op,
    expr: expr,
  }
}

func (node UnaryOperatorNode) String() string {
  var buf bytes.Buffer

  buf.WriteString(fmt.Sprintf("- %s\n", node.op))
  buf.WriteString(fmt.Sprintf("- %s\n", node.expr))

  return buf.String()
}

type BinaryOperatorNode struct {
	pos   Position
	op    BinaryOperator
	expr1 ExpressionNode
	expr2 ExpressionNode
}

func NewBinaryOperatorNode(pos Position, op BinaryOperator, expr1 ExpressionNode, expr2 ExpressionNode) BinaryOperatorNode {
  return BinaryOperatorNode {
    pos: pos,
    op: op,
    expr1: expr1,
    expr2: expr2,
  }
}

func (node BinaryOperatorNode) String() string {
  var buf bytes.Buffer

  buf.WriteString(fmt.Sprintf("- %s\n", node.op))
  buf.WriteString(fmt.Sprintf("- %s\n", node.expr1))
  buf.WriteString(fmt.Sprintf("- %s\n", node.expr2))

  return buf.String()
}
