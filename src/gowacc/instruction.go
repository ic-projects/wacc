package main

import (
	"bytes"
	"fmt"
	"utils"
)

const (
	// ERROR is the string to be printed when no instruction exists.
	ERROR string = "ERROR"
	// INDENT is the string to be printed before each line of assembly.
	INDENT string = "\t"
)

// Instruction is an interface for assembly instructions to implement.
type Instruction interface {
	// armAssembly returns an ARM assembly instruction
	armAssembly() string
}

/**************** CONDITIONS ****************/

// Condition is an enum of all possible instruction conditions.
type Condition int

const (
	// AL Default (Always / No condition)
	AL Condition = iota
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
	// VS Overflow
	VS
)

func (cond Condition) armCondition() string {
	switch cond {
	case EQ:
		return "EQ"
	case NE:
		return "NE"
	case GE:
		return "GE"
	case GT:
		return "GT"
	case LE:
		return "LE"
	case LT:
		return "LT"
	case VS:
		return "VS"
	default:
		return ""
	}
}

/**************** SET ****************/

func armSet(set bool) string {
	if set {
		return "S"
	}
	return ""
}

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

func (shift Shift) armShift() string {
	switch shift {
	case LSL:
		return "LSL"
	case ASR:
		return "ASR"
	default:
		return ""
	}
}

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

// Reg{, ShiftType #ShiftVal}
func (op RegisterOperand) armOperand() string {
	if op.ShiftType == NONE || op.ShiftVal == 0 {
		return op.Reg.String()
	}
	return fmt.Sprintf(
		"%s, %s %d",
		op.Reg.String(),
		op.ShiftType.armShift(),
		op.ShiftVal,
	)
}

// ImmediateOperand is an immediate integer value.
type ImmediateOperand int

// #ImmediateOperand
func (op ImmediateOperand) armOperand() string {
	return fmt.Sprintf("#%d", int(op))
}

// ImmediateCharOperand is an immediate character value.
type ImmediateCharOperand string

// #'ImmediateCharOperand'
func (op ImmediateCharOperand) armOperand() string {
	return fmt.Sprintf("#'%s'", string(op))
}

/**************** ADDRESS ****************/

// Address is a memory address.
type Address interface {
	// armOperand returns an ARM assembly address
	armAddress() string
}

// LabelAddress is a label.
type LabelAddress string

// =LabelAddress
func (label LabelAddress) armAddress() string {
	return fmt.Sprintf("=%s", string(label))
}

// ConstantAddress is an integer address.
type ConstantAddress int

// =ConstantAddress
func (addr ConstantAddress) armAddress() string {
	return fmt.Sprintf("=%d", int(addr))
}

// RegisterAddress is a register value + an offset.
type RegisterAddress struct {
	Reg    utils.Register
	Offset int
}

