package location

import (
	"bytes"
	"fmt"
)

// Register is an enum which defines the registers according to the ARM11
// specification.
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

// FreeRegisters returns an array of the registers that functions are allowed to
// use.
func FreeRegisters() []Register {
	return []Register{R11, R10, R9, R8, R7, R6, R5, R4}
}

// ReturnRegisters returns an array of the registers used for passing and
// returning parameters from functions.
func ReturnRegisters() []Register {
	return []Register{R0, R1, R2, R3}
}

// String returns the string representation of this Register.
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

// NewRegisterStack creates and returns an empty RegisterStack.
func NewRegisterStack() *RegisterStack {
	return &RegisterStack{
		stack: []Register{},
	}
}

// NewRegisterStackWith creates and returns a RegisterStack with the stack set
// to the list of registers provided in the registers parameter.
func NewRegisterStackWith(registers []Register) *RegisterStack {
	return &RegisterStack{
		stack: registers,
	}
}

// Length returns the number of registers in this RegisterStack.
func (registerStack *RegisterStack) Length() int {
	return len(registerStack.stack)
}

// Pop removes and returns the Register at the top of this RegisterStack.
func (registerStack *RegisterStack) Pop() Register {
	if len(registerStack.stack) != 0 {
		register := registerStack.stack[len(registerStack.stack)-1]
		registerStack.stack = registerStack.stack[:len(registerStack.stack)-1]
		return register
	}
	fmt.Println("Internal compiler error")
	return UNDEFINED
}

// Peek returns the Register at the top of this RegisterStack.
func (registerStack *RegisterStack) Peek() Register {
	if len(registerStack.stack) != 0 {
		register := registerStack.stack[len(registerStack.stack)-1]
		return register
	}
	fmt.Println("Internal compiler error")
	return UNDEFINED
}

// Push adds the specified register to the top of this RegisterStack.
func (registerStack *RegisterStack) Push(register Register) {
	registerStack.stack = append(registerStack.stack, register)
}

// String returns the string representation of this RegisterStack in the form:
// "[ R0 R1 ... ]"
func (registerStack *RegisterStack) String() string {
	var buf bytes.Buffer
	buf.WriteString("[ ")
	for _, r := range registerStack.stack {
		buf.WriteString(r.String() + " ")
	}
	buf.WriteString("]")
	return buf.String()
}
