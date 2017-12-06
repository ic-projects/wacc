package main

import (
	"ast"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"utils"
)

/**************** CODE GENERATOR STRUCTS ****************/

// ASCIIWord is a struct that stores the length and string of an ASCII string.
type ASCIIWord struct {
	length int
	text   string
}

// NewASCIIWord builds an ASCIIWord.
func NewASCIIWord(length int, text string) ASCIIWord {
	return ASCIIWord{
		length: length,
		text:   text,
	}
}

// Assembly is a struct that stores the different parts of the assembly. It
// stores the .data, .text and global.
type Assembly struct {
	data        map[string](ASCIIWord)
	dataCounter int
	text        []string
	global      map[string]([]string)
}

// NewAssembly builds an Assembly object.
func NewAssembly() *Assembly {
	return &Assembly{
		data:        make(map[string]ASCIIWord),
		dataCounter: 0,
		text:        make([]string, 0),
		global:      make(map[string]([]string)),
	}
}

// String will return the string format of the Assembly code.
func (asm *Assembly) String() string {
	var buf bytes.Buffer
	buf.WriteString(".data\n")
	for dname, d := range asm.data {
		buf.WriteString(dname + ":\n")
		buf.WriteString(fmt.Sprintf("   .word %d\n", d.length))
		buf.WriteString(fmt.Sprintf("   .ascii \"%s\"\n", d.text))
	}
	buf.WriteString(".text\n")
	for _, s := range asm.text {
		buf.WriteString(utils.Indent(s, "  "))
	}
	buf.WriteString(".global main\n")
	for fname, f := range asm.global {
		buf.WriteString(fname + ":\n")
		for _, s := range f {
			buf.WriteString(utils.Indent(s, "  "))
		}
	}
	return buf.String()
}

// NumberedCode will return the string format of the Assembly code with line
// numbers.
func (asm *Assembly) NumberedCode() string {
	var buf bytes.Buffer
	for i, line := range strings.Split(asm.String(), "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%d\t%s\n", i, line))
		}
	}
	return buf.String()
}

// SaveToFile is a function that will save the assembly to the given savepath
// overwriting any file already there.
func (asm *Assembly) SaveToFile(savepath string) error {
	file, err := os.Create(savepath)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	if _, err = w.WriteString(asm.String()); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return file.Close()
}

// LocationOf will return the location of a
func (v *CodeGenerator) LocationOf(loc *utils.Location) string {
	// Location is a register
	if loc.Register != utils.UNDEFINED {
		return loc.Register.String()
	}

	// Location is an address on the heap
	if loc.Address != 0 {
		return "#" + strconv.Itoa(loc.Address)
	}

	// Location is a stack offset
	return "[sp, #" +
		strconv.Itoa(v.currentStackPos-loc.CurrentPos) +
		"]"
}

// PointerTo returns a string, that when added gives the object's location in
// memory
func (v *CodeGenerator) PointerTo(location *utils.Location) string {
	return "sp, #" +
		strconv.Itoa(v.currentStackPos-location.CurrentPos)
}

// GenerateCode is a function that will generate and return the finished
// assembly code for a given AST.
func GenerateCode(
	tree ast.ProgramNode,
	symbolTable *ast.SymbolTable,
) *Assembly {
	codeGen := NewCodeGenerator(symbolTable)

	ast.Walk(codeGen, tree)

	return codeGen.asm
}

// CodeGenerator is a struct that implements EntryExitVisitor to be called with
// Walk. It stores the assembly, a label count (used for ensuring distinct
// labels) the currentFunction (so when we add code we add to the
// currentFunction), a symbolTable, a stack of freeRegisters and
// returnRegisters, a predefined function Library, and the current position of
// the stack.
type CodeGenerator struct {
	asm             *Assembly
	labelCount      int
	currentFunction string
	symbolTable     *ast.SymbolTable
	freeRegisters   *utils.RegisterStack
	returnRegisters *utils.RegisterStack
	library         *Library
	currentStackPos int
}

