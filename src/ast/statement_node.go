package ast

import (
	"bytes"
	"fmt"
	"utils"
)

// StatementNode is an empty interface for statement nodes to implement.
type StatementNode interface {
	fmt.Stringer
}

/**************** STATEMENT HELPER FUNCTIONS ****************/

// FinalStatIsValid given the last statement from a statement list, this
// function traverses to the last statement checks that statement is a valid end
// statement, such as a return or exit.
func FinalStatIsValid(s StatementNode) bool {
	switch s.(type) {
	case *ReturnNode:
		return true
	case *ExitNode:
		return true
	case *ScopeNode:
		stats := s.(*ScopeNode).Stats
		finalStat := stats[len(stats)-1]
		return FinalStatIsValid(finalStat)
	case *IfNode:
		ifStats := s.(*IfNode).IfStats
		ifFinalStat := ifStats[len(ifStats)-1]
		elseStats := s.(*IfNode).ElseStats
		elseFinalStat := elseStats[len(elseStats)-1]
		return FinalStatIsValid(ifFinalStat) && FinalStatIsValid(elseFinalStat)
	default:
		return false
	}
}

/**************** SKIP NODE ****************/

// SkipNode is a struct that stores the position of a skip statement.
type SkipNode struct {
	Pos utils.Position
}

// NewSkipNode builds a SkipNode
func NewSkipNode(pos utils.Position) *SkipNode {
	return &SkipNode{
		Pos: pos,
	}
}

func (node SkipNode) String() string {
	return "- SKIP\n"
}

/**************** DECLARE NODE ****************/

// DeclareNode is a struct that stores the position, type, identifier and
// assignment of a declaration.
//
// E.g.
//
//  int i = 5
type DeclareNode struct {
	Pos   utils.Position
	T     TypeNode
	Ident *IdentifierNode
	RHS   RHSNode
}

// NewDeclareNode builds a DeclareNode
func NewDeclareNode(
	pos utils.Position,
	t TypeNode,
	ident *IdentifierNode,
	rhs RHSNode,
) *DeclareNode {
	return &DeclareNode{
		Pos:   pos,
		T:     t,
		Ident: ident,
		RHS:   rhs,
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
	buf.WriteString(utils.Indent(fmt.Sprintf("%s\n", node.RHS), "    "))
	return buf.String()
}

/**************** ASSIGN NODE ****************/

// AssignNode stores the position, left hand side and right hand side of an
// assignment statement.
//
// E.g.
//  i = 4
type AssignNode struct {
	Pos utils.Position
	LHS LHSNode
	RHS RHSNode
}

// NewAssignNode builds a AssignNode
func NewAssignNode(pos utils.Position, lhs LHSNode, rhs RHSNode) *AssignNode {
	return &AssignNode{
		Pos: pos,
		LHS: lhs,
		RHS: rhs,
	}
}

func (node AssignNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- ASSIGNMENT"))
	buf.WriteString(fmt.Sprintln("  - LHS"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.LHS), "    "))
	buf.WriteString(fmt.Sprintln("  - RHS"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.RHS), "    "))
	return buf.String()
}

/**************** READ NODE ****************/

// ReadNode is a struct that stores the position and expression of a read
// statement.
//
// E.g.
//  read i
type ReadNode struct {
	Pos utils.Position
	LHS LHSNode
}

// NewReadNode builds a ReadNode
func NewReadNode(pos utils.Position, lhs LHSNode) *ReadNode {
	return &ReadNode{
		Pos: pos,
		LHS: lhs,
	}
}

func (node ReadNode) String() string {
	return writeSimpleString("READ", node.LHS)
}

/**************** FREE NODE ****************/

// FreeNode stores the position and expression of a free statement.
//
// E.g.
//
//  free p
type FreeNode struct {
	Pos  utils.Position
	Expr ExpressionNode
}

// NewFreeNode builds a FreeNode
func NewFreeNode(pos utils.Position, expr ExpressionNode) *FreeNode {
	return &FreeNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node FreeNode) String() string {
	return writeSimpleString("FREE", node.Expr)
}

/**************** RETURN NODE ****************/

// ReturnNode stores the position and expression of a return statement.
//
// E.g.
//
//  return 5
type ReturnNode struct {
	Pos  utils.Position
	Expr ExpressionNode
}

// NewReturnNode builds a ReturnNode
func NewReturnNode(pos utils.Position, expr ExpressionNode) *ReturnNode {
	return &ReturnNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node ReturnNode) String() string {
	return writeSimpleString("RETURN", node.Expr)
}

/**************** EXIT NODE ****************/

// ExitNode stores the position and expression of an exit statement.
//
// E.g.
//
//  exit 255
type ExitNode struct {
	Pos  utils.Position
	Expr ExpressionNode
}

// NewExitNode builds a ExitNode
func NewExitNode(pos utils.Position, expr ExpressionNode) *ExitNode {
	return &ExitNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node ExitNode) String() string {
	return writeSimpleString("EXIT", node.Expr)
}

/**************** PRINT NODE ****************/

