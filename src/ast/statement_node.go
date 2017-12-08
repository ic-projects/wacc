package ast

import (
	"bytes"
	"fmt"
	"utils"
)

// StatementNode is an interface for statement nodes to implement.
type StatementNode interface {
	ProgramNode
}

/**************** STATEMENT NODE SLICE ****************/

// Statements is a slice of StatementNodes. We need this type so that visitors
// can easily tell when to change scope.
type Statements []StatementNode

func (stats Statements) String() string {
	var buf bytes.Buffer
	for _, s := range stats {
		buf.WriteString(s.String())
	}
	return buf.String()
}

func (stats Statements) walkNode(visitor Visitor) {
	for _, s := range stats {
		Walk(visitor, s)
	}
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
	case *SwitchNode:
		for _, c := range s.(*SwitchNode).Cases {
			if !FinalStatIsValid(c.Stats[len(c.Stats)-1]) {
				return false
			}
		}
		return true
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

func (node *SkipNode) String() string {
	return "- SKIP\n"
}

func (node *SkipNode) walkNode(visitor Visitor) {
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

func (node *DeclareNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- DECLARE"))
	buf.WriteString(fmt.Sprintln("  - TYPE"))
	buf.WriteString(utils.Indent(fmt.Sprintf("- %s", node.T), "    "))
	buf.WriteString(fmt.Sprintln("  - LHS"))
	buf.WriteString(utils.Indent(fmt.Sprintln(node.Ident), "    "))
	buf.WriteString(fmt.Sprintln("  - RHS"))
	buf.WriteString(utils.Indent(fmt.Sprintln(node.RHS), "    "))
	return buf.String()
}

func (node *DeclareNode) MapExpressions(m Mapper) {
	node.RHS = m(node.RHS)
}

func (node *DeclareNode) walkNode(visitor Visitor) {
	Walk(visitor, node.RHS)
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

func (node *AssignNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- ASSIGNMENT"))
	buf.WriteString(fmt.Sprintln("  - LHS"))
	buf.WriteString(utils.Indent(node.LHS.String(), "    "))
	buf.WriteString(fmt.Sprintln("  - RHS"))
	buf.WriteString(utils.Indent(node.RHS.String(), "    "))
	return buf.String()
}

func (node *AssignNode) MapExpressions(m Mapper) {
	node.RHS = m(node.RHS)
}

func (node *AssignNode) walkNode(visitor Visitor) {
	Walk(visitor, node.LHS)
	Walk(visitor, node.RHS)
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

func (node *ReadNode) String() string {
	return writeSimpleString("READ", node.LHS)
}

func (node *ReadNode) MapExpressions(m Mapper) {
}

func (node *ReadNode) walkNode(visitor Visitor) {
	Walk(visitor, node.LHS)
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

func (node *FreeNode) String() string {
	return writeSimpleString("FREE", node.Expr)
}

func (node *FreeNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *FreeNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
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

func (node *ReturnNode) String() string {
	return writeSimpleString("RETURN", node.Expr)
}

func (node *ReturnNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *ReturnNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
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

func (node *ExitNode) String() string {
	return writeSimpleString("EXIT", node.Expr)
}

func (node *ExitNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *ExitNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
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

func (node *PrintNode) String() string {
	return writeSimpleString("PRINT", node.Expr)
}

func (node *PrintNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *PrintNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
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

func (node *PrintlnNode) String() string {
	return writeSimpleString("PRINTLN", node.Expr)
}

func (node *PrintlnNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *PrintlnNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
}

/**************** SWITCH NODE ****************/

// SwitchNode stores the position, condition and the branches of a switch
// statement.
//
// E.g.
//
//  when i case 0, 1: skip end else: skip end fi
type SwitchNode struct {
	Pos   utils.Position
	Expr  ExpressionNode
	Cases []CaseNode
}

// NewSwitchNode builds a SwitchNode.
func NewSwitchNode(
	pos utils.Position,
	expr ExpressionNode,
	cases []CaseNode,
) *SwitchNode {
	return &SwitchNode{
		Pos:   pos,
		Expr:  expr,
		Cases: cases,
	}
}

func (node *SwitchNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- SWITCH"))
	buf.WriteString(utils.Indent(fmt.Sprintln("- EXPRESSION"), "  "))
	buf.WriteString(utils.Indent(node.Expr.String(), "    "))
	for _, s := range node.Cases {
		buf.WriteString(utils.Indent(s.String(), "    "))
	}
	return buf.String()
}

func (node *SwitchNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *SwitchNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
	for _, c := range node.Cases {
		if !c.IsDefault {
			for _, e := range c.Exprs {
				Walk(visitor, e)
			}
		}
		Walk(visitor, c.Stats)
	}
}

/**************** CASE NODE ****************/

// CaseNode is a struct that holds the case and the statements that should be
// executed.
type CaseNode struct {
	Pos       utils.Position
	Exprs     []ExpressionNode
	Stats     Statements
	IsDefault bool
}

// NewDefaultCaseNode builds a CaseNode for the default case.
func NewDefaultCaseNode(pos utils.Position, stats []StatementNode) CaseNode {
	return CaseNode{
		Pos:       pos,
		Stats:     stats,
		IsDefault: true,
	}
}

// NewCaseNode builds a CaseNode for a non-default case.
func NewCaseNode(
	pos utils.Position,
	exprs []ExpressionNode,
	stats []StatementNode,
) CaseNode {
	return CaseNode{
		Pos:       pos,
		Exprs:     exprs,
		Stats:     stats,
		IsDefault: false,
	}
}

func (node *CaseNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- CASE"))
	if node.IsDefault {
		buf.WriteString(utils.Indent(fmt.Sprintln("- DEFAULT"), "  "))
	} else {
		for _, e := range node.Exprs {
			buf.WriteString(utils.Indent(fmt.Sprintln("- EXPRESSION"), "  "))
			buf.WriteString(utils.Indent(e.String(), "    "))
		}
	}
	buf.WriteString(utils.Indent(fmt.Sprintln("- THEN"), "  "))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(s.String(), "    "))
	}
	return buf.String()
}

func (node *CaseNode) MapExpressions(m Mapper) {
	for _, e := range node.Exprs {
		e = m(e)
	}
}

func (node *CaseNode) walkNode(visitor Visitor) {
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
	IfStats   Statements
	ElseStats Statements
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

func (node *IfNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- IF"))
	buf.WriteString(utils.Indent(fmt.Sprintln("- CONDITION"), "  "))
	buf.WriteString(utils.Indent(node.Expr.String(), "    "))
	buf.WriteString(utils.Indent(fmt.Sprintln("- THEN"), "  "))
	for _, s := range node.IfStats {
		buf.WriteString(utils.Indent(s.String(), "    "))
	}
	buf.WriteString(utils.Indent(fmt.Sprintln("- ELSE"), "  "))
	for _, s := range node.ElseStats {
		buf.WriteString(utils.Indent(s.String(), "    "))
	}
	return buf.String()
}

func (node *IfNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *IfNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
	Walk(visitor, node.IfStats)
	Walk(visitor, node.ElseStats)
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
	Stats Statements
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

func (node *LoopNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- LOOP"))
	buf.WriteString(fmt.Sprintln("  - CONDITION"))
	buf.WriteString(utils.Indent(node.Expr.String(), "    "))
	buf.WriteString(fmt.Sprintln("  - DO"))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(s.String(), "    "))
	}
	return buf.String()
}

func (node *LoopNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *LoopNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Expr)
	Walk(visitor, node.Stats)
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
	Stats   Statements
}

// NewForLoopNode builds a ForLoopNode.
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

func (node *ForLoopNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- FOR"))
	buf.WriteString(fmt.Sprintln("  - INITIAL"))
	buf.WriteString(utils.Indent(node.Initial.String(), "    "))
	buf.WriteString(fmt.Sprintln("  - CONDITION"))
	buf.WriteString(utils.Indent(node.Expr.String(), "    "))
	buf.WriteString(fmt.Sprintln("  - UPDATE"))
	buf.WriteString(utils.Indent(node.Update.String(), "    "))
	buf.WriteString(fmt.Sprintln("  - DO"))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(s.String(), "    "))
	}
	return buf.String()
}

func (node *ForLoopNode) MapExpressions(m Mapper) {
	node.Expr = m(node.Expr)
}

func (node *ForLoopNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Initial)
	Walk(visitor, node.Expr)
	Walk(visitor, node.Update)
	Walk(visitor, node.Stats)
}

/**************** SCOPE NODE ****************/

// ScopeNode stores the position and statement of a new scope.
//
// E.g.
//
//  begin skip end
type ScopeNode struct {
	Pos   utils.Position
	Stats Statements
}

// NewScopeNode builds a ScopeNode
func NewScopeNode(pos utils.Position, stats []StatementNode) *ScopeNode {
	return &ScopeNode{
		Pos:   pos,
		Stats: stats,
	}
}

func (node *ScopeNode) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintln("- SCOPE"))
	for _, s := range node.Stats {
		buf.WriteString(utils.Indent(s.String(), "  "))
	}
	return buf.String()
}

func (node *ScopeNode) MapExpressions(m Mapper) {
}

func (node *ScopeNode) walkNode(visitor Visitor) {
	Walk(visitor, node.Stats)
}
