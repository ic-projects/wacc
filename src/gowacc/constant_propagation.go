package main

import (
	"ast"
)

func SimplifiyTree(
	tree ast.ProgramNode,
	symbolTable *ast.SymbolTable) {

	propagator := NewPropagator(symbolTable)
	ast.Walk(propagator, tree)
	symbolTable.Reset()
}

type Propagator struct {
	symbolTable *ast.SymbolTable
}

func NewPropagator(symbolTable *ast.SymbolTable) *Propagator {
	return &Propagator{
		symbolTable: symbolTable,
	}
}

func (v *Propagator) Visit(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case *ast.FunctionNode:
		v.symbolTable.MoveNextScope()
	case ast.Parameters:
		for _, e := range node {
			dec, _ := v.symbolTable.SearchForIdent(e.Ident.Ident)
			dec.IsDeclared = true
		}
	case ast.ExpressionHolderNode:
		node.MapExpressions(v.simulate)
	case ast.Statements:
		v.symbolTable.MoveNextScope()

	}
}

func (v *Propagator) Leave(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case ast.Statements:
		v.symbolTable.MoveUpScope()
	case *ast.FunctionNode:
		v.symbolTable.MoveUpScope()
	case *ast.DeclareNode:
		dec, _ := v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident)
		dec.IsDeclared = true
		identDec := v.symbolTable.SearchForDeclaredIdent(node.Ident.Ident)
		switch rhs := node.RHS.(type) {
		case *ast.BooleanLiteralNode,
			*ast.IntegerLiteralNode,
			*ast.CharacterLiteralNode:
			identDec.SetValue(rhs)
		default:
			identDec.RemoveValue()
		}
	case *ast.AssignNode:
		if ident, ok := node.LHS.(*ast.IdentifierNode); ok {
			identDec := v.symbolTable.SearchForDeclaredIdent(ident.Ident)
			switch rhs := node.RHS.(type) {
			case *ast.BooleanLiteralNode,
				*ast.IntegerLiteralNode,
				*ast.CharacterLiteralNode:
				identDec.SetValue(rhs)
			default:
				identDec.RemoveValue()
			}
		}
	}
}

func (v *Propagator) simulate(node ast.ProgramNode) ast.ExpressionNode {
	if result, ok := v.simulateFull(node); ok {
		return result
	}
	return node
}

func (v *Propagator) simulateFull(node ast.ProgramNode) (ast.ExpressionNode, bool) {
	switch t := node.(type) {
	case *ast.IdentifierNode:
		identDec := v.symbolTable.SearchForDeclaredIdent(t.Ident)
		if identDec.HasValue {
			return identDec.Value, true
		}
	case *ast.BooleanLiteralNode,
		*ast.IntegerLiteralNode,
		*ast.CharacterLiteralNode:
		return t, true
	case *ast.BinaryOperatorNode:
		expr1, ok1 := v.simulateFull(t.Expr1)
		expr2, ok2 := v.simulateFull(t.Expr2)
		if ok1 && ok2 {
			return t.Op.Apply(expr1, expr2)
		}
	case *ast.UnaryOperatorNode:
		if expr, ok := v.simulateFull(t.Expr); ok {
			return t.Op.Apply(expr)
		}
	}
	return node, false
}
