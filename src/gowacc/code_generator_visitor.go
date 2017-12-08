package main

import (
	"ast"
	"bytes"
	"fmt"
	"utils"
)

/**************** FUNCTION NODE ****************/

func (v *CodeGenerator) visitFunctionNode(node *ast.FunctionNode) {
	v.currentStackPos = 0
	v.freeRegisters = utils.NewRegisterStackWith(utils.FreeRegisters())
	v.symbolTable.MoveNextScope()
	if node.Ident.Ident == "" {
		v.addFunction("main")
	} else {
		var buff bytes.Buffer
		for _, p := range node.Params {
			buff.WriteString(p.T.Hash())
		}
		v.addFunction("f" + buff.String() + "_" + node.Ident.Ident)
	}
	v.addCode(NewPush(utils.LR).armAssembly())
}

func (v *CodeGenerator) leaveFunctionNode(node *ast.FunctionNode) {
	if v.symbolTable.CurrentScope.ScopeSize > 0 {
		// Cannot add more than 1024 from SP at once, so do it in multiple
		// iterations.

		// Not using addToStackPointer function as we want v.currentStackPos
		// to be unchanged.
		i := v.symbolTable.CurrentScope.ScopeSize
		for ; i > 1024; i -= 1024 {
			v.addCode(NewAdd(utils.SP, utils.SP, 1024).armAssembly())
		}
		v.addCode(NewAdd(utils.SP, utils.SP, i).armAssembly())
	}
	v.symbolTable.MoveUpScope()
	if node.Ident.Ident == "" {
		v.addCode(NewLoad(W, utils.R0, ConstantAddress(0)).armAssembly())
		v.addCode(NewPop(utils.PC).armAssembly())
	}
	v.addCode(".ltorg")
}

/**************** PARAMETERS ****************/

func (v *CodeGenerator) visitParameters(node ast.Parameters) {
	registers := utils.ReturnRegisters()
	parametersFromRegsSize := 0
	parametersFromStackSize := 0
	for n, e := range node {
		dec, _ := v.symbolTable.SearchForIdent(e.Ident.Ident)
		dec.IsDeclared = true
		if n < len(registers) {
			// Go through parameters stored in R0 - R4 first
			parametersFromRegsSize += ast.SizeOf(e.T)
			dec.AddLocation(utils.NewStackOffsetLocation(parametersFromRegsSize))
		} else {
			// Then go through parameters stored on stack
			dec.AddLocation(
				utils.NewStackOffsetLocation(parametersFromStackSize -
					ast.SizeOf(ast.NewBaseTypeNode(ast.INT))))
			parametersFromStackSize -= ast.SizeOf(e.T)
		}
	}

	if parametersFromRegsSize > 0 {
		v.subtractFromStackPointer(parametersFromRegsSize)
	}

	v.symbolTable.CurrentScope.ScopeSize = parametersFromRegsSize
	for n, e := range node {
		dec, _ := v.symbolTable.SearchForIdent(e.Ident.Ident)
		if n < len(registers) {
			v.addCode(NewStore(
				store(ast.SizeOf(e.T)),
				registers[n],
				v.LocationOf(dec.Location),
			).armAssembly())
		}
	}
}

/**************** ASSIGN NODE ****************/

