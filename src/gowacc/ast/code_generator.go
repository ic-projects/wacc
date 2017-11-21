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
	text string
}

func NewAsciiWord(length int, text string) AsciiWord {
	return AsciiWord{
		length: length,
		text: text,
	}
}


type Assembly struct {
	data   			map[string](AsciiWord)
	dataCounter int
	text   			[]string
	global 			map[string]([]string)
}

func NewAssembly() *Assembly {
	return &Assembly{
		data:   make(map[string]AsciiWord),
		dataCounter: 0,
		text:   make([]string, 0),
		global: make(map[string]([]string)),
	}
}

// String will return the string format of the Assembly code with line numbers.
func (asm *Assembly) String() string {
	var buf bytes.Buffer
	buf.WriteString(".data\n")
	for dname, d := range asm.data {
		buf.WriteString(dname + ":\n")
		buf.WriteString(fmt.Sprintf("   .word %d", d.length))
		buf.WriteString(fmt.Sprintf("   .ascii \"%s\"", d.text))
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
	currentFunction string
	symbolTable     *SymbolTable
	freeRegisters   *RegisterStack
	returnRegisters *RegisterStack
	library					*Library
}

// NewCodeGenerator returns an initialised CodeGenerator
func NewCodeGenerator(symbolTable *SymbolTable) *CodeGenerator {
	return &CodeGenerator{
		asm:             NewAssembly(),
		symbolTable:     symbolTable,
		freeRegisters:   NewRegisterStackWith([]Register{R4, R5, R6, R7, R8, R9, R10, R11, R12}),
		returnRegisters: NewRegisterStack(),
		library:				 GetLibrary(),
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
	register := registerStack.stack[len(registerStack.stack)-1]
	registerStack.stack = registerStack.stack[:len(registerStack.stack)-1]
	return register
}

func (registerStack *RegisterStack) Push(register Register) {
	registerStack.stack = append(registerStack.stack, register)
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
	length := 1 // Get length of text
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

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *CodeGenerator) Visit(programNode ProgramNode) {
	switch node := programNode.(type) {
	case Program:

	case FunctionNode:
		v.symbolTable.MoveNextScope()
		if node.ident.ident == "" {
			v.addFunction("main")
		} else {
			v.addFunction("f_" + node.ident.ident)
		}
		v.addCode("PUSH {lr}")
	case ParameterNode:

	case SkipNode:

	case DeclareNode:

	case AssignNode:

	case ReadNode:

	case FreeNode:

	case ReturnNode:

	case ExitNode:

	case PrintNode:

	case PrintlnNode:

	case IfNode:

	case LoopNode:

	case ScopeNode:

	case IdentifierNode:

	case PairFirstElementNode:

	case PairSecondElementNode:

	case ArrayElementNode:

	case ArrayLiteralNode:

	case NewPairNode:

	case FunctionCallNode:

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

	case StringLiteralNode:

	case PairLiteralNode:

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
	}
}

// Leave will be called to leave the current node.
func (v *CodeGenerator) Leave(programNode ProgramNode) {
	switch node := programNode.(type) {
	case []StatementNode:
		v.symbolTable.MoveUpScope()

	case FunctionNode:
		v.symbolTable.MoveUpScope()
		if node.ident.ident == "" {
			v.addCode("LDR r0, =0",
				"POP {pc}")
		}
		v.addCode(".ltorg")
	case ArrayLiteralNode:

	case ExitNode:
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		v.addCode("MOV r0, "+register.String(),
			"BL exit")
	case ReturnNode:
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		v.addCode("MOV r0, "+register.String(),
			"POP {pc}")
	case UnaryOperatorNode:
		operand := v.returnRegisters.Pop()
		returnRegister := v.freeRegisters.Pop()
		switch node.op {
		case NOT:
			v.addCode("EOR " + returnRegister.String() + ", " + operand.String() + ", #1")
		case NEG:

		case LEN:

		case ORD:

		case CHR:

		}
		v.freeRegisters.Push(operand)
		v.returnRegisters.Push(returnRegister)
	case BinaryOperatorNode:
		operand2 := v.returnRegisters.Pop()
		operand1 := v.returnRegisters.Pop()
		returnRegister := v.freeRegisters.Pop()
		switch node.op {
		case MUL:
			v.addCode("MUL " + returnRegister.String() + ", " + operand1.String() + ", " + operand2.String())
		case DIV:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL __aeabi_idiv",
				"MOV "+returnRegister.String()+", r0")
		case MOD:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL __aeabi_idivmod",
				"MOV "+returnRegister.String()+", r1")
		case ADD:
			v.addCode("ADD " + returnRegister.String() + ", " + operand1.String() + ", " + operand2.String())
		case SUB:
			v.addCode("SUB " + returnRegister.String() + ", " + operand1.String() + ", " + operand2.String())
		case GT:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVGT "+returnRegister.String()+", #1",
				"MOVLE "+returnRegister.String()+", #0")
		case GEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVGE "+returnRegister.String()+", #1",
				"MOVLT "+returnRegister.String()+", #0")
		case LT:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVLT "+returnRegister.String()+", #1",
				"MOVGE "+returnRegister.String()+", #0")
		case LEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVLE "+returnRegister.String()+", #1",
				"MOVGT "+returnRegister.String()+", #0")
		case EQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVEQ "+returnRegister.String()+", #1",
				"MOVNE "+returnRegister.String()+", #0")
		case NEQ:
			v.addCode("CMP "+operand1.String()+", "+operand2.String(),
				"MOVNE "+returnRegister.String()+", #1",
				"MOVEQ "+returnRegister.String()+", #0")
		case AND:
			v.addCode("AND " + returnRegister.String() + ", " + operand1.String() + ", " + operand2.String())
		case OR:
			v.addCode("ORR " + returnRegister.String() + ", " + operand1.String() + ", " + operand2.String())
		}
		v.freeRegisters.Push(operand1)
		v.freeRegisters.Push(operand2)
		v.returnRegisters.Push(returnRegister)
	}
}

type Location struct {
	register    Register
	address     int
	stackOffset int
}

func NewRegisterLocation(register Register) *Location {
	return &Location{
		register:    register,
		address:     -1,
		stackOffset: -1,
	}
}

func NewAddressLocation(address int) *Location {
	return &Location{
		register:    UNDEFINED,
		address:     address,
		stackOffset: -1,
	}
}

func NewStackOffsetLocation(offset int) *Location {
	return &Location{
		register:    UNDEFINED,
		address:     -1,
		stackOffset: offset,
	}
}

func (location *Location) UpdateStackOffsetPush() {
	// Only updating if the location is a StackOffsetLocation
	if location.stackOffset != -1 {
		location.stackOffset++
	}
}

func (location *Location) UpdateStackOffsetPop() {
	// Only updating if the location is a StackOffsetLocation
	if location.stackOffset != -1 {
		location.stackOffset--
	}
}

func (location *Location) String() string {
	// Location is a register
	if location.register == UNDEFINED {
		return location.register.String()
	}

	var buf bytes.Buffer

	// Location is an address on the heap
	if location.address != -1 {
		buf.WriteString("#")
		buf.WriteString(strconv.Itoa(location.address))
		return buf.String()
	}

	// Location is a stack offset
	buf.WriteString("[sp, #")
	buf.WriteString(strconv.Itoa(location.stackOffset))
	buf.WriteString("]")
	return buf.String()
}
