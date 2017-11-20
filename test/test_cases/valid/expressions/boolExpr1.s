-- Compiling...
-- Printing Assembly...
boolExpr1.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 7
4		.ascii	"Correct"
5	msg_1:
6		.word 5
7		.ascii	"Wrong"
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 1
13		.ascii	"\0"
14	
15	.text
16	
17	.global main
18	main:
19		PUSH {lr}
20		SUB sp, sp, #1
21		MOV r4, #1
22		MOV r5, #0
23		AND r4, r4, r5
24		MOV r5, #1
25		MOV r6, #0
26		AND r5, r5, r6
27		ORR r4, r4, r5
28		EOR r4, r4, #1
29		STRB r4, [sp]
30		LDRSB r4, [sp]
31		MOV r5, #1
32		CMP r4, r5
33		MOVEQ r4, #1
34		MOVNE r4, #0
35		CMP r4, #0
36		BEQ L0
37		LDR r4, =msg_0
38		MOV r0, r4
39		BL p_print_string
40		BL p_print_ln
41		B L1
42	L0:
43		LDR r4, =msg_1
44		MOV r0, r4
45		BL p_print_string
46		BL p_print_ln
47	L1:
48		ADD sp, sp, #1
49		LDR r0, =0
50		POP {pc}
51		.ltorg
52	p_print_string:
53		PUSH {lr}
54		LDR r1, [r0]
55		ADD r2, r0, #4
56		LDR r0, =msg_2
57		ADD r0, r0, #4
58		BL printf
59		MOV r0, #0
60		BL fflush
61		POP {pc}
62	p_print_ln:
63		PUSH {lr}
64		LDR r0, =msg_3
65		ADD r0, r0, #4
66		BL puts
67		MOV r0, #0
68		BL fflush
69		POP {pc}
70	
===========================================================
-- Finished
