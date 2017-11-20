package ast

import (
	"bytes"
	"fmt"
	"strings"
)

type Assembly []string

func (asm Assembly) String() string {
	var tempbuf bytes.Buffer
	for _, s := range asm {
		tempbuf.WriteString(indent(s, "  "))
	}
	var buf bytes.Buffer
	for i, line := range strings.Split(tempbuf.String(), "\n") {
		if line != "" {
			buf.WriteString(fmt.Sprintf("%d\t%s\n", i, line))
		}
	}
	return buf.String()
}

// GenerateCode is a function that will generate and return the finished assembly
// code for a given AST.
func GenerateCode(tree ProgramNode) Assembly {
	codeGen := NewCodeGenerator()

	Walk(codeGen, tree)

	codeGen.add("code")
	codeGen.add("assembly")

	return codeGen.asm
}

// CodeGenerator is a struct that implements EntryExitVisitor to be called with
// Walk. It stores
type CodeGenerator struct {
	asm Assembly
}

// NewCodeGenerator returns an initialised CodeGenerator
func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{
		asm: make([]string, 0),
	}
}

func (v *CodeGenerator) add(line string) {
	v.asm = append(v.asm, line)
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
