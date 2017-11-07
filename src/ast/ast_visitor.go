package ast

// Visitor is an interface that must be implemented to use the walker.
type Visitor interface {
	Visit(ProgramNode)
}

// EntryExitVisitor is an extension of the Visitor interface to allow exiting
// nodes. This is very useful for the semantic checker, for moving up scopes at
// the end of the scope.
type EntryExitVisitor interface {
	Visit(ProgramNode)
	Leave(ProgramNode)
}

// Walk will take a visitor and a programNode, and recurse downwards, calling
// visit on the programNodes below the programNode and the current programNode.
func Walk(visitor Visitor, programNode ProgramNode) {
	visitor.Visit(programNode)
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
		Walk(visitor, node.lhs)
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
  // If we have a EntryExitVisitor, use it to call leave.
	if v, ok := visitor.(EntryExitVisitor); ok {
		v.Leave(programNode)
	}
}
