package codegen

import (
	"ast"
	"bufio"
	"bytes"
	"fmt"
	"location"
	"os"
	"strconv"
	"strings"
	"utils"
)

type AsciiWord struct {
	length int
	text   string
}

func NewAsciiWord(length int, text string) AsciiWord {
	return AsciiWord{
		length: length,
		text:   text,
	}
}

type Assembly struct {
	data        map[string](AsciiWord)
	dataCounter int
	text        []string
	global      map[string]([]string)
}

func NewAssembly() *Assembly {
	return &Assembly{
		data:        make(map[string]AsciiWord),
		dataCounter: 0,
		text:        make([]string, 0),
		global:      make(map[string]([]string)),
	}
}

// String will return the string format of the Assembly code with line numbers.
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
	defer file.Close()
	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	defer w.Flush()
	_, err = w.WriteString(asm.String())
	if err != nil {
		return err
	}

	return nil
}

func (codeGenerator *CodeGenerator) LocationOf(loc *location.Location) string {
	// Location is a register
	if loc.Register != location.UNDEFINED {
		return loc.Register.String()
	}

	// Location is an address on the heap
	if loc.Address != 0 {
		return "#" + strconv.Itoa(loc.Address)
	}

	// Location is a stack offset
	return "[sp, #" + strconv.Itoa(codeGenerator.currentStackPos-loc.CurrentPos) + "]"
}

func (codeGenerator *CodeGenerator) PointerTo(location *location.Location) string {
	return "sp, #" + strconv.Itoa(codeGenerator.currentStackPos-location.CurrentPos)
}

// GenerateCode is a function that will generate and return the finished assembly
// code for a given AST.
func GenerateCode(tree ast.ProgramNode, symbolTable *ast.SymbolTable) *Assembly {
	codeGen := NewCodeGenerator(symbolTable)

	ast.Walk(codeGen, tree)

	return codeGen.asm
}

// CodeGenerator is a struct that implements EntryExitVisitor to be called with
// Walk. It stores
type CodeGenerator struct {
	asm             *Assembly
	labelCount      int
	currentFunction string
	symbolTable     *ast.SymbolTable
	freeRegisters   *location.RegisterStack
	returnRegisters *location.RegisterStack
	library         *Library
	currentStackPos int
}

// NewCodeGenerator returns an initialised CodeGenerator
func NewCodeGenerator(symbolTable *ast.SymbolTable) *CodeGenerator {
	return &CodeGenerator{
		asm:             NewAssembly(),
		labelCount:      0,
		symbolTable:     symbolTable,
		freeRegisters:   location.NewRegisterStackWith(location.FreeRegisters()),
		returnRegisters: location.NewRegisterStack(),
		library:         GetLibrary(),
		currentStackPos: 0,
	}
}

func (v *CodeGenerator) addPrint(t ast.TypeNode) {
	switch node := t.(type) {
	case ast.BaseTypeNode:
		switch node.T {
		case ast.BOOL:
			v.addCode("BL " + PRINT_BOOL.String())
			v.usesFunction(PRINT_BOOL)
		case ast.INT:
			v.addCode("BL " + PRINT_INT.String())
			v.usesFunction(PRINT_INT)
		case ast.CHAR:
			v.addCode("BL putchar")
		case ast.PAIR:
			v.addCode("BL " + PRINT_REFERENCE.String())
			v.usesFunction(PRINT_REFERENCE)
		}
	case ast.ArrayTypeNode:
		if arr, ok := node.T.(ast.BaseTypeNode); ok {
			if arr.T == ast.CHAR && node.Dim == 1 {
				v.addCode("BL " + PRINT_STRING.String())
				v.usesFunction(PRINT_STRING)
				return
			}
		}
		v.addCode("BL " + PRINT_REFERENCE.String())
		v.usesFunction(PRINT_REFERENCE)
	case ast.PairTypeNode:
		v.addCode("BL " + PRINT_REFERENCE.String())
		v.usesFunction(PRINT_REFERENCE)
	}
}

// addDataWithLabel adds a ascii word to the data section generating a unique label
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
	v.asm.data[label] = NewAsciiWord(length, text)
}

// addText add lines of assembly to the already text part of the generated
// assembly code
func (v *CodeGenerator) addText(lines ...string) {
	for _, line := range lines {
		v.asm.text = append(v.asm.text, line+"\n")
	}
}