// NewCodeGenerator returns an initialised CodeGenerator.
func NewCodeGenerator(symbolTable *ast.SymbolTable) *CodeGenerator {
	return &CodeGenerator{
		asm:             NewAssembly(),
		labelCount:      0,
		symbolTable:     symbolTable,
		freeRegisters:   utils.NewRegisterStackWith(utils.FreeRegisters()),
		returnRegisters: utils.NewRegisterStack(),
		library:         GetLibrary(),
		currentStackPos: 0,
	}
}

// addPrint will add the correct type of print function for the type given.
func (v *CodeGenerator) addPrint(t ast.TypeNode) {
	switch node := ast.ToValue(t).(type) {
	case ast.BaseTypeNode:
		switch node.T {
		case ast.BOOL:
			v.callLibraryFunction("BL", printBool)
		case ast.INT:
			v.callLibraryFunction("BL", printInt)
		case ast.CHAR:
			v.addCode("BL putchar")
		case ast.PAIR:
			v.callLibraryFunction("BL", printReference)
		}
	case ast.ArrayTypeNode:
		if arr, ok := ast.ToValue(node.T).(ast.BaseTypeNode); ok {
			if arr.T == ast.CHAR {
				v.callLibraryFunction("BL", printString)
				return
			}
		}
		v.callLibraryFunction("BL", printReference)
	case ast.PairTypeNode, ast.StructTypeNode, ast.NullTypeNode:
		v.callLibraryFunction("BL", printReference)
	}
}

// addDataWithLabel adds a ascii word to the data section generating a unique
// label
func (v *CodeGenerator) addData(text string) string {
	label := "msg_" + strconv.Itoa(v.asm.dataCounter)
	v.asm.dataCounter++
	v.addDataWithLabel(label, text)
	return label
}

// addDataWithLabel adds a ascii word to the data section using a given label
func (v *CodeGenerator) addDataWithLabel(label string, text string) {
	length := 0
	for i := 0; i < len(text); i++ {
		length++
		if text[i] == '\\' {
			i++
		}
	}
	v.asm.data[label] = NewASCIIWord(length, text)
}

// addCode formats and adds one line of assembly to the correct location in then
// generated assembly code.
func (v *CodeGenerator) addCode(lineFormat string, inputs ...interface{}) {
	v.asm.global[v.currentFunction] = append(v.asm.global[v.currentFunction],
		fmt.Sprintf(lineFormat+"\n", inputs...))
}

// addCode add lines of assembly to the already code part of the generated
// assembly code.
func (v *CodeGenerator) addFunction(name string) {
	v.asm.global[name] = make([]string, 0)
	v.currentFunction = name
}

// callLibraryFunction adds the corresponding predefined function to the
// assembly if it is not already added. It also adds the call to the function to
// the assembly, using call as the instruction before the label.
func (v *CodeGenerator) callLibraryFunction(
	call string,
	function LibraryFunction,
) {
	v.addCode("%s %s", call, function)
	v.library.add(v, function)
}

/**************** WALKER FUNCTIONS ****************/

