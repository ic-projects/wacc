package ast

import (
	"bytes"
	"fmt"
	"os"
)

type SemanticCheck struct {
	symbolTable *SymbolTable
	typeChecker *TypeChecker
	Errors      []GenericError
}

func NewSemanticCheck() *SemanticCheck {
	return &SemanticCheck{
		symbolTable: NewSymbolTable(),
		typeChecker: NewTypeChecker(),
		Errors:      make([]GenericError, 0),
	}
}

type GenericError interface {
	String() string
	Pos() Position
}

type TypeError struct {
	pos      Position
	got      TypeNode
	expected map[TypeNode]bool
}

func (e TypeError) Pos() Position {
	return e.pos
}

func (e DeclarationError) Pos() Position {
	return e.pos
}

func (e TypeError) String() string {
	var b bytes.Buffer
	b.WriteString("Expected type ")
	i := 1
	for t := range e.expected {
		if i == len(e.expected) {
			b.WriteString(fmt.Sprintf("\"%s\"", t))
		} else {
			b.WriteString(fmt.Sprintf("\"%s\" or ", t))
		}
		i++
	}

	b.WriteString(fmt.Sprintf(" but got type \"%s\"", e.got))
	return b.String()
}

func NewTypeError(got TypeNode, expected map[TypeNode]bool) TypeError {
	return TypeError{
		got:      got,
		expected: expected,
	}
}

func (e TypeError) addPos(pos Position) TypeError {
	e.pos = pos
	return e
}

type DeclarationError struct {
	pos      Position
	isFunction bool
	isDefined bool
	identifier string
}

func NewDeclarationError(pos Position, isFunction bool, isDefined bool, identifier string) DeclarationError {
	return DeclarationError{
		pos: pos,
		isFunction: isFunction,
		isDefined: isDefined,
		identifier: identifier,
	}
}

func (e DeclarationError) String() string {
	var b bytes.Buffer
	if e.isFunction {
		if e.isDefined {
			b.WriteString(fmt.Sprintf("Function \"%s\" is already defined", e.identifier))
		} else {
			b.WriteString(fmt.Sprintf("Function \"%s\" is not defined", e.identifier))
		}
	} else {
		if e.isDefined {
			b.WriteString(fmt.Sprintf("Variable \"%s\" is already defined in the current scope", e.identifier))
		} else {
			b.WriteString(fmt.Sprintf("Variable \"%s\" is not defined in the current scope", e.identifier))
		}
	}
	return b.String()
}

func (v *SemanticCheck) PrintSymbolTable() {
	fmt.Println(v.symbolTable.String())
}


