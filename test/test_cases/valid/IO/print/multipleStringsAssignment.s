-- Compiling...
-- Printing Assembly...
multipleStringsAssignment.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 2
4		.ascii	"Hi"
5	msg_1:
6		.word 2
7		.ascii	"Hi"
8	msg_2:
9		.word 6
10		.ascii	"s1 is "
11	msg_3:
12		.word 6
13		.ascii	"s2 is "
14	msg_4:
15		.word 18
16		.ascii	"They are the same."
17	msg_5:
18		.word 22
19		.ascii	"They are not the same."
20	msg_6:
21		.word 22
22		.ascii	"Now modify s1[0] = \'h\'"
23	msg_7:
24		.word 6
25		.ascii	"s1 is "
26	msg_8:
27		.word 6
28		.ascii	"s2 is "
29	msg_9:
30		.word 18
31		.ascii	"They are the same."
32	msg_10:
33		.word 22
34		.ascii	"They are not the same."
35	msg_11:
36		.word 5
37		.ascii	"%.*s\0"
38	msg_12:
39		.word 1
40		.ascii	"\0"
41	msg_13:
42		.word 44
43		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
44	msg_14:
45		.word 45
46		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
47	
48	.text
49	
50	.global main
51	main:
52		PUSH {lr}
53		SUB sp, sp, #8
54		LDR r4, =msg_0
55		STR r4, [sp, #4]
56		LDR r4, =msg_1
57		STR r4, [sp]
58		LDR r4, =msg_2
59		MOV r0, r4
60		BL p_print_string
61		LDR r4, [sp, #4]
62		MOV r0, r4
63		BL p_print_string
64		BL p_print_ln
65		LDR r4, =msg_3
66		MOV r0, r4
67		BL p_print_string
68		LDR r4, [sp]
69		MOV r0, r4
70		BL p_print_string
71		BL p_print_ln
72		LDR r4, [sp, #4]
73		LDR r5, [sp]
74		CMP r4, r5
75		MOVEQ r4, #1
76		MOVNE r4, #0
77		CMP r4, #0
78		BEQ L0
79		LDR r4, =msg_4
80		MOV r0, r4
81		BL p_print_string
82		BL p_print_ln
83		B L1
84	L0:
85		LDR r4, =msg_5
86		MOV r0, r4
87		BL p_print_string
88		BL p_print_ln
89	L1:
90		LDR r4, =msg_6
91		MOV r0, r4
92		BL p_print_string
93		BL p_print_ln
94		MOV r4, #'h'
95		ADD r5, sp, #4
96		LDR r6, =0
97		LDR r5, [r5]
98		MOV r0, r6
99		MOV r1, r5
100		BL p_check_array_bounds
101		ADD r5, r5, #4
102		ADD r5, r5, r6
103		STRB r4, [r5]
104		LDR r4, =msg_7
105		MOV r0, r4
106		BL p_print_string
107		LDR r4, [sp, #4]
108		MOV r0, r4
109		BL p_print_string
110		BL p_print_ln
111		LDR r4, =msg_8
112		MOV r0, r4
113		BL p_print_string
114		LDR r4, [sp]
115		MOV r0, r4
116		BL p_print_string
117		BL p_print_ln
118		LDR r4, [sp, #4]
119		LDR r6, [sp]
120		CMP r4, r6
121		MOVEQ r4, #1
122		MOVNE r4, #0
123		CMP r4, #0
124		BEQ L2
125		LDR r4, =msg_9
126		MOV r0, r4
127		BL p_print_string
128		BL p_print_ln
129		B L3
130	L2:
131		LDR r4, =msg_10
132		MOV r0, r4
133		BL p_print_string
134		BL p_print_ln
135	L3:
136		ADD sp, sp, #8
137		LDR r0, =0
138		POP {pc}
139		.ltorg
140	p_print_string:
141		PUSH {lr}
142		LDR r1, [r0]
143		ADD r2, r0, #4
144		LDR r0, =msg_11
145		ADD r0, r0, #4
146		BL printf
147		MOV r0, #0
148		BL fflush
149		POP {pc}
150	p_print_ln:
151		PUSH {lr}
152		LDR r0, =msg_12
153		ADD r0, r0, #4
154		BL puts
155		MOV r0, #0
156		BL fflush
157		POP {pc}
158	p_check_array_bounds:
159		PUSH {lr}
160		CMP r0, #0
161		LDRLT r0, =msg_13
162		BLLT p_throw_runtime_error
163		LDR r1, [r1]
164		CMP r0, r1
165		LDRCS r0, =msg_14
166		BLCS p_throw_runtime_error
167		POP {pc}
168	p_throw_runtime_error:
169		BL p_print_string
170		MOV r0, #-1
171		BL exit
172	
===========================================================
-- Finished
