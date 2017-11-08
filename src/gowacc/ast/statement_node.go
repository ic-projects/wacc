package ast

import (
	"bytes"
	"fmt"
)

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
	ident IdentifierNode
	rhs   RHSNode
}

func NewDeclareNode(pos Position, t TypeNode, ident IdentifierNode, rhs RHSNode) DeclareNode {
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
	buf.WriteString(fmt.Sprintln("  - TYPE"))
	buf.WriteString(indent(fmt.Sprintf("- %s\n", node.t), "    "))
	buf.WriteString(fmt.Sprintln("  - LHS"))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.ident), "    "))
	buf.WriteString(fmt.Sprintln("  - RHS"))
	buf.WriteString(indent(fmt.Sprintf("%s\n", node.rhs), "    "))
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
	buf.WriteString(fmt.Sprintln("  - LHS"))
	buf.WriteString(indent(fmt.Sprintf("%s", node.lhs), "    "))
	buf.WriteString(fmt.Sprintln("  - RHS"))
	buf.WriteString(indent(fmt.Sprintf("%s", node.rhs), "    "))
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
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "  "))
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
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "  "))
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
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "  "))
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
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "  "))
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
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "  "))
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
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "  "))
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
	buf.WriteString(fmt.Sprintln("- IF"))
	buf.WriteString(indent(fmt.Sprintln("- CONDITION"), "  "))
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "    "))
	buf.WriteString(indent(fmt.Sprintln("- THEN"), "  "))
	for _, s := range node.ifStats {
		buf.WriteString(indent(fmt.Sprintf("%s", s), "    "))
	}
	buf.WriteString(indent(fmt.Sprintln("- ELSE"), "  "))
	for _, s := range node.elseStats {
		buf.WriteString(indent(fmt.Sprintf("%s", s), "    "))
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
	buf.WriteString(fmt.Sprintln("- LOOP"))
	buf.WriteString(fmt.Sprintln("  - CONDITION"))
	buf.WriteString(indent(fmt.Sprintf("%s", node.expr), "    "))
	buf.WriteString(fmt.Sprintln("  - DO"))
	for _, s := range node.stats {
		buf.WriteString(indent(fmt.Sprintf("%s", s), "    "))
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
		buf.WriteString(indent(fmt.Sprintf("%s", s), "  "))
	}
	return buf.String()
}