// NoRecurse defines the nodes of the AST which should not be automatically
// walked. This means we can Walk the children in any way we choose.
func (v *CodeGenerator) NoRecurse(programNode ast.ProgramNode) bool {
	switch programNode.(type) {
	case *ast.IfNode,
		*ast.SwitchNode,
		*ast.AssignNode,
		*ast.ArrayLiteralNode,
		*ast.ArrayElementNode,
		*ast.LoopNode,
		*ast.ForLoopNode,
		*ast.NewPairNode,
		*ast.StructNewNode,
		*ast.ReadNode,
		*ast.FunctionCallNode,
		*ast.BinaryOperatorNode:
		return true
	default:
		return false
	}
}

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *CodeGenerator) Visit(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case *ast.FunctionNode:
		v.currentStackPos = 0
		v.freeRegisters = utils.NewRegisterStackWith(utils.FreeRegisters())
		v.symbolTable.MoveNextScope()
		if node.Ident.Ident == "" {
			v.addFunction("main")
		} else {
			v.addFunction("f_" + node.Ident.Ident)
		}
		v.addCode("PUSH {lr}")
	case []*ast.ParameterNode:
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
				v.addCode("%s %s, %s",
					store(ast.SizeOf(e.T)),
					registers[n],
					v.LocationOf(dec.Location))
			}
		}
	case *ast.AssignNode:
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
			if ast.SizeOf(arr.GetDimElement(len(lhsNode.Exprs))) == 1 {
				v.addCode("STRB %s, [%s]", rhsRegister, lhsRegister)
			} else {
				v.addCode("STR %s, [%s]", rhsRegister, lhsRegister)
			}
		case *ast.StructElementNode:
			ast.Walk(v, lhsNode)
			lhsRegister := v.getReturnRegister()
			v.addCode("%s %s, [%s]",
				store(ast.SizeOf(lhsNode.StructType.T)),
				rhsRegister,
				lhsRegister)
		case *ast.PairFirstElementNode:
			ast.Walk(v, lhsNode)
			lhsRegister := v.getReturnRegister()
			v.addCode("%s %s, [%s]",
				store(ast.SizeOf(ast.Type(lhsNode.Expr, v.symbolTable))),
				rhsRegister,
				lhsRegister)
		case *ast.PairSecondElementNode:
			ast.Walk(v, lhsNode)
			lhsRegister := v.getReturnRegister()
			v.addCode("%s %s, [%s]",
				store(ast.SizeOf(ast.Type(lhsNode.Expr, v.symbolTable))),
				rhsRegister,
				lhsRegister)
		case *ast.IdentifierNode:
			ident := v.symbolTable.SearchForDeclaredIdent(lhsNode.Ident)
			if ident.Location != nil {
				v.addCode("%s %s, %s",
					store(ast.SizeOf(ident.T)),
					rhsRegister,
					v.LocationOf(ident.Location))
			}
		}
		v.freeRegisters.Push(rhsRegister)
	case *ast.ReadNode:
		if ident, ok := node.LHS.(*ast.IdentifierNode); ok {
			dec := v.symbolTable.SearchForDeclaredIdent(ident.Ident)
			v.addCode(
				"ADD %s, %s", v.getFreeRegister(),
				v.PointerTo(dec.Location),
			)
		} else {
			ast.Walk(v, node.LHS)
		}
		v.addCode("MOV r0, %s", v.getReturnRegister())
		if ast.SizeOf(ast.Type(node.LHS, v.symbolTable)) == 1 {
			v.callLibraryFunction("BL", readChar)
		} else {
			v.callLibraryFunction("BL", readInt)
		}
	case *ast.IfNode:
		// Labels
		elseLabel := v.labelCount + 1
		endifLabel := v.labelCount + 1
		v.labelCount += 2
		// Cond
		ast.Walk(v, node.Expr)
		v.addCode("CMP %s, #0", v.getReturnRegister())
		v.addCode("BEQ ELSE%d", elseLabel)
		// If
		ast.Walk(v, node.IfStats)
		v.addCode("B ENDIF%d", endifLabel)
		// Else
		v.addCode("ELSE%d:", elseLabel)
		ast.Walk(v, node.ElseStats)
		// Fi
		v.addCode("ENDIF%d:", endifLabel)
	case *ast.SwitchNode:
		// Labels
		endLabel := v.labelCount + 1
		v.labelCount++
		ast.Walk(v, node.Expr)
		r := v.returnRegisters.Pop()
		for i, c := range node.Cases {
			if !c.IsDefault {
				ast.Walk(v, c.Expr)
				v.addCode("CMP %s, %s", v.getReturnRegister(), r)
				v.addCode("BEQ CASE%d_%d", endLabel, i)
			} else {
				v.addCode("B CASE%d_%d", endLabel, i)
			}
		}
		v.addCode("B END%d", endLabel)
		v.freeRegisters.Push(r)
		for i, c := range node.Cases {
			v.addCode("CASE%d_%d:", endLabel, i)
			ast.Walk(v, c.Stats)
			v.addCode("B END%d", endLabel)
		}
		v.addCode("END%d:", endLabel)
	case *ast.LoopNode:
		// Labels
		doLabel := v.labelCount + 1
		whileLabel := v.labelCount + 1
		v.labelCount += 2
		v.addCode("B WHILE%d", whileLabel)
		// Do
		v.addCode("DO%d:", doLabel)
		ast.Walk(v, node.Stats)
		// While
		v.addCode("WHILE%d:", whileLabel)
		ast.Walk(v, node.Expr)
		v.addCode("CMP %s, #1", v.getReturnRegister())
		v.addCode("BEQ DO%d", doLabel)
	case *ast.ForLoopNode:
		ast.Walk(v, node.Initial)
		// Labels
		doLabel := v.labelCount + 1
		whileLabel := v.labelCount + 1
		v.labelCount += 2
		v.addCode("B WHILE%d", whileLabel)
		// Do
		v.addCode("DO%d:", doLabel)
		ast.Walk(v, node.Stats)
		ast.Walk(v, node.Update)
		// While
		v.addCode("WHILE%d:", whileLabel)
		ast.Walk(v, node.Expr)
		v.addCode("CMP %s, #1", v.getReturnRegister())
		v.addCode("BEQ DO%d", doLabel)
	case *ast.IdentifierNode:
		dec := v.symbolTable.SearchForDeclaredIdent(node.Ident)
		v.addCode("%s %s, %s",
			load(ast.SizeOf(dec.T)),
			v.getFreeRegister(),
			v.LocationOf(dec.Location))
	case *ast.ArrayElementNode:
		ast.Walk(v, node.Ident)
		identRegister := v.returnRegisters.Pop()

		length := len(node.Exprs)
		symbol := v.symbolTable.SearchForDeclaredIdent(node.Ident.Ident)
		lastIsCharOrBool := ast.SizeOf(ast.ToValue(symbol.T).(ast.ArrayTypeNode).GetDimElement(length)) == 1

		for i := 0; i < length; i++ {
			expr := node.Exprs[i]
			ast.Walk(v, expr)
			exprRegister := v.getReturnRegister()
			v.addCode("MOV r0, %s", exprRegister)
			v.addCode("MOV r1, %s", identRegister)
			v.callLibraryFunction("BL", checkArrayIndex)
			v.addCode("ADD %s, %s, #4", identRegister, identRegister)

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
					v.addCode("LDRSB %s, [%s]", identRegister, identRegister)
				} else {
					v.addCode("LDR %s, [%s]", identRegister, identRegister)
				}
			}
		}

		v.returnRegisters.Push(identRegister)
	case *ast.StructElementNode:
		ast.Walk(v, node.Struct)

		register := v.returnRegisters.Peek()
		v.addCode("MOV r0, %s", register)
		v.callLibraryFunction("BL", checkNullPointer)
		v.addCode(
			"ADD %s, %s, #%d",
			register,
			register,
			node.StructType.MemoryOffset,
		)
		// If we don't want a Pointer then don't retrieve the value
		if !node.Pointer {
			v.addCode("%s %s, [%s]",
				load(ast.SizeOf(node.StructType.T)),
				register,
				register)
		}
	case *ast.ArrayLiteralNode:
		register := v.getFreeRegister()
		length := len(node.Exprs)
		size := 0
		if length > 0 {
			size = ast.SizeOf(ast.Type(node.Exprs[0], v.symbolTable))
		}
		v.addCode("LDR r0, =%d", length*size+4)
		v.addCode("BL malloc")
		v.addCode("MOV %s, r0", register)
		for i := 0; i < length; i++ {
			ast.Walk(v, node.Exprs[i])
			exprRegister := v.getReturnRegister()
			v.addCode("%s %s, [%s, #%d]",
				store(size),
				exprRegister,
				register,
				4+i*size)
		}
		lengthRegister := v.freeRegisters.Peek()
		v.addCode("LDR %s, =%d", lengthRegister, length)
		v.addCode("STR %s, [%s]", lengthRegister, register)
	case *ast.StructNewNode:
		register := v.getFreeRegister()

		v.addCode("LDR r0, =%d", node.StructNode.MemorySize)
		v.addCode("BL malloc")
		v.addCode("MOV %s, r0", register)
		for index, n := range node.StructNode.Types {
			ast.Walk(v, node.Exprs[index])
			exprRegister := v.getReturnRegister()
			v.addCode("%s %s, [%s, #%d]",
				store(ast.SizeOf(n.T)),
				exprRegister,
				register,
				n.MemoryOffset,
			)
		}
	case *ast.NewPairNode:
		register := v.getFreeRegister()

		// Make space for 2 new pointers on heap
		v.addCode("LDR r0, =8")
		v.addCode("BL malloc")
		v.addCode("MOV %s, r0", register)

		// Store first element
		ast.Walk(v, node.Fst)
		fst := v.getReturnRegister()
		fstSize := ast.SizeOf(ast.Type(node.Fst, v.symbolTable))
		v.addCode("LDR r0, =%d", fstSize)
		v.addCode("BL malloc")
		v.addCode("STR r0, [%s]", register)
		v.addCode("%s %s, [r0]", store(fstSize), fst)

		// Store second element
		ast.Walk(v, node.Snd)
		snd := v.getReturnRegister()
		sndSize := ast.SizeOf(ast.Type(node.Snd, v.symbolTable))
		v.addCode("LDR r0, =%d", sndSize)
		v.addCode("BL malloc")
		v.addCode("STR r0, [%s, #4]", register)
		v.addCode("%s %s, [r0]", store(sndSize), snd)

	case *ast.FunctionCallNode:
		registers := utils.ReturnRegisters()
		size := 0
		for i := len(node.Exprs) - 1; i >= 0; i-- {
			ast.Walk(v, node.Exprs[i])
			register := v.getReturnRegister()
			if i < len(registers) {
				v.addCode("MOV %s, %s", registers[i], register)
			} else {
				f, _ := v.symbolTable.SearchForFunction(node.Ident.Ident)
				v.subtractFromStackPointer(ast.SizeOf(f.Params[i].T))
				v.addCode("%s %s, [sp]", store(ast.SizeOf(f.Params[i].T)), register)
				size += ast.SizeOf(f.Params[i].T)
			}
		}
		v.addCode("BL f_%s", node.Ident.Ident)
		if size > 0 {
			v.addToStackPointer(size)
		}

		v.addCode("MOV %s, r0", v.getFreeRegister())
	case *ast.IntegerLiteralNode:
		v.addCode("LDR %s, =%d", v.getFreeRegister(), node.Val)
	case *ast.BooleanLiteralNode:
		register := v.getFreeRegister()
		if node.Val {
			v.addCode("MOV %s, #1", register) // True
		} else {
			v.addCode("MOV %s, #0", register) // False
		}
	case *ast.CharacterLiteralNode:
		v.addCode("MOV %s, #'%s'", v.getFreeRegister(), string(node.Val))
	case *ast.StringLiteralNode:
		label := v.addData(node.Val)
		v.addCode("LDR %s, =%s", v.getFreeRegister(), label)
	case *ast.NullNode:
		v.addCode("LDR %s, =0", v.getFreeRegister())
	case *ast.UnaryOperatorNode:
		switch node.Op {
		case ast.NOT:
		case ast.NEG:
		case ast.LEN:
		case ast.ORD:
		case ast.CHR:
		}
	case *ast.BinaryOperatorNode:
		var operand2 utils.Register
		var operand1 utils.Register

		if v.freeRegisters.Length() == 2 {
			ast.Walk(v, node.Expr2)
			operand2 = v.returnRegisters.Pop()
			v.addCode("PUSH {%s}", operand2)
			v.currentStackPos += ast.SizeOf(ast.Type(node.Expr1, v.symbolTable))
			v.freeRegisters.Push(operand2)

			ast.Walk(v, node.Expr1)
			operand1 = v.returnRegisters.Pop()
			operand2 = v.freeRegisters.Pop()
			v.addCode("POP {%s}", operand2)
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
		switch node.Op {
		case ast.MUL:
			v.addCode(
				"SMULL %s, %s, %s, %s",
				operand1,
				operand2,
				operand1,
				operand2,
			)
			v.addCode("CMP %s, %s, ASR #31", operand2, operand1)
			v.callLibraryFunction("BLNE", checkOverflow)
		case ast.DIV:
			v.addCode("MOV r0, %s", operand1)
			v.addCode("MOV r1, %s", operand2)
			v.callLibraryFunction("BL", checkDivide)
			v.addCode("BL __aeabi_idiv")
			v.addCode("MOV %s, r0", operand1)
		case ast.MOD:
			v.addCode("MOV r0, %s", operand1)
			v.addCode("MOV r1, %s", operand2)
			v.callLibraryFunction("BL", checkDivide)
			v.addCode("BL __aeabi_idivmod")
			v.addCode("MOV %s, r1", operand1)
		case ast.ADD:
			v.addCode("ADDS %s, %s, %s", operand1, operand1, operand2)
			v.callLibraryFunction("BLVS", checkOverflow)
		case ast.SUB:
			v.addCode("SUBS %s, %s, %s", operand1, operand1, operand2)
			v.callLibraryFunction("BLVS", checkOverflow)
		case ast.GT:
			v.addCode("CMP %s, %s", operand1, operand2)
			v.addCode("MOVGT %s, #1", operand1)
			v.addCode("MOVLE %s, #0", operand1)
		case ast.GEQ:
			v.addCode("CMP %s, %s", operand1, operand2)
			v.addCode("MOVGE %s, #1", operand1)
			v.addCode("MOVLT %s, #0", operand1)
		case ast.LT:
			v.addCode("CMP %s, %s", operand1, operand2)
			v.addCode("MOVLT %s, #1", operand1)
			v.addCode("MOVGE %s, #0", operand1)
		case ast.LEQ:
			v.addCode("CMP %s, %s", operand1, operand2)
			v.addCode("MOVLE %s, #1", operand1)
			v.addCode("MOVGT %s, #0", operand1)
		case ast.EQ:
			v.addCode("CMP %s, %s", operand1, operand2)
			v.addCode("MOVEQ %s, #1", operand1)
			v.addCode("MOVNE %s, #0", operand1)
		case ast.NEQ:
			v.addCode("CMP %s, %s", operand1, operand2)
			v.addCode("MOVNE %s, #1", operand1)
			v.addCode("MOVEQ %s, #0", operand1)
		case ast.AND:
			v.addCode("AND %s, %s, %s", operand1, operand1, operand2)
		case ast.OR:
			v.addCode("ORR %s, %s, %s", operand1, operand1, operand2)
		}
		v.freeRegisters.Push(operand2)
		v.returnRegisters.Push(operand1)
	case []ast.StatementNode:
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
}

