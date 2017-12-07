package main

import (
	"ast"
	"fmt"
)

func SimplifiyTree(
	tree ast.ProgramNode,
	symbolTable *ast.SymbolTable) {

	propagator := NewPropagator(symbolTable)
	ast.Walk(propagator, tree)
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
	case *ast.DeclareNode:
		if result, ok := v.simulate(node.RHS); ok {
			fmt.Println("simplified")
			fmt.Println(node.RHS)
			fmt.Println(result)
			node.RHS = result
		}
	case *ast.AssignNode:
		if result, ok := v.simulate(node.RHS); ok {
			fmt.Println("simplified")
			fmt.Println(node.RHS)
			fmt.Println(result)
			node.RHS = result
		}
	case *ast.ArrayElementNode:
		// TODO
	case *ast.FunctionCallNode:
		// TODO
	}
}

func (v *Propagator) simulate(node ast.ProgramNode) (ast.ExpressionNode, bool) {
	switch t := node.(type) {
	case *ast.IdentifierNode:
		// TODO
	case *ast.BooleanLiteralNode,
		*ast.IntegerLiteralNode,
		*ast.CharacterLiteralNode:
		return t, true
	case *ast.BinaryOperatorNode:
		expr1, ok1 := v.simulate(t.Expr1)
		expr2, ok2 := v.simulate(t.Expr2)
		if ok1 && ok2 {
			return t.Op.Apply(expr1, expr2)
		}
	case *ast.UnaryOperatorNode:
		if expr, ok := v.simulate(t.Expr); ok {
			return t.Op.Apply(expr)
		}
	}
	return node, false
}
