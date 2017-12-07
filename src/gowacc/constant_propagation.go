package main

import (
	"ast"
	"utils"
)

// SimplifyTree will simplifiy a given tree, changing branches of the
// tree to their immediate values if they can be calculated.
func SimplifyTree(
	tree ast.ProgramNode,
	checker *SemanticCheck,
) *SemanticCheck {

	propagator := NewPropagator(checker)
	ast.Walk(propagator, tree)
	checker.symbolTable.Reset()
	return checker
}

// Propagator is the struct used when simplifying the tree, it links
// the propagator to the symbol table and error list.
type Propagator struct {
	symbolTable *ast.SymbolTable
	errors      *[]GenericError
}

// NewPropagator returns a initialised Propagator struct from a given
// semantic checker.
func NewPropagator(checker *SemanticCheck) *Propagator {
	return &Propagator{
		symbolTable: checker.symbolTable,
		errors:      checker.Errors,
	}
}

func (v *Propagator) Visit(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case *ast.FunctionNode:
		v.symbolTable.DestroyAllConstants()
		v.symbolTable.MoveNextScope()
	case *ast.ReadNode:
		if i, ok := node.LHS.(*ast.IdentifierNode); ok {
			v.symbolTable.DestroyConstant(i.Ident)
		} else if i, ok := node.LHS.(*ast.ArrayElementNode); ok {
			v.symbolTable.DestroyConstant(i.Ident.Ident)
		}
	case ast.Parameters:
		for _, e := range node {
			dec, _ := v.symbolTable.SearchForIdent(e.Ident.Ident)
			dec.IsDeclared = true
		}
	case *ast.LoopNode,
		*ast.ForLoopNode:
		v.symbolTable.DestroyAllConstants()
	case ast.Statements:
		v.symbolTable.DestroyAllConstants()
		v.symbolTable.MoveNextScope()
	case ast.ExpressionHolderNode:
		node.MapExpressions(v.simulate)
	}
}

// SetValue is used to set the internal value of an identifier in the
// symbol table if it detected to have a constant value.
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
			v.symbolTable.DestroyConstant(identDec.Ident.Ident)
		}
	default:
		v.symbolTable.DestroyConstant(identDec.Ident.Ident)
	}
}

func (v *Propagator) Leave(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case ast.Statements:
		v.symbolTable.DestroyAllConstants()
		v.symbolTable.MoveUpScope()
	case *ast.FunctionNode:
		v.symbolTable.DestroyAllConstants()
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
		} else if i, ok := node.LHS.(*ast.ArrayElementNode); ok {
			v.symbolTable.DestroyConstant(i.Ident.Ident)
		}
	}
}

// simulate returns the simplified version of a expression node branch.
func (v *Propagator) simulate(node ast.ProgramNode) ast.ExpressionNode {
	if result, ok := v.simulateFull(node); ok {
		return result
	}
	return node
}

// simulateFull returns the simplified version of a expression node branch and
// a boolean indicating if a prune or change occured.
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

					// Array index bounds checking
					if index >= int64(len(arrLiteral.Exprs)) {
						*v.errors = append(*v.errors,
							NewCustomStringError(t.Pos,
								ArrayIndexTooLarge,
								index,
								len(arrLiteral.Exprs)))
						return node, false
					} else if index < 0 {
						*v.errors = append(*v.errors,
							NewCustomStringError(t.Pos, ArrayIndexNegative, index))
						return node, false
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
			if _, ok := identDec.Value.(*ast.ArrayLiteralNode); !ok {
				return identDec.Value, true
			}
		}
	case *ast.BooleanLiteralNode,
		*ast.IntegerLiteralNode,
		*ast.CharacterLiteralNode:
		return t, true
	case *ast.BinaryOperatorNode:
		expr1, ok1 := v.simulateFull(t.Expr1)
		expr2, ok2 := v.simulateFull(t.Expr2)
		if ok1 && ok2 {
			return v.ApplyBinary(t, expr1, expr2)
		}
	case *ast.UnaryOperatorNode:
		if expr, ok := v.simulateFull(t.Expr); ok {
			return v.ApplyUnary(t, expr)
		}
	}
	return node, false
}