// [Reg{, Offset}]
func (addr RegisterAddress) armAddress() string {
	if addr.Offset == 0 {
		return fmt.Sprintf("[%s]", addr.Reg.String())
	}
	return fmt.Sprintf("[%s, #%d]", addr.Reg.String(), addr.Offset)
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
func (instr Move) armAssembly() string {
	return fmt.Sprintf(
		"%sMOV%s%s %s, %s",
		INDENT,
		instr.Cond.armCondition(),
		armSet(instr.Set),
		instr.Rd.String(),
		instr.Op2.armOperand(),
	)
}

/**************** ARITHMETIC INSTRUCTIONS ****************/

// ArithmeticInstructionType represents a single arithmetic instruction.
type ArithmeticInstructionType int

const (
	// ADD Add
	ADD ArithmeticInstructionType = iota
	// SUB Subtract
	SUB
	// RSB Reverse subtract
	RSB
)

func (instr ArithmeticInstructionType) armInstruction() string {
	switch instr {
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case RSB:
		return "RSB"
	}
	return ERROR
}

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
func (instr ArithmeticInstruction) armAssembly() string {
	return fmt.Sprintf(
		"%s%s%s%s %s, %s, %s",
		INDENT,
		instr.Instr.armInstruction(),
		instr.Cond.armCondition(),
		armSet(instr.Set),
		instr.Rd.String(),
		instr.Rs.String(),
		instr.Op2.armOperand(),
	)
}

/**************** COMPARE INSTRUCTIONS ****************/

// Compare is a struct for the CMP instruction.
//
// E.g. CMP: if Cond then update flags using Rn - Op2
type Compare struct {
	Cond Condition
	Rn   utils.Register
	Op2  Operand
}

// CMP{Cond} Rn, Op2
func (instr Compare) armAssembly() string {
	return fmt.Sprintf(
		"%sCMP%s %s, %s",
		INDENT,
		instr.Cond.armCondition(),
		instr.Rn.String(),
		instr.Op2.armOperand(),
	)
}

/**************** LOGICAL INSTRUCTIONS ****************/

// LogicalInstructionType represents a single logical instruction.
type LogicalInstructionType int

const (
	// AND Logical and
	AND LogicalInstructionType = iota
	// EOR Exclusive or
	EOR
	// ORR Logical or
	ORR
)

func (instr LogicalInstructionType) armInstruction() string {
	switch instr {
	case AND:
		return "AND"
	case EOR:
		return "EOR"
	case ORR:
		return "ORR"
	}
	return ERROR
}

// LogicalInstruction is the base struct for all logical instructions.
//
// E.g. AND: if Cond then Rd := Rs AND Op2; if Set then update flags
// E.g. EOR: if Cond then Rd := Rs EOR Op2; if Set then update flags
// E.g. ORR: if Cond then Rd := Rs OR Op2; if Set then update flags
type LogicalInstruction struct {
	Instr LogicalInstructionType
	Cond  Condition
	Set   bool
	Rd    utils.Register
	Rs    utils.Register
	Op2   Operand
}

// Instr{Cond}{S} Rd, Rs, Op2
func (instr LogicalInstruction) armAssembly() string {
	return fmt.Sprintf(
		"%s%s%s%s %s, %s, %s",
		INDENT,
		instr.Instr.armInstruction(),
		instr.Cond.armCondition(),
		armSet(instr.Set),
		instr.Rd.String(),
		instr.Rs.String(),
		instr.Op2.armOperand(),
	)
}

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

// SMULL{Cond}{S} RdLo, RdHi, Rm{, Rs}
func (instr SignedMultiply) armAssembly() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(
		"%sSMULL%s%s %s, %s, %s",
		INDENT,
		instr.Cond.armCondition(),
		armSet(instr.Set),
		instr.RdLo.String(),
		instr.RdHi.String(),
		instr.Rm.String(),
	))
	if instr.Rs != utils.UNDEFINED {
		buf.WriteString(fmt.Sprintf(", %s", instr.Rs.String()))
	}
	return buf.String()
}

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

// NewBranch builds a Branch.
func NewBranch(label string) Branch {
	return NewCondBranch(AL, label)
}

// NewCondBranch builds a Branch with a condition.
func NewCondBranch(cond Condition, label string) Branch {
	return Branch{cond, true, LabelAddress(label)}
}

// NewBranchL builds a Branch with a link.
func NewBranchL(label string) Branch {
	return NewCondBranchL(AL, label)
}

// NewCondBranchL builds a Branch with a condition and a link.
func NewCondBranchL(cond Condition, label string) Branch {
	return Branch{cond, true, LabelAddress(label)}
}

func armLink(link bool) string {
	if link {
		return "L"
	}
	return ""
}

// B{Link}{Cond} Addr
func (instr Branch) armAssembly() string {
	return fmt.Sprintf(
		"%sB%s%s %s",
		INDENT,
		armLink(instr.Link),
		instr.Cond.armCondition(),
		instr.Addr.armAddress()[1:],
	)
}

/**************** SINGLE REGISTER DATA TRANSFER ****************/

