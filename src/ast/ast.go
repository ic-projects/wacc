ProgramNode {
	[FunctionNode]
}

FunctionNode {
	Position
	TypeNode
	Identifier
	[StatementNode]
}

// ---- StatementNodes ----

DeclareNode {
	Position
	TypeNode
	Identifier
	RHSNode
}

AssignNode {
	Position
	LHSNode
	RHSNode
}

ReadNode {
	Position
	ExpressionNode
}

FreeNode {
	Position
	ExpressionNode
}

ReturnNode {
	Position
	ExpressionNode
}

ExitNode {
	Position
	ExpressionNode
}

PrintNode {
	Position
	ExpressionNode
}

PrintlnNode {
	Position
	ExpressionNode
}

IfNode {
	Position
	ExpressionNode
	[StatementNodes]
	[StatementNodes]
}

LoopNode {
	Position
	ExpressionNode
	[StatementNodes]
}

ScopeNode {
	Position
	[StatementNodes]
}

// LHSNodes

IdentifierNode {
	Position
	Identifier
}

PairFirstElementNode {
	Position
	Identifier
	ExpressionNode
}

PairSecondElementNode {
	Position
	Identifier
	ExpressionNode
}

ArrayElementNode {
	Position
	Identifier
	ExpressionNode
}

// ---- RHSNodes ----

// ExpressionNode

ArrayLiteralNode {
	Position
	[ExpressionNode]
}

PairLiteralNode {
	Position
	ExpressionNode
	ExpressionNode
}

// PairFirstElementNode

// PairSecondElementNode

FunctionCallNode {
	Position
	Identifier
	[ExpressionNode]
}

// ---- TypeNodes ----

BaseTypeNode {
	BaseType
}

ArrayTypeNode {
	Dimension
	ArrayableTypeNode
}

PairTypeNode {
	PairableTypeNode
	PairableTypeNode
}

// ---- ArrayableTypeNodes ----

// BaseTypeNode

// PairTypeNode

// ---- PairableTypeNodes ----

// BaseTypeNode

// ArrayTypeNode

PairablePairTypeNode {}

// ---- ExpressionNodes ----

IntegerLiteralNode {
	Position
	Integer
}

BooleanLiteralNode {
	Position
	Boolean
}

CharacterLiteralNode {
	Position
	Character
}

StringLiteralNode {
	Position
	String
}

// PairLiteralNode

// IdentifierNode

// ArrayElementNode

UnaryOperatorNode {
	Position
	UnaryOperator
	ExpressionNode
}

BinaryOperatorNode {
	Position
	BinaryOperator
	ExpressionNode
	ExpressionNode
}