func (v *CodeGenerator) visitAssignNode(node *ast.AssignNode) {
	// Rhs
	ast.Walk(v, node.RHS)
	rhsRegister := v.returnRegisters.Pop()
	// Lhs
	switch lhsNode := node.LHS.(type) {
	case *ast.ArrayElementNode:
		ast.Walk(v, lhsNode)
		lhsRegister := v.getReturnRegister()
		dec := v.symbolTable.SearchForDeclaredIdent(lhsNode.Ident.Ident)
		arr := ast.ToValue(dec.T).(ast.ArrayTypeNode)
		v.addCode(NewStoreReg(
			store(ast.SizeOf(arr.GetDimElement(len(lhsNode.Exprs)))),
			rhsRegister,
			lhsRegister,
		).armAssembly())
	case *ast.StructElementNode:
		ast.Walk(v, lhsNode)
		lhsRegister := v.getReturnRegister()
		v.addCode(NewStoreReg(
			store(ast.SizeOf(lhsNode.StructType.T)),
			rhsRegister,
			lhsRegister,
		).armAssembly())
	case *ast.PairFirstElementNode:
		ast.Walk(v, lhsNode)
		lhsRegister := v.getReturnRegister()
		v.addCode(NewStoreReg(
			store(ast.SizeOf(ast.Type(lhsNode.Expr, v.symbolTable))),
			rhsRegister,
			lhsRegister,
		).armAssembly())
	case *ast.PairSecondElementNode:
		ast.Walk(v, lhsNode)
		lhsRegister := v.getReturnRegister()
		v.addCode(NewStoreReg(
			store(ast.SizeOf(ast.Type(lhsNode.Expr, v.symbolTable))),
			rhsRegister,
			lhsRegister,
		).armAssembly())
	case *ast.IdentifierNode:
		ident := v.symbolTable.SearchForDeclaredIdent(lhsNode.Ident)
		if ident.Location != nil {
			v.addCode(NewStore(
				store(ast.SizeOf(ident.T)),
				rhsRegister,
				v.LocationOf(ident.Location),
			).armAssembly())
		}
	}
	v.freeRegisters.Push(rhsRegister)
}

/**************** READ NODE ****************/

func (v *CodeGenerator) visitReadNode(node *ast.ReadNode) {
	if ident, ok := node.LHS.(*ast.IdentifierNode); ok {
		dec := v.symbolTable.SearchForDeclaredIdent(ident.Ident)
		v.addCode(NewAdd(
			v.getFreeRegister(),
			utils.SP,
			v.PointerTo(dec.Location),
		).armAssembly())
	} else {
		ast.Walk(v, node.LHS)
	}
	v.addCode(NewMove(utils.R0, v.getReturnRegister()).armAssembly())
	if ast.SizeOf(ast.Type(node.LHS, v.symbolTable)) == 1 {
		v.callLibraryFunction(AL, readChar)
	} else {
		v.callLibraryFunction(AL, readInt)
	}
}

/**************** IF NODE ****************/

func (v *CodeGenerator) visitIfNode(node *ast.IfNode) {
	// Labels
	elseLabel := fmt.Sprintf("ELSE%d", v.labelCount)
	endifLabel := fmt.Sprintf("ENDIF%d", v.labelCount)
	v.labelCount++
	// Cond
	ast.Walk(v, node.Expr)
	v.addCode(NewCompareInt(v.getReturnRegister(), 0).armAssembly())
	v.addCode(NewCondBranch(EQ, elseLabel).armAssembly())
	// If
	ast.Walk(v, node.IfStats)
	v.addCode(NewBranch(endifLabel).armAssembly())
	// Else
	v.addCode("%s:", elseLabel)
	ast.Walk(v, node.ElseStats)
	// Fi
	v.addCode("%s:", endifLabel)
}

/**************** SWITCH NODE ****************/

func (v *CodeGenerator) visitSwitchNode(node *ast.SwitchNode) {
	labelNumber := v.labelCount
	endLabel := fmt.Sprintf("END%d", labelNumber)
	v.labelCount++
	ast.Walk(v, node.Expr)
	r := v.returnRegisters.Pop()
	for i, c := range node.Cases {
		caseLabel := fmt.Sprintf("CASE%d_%d", labelNumber, i)
		if !c.IsDefault {
			ast.Walk(v, c.Expr)
			v.addCode(NewCompare(v.getReturnRegister(), r).armAssembly())
			v.addCode(NewCondBranch(EQ, caseLabel).armAssembly())
		} else {
			v.addCode(NewBranch(caseLabel).armAssembly())
		}
	}
	v.addCode(NewBranch(endLabel).armAssembly())
	v.freeRegisters.Push(r)
	for i, c := range node.Cases {
		caseLabel := fmt.Sprintf("CASE%d_%d", labelNumber, i)
		v.addCode("%s:", caseLabel)
		ast.Walk(v, c.Stats)
		v.addCode(NewBranch(endLabel).armAssembly())
	}
	v.addCode("%s:", endLabel)
}

/**************** LOOP NODE ****************/

