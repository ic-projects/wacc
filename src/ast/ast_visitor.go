package ast

import (
	"fmt"
)

type Visitor interface {
	Visit(ProgramNode) Visitor
}

type EntryExitVisitor interface {
	Visit(ProgramNode) Visitor
	Leave(ProgramNode) Visitor
}

func Walk(visitor Visitor, programNode ProgramNode) {
	visitor = visitor.Visit(programNode)
	switch node := programNode.(type) {
	case []StatementNode:
		for _, s := range node {
			Walk(visitor, s)
		}
	case Program:
		for _, f := range node.functions {
			Walk(visitor, f)
		}
	case FunctionNode:
		for _, p := range node.params {
			Walk(visitor, p)
		}
		Walk(visitor, node.stats)
	case ParameterNode:

	case SkipNode:

	case DeclareNode:
		Walk(visitor, node.rhs)
	case AssignNode:
		Walk(visitor, node.rhs)
	case ReadNode:
		Walk(visitor, node.expr)
	case FreeNode:
		Walk(visitor, node.expr)
	case ReturnNode:
		Walk(visitor, node.expr)
	case ExitNode:
		Walk(visitor, node.expr)
	case PrintNode:
		Walk(visitor, node.expr)
	case PrintlnNode:
		Walk(visitor, node.expr)
	case IfNode:
		Walk(visitor, node.expr)
		Walk(visitor, node.ifStats)
		Walk(visitor, node.elseStats)
	case LoopNode:
		Walk(visitor, node.expr)
		Walk(visitor, node.stats)
	case ScopeNode:
		Walk(visitor, node.stats)
	case IdentifierNode:

	case PairFirstElementNode:
		Walk(visitor, node.expr)
	case PairSecondElementNode:
		Walk(visitor, node.expr)
	case ArrayElementNode:
		for _, e := range node.exprs {
			Walk(visitor, e)
		}
	case ArrayLiteralNode:
		for _, e := range node.exprs {
			Walk(visitor, e)
		}
	case NewPairNode:
		Walk(visitor, node.fst)
		Walk(visitor, node.snd)
	case FunctionCallNode:
		for _, e := range node.exprs {
			Walk(visitor, e)
		}
	case BaseType:

	case BaseTypeNode:

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
		Walk(visitor, node.expr)
	case BinaryOperatorNode:
		Walk(visitor, node.expr1)
		Walk(visitor, node.expr2)
	default:

	}
	if v, ok := visitor.(EntryExitVisitor); ok {
		v.Leave(programNode)
	}
}

type Printer struct {
}

func NewPrinter() Printer {
	return Printer{}
}

func (v Printer) Visit(programNode ProgramNode) Visitor {
	switch programNode.(type) {
	case Program:
		fmt.Println("Program")
	case FunctionNode:
		fmt.Println("FunctionNode")
	case ParameterNode:
		fmt.Println("ParamaterNode")
	case SkipNode:
		fmt.Println("SkipNode")
	case DeclareNode:
		fmt.Println("DeclareNode")
	case AssignNode:
		fmt.Println("AssignNode")
	case ReadNode:
		fmt.Println("ReadNode")
	case FreeNode:
		fmt.Println("FreeNode")
	case ReturnNode:
		fmt.Println("ReturnNode")
	case ExitNode:
		fmt.Println("ExitNode")
	case PrintNode:
		fmt.Println("PrintNode")
	case PrintlnNode:
		fmt.Println("PrintlnNode")
	case IfNode:
		fmt.Println("IfNode")
	case LoopNode:
		fmt.Println("LoopNode")
	case ScopeNode:
		fmt.Println("ScopeNode")
	case IdentifierNode:
		fmt.Println("IdentifierNode")
	case PairFirstElementNode:
		fmt.Println("PairFirstElementNode")
	case PairSecondElementNode:
		fmt.Println("PairSecondElementNode")
	case ArrayElementNode:
		fmt.Println("ArrayElementNode")
	case ArrayLiteralNode:
		fmt.Println("ArrayLiteralNode")
	case NewPairNode:
		fmt.Println("NewPairNode")
	case FunctionCallNode:
		fmt.Println("FunctionCallNode")
	case BaseType:
		fmt.Println("BaseType")
	case ArrayTypeNode:
		fmt.Println("ArrayTypeNode")
	case PairTypeNode:
		fmt.Println("PairTypeNode")
	case UnaryOperator:
		fmt.Println("UnOp")
	case BinaryOperator:
		fmt.Println("BinOp")
	case IntegerLiteralNode:
		fmt.Println("IntegerLiteralNode")
	case BooleanLiteralNode:
		fmt.Println("BooleanLiteralNode")
	case CharacterLiteralNode:
		fmt.Println("CharacterLiteralNode")
	case StringLiteralNode:
		fmt.Println("StringLiteralNode")
	case PairLiteralNode:
		fmt.Println("PairLiteralNode")
	case UnaryOperatorNode:
		fmt.Println("UnOpNode")
	case BinaryOperatorNode:
		fmt.Println("BinOpNode")
	case []StatementNode:
		fmt.Println("Statements")
	default:
		fmt.Println("Unknown")
	}
	return v
}
