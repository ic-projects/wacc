.text

.global main
main:
	PUSH {lr}
	LDR r4, =256
	MOV r0, r4
	BL exit
	LDR r0, =0
	POP {pc}
	.ltorg