func (v *CodeGenerator) visitLoopNode(node *ast.LoopNode) {
	// Labels
	doLabel := fmt.Sprintf("DO%d", v.labelCount)
	whileLabel := fmt.Sprintf("WHILE%d", v.labelCount)
	v.labelCount++
	v.addCode(NewBranch(whileLabel).armAssembly())
	// Do
	v.addCode("%s:", doLabel)
	ast.Walk(v, node.Stats)
	// While
	v.addCode("%s:", whileLabel)
	ast.Walk(v, node.Expr)
	v.addCode(NewCompareInt(v.getReturnRegister(), 1).armAssembly())
	v.addCode(NewCondBranch(EQ, doLabel).armAssembly())
}

/**************** FOR LOOP NODE ****************/

func (v *CodeGenerator) visitForLoopNode(node *ast.ForLoopNode) {
	ast.Walk(v, node.Initial)
	// Labels
	doLabel := fmt.Sprintf("DO%d", v.labelCount)
	whileLabel := fmt.Sprintf("WHILE%d", v.labelCount)
	v.labelCount++
	v.addCode(NewBranch(whileLabel).armAssembly())
	// Do
	v.addCode("%s:", doLabel)
	ast.Walk(v, node.Stats)
	ast.Walk(v, node.Update)
	// While
	v.addCode("%s:", whileLabel)
	ast.Walk(v, node.Expr)
	v.addCode(NewCompareInt(v.getReturnRegister(), 1).armAssembly())
	v.addCode(NewCondBranch(EQ, doLabel).armAssembly())
}

/**************** IDENTIFIER NODE ****************/

func (v *CodeGenerator) visitIdentifierNode(node *ast.IdentifierNode) {
	dec := v.symbolTable.SearchForDeclaredIdent(node.Ident)
	v.addCode(NewLoad(
		load(ast.SizeOf(dec.T)),
		v.getFreeRegister(),
		v.LocationOf(dec.Location),
	).armAssembly())
}

/**************** ARRAY ELEMENT NODE ****************/

func (v *CodeGenerator) visitArrayElementNode(node *ast.ArrayElementNode) {
	ast.Walk(v, node.Ident)
	identRegister := v.returnRegisters.Pop()

	length := len(node.Exprs)
	symbol := v.symbolTable.SearchForDeclaredIdent(node.Ident.Ident)
	lastIsCharOrBool := ast.SizeOf(
		ast.ToValue(symbol.T).(ast.ArrayTypeNode).GetDimElement(length),
	) == 1

	for i := 0; i < length; i++ {
		expr := node.Exprs[i]
		ast.Walk(v, expr)
		exprRegister := v.getReturnRegister()
		v.addCode(NewMove(utils.R0, exprRegister).armAssembly())
		v.addCode(NewMove(utils.R1, identRegister).armAssembly())
		v.callLibraryFunction(AL, checkArrayIndex)
		v.addCode(NewAdd(identRegister, identRegister, 4).armAssembly())

		if i == length-1 && lastIsCharOrBool {
			v.addCode(
				"ADD %s, %s, %s", identRegister,
				identRegister,
				exprRegister,
			)
		} else {
			v.addCode("ADD %s, %s, %s, LSL #2",
				identRegister,
				identRegister,
				exprRegister)
		}

		// If it is an assignment leave the Pointer to the element in the
		// register otherwise convert to value
		if !node.Pointer {
			if i == length-1 && lastIsCharOrBool {
				v.addCode(NewLoadReg(SB, identRegister, identRegister).armAssembly())
			} else {
				v.addCode(NewLoadReg(W, identRegister, identRegister).armAssembly())
			}
		}
	}

	v.returnRegisters.Push(identRegister)
}

/**************** STRUCT ELEMENT NODE ****************/

func (v *CodeGenerator) visitStructElementNode(node *ast.StructElementNode) {
	ast.Walk(v, node.Struct)

	register := v.returnRegisters.Peek()
	v.addCode(NewMove(utils.R0, register).armAssembly())
	v.callLibraryFunction(AL, checkNullPointer)
	v.addCode(
		"ADD %s, %s, #%d",
		register,
		register,
		node.StructType.MemoryOffset,
	)
	// If we don't want a Pointer then don't retrieve the value
	if !node.Pointer {
		v.addCode(
			NewLoadReg(
				load(ast.SizeOf(node.StructType.T)),
				register,
				register,
			).armAssembly())
	}
}

