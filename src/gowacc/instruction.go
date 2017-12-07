package main

import "utils"

// Instruction is an interface for assembly instructions to implement.
type Instruction interface {
	// armAssembly returns an ARM assembly instruction
	armAssembly() string
}

/**************** CONDITIONS ****************/

// Condition is an enum of all possible instruction conditions.
type Condition int

const (
	// ALWAYS Default (No condition)
	ALWAYS Condition = iota
	// EQ Equal / zero
	EQ
	// NE Not equal / non-zero
	NE
	// GE Greater than or equal to
	GE
	// GT Greater than
	GT
	// LE Less than or equal to
	LE
	// LT Less than
	LT
)

/**************** OPERAND ****************/

// Shift is an enum of supported logical / arithmetic / rotate shifts.
type Shift int

const (
	// NONE Default (No shift)
	NONE Shift = iota
	// LSL Logical Shift Left
	LSL
	// ASR Arithmetic Shift Right
	ASR
)

// Operand is a shifted register or an immediate operand.
type Operand interface {
	// armOperand returns an ARM assembly operand2
	armOperand() string
}

// RegisterOperand is a struct holding a register and shift information.
type RegisterOperand struct {
	Reg       utils.Register
	ShiftType Shift
	ShiftVal  int
}

// ImmediateOperand is an immediate value.
type ImmediateOperand int

/**************** ADDRESS ****************/

// Address is a memory address.
type Address interface {
	// armOperand returns an ARM assembly address
	armAddress() string
}

// LabelAddress is a label.
type LabelAddress string

// ConstantAddress is an integer address.
type ConstantAddress int

// PreIndexedAddress is a register value + an offset.
type PreIndexedAddress struct {
	Reg    utils.Register
	Offset int
}

/**************** MOVEMENT ****************/

// Move is a struct for the MOV instruction.
//
// E.g. MOV: if Cond then Rd := Op2; if Set then update flags
type Move struct {
	Cond Condition
	Set  bool
	Rd   utils.Register
	Op2  Operand
}

// MOV{Cond}{S} Rd, Op2

/**************** ARITHMETIC INSTRUCTIONS ****************/

// ArithmeticInstructionType represents a single arithmetic instruction.
type ArithmeticInstructionType int

const (
	// ADD Add
	ADD Condition = iota
	// SUB Subtract
	SUB
	// RSB Reverse subtract
	RSB
)

// ArithmeticInstruction is the base struct for all arithmetic instructions.
//
// E.g. ADD: if Cond then Rd := Rs + Op2; if Set then update flags
// E.g. SUB: if Cond then Rd := Rs - Op2; if Set then update flags
// E.g. RSB: if Cond then Rd := Op2 - Rs; if Set then update flags
type ArithmeticInstruction struct {
	Instr ArithmeticInstructionType
	Cond  Condition
	Set   bool
	Rd    utils.Register
	Rs    utils.Register
	Op2   Operand
}

// Instr{Cond}{S} Rd, Rs, Op2

/**************** COMPARE INSTRUCTIONS ****************/

// Compare is a struct for the CMP instruction.
//
// E.g. CMP: if Cond then update flags using Rn - Op2
type Compare struct {
	Cond Condition
	Rs   utils.Register
	Op2  Operand
}

// CMP{Cond} Rn, Op2

/**************** LOGICAL INSTRUCTIONS ****************/

// LogicalInstructionType represents a single logical instruction.
type LogicalInstructionType int

const (
	// AND Logical and
	AND Condition = iota
	// EOR Exclusive or
	EOR
	// ORR Logical or
	ORR
)

// LogicalInstruction is the base struct for all logical instructions.
//
// E.g. AND: if Cond then Rd := Rs AND Op2; if Set then update flags
// E.g. EOR: if Cond then Rd := Rs EOR Op2; if Set then update flags
// E.g. ORR: if Cond then Rd := Rs OR Op2; if Set then update flags
type LogicalInstruction struct {
	Cond Condition
	Set  bool
	Rd   utils.Register
	Rs   utils.Register
	Op2  Operand
}

// Instr{Cond}{S} Rd, Rs, Op2

/**************** MULTIPLY INSTRUCTIONS ****************/

// SignedMultiply is a struct for the SMULL instruction.
// E.g. SMULL: if Cond then RdHi,RdLo := Rm * Rs; if Set then update flags
type SignedMultiply struct {
	Cond Condition
	Set  bool
	RdLo utils.Register
	RdHi utils.Register
	Rm   utils.Register
	Rs   utils.Register
}

// SMULL{Cond}{S} RdLo, RdHi, Rm {, Rs}

/**************** BRANCHING ****************/

// Branch is a struct for the B instruction.
//
// E.g. B: if Cond then PC := Addr
// E.g. BL: if Cond then R14 := address of next instruction, PC := Addr
type Branch struct {
	Cond Condition
	Link bool
	Addr Address
}

// B{Link}{Cond} Addr

/**************** SINGLE REGISTER DATA TRANSFER ****************/

// DataTransferInstructionType represents a single data transfer instruction.
type DataTransferInstructionType int

const (
	// LDR Load
	LDR DataTransferInstructionType = iota
	// STR Store
	STR
)

// Size is an enum of load / store sizes.
type Size int

const (
	// W Default (Word)
	W Size = iota
	// B Unsigned Byte
	B
	// SB Signed Byte
	SB
)

// DataTransferInstruction is the base struct for all data transfer
// instructions.
//
// E.g. LDR: if Cond then Rd := [Addr]
// E.g. STR: if Cond then [Addr] := Rd
type DataTransferInstruction struct {
	Instr DataTransferInstructionType
	Cond  Condition
	Size  Size
	Rd    utils.Register
	Addr  Address
}

// Instr{Cond}{Size} Rd, Addr

/**************** STACK ****************/

// StackInstructionType represents a single stack instruction.
type StackInstructionType int

const (
	// PUSH Push
	PUSH DataTransferInstructionType = iota
	// POP Pop
	POP
)

// StackInstruction is the base struct for all stack instructions.
//
// E.g. PUSH: if Cond then PUSH Reglist
// E.g. POP: if Cond then POP Reglist
type StackInstruction struct {
	Cond    Condition
	Reglist []utils.Register
}

// Instr{Cond} Reglist