// Leave will be called to leave the current node.
func (v *CodeGenerator) Leave(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case []ast.StatementNode:
		if v.symbolTable.CurrentScope.ScopeSize != 0 {
			v.addToStackPointer(v.symbolTable.CurrentScope.ScopeSize)
		}
		v.symbolTable.MoveUpScope()
	case *ast.FunctionNode:
		if v.symbolTable.CurrentScope.ScopeSize > 0 {
			// Cannot add more than 1024 from SP at once, so do it in multiple
			// iterations.

			// Not using addToStackPointer function as we want v.currentStackPos
			// to be unchanged.
			i := v.symbolTable.CurrentScope.ScopeSize
			for ; i > 1024; i -= 1024 {
				v.addCode("ADD sp, sp, #1024")
			}
			v.addCode("ADD sp, sp, #%d", i)
		}
		v.symbolTable.MoveUpScope()
		if node.Ident.Ident == "" {
			v.addCode("LDR r0, =0")
			v.addCode("POP {pc}")
		}
		v.addCode(".ltorg")
	case *ast.ArrayLiteralNode:
	case *ast.FreeNode:
		v.addCode("MOV r0, %s", v.getReturnRegister())
		v.callLibraryFunction("BL", free)
	case *ast.DeclareNode:
		dec, _ := v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident)
		dec.IsDeclared = true
		if dec.Location != nil {
			v.addCode("%s %s, %s",
				store(ast.SizeOf(dec.T)),
				v.getReturnRegister(),
				v.LocationOf(dec.Location))
		}
	case *ast.PrintNode:
		v.addCode("MOV r0, %s", v.getReturnRegister())
		v.addPrint(ast.Type(node.Expr, v.symbolTable))
	case *ast.PrintlnNode:
		v.addCode("MOV r0, %s", v.getReturnRegister())
		v.addPrint(ast.Type(node.Expr, v.symbolTable))
		v.callLibraryFunction("BL", printLn)
	case *ast.ExitNode:
		v.addCode("MOV r0, %s", v.getReturnRegister())
		v.addCode("BL exit")
	case *ast.ReturnNode:
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
				v.addCode("ADD sp, sp, #1024")
			}
			v.addCode("ADD sp, sp, #%d", i)
		}
		v.addCode("MOV r0, %s", v.getReturnRegister())
		v.addCode("POP {pc}")
	case *ast.PairFirstElementNode:
		register := v.returnRegisters.Peek()
		v.addCode("MOV r0, %s", register)
		v.callLibraryFunction("BL", checkNullPointer)
		v.addCode("LDR %s, [%s]", register, register)
		// If we don't want a Pointer then don't retrieve the value
		if !node.Pointer {
			v.addCode("%s %s, [%s]",
				load(ast.SizeOf(ast.Type(node.Expr, v.symbolTable))),
				register,
				register)
		}
	case *ast.PairSecondElementNode:
		register := v.returnRegisters.Peek()
		v.addCode("MOV r0, %s", register)
		v.callLibraryFunction("BL", checkNullPointer)
		v.addCode("LDR %s, [%s, #4]", register, register)

		// If we don't want a Pointer then don't retrieve the value
		if !node.Pointer {
			v.addCode("%s %s, [%s]",
				load(ast.SizeOf(ast.Type(node.Expr, v.symbolTable))),
				register,
				register)
		}
	case *ast.UnaryOperatorNode:
		register := v.returnRegisters.Peek()
		switch node.Op {
		case ast.NOT:
			v.addCode("EOR %s, %s, #1", register, register)
		case ast.NEG:
			v.addCode("RSBS %s, %s, #0", register, register)
			v.callLibraryFunction("BLVS", checkOverflow)
		case ast.LEN:
			v.addCode("LDR %s, [%s]", register, register)
		case ast.ORD:
		case ast.CHR:
		}
	}
}

