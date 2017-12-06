package main

import (
	"ast"
	"fmt"
)

func validType(T ast.TypeNode, i *ast.IdentifierNode) GenericError {
	switch t := ast.ToValue(T).(type) {
	case ast.ArrayTypeNode:
		if validType(t.T, i) != nil {
			return NewCustomError(i.Pos, fmt.Sprint("Unknown dynamic type for ident %s, has type %s", i.Ident, t))
		}
		return nil
	case ast.DynamicTypeNode:
		if len(t.T.Poss) > 1 {
			return NewCustomError(i.Pos, fmt.Sprint("Ambiguous dynamic type for ident %s, could be of types %s", i.Ident, t.T.Poss))
		}
		if len(t.T.Poss) != 1 || validType(t.T.Poss[0], i) != nil {
			return NewCustomError(i.Pos, fmt.Sprint("Unknown dynamic type for ident %s, has type %s", i.Ident, t))
		}
		return nil
	case ast.PairTypeNode:
		if validType(t.T1, i) != nil || validType(t.T2, i) != nil {
			return NewCustomError(i.Pos, fmt.Sprint("Unknown dynamic type for ident %s, has type %s", i.Ident, t))
		}
		return nil
	default:
		return nil
	}
}

func (v *SemanticCheck) checkForDynamicErrors(e *[]GenericError) bool {
	for _, f := range v.symbolTable.Structs {
		for _, t := range f.Types {
			err := validType(t.T, t.Ident)
			if err != nil {
				*e = append(*e, err)
			}
		}
	}

	for _, f := range v.symbolTable.Functions {
		for _, t := range f.Params {
			err := validType(t.T, t.Ident)
			if err != nil {
				*e = append(*e, err)
			}
		}
	}

	for _, s := range v.symbolTable.Head.ChildScopes {
		err := checkForValidTypes(s)
		if err != nil {
			*e = append(*e, err...)
		}
	}

	return len(*e) > 0
}

func checkForValidTypes(node *ast.SymbolTableNode) []GenericError {
	e := make([]GenericError, 0)
	for _, ident := range node.Scope {
		err := validType(ident.T, ident.Ident)
		if err != nil {
			e = append(e, err)
		}
	}

	for _, s := range node.ChildScopes {
		err := checkForValidTypes(s)
		if err != nil {
			e = append(e, err...)
		}
	}

	return e
}

/**************** SEMANTIC CHECK ****************/

// SemanticCheck is a struct that implements EntryExitVisitor to be called with
// Walk. It stores a SymbolTable, a TypeChecker, and a list of GenericErrors.
type SemanticCheck struct {
	symbolTable *ast.SymbolTable
	typeTable   map[string]*ast.StructNode
	typeChecker *TypeChecker
	Errors      []GenericError
}

// NewSemanticCheck returns an initialised SemanticCheck
func NewSemanticCheck() *SemanticCheck {
	return &SemanticCheck{
		symbolTable: ast.NewSymbolTable(),
		typeTable:   make(map[string]*ast.StructNode),
		typeChecker: NewTypeChecker(),
		Errors:      make([]GenericError, 0),
	}
}

/**************** GETTER / PRINTING FUNCTIONS ****************/

// SymbolTable returns a semantic check's symbol table object
func (v *SemanticCheck) SymbolTable() *ast.SymbolTable {
	return v.symbolTable
}

// PrintSymbolTable prints the string representation for a symbol table
func (v *SemanticCheck) PrintSymbolTable() {
	fmt.Println(v.symbolTable.String())
}

// PrintErrors pretty prints semantic errors using the PrintErrors function in
// semantic_error.go
func (v *SemanticCheck) PrintErrors(filepath string) {
	PrintErrors(v.Errors, filepath)
}

func (v *SemanticCheck) hasErrors() bool {
	if len(v.Errors) > 0 {
		return true
	}
	return v.checkForDynamicErrors(&v.Errors)
}