/**************** ARRAY LITERAL NODE ****************/

func (v *CodeGenerator) visitArrayLiteralNode(node *ast.ArrayLiteralNode) {
	register := v.getFreeRegister()
	length := len(node.Exprs)
	size := 0
	if length > 0 {
		size = ast.SizeOf(ast.Type(node.Exprs[0], v.symbolTable))
	}
	v.addCode(NewLoad(W, utils.R0, ConstantAddress(length*size+4)).armAssembly())
	v.addCode(NewBranchL("malloc").armAssembly())
	v.addCode(NewMove(register, utils.R0).armAssembly())
	for i := 0; i < length; i++ {
		ast.Walk(v, node.Exprs[i])
		exprRegister := v.getReturnRegister()
		v.addCode(NewStoreRegOffset(
			store(size),
			exprRegister,
			register,
			4+i*size,
		).armAssembly())
	}
	lengthRegister := v.freeRegisters.Peek()
	v.addCode(NewLoad(W, lengthRegister, ConstantAddress(length)).armAssembly())
	v.addCode(NewStoreReg(W, lengthRegister, register).armAssembly())
}

/**************** STRUCT NEW NODE ****************/

func (v *CodeGenerator) visitStructNewNode(node *ast.StructNewNode) {
	register := v.getFreeRegister()

	v.addCode(NewLoad(
		W,
		utils.R0,
		ConstantAddress(node.StructNode.MemorySize),
	).armAssembly())
	v.addCode(NewBranchL("malloc").armAssembly())
	v.addCode(NewMove(register, utils.R0).armAssembly())
	for index, n := range node.StructNode.Types {
		ast.Walk(v, node.Exprs[index])
		exprRegister := v.getReturnRegister()
		v.addCode(NewStoreRegOffset(
			store(ast.SizeOf(n.T)),
			exprRegister,
			register,
			n.MemoryOffset,
		).armAssembly())
	}
}

/**************** NEW PAIR NODE ****************/

func (v *CodeGenerator) visitNewPairNode(node *ast.NewPairNode) {
	register := v.getFreeRegister()

	// Make space for 2 new pointers on heap
	v.addCode(NewLoad(W, utils.R0, ConstantAddress(8)).armAssembly())
	v.addCode(NewBranchL("malloc").armAssembly())
	v.addCode(NewMove(register, utils.R0).armAssembly())

	// Store first element
	ast.Walk(v, node.Fst)
	fst := v.getReturnRegister()
	fstSize := ast.SizeOf(ast.Type(node.Fst, v.symbolTable))
	v.addCode(NewLoad(W, utils.R0, ConstantAddress(fstSize)).armAssembly())
	v.addCode(NewBranchL("malloc").armAssembly())
	v.addCode(NewStoreReg(W, utils.R0, register).armAssembly())
	v.addCode(NewStoreReg(store(fstSize), fst, utils.R0).armAssembly())

	// Store second element
	ast.Walk(v, node.Snd)
	snd := v.getReturnRegister()
	sndSize := ast.SizeOf(ast.Type(node.Snd, v.symbolTable))
	v.addCode(NewLoad(W, utils.R0, ConstantAddress(sndSize)).armAssembly())
	v.addCode(NewBranchL("malloc").armAssembly())
	v.addCode(NewStoreRegOffset(W, utils.R0, register, 4).armAssembly())
	v.addCode(NewStoreReg(store(sndSize), snd, utils.R0).armAssembly())
}

/**************** FUNCTION CALL NODE ****************/

