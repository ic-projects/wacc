-- Compiling...
-- Printing Assembly...
mutualRecursion.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 12
4		.ascii	"r1: sending "
5	msg_1:
6		.word 13
7		.ascii	"r2: received "
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 3
13		.ascii	"%d\0"
14	msg_4:
15		.word 1
16		.ascii	"\0"
17	msg_5:
18		.word 82
19		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
20	
21	.text
22	
23	.global main
24	f_r1:
25		PUSH {lr}
26		LDR r4, [sp, #4]
27		LDR r5, =0
28		CMP r4, r5
29		MOVEQ r4, #1
30		MOVNE r4, #0
31		CMP r4, #0
32		BEQ L0
33		B L1
34	L0:
35		SUB sp, sp, #4
36		LDR r4, =msg_0
37		MOV r0, r4
38		BL p_print_string
39		LDR r4, [sp, #8]
40		MOV r0, r4
41		BL p_print_int
42		BL p_print_ln
43		LDR r4, [sp, #8]
44		STR r4, [sp, #-4]!
45		BL f_r2
46		ADD sp, sp, #4
47		MOV r4, r0
48		STR r4, [sp]
49		ADD sp, sp, #4
50	L1:
51		LDR r4, =42
52		MOV r0, r4
53		POP {pc}
54		POP {pc}
55		.ltorg
56	f_r2:
57		PUSH {lr}
58		SUB sp, sp, #4
59		LDR r4, =msg_1
60		MOV r0, r4
61		BL p_print_string
62		LDR r4, [sp, #8]
63		MOV r0, r4
64		BL p_print_int
65		BL p_print_ln
66		LDR r4, [sp, #8]
67		LDR r5, =1
68		SUBS r4, r4, r5
69		BLVS p_throw_overflow_error
70		STR r4, [sp, #-4]!
71		BL f_r1
72		ADD sp, sp, #4
73		MOV r4, r0
74		STR r4, [sp]
75		LDR r4, =44
76		MOV r0, r4
77		ADD sp, sp, #4
78		POP {pc}
79		POP {pc}
80		.ltorg
81	main:
82		PUSH {lr}
83		SUB sp, sp, #4
84		LDR r4, =0
85		STR r4, [sp]
86		LDR r4, =8
87		STR r4, [sp, #-4]!
88		BL f_r1
89		ADD sp, sp, #4
90		MOV r4, r0
91		STR r4, [sp]
92		ADD sp, sp, #4
93		LDR r0, =0
94		POP {pc}
95		.ltorg
96	p_print_string:
97		PUSH {lr}
98		LDR r1, [r0]
99		ADD r2, r0, #4
100		LDR r0, =msg_2
101		ADD r0, r0, #4
102		BL printf
103		MOV r0, #0
104		BL fflush
105		POP {pc}
106	p_print_int:
107		PUSH {lr}
108		MOV r1, r0
109		LDR r0, =msg_3
110		ADD r0, r0, #4
111		BL printf
112		MOV r0, #0
113		BL fflush
114		POP {pc}
115	p_print_ln:
116		PUSH {lr}
117		LDR r0, =msg_4
118		ADD r0, r0, #4
119		BL puts
120		MOV r0, #0
121		BL fflush
122		POP {pc}
123	p_throw_overflow_error:
124		LDR r0, =msg_5
125		BL p_throw_runtime_error
126	p_throw_runtime_error:
127		BL p_print_string
128		MOV r0, #-1
129		BL exit
130	
===========================================================
-- Finished
