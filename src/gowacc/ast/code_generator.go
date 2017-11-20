package ast

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Assembly struct {
	data []string
	text []string
	global map[string]([]string)
}

func NewAssembly() *Assembly {
	return &Assembly{
		data: make([]string, 0),
		text: make([]string, 0),
		global: make(map[string]([]string)),
	}
}

// String will return the string format of the Assembly code with line numbers.
func (asm *Assembly) String() string {
	var buf bytes.Buffer
	buf.WriteString(".data\n")
	for _, s := range asm.data {
		buf.WriteString(indent(s, "  "))
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
	R0    Register = iota
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
	asm *Assembly
	currentFunction string
	symbolTable *SymbolTable
	freeRegisters []Register
	returnRegisters *RegisterStack
}

// NewCodeGenerator returns an initialised CodeGenerator
func NewCodeGenerator(symbolTable *SymbolTable) *CodeGenerator {
	return &CodeGenerator{
		asm: NewAssembly(),
		symbolTable: symbolTable,
		freeRegisters: []Register{R3,R4,R5,R6,R7,R8,R9,R10,R11,R12},
		returnRegisters: NewRegisterStack(),
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

func (registerStack *RegisterStack) Pop() Register {
	register := registerStack.stack[len(registerStack.stack) - 1];
	registerStack.stack = registerStack.stack[:len(registerStack.stack) - 1]
	return register;
}

func (registerStack *RegisterStack) Push(register Register) {
	registerStack.stack = append(registerStack.stack, register)
}

// addData add lines of assembly to the already data part of the generated
// assembly code
func (v *CodeGenerator) addData(lines ...string) {
	for _, line := range lines {
		v.asm.data = append(v.asm.data, line + "\n")
	}
}

// addText add lines of assembly to the already text part of the generated
// assembly code
func (v *CodeGenerator) addText(lines ...string) {
	for _, line := range lines {
		v.asm.text = append(v.asm.text, line + "\n")
	}
}

// addCode add lines of assembly to the already code part of the generated
// assembly code
func (v *CodeGenerator) addCode(lines ...string) {
	for _, line := range lines {
		v.asm.global[v.currentFunction] = append(v.asm.global[v.currentFunction], line + "\n")
	}
}

// addCode add lines of assembly to the already code part of the generated
// assembly code
func (v *CodeGenerator) addFunction(name string) {
	v.asm.global[name] = make([]string, 0)
	v.currentFunction = name
}

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *CodeGenerator) Visit(programNode ProgramNode) {
	switch node := programNode.(type) {
	case Program:

	case FunctionNode:
    v.symbolTable.MoveNextScope()
		if (node.ident.ident == "") {
			v.addFunction("main")
		} else {
			v.addFunction("f_"+node.ident.ident)
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

	case BooleanLiteralNode:

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
		if (node.ident.ident == "") {
			v.addCode("LDR r0, =0")
		}
		v.addCode("POP {pc}", ".ltorg")
	case ArrayLiteralNode:
	}
}