func (v *SemanticCheck) Visit(programNode ProgramNode) {
	var typeError TypeError
	var declarationError DeclarationError
	switch node := programNode.(type) {
	case Program:
		for _, f := range node.functions {
			_, ok := v.symbolTable.SearchForFunction(f.ident.ident)
			if ok {
				declarationError = NewDeclarationError(f.pos, true, true, f.ident.ident)
			} else {
				v.symbolTable.AddFunction(f.ident.ident, f)
			}
		}
	case FunctionNode:
		v.symbolTable.MoveDownScope()
		v.typeChecker.expectRepeatUntilForce(node.t)
	case ParameterNode:
		_, ok := v.symbolTable.SearchForIdent(node.ident.ident)
		if ok {
			declarationError = NewDeclarationError(node.pos, false, true, node.ident.ident)
		} else {
			v.symbolTable.AddToScope(node.ident.ident, node)
		}
	case SkipNode:
	case DeclareNode:
		_, ok := v.symbolTable.SearchForIdentInCurrentScope(node.ident.ident)
		if ok {
			declarationError = NewDeclarationError(node.pos, false, true, node.ident.ident)
			v.typeChecker.freeze(node)
		} else {
			v.symbolTable.AddToScope(node.ident.ident, node)
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
		identDec, ok := v.symbolTable.SearchForIdent(node.ident)
		if !ok {
			declarationError = NewDeclarationError(node.pos, false, false, node.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			typeError = v.typeChecker.seen(identDec.t).addPos(node.pos)
		}
	case PairFirstElementNode:
		// Look up type for pair call seen
		if identNode, ok := node.expr.(IdentifierNode); !ok {
			fmt.Printf("Not an identifier for a pair %s", identNode.ident)
			os.Exit(200)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.ident); !ok {
			declarationError = NewDeclarationError(identNode.pos, false, false, identNode.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			typeError = v.typeChecker.seen(identDec.t.(PairTypeNode).t1).addPos(node.pos)
			v.typeChecker.expect(identDec.t)
		}
	case PairSecondElementNode:
		if identNode, ok := node.expr.(IdentifierNode); !ok {
			fmt.Printf("Not an identifier for a pair %s", identNode.ident)
			os.Exit(200)
		} else if identDec, ok := v.symbolTable.SearchForIdent(identNode.ident); !ok {
			declarationError = NewDeclarationError(identNode.pos, false, false, identNode.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else {
			v.typeChecker.seen(identDec.t.(PairTypeNode).t2)
			v.typeChecker.expect(identDec.t)
		}
	case ArrayElementNode:
		// Check identifier
		identDec, ok := v.symbolTable.SearchForIdent(node.ident.ident)
		if !ok {
			declarationError = NewDeclarationError(node.pos, false, false, node.ident.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if arrayNode, ok := identDec.t.(ArrayTypeNode); !ok {
			fmt.Printf("Array access on non-array variable %s", node.ident.ident)
			os.Exit(200)
		} else {
			if dimLeft := arrayNode.dim - len(node.exprs); dimLeft == 0 {
				typeError = v.typeChecker.seen(arrayNode.t).addPos(node.pos)
			} else {
				typeError = v.typeChecker.seen(NewArrayTypeNode(arrayNode.t, dimLeft)).addPos(node.pos)
			}
		}
		for i := 0; i < len(node.exprs); i++ {
			v.typeChecker.expect(NewBaseTypeNode(INT))
		}

	case ArrayLiteralNode:
		typeError = v.typeChecker.seen(ArrayTypeNode{}).addPos(node.pos)
	case NewPairNode:
		typeError = v.typeChecker.seen(PairTypeNode{}).addPos(node.pos)
	case FunctionCallNode:
		f, ok := v.symbolTable.SearchForFunction(node.ident.ident)
		if !ok {
			declarationError = NewDeclarationError(node.pos, true, false, node.ident.ident)
			v.typeChecker.seen(nil)
			v.typeChecker.freeze(node)
		} else if len(f.params) != len(node.exprs) {
			typeError = v.typeChecker.seen(f.t).addPos(node.pos)
			fmt.Printf("Incorrect number of parameters in")
			os.Exit(200)
		} else {
			typeError = v.typeChecker.seen(f.t).addPos(node.pos)
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
		typeError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
	case BooleanLiteralNode:
		typeError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
	case CharacterLiteralNode:
		typeError = v.typeChecker.seen(NewBaseTypeNode(CHAR)).addPos(node.pos)
	case StringLiteralNode:
		typeError = v.typeChecker.seen(NewArrayTypeNode(NewBaseTypeNode(CHAR), 1)).addPos(node.pos)
	case PairLiteralNode:
		typeError = v.typeChecker.seen(NewBaseTypeNode(PAIR)).addPos(node.pos)
	case UnaryOperatorNode:
		switch node.op {
		case NOT:
			typeError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
		case NEG:
			typeError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case LEN:
			typeError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(ArrayTypeNode{})
		case ORD:
			typeError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(CHAR))
		case CHR:
			typeError = v.typeChecker.seen(NewBaseTypeNode(CHAR)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
		}
	case BinaryOperatorNode:
		switch node.op {
		case MUL, DIV, MOD, ADD, SUB:
			typeError = v.typeChecker.seen(NewBaseTypeNode(INT)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(INT))
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case GT, GEQ, LT, LEQ:
			typeError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
			v.typeChecker.expectTwiceSame(NewSetExpectance([]TypeNode{NewBaseTypeNode(INT), NewBaseTypeNode(CHAR)}))
		case EQ, NEQ:
			typeError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
			v.typeChecker.expectTwiceSame(NewAnyExpectance())
		case AND, OR:
			typeError = v.typeChecker.seen(NewBaseTypeNode(BOOL)).addPos(node.pos)
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
		}
	case []StatementNode:
		v.symbolTable.MoveDownScope()
	}

	if declarationError != (DeclarationError{}) {
		v.Errors = append(v.Errors, declarationError)
	} else if typeError.got != nil {
		v.Errors = append(v.Errors, typeError)
	}
}

func (v *SemanticCheck) Leave(programNode ProgramNode) {
	switch programNode.(type) {
	case []StatementNode:
		v.symbolTable.MoveUpScope()
	case FunctionNode:
		v.symbolTable.MoveUpScope()
		v.typeChecker.forcePop()
	case ArrayLiteralNode:
		v.typeChecker.forcePop()
	}
	v.typeChecker.unfreeze(programNode)
}
