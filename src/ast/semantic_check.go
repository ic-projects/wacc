package ast

import (
	"fmt"
	"reflect"
)

type SemanticCheck struct {
	symbolTable  SymbolTable
	expectedType []TypeNode
}

func NewSemanticCheck() SemanticCheck {
	return SemanticCheck{
		symbolTable:  NewSymbolTable(),
		expectedType: make([]TypeNode, 0, 2),
	}
}

func (v SemanticCheck) Visit(programNode ProgramNode) Visitor {
	switch node := programNode.(type) {
	case Program:

	case FunctionNode:
		_, ok := v.symbolTable.SearchFor(node.ident.ident)
		if ok {

		} else {
			v.symbolTable.AddToScope(node.ident.ident, node)
		}
		v.symbolTable.MoveDownScope()
	case ParameterNode:
		_, ok := v.symbolTable.SearchFor(node.ident.ident)
		if ok {

		} else {
			v.symbolTable.AddToScope(node.ident.ident, node)
		}
	case SkipNode:
	case DeclareNode:
		_, ok := v.symbolTable.SearchFor(node.ident.ident)
		if ok {

		} else {
			v.symbolTable.AddToScope(node.ident.ident, node)
		}
	case AssignNode:
		v.expectedType[0] = NewLHSNode() //Maybe use list instead of empty struct
		v.expectedType[1] = NewRHSNode()
	case ReadNode:
		v.expectedType[0] = NewLHSNode()
	case FreeNode:
		v.expectedType[0] = NewBaseTypeNode(PAIR) //Or array
	case ReturnNode:

	case ExitNode:
		v.expectedType[0] = NewBaseTypeNode(INT)
	case PrintNode:
		v.expectedType[0] = nil
	case PrintlnNode:
		v.expectedType[0] = nil
	case IfNode:
		v.expectedType[0] = NewBaseTypeNode(BOOL)
	case LoopNode:
		v.expectedType[0] = NewBaseTypeNode(BOOL)
	case ScopeNode:
	case IdentifierNode:
		programNode, ok := v.symbolTable.SearchFor(node.ident)
		if !ok {

		} else if declareNode, ok := programNode.(DeclareNode); ok {

		} else if reflect.DeepEqual(v.expectedType[0], declareNode.t) {

		}
	case PairFirstElementNode:
		v.expectedType[0] = NewBaseTypeNode(PAIR)
		//Is it a assignlhs or assign rhs
	case PairSecondElementNode:
		v.expectedType[0] = NewBaseTypeNode(PAIR)
	case ArrayElementNode:
		//Check identifier
		v.expectedType[0] = NewBaseTypeNode(INT)
	case ArrayLiteralNode:
		v.expectedType[1] = v.expectedType[0] //For as length of epressions
	case NewPairNode:
	case FunctionCallNode:
		programNode, ok := v.symbolTable.SearchFor(node.ident.ident)
		if !ok {

		} else if functionNode, ok := programNode.(FunctionNode); ok {

		} else if reflect.DeepEqual(v.expectedType[0], functionNode.t) {
			//Add expected types for the paramaters
		}
	case BaseTypeNode:

	case ArrayTypeNode:

	case PairTypeNode:

	case UnaryOperator:

	case BinaryOperator:

	case IntegerLiteralNode:
		if v.expectedType[0] != NewBaseTypeNode(INT) {

		} else {

		}
	case BooleanLiteralNode:
		if v.expectedType[0] != NewBaseTypeNode(BOOL) {

		} else {

		}
	case CharacterLiteralNode:
		if v.expectedType[0] != NewBaseTypeNode(CHAR) {

		} else {

		}
	case StringLiteralNode:
		if v.expectedType[0] != NewBaseTypeNode(STRING) {

		} else {

		}
	case PairLiteralNode:
		if v.expectedType[0] != NewBaseTypeNode(PAIR) {

		} else {

		}
	case UnaryOperatorNode:
		switch node.op {
		case NOT:
			if v.expectedType[0] != NewBaseTypeNode(BOOL) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(BOOL)
			}
		case NEG:
			if v.expectedType[0] != NewBaseTypeNode(INT) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(INT)
			}
		case LEN:
			if v.expectedType[0] != NewBaseTypeNode(INT) {

			} else {
				v.expectedType[0] = NewArrayTypeNode(nil, 1) //Can be any
			}
		case ORD:
			if v.expectedType[0] != NewBaseTypeNode(INT) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(CHAR)
			}
		case CHR:
			if v.expectedType[0] != NewBaseTypeNode(CHAR) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(INT)
			}
		}
	case BinaryOperatorNode:
		switch node.op {
		case MUL, DIV, MOD, ADD, SUB:
			if v.expectedType[0] != NewBaseTypeNode(INT) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(INT)
				v.expectedType[0] = NewBaseTypeNode(INT)
			}
		case GT, GEQ, LT, LEQ:
			if v.expectedType[0] != NewBaseTypeNode(BOOL) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(INT) //Or Char
				v.expectedType[0] = NewBaseTypeNode(INT) //Has to be the same as the first
			}
		case EQ, NEQ:
			if v.expectedType[0] != NewBaseTypeNode(BOOL) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(BOOL) //Or any other
				v.expectedType[0] = NewBaseTypeNode(BOOL) //The same as the first
			}
		case AND, OR:
			if v.expectedType[0] != NewBaseTypeNode(BOOL) {

			} else {
				v.expectedType[0] = NewBaseTypeNode(BOOL)
				v.expectedType[0] = NewBaseTypeNode(BOOL)
			}
		}
	case []StatementNode:
		v.symbolTable.MoveDownScope()
	default:
		fmt.Println("UnknownNode")
	}
	return v
}

func (v SemanticCheck) Leave(programNode ProgramNode) Visitor {
	switch programNode.(type) {
	case []StatementNode:
		v.symbolTable.MoveUpScope()
	case FunctionNode:
		v.symbolTable.MoveUpScope()
	}
	return v
}
