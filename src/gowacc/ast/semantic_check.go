package ast

import (
	"fmt"
)

// SemanticCheck is a struct that implements EntryExitVisitor to be called with
// Walk. It stores a SymbolTable, a TypeChecker, and a list of GenericErrors.
type SemanticCheck struct {
	symbolTable *SymbolTable
	typeChecker *TypeChecker
	Errors      []GenericError
}

// NewSemanticCheck returns an initialised SemanticCheck
func NewSemanticCheck() *SemanticCheck {
	return &SemanticCheck{
		symbolTable: NewSymbolTable(),
		typeChecker: NewTypeChecker(),
		Errors:      make([]GenericError, 0),
	}
}

func (v *SemanticCheck) PrintSymbolTable() {
	fmt.Println(v.symbolTable.String())
}

func (v *SemanticCheck) SymbolTable() *SymbolTable {
	return v.symbolTable
}

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *SemanticCheck) Visit(programNode ProgramNode) {
	var foundError GenericError
	foundError = nil
	switch node := programNode.(type) {
	case Program:
		// Add the Functions when hitting program instead of each function so that
		// Functions can be declared in any order.
		for _, f := range node.Functions {
			if functionNode, ok := v.symbolTable.SearchForFunction(f.Ident.Ident); ok {
				foundError = NewPreviouslyDeclared(NewDeclarationError(f.Pos, true, true, f.Ident.Ident), functionNode.Pos)
			} else {
				v.symbolTable.AddFunction(f.Ident.Ident, f)
			}
		}
	case FunctionNode:
		// Move down scope so that the parameters are on a new scope.
		v.symbolTable.MoveDownScope()
		v.typeChecker.expectRepeatUntilForce(node.T)
	case ParameterNode:
		if declareNode, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); ok {
			foundError = NewPreviouslyDeclared(NewDeclarationError(node.Pos, false, true, node.Ident.Ident), declareNode.pos)
		}
	case SkipNode:
	case DeclareNode:
		if declareNode, ok := v.symbolTable.SearchForIdentInCurrentScope(node.ident.Ident); ok {
			foundError = NewPreviouslyDeclared(NewDeclarationError(node.pos, false, true, node.ident.Ident), declareNode.pos)
			v.typeChecker.freeze(node)
		} else {
			v.typeChecker.expect(node.t)
		}
	case AssignNode:
		v.typeChecker.expectTwiceSame(NewAnyExpectance())
	case ReadNode:
		v.typeChecker.expectSet([]TypeNode{NewBaseTypeNode(INT), NewBaseTypeNode(CHAR)})
	case FreeNode:
		v.typeChecker.expectSet([]TypeNode{PairTypeNode{}, ArrayTypeNode{}})
	case ReturnNode:
	case ExitNode:
		v.typeChecker.expect(NewBaseTypeNode(INT))
	case PrintNode:
		v.typeChecker.expectAny()
	case PrintlnNode:
		v.typeChecker.expectAny()
	case IfNode:
		v.typeChecker.expect(NewBaseTypeNode(BOOL))
	case LoopNode:
		v.typeChecker.expect(NewBaseTypeNode(BOOL))
	case ScopeNode:
	case IdentifierNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.Ident); !ok {
			foundError = NewDeclarationError(node.Pos, false, false, node.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(identDec.t).addPos(node.Pos)
			if foundError != nil {
				foundError = NewTypeErrorDeclaration(foundError.(TypeError), identDec.pos)
			}
		}
	case PairFirstElementNode:
		//  Look up type for pair call seen
		if identNode, ok := node.Expr.(IdentifierNode); !ok {
			foundError = NewCustomError(node.Pos, "Cannot access first element of null")
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.Ident); !ok {
			foundError = NewDeclarationError(identNode.Pos, false, false, identNode.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(identDec.t.(PairTypeNode).t1).addPos(node.Pos)
			v.typeChecker.expect(identDec.t)
		}
	case PairSecondElementNode:
		if identNode, ok := node.Expr.(IdentifierNode); !ok {
			foundError = NewCustomError(node.Pos, "Cannot access second element of null")
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.Ident); !ok {
			foundError = NewDeclarationError(identNode.Pos, false, false, identNode.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			v.typeChecker.seen(identDec.t.(PairTypeNode).t2)
			v.typeChecker.expect(identDec.t)
		}
	case ArrayElementNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); !ok {
			foundError = NewDeclarationError(node.Pos, false, false, node.Ident.Ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if arrayNode, ok := identDec.t.(ArrayTypeNode); !ok {
			foundError = NewCustomError(node.Pos, fmt.Sprintf("Array access on non-array variable \"%s\"", node.Ident.Ident))
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			// If we have an array or a single element (for use in newsted arrays).
			if dimLeft := arrayNode.dim - len(node.Exprs); dimLeft == 0 {
				foundError = v.typeChecker.seen(arrayNode.t).addPos(node.Pos)
			} else {
				foundError = v.typeChecker.seen(NewArrayTypeNode(arrayNode.t, dimLeft)).addPos(node.Pos)
			}
		}
		for i := 0; i < len(node.Exprs); i++ {
			v.typeChecker.expect(NewBaseTypeNode(INT))
		}
	case ArrayLiteralNode:
		foundError = v.typeChecker.seen(ArrayTypeNode{}).addPos(node.Pos)
	case NewPairNode:
		foundError = v.typeChecker.seen(PairTypeNode{}).addPos(node.Pos)
	case FunctionCallNode:
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
	case BaseTypeNode:
	case ArrayTypeNode:
	case PairTypeNode:
	case UnaryOperator:
	case BinaryOperator:
	case IntegerLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.Pos)
	case BooleanLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.Pos)
	case CharacterLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(CHAR)).addPos(node.Pos)
	case StringLiteralNode:
		foundError = v.typeChecker.seen(NewStringArrayTypeNode()).addPos(node.Pos)
	case PairLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(PAIR)).addPos(node.Pos)
	case UnaryOperatorNode:
		switch node.Op {
		case NOT:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.Pos)
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
		case NEG:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.Pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case LEN:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.Pos)
			v.typeChecker.expect(ArrayTypeNode{})
		case ORD:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.Pos)
			v.typeChecker.expect(NewBaseTypeNode(CHAR))
		case CHR:
			foundError = v.typeChecker.seen(NewBaseTypeNode(CHAR)).addPos(node.Pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
		}
	case BinaryOperatorNode:
		switch node.Op {
		case MUL, DIV, MOD, ADD, SUB:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.Pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case GT, GEQ, LT, LEQ:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.Pos)
			v.typeChecker.expectTwiceSame(NewSetExpectance([]TypeNode{NewBaseTypeNode(INT), NewBaseTypeNode(CHAR)}))
		case EQ, NEQ:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.Pos)
			v.typeChecker.expectTwiceSame(NewAnyExpectance())
		case AND, OR:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.Pos)
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
		}
	case []StatementNode:
		v.symbolTable.MoveDownScope()
	}

	// If we have an error, add it to the list of errors.
	if foundError != nil {
		v.Errors = append(v.Errors, foundError)
	}
}

// Leave will be called to leave the current node.
func (v *SemanticCheck) Leave(programNode ProgramNode) {
	switch node := programNode.(type) {
	case []StatementNode:
		v.symbolTable.MoveUpScope()
	case FunctionNode:
		v.symbolTable.MoveUpScope()
		v.typeChecker.forcePop()
	case ArrayLiteralNode:
		v.typeChecker.forcePop()
	case DeclareNode:
		if _, ok := v.symbolTable.SearchForIdentInCurrentScope(node.ident.Ident); !ok {
			v.symbolTable.AddToScope(node.ident.Ident, node)
		}
	case ParameterNode:
		if _, ok := v.symbolTable.SearchForIdent(node.Ident.Ident); !ok {
			v.symbolTable.AddToScope(node.Ident.Ident, node)
		}
	}
	v.typeChecker.unfreeze(programNode)
}

func (v *SemanticCheck) PrintErrors(filepath string) {
	PrintErrors(v.Errors, filepath)
}
