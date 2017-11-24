package ast

import (
	"bytes"
	"fmt"
	"utils"
)

// StatementNode is an empty interface for statement nodes to implement.
type StatementNode interface {
}

/******************** STATEMENT HELPER FUNCTIONS ****************/

// FinalStatIsValid given the last statement from a statement list, this function
// traverses to the last statement checks that statement is a valid end statement,
// such as a return or exit.
func FinalStatIsValid(s StatementNode) bool {
	switch s.(type) {
	case ReturnNode:
		return true
	case ExitNode:
		return true
	case ScopeNode:
		stats := s.(ScopeNode).Stats
		finalStat := stats[len(stats)-1]
		return FinalStatIsValid(finalStat)
	case IfNode:
		ifStats := s.(IfNode).IfStats
		ifFinalStat := ifStats[len(ifStats)-1]
		elseStats := s.(IfNode).ElseStats
		elseFinalStat := elseStats[len(elseStats)-1]
		return FinalStatIsValid(ifFinalStat) && FinalStatIsValid(elseFinalStat)
	default:
		return false
	}
}

/******************** SKIP NODE ****************/

// SkipNode is a struct that stores the position of a skip statement.
type SkipNode struct {
	Pos Position
}

func NewSkipNode(pos Position) SkipNode {
	return SkipNode{
		Pos: pos,
	}
}

func (node SkipNode) String() string {
	return "- SKIP\n"
}

/******************** DECLARE NODE ****************/

// DeclareNode is a struct that stores the position, type, identifier and
// assignment of a declaration.
//
// E.g.
//
//  int i = 5
type DeclareNode struct {
	Pos   Position
	T     TypeNode
	Ident IdentifierNode
	Rhs   RHSNode
}

func NewDeclareNode(pos Position, t TypeNode, ident IdentifierNode, rhs RHSNode) DeclareNode {
	return DeclareNode{
		Pos:   pos,
		T:     t,
		Ident: ident,
		Rhs:   rhs,
	}
}

func (node DeclareNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- DECLARE"))
	buf.WriteString(fmt.Sprintln("  - TYPE"))
	buf.WriteString(utils.Indent(fmt.Sprintf("- %s\n", node.T), "    "))
	buf.WriteString(fmt.Sprintln("  - LHS"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.Ident), "    "))
	buf.WriteString(fmt.Sprintln("  - RHS"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.Rhs), "    "))
	return buf.String()
}

/******************** ASSIGN NODE ****************/

// AssignNode stores the position, left hand side and right hand side of an
// assignment statement.
//
// E.g.
//  i = 4
type AssignNode struct {
	Pos Position
	Lhs LHSNode
	Rhs RHSNode
}

func NewAssignNode(pos Position, lhs LHSNode, rhs RHSNode) AssignNode {
	return AssignNode{
		Pos: pos,
		Lhs: lhs,
		Rhs: rhs,
	}
}

func (node AssignNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- ASSIGNMENT"))
	buf.WriteString(fmt.Sprintln("  - LHS"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Lhs), "    "))
	buf.WriteString(fmt.Sprintln("  - RHS"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Rhs), "    "))
	return buf.String()
}

/******************** READ NODE ****************/

// ReadNode is a struct that stores the position and expression of a read
// statement.
//
// E.g.
//  read i
type ReadNode struct {
	Pos Position
	Lhs LHSNode
}

func NewReadNode(pos Position, lhs LHSNode) ReadNode {
	return ReadNode{
		Pos: pos,
		Lhs: lhs,
	}
}

func (node ReadNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- READ"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Lhs), "  "))
	return buf.String()
}

/******************** FREE NODE ****************/

// FreeNode stores the position and expression of a free statement.
//
// E.g.
//
//  free p
type FreeNode struct {
	Pos  Position
	Expr ExpressionNode
}

func NewFreeNode(pos Position, expr ExpressionNode) FreeNode {
	return FreeNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node FreeNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- FREE"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "  "))
	return buf.String()
}

