.data

msg_0:
	.word 5
	.ascii	"hello"
msg_1:
	.word 4
	.ascii	"test"

.text

.global main
f_f:
	PUSH {lr}
	LDR r4, [sp, #4]
	MOV r0, r4
	POP {pc}
	POP {pc}
	.ltorg
f_f2:
	PUSH {lr}
	LDR r4, [sp, #4]
	MOV r0, r4
	POP {pc}
	POP {pc}
	.ltorg
main:
	PUSH {lr}
	SUB sp, sp, #16
	LDR r4, =msg_0
	STR r4, [sp, #12]
	LDR r0, =7
	BL malloc
	MOV r4, r0
	MOV r5, #'a'
	STRB r5, [r4, #4]
	MOV r5, #'b'
	STRB r5, [r4, #5]
	MOV r5, #'c'
	STRB r5, [r4, #6]
	LDR r5, =3
	STR r5, [r4]
	STR r4, [sp, #8]
	LDR r4, [sp, #12]
	STR r4, [sp, #-4]!
	BL f_f
	ADD sp, sp, #4
	MOV r4, r0
	STR r4, [sp, #4]
	LDR r4, [sp, #8]
	STR r4, [sp, #-4]!
	BL f_f
	ADD sp, sp, #4
	MOV r4, r0
	STR r4, [sp, #12]
	LDR r4, [sp, #12]
	STR r4, [sp, #-4]!
	BL f_f2
	ADD sp, sp, #4
	MOV r4, r0
	STR r4, [sp, #8]
	LDR r4, [sp, #8]
	STR r4, [sp, #-4]!
	BL f_f2
	ADD sp, sp, #4
	MOV r4, r0
	STR r4, [sp, #4]
	LDR r4, =msg_1
	STR r4, [sp]
	ADD sp, sp, #16
	LDR r0, =0
	POP {pc}
	.ltorg

