-- Compiling...
-- Printing Assembly...
stringEqualsExpr.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 5
4		.ascii	"Hello"
5	msg_1:
6		.word 3
7		.ascii	"foo"
8	msg_2:
9		.word 3
10		.ascii	"foo"
11	msg_3:
12		.word 5
13		.ascii	"true\0"
14	msg_4:
15		.word 6
16		.ascii	"false\0"
17	msg_5:
18		.word 1
19		.ascii	"\0"
20	
21	.text
22	
23	.global main
24	main:
25		PUSH {lr}
26		SUB sp, sp, #13
27		LDR r4, =msg_0
28		STR r4, [sp, #9]
29		LDR r4, =msg_1
30		STR r4, [sp, #5]
31		LDR r4, =msg_2
32		STR r4, [sp, #1]
33		LDR r4, [sp, #9]
34		LDR r5, [sp, #9]
35		CMP r4, r5
36		MOVEQ r4, #1
37		MOVNE r4, #0
38		STRB r4, [sp]
39		LDRSB r4, [sp]
40		MOV r0, r4
41		BL p_print_bool
42		BL p_print_ln
43		LDR r4, [sp, #9]
44		LDR r5, [sp, #5]
45		CMP r4, r5
46		MOVEQ r4, #1
47		MOVNE r4, #0
48		MOV r0, r4
49		BL p_print_bool
50		BL p_print_ln
51		LDR r4, [sp, #5]
52		LDR r5, [sp, #1]
53		CMP r4, r5
54		MOVEQ r4, #1
55		MOVNE r4, #0
56		MOV r0, r4
57		BL p_print_bool
58		BL p_print_ln
59		ADD sp, sp, #13
60		LDR r0, =0
61		POP {pc}
62		.ltorg
63	p_print_bool:
64		PUSH {lr}
65		CMP r0, #0
66		LDRNE r0, =msg_3
67		LDREQ r0, =msg_4
68		ADD r0, r0, #4
69		BL printf
70		MOV r0, #0
71		BL fflush
72		POP {pc}
73	p_print_ln:
74		PUSH {lr}
75		LDR r0, =msg_5
76		ADD r0, r0, #4
77		BL puts
78		MOV r0, #0
79		BL fflush
80		POP {pc}
81	
===========================================================
-- Finished
