package location

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
