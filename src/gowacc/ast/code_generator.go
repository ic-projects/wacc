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
		}
	}
}

// addDataWithLabel adds a ascii word to the data section generating a unique label
func (v *CodeGenerator) addData(text string) string {
	label := "msg_" + strconv.Itoa(v.asm.dataCounter)
	v.asm.dataCounter++
	v.addDataWithLabel(label, text, len(text))
	return label
}

// addDataWithLabel adds a ascii word to the data section using a given label
func (v *CodeGenerator) addDataWithLabel(label string, text string, length int) {
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
// walked. This means we can visit the children in any way we choose.
func (v *CodeGenerator) NoRecurse(programNode ProgramNode) bool {
	switch programNode.(type) {
	case IfNode:
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
		// Cond
		Walk(v, node.expr)
		r := v.returnRegisters.Pop()
		v.addCode(fmt.Sprintf("CMP %s, #0", r))
		v.freeRegisters.Push(r)
		v.addCode(fmt.Sprintf("BEQ L%d", v.labelCount))
		// If
		Walk(v, node.ifStats)
		v.addCode(fmt.Sprintf("B L%d", v.labelCount))
		// Else
		v.addCode(fmt.Sprintf("L%d:", v.labelCount))
		v.labelCount++
		Walk(v, node.elseStats)
		// Fi
		v.addCode(fmt.Sprintf("L%d:", v.labelCount))
		v.labelCount++
	case LoopNode:

	case ScopeNode:

	case IdentifierNode:
		register := v.freeRegisters.Pop()
		dec, _ := v.symbolTable.SearchForIdent(node.ident)
		v.addCode("LDR " + register.String() + ", " + dec.location.String())
		v.returnRegisters.Push(register)
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
	case CharacterLiteralNode:
		register := v.freeRegisters.Pop()
		v.addCode("LDR " + register.String() + ", #" + string(node.val))
		v.returnRegisters.Push(register)
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
		i := 0
		for _, dec := range v.symbolTable.currentScope.scope {
			switch typeNode := dec.t.(type) {
			case BaseTypeNode:
				switch typeNode.t {
				case STRING:
				case PAIR:
				default:
					dec.AddLocation(NewStackOffsetLocation(i, v))
					i += sizeOf(dec.t)
				}
			}
		}
		if i != 0 {
			v.addCode("SUB sp, sp, #" + strconv.Itoa(i))
		}
		v.symbolTable.currentScope.scopeSize = i
	}
}

// Leave will be called to leave the current node.
func (v *CodeGenerator) Leave(programNode ProgramNode) {
	switch node := programNode.(type) {
	case []StatementNode:
		if v.symbolTable.currentScope.scopeSize != 0 {
			v.addCode("ADD sp, sp, #" + strconv.Itoa(v.symbolTable.currentScope.scopeSize))
		}
		v.symbolTable.MoveUpScope()
	case FunctionNode:
		v.symbolTable.MoveUpScope()
		if node.ident.ident == "" {
			v.addCode("LDR r0, =0",
				"POP {pc}")
		}
		v.addCode(".ltorg")
	case ArrayLiteralNode:

	case DeclareNode:
		dec, _ := v.symbolTable.SearchForIdentInCurrentScope(node.ident.ident)
		register := v.returnRegisters.Pop()
		v.freeRegisters.Push(register)
		if dec.location != nil {
			if sizeOf(dec.t) == 1 {
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
		v.addCode("MOV r0, "+register.String(),
			"POP {pc}")
	case UnaryOperatorNode:
		operand := v.returnRegisters.Pop()
		returnRegister := v.freeRegisters.Pop()
		switch node.op {
		case NOT:
			v.addCode("EOR " + returnRegister.String() + ", " + operand.String() + ", #1")
		case NEG:
			register := v.returnRegisters.Peek()
			v.addCode("RSBS " + register.String() + ", " + register.String() + ", #0")
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
			v.addCode("SMULL " + returnRegister.String() + ", " + operand2.String() + ", " + operand1.String() + ", " + operand2.String(),
				"CMP "+ operand2.String() + ", " + returnRegister.String() + ", ASR #31",
				"BLNE " + CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case DIV:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL "+CHECK_DIVIDE.String(),
				"BL __aeabi_idiv",
				"MOV "+returnRegister.String()+", r0")
			v.usesFunction(CHECK_DIVIDE)
		case MOD:
			v.addCode("MOV r0, "+operand1.String(),
				"MOV r1, "+operand2.String(),
				"BL "+CHECK_DIVIDE.String(),
				"BL __aeabi_idivmod",
				"MOV "+returnRegister.String()+", r1")
			v.usesFunction(CHECK_DIVIDE)
		case ADD:
			v.addCode("ADDS " + returnRegister.String() + ", " + operand1.String() + ", " + operand2.String(),
				"BLVS " + CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
		case SUB:
			v.addCode("SUB " + returnRegister.String() + ", " + operand1.String() + ", " + operand2.String(),
				"BLVS " + CHECK_OVERFLOW.String())
			v.usesFunction(CHECK_OVERFLOW)
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
