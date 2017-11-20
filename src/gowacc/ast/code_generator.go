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
	global []string
}

func NewAssembly() *Assembly {
	return &Assembly{
		data: make([]string, 0),
		text: make([]string, 0),
		global: make([]string, 0),
	}
}

// String will return the string format of the Assembly code with line numbers.
func (asm *Assembly) String() string {
	var buf bytes.Buffer
	buf.WriteString(indent(".data", "  "))
	for _, s := range asm.data {
		buf.WriteString(indent(s, "  "))
	}
	buf.WriteString(indent(".text", "  "))
	for _, s := range asm.text {
		buf.WriteString(indent(s, "  "))
	}
	buf.WriteString(indent(".global main", "  "))
	for _, s := range asm.global {
		buf.WriteString(indent(s, "  "))
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
)

// GenerateCode is a function that will generate and return the finished assembly
// code for a given AST.
func GenerateCode(tree ProgramNode, symbolTable *SymbolTable) *Assembly {
	codeGen := NewCodeGenerator(symbolTable)

	Walk(codeGen, tree)

	codeGen.addCode("code", "code2")
	codeGen.addCode("assembly")

	return codeGen.asm
}

// CodeGenerator is a struct that implements EntryExitVisitor to be called with
// Walk. It stores
type CodeGenerator struct {
	asm *Assembly
	symbolTable *SymbolTable
	freeRegisters []Register
	returnRegisters []Register
}

// NewCodeGenerator returns an initialised CodeGenerator
func NewCodeGenerator(symbolTable *SymbolTable) *CodeGenerator {
	return &CodeGenerator{
		asm: NewAssembly(),
		symbolTable: symbolTable,
		freeRegisters: []Register{R3,R4,R5,R6,R7,R8,R9,R10,R11,R12},
		returnRegisters: make([]Register, 0),
	}
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
		v.asm.global = append(v.asm.global, line + "\n")
	}
}

// Visit will apply the correct rule for the programNode given, to be used with
// Walk.
func (v *CodeGenerator) Visit(programNode ProgramNode) {
	switch node := programNode.(type) {
	case Program:

	case FunctionNode:

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

	}
}

// Leave will be called to leave the current node.
func (v *CodeGenerator) Leave(programNode ProgramNode) {
	switch programNode.(type) {
	case []StatementNode:

	case FunctionNode:

	case ArrayLiteralNode:
	}
}
