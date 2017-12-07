package main

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

/**************** SIZES ****************/

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

/**************** SHIFTS ****************/

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

/**************** OPERAND 2 ****************/

// Operand2

/**************** ADDRESS ****************/

// Address

/**************** MOVEMENT ****************/

// Move
// E.g. MOV: if Cond then Rd := Op2; if Set then update flags

// MOV{Cond}{S} Rd, Op2

/**************** ARITHMETIC INSTRUCTIONS ****************/

// Add
// E.g. ADD: if Cond then Rd := Rs + Op2; if Set then update flags

// ADD{Cond}{S} Rd, Rn, Op2

// Subtract
// E.g. SUB: if Cond then Rd := Rs - Op2; if Set then update flags

// SUB{Cond}{S} Rd, Rn, Op2

// ReverseSubtract
// E.g. RSB: if Cond then Rd := Op2 - Rs; if Set then update flags

// RSB{Cond}{S} Rd, Rn, Op2

/**************** COMPARE INSTRUCTIONS ****************/

// Compare
// E.g. CMP: if Cond then update flags using Rn - Op2

// AND{Cond} Rn, Op2

/**************** LOGICAL INSTRUCTIONS ****************/

// LogicalAnd
// E.g. AND: if Cond then Rd := Rs AND Op2; if Set then update flags

// AND{Cond}{S} Rd, Rn, Op2

// ExclusiveOr
// E.g. EOR: if Cond then Rd := Rs EOR Op2; if Set then update flags

// EOR{Cond}{S} Rd, Rn, Op2

// LogicalOr
// E.g. ORR: if Cond then Rd := Rs OR Op2; if Set then update flags

// ORR{Cond}{S} Rd, Rn, Op2

/**************** MULTIPLY INSTRUCTIONS ****************/

// SignedMultiply
// E.g. SMULL: if Cond then RdHi,RdLo := Rm * Rs; if Set then update flags

// SMULL{Cond}{S} Rd, Rm, Rs {, Rn}

/**************** BRANCHING ****************/

// Branch
// E.g. B: if Cond then PC := Addr

// B{Cond} Addr

// BranchWithLink
// E.g BL: if Cond then R14 := address of next instruction, PC := Addr

// BL{Cond} Addr

/**************** SINGLE REGISTER DATA TRANSFER ****************/

// Load
// E.g. LDR: if Cond then Rd := [Addr]

// LDR{Cond}{Size} Rd, Addr

// Store
// E.g. STR: if Cond then [Addr] := Rd

// STR{Cond}{Size} Rd, Addr

/**************** STACK ****************/

// Push
// E.g. PUSH: if cond then PUSH Reglist

// PUSH{Cond} Reglist

// Pop
// E.g. POP: if cond then POP Reglist

// POP{Cond} Reglist
