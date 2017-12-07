.data

msg_0:
	.word 3
	.ascii	"%p\0"
msg_1:
	.word 1
	.ascii	"\0"
msg_2:
	.word 5
	.ascii	"%.*s\0"
msg_3:
	.word 44
	.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
msg_4:
	.word 45
	.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
msg_5:
	.word 50
	.ascii	"NullReferenceError: dereference a null reference\n\0"
msg_6:
	.word 3
	.ascii	"%d\0"

.text

.global main
main:
	PUSH {lr}
	SUB sp, sp, #20
	LDR r0, =16
	BL malloc
	MOV r4, r0
	LDR r5, =1
	STR r5, [r4, #4]
	LDR r5, =2
	STR r5, [r4, #8]
	LDR r5, =3
	STR r5, [r4, #12]
	LDR r5, =3
	STR r5, [r4]
	STR r4, [sp, #16]
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
	STR r4, [sp, #12]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #16]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, [sp, #12]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #8]
	LDR r4, [sp, #16]
	MOV r0, r4
	BL p_print_reference
	BL p_print_ln
	LDR r4, [sp, #12]
	MOV r0, r4
	BL p_print_string
	BL p_print_ln
	ADD r4, sp, #12
	LDR r5, =2
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5
	LDRSB r4, [r4]
	MOV r0, r4
	BL putchar
	BL p_print_ln
	LDR r4, [sp, #8]
	MOV r0, r4
	BL p_check_null_pointer
	LDR r4, [r4, #4]
	LDR r4, [r4]
	STR r4, [sp, #4]
	LDR r4, [sp, #4]
	MOV r0, r4
	BL p_print_string
	BL p_print_ln
	ADD r4, sp, #4
	LDR r5, =0
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5
	LDRSB r4, [r4]
	MOV r0, r4
	BL putchar
	BL p_print_ln
	ADD r4, sp, #4
	LDR r5, =1
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5
	LDRSB r4, [r4]
	MOV r0, r4
	BL putchar
	BL p_print_ln
	ADD r4, sp, #4
	LDR r5, =2
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5
	LDRSB r4, [r4]
	MOV r0, r4
	BL putchar
	BL p_print_ln
	LDR r4, [sp, #8]
	MOV r0, r4
	BL p_check_null_pointer
	LDR r4, [r4]
	LDR r4, [r4]
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
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	ADD r4, sp, #16
	LDR r5, =0
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	ADD r4, sp, #0
	LDR r5, =1
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	ADD r4, sp, #16
	LDR r5, =1
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	ADD r4, sp, #0
	LDR r5, =2
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	ADD r4, sp, #16
	LDR r5, =2
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	ADD sp, sp, #20
	LDR r0, =0
	POP {pc}
	.ltorg
p_print_reference:
	PUSH {lr}
	MOV r1, r0
	LDR r0, =msg_0
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}
p_print_ln:
	PUSH {lr}
	LDR r0, =msg_1
	ADD r0, r0, #4
	BL puts
	MOV r0, #0
	BL fflush
	POP {pc}
p_print_string:
	PUSH {lr}
	LDR r1, [r0]
	ADD r2, r0, #4
	LDR r0, =msg_2
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}
p_check_array_bounds:
	PUSH {lr}
	CMP r0, #0
	LDRLT r0, =msg_3
	BLLT p_throw_runtime_error
	LDR r1, [r1]
	CMP r0, r1
	LDRCS r0, =msg_4
	BLCS p_throw_runtime_error
	POP {pc}
p_check_null_pointer:
	PUSH {lr}
	CMP r0, #0
	LDREQ r0, =msg_5
	BLEQ p_throw_runtime_error
	POP {pc}
p_print_int:
	PUSH {lr}
	MOV r1, r0
	LDR r0, =msg_6
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}
p_throw_runtime_error:
	BL p_print_string
	MOV r0, #-1
	BL exit