// PrintNode stores the position and expression of an print statement.
//
// E.g.
//
//  print "printing"
type PrintNode struct {
	Pos  utils.Position
	Expr ExpressionNode
}

// NewPrintNode builds a PrintNode
func NewPrintNode(pos utils.Position, expr ExpressionNode) *PrintNode {
	return &PrintNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node PrintNode) String() string {
	return writeSimpleString("PRINT", node.Expr)
}

/**************** PRINTLN NODE ****************/

// PrintlnNode stores the position and expression of an println statement.
//
// E.g.
//
//  println "printing"
type PrintlnNode struct {
	Pos  utils.Position
	Expr ExpressionNode
}

// NewPrintlnNode builds a PrintlnNode
func NewPrintlnNode(pos utils.Position, expr ExpressionNode) *PrintlnNode {
	return &PrintlnNode{
		Pos:  pos,
		Expr: expr,
	}
}

func (node PrintlnNode) String() string {
	return writeSimpleString("PRINTLN", node.Expr)
}

/**************** SWITCH NODE ****************/

// SwitchNode stores the position, condition and the branches of a switch
// statement.
//
// E.g.
//
//  todo
type SwitchNode struct {
	Pos   utils.Position
	Expr  ExpressionNode
	Cases []CaseNode
}

func NewSwitchNode(pos utils.Position, expr ExpressionNode, cases []CaseNode) *SwitchNode {
	return &SwitchNode{
		Pos:   pos,
		Expr:  expr,
		Cases: cases,
	}
}

func (node SwitchNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- SWITCH"))
	buf.WriteString(utils.Indent(fmt.Sprintln("- EXPRESSION"), "  "))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "    "))
	//	buf.WriteString(utils.Indent(fmt.Sprintln("- THEN"), "  "))
	for _, s := range node.Cases {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "    "))
	}
	return buf.String()
}

type CaseNode struct {
	Pos   utils.Position
	Expr  ExpressionNode
	Stats []StatementNode
}

func NewCaseNode(pos utils.Position, expr ExpressionNode, stats []StatementNode) CaseNode {
	return CaseNode{
		Pos:   pos,
		Expr:  expr,
		Stats: stats,
	}
}

func (node CaseNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- CASE"))
	buf.WriteString(utils.Indent(fmt.Sprintln("- EXPRESSION"), "  "))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "    "))
	buf.WriteString(utils.Indent(fmt.Sprintln("- THEN"), "  "))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "    "))
	}
	return buf.String()
}

/**************** IF NODE ****************/

// IfNode stores the position, condition and the two branches of an if else
// statement.
//
// E.g.
//
//  if true then skip else skip fi
type IfNode struct {
	Pos       utils.Position
	Expr      ExpressionNode
	IfStats   []StatementNode
	ElseStats []StatementNode
}

// NewIfNode builds a IfNode
func NewIfNode(
	pos utils.Position,
	expr ExpressionNode,
	ifStats []StatementNode,
	elseStats []StatementNode,
) *IfNode {
	return &IfNode{
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

/**************** LOOP NODE ****************/

// LoopNode stores the position, condition and loop statements for a loop
// while loop statement.
//
// E.g.
//
//  while true do skip done
type LoopNode struct {
	Pos   utils.Position
	Expr  ExpressionNode
	Stats []StatementNode
}

// NewLoopNode builds a LoopNode
func NewLoopNode(
	pos utils.Position,
	expr ExpressionNode,
	stats []StatementNode,
) *LoopNode {
	return &LoopNode{
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

/**************** FOR LOOP NODE ****************/

// ForLoopNode stores the position, initial statement, condition, update
// statement and loop statements for a for loop statement
//
// E.g.
//
//  for int i = 0; i > 3; i = i + 1 do skip done
type ForLoopNode struct {
	Pos     utils.Position
	Initial StatementNode
	Expr    ExpressionNode
	Update  StatementNode
	Stats   []StatementNode
}

func NewForLoopNode(
	pos utils.Position,
	initial StatementNode,
	expr ExpressionNode,
	update StatementNode,
	stats []StatementNode,
) *ForLoopNode {
	return &ForLoopNode{
		Pos:     pos,
		Initial: initial,
		Expr:    expr,
		Update:  update,
		Stats:   stats,
	}
}

func (node ForLoopNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- FOR"))
	buf.WriteString(fmt.Sprintln("  - INITIAL"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Initial), "    "))
	buf.WriteString(fmt.Sprintln("  - CONDITION"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Expr), "    "))
	buf.WriteString(fmt.Sprintln("  - UPDATE"))
	buf.WriteString(utils.Indent(fmt.Sprintf("%s", node.Update), "    "))
	buf.WriteString(fmt.Sprintln("  - DO"))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(fmt.Sprintf("%s", s), "    "))
	}
	return buf.String()
}

/**************** SCOPE NODE ****************/

// ScopeNode stores the position and statement of a new scope.
//
// E.g.
//
//  begin skip end
type ScopeNode struct {
	Pos   utils.Position
	Stats []StatementNode
}

// NewScopeNode builds a ScopeNode
func NewScopeNode(pos utils.Position, stats []StatementNode) *ScopeNode {
	return &ScopeNode{
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
