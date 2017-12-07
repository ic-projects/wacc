.data

msg_0:
	.word 82
	.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
msg_1:
	.word 3
	.ascii	"%d\0"
msg_2:
	.word 1
	.ascii	"\0"
msg_3:
	.word 45
	.ascii	"DivideByZeroError: divide or modulo by zero\n\0"
msg_4:
	.word 5
	.ascii	"%.*s\0"

.text

.global main
main:
	PUSH {lr}
	LDR r4, =1
	LDR r5, =2
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =3
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =4
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =5
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	SMULL r4, r5, r4, r5
	CMP r5, r4, ASR #31
	BLNE p_throw_overflow_error
	LDR r5, =3
	SMULL r4, r5, r4, r5
	CMP r5, r4, ASR #31
	BLNE p_throw_overflow_error
	LDR r5, =4
	SMULL r4, r5, r4, r5
	CMP r5, r4, ASR #31
	BLNE p_throw_overflow_error
	LDR r5, =5
	SMULL r4, r5, r4, r5
	CMP r5, r4, ASR #31
	BLNE p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =3
	LDR r6, =4
	SMULL r5, r6, r5, r6
	CMP r6, r5, ASR #31
	BLNE p_throw_overflow_error
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =5
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =3
	LDR r6, =4
	SMULL r5, r6, r5, r6
	CMP r6, r5, ASR #31
	BLNE p_throw_overflow_error
	LDR r6, =5
	MOV r0, r5
	MOV r1, r6
	BL p_check_divide_by_zero
	BL __aeabi_idiv
	MOV r5, r0
	SUBS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =6
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =7
	SUBS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =3
	LDR r6, =4
	SMULL r5, r6, r5, r6
	CMP r6, r5, ASR #31
	BLNE p_throw_overflow_error
	LDR r6, =5
	MOV r0, r5
	MOV r1, r6
	BL p_check_divide_by_zero
	BL __aeabi_idiv
	MOV r5, r0
	SUBS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =6
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =7
	SUBS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	SMULL r4, r5, r4, r5
	CMP r5, r4, ASR #31
	BLNE p_throw_overflow_error
	LDR r5, =3
	LDR r6, =4
	SMULL r5, r6, r5, r6
	CMP r6, r5, ASR #31
	BLNE p_throw_overflow_error
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	MOV r0, r4
	MOV r1, r5
	BL p_check_divide_by_zero
	BL __aeabi_idiv
	MOV r4, r0
	LDR r5, =3
	LDR r6, =4
	SMULL r5, r6, r5, r6
	CMP r6, r5, ASR #31
	BLNE p_throw_overflow_error
	LDR r6, =5
	SMULL r5, r6, r5, r6
	CMP r6, r5, ASR #31
	BLNE p_throw_overflow_error
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	LDR r6, =3
	ADDS r5, r5, r6
	BLVS p_throw_overflow_error
	LDR r6, =4
	SMULL r5, r6, r5, r6
	CMP r6, r5, ASR #31
	BLNE p_throw_overflow_error
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =5
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =3
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	LDR r5, =4
	SMULL r4, r5, r4, r5
	CMP r5, r4, ASR #31
	BLNE p_throw_overflow_error
	LDR r5, =5
	ADDS r4, r4, r5
	BLVS p_throw_overflow_error
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r4, =1
	LDR r5, =2
	LDR r6, =3
	LDR r7, =4
	SMULL r6, r7, r6, r7
	CMP r7, r6, ASR #31
	BLNE p_throw_overflow_error
	ADDS r5, r5, r6
	BLVS p_throw_overflow_error
	LDR r6, =5
	ADDS r5, r5, r6
	BLVS p_throw_overflow_error
	MOV r0, r4
	MOV r1, r5
	BL p_check_divide_by_zero
	BL __aeabi_idiv
	MOV r4, r0
	MOV r0, r4
	BL p_print_int
	BL p_print_ln
	LDR r0, =0
	POP {pc}
	.ltorg
p_throw_overflow_error:
	LDR r0, =msg_0
	BL p_throw_runtime_error
p_print_int:
	PUSH {lr}
	MOV r1, r0
	LDR r0, =msg_1
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}
p_print_ln:
	PUSH {lr}
	LDR r0, =msg_2
	ADD r0, r0, #4
	BL puts
	MOV r0, #0
	BL fflush
	POP {pc}
p_check_divide_by_zero:
	PUSH {lr}
	CMP r1, #0
	LDREQ r0, =msg_3
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
	LDR r0, =msg_4
	ADD r0, r0, #4
	BL printf
	MOV r0, #0
	BL fflush
	POP {pc}