// Visit will apply the correct rule for the programNode given, to be used with
// ast.Walk.
func (v *SemanticCheck) Visit(programNode ast.ProgramNode) {
	var foundError GenericError
	switch node := programNode.(type) {
	case *ast.Program:
		for _, s := range node.Structs {
			if _, ok := v.symbolTable.SearchForStruct(s.Ident.Ident); ok {
				foundError = NewDeclarationError(
					s.Pos,
					false,
					true,
					s.Ident.Ident,
				)
			} else {
				v.symbolTable.AddStruct(s.Ident.Ident, s)
			}
		}
		// Add the Functions when hitting program instead of each function so
		// that Functions can be declared in any order.
		for _, f := range node.Functions {
			if functionNode, ok :=
				v.symbolTable.SearchForFunction(f.Ident.Ident); ok {
				foundError = NewPreviouslyDeclared(NewDeclarationError(
					f.Pos,
					true,
					true,
					f.Ident.Ident,
				), functionNode.Pos)
			} else {
				v.symbolTable.AddFunction(f.Ident.Ident, f)
			}
		}
	case *ast.FunctionNode:
		// Move down Scope so that the parameters are on a new Scope.
		v.symbolTable.MoveDownScope()
		v.typeChecker.expectRepeatUntilForce(node.T)
	case *ast.StructNode:
		for _, t := range node.Types {
			if s, ok := t.T.(*ast.StructTypeNode); ok {
				if _, ok := v.symbolTable.SearchForStruct(s.Ident); !ok {
					foundError = NewDeclarationError(
						node.Pos,
						false,
						false,
						s.Ident,
					)
				}
			}
		}
	case *ast.ParameterNode:
		if declareNode, ok :=
			v.symbolTable.SearchForIdent(node.Ident.Ident); ok {
			foundError = NewPreviouslyDeclared(NewDeclarationError(
				node.Pos,
				false,
				true,
				node.Ident.Ident,
			), declareNode.Pos)
		}
	case *ast.SkipNode:
	case *ast.DeclareNode:
		if declareNode, ok := v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident); ok {
			if _, ok := node.T.(*ast.DynamicTypeNode); !ok {
				foundError = NewPreviouslyDeclared(
					NewDeclarationError(node.Pos, false, true, node.Ident.Ident),
					declareNode.Pos)
				v.typeChecker.freeze(node)
			} else {
				v.typeChecker.expect(declareNode.T)
			}
		} else {
			v.typeChecker.expect(node.T)
		}
	case *ast.PointerNewNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); !ok {
			foundError = NewDeclarationError(node.Pos, false, false, node.Ident.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(ast.NewPointerTypeNode(identDec.T)).addPos(node.Pos)
			if foundError != nil {
				foundError = NewTypeErrorDeclaration(foundError.(TypeError), identDec.Pos)
			}
		}
	case *ast.AssignNode:
		v.typeChecker.expectTwiceSame(NewAnyExpectance())
	case *ast.ReadNode:
		v.typeChecker.expectSet([]ast.TypeNode{
			ast.NewBaseTypeNode(ast.INT),
			ast.NewBaseTypeNode(ast.CHAR),
		})
	case *ast.FreeNode:
		v.typeChecker.expectSet([]ast.TypeNode{&ast.PairTypeNode{}, &ast.ArrayTypeNode{}})
	case *ast.ReturnNode:
	case *ast.ExitNode:
		v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
	case *ast.PrintNode:
		v.typeChecker.expectAny()
	case *ast.PrintlnNode:
		v.typeChecker.expectAny()
	case *ast.IfNode:
		v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
	case *ast.SwitchNode:
		v.typeChecker.expectRepeatUntilForce(ast.NewBaseTypeNode(ast.INT))
	case *ast.LoopNode:
		v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
	case *ast.ForLoopNode:
		v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
		v.symbolTable.MoveDownScope()
	case *ast.ScopeNode:
	case *ast.IdentifierNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.Ident); !ok {
			foundError = NewDeclarationError(node.Pos, false, false, node.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(identDec.T).addPos(node.Pos)
			if foundError != nil {
				foundError = NewTypeErrorDeclaration(
					foundError.(TypeError),
					identDec.Pos,
				)
			}
		}
	case *ast.PairFirstElementNode:
		//  Look up type for pair call seen
		if identNode, ok := node.Expr.(*ast.IdentifierNode); !ok {
			foundError = NewCustomError(
				node.Pos,
				"Cannot access first element of null",
			)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok :=
			v.symbolTable.SearchForIdent(identNode.Ident); !ok {
			foundError = NewDeclarationError(
				identNode.Pos,
				false,
				false,
				identNode.Ident,
			)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			if dyn, ok := ast.ToValue(identDec.T).(*ast.DynamicTypeNode); ok {
				dyn.ReduceSet([]ast.TypeNode{ast.NewPairTypeNode(
					ast.NewDynamicTypeInsidePairNode(),
					ast.NewDynamicTypeInsidePairNode())})
			}
			foundError = v.typeChecker.seen(ast.ToValue(identDec.T).(ast.PairTypeNode).T1).addPos(node.Pos)
			v.typeChecker.expect(identDec.T)
		}
	case *ast.PairSecondElementNode:
		if identNode, ok := node.Expr.(*ast.IdentifierNode); !ok {
			foundError = NewCustomError(
				node.Pos,
				"Cannot access second element of null",
			)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok :=
			v.symbolTable.SearchForIdent(identNode.Ident); !ok {
			foundError = NewDeclarationError(
				identNode.Pos,
				false,
				false,
				identNode.Ident,
			)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			if dyn, ok := ast.ToValue(identDec.T).(*ast.DynamicTypeNode); ok {
				dyn.ReduceSet([]ast.TypeNode{ast.NewPairTypeNode(
					ast.NewDynamicTypeInsidePairNode(),
					ast.NewDynamicTypeInsidePairNode())})
			}
			v.typeChecker.seen(ast.ToValue(identDec.T).(ast.PairTypeNode).T2)
			v.typeChecker.expect(identDec.T)
		}
	case *ast.StructElementNode:
		if id, ok := v.symbolTable.SearchForIdent(node.Struct.Ident); !ok {
			foundError = NewDeclarationError(
				node.Pos,
				false,
				false,
				node.Struct.Ident,
			)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			if dyn, ok := ast.ToValue(id.T).(*ast.DynamicTypeNode); ok {
				poss := v.symbolTable.SearchForStructByUsage(node.Struct.Ident)
				dyn.ReduceSet(poss)
				v.typeChecker.seen(nil)
			} else {
				if struc, ok := ast.ToValue(id.T).(ast.StructTypeNode); !ok {
					foundError = NewCustomError(node.Pos, fmt.Sprintf("Struct access on non-struct variable \"%s\"", node.Ident.Ident))
					v.typeChecker.seen(nil)
					v.typeChecker.freeze(node)
				} else if structNode, ok := v.symbolTable.SearchForStruct(struc.Ident); !ok {
					foundError = NewCustomError(node.Pos, fmt.Sprintf("Struct access on non-struct variable \"%s\"", node.Ident.Ident))
					v.typeChecker.seen(nil)
					v.typeChecker.freeze(node)
				} else if found, ok := structNode.TypesMap[node.Ident.Ident]; !ok {
					foundError = NewDeclarationError(node.Pos, false, false, node.Ident.Ident)
					v.typeChecker.seen(nil)
					v.typeChecker.freeze(node)
				} else {
					node.SetStructType(structNode.Types[found])
					foundError = v.typeChecker.seen(structNode.Types[found].T).addPos(node.Pos)
				}
			}
		}
	case *ast.ArrayElementNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); !ok {
			foundError = NewDeclarationError(
				node.Pos,
				false,
				false,
				node.Ident.Ident,
			)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			if dyn, ok := ast.ToValue(identDec.T).(*ast.DynamicTypeNode); ok {
				dyn.ReduceSet([]ast.TypeNode{ast.NewArrayTypeDimNode(ast.NewDynamicTypeNode(), len(node.Exprs))})
			}

			if arrayNode, ok := ast.ToValue(identDec.T).(ast.ArrayTypeNode); !ok {
				foundError = NewCustomError(node.Pos, fmt.Sprintf("Array access on non-array variable \"%s\" of type %s", node.Ident.Ident, identDec.T))
				v.typeChecker.seen(nil)
				v.typeChecker.freeze(node)
			} else {
				// If we have an array or a single element (for use in newsted arrays).
				foundError = v.typeChecker.seen(arrayNode.GetDimElement(len(node.Exprs))).addPos(node.Pos)
			}
		}
		for i := 0; i < len(node.Exprs); i++ {
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		}

	case *ast.ArrayLiteralNode:
		foundError = v.typeChecker.seen(&ast.ArrayTypeNode{}).addPos(node.Pos)
	case *ast.NewPairNode:
		foundError = v.typeChecker.seen(&ast.PairTypeNode{}).addPos(node.Pos)
	case *ast.StructNewNode:
		foundError = v.typeChecker.seen(node.T).addPos(node.Pos)
		if structNode, ok := v.symbolTable.SearchForStruct(node.T.Ident); !ok {
			foundError = NewCustomError(node.Pos, fmt.Sprintf(
				"Struct init on non-struct \"%s\"",
				node.T,
			))
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if len(structNode.Types) != len(node.Exprs) {
			foundError = NewCustomError(node.Pos, fmt.Sprintf(
				"Incorrect number of parameters for struct \"%s\" "+
					"(Expected: %d, Given: %d)",
				structNode.Ident.Ident,
				len(structNode.Types),
				len(node.Exprs),
			))
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			node.SetStructType(structNode)
			for i := len(structNode.Types) - 1; i >= 0; i-- {
				v.typeChecker.expect(structNode.Types[i].T)
			}
		}
	case *ast.FunctionCallNode:
		if f, ok := v.symbolTable.SearchForFunction(node.Ident.Ident); !ok {
			foundError = NewDeclarationError(
				node.Pos,
				true,
				false,
				node.Ident.Ident,
			)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if len(f.Params) != len(node.Exprs) {
			foundError = NewCustomError(node.Pos, fmt.Sprintf(
				"Incorrect number of parameters for function \"%s\" "+
					"(Expected: %d, Given: %d)",
				node.Ident.Ident,
				len(f.Params),
				len(node.Exprs),
			))
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
	case *ast.UnaryOperator:
	case *ast.BinaryOperator:
	case *ast.IntegerLiteralNode:
		foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.INT)).addPos(node.Pos)
	case *ast.BooleanLiteralNode:
		foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.BOOL)).addPos(node.Pos)
	case *ast.CharacterLiteralNode:
		foundError = v.typeChecker.seen(ast.NewBaseTypeNode(ast.CHAR)).addPos(node.Pos)
	case *ast.StringLiteralNode:
		foundError = v.typeChecker.seen(
			ast.NewStringArrayTypeNode(),
		).addPos(node.Pos)
	case *ast.NullNode:
		foundError = v.typeChecker.seen(ast.NewNullTypeNode()).addPos(node.Pos)
	case *ast.UnaryOperatorNode:
		switch node.Op {
		case ast.NOT:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.BOOL),
			).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.BOOL))
		case ast.NEG:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.INT),
			).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		case ast.LEN:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.INT),
			).addPos(node.Pos)
			v.typeChecker.expect(&ast.ArrayTypeNode{})
		case ast.ORD:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.INT),
			).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.CHAR))
		case ast.CHR:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.CHAR),
			).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		}
	case *ast.BinaryOperatorNode:
		switch node.Op {
		case ast.MUL, ast.DIV, ast.MOD, ast.ADD, ast.SUB:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.INT),
			).addPos(node.Pos)
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
			v.typeChecker.expect(ast.NewBaseTypeNode(ast.INT))
		case ast.GT, ast.GEQ, ast.LT, ast.LEQ:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.BOOL),
			).addPos(node.Pos)
			v.typeChecker.expectTwiceSame(NewSetExpectance([]ast.TypeNode{
				ast.NewBaseTypeNode(ast.INT),
				ast.NewBaseTypeNode(ast.CHAR),
			}))
		case ast.EQ, ast.NEQ:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.BOOL),
			).addPos(node.Pos)
			v.typeChecker.expectTwiceSame(NewAnyExpectance())
		case ast.AND, ast.OR:
			foundError = v.typeChecker.seen(
				ast.NewBaseTypeNode(ast.BOOL),
			).addPos(node.Pos)
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
	case *ast.ForLoopNode:
		v.symbolTable.MoveUpScope()
	case ast.SwitchNode:
		v.typeChecker.forcePop()
	case *ast.FunctionNode:
		v.symbolTable.MoveUpScope()
		v.typeChecker.forcePop()
	case *ast.ArrayLiteralNode:
		v.typeChecker.forcePop()
	case *ast.DeclareNode:
		if _, ok :=
			v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident); !ok {
			v.symbolTable.AddToScope(node.Ident.Ident, node)
		}
	case *ast.ParameterNode:
		if _, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); !ok {
			v.symbolTable.AddToScope(node.Ident.Ident, node)
		}
	}
	v.typeChecker.unfreeze(programNode)
}