// addCode add lines of assembly to the already code part of the generated
// assembly code
func (v *CodeGenerator) addCode(lines ...string) {
	for _, line := range lines {
		v.asm.global[v.currentFunction] = append(v.asm.global[v.currentFunction], line+"\n")
	}
}

// addCode add lines of assembly to the already code part of the generated
// assembly code
func (v *CodeGenerator) addFunction(name string) {
	v.asm.global[name] = make([]string, 0)
	v.currentFunction = name
}

// usesFunction adds the corresponding predefined function to the assembly if
// it is not already added
func (v *CodeGenerator) usesFunction(f LibraryFunction) {
	v.library.add(v, f)
}

// NoRecurse defines the nodes of the AST which should not be automatically
// walked. This means we can Walk the children in any way we choose.
func (v *CodeGenerator) NoRecurse(programNode ast.ProgramNode) bool {
	switch programNode.(type) {
	case ast.IfNode,
		ast.AssignNode,
		ast.ArrayLiteralNode,
		ast.ArrayElementNode,
		ast.LoopNode,
		ast.NewPairNode,
		ast.ReadNode,
		ast.FunctionCallNode,
		ast.BinaryOperatorNode:
		return true
	default:
		return false
	}
}

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *CodeGenerator) Visit(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case ast.Program:

	case ast.FunctionNode:
		v.currentStackPos = 0
		v.freeRegisters = location.NewRegisterStackWith(location.FreeRegisters())
		v.symbolTable.MoveNextScope()
		if node.Ident.Ident == "" {
			v.addFunction("main")
		} else {
			v.addFunction("f_" + node.Ident.Ident)
		}
		v.addCode("PUSH {lr}")
	case []ast.ParameterNode:
		registers := location.ReturnRegisters()
		i := 0
		j := 0
		for n, e := range node {
			dec, _ := v.symbolTable.SearchForIdent(e.Ident.Ident)
			dec.IsDeclared = true
			if n < len(registers) {
				i += ast.SizeOf(e.T)
				dec.AddLocation(location.NewStackOffsetLocation(i))
			} else {
				dec.AddLocation(location.NewStackOffsetLocation(j - ast.SizeOf(ast.NewBaseTypeNode(ast.INT))))
				j -= ast.SizeOf(e.T)
			}
		}

		if i > 0 {
			v.addCode("SUB sp, sp, #" + strconv.Itoa(i))
		}
		v.symbolTable.CurrentScope.ScopeSize = i
		v.currentStackPos += i
		for n, e := range node {
			dec, _ := v.symbolTable.SearchForIdent(e.Ident.Ident)
			if n < len(registers) {
				v.addCode(store(ast.SizeOf(e.T), registers[n].String(), v.LocationOf(dec.Location)))
			}
		}
	case ast.ParameterNode:

	case ast.SkipNode:

	case ast.DeclareNode:

	case ast.AssignNode:
		// Rhs
		ast.Walk(v, node.Rhs)
		rhsRegister := v.returnRegisters.Pop()
		// Lhs
		switch lhsNode := node.Lhs.(type) {
		case ast.ArrayElementNode:
			ast.Walk(v, lhsNode)
			lhsRegister := v.getReturnRegister()
			dec := v.symbolTable.SearchForDeclaredIdent(lhsNode.Ident.Ident)
			arr := dec.T.(ast.ArrayTypeNode)

			if len(lhsNode.Exprs) == arr.Dim {
				v.addCode(store(ast.SizeOf(arr.T), rhsRegister.String(), "["+lhsRegister.String()+"]"))
			}
		case ast.PairFirstElementNode:
			ast.Walk(v, lhsNode)
			lhsRegister := v.getReturnRegister()
			v.addCode(store(ast.SizeOf(ast.Type(lhsNode.Expr, v.symbolTable)), rhsRegister.String(), "["+lhsRegister.String()+"]"))
		case ast.PairSecondElementNode:
			ast.Walk(v, lhsNode)
			lhsRegister := v.getReturnRegister()
			v.addCode(store(ast.SizeOf(ast.Type(lhsNode.Expr, v.symbolTable)), rhsRegister.String(), "["+lhsRegister.String()+"]"))
		case ast.IdentifierNode:
			ident := v.symbolTable.SearchForDeclaredIdent(lhsNode.Ident)
			if ident.Location != nil {
				v.addCode(store(ast.SizeOf(ident.T), rhsRegister.String(), v.LocationOf(ident.Location)))
			}
		}
		v.freeRegisters.Push(rhsRegister)
	case ast.ReadNode:

		if ident, ok := node.Lhs.(ast.IdentifierNode); ok {
			dec := v.symbolTable.SearchForDeclaredIdent(ident.Ident)
			v.addCode("ADD " + v.getFreeRegister().String() + ", " + v.PointerTo(dec.Location))
		} else {
			ast.Walk(v, node.Lhs)
		}
		v.addCode("MOV r0, " + v.getReturnRegister().String())
		if ast.SizeOf(ast.Type(node.Lhs, v.symbolTable)) == 1 {
			v.addCode("BL " + READ_CHAR.String())
			v.usesFunction(READ_CHAR)
		} else {
			v.addCode("BL " + READ_INT.String())
			v.usesFunction(READ_INT)
		}
	case ast.FreeNode:

	case ast.ReturnNode:

	case ast.ExitNode:

	case ast.PrintNode:

	case ast.PrintlnNode:

	case ast.IfNode:
		// Labels
		elseLabel := v.labelCount + 1
		endifLabel := v.labelCount + 1
		v.labelCount += 2
		// Cond
		ast.Walk(v, node.Expr)
		v.addCode(fmt.Sprintf("CMP %s, #0", v.getReturnRegister()))
		v.addCode(fmt.Sprintf("BEQ ELSE%d", elseLabel))
		// If
		ast.Walk(v, node.IfStats)
		v.addCode(fmt.Sprintf("B ENDIF%d", endifLabel))
		// Else
		v.addCode(fmt.Sprintf("ELSE%d:", elseLabel))
		ast.Walk(v, node.ElseStats)
		// Fi
		v.addCode(fmt.Sprintf("ENDIF%d:", endifLabel))
	case ast.LoopNode:
		// Labels
		doLabel := v.labelCount + 1
		whileLabel := v.labelCount + 1
		v.labelCount += 2
		v.addCode(fmt.Sprintf("B WHILE%d", whileLabel))
		// Do
		v.addCode(fmt.Sprintf("DO%d:", doLabel))
		v.labelCount++
		ast.Walk(v, node.Stats)
		// While
		v.addCode(fmt.Sprintf("WHILE%d:", whileLabel))
		v.labelCount++
		ast.Walk(v, node.Expr)
		v.addCode(fmt.Sprintf("CMP %s, #1", v.getReturnRegister()))
		v.addCode(fmt.Sprintf("BEQ DO%d", doLabel))
	case ast.ScopeNode:

	case ast.IdentifierNode:
		dec := v.symbolTable.SearchForDeclaredIdent(node.Ident)
		v.addCode(load(ast.SizeOf(dec.T), v.getFreeRegister().String(), v.LocationOf(dec.Location)))
	case ast.PairFirstElementNode:

	case ast.PairSecondElementNode:

	case ast.ArrayElementNode:
		ast.Walk(v, node.Ident)
		identRegister := v.returnRegisters.Pop()

		length := len(node.Exprs)
		symbol := v.symbolTable.SearchForDeclaredIdent(node.Ident.Ident)
		lastIsCharOrBool := ast.SizeOf(symbol.T.(ast.ArrayTypeNode).T) == 1 && symbol.T.(ast.ArrayTypeNode).Dim == length

		for i := 0; i < length; i++ {
			expr := node.Exprs[i]
			ast.Walk(v, expr)
			exprRegister := v.getReturnRegister()
			if i > 0 {
				//v.addCode("LDR " + identRegister.String() + ", [" + identRegister.String() + "]")
			}
			v.addCode(
				"MOV r0, "+exprRegister.String(),
				"MOV r1, "+identRegister.String(),
				"BL "+CHECK_ARRAY_INDEX.String(),
				"ADD "+identRegister.String()+", "+identRegister.String()+", #4")

			if i == length-1 && lastIsCharOrBool {
				v.addCode(fmt.Sprintf("ADD %s, %s, %s", identRegister, identRegister, exprRegister))
			} else {
				v.addCode(fmt.Sprintf("ADD %s, %s, %s, LSL #2", identRegister, identRegister, exprRegister))
			}

			// If it is an assignment leave the Pointer to the element in the register
			// otherwise convert to value
			if !node.Pointer {
				if i == length-1 && lastIsCharOrBool {
					v.addCode(fmt.Sprintf("LDRSB %s, [%s]", identRegister, identRegister))
				} else {
					v.addCode(fmt.Sprintf("LDR %s, [%s]", identRegister, identRegister))
				}
			}
		}

		v.usesFunction(CHECK_ARRAY_INDEX)

		v.returnRegisters.Push(identRegister)
	case ast.ArrayLiteralNode:
		register := v.getFreeRegister()
		length := len(node.Exprs)
		size := 0
		if length > 0 {
			size = ast.SizeOf(ast.Type(node.Exprs[0], v.symbolTable))
		}
		v.addCode(
			"LDR r0, ="+strconv.Itoa(length*size+4),
			"BL malloc",
			"MOV "+register.String()+", r0")
		for i := 0; i < length; i++ {
			ast.Walk(v, node.Exprs[i])
			exprRegister := v.getReturnRegister()
			v.addCode(store(size, exprRegister.String(), "["+register.String()+", #"+strconv.Itoa(4+i*size)+"]"))
		}
		lengthRegister := v.freeRegisters.Peek()
		v.addCode(
			"LDR "+lengthRegister.String()+", ="+strconv.Itoa(length),
			"STR "+lengthRegister.String()+", ["+register.String()+"]")
	case ast.NewPairNode:
		register := v.getFreeRegister()

		// Make space for 2 new pointers on heap
		v.addCode("LDR r0, =8",
			"BL malloc",
			"MOV "+register.String()+", r0")

		// Store first element
		ast.Walk(v, node.Fst)
		fst := v.getReturnRegister()
		fstSize := ast.SizeOf(ast.Type(node.Fst, v.symbolTable))
		v.addCode("LDR r0, ="+strconv.Itoa(fstSize),
			"BL malloc",
			"STR r0, ["+register.String()+"]")
		v.addCode(store(fstSize, fst.String(), "[r0]"))

		// Store second element
		ast.Walk(v, node.Snd)
		v.getReturnRegister()
		sndSize := ast.SizeOf(ast.Type(node.Snd, v.symbolTable))
		v.addCode("LDR r0, ="+strconv.Itoa(sndSize),
			"BL malloc",
			"STR r0, ["+register.String()+", #4]")
		v.addCode(store(sndSize, fst.String(), "[r0]"))
	case ast.FunctionCallNode:
		registers := location.ReturnRegisters()
		size := 0
		for i := len(node.Exprs) - 1; i >= 0; i-- {
			ast.Walk(v, node.Exprs[i])
			register := v.getReturnRegister()
			if i < len(registers) {
				v.addCode("MOV " + registers[i].String() + ", " + register.String())
			} else {
				f, _ := v.symbolTable.SearchForFunction(node.Ident.Ident)
				v.addCode("SUB sp, sp, #" + strconv.Itoa(ast.SizeOf(f.Params[i].T)))
				v.addCode(store(ast.SizeOf(f.Params[i].T), register.String(), ", [sp]"))
				//		v.addCode("PUSH {" + register.String() + "}")
				size += ast.SizeOf(f.Params[i].T)
				v.currentStackPos += ast.SizeOf(f.Params[i].T)
			}
		}
		v.addCode("BL f_" + node.Ident.Ident)
		if size > 0 {
			v.addCode("ADD sp, sp, #" + strconv.Itoa(size))
			v.currentStackPos -= size
		}

		register := v.getFreeRegister()
		v.addCode("MOV " + register.String() + ", " + location.R0.String())
	case ast.BaseTypeNode:

	case ast.ArrayTypeNode:

	case ast.PairTypeNode:

	case ast.UnaryOperator:

	case ast.BinaryOperator:

	case ast.IntegerLiteralNode:
		v.addCode("LDR " + v.getFreeRegister().String() + ", =" + strconv.Itoa(node.Val))
	case ast.BooleanLiteralNode:
		register := v.getFreeRegister()
		if node.Val {
			v.addCode("MOV " + register.String() + ", #1") // True
		} else {
			v.addCode("MOV " + register.String() + ", #0") // False
		}
	case ast.CharacterLiteralNode:
		v.addCode("MOV " + v.getFreeRegister().String() + ", #'" + string(node.Val) + "'")
	case ast.StringLiteralNode:
		label := v.addData(node.Val)
		v.addCode("LDR " + v.getFreeRegister().String() + ", =" + label)
	case ast.PairLiteralNode:
		v.addCode("LDR " + v.getFreeRegister().String() + ", =0")
	case ast.UnaryOperatorNode:
		switch node.Op {
		case ast.NOT:

		case ast.NEG:

		case ast.LEN:

		case ast.ORD:

		case ast.CHR:

		}
	case ast.BinaryOperatorNode:
		operand2 := location.UNDEFINED
		operand1 := location.UNDEFINED

		if v.freeRegisters.Length() == 2 {
			ast.Walk(v, node.Expr2)
			operand2 = v.returnRegisters.Pop()
			v.addCode("PUSH {" + operand2.String() + "}")
			v.currentStackPos += ast.SizeOf(ast.Type(node.Expr1, v.symbolTable))
			v.freeRegisters.Push(operand2)

			ast.Walk(v, node.Expr1)
			operand1 = v.returnRegisters.Pop()
			operand2 = v.freeRegisters.Pop()
			v.addCode("POP {" + operand2.String() + "}")
			v.currentStackPos -= ast.SizeOf(ast.Type(node.Expr1, v.symbolTable))
		} else {
			ast.Walk(v, node.Expr2)
			operand2 = v.returnRegisters.Pop()
			ast.Walk(v, node.Expr1)
			operand1 = v.returnRegisters.Pop()
		}
		switch node.Op {
		case ast.MUL:
			v.addCode("SMULL "+operand1.String()+", "+operand2.String()+", "+operand1.String()+", "+operand2.String(),
				"CMP "+operand2.String()+", "+operand1.String()+", ASR #31",
				"BLNE "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case ast.DIV:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL "+CHECK_DIVIDE.String(),
				"BL __aeabi_idiv",
				"MOV "+operand1.String()+", r0")
			v.usesFunction(CHECK_DIVIDE)
		case ast.MOD:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL "+CHECK_DIVIDE.String(),
				"BL __aeabi_idivmod",
				"MOV "+operand1.String()+", r1")
			v.usesFunction(CHECK_DIVIDE)
		case ast.ADD:
			v.addCode("ADDS "+operand1.String()+", "+operand1.String()+", "+operand2.String(),
				"BLVS "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case ast.SUB:
			v.addCode("SUBS "+operand1.String()+", "+operand1.String()+", "+operand2.String(),
				"BLVS "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case ast.GT:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVGT "+operand1.String()+", #1",
				"MOVLE "+operand1.String()+", #0")
		case ast.GEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVGE "+operand1.String()+", #1",
				"MOVLT "+operand1.String()+", #0")
		case ast.LT:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVLT "+operand1.String()+", #1",
				"MOVGE "+operand1.String()+", #0")
		case ast.LEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVLE "+operand1.String()+", #1",
				"MOVGT "+operand1.String()+", #0")
		case ast.EQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVEQ "+operand1.String()+", #1",
				"MOVNE "+operand1.String()+", #0")
		case ast.NEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVNE "+operand1.String()+", #1",
				"MOVEQ "+operand1.String()+", #0")
		case ast.AND:
			v.addCode("AND " + operand1.String() + ", " + operand1.String() + ", " + operand2.String())
		case ast.OR:
			v.addCode("ORR " + operand1.String() + ", " + operand1.String() + ", " + operand2.String())
		}
		v.freeRegisters.Push(operand2)
		v.returnRegisters.Push(operand1)
	case []ast.StatementNode:
		v.symbolTable.MoveNextScope()
		size := 0
		for _, dec := range v.symbolTable.CurrentScope.Scope {
			size += ast.SizeOf(dec.T)
			dec.AddLocation(location.NewStackOffsetLocation(v.currentStackPos + size))
		}
		if size != 0 {
			i := size
			for ; i > 1024; i -= 1024 {
				v.addCode("SUB sp, sp, #1024")
			}
			v.addCode("SUB sp, sp, #" + strconv.Itoa(i))
		}
		v.symbolTable.CurrentScope.ScopeSize = size
		v.currentStackPos += size
	}
}

