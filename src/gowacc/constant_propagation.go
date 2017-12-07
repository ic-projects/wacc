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

func (v *Propagator) SetValue(node ast.RHSNode, identDec *ast.IdentifierDeclaration) {
	switch rhs := node.(type) {
	case *ast.BooleanLiteralNode,
		*ast.IntegerLiteralNode,
		*ast.CharacterLiteralNode:
		identDec.SetValue(rhs)
	case *ast.ArrayLiteralNode:
		flag := true
		for _, e := range rhs.Exprs {
			if _, ok := e.(*ast.IntegerLiteralNode); !ok {
				flag = false
				break
			}
		}
		if flag {
			identDec.SetValue(rhs)
		} else {
			identDec.RemoveValue()
		}
	default:
		identDec.RemoveValue()
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
		v.SetValue(node.RHS, identDec)
	case *ast.AssignNode:
		if ident, ok := node.LHS.(*ast.IdentifierNode); ok {
			identDec := v.symbolTable.SearchForDeclaredIdent(ident.Ident)
			v.SetValue(node.RHS, identDec)
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
	case *ast.ArrayElementNode:
		identDec := v.symbolTable.SearchForDeclaredIdent(t.Ident.Ident)
		if identDec.HasValue {
			cur := identDec.Value
			arrLiteral := identDec.Value.(*ast.ArrayLiteralNode)
			for _, e := range t.Exprs {
				if expr, ok := v.simulateFull(e); ok {
					index := expr.(*ast.IntegerLiteralNode).Val
					if index >= len(arrLiteral.Exprs) {
						// TODO Array out of bounds access
					}
					cur = arrLiteral.Exprs[index]
				} else {
					return node, false
				}
			}
			return cur, true
		}
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
