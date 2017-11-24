package semantics

import (
	"fmt"
	"ast"
)

// SemanticCheck is a struct that implements EntryExitVisitor to be called with
// Walk. It stores a SymbolTable, a TypeChecker, and a list of GenericErrors.
type SemanticCheck struct {
	symbolTable *ast.SymbolTable
	typeChecker *TypeChecker
	Errors      []GenericError
}

// NewSemanticCheck returns an initialised SemanticCheck
func NewSemanticCheck() *SemanticCheck {
	return &SemanticCheck{
		symbolTable: ast.NewSymbolTable(),
		typeChecker: NewTypeChecker(),
		Errors:      make([]GenericError, 0),
	}
}

func (v *SemanticCheck) PrintSymbolTable() {
	fmt.Println(v.symbolTable.String())
}

func (v *SemanticCheck) SymbolTable() *ast.SymbolTable {
	return v.symbolTable
}

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *SemanticCheck) Visit(programNode ast.ProgramNode) {
	var foundError GenericError
	foundError = nil
	switch node := programNode.(type) {
	case ast.Program:
		// Add the Functions when hitting program instead of each function so that
		// Functions can be declared in any order.
		for _, f := range node.Functions {
			if functionNode, ok := v.symbolTable.SearchForFunction(f.Ident.Ident); ok {
				foundError = NewPreviouslyDeclared(NewDeclarationError(f.Pos, true, true, f.Ident.Ident), functionNode.Pos)
			} else {
				v.symbolTable.AddFunction(f.Ident.Ident, f)
			}
		}
	case ast.FunctionNode:
		// Move down Scope so that the parameters are on a new Scope.
		v.symbolTable.MoveDownScope()
		v.typeChecker.expectRepeatUntilForce(node.T)
	case ast.ParameterNode:
		if declareNode, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); ok {
			foundError = NewPreviouslyDeclared(NewDeclarationError(node.Pos, false, true, node.Ident.Ident), declareNode.Pos)
		}
	case ast.SkipNode:
	case ast.DeclareNode:
		if declareNode, ok := v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident); ok {
			foundError = NewPreviouslyDeclared(NewDeclarationError(node.Pos, false, true, node.Ident.Ident), declareNode.Pos)
			v.typeChecker.freeze(node)
		} else {
			v.typeChecker.expect(node.T)
		}
	case ast.AssignNode:
		v.typeChecker.expectTwiceSame(NewAnyExpectance())
	case ast.ReadNode:
		v.typeChecker.expectSet([]ast.TypeNode{ast.NewBaseTypeNode(ast.INT), ast.NewBaseTypeNode(ast.CHAR)})
	case ast.FreeNode:
		v.typeChecker.expectSet([]ast.TypeNode{ast.PairTypeNode{}, ast.ArrayTypeNode{}})
	case ast.ReturnNode:
	case ast.ExitNode:
		v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
	case ast.PrintNode:
		v.typeChecker.expectAny()
	case ast.PrintlnNode:
		v.typeChecker.expectAny()
	case ast.IfNode:
		v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
	case ast.LoopNode:
		v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
	case ast.ScopeNode:
	case ast.IdentifierNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.Ident); !ok {
			foundError = NewDeclarationError(node.Pos, false, false, node.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(identDec.T).addPos(node.Pos)
			if foundError != nil {
				foundError = NewTypeErrorDeclaration(foundError.(TypeError), identDec.Pos)
			}
		}
	case ast.PairFirstElementNode:
		//  Look up type for pair call seen
		if identNode, ok := node.Expr.(ast.IdentifierNode); !ok {
			foundError = NewCustomError(node.Pos, "Cannot access first element of null")
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.Ident); !ok {
			foundError = NewDeclarationError(identNode.Pos, false, false, identNode.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(identDec.T.(ast.PairTypeNode).T1).addPos(node.Pos)
			v.typeChecker.expect(identDec.T)
		}
	case ast.PairSecondElementNode:
		if identNode, ok := node.Expr.(ast.IdentifierNode); !ok {
			foundError = NewCustomError(node.Pos, "Cannot access second element of null")
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.Ident); !ok {
			foundError = NewDeclarationError(identNode.Pos, false, false, identNode.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			v.typeChecker.seen(identDec.T.(ast.PairTypeNode).T2)
			v.typeChecker.expect(identDec.T)
		}
	case ast.ArrayElementNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); !ok {
			foundError = NewDeclarationError(node.Pos, false, false, node.Ident.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if arrayNode, ok := identDec.T.(ast.ArrayTypeNode); !ok {
			foundError = NewCustomError(node.Pos, fmt.Sprintf("Array access on non-array variable \"%s\"", node.Ident.Ident))
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			// If we have an array or a single element (for use in newsted arrays).
			if dimLeft := arrayNode.Dim - len(node.Exprs); dimLeft == 0 {
				foundError = v.typeChecker.seen(arrayNode.T).addPos(node.Pos)
			} else {
				foundError = v.typeChecker.seen(ast.NewArrayTypeNode(arrayNode.T, dimLeft)).addPos(node.Pos)
			}
		}
		for i := 0; i < len(node.Exprs); i++ {
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		}
	case ast.ArrayLiteralNode:
		foundError = v.typeChecker.seen(ast.ArrayTypeNode{}).addPos(node.Pos)
	case ast.NewPairNode:
		foundError = v.typeChecker.seen(ast.PairTypeNode{}).addPos(node.Pos)
	case ast.FunctionCallNode:
		if f, ok := v.symbolTable.SearchForFunction(node.Ident.Ident); !ok {
			foundError = NewDeclarationError(node.Pos, true, false, node.Ident.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if len(f.Params) != len(node.Exprs) {
			foundError = NewCustomError(node.Pos, fmt.Sprintf("Incorrect number of parameters for function \"%s\" (Expected: %d, Given: %d)", node.Ident.Ident, len(f.Params), len(node.Exprs)))
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(f.T).addPos(node.Pos)
			for i := len(f.Params) - 1; i >= 0; i-- {
				v.typeChecker.expect(f.Params[i].T)
			}
		}
	case ast.BaseTypeNode:
	case ast.ArrayTypeNode:
	case ast.PairTypeNode:
	case ast.UnaryOperator:
	case ast.BinaryOperator:
	case ast.IntegerLiteralNode:
		foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.INT)).addPos(node.Pos)
	case ast.BooleanLiteralNode:
		foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.BOOL)).addPos(node.Pos)
	case ast.CharacterLiteralNode:
		foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.CHAR)).addPos(node.Pos)
	case ast.StringLiteralNode:
		foundError = v.typeChecker.seen(ast.NewStringArrayTypeNode()).addPos(node.Pos)
	case ast.PairLiteralNode:
		foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.PAIR)).addPos(node.Pos)
	case ast.UnaryOperatorNode:
		switch node.Op {
		case ast.NOT:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.BOOL)).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
		case ast.NEG:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.INT)).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		case ast.LEN:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.INT)).addPos(node.Pos)
			v.typeChecker.expect(ast.ArrayTypeNode{})
		case ast.ORD:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.INT)).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.CHAR))
		case ast.CHR:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.CHAR)).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		}
	case ast.BinaryOperatorNode:
		switch node.Op {
		case ast.MUL, ast.DIV, ast.MOD, ast.ADD, ast.SUB:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.INT)).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		case ast.GT, ast.GEQ, ast.LT, ast.LEQ:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.BOOL)).addPos(node.Pos)
			v.typeChecker.expectTwiceSame(NewSetExpectance([]ast.TypeNode{ast.NewBaseTypeNode(ast.INT), ast.NewBaseTypeNode(ast.CHAR)}))
		case ast.EQ, ast.NEQ:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.BOOL)).addPos(node.Pos)
			v.typeChecker.expectTwiceSame(NewAnyExpectance())
		case ast.AND, ast.OR:
			foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.BOOL)).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
		}
	case []ast.StatementNode:
		v.symbolTable.MoveDownScope()
	}

	// If we have an error, add it to the list of errors.
	if foundError != nil {
		v.Errors = append(v.Errors, foundError)
	}
}

// Leave will be called to leave the current node.
func (v *SemanticCheck) Leave(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case []ast.StatementNode:
		v.symbolTable.MoveUpScope()
	case ast.FunctionNode:
		v.symbolTable.MoveUpScope()
		v.typeChecker.forcePop()
	case ast.ArrayLiteralNode:
		v.typeChecker.forcePop()
	case ast.DeclareNode:
		if _, ok := v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident); !ok {
			v.symbolTable.AddToScope(node.Ident.Ident, node)
		}
	case ast.ParameterNode:
		if _, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); !ok {
			v.symbolTable.AddToScope(node.Ident.Ident, node)
		}
	}
	v.typeChecker.unfreeze(programNode)
}

func (v *SemanticCheck) PrintErrors(filepath string) {
	PrintErrors(v.Errors, filepath)
}
