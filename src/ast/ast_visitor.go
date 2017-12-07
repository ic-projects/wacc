package ast

// Visitor is an interface that must be implemented to use the walker.
type Visitor interface {
	Visit(ProgramNode)
}

// EntryExitVisitor is an extension of the Visitor interface to allow exiting
// nodes. This is very useful for the semantic checker, for moving up scopes at
// the end of the scope for example.
type EntryExitVisitor interface {
	Visitor
	Leave(ProgramNode)
}

// EntryExitVisitorOptional is an extension of the EntryExitVisitor interface to
// allow manual walking through the AST
type EntryExitVisitorOptional interface {
	EntryExitVisitor
	NoRecurse(ProgramNode) bool
}

// Walk will take a visitor and a programNode, and recurse downwards, calling
// visit on the programNodes below the programNode and the current programNode.
func Walk(visitor Visitor, programNode ProgramNode) {
	visitor.Visit(programNode)
	if v, ok := visitor.(EntryExitVisitorOptional); !ok ||
		!v.NoRecurse(programNode) {
		programNode.walkNode(visitor)
	}
	// If we have a EntryExitVisitor, use it to call leave.
	if v, ok := visitor.(EntryExitVisitor); ok {
		v.Leave(programNode)
	}
}
