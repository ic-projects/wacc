-- Compiling...
-- Printing Assembly...
lessCharExpr.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 5
4		.ascii	"true\0"
5	msg_1:
6		.word 6
7		.ascii	"false\0"
8	msg_2:
9		.word 1
10		.ascii	"\0"
11	
12	.text
13	
14	.global main
15	main:
16		PUSH {lr}
17		SUB sp, sp, #3
18		MOV r4, #'a'
19		STRB r4, [sp, #2]
20		MOV r4, #'e'
21		STRB r4, [sp, #1]
22		MOV r4, #'c'
23		STRB r4, [sp]
24		LDRSB r4, [sp, #2]
25		LDRSB r5, [sp, #1]
26		CMP r4, r5
27		MOVLT r4, #1
28		MOVGE r4, #0
29		MOV r0, r4
30		BL p_print_bool
31		BL p_print_ln
32		LDRSB r4, [sp, #1]
33		LDRSB r5, [sp]
34		CMP r4, r5
35		MOVLT r4, #1
36		MOVGE r4, #0
37		MOV r0, r4
38		BL p_print_bool
39		BL p_print_ln
40		ADD sp, sp, #3
41		LDR r0, =0
42		POP {pc}
43		.ltorg
44	p_print_bool:
45		PUSH {lr}
46		CMP r0, #0
47		LDRNE r0, =msg_0
48		LDREQ r0, =msg_1
49		ADD r0, r0, #4
50		BL printf
51		MOV r0, #0
52		BL fflush
53		POP {pc}
54	p_print_ln:
55		PUSH {lr}
56		LDR r0, =msg_2
57		ADD r0, r0, #4
58		BL puts
59		MOV r0, #0
60		BL fflush
61		POP {pc}
62	
===========================================================
-- Finished
