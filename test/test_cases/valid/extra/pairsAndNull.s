.data

msg_0:
	.word 50
	.ascii	"NullReferenceError: dereference a null reference\n\0"
msg_1:
	.word 5
	.ascii	"%.*s\0"

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
	SUB sp, sp, #48
	LDR r4, =0
	STR r4, [sp, #44]
	LDR r4, [sp, #44]
	MOV r0, r4
	BL p_check_null_pointer
	LDR r4, [r4]
	LDR r4, [r4]
	STR r4, [sp, #40]
	LDR r4, [sp, #40]
	STR r4, [sp, #-4]!
	BL f_f
	ADD sp, sp, #4
	MOV r4, r0
	STR r4, [sp, #36]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #44]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, [sp, #40]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #32]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #32]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, =0
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #28]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, =0
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, =0
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #24]
	LDR r4, =0
	STR r4, [sp, #-4]!
	BL f_f
	ADD sp, sp, #4
	MOV r4, r0
	STR r4, [sp, #20]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, =1
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, =0
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #16]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, =1
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, [sp, #16]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #12]
	LDR r4, [sp, #12]
	MOV r0, r4
	BL p_check_null_pointer
	LDR r4, [r4, #4]
	LDR r4, [r4]
	STR r4, [sp, #8]
	LDR r4, [sp, #16]
	STR r4, [sp, #-4]!
	BL f_f2
	ADD sp, sp, #4
	MOV r4, r0
	STR r4, [sp, #4]
	LDR r4, [sp, #44]
	MOV r0, r4
	BL p_check_null_pointer
	LDR r4, [r4]
	LDR r4, [r4]
	STR r4, [sp]
	ADD sp, sp, #48
	LDR r0, =0
	POP {pc}
	.ltorg
p_check_null_pointer:
	PUSH {lr}
	CMP r0, #0
	LDREQ r0, =msg_0
	BLEQ p_throw_runtime_error
	POP {pc}
p_throw_runtime_error:
	BL p_print_string
	MOV r0, #-1
	BL exit
p_print_string:
	PUSH {lr}
	LDR r1, [r0]
	ADD r2, r0, #4
	LDR r0, =msg_1
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}

