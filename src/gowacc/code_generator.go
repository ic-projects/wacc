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

const (
	malloc  string = "malloc"
	exit    string = "exit"
	div     string = "__aeabi_idiv"
	divmod  string = "__aeabi_idivmod"
	putchar string = "putchar"
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
	buf.WriteString(".data\n\n")
	for dname, d := range asm.data {
		buf.WriteString(dname + ":\n")
		buf.WriteString(fmt.Sprintf("\t.word %d\n", d.length))
		buf.WriteString(fmt.Sprintf("\t.ascii \"%s\"\n", d.text))
	}
	buf.WriteString("\n.text\n\n")
	for _, s := range asm.text {
		buf.WriteString(utils.Indent(s, "\t"))
	}
	buf.WriteString(".global main\n")
	for fname, f := range asm.global {
		buf.WriteString(fname + ":\n")
		for _, s := range f {
			buf.WriteString(s)
		}
	}
	buf.WriteString("\n")
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

// LocationOf will return the string representation of a Location.
//
// If the Location is an address (that is stored on the heap), it will return
// the stack offset.
func (v *CodeGenerator) LocationOf(loc *utils.Location) Address {
	// Location is a register
	if loc.IsRegister() {
		return RegisterAddress{loc.Register, 0}
	}

	// Location is either a stack offset or an address stored on the stack.
	return RegisterAddress{utils.SP, v.currentStackPos - loc.CurrentPos}
}

// PointerTo returns an int, that when added to SP gives the object's location in
// memory
func (v *CodeGenerator) PointerTo(location *utils.Location) int {
	return v.currentStackPos - location.CurrentPos
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
			v.callLibraryFunction(AL, printBool)
		case ast.INT:
			v.callLibraryFunction(AL, printInt)
		case ast.CHAR:
			v.addCode(NewBranchL(putchar).armAssembly())
		case ast.PAIR:
			v.callLibraryFunction(AL, printReference)
		}
	case ast.ArrayTypeNode:
		if arr, ok := ast.ToValue(node.T).(ast.BaseTypeNode); ok {
			if arr.T == ast.CHAR {
				v.callLibraryFunction(AL, printString)
				return
			}
		}
		v.callLibraryFunction(AL, printReference)
	case
		ast.PairTypeNode,
		ast.StructTypeNode,
		ast.NullTypeNode,
		ast.PointerTypeNode:
		v.callLibraryFunction(AL, printReference)
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
	cond Condition,
	function LibraryFunction,
) {
	v.addCode(NewCondBranchL(cond, function.String()).armAssembly())
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
		*ast.PointerNewNode,
		*ast.PointerDereferenceNode,
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
		v.visitFunctionNode(node)
	case ast.Parameters:
		v.visitParameters(node)
	case *ast.AssignNode:
		v.visitAssignNode(node)
	case *ast.ReadNode:
		v.visitReadNode(node)
	case *ast.IfNode:
		v.visitIfNode(node)
	case *ast.SwitchNode:
		v.visitSwitchNode(node)
	case *ast.LoopNode:
		v.visitLoopNode(node)
	case *ast.ForLoopNode:
		v.visitForLoopNode(node)
	case *ast.IdentifierNode:
		v.visitIdentifierNode(node)
	case *ast.ArrayElementNode:
		v.visitArrayElementNode(node)
	case *ast.StructElementNode:
		v.visitStructElementNode(node)
	case *ast.ArrayLiteralNode:
		v.visitArrayLiteralNode(node)
	case *ast.StructNewNode:
		v.visitStructNewNode(node)
	case *ast.PointerNewNode:
		v.visitPointerNewNode(node)
	case *ast.PointerDereferenceNode:
		v.visitPointerDereferenceNode(node)
	case *ast.NewPairNode:
		v.visitNewPairNode(node)
	case *ast.FunctionCallNode:
		v.visitFunctionCallNode(node)
	case *ast.IntegerLiteralNode:
		v.visitIntegerLiteralNode(node)
	case *ast.BooleanLiteralNode:
		v.visitBooleanLiteralNode(node)
	case *ast.CharacterLiteralNode:
		v.visitCharacterLiteralNode(node)
	case *ast.StringLiteralNode:
		v.visitStringLiteralNode(node)
	case *ast.NullNode:
		v.visitNullNode()
	case *ast.BinaryOperatorNode:
		v.visitBinaryOperatorNode(node)
	case ast.Statements:
		v.visitStatements()
	}
}

// Leave will be called to leave the current node.
func (v *CodeGenerator) Leave(programNode ast.ProgramNode) {
	switch node := programNode.(type) {
	case ast.Statements:
		v.leaveStatements()
	case *ast.FunctionNode:
		v.leaveFunctionNode(node)
	case *ast.FreeNode:
		v.leaveFreeNode()
	case *ast.DeclareNode:
		v.leaveDeclareNode(node)
	case *ast.PrintNode:
		v.leavePrintNode(node)
	case *ast.PrintlnNode:
		v.leavePrintlnNode(node)
	case *ast.ExitNode:
		v.leaveExitNode()
	case *ast.ReturnNode:
		v.leaveReturnNode()
	case *ast.PairFirstElementNode:
		v.leavePairFirstElementNode(node)
	case *ast.PairSecondElementNode:
		v.leavePairSecondElementNode(node)
	case *ast.UnaryOperatorNode:
		v.leaveUnaryOperatorNode(node)
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
		v.addCode(NewAdd(utils.SP, utils.SP, 1024).armAssembly())
	}
	v.addCode(NewAdd(utils.SP, utils.SP, i).armAssembly())
}

// subtractFromStackPointer decrements the stack pointer by the size parameter.
// If size is greater than 1024 then it will decrement in multiple iterations.
func (v *CodeGenerator) subtractFromStackPointer(size int) {
	// Cannot subtract more than 1024 from SP at once, so do it in multiple
	// iterations.
	i := size
	v.currentStackPos += size
	for ; i > 1024; i -= 1024 {
		v.addCode(NewSub(utils.SP, utils.SP, 1024).armAssembly())
	}
	v.addCode(NewSub(utils.SP, utils.SP, i).armAssembly())
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
func store(size int) Size {
	if size == 1 {
		return B
	}
	return W
}

// load will return "LDRB" if the passed paramater is one, "LDR" otherwise.
func load(size int) Size {
	if size == 1 {
		return SB
	}
	return W
}