func (v *CodeGenerator) visitFunctionCallNode(node *ast.FunctionCallNode) {
	registers := utils.ReturnRegisters()
	size := 0
	for i := len(node.Exprs) - 1; i >= 0; i-- {
		ast.Walk(v, node.Exprs[i])
		register := v.getReturnRegister()
		if i < len(registers) {
			v.addCode(NewMove(registers[i], register).armAssembly())
		} else {
			f, _ := v.symbolTable.SearchForFunction(node.Ident.Ident, node.Exprs)
			v.subtractFromStackPointer(ast.SizeOf(f.Params[i].T))
			v.addCode(NewStoreReg(
				store(ast.SizeOf(f.Params[i].T)),
				register,
				utils.SP,
			).armAssembly())
			size += ast.SizeOf(f.Params[i].T)
		}
	}
	var buff bytes.Buffer
	for _, e := range node.Exprs {
		buff.WriteString(ast.Type(e, v.symbolTable).Hash())
	}
	functionLabel := fmt.Sprintf("f%s_%s", buff.String(), node.Ident.Ident)
	v.addCode(NewBranchL(functionLabel).armAssembly())
	if size > 0 {
		v.addToStackPointer(size)
	}

	v.addCode(NewMove(v.getFreeRegister(), utils.R0).armAssembly())
}

/**************** INTEGER LITERAL NODE ****************/

func (v *CodeGenerator) visitIntegerLiteralNode(node *ast.IntegerLiteralNode) {
	v.addCode(NewLoad(W, v.getFreeRegister(), ConstantAddress(node.Val)).armAssembly())
}

/**************** BOOLEAN LITERAL NODE ****************/

func (v *CodeGenerator) visitBooleanLiteralNode(node *ast.BooleanLiteralNode) {
	register := v.getFreeRegister()
	if node.Val {
		// True
		v.addCode(NewMoveInt(register, 1).armAssembly())
	} else {
		// False
		v.addCode(NewMoveInt(register, 0).armAssembly())
	}
}

/**************** CHARACTER LITERAL NODE ****************/

func (v *CodeGenerator) visitCharacterLiteralNode(
	node *ast.CharacterLiteralNode,
) {
	v.addCode(NewMoveChar(v.getFreeRegister(), string(node.Val)).armAssembly())
}

/**************** STRING LITERAL NODE ****************/

func (v *CodeGenerator) visitStringLiteralNode(node *ast.StringLiteralNode) {
	label := v.addData(node.Val)
	v.addCode(NewLoad(W, v.getFreeRegister(), LabelAddress(label)).armAssembly())
}

/**************** NULL NODE ****************/

func (v *CodeGenerator) visitNullNode() {
	v.addCode(NewLoad(W, v.getFreeRegister(), ConstantAddress(0)).armAssembly())
}

/**************** BINARY OPERATOR NODE ****************/

func (v *CodeGenerator) visitBinaryOperatorNode(node *ast.BinaryOperatorNode) {
	var operand2 utils.Register
	var operand1 utils.Register

	if v.freeRegisters.Length() == 2 {
		ast.Walk(v, node.Expr2)
		operand2 = v.returnRegisters.Pop()
		v.addCode(NewPush(operand2).armAssembly())
		v.currentStackPos += ast.SizeOf(ast.Type(node.Expr1, v.symbolTable))
		v.freeRegisters.Push(operand2)

		ast.Walk(v, node.Expr1)
		operand1 = v.returnRegisters.Pop()
		operand2 = v.freeRegisters.Pop()
		v.addCode(NewPop(operand2).armAssembly())
		v.currentStackPos -= ast.SizeOf(ast.Type(node.Expr1, v.symbolTable))
	} else {
		// Evaluate the expression with the largest weight first
		if ast.Weight(node.Expr1) > ast.Weight(node.Expr2) {
			ast.Walk(v, node.Expr1)
			operand1 = v.returnRegisters.Pop()
			ast.Walk(v, node.Expr2)
			operand2 = v.returnRegisters.Pop()
		} else {
			ast.Walk(v, node.Expr2)
			operand2 = v.returnRegisters.Pop()
			ast.Walk(v, node.Expr1)
			operand1 = v.returnRegisters.Pop()
		}
	}
	v.visitBinaryOperator(node.Op, operand1, operand2)
	v.freeRegisters.Push(operand2)
	v.returnRegisters.Push(operand1)
}

