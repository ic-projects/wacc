package main

// Visitor is an interface that must be implemented to use the walker.
type Visitor interface {
	Visit(ProgramNode)
}

// EntryExitVisitor is an extension of the Visitor interface to allow exiting
// nodes. This is very useful for the semantic checker, for moving up scopes at
// the end of the scope for example.
type EntryExitVisitor interface {
	Visit(ProgramNode)
	Leave(ProgramNode)
}

// EntryExitVisitorOptional is an extension of the EntryExitVisitor interface to
// allow manual walking through the AST
type EntryExitVisitorOptional interface {
	Visit(ProgramNode)
	Leave(ProgramNode)
	NoRecurse(ProgramNode) bool
}

// Walk will take a visitor and a programNode, and recurse downwards, calling
// visit on the programNodes below the programNode and the current programNode.
func Walk(visitor Visitor, programNode ProgramNode) {
	visitor.Visit(programNode)
	if v, ok := visitor.(EntryExitVisitorOptional); !ok ||
		!v.NoRecurse(programNode) {
		switch node := (programNode).(type) {
		case []StatementNode:
			for _, s := range node {
				Walk(visitor, s)
			}
		case *Program:
			for _, f := range node.Structs {
				Walk(visitor, f)
			}
			for _, f := range node.Functions {
				Walk(visitor, f)
			}
		case *StructNode:
			for _, f := range node.Types {
				Walk(visitor, f)
			}
		case *FunctionNode:
			Walk(visitor, node.Params)
			Walk(visitor, node.Stats)
		case []*ParameterNode:
			for _, p := range node {
				Walk(visitor, p)
			}
		case *DeclareNode:
			Walk(visitor, node.Rhs)
		case *AssignNode:
			Walk(visitor, node.Lhs)
			Walk(visitor, node.Rhs)
		case *ReadNode:
			Walk(visitor, node.Lhs)
		case *FreeNode:
			Walk(visitor, node.Expr)
		case *ReturnNode:
			Walk(visitor, node.Expr)
		case *ExitNode:
			Walk(visitor, node.Expr)
		case *PrintNode:
			Walk(visitor, node.Expr)
		case *PrintlnNode:
			Walk(visitor, node.Expr)
		case *IfNode:
			Walk(visitor, node.Expr)
			Walk(visitor, node.IfStats)
			Walk(visitor, node.ElseStats)
		case *LoopNode:
			Walk(visitor, node.Expr)
			Walk(visitor, node.Stats)
		case *ForLoopNode:
			Walk(visitor, node.Initial)
			Walk(visitor, node.Expr)
			Walk(visitor, node.Update)
			Walk(visitor, node.Stats)
		case *ScopeNode:
			Walk(visitor, node.Stats)
		case *PairFirstElementNode:
			Walk(visitor, node.Expr)
		case *PairSecondElementNode:
			Walk(visitor, node.Expr)
		case *ArrayElementNode:
			for _, e := range node.Exprs {
				Walk(visitor, e)
			}
		case *ArrayLiteralNode:
			for _, e := range node.Exprs {
				Walk(visitor, e)
			}
		case *NewPairNode:
			Walk(visitor, node.Fst)
			Walk(visitor, node.Snd)
		case *StructNewNode:
			for _, e := range node.Exprs {
				Walk(visitor, e)
			}
		case *FunctionCallNode:
			for _, e := range node.Exprs {
				Walk(visitor, e)
			}
		case *UnaryOperatorNode:
			Walk(visitor, node.Expr)
		case *BinaryOperatorNode:
			Walk(visitor, node.Expr1)
			Walk(visitor, node.Expr2)
		}
	}
	// If we have a EntryExitVisitor, use it to call leave.
	if v, ok := visitor.(EntryExitVisitor); ok {
		v.Leave(programNode)
	}
}
