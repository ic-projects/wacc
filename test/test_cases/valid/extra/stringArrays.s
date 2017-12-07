.data

msg_0:
	.word 3
	.ascii	"fst"
msg_1:
	.word 3
	.ascii	"snd"
msg_2:
	.word 44
	.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
msg_3:
	.word 45
	.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
msg_4:
	.word 5
	.ascii	"true\0"
msg_5:
	.word 6
	.ascii	"false\0"
msg_6:
	.word 1
	.ascii	"\0"
msg_7:
	.word 5
	.ascii	"%.*s\0"

.text

.global main
main:
	PUSH {lr}
	SUB sp, sp, #12
	LDR r4, =msg_0
	STR r4, [sp, #8]
	LDR r4, =msg_1
	STR r4, [sp, #4]
	LDR r0, =12
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #8]
	STR r5, [r4, #4]
	LDR r5, [sp, #4]
	STR r5, [r4, #8]
	LDR r5, =2
	STR r5, [r4]
	STR r4, [sp]
	ADD r4, sp, #0
	LDR r5, =0
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	LDR r5, [sp, #8]
	CMP r4, r5
	MOVEQ r4, #1
	MOVNE r4, #0
	MOV r0, r4
	BL p_print_bool
	BL p_print_ln
	ADD r4, sp, #0
	LDR r5, =0
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r5, =1
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5
	LDRSB r4, [r4]
	MOV r5, #'f'
	CMP r4, r5
	MOVEQ r4, #1
	MOVNE r4, #0
	MOV r0, r4
	BL p_print_bool
	BL p_print_ln
	ADD r4, sp, #0
	LDR r5, =1
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r5, =2
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5
	LDRSB r4, [r4]
	MOV r5, #'d'
	CMP r4, r5
	MOVEQ r4, #1
	MOVNE r4, #0
	MOV r0, r4
	BL p_print_bool
	BL p_print_ln
	ADD sp, sp, #12
	LDR r0, =0
	POP {pc}
	.ltorg
p_check_array_bounds:
	PUSH {lr}
	CMP r0, #0
	LDRLT r0, =msg_2
	BLLT p_throw_runtime_error
	LDR r1, [r1]
	CMP r0, r1
	LDRCS r0, =msg_3
	BLCS p_throw_runtime_error
	POP {pc}
p_print_bool:
	PUSH {lr}
	CMP r0, #0
	LDRNE r0, =msg_4
	LDREQ r0, =msg_5
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}
p_print_ln:
	PUSH {lr}
	LDR r0, =msg_6
	ADD r0, r0, #4
	BL puts
	MOV r0, #0
	BL fflush
	POP {pc}
p_throw_runtime_error:
	BL p_print_string
	MOV r0, #-1
	BL exit
p_print_string:
	PUSH {lr}
	LDR r1, [r0]
	ADD r2, r0, #4
	LDR r0, =msg_7
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}