func (v *CodeGenerator) visitBinaryOperator(
	op ast.BinaryOperator,
	r1 utils.Register,
	r2 utils.Register,
) {
	switch op {
	case ast.MUL:
		v.visitMul(r1, r2)
	case ast.DIV:
		v.visitDivMod("__aeabi_idiv", r1, r2, utils.R0)
	case ast.MOD:
		v.visitDivMod("__aeabi_idivmod", r1, r2, utils.R1)
	case ast.ADD:
		v.visitADD(r1, r2)
	case ast.SUB:
		v.visitSUB(r1, r2)
	case ast.GT:
		v.visitCompare(r1, r2, GT, LE)
	case ast.GEQ:
		v.visitCompare(r1, r2, GE, LT)
	case ast.LT:
		v.visitCompare(r1, r2, LT, GE)
	case ast.LEQ:
		v.visitCompare(r1, r2, LE, GT)
	case ast.EQ:
		v.visitCompare(r1, r2, EQ, NE)
	case ast.NEQ:
		v.visitCompare(r1, r2, NE, EQ)
	case ast.AND:
		v.visitAndOr(AND, r1, r2)
	case ast.OR:
		v.visitAndOr(ORR, r1, r2)
	}
}

func (v *CodeGenerator) visitMul(r1 utils.Register, r2 utils.Register) {
	v.addCode(NewSignedMultiply(r1, r2).armAssembly())
	v.addCode(NewCompareASR(r2, r1, 31).armAssembly())
	v.callLibraryFunction(NE, checkOverflow)
}

func (v *CodeGenerator) visitDivMod(
	label string,
	r1 utils.Register,
	r2 utils.Register,
	returnReg utils.Register,
) {
	v.addCode(NewMove(utils.R0, r1).armAssembly())
	v.addCode(NewMove(utils.R1, r2).armAssembly())
	v.callLibraryFunction(AL, checkDivide)
	v.addCode(NewBranchL(label).armAssembly())
	v.addCode(NewMove(r1, returnReg).armAssembly())
}

func (v *CodeGenerator) visitADD(r1 utils.Register, r2 utils.Register) {
	v.addCode("ADDS %s, %s, %s", r1, r1, r2)
	v.callLibraryFunction(VS, checkOverflow)
}

func (v *CodeGenerator) visitSUB(r1 utils.Register, r2 utils.Register) {
	v.addCode(NewSubtract(r1, r2).armAssembly())
	v.callLibraryFunction(VS, checkOverflow)
}

func (v *CodeGenerator) visitCompare(
	r1 utils.Register,
	r2 utils.Register,
	cond Condition,
	opp Condition,
) {
	v.addCode(NewCompare(r1, r2).armAssembly())
	v.addCode(NewMoveCond(cond, r1, 1).armAssembly())
	v.addCode(NewMoveCond(opp, r1, 0).armAssembly())
}

func (v *CodeGenerator) visitAndOr(
	instr LogicalInstructionType,
	r1 utils.Register,
	r2 utils.Register,
) {
	v.addCode(NewLogicalInstr(instr, r1, r1, r2).armAssembly())
}

/**************** STATEMENTS ****************/

func (v *CodeGenerator) visitStatements() {
	v.symbolTable.MoveNextScope()
	size := 0
	for _, dec := range v.symbolTable.CurrentScope.Scope {
		size += ast.SizeOf(dec.T)
		dec.AddLocation(utils.NewStackOffsetLocation(v.currentStackPos + size))
	}
	if size != 0 {
		v.subtractFromStackPointer(size)
	}
	v.symbolTable.CurrentScope.ScopeSize = size
}

func (v *CodeGenerator) leaveStatements() {
	if v.symbolTable.CurrentScope.ScopeSize != 0 {
		v.addToStackPointer(v.symbolTable.CurrentScope.ScopeSize)
	}
	v.symbolTable.MoveUpScope()
}

/**************** FREE NODE ****************/

func (v *CodeGenerator) leaveFreeNode() {
	v.addCode(NewMove(utils.R0, v.getReturnRegister()).armAssembly())
	v.callLibraryFunction(AL, free)
}

/**************** DECLARE NODE ****************/

