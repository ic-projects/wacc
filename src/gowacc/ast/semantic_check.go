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
		// Add the functions when hitting program instead of each function so that
		// functions can be declared in any order.
		for _, f := range node.functions {
			if functionNode, ok := v.symbolTable.SearchForFunction(f.ident.ident); ok {
				foundError = NewPreviouslyDeclared(NewDeclarationError(f.pos, true, true, f.ident.ident), functionNode.pos)
			} else {
				v.symbolTable.AddFunction(f.ident.ident, f)
			}
		}
	case FunctionNode:
		// Move down scope so that the parameters are on a new scope.
		v.symbolTable.MoveDownScope()
		v.typeChecker.expectRepeatUntilForce(node.t)
	case ParameterNode:
		if declareNode, ok := v.symbolTable.SearchForIdent(node.ident.ident); ok {
			foundError = NewPreviouslyDeclared(NewDeclarationError(node.pos, false, true, node.ident.ident), declareNode.pos)
		}
	case SkipNode:
	case DeclareNode:
		if declareNode, ok := v.symbolTable.SearchForIdentInCurrentScope(node.ident.ident); ok {
			foundError = NewPreviouslyDeclared(NewDeclarationError(node.pos, false, true, node.ident.ident), declareNode.pos)
			v.typeChecker.freeze(node)
		} else {
			v.typeChecker.expect(node.t)
		}
	case AssignNode:
		if arr, ok := node.lhs.(ArrayElementNode); ok {
			arr.assign = true
		}
		if fst, ok := node.lhs.(PairFirstElementNode); ok {
			fst.assign = true
		}
		if snd, ok := node.lhs.(PairSecondElementNode); ok {
			snd.assign = true
		}
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
		if identDec, ok := v.symbolTable.SearchForIdent(node.ident); !ok {
			foundError = NewDeclarationError(node.pos, false, false, node.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(identDec.t).addPos(node.pos)
			if foundError != nil {
				foundError = NewTypeErrorDeclaration(foundError.(TypeError), identDec.pos)
			}
		}
	case PairFirstElementNode:
		//  Look up type for pair call seen
		if identNode, ok := node.expr.(IdentifierNode); !ok {
			foundError = NewCustomError(node.pos, "Cannot access first element of null")
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.ident); !ok {
			foundError = NewDeclarationError(identNode.pos, false, false, identNode.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(identDec.t.(PairTypeNode).t1).addPos(node.pos)
			v.typeChecker.expect(identDec.t)
		}
	case PairSecondElementNode:
		if identNode, ok := node.expr.(IdentifierNode); !ok {
			foundError = NewCustomError(node.pos, "Cannot access second element of null")
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.ident); !ok {
			foundError = NewDeclarationError(identNode.pos, false, false, identNode.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			v.typeChecker.seen(identDec.t.(PairTypeNode).t2)
			v.typeChecker.expect(identDec.t)
		}
	case ArrayElementNode:
		if identDec, ok := v.symbolTable.SearchForIdent(node.ident.ident); !ok {
			foundError = NewDeclarationError(node.pos, false, false, node.ident.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if arrayNode, ok := identDec.t.(ArrayTypeNode); !ok {
			foundError = NewCustomError(node.pos, fmt.Sprintf("Array access on non-array variable \"%s\"", node.ident.ident))
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			// If we have an array or a single element (for use in newsted arrays).
			if dimLeft := arrayNode.dim - len(node.exprs); dimLeft == 0 {
				foundError = v.typeChecker.seen(arrayNode.t).addPos(node.pos)
			} else {
				foundError = v.typeChecker.seen(NewArrayTypeNode(arrayNode.t, dimLeft)).addPos(node.pos)
			}
		}
		for i := 0; i < len(node.exprs); i++ {
			v.typeChecker.expect(NewBaseTypeNode(INT))
		}
	case ArrayLiteralNode:
		foundError = v.typeChecker.seen(ArrayTypeNode{}).addPos(node.pos)
	case NewPairNode:
		foundError = v.typeChecker.seen(PairTypeNode{}).addPos(node.pos)
	case FunctionCallNode:
		if f, ok := v.symbolTable.SearchForFunction(node.ident.ident); !ok {
			foundError = NewDeclarationError(node.pos, true, false, node.ident.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if len(f.params) != len(node.exprs) {
			foundError = NewCustomError(node.pos, fmt.Sprintf("Incorrect number of parameters for function \"%s\" (Expected: %d, Given: %d)", node.ident.ident, len(f.params), len(node.exprs)))
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			foundError = v.typeChecker.seen(f.t).addPos(node.pos)
			for i := len(f.params) - 1; i >= 0; i-- {
				v.typeChecker.expect(f.params[i].t)
			}
		}
	case BaseTypeNode:
	case ArrayTypeNode:
	case PairTypeNode:
	case UnaryOperator:
	case BinaryOperator:
	case IntegerLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
	case BooleanLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
	case CharacterLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(CHAR)).addPos(node.pos)
	case StringLiteralNode:
		foundError = v.typeChecker.seen(NewStringArrayTypeNode()).addPos(node.pos)
	case PairLiteralNode:
		foundError = v.typeChecker.seen(NewBaseTypeNode(PAIR)).addPos(node.pos)
	case UnaryOperatorNode:
		switch node.op {
		case NOT:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
		case NEG:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case LEN:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(ArrayTypeNode{})
		case ORD:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(CHAR))
		case CHR:
			foundError = v.typeChecker.seen(NewBaseTypeNode(CHAR)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
		}
	case BinaryOperatorNode:
		switch node.op {
		case MUL, DIV, MOD, ADD, SUB:
			foundError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case GT, GEQ, LT, LEQ:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
			v.typeChecker.expectTwiceSame(NewSetExpectance([]TypeNode{NewBaseTypeNode(INT), NewBaseTypeNode(CHAR)}))
		case EQ, NEQ:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
			v.typeChecker.expectTwiceSame(NewAnyExpectance())
		case AND, OR:
			foundError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
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
		if _, ok := v.symbolTable.SearchForIdentInCurrentScope(node.ident.ident); !ok {
			v.symbolTable.AddToScope(node.ident.ident, node)
		}
	case ParameterNode:
		if _, ok := v.symbolTable.SearchForIdent(node.ident.ident); !ok {
			v.symbolTable.AddToScope(node.ident.ident, node)
		}
	}
	v.typeChecker.unfreeze(programNode)
}

func (v *SemanticCheck) PrintErrors(filepath string) {
	PrintErrors(v.Errors, filepath)
}
