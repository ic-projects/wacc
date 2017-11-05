package ast

type SemanticCheck struct {
	symbolTable  SymbolTable
	typeChecker  *TypeChecker
}

func NewSemanticCheck() SemanticCheck {
	return SemanticCheck{
		symbolTable: NewSymbolTable(),
		typeChecker: NewTypeChecker(),
	}
}

func (v SemanticCheck) Visit(programNode ProgramNode) Visitor {
	switch node := programNode.(type) {
	case Program:

	/*case FunctionNode:
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
	case SkipNode:*/
	case DeclareNode:
		_, ok := v.symbolTable.SearchFor(node.ident.ident)
		if ok {

		} else {
			v.symbolTable.AddToScope(node.ident.ident, node)

		}
		v.typeChecker.expect(node.t)
	case AssignNode:
		//v.expectedType[0] = NewLHSNode() //Maybe use list instead of empty struct
		//v.expectedType[1] = NewRHSNode()
	case ReadNode:
		//v.expectedType[0] = NewLHSNode()
	case FreeNode:
		v.typeChecker.expectSet([]TypeNode{NewBaseTypeNode(PAIR), ArrayTypeNode{}})
		//v.typeChecker.expect("pair") //Or array
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
		programNode, ok := v.symbolTable.SearchFor(node.ident)
		if !ok {

		} else if declareNode, ok := programNode.(DeclareNode); ok {
			v.typeChecker.seen(declareNode.t)
		}

	case PairFirstElementNode:
		//LOOK UP TYPE FOR PAIR CALL SEEN
		v.typeChecker.expect(NewBaseTypeNode(PAIR))
		//Is it a assignlhs or assign rhs
	case PairSecondElementNode:
		//LOOK UP TYPE FOR PAIR CALL SEEN
		v.typeChecker.expect(NewBaseTypeNode(PAIR))
	case ArrayElementNode:
		//Check identifier
		/*
		v.typeChecker.seen(type of array)
		for i := 0; i < dimensions; i++ {
			v.typeChecker.expect(NewBaseTypeNode(INT))
		}*/

	case ArrayLiteralNode:
		v.typeChecker.seen(ArrayTypeNode{})
	case NewPairNode:

	/*
	case FunctionCallNode:
		programNode, ok := v.symbolTable.SearchFor(node.ident.ident)
		if !ok {

		} else if functionNode, ok := programNode.(FunctionNode); ok {

		} else if reflect.DeepEqual(v.expectedType[0], functionNode.t) {
			//Add expected types for the paramaters
		}*/
	case BaseTypeNode:

	case ArrayTypeNode:

	case PairTypeNode:

	case UnaryOperator:

	case BinaryOperator:

	case IntegerLiteralNode:
		v.typeChecker.seen(NewBaseTypeNode(INT))
	case BooleanLiteralNode:
		v.typeChecker.seen(NewBaseTypeNode(BOOL))
	case CharacterLiteralNode:
		v.typeChecker.seen(NewBaseTypeNode(CHAR))
	case StringLiteralNode:
		v.typeChecker.seen(NewBaseTypeNode(STRING))
	case PairLiteralNode:
		v.typeChecker.seen(NewBaseTypeNode(PAIR))
	case UnaryOperatorNode:
		switch node.op {
		case NOT:
			v.typeChecker.seen(NewBaseTypeNode(BOOL))
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
		case NEG:
			v.typeChecker.seen(NewBaseTypeNode(INT))
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case LEN:
			v.typeChecker.seen(NewBaseTypeNode(INT))
			v.typeChecker.expectAny()
			v.typeChecker.expect(ArrayTypeNode{})
		case ORD:
			v.typeChecker.seen(NewBaseTypeNode(INT))
			v.typeChecker.expect(NewBaseTypeNode(CHAR))
		case CHR:
			v.typeChecker.seen(NewBaseTypeNode(CHAR))
			v.typeChecker.seen(NewBaseTypeNode(INT))
		}
	case BinaryOperatorNode:
		switch node.op {
		case MUL, DIV, MOD, ADD, SUB:
			v.typeChecker.seen(NewBaseTypeNode(INT))
			v.typeChecker.expect(NewBaseTypeNode(INT))
			v.typeChecker.expect(NewBaseTypeNode(INT))
		case GT, GEQ, LT, LEQ:
			v.typeChecker.seen(NewBaseTypeNode(BOOL))
			v.typeChecker.expectTwiceSame(NewSetExpectance([]TypeNode{NewBaseTypeNode(INT), NewBaseTypeNode(CHAR)}))
		case EQ, NEQ:
			v.typeChecker.seen(NewBaseTypeNode(BOOL))
			v.typeChecker.expectTwiceSame(NewAnyExpectance())
		case AND, OR:
			v.typeChecker.seen(NewBaseTypeNode(BOOL))
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
			v.typeChecker.expect(NewBaseTypeNode(BOOL))
		}
	case []StatementNode:
		v.symbolTable.MoveDownScope()
	default:
		//fmt.Println("UnknownNode")
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
