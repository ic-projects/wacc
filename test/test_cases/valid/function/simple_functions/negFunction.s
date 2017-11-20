-- Compiling...
-- Printing Assembly...
negFunction.s contents are:
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
15	f_neg:
16		PUSH {lr}
17		LDRSB r4, [sp, #4]
18		EOR r4, r4, #1
19		MOV r0, r4
20		POP {pc}
21		POP {pc}
22		.ltorg
23	main:
24		PUSH {lr}
25		SUB sp, sp, #1
26		MOV r4, #1
27		STRB r4, [sp]
28		LDRSB r4, [sp]
29		MOV r0, r4
30		BL p_print_bool
31		BL p_print_ln
32		LDRSB r4, [sp]
33		STRB r4, [sp, #-1]!
34		BL f_neg
35		ADD sp, sp, #1
36		MOV r4, r0
37		STRB r4, [sp]
38		LDRSB r4, [sp]
39		MOV r0, r4
40		BL p_print_bool
41		BL p_print_ln
42		LDRSB r4, [sp]
43		STRB r4, [sp, #-1]!
44		BL f_neg
45		ADD sp, sp, #1
46		MOV r4, r0
47		STRB r4, [sp]
48		LDRSB r4, [sp]
49		STRB r4, [sp, #-1]!
50		BL f_neg
51		ADD sp, sp, #1
52		MOV r4, r0
53		STRB r4, [sp]
54		LDRSB r4, [sp]
55		STRB r4, [sp, #-1]!
56		BL f_neg
57		ADD sp, sp, #1
58		MOV r4, r0
59		STRB r4, [sp]
60		LDRSB r4, [sp]
61		MOV r0, r4
62		BL p_print_bool
63		BL p_print_ln
64		ADD sp, sp, #1
65		LDR r0, =0
66		POP {pc}
67		.ltorg
68	p_print_bool:
69		PUSH {lr}
70		CMP r0, #0
71		LDRNE r0, =msg_0
72		LDREQ r0, =msg_1
73		ADD r0, r0, #4
74		BL printf
75		MOV r0, #0
76		BL fflush
77		POP {pc}
78	p_print_ln:
79		PUSH {lr}
80		LDR r0, =msg_2
81		ADD r0, r0, #4
82		BL puts
83		MOV r0, #0
84		BL fflush
85		POP {pc}
86	
===========================================================
-- Finished