// Leave will be called to leave the current node.
func (v *CodeGenerator) Leave(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case []ast.StatementNode:
		if v.symbolTable.CurrentScope.ScopeSize != 0 {
			i := v.symbolTable.CurrentScope.ScopeSize
			v.currentStackPos -= i
			for ; i > 1024; i -= 1024 {
				v.addCode("ADD sp, sp, #1024")
			}
			v.addCode("ADD sp, sp, #" + strconv.Itoa(i))
		}
		v.symbolTable.MoveUpScope()
	case ast.FunctionNode:
		if v.symbolTable.CurrentScope.ScopeSize > 0 {
			v.addCode("ADD sp, sp, #" + strconv.Itoa(v.symbolTable.CurrentScope.ScopeSize))
		}
		v.symbolTable.MoveUpScope()
		if node.Ident.Ident == "" {
			v.addCode("LDR r0, =0",
				"POP {pc}")
		}
		v.addCode(".ltorg")
	case ast.ArrayLiteralNode:
	case ast.FreeNode:
		v.addCode("MOV r0, "+v.getReturnRegister().String(),
			"BL "+FREE.String())
		v.usesFunction(FREE)
	case ast.DeclareNode:
		dec, _ := v.symbolTable.SearchForIdentInCurrentScope(node.Ident.Ident)
		dec.IsDeclared = true
		if dec.Location != nil {
			v.addCode(store(ast.SizeOf(dec.T), v.getReturnRegister().String(), v.LocationOf(dec.Location)))
		}
	case ast.PrintNode:
		v.addCode("MOV r0, " + v.getReturnRegister().String())
		v.addPrint(ast.Type(node.Expr, v.symbolTable))
	case ast.PrintlnNode:
		v.addCode("MOV r0, " + v.getReturnRegister().String())
		v.addPrint(ast.Type(node.Expr, v.symbolTable))
		v.addCode("BL " + PRINT_LN.String())
		v.usesFunction(PRINT_LN)
	case ast.ExitNode:
		v.addCode("MOV r0, "+v.getReturnRegister().String(),
			"BL exit")
	case ast.ReturnNode:
		i := 0
		for t := v.symbolTable.CurrentScope; t != v.symbolTable.Head; t = t.ParentScope {
			i += t.ScopeSize
		}
		if i > 0 {
			v.addCode("ADD sp, sp, #" + strconv.Itoa(i))
		}
		v.addCode("MOV r0, "+v.getReturnRegister().String(),
			"POP {pc}")

	case ast.PairFirstElementNode:
		register := v.returnRegisters.Peek()
		v.addCode("MOV r0, "+register.String(),
			"BL "+CHECK_NULL_POINTER.String(),
			"LDR "+register.String()+", ["+register.String()+"]")

		// If we don't want a Pointer then don't retrieve the value
		if !node.Pointer {
			v.addCode(load(ast.SizeOf(ast.Type(node.Expr, v.symbolTable)), register.String(), "["+register.String()+"]"))
		}
		v.usesFunction(CHECK_NULL_POINTER)
	case ast.PairSecondElementNode:
		register := v.returnRegisters.Peek()
		v.addCode("MOV r0, "+register.String(),
			"BL "+CHECK_NULL_POINTER.String(),
			"LDR "+register.String()+", ["+register.String()+", #4]")

		// If we don't want a Pointer then don't retrieve the value
		if !node.Pointer {
			v.addCode(load(ast.SizeOf(ast.Type(node.Expr, v.symbolTable)), register.String(), "["+register.String()+"]"))
		}
		v.usesFunction(CHECK_NULL_POINTER)
	case ast.UnaryOperatorNode:
		register := v.returnRegisters.Peek()
		switch node.Op {
		case ast.NOT:
			v.addCode("EOR " + register.String() + ", " + register.String() + ", #1")
		case ast.NEG:
			v.addCode("RSBS "+register.String()+", "+register.String()+", #0",
				"BLVS "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case ast.LEN:
			v.addCode("LDR " + register.String() + ", [" + register.String() + "]")
		case ast.ORD:

		case ast.CHR:

		}
	}
}

func (v *CodeGenerator) getFreeRegister() location.Register {
	register := v.freeRegisters.Pop()
	v.returnRegisters.Push(register)
	return register
}

func (v *CodeGenerator) getReturnRegister() location.Register {
	register := v.returnRegisters.Pop()
	v.freeRegisters.Push(register)
	return register
}

func store(size int, from string, dest string) string {
	if size == 1 {
		return fmt.Sprintf("STRB %s, %s", from, dest)
	}
	return fmt.Sprintf("STR %s, %s", from, dest)
}

func load(size int, from string, dest string) string {
	if size == 1 {
		return fmt.Sprintf("LDRSB %s, %s", from, dest)
	}
	return fmt.Sprintf("LDR %s, %s", from, dest)
}
