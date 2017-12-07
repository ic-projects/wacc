.data

msg_0:
	.word 3
	.ascii	"%p\0"
msg_1:
	.word 1
	.ascii	"\0"
msg_2:
	.word 44
	.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
msg_3:
	.word 45
	.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
msg_4:
	.word 50
	.ascii	"NullReferenceError: dereference a null reference\n\0"
msg_5:
	.word 3
	.ascii	"%d\0"
msg_6:
	.word 5
	.ascii	"%.*s\0"

.text

.global main
main:
	PUSH {lr}
	SUB sp, sp, #76
	LDR r4, =0
	STR r4, [sp, #72]
	LDR r4, [sp, #72]
	STR r4, [sp, #68]
	LDR r4, [sp, #72]
	STR r4, [sp, #64]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #72]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, [sp, #68]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #60]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #60]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, =0
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #56]
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
	STR r4, [sp, #52]
	LDR r4, [sp, #52]
	STR r4, [sp, #48]
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
	STR r4, [sp, #44]
	LDR r4, [sp, #44]
	MOV r0, r4
	BL p_print_reference
	BL p_print_ln
	LDR r0, =16
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #72]
	STR r5, [r4, #4]
	LDR r5, [sp, #68]
	STR r5, [r4, #8]
	LDR r5, [sp, #64]
	STR r5, [r4, #12]
	LDR r5, =3
	STR r5, [r4]
	STR r4, [sp, #40]
	LDR r4, [sp, #40]
	STR r4, [sp, #36]
	LDR r0, =16
	BL malloc
	MOV r4, r0
	LDR r5, =0
	STR r5, [r4, #4]
	LDR r5, =0
	STR r5, [r4, #8]
	LDR r5, [sp, #60]
	STR r5, [r4, #12]
	LDR r5, =3
	STR r5, [r4]
	STR r4, [sp, #32]
	LDR r0, =16
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #36]
	STR r5, [r4, #4]
	LDR r5, [sp, #32]
	STR r5, [r4, #8]
	LDR r5, [sp, #40]
	STR r5, [r4, #12]
	LDR r5, =3
	STR r5, [r4]
	STR r4, [sp, #28]
	LDR r4, [sp, #28]
	MOV r0, r4
	BL p_print_reference
	BL p_print_ln
	ADD r4, sp, #28
	LDR r5, =0
	LDR r4, [r4]
	MOV r0, r5
	MOV r1, r4
	BL p_check_array_bounds
	ADD r4, r4, #4
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	MOV r0, r4
	BL p_print_reference
	BL p_print_ln
	ADD r4, sp, #28
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
	ADD r4, r4, r5, LSL #2
	LDR r4, [r4]
	MOV r0, r4
	BL p_print_reference
	BL p_print_ln
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #32]
	LDR r5, [r5]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, =1
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #24]
	LDR r4, [sp, #24]
	MOV r0, r4
	BL p_check_null_pointer
	LDR r4, [r4]
	LDR r4, [r4]
	STR r4, [sp, #20]
	LDR r4, [sp, #20]
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r0, =12
	BL malloc
	MOV r4, r0
	ADD r5, sp, #28
	LDR r6, =0
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	ADD r6, sp, #28
	LDR r7, [sp, #32]
	LDR r7, [r7]
	LDR r6, [r6]
	MOV r0, r7
	MOV r1, r6
	BL p_check_array_bounds
	ADD r6, r6, #4
	ADD r6, r6, r7, LSL #2
	LDR r6, [r6]
	LDR r6, [r6]
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r5, [r5]
	STR r5, [r4, #4]
	ADD r5, sp, #28
	LDR r6, =1
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r6, =2
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r5, [r5]
	STR r5, [r4, #8]
	LDR r5, =2
	STR r5, [r4]
	STR r4, [sp, #16]
	LDR r0, =12
	BL malloc
	MOV r4, r0
	ADD r5, sp, #28
	LDR r6, [sp, #32]
	LDR r6, [r6]
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r5, [r5]
	STR r5, [r4, #4]
	LDR r5, [sp, #36]
	STR r5, [r4, #8]
	LDR r5, =2
	STR r5, [r4]
	STR r4, [sp, #12]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	ADD r5, sp, #12
	LDR r6, =1
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r6, =2
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r5, [r5]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	LDR r5, =0
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #8]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, =0
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	ADD r5, sp, #16
	ADD r6, sp, #28
	ADD r7, sp, #28
	ADD r8, sp, #28
	LDR r9, [sp, #28]
	LDR r9, [r9]
	LDR r8, [r8]
	MOV r0, r9
	MOV r1, r8
	BL p_check_array_bounds
	ADD r8, r8, #4
	ADD r8, r8, r9, LSL #2
	LDR r8, [r8]
	LDR r8, [r8]
	LDR r7, [r7]
	MOV r0, r8
	MOV r1, r7
	BL p_check_array_bounds
	ADD r7, r7, #4
	ADD r7, r7, r8, LSL #2
	LDR r7, [r7]
	LDR r7, [r7]
	LDR r6, [r6]
	MOV r0, r7
	MOV r1, r6
	BL p_check_array_bounds
	ADD r6, r6, #4
	ADD r6, r6, r7, LSL #2
	LDR r6, [r6]
	LDR r6, [r6]
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r5, [r5]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp, #4]
	LDR r0, =8
	BL malloc
	MOV r4, r0
	LDR r5, [sp, #20]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4]
	ADD r5, sp, #12
	LDR r6, =0
	LDR r5, [r5]
	MOV r0, r6
	MOV r1, r5
	BL p_check_array_bounds
	ADD r5, r5, #4
	ADD r5, r5, r6, LSL #2
	LDR r5, [r5]
	LDR r0, =4
	BL malloc
	STR r5, [r0]
	STR r0, [r4, #4]
	STR r4, [sp]
	ADD sp, sp, #76
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
p_check_null_pointer:
	PUSH {lr}
	CMP r0, #0
	LDREQ r0, =msg_4
	BLEQ p_throw_runtime_error
	POP {pc}
p_print_int:
	PUSH {lr}
	MOV r1, r0
	LDR r0, =msg_5
	ADD r0, r0, #4
	BL printf
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
	LDR r0, =msg_6
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}