/**************** HELPER FUNCTIONS ****************/

// addToStackPointer increments the stack pointer by the size parameter.
// If size is greater than 1024 then it will increment in multiple iterations.
func (v *CodeGenerator) addToStackPointer(size int) {
	// Cannot add more than 1024 from SP at once, so do it in multiple
	// iterations.
	i := size
	v.currentStackPos -= i
	for ; i > 1024; i -= 1024 {
		v.addCode("ADD sp, sp, #1024")
	}
	v.addCode("ADD sp, sp, #%d", i)
}

// subtractFromStackPointer decrements the stack pointer by the size parameter.
// If size is greater than 1024 then it will decrement in multiple iterations.
func (v *CodeGenerator) subtractFromStackPointer(size int) {
	// Cannot subtract more than 1024 from SP at once, so do it in multiple
	// iterations.
	i := size
	v.currentStackPos += size
	for ; i > 1024; i -= 1024 {
		v.addCode("SUB sp, sp, #1024")
	}
	v.addCode("SUB sp, sp, #%d", i)
}

// getFreeRegister pops a register from freeRegisters, and returns it
// after pushing it onto the returnRegisters.
func (v *CodeGenerator) getFreeRegister() utils.Register {
	register := v.freeRegisters.Pop()
	v.returnRegisters.Push(register)
	return register
}

// getFreeRegister pops a register from returnRegister, and returns it
// after pushing it onto the freeRegisters.
func (v *CodeGenerator) getReturnRegister() utils.Register {
	register := v.returnRegisters.Pop()
	v.freeRegisters.Push(register)
	return register
}

// store will return "STRB" if the passed paramater is one, "STR" otherwise.
func store(size int) string {
	if size == 1 {
		return "STRB"
	}
	return "STR"
}

// load will return "LDRB" if the passed paramater is one, "LDR" otherwise.
func load(size int) string {
	if size == 1 {
		return "LDRSB"
	}
	return "LDR"
}
