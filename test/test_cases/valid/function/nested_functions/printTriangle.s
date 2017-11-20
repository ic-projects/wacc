-- Compiling...
-- Printing Assembly...
printTriangle.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 1
4		.ascii	"-"
5	msg_1:
6		.word 0
7		.ascii	""
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 82
13		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
14	msg_4:
15		.word 1
16		.ascii	"\0"
17	
18	.text
19	
20	.global main
21	f_f:
22		PUSH {lr}
23		LDR r4, [sp, #4]
24		LDR r5, =0
25		CMP r4, r5
26		MOVEQ r4, #1
27		MOVNE r4, #0
28		CMP r4, #0
29		BEQ L0
30		B L1
31	L0:
32		SUB sp, sp, #8
33		LDR r4, [sp, #12]
34		STR r4, [sp, #4]
35		B L2
36	L3:
37		LDR r4, =msg_0
38		MOV r0, r4
39		BL p_print_string
40		LDR r4, [sp, #4]
41		LDR r5, =1
42		SUBS r4, r4, r5
43		BLVS p_throw_overflow_error
44		STR r4, [sp, #4]
45	L2:
46		LDR r4, [sp, #4]
47		LDR r5, =0
48		CMP r4, r5
49		MOVGT r4, #1
50		MOVLE r4, #0
51		CMP r4, #1
52		BEQ L3
53		LDR r4, =msg_1
54		MOV r0, r4
55		BL p_print_string
56		BL p_print_ln
57		LDR r4, [sp, #12]
58		LDR r5, =1
59		SUBS r4, r4, r5
60		BLVS p_throw_overflow_error
61		STR r4, [sp, #-4]!
62		BL f_f
63		ADD sp, sp, #4
64		MOV r4, r0
65		STR r4, [sp]
66		ADD sp, sp, #8
67	L1:
68		LDR r4, =0
69		MOV r0, r4
70		POP {pc}
71		POP {pc}
72		.ltorg
73	main:
74		PUSH {lr}
75		SUB sp, sp, #4
76		LDR r4, =8
77		STR r4, [sp, #-4]!
78		BL f_f
79		ADD sp, sp, #4
80		MOV r4, r0
81		STR r4, [sp]
82		ADD sp, sp, #4
83		LDR r0, =0
84		POP {pc}
85		.ltorg
86	p_print_string:
87		PUSH {lr}
88		LDR r1, [r0]
89		ADD r2, r0, #4
90		LDR r0, =msg_2
91		ADD r0, r0, #4
92		BL printf
93		MOV r0, #0
94		BL fflush
95		POP {pc}
96	p_throw_overflow_error:
97		LDR r0, =msg_3
98		BL p_throw_runtime_error
99	p_print_ln:
100		PUSH {lr}
101		LDR r0, =msg_4
102		ADD r0, r0, #4
103		BL puts
104		MOV r0, #0
105		BL fflush
106		POP {pc}
107	p_throw_runtime_error:
108		BL p_print_string
109		MOV r0, #-1
110		BL exit
111	
===========================================================
-- Finished
