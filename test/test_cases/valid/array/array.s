-- Compiling...
-- Printing Assembly...
array.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 4
4		.ascii	" = {"
5	msg_1:
6		.word 2
7		.ascii	", "
8	msg_2:
9		.word 1
10		.ascii	"}"
11	msg_3:
12		.word 44
13		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
14	msg_4:
15		.word 45
16		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
17	msg_5:
18		.word 82
19		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
20	msg_6:
21		.word 3
22		.ascii	"%p\0"
23	msg_7:
24		.word 5
25		.ascii	"%.*s\0"
26	msg_8:
27		.word 3
28		.ascii	"%d\0"
29	msg_9:
30		.word 1
31		.ascii	"\0"
32	
33	.text
34	
35	.global main
36	main:
37		PUSH {lr}
38		SUB sp, sp, #8
39		LDR r0, =44
40		BL malloc
41		MOV r4, r0
42		LDR r5, =0
43		STR r5, [r4, #4]
44		LDR r5, =0
45		STR r5, [r4, #8]
46		LDR r5, =0
47		STR r5, [r4, #12]
48		LDR r5, =0
49		STR r5, [r4, #16]
50		LDR r5, =0
51		STR r5, [r4, #20]
52		LDR r5, =0
53		STR r5, [r4, #24]
54		LDR r5, =0
55		STR r5, [r4, #28]
56		LDR r5, =0
57		STR r5, [r4, #32]
58		LDR r5, =0
59		STR r5, [r4, #36]
60		LDR r5, =0
61		STR r5, [r4, #40]
62		LDR r5, =10
63		STR r5, [r4]
64		STR r4, [sp, #4]
65		LDR r4, =0
66		STR r4, [sp]
67		B L0
68	L1:
69		LDR r4, [sp]
70		ADD r5, sp, #4
71		LDR r6, [sp]
72		LDR r5, [r5]
73		MOV r0, r6
74		MOV r1, r5
75		BL p_check_array_bounds
76		ADD r5, r5, #4
77		ADD r5, r5, r6, LSL #2
78		STR r4, [r5]
79		LDR r4, [sp]
80		LDR r6, =1
81		ADDS r4, r4, r6
82		BLVS p_throw_overflow_error
83		STR r4, [sp]
84	L0:
85		LDR r4, [sp]
86		LDR r6, =10
87		CMP r4, r6
88		MOVLT r4, #1
89		MOVGE r4, #0
90		CMP r4, #1
91		BEQ L1
92		LDR r4, [sp, #4]
93		MOV r0, r4
94		BL p_print_reference
95		LDR r4, =msg_0
96		MOV r0, r4
97		BL p_print_string
98		LDR r4, =0
99		STR r4, [sp]
100		B L2
101	L3:
102		ADD r4, sp, #4
103		LDR r6, [sp]
104		LDR r4, [r4]
105		MOV r0, r6
106		MOV r1, r4
107		BL p_check_array_bounds
108		ADD r4, r4, #4
109		ADD r4, r4, r6, LSL #2
110		LDR r4, [r4]
111		MOV r0, r4
112		BL p_print_int
113		LDR r4, [sp]
114		LDR r6, =9
115		CMP r4, r6
116		MOVLT r4, #1
117		MOVGE r4, #0
118		CMP r4, #0
119		BEQ L4
120		LDR r4, =msg_1
121		MOV r0, r4
122		BL p_print_string
123		B L5
124	L4:
125	L5:
126		LDR r4, [sp]
127		LDR r6, =1
128		ADDS r4, r4, r6
129		BLVS p_throw_overflow_error
130		STR r4, [sp]
131	L2:
132		LDR r4, [sp]
133		LDR r6, =10
134		CMP r4, r6
135		MOVLT r4, #1
136		MOVGE r4, #0
137		CMP r4, #1
138		BEQ L3
139		LDR r4, =msg_2
140		MOV r0, r4
141		BL p_print_string
142		BL p_print_ln
143		ADD sp, sp, #8
144		LDR r0, =0
145		POP {pc}
146		.ltorg
147	p_check_array_bounds:
148		PUSH {lr}
149		CMP r0, #0
150		LDRLT r0, =msg_3
151		BLLT p_throw_runtime_error
152		LDR r1, [r1]
153		CMP r0, r1
154		LDRCS r0, =msg_4
155		BLCS p_throw_runtime_error
156		POP {pc}
157	p_throw_overflow_error:
158		LDR r0, =msg_5
159		BL p_throw_runtime_error
160	p_print_reference:
161		PUSH {lr}
162		MOV r1, r0
163		LDR r0, =msg_6
164		ADD r0, r0, #4
165		BL printf
166		MOV r0, #0
167		BL fflush
168		POP {pc}
169	p_print_string:
170		PUSH {lr}
171		LDR r1, [r0]
172		ADD r2, r0, #4
173		LDR r0, =msg_7
174		ADD r0, r0, #4
175		BL printf
176		MOV r0, #0
177		BL fflush
178		POP {pc}
179	p_print_int:
180		PUSH {lr}
181		MOV r1, r0
182		LDR r0, =msg_8
183		ADD r0, r0, #4
184		BL printf
185		MOV r0, #0
186		BL fflush
187		POP {pc}
188	p_print_ln:
189		PUSH {lr}
190		LDR r0, =msg_9
191		ADD r0, r0, #4
192		BL puts
193		MOV r0, #0
194		BL fflush
195		POP {pc}
196	p_throw_runtime_error:
197		BL p_print_string
198		MOV r0, #-1
199		BL exit
200	
===========================================================
-- Finished
