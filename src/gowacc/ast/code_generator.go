package ast

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		buf.WriteString(indent(s, "  "))
	}
	buf.WriteString(".global main\n")
	for fname, f := range asm.global {
		buf.WriteString(fname + ":\n")
		for _, s := range f {
			buf.WriteString(indent(s, "  "))
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

type Register int

const (
	R0 Register = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	R10
	R11
	R12
	SP
	LR
	PC
	APSR
	UNDEFINED
)

func (r Register) String() string {
	switch r {
	case R0:
		return "r0"
	case R1:
		return "r1"
	case R2:
		return "r2"
	case R3:
		return "r3"
	case R4:
		return "r4"
	case R5:
		return "r5"
	case R6:
		return "r6"
	case R7:
		return "r7"
	case R8:
		return "r8"
	case R9:
		return "r9"
	case R10:
		return "r10"
	case R11:
		return "r11"
	case R12:
		return "r12"
	case SP:
		return "sp"
	case LR:
		return "lr"
	case PC:
		return "pc"
	case APSR:
		return "apsr"
	default:
		return "UNDEFINED"
	}
}

// GenerateCode is a function that will generate and return the finished assembly
// code for a given AST.
func GenerateCode(tree ProgramNode, symbolTable *SymbolTable) *Assembly {
	codeGen := NewCodeGenerator(symbolTable)

	Walk(codeGen, tree)

	return codeGen.asm
}

// CodeGenerator is a struct that implements EntryExitVisitor to be called with
// Walk. It stores
type CodeGenerator struct {
	asm             *Assembly
	labelCount      int
	currentFunction string
	symbolTable     *SymbolTable
	freeRegisters   *RegisterStack
	returnRegisters *RegisterStack
	library         *Library
	currentStackPos int
}

// NewCodeGenerator returns an initialised CodeGenerator
func NewCodeGenerator(symbolTable *SymbolTable) *CodeGenerator {
	return &CodeGenerator{
		asm:             NewAssembly(),
		labelCount:      0,
		symbolTable:     symbolTable,
		freeRegisters:   NewRegisterStackWith([]Register{R11, R10, R9, R8, R7, R6, R5, R4}),
		returnRegisters: NewRegisterStack(),
		library:         GetLibrary(),
		currentStackPos: 0,
	}
}

// RegisterStack is a struct that represents a stack of regsters.
// It is used to keep track of which register is used for returning a value.
//
// When a callee returns with value, it pushes the register used to store the
// return value to the stack.
//
// The caller pops a register off the stack to determine the register where the
// return value is stored.
type RegisterStack struct {
	stack []Register
}

func NewRegisterStack() *RegisterStack {
	return &RegisterStack{
		stack: []Register{},
	}
}

func NewRegisterStackWith(registers []Register) *RegisterStack {
	return &RegisterStack{
		stack: registers,
	}
}

func (registerStack *RegisterStack) Pop() Register {
	if len(registerStack.stack) != 0 {
		register := registerStack.stack[len(registerStack.stack)-1]
		registerStack.stack = registerStack.stack[:len(registerStack.stack)-1]
		return register
	}
	fmt.Println("Internal compiler error")
	return UNDEFINED
}

func (registerStack *RegisterStack) Peek() Register {
	if len(registerStack.stack) != 0 {
		register := registerStack.stack[len(registerStack.stack)-1]
		return register
	}
	fmt.Println("Internal compiler error")
	return UNDEFINED
}

func (registerStack *RegisterStack) Push(register Register) {
	registerStack.stack = append(registerStack.stack, register)
}

func (registerStack *RegisterStack) String() string {
	var buf bytes.Buffer
	buf.WriteString("[ ")
	for _, r := range registerStack.stack {
		buf.WriteString(r.String() + " ")
	}
	buf.WriteString("]")
	return buf.String()
}

func (v *CodeGenerator) addPrint(t TypeNode) {
	switch node := t.(type) {
	case BaseTypeNode:
		switch node.t {
		case BOOL:
			v.addCode("BL " + PRINT_BOOL.String())
			v.usesFunction(PRINT_BOOL)
		case INT:
			v.addCode("BL " + PRINT_INT.String())
			v.usesFunction(PRINT_INT)
		case CHAR:
			v.addCode("BL putchar")
		case PAIR:
			v.addCode("BL " + PRINT_REFERENCE.String())
			v.usesFunction(PRINT_REFERENCE)
		}
	case ArrayTypeNode:
		if arr, ok := node.t.(BaseTypeNode); ok {
			if arr.t == CHAR && node.dim == 1 {
				v.addCode("BL " + PRINT_STRING.String())
				v.usesFunction(PRINT_STRING)
				return
			}
		}
		v.addCode("BL " + PRINT_REFERENCE.String())
		v.usesFunction(PRINT_REFERENCE)
	case PairTypeNode:
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
func (v *CodeGenerator) NoRecurse(programNode ProgramNode) bool {
	switch programNode.(type) {
	case IfNode,
		AssignNode,
		ArrayLiteralNode,
		ArrayElementNode,
		LoopNode,
		NewPairNode,
		ReadNode:
		return true
	case FunctionCallNode:
		return true
	default:
		return false
	}
}

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *CodeGenerator) Visit(programNode ProgramNode) {
	switch node := programNode.(type) {
	case Program:

	case FunctionNode:
		v.currentStackPos = 0
		v.freeRegisters.stack = []Register{R11, R10, R9, R8, R7, R6, R5, R4}
		v.symbolTable.MoveNextScope()
		if node.Ident.ident == "" {
			v.addFunction("main")
		} else {
			v.addFunction("f_" + node.Ident.ident)
		}
		v.addCode("PUSH {lr}")
	case []ParameterNode:
		registers := []Register{R0, R1, R2, R3}
		i := 0
		j := 0
		for n, e := range node {
			dec, _ := v.symbolTable.SearchForIdent(e.Ident.ident)
			dec.isDeclared = true
			if n < len(registers) {
				i += SizeOf(e.T)
				dec.AddLocation(NewStackOffsetLocation(i, v))
			} else {
				dec.AddLocation(NewStackOffsetLocation(j-4, v))
				j -= SizeOf(e.T)
			}
		}

		if i > 0 {
			v.addCode("SUB sp, sp, #" + strconv.Itoa(i))
		}
		v.symbolTable.currentScope.scopeSize = i
		v.currentStackPos += i
		for n, e := range node {
			dec, _ := v.symbolTable.SearchForIdent(e.Ident.ident)
			if n < len(registers) {
				if SizeOf(e.T) == 1 {
					v.addCode("STRB " + registers[n].String() + ", " + dec.location.String())
				} else {
					v.addCode("STR " + registers[n].String() + ", " + dec.location.String())
				}
			}
		}
	case ParameterNode:

	case SkipNode:

	case DeclareNode:

	case AssignNode:
		// Rhs
		Walk(v, node.rhs)
		rhsRegister := v.returnRegisters.Pop()
		// Lhs
		switch lhsNode := node.lhs.(type) {
		case ArrayElementNode:
			Walk(v, lhsNode)
			lhsRegister := v.returnRegisters.Pop()
			dec := v.symbolTable.SearchForDeclaredIdent(lhsNode.ident.ident)
			arr := dec.t.(ArrayTypeNode)
			if SizeOf(arr.t) == 1 && len(lhsNode.exprs) == arr.dim {
				v.addCode(fmt.Sprintf("STRB %s, [%s]", rhsRegister, lhsRegister))
			} else {
				v.addCode(fmt.Sprintf("STR %s, [%s]", rhsRegister, lhsRegister))
			}
		case PairFirstElementNode:
			Walk(v, lhsNode)
			lhsRegister := v.returnRegisters.Pop()
			if SizeOf(Type(lhsNode.expr, v.symbolTable)) == 1 {
				v.addCode(fmt.Sprintf("STRB %s, [%s]", rhsRegister, lhsRegister))
			} else {
				v.addCode(fmt.Sprintf("STR %s, [%s]", rhsRegister, lhsRegister))
			}
		case PairSecondElementNode:
			Walk(v, lhsNode)
			lhsRegister := v.returnRegisters.Pop()
			if SizeOf(Type(lhsNode.expr, v.symbolTable)) == 1 {
				v.addCode(fmt.Sprintf("STRB %s, [%s]", rhsRegister, lhsRegister))
			} else {
				v.addCode(fmt.Sprintf("STR %s, [%s]", rhsRegister, lhsRegister))
			}
		case IdentifierNode:
			ident := v.symbolTable.SearchForDeclaredIdent(lhsNode.ident)
			if ident.location != nil {
				if SizeOf(ident.t) == 1 {
					v.addCode("STRB " + rhsRegister.String() + ", " + ident.location.String())
				} else {
					v.addCode("STR " + rhsRegister.String() + ", " + ident.location.String())
				}
			}
		}
		v.freeRegisters.Push(rhsRegister)
	case ReadNode:

		if ident, ok := node.lhs.(IdentifierNode); ok {
			register := v.freeRegisters.Pop()
			dec := v.symbolTable.SearchForDeclaredIdent(ident.ident)
			v.addCode("ADD " + register.String() + ", " + dec.location.PointerTo())
			v.returnRegisters.Push(register)
		} else {
			Walk(v, node.lhs)
		}
		register := v.returnRegisters.Pop()
		v.addCode("MOV r0, " + register.String())
		if SizeOf(Type(node.lhs, v.symbolTable)) == 1 {
			v.addCode("BL " + READ_CHAR.String())
			v.usesFunction(READ_CHAR)
		} else {
			v.addCode("BL " + READ_INT.String())
			v.usesFunction(READ_INT)
		}
		v.freeRegisters.Push(register)
	case FreeNode:

	case ReturnNode:

	case ExitNode:

	case PrintNode:

	case PrintlnNode:

	case IfNode:
		// Labels
		elseLabel := v.labelCount + 1
		endifLabel := v.labelCount + 1
		v.labelCount += 2
		// Cond
		Walk(v, node.expr)
		r := v.returnRegisters.Pop()
		v.addCode(fmt.Sprintf("CMP %s, #0", r))
		v.freeRegisters.Push(r)
		v.addCode(fmt.Sprintf("BEQ ELSE%d", elseLabel))
		// If
		Walk(v, node.ifStats)
		v.addCode(fmt.Sprintf("B ENDIF%d", endifLabel))
		// Else
		v.addCode(fmt.Sprintf("ELSE%d:", elseLabel))
		Walk(v, node.elseStats)
		// Fi
		v.addCode(fmt.Sprintf("ENDIF%d:", endifLabel))
	case LoopNode:
		// Labels
		doLabel := v.labelCount + 1
		whileLabel := v.labelCount + 1
		v.labelCount += 2
		v.addCode(fmt.Sprintf("B WHILE%d", whileLabel))
		// Do
		v.addCode(fmt.Sprintf("DO%d:", doLabel))
		v.labelCount++
		Walk(v, node.stats)
		// While
		v.addCode(fmt.Sprintf("WHILE%d:", whileLabel))
		v.labelCount++
		Walk(v, node.expr)
		r := v.returnRegisters.Pop()
		v.addCode(fmt.Sprintf("CMP %s, #1", r))
		v.freeRegisters.Push(r)
		v.addCode(fmt.Sprintf("BEQ DO%d", doLabel))
	case ScopeNode:

	case IdentifierNode:
		register := v.freeRegisters.Pop()
		dec := v.symbolTable.SearchForDeclaredIdent(node.ident)
		if SizeOf(dec.t) == 1 {
			v.addCode("LDRSB " + register.String() + ", " + dec.location.String())
		} else {
			v.addCode("LDR " + register.String() + ", " + dec.location.String())
		}
		v.returnRegisters.Push(register)
	case PairFirstElementNode:

	case PairSecondElementNode:

	case ArrayElementNode:
		Walk(v, node.ident)
		identRegister := v.returnRegisters.Pop()

		length := len(node.exprs)
		symbol := v.symbolTable.SearchForDeclaredIdent(node.ident.ident)
		lastIsCharOrBool := SizeOf(symbol.t.(ArrayTypeNode).t) == 1 && symbol.t.(ArrayTypeNode).dim == length

		for i := 0; i < length; i++ {
			expr := node.exprs[i]
			Walk(v, expr)
			exprRegister := v.returnRegisters.Pop()
			v.freeRegisters.Push(exprRegister)
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

			// If it is an assignment leave the pointer to the element in the register
			// otherwise convert to value
			if !node.pointer {
				if i == length-1 && lastIsCharOrBool {
					v.addCode(fmt.Sprintf("LDRSB %s, [%s]", identRegister, identRegister))
				} else {
					v.addCode(fmt.Sprintf("LDR %s, [%s]", identRegister, identRegister))
				}
			}
		}

		v.usesFunction(CHECK_ARRAY_INDEX)

		v.returnRegisters.Push(identRegister)
	case ArrayLiteralNode:
		register := v.freeRegisters.Pop()
		length := len(node.exprs)
		size := 0
		if length > 0 {
			size = SizeOf(Type(node.exprs[0], v.symbolTable))
		}
		v.addCode(
			"LDR r0, ="+strconv.Itoa(length*size+4),
			"BL malloc",
			"MOV "+register.String()+", r0")
		for i := 0; i < length; i++ {
			Walk(v, node.exprs[i])
			exprRegister := v.returnRegisters.Pop()
			v.freeRegisters.Push(exprRegister)
			if size == 1 {
				v.addCode("STRB " + exprRegister.String() + ", [" + register.String() + ", #" + strconv.Itoa(4+i*size) + "]")
			} else {
				v.addCode("STR " + exprRegister.String() + ", [" + register.String() + ", #" + strconv.Itoa(4+i*size) + "]")
			}
		}
		lengthRegister := v.freeRegisters.Pop()
		v.addCode(
			"LDR "+lengthRegister.String()+", ="+strconv.Itoa(length),
			"STR "+lengthRegister.String()+", ["+register.String()+"]")
		v.freeRegisters.Push(lengthRegister)
		v.returnRegisters.Push(register)
	case NewPairNode:
		register := v.freeRegisters.Pop()

		// Make space for 2 new pointers on heap
		v.addCode("LDR r0, =8",
			"BL malloc",
			"MOV "+register.String()+", r0")

		// Store first element
		Walk(v, node.fst)
		fst := v.returnRegisters.Pop()
		fstSize := SizeOf(Type(node.fst, v.symbolTable))
		v.addCode("LDR r0, ="+strconv.Itoa(fstSize),
			"BL malloc",
			"STR r0, ["+register.String()+"]")
		if fstSize == 1 {
			v.addCode("STRB " + fst.String() + ", [r0]")
		} else {
			v.addCode("STR " + fst.String() + ", [r0]")
		}
		v.freeRegisters.Push(fst)

		// Store second element
		Walk(v, node.snd)
		snd := v.returnRegisters.Pop()
		sndSize := SizeOf(Type(node.snd, v.symbolTable))
		v.addCode("LDR r0, ="+strconv.Itoa(sndSize),
			"BL malloc",
			"STR r0, ["+register.String()+", #4]")
		if sndSize == 1 {
			v.addCode("STRB " + snd.String() + ", [r0]")
		} else {
			v.addCode("STR " + snd.String() + ", [r0]")
		}
		v.freeRegisters.Push(snd)
		v.returnRegisters.Push(register)
	case FunctionCallNode:
		registers := []Register{R0, R1, R2, R3}
		size := 0
		for i := len(node.exprs) - 1; i >= 0; i-- {
			Walk(v, node.exprs[i])
			register := v.returnRegisters.Pop()
			v.freeRegisters.Push(register)
			if i < len(registers) {
				v.addCode("MOV " + registers[i].String() + ", " + register.String())
			} else {
				f, _ := v.symbolTable.SearchForFunction(node.ident.ident)
				v.addCode("SUB sp, sp, #" + strconv.Itoa(SizeOf(f.Params[i].T)))
				if SizeOf(f.Params[i].T) == 1 {
					v.addCode("STRB " + register.String() + ", [sp]")
				} else {
					v.addCode("STR " + register.String() + ", [sp]")
				}
				//		v.addCode("PUSH {" + register.String() + "}")
				size += SizeOf(f.Params[i].T)
				v.currentStackPos += SizeOf(f.Params[i].T)
			}
		}
		v.addCode("BL f_" + node.ident.ident)
		if size > 0 {
			v.addCode("ADD sp, sp, #" + strconv.Itoa(size))
			v.currentStackPos -= size
		}

		register := v.freeRegisters.Pop()
		v.returnRegisters.Push(register)
		v.addCode("MOV " + register.String() + ", " + R0.String())
	case BaseTypeNode:

	case ArrayTypeNode:

	case PairTypeNode:

	case UnaryOperator:

	case BinaryOperator:

	case IntegerLiteralNode:
		register := v.freeRegisters.Pop()
		v.addCode("LDR " + register.String() + ", =" + strconv.Itoa(node.val))
		v.returnRegisters.Push(register)
	case BooleanLiteralNode:
		register := v.freeRegisters.Pop()
		if node.val {
			v.addCode("MOV " + register.String() + ", #1") // True
		} else {
			v.addCode("MOV " + register.String() + ", #0") // False
		}
		v.returnRegisters.Push(register)
	case CharacterLiteralNode:
		register := v.freeRegisters.Pop()
		v.addCode("MOV " + register.String() + ", #'" + string(node.val) + "'")
		v.returnRegisters.Push(register)
	case StringLiteralNode:
		register := v.freeRegisters.Pop()
		label := v.addData(node.val)
		v.addCode("LDR " + register.String() + ", =" + label)
		v.returnRegisters.Push(register)
	case PairLiteralNode:
		register := v.freeRegisters.Pop()
		v.returnRegisters.Push(register)
		v.addCode("LDR " + register.String() + ", =0")
	case UnaryOperatorNode:
		switch node.op {
		case NOT:

		case NEG:

		case LEN:

		case ORD:

		case CHR:

		}
	case BinaryOperatorNode:
		switch node.op {
		case MUL, DIV, MOD, ADD, SUB:

		case GT, GEQ, LT, LEQ:

		case EQ, NEQ:

		case AND, OR:

		}
	case []StatementNode:
		v.symbolTable.MoveNextScope()
		size := 0
		for _, dec := range v.symbolTable.currentScope.scope {
			size += SizeOf(dec.t)
			dec.AddLocation(NewStackOffsetLocation(v.currentStackPos+size, v))
		}
		if size != 0 {
			i := size
			for ; i > 1024; i -= 1024 {
				v.addCode("SUB sp, sp, #1024")
			}
			v.addCode("SUB sp, sp, #" + strconv.Itoa(i))
		}
		v.symbolTable.currentScope.scopeSize = size
		v.currentStackPos += size
	}
}

// Leave will be called to leave the current node.
func (v *CodeGenerator) Leave(programNode ProgramNode) {
	switch node := programNode.(type) {
	case []StatementNode:
		if v.symbolTable.currentScope.scopeSize != 0 {
			i := v.symbolTable.currentScope.scopeSize
			v.currentStackPos -= i
			for ; i > 1024; i -= 1024 {
				v.addCode("ADD sp, sp, #1024")
			}
			v.addCode("ADD sp, sp, #" + strconv.Itoa(i))
		}
		v.symbolTable.MoveUpScope()
	case FunctionNode:
		if v.symbolTable.currentScope.scopeSize > 0 {
			v.addCode("ADD sp, sp, #" + strconv.Itoa(v.symbolTable.currentScope.scopeSize))
		}
		v.symbolTable.MoveUpScope()
		if node.Ident.ident == "" {
			v.addCode("LDR r0, =0",
				"POP {pc}")
		}
		v.addCode(".ltorg")
	case ArrayLiteralNode:
	case FreeNode:
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		v.addCode("MOV r0, "+register.String(),
			"BL "+FREE.String())
		v.usesFunction(FREE)
	case DeclareNode:
		dec, _ := v.symbolTable.SearchForIdentInCurrentScope(node.ident.ident)
		dec.isDeclared = true
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		if dec.location != nil {
			if SizeOf(dec.t) == 1 {
				v.addCode("STRB " + register.String() + ", " + dec.location.String())
			} else {
				v.addCode("STR " + register.String() + ", " + dec.location.String())
			}
		}
	case PrintNode:
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		v.addCode("MOV r0, " + register.String())
		v.addPrint(Type(node.expr, v.symbolTable))
	case PrintlnNode:
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		v.addCode("MOV r0, " + register.String())
		v.addPrint(Type(node.expr, v.symbolTable))
		v.addCode("BL " + PRINT_LN.String())
		v.usesFunction(PRINT_LN)
	case ExitNode:
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		v.addCode("MOV r0, "+register.String(),
			"BL exit")
	case ReturnNode:
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		i := 0
		for t := v.symbolTable.currentScope; t != v.symbolTable.head; t = t.parentScope {
			i += t.scopeSize
		}
		if i > 0 {
			v.addCode("ADD sp, sp, #" + strconv.Itoa(i))
		}
		v.addCode("MOV r0, "+register.String(),
			"POP {pc}")

	case PairFirstElementNode:
		register := v.returnRegisters.Peek()
		v.addCode("MOV r0, "+register.String(),
			"BL "+CHECK_NULL_POINTER.String(),
			"LDR "+register.String()+", ["+register.String()+"]")

		// If we don'T want a pointer then don'T retrieve the value
		if !node.pointer {
			if SizeOf(Type(node.expr, v.symbolTable)) == 1 {
				v.addCode("LDRSB " + register.String() + ", [" + register.String() + "]")
			} else {
				v.addCode("LDR " + register.String() + ", [" + register.String() + "]")
			}
		}
		v.usesFunction(CHECK_NULL_POINTER)
	case PairSecondElementNode:
		register := v.returnRegisters.Peek()
		v.addCode("MOV r0, "+register.String(),
			"BL "+CHECK_NULL_POINTER.String(),
			"LDR "+register.String()+", ["+register.String()+", #4]")

		// If we don'T want a pointer then don'T retrieve the value
		if !node.pointer {
			if SizeOf(Type(node.expr, v.symbolTable)) == 1 {
				v.addCode("LDRSB " + register.String() + ", [" + register.String() + "]")
			} else {
				v.addCode("LDR " + register.String() + ", [" + register.String() + "]")
			}
		}
		v.usesFunction(CHECK_NULL_POINTER)
	case UnaryOperatorNode:
		register := v.returnRegisters.Peek()
		switch node.op {
		case NOT:
			v.addCode("EOR " + register.String() + ", " + register.String() + ", #1")
		case NEG:
			v.addCode("RSBS "+register.String()+", "+register.String()+", #0",
				"BLVS "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case LEN:
			v.addCode("LDR " + register.String() + ", [" + register.String() + "]")
		case ORD:

		case CHR:

		}
	case BinaryOperatorNode:
		operand2 := v.returnRegisters.Pop()
		operand1 := v.returnRegisters.Pop()
		switch node.op {
		case MUL:
			v.addCode("SMULL "+operand1.String()+", "+operand2.String()+", "+operand1.String()+", "+operand2.String(),
				"CMP "+operand2.String()+", "+operand1.String()+", ASR #31",
				"BLNE "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case DIV:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL "+CHECK_DIVIDE.String(),
				"BL __aeabi_idiv",
				"MOV "+operand1.String()+", r0")
			v.usesFunction(CHECK_DIVIDE)
		case MOD:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL "+CHECK_DIVIDE.String(),
				"BL __aeabi_idivmod",
				"MOV "+operand1.String()+", r1")
			v.usesFunction(CHECK_DIVIDE)
		case ADD:
			v.addCode("ADDS "+operand1.String()+", "+operand1.String()+", "+operand2.String(),
				"BLVS "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case SUB:
			v.addCode("SUBS "+operand1.String()+", "+operand1.String()+", "+operand2.String(),
				"BLVS "+CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case GT:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVGT "+operand1.String()+", #1",
				"MOVLE "+operand1.String()+", #0")
		case GEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVGE "+operand1.String()+", #1",
				"MOVLT "+operand1.String()+", #0")
		case LT:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVLT "+operand1.String()+", #1",
				"MOVGE "+operand1.String()+", #0")
		case LEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVLE "+operand1.String()+", #1",
				"MOVGT "+operand1.String()+", #0")
		case EQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVEQ "+operand1.String()+", #1",
				"MOVNE "+operand1.String()+", #0")
		case NEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVNE "+operand1.String()+", #1",
				"MOVEQ "+operand1.String()+", #0")
		case AND:
			v.addCode("AND " + operand1.String() + ", " + operand1.String() + ", " + operand2.String())
		case OR:
			v.addCode("ORR " + operand1.String() + ", " + operand1.String() + ", " + operand2.String())
		}
		v.freeRegisters.Push(operand2)
		v.returnRegisters.Push(operand1)
	}
}

type Location struct {
	register Register
	address  int

	// Stores information needed to determine stack offset
	currentPos    int
	codeGenerator *CodeGenerator
}

func NewRegisterLocation(register Register) *Location {
	return &Location{
		register: register,
	}
}

func NewAddressLocation(address int) *Location {
	return &Location{
		register: UNDEFINED,
		address:  address,
	}
}

func NewStackOffsetLocation(currentPos int, v *CodeGenerator) *Location {
	return &Location{
		register:      UNDEFINED,
		currentPos:    currentPos,
		codeGenerator: v,
	}
}

func (location *Location) String() string {
	// Location is a register
	if location.register != UNDEFINED {
		return location.register.String()
	}

	// Location is an address on the heap
	if location.address != 0 {
		return "#" + strconv.Itoa(location.address)
	}

	// Location is a stack offset
	return "[sp, #" + strconv.Itoa(location.codeGenerator.currentStackPos-location.currentPos) + "]"
}

func (location *Location) PointerTo() string {
	return "sp, #" + strconv.Itoa(+location.codeGenerator.currentStackPos-location.currentPos)
}
