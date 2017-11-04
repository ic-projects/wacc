package ast

import "fmt"

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

	case ParameterNode:

	case SkipNode:
	case DeclareNode:
		_, ok := v.symbolTable.SearchFor(node.ident.String())
		if ok {

		} else {
			v.symbolTable.AddToScope(node.ident.String(), node)
		}

	case AssignNode:

	case ReadNode:

	case FreeNode:

	case ReturnNode:

	case ExitNode:

	case PrintNode:

	case PrintlnNode:

	case IfNode:

	case LoopNode:

	case ScopeNode:
	case IdentifierNode:
	case PairFirstElementNode:

	case PairSecondElementNode:

	case ArrayElementNode:

	case ArrayLiteralNode:

	case NewPairNode:

	case FunctionCallNode:

	case BaseType:

	case ArrayTypeNode:

	case PairTypeNode:

	case UnaryOperator:

	case BinaryOperator:

	case IntegerLiteralNode:

	case BooleanLiteralNode:

	case CharacterLiteralNode:

	case StringLiteralNode:

	case PairLiteralNode:

	case UnaryOperatorNode:

	case BinaryOperatorNode:

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
	}
	return v
}