/******************** RETURN NODE ****************/

// ReturnNode stores the position and expression of a return statement.
//
// E.g.
//
//  return 5
type ReturnNode struct {
	Pos  Position
	Expr ExpressionNode
}

func NewReturnNode(pos Position, expr ExpressionNode) ReturnNode {
	return ReturnNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node ReturnNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- RETURN"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "  "))
	return buf.String()
}

/******************** EXIT NODE ****************/

// ExitNode stores the position and expression of an exit statement.
//
// E.g.
//
//  exit 255
type ExitNode struct {
	Pos  Position
	Expr ExpressionNode
}

func NewExitNode(pos Position, expr ExpressionNode) ExitNode {
	return ExitNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node ExitNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- EXIT"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "  "))
	return buf.String()
}

/******************** PRINT NODE ****************/

// PrintNode stores the position and expression of an print statement.
//
// E.g.
//
//  print "printing"
type PrintNode struct {
	Pos  Position
	Expr ExpressionNode
}

func NewPrintNode(pos Position, expr ExpressionNode) PrintNode {
	return PrintNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node PrintNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- PRINT"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "  "))
	return buf.String()
}

/******************** PRINTLN NODE ****************/

// PrintlnNode stores the position and expression of an println statement.
//
// E.g.
//
//  println "printing"
type PrintlnNode struct {
	Pos  Position
	Expr ExpressionNode
}

func NewPrintlnNode(pos Position, expr ExpressionNode) PrintlnNode {
	return PrintlnNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node PrintlnNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- PRINTLN"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "  "))
	return buf.String()
}

/******************** IF NODE ****************/

// IfNode stores the position, condition and the two branches of an if else
// statement.
//
// E.g.
//
//  if true then skip else skip fi
type IfNode struct {
	Pos       Position
	Expr      ExpressionNode
	IfStats   []StatementNode
	ElseStats []StatementNode
}

func NewIfNode(pos Position, expr ExpressionNode, ifStats []StatementNode, elseStats []StatementNode) IfNode {
	return IfNode{
		Pos:       pos,
		Expr:      expr,
		IfStats:   ifStats,
		ElseStats: elseStats,
	}
}

func (node IfNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- IF"))
	buf.WriteString(utils.Indent(fmt.Sprintln("- CONDITION"), "  "))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "    "))
	buf.WriteString(utils.Indent(fmt.Sprintln("- THEN"), "  "))
	for _, s := range node.IfStats {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "    "))
	}
	buf.WriteString(utils.Indent(fmt.Sprintln("- ELSE"), "  "))
	for _, s := range node.ElseStats {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "    "))
	}
	return buf.String()
}

/******************** LOOP NODE ****************/

// LoopNode stores the position, condition and loop statements for a loop
// while loop statement.
//
// E.g.
//
//  while true do skip done
type LoopNode struct {
	Pos   Position
	Expr  ExpressionNode
	Stats []StatementNode
}

func NewLoopNode(pos Position, expr ExpressionNode, stats []StatementNode) LoopNode {
	return LoopNode{
		Pos:   pos,
		Expr:  expr,
		Stats: stats,
	}
}

func (node LoopNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- LOOP"))
	buf.WriteString(fmt.Sprintln("  - CONDITION"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "    "))
	buf.WriteString(fmt.Sprintln("  - DO"))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "    "))
	}
	return buf.String()
}

/******************** SCOPE NODE ****************/

// ScopeNode stores the position and statement of a new scope.
//
// E.g.
//
//  begin skip end
type ScopeNode struct {
	Pos   Position
	Stats []StatementNode
}

func NewScopeNode(pos Position, stats []StatementNode) ScopeNode {
	return ScopeNode{
		Pos:   pos,
		Stats: stats,
	}
}

func (node ScopeNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- SCOPE"))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "  "))
	}
	return buf.String()
}