// DataTransferInstructionType represents a single data transfer instruction.
type DataTransferInstructionType int

const (
	// LDR Load
	LDR DataTransferInstructionType = iota
	// STR Store
	STR
)

func (instr DataTransferInstructionType) armInstruction() string {
	switch instr {
	case LDR:
		return "LDR"
	case STR:
		return "STR"
	}
	return ERROR
}

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

func (size Size) armSize() string {
	switch size {
	case B:
		return "B"
	case SB:
		return "SB"
	default:
		return ""
	}
}

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

// NewLoad builds a LDR instruction.
func NewLoad(
	size Size,
	reg utils.Register,
	addr Address,
) DataTransferInstruction {
	return DataTransferInstruction{LDR, AL, size, reg, addr}
}

// NewLoadReg builds a LDR instruction from an address held by a register.
func NewLoadReg(
	size Size,
	r1 utils.Register,
	r2 utils.Register,
) DataTransferInstruction {
	return NewLoad(size, r1, RegisterAddress{r2, 0})
}

// NewLoadRegOffset builds a LDR instruction to an address held by a register.
func NewLoadRegOffset(
	size Size,
	r1 utils.Register,
	r2 utils.Register,
	offset int,
) DataTransferInstruction {
	return NewLoad(size, r1, RegisterAddress{r2, offset})
}

// NewStore builds a STR instruction.
func NewStore(
	size Size,
	reg utils.Register,
	addr Address,
) DataTransferInstruction {
	return DataTransferInstruction{STR, AL, size, reg, addr}
}

// NewStoreReg builds a STR instruction to an address held by a register.
func NewStoreReg(
	size Size,
	r1 utils.Register,
	r2 utils.Register,
) DataTransferInstruction {
	return NewStore(size, r1, RegisterAddress{r2, 0})
}

// NewStoreRegOffset builds a STR instruction to an address held by a register.
func NewStoreRegOffset(
	size Size,
	r1 utils.Register,
	r2 utils.Register,
	offset int,
) DataTransferInstruction {
	return NewStore(size, r1, RegisterAddress{r2, offset})
}

// Instr{Cond}{Size} Rd, Addr
func (instr DataTransferInstruction) armAssembly() string {
	return fmt.Sprintf(
		"%s%s%s%s %s, %s",
		INDENT,
		instr.Instr.armInstruction(),
		instr.Cond.armCondition(),
		instr.Size.armSize(),
		instr.Rd.String(),
		instr.Addr.armAddress(),
	)
}

/**************** STACK ****************/

// StackInstructionType represents a single stack instruction.
type StackInstructionType int

const (
	// PUSH Push
	PUSH StackInstructionType = iota
	// POP Pop
	POP
)

func (instr StackInstructionType) armInstruction() string {
	switch instr {
	case PUSH:
		return "PUSH"
	case POP:
		return "POP"
	}
	return ERROR
}

// StackInstruction is the base struct for all stack instructions.
//
// E.g. PUSH: if Cond then PUSH Reglist
// E.g. POP: if Cond then POP Reglist
type StackInstruction struct {
	Instr   StackInstructionType
	Cond    Condition
	Reglist []utils.Register
}

// NewPush builds a PUSH instruction with a single register.
func NewPush(reg utils.Register) StackInstruction {
	return StackInstruction{PUSH, AL, []utils.Register{reg}}
}

// NewPop builds a POP instruction with a single register.
func NewPop(reg utils.Register) StackInstruction {
	return StackInstruction{POP, AL, []utils.Register{reg}}
}

// Instr{Cond} Reglist
func (instr StackInstruction) armAssembly() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf(
		"%s%s%s {",
		INDENT,
		instr.Instr.armInstruction(),
		instr.Cond.armCondition(),
	))
	for i, r := range instr.Reglist {
		if i == 0 {
			buf.WriteString(r.String())
		} else {
			buf.WriteString(fmt.Sprintf(", %s", r.String()))
		}
	}
	buf.WriteString("}")
	return buf.String()
}