func (v *Propagator) ApplyUnary(
	un *ast.UnaryOperatorNode,
	e ast.ExpressionNode,
) (ast.ExpressionNode, bool) {
	switch un.Op {
	case ast.NOT:
		if b, ok := e.(*ast.BooleanLiteralNode); ok {
			return ast.NewBooleanLiteralNode(b.Pos, !b.Val), true
		}
	case ast.NEG:
		if b, ok := e.(*ast.IntegerLiteralNode); ok {
			if v.CheckOverflow(-b.Val, un.Pos) {
				return nil, false
			}
			return ast.NewIntegerLiteralNode(b.Pos, -b.Val), true
		}
	case ast.LEN:
		// TODO
	case ast.ORD:
		if b, ok := e.(*ast.CharacterLiteralNode); ok {
			return ast.NewIntegerLiteralNode(b.Pos, int64(b.Val)), true
		}
	case ast.CHR:
		if b, ok := e.(*ast.IntegerLiteralNode); ok {
			return ast.NewCharacterLiteralNode(b.Pos, rune(b.Val)), true
		}
	}
	return nil, false
}

func (v *Propagator) ApplyBinary(
	binOp *ast.BinaryOperatorNode,
	e1 ast.ExpressionNode,
	e2 ast.ExpressionNode,
) (ast.ExpressionNode, bool) {
	switch binOp.Op {
	case ast.OR:
		b1, ok1 := e1.(*ast.BooleanLiteralNode)
		b2, ok2 := e2.(*ast.BooleanLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val || b2.Val), true
		}
	case ast.AND:
		b1, ok1 := e1.(*ast.BooleanLiteralNode)
		b2, ok2 := e2.(*ast.BooleanLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val && b2.Val), true
		}
	case ast.MUL:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			if v.CheckOverflow(b1.Val*b2.Val, binOp.Pos) {
				return nil, false
			}
			return ast.NewIntegerLiteralNode(b1.Pos, b1.Val*b2.Val), true
		}
	case ast.DIV:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			if b2.Val == 0 {
				*v.errors = append(*v.errors, NewCustomStringError(binOp.Pos, DivideByZero))
				return nil, false
			}
			return ast.NewIntegerLiteralNode(b1.Pos, b1.Val/b2.Val), true
		}
	case ast.MOD:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			if b2.Val == 0 {
				*v.errors = append(*v.errors, NewCustomStringError(binOp.Pos, ModByZero))
				return nil, false
			}
			return ast.NewIntegerLiteralNode(b1.Pos, b1.Val%b2.Val), true
		}
	case ast.SUB:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			if v.CheckOverflow(b1.Val-b2.Val, binOp.Pos) {
				return nil, false
			}
			return ast.NewIntegerLiteralNode(b1.Pos, b1.Val-b2.Val), true
		}
	case ast.ADD:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			if v.CheckOverflow(b1.Val+b2.Val, binOp.Pos) {
				return nil, false
			}
			return ast.NewIntegerLiteralNode(b1.Pos, b1.Val+b2.Val), true
		}
	case ast.GEQ:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val >= b2.Val), true
		}
	case ast.GT:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val > b2.Val), true
		}
	case ast.LEQ:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val <= b2.Val), true
		}
	case ast.LT:
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val < b2.Val), true
		}
	case ast.EQ:
		//TODO make for all types
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val == b2.Val), true
		}
	case ast.NEQ:
		//TODO make for all types
		b1, ok1 := e1.(*ast.IntegerLiteralNode)
		b2, ok2 := e2.(*ast.IntegerLiteralNode)
		if ok1 && ok2 {
			return ast.NewBooleanLiteralNode(b1.Pos, b1.Val != b2.Val), true
		}
	}
	return nil, false
}

func (v *Propagator) CheckOverflow(val int64, pos utils.Position) bool {
	if int64(int32(val)) != val {
		*v.errors = append(*v.errors,
			NewCustomStringError(pos, OverFlow, val))
		return true
	}
	return false
}