func (v *CodeGenerator) leaveDeclareNode(node *ast.DeclareNode) {
	dec, _ := v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident)
	dec.IsDeclared = true
	if dec.Location != nil {
		v.addCode(NewStore(
			store(ast.SizeOf(dec.T)),
			v.getReturnRegister(),
			v.LocationOf(dec.Location),
		).armAssembly())
	}
}

/**************** PRINT NODE ****************/

func (v *CodeGenerator) leavePrintNode(node *ast.PrintNode) {
	v.addCode(NewMove(utils.R0, v.getReturnRegister()).armAssembly())
	v.addPrint(ast.Type(node.Expr, v.symbolTable))
}

/**************** PRINTLN NODE ****************/

func (v *CodeGenerator) leavePrintlnNode(node *ast.PrintlnNode) {
	v.addCode(NewMove(utils.R0, v.getReturnRegister()).armAssembly())
	v.addPrint(ast.Type(node.Expr, v.symbolTable))
	v.callLibraryFunction(AL, printLn)
}

/**************** EXIT NODE ****************/

func (v *CodeGenerator) leaveExitNode() {
	v.addCode(NewMove(utils.R0, v.getReturnRegister()).armAssembly())
	v.addCode(NewBranchL("exit").armAssembly())
}

/**************** RETURN NODE ****************/

func (v *CodeGenerator) leaveReturnNode() {
	sizeOfAllVariablesInScope := 0
	for scope := v.symbolTable.CurrentScope; scope !=
		v.symbolTable.Head; scope = scope.ParentScope {
		sizeOfAllVariablesInScope += scope.ScopeSize
	}
	if sizeOfAllVariablesInScope > 0 {
		// Cannot add more than 1024 from SP at once, so do it in multiple
		// iterations.

		// Not using addToStackPointer function as we want v.currentStackPos
		// to be unchanged.
		i := sizeOfAllVariablesInScope
		for ; i > 1024; i -= 1024 {
			v.addCode(NewAdd(utils.SP, utils.SP, 1024).armAssembly())
		}
		v.addCode(NewAdd(utils.SP, utils.SP, i).armAssembly())
	}
	v.addCode(NewMove(utils.R0, v.getReturnRegister()).armAssembly())
	v.addCode(NewPop(utils.PC).armAssembly())
}

/**************** PAIR FIRST ELEMENT NODE ****************/

func (v *CodeGenerator) leavePairFirstElementNode(
	node *ast.PairFirstElementNode,
) {
	register := v.returnRegisters.Peek()
	v.addCode(NewMove(utils.R0, register).armAssembly())
	v.callLibraryFunction(AL, checkNullPointer)
	v.addCode(NewLoadReg(W, register, register).armAssembly())
	// If we don't want a Pointer then don't retrieve the value
	if !node.Pointer {
		v.addCode(NewLoadReg(
			load(ast.SizeOf(ast.Type(node.Expr, v.symbolTable))),
			register,
			register,
		).armAssembly())
	}
}

/**************** PAIR SECOND ELEMENT NODE ****************/

func (v *CodeGenerator) leavePairSecondElementNode(
	node *ast.PairSecondElementNode,
) {
	register := v.returnRegisters.Peek()
	v.addCode(NewMove(utils.R0, register).armAssembly())
	v.callLibraryFunction(AL, checkNullPointer)
	v.addCode(NewLoadRegOffset(W, register, register, 4).armAssembly())

	// If we don't want a Pointer then don't retrieve the value
	if !node.Pointer {
		v.addCode(NewLoadReg(
			load(ast.SizeOf(ast.Type(node.Expr, v.symbolTable))),
			register,
			register,
		).armAssembly())
	}
}

/**************** UNARY OPERATOR NODE ****************/

func (v *CodeGenerator) leaveUnaryOperatorNode(node *ast.UnaryOperatorNode) {
	register := v.returnRegisters.Peek()
	switch node.Op {
	case ast.NOT:
		v.addCode(NewLogicalInstrInt(EOR, register, register, 1).armAssembly())
	case ast.NEG:
		v.addCode(NewNegate(register).armAssembly())
		v.callLibraryFunction(VS, checkOverflow)
	case ast.LEN:
		v.addCode(NewLoadReg(W, register, register).armAssembly())
	}
}
