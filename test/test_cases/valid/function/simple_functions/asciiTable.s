-- Compiling...
-- Printing Assembly...
asciiTable.s contents are:
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
9		.word 3
10		.ascii	"|  "
11	msg_3:
12		.word 1
13		.ascii	" "
14	msg_4:
15		.word 3
16		.ascii	" = "
17	msg_5:
18		.word 3
19		.ascii	"  |"
20	msg_6:
21		.word 28
22		.ascii	"Asci character lookup table:"
23	msg_7:
24		.word 5
25		.ascii	"%.*s\0"
26	msg_8:
27		.word 82
28		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
29	msg_9:
30		.word 1
31		.ascii	"\0"
32	msg_10:
33		.word 3
34		.ascii	"%d\0"
35	
36	.text
37	
38	.global main
39	f_printLine:
40		PUSH {lr}
41		SUB sp, sp, #4
42		LDR r4, =0
43		STR r4, [sp]
44		B L0
45	L1:
46		LDR r4, =msg_0
47		MOV r0, r4
48		BL p_print_string
49		LDR r4, [sp]
50		LDR r5, =1
51		ADDS r4, r4, r5
52		BLVS p_throw_overflow_error
53		STR r4, [sp]
54	L0:
55		LDR r4, [sp]
56		LDR r5, [sp, #8]
57		CMP r4, r5
58		MOVLT r4, #1
59		MOVGE r4, #0
60		CMP r4, #1
61		BEQ L1
62		LDR r4, =msg_1
63		MOV r0, r4
64		BL p_print_string
65		BL p_print_ln
66		MOV r4, #1
67		MOV r0, r4
68		ADD sp, sp, #4
69		POP {pc}
70		POP {pc}
71		.ltorg
72	f_printMap:
73		PUSH {lr}
74		LDR r4, =msg_2
75		MOV r0, r4
76		BL p_print_string
77		LDR r4, [sp, #4]
78		LDR r5, =100
79		CMP r4, r5
80		MOVLT r4, #1
81		MOVGE r4, #0
82		CMP r4, #0
83		BEQ L2
84		LDR r4, =msg_3
85		MOV r0, r4
86		BL p_print_string
87		B L3
88	L2:
89	L3:
90		LDR r4, [sp, #4]
91		MOV r0, r4
92		BL p_print_int
93		LDR r4, =msg_4
94		MOV r0, r4
95		BL p_print_string
96		LDR r4, [sp, #4]
97		MOV r0, r4
98		BL putchar
99		LDR r4, =msg_5
100		MOV r0, r4
101		BL p_print_string
102		BL p_print_ln
103		MOV r4, #1
104		MOV r0, r4
105		POP {pc}
106		POP {pc}
107		.ltorg
108	main:
109		PUSH {lr}
110		SUB sp, sp, #5
111		LDR r4, =msg_6
112		MOV r0, r4
113		BL p_print_string
114		BL p_print_ln
115		LDR r4, =13
116		STR r4, [sp, #-4]!
117		BL f_printLine
118		ADD sp, sp, #4
119		MOV r4, r0
120		STRB r4, [sp, #4]
121		MOV r4, #' '
122		STR r4, [sp]
123		B L4
124	L5:
125		LDR r4, [sp]
126		STR r4, [sp, #-4]!
127		BL f_printMap
128		ADD sp, sp, #4
129		MOV r4, r0
130		STRB r4, [sp, #4]
131		LDR r4, [sp]
132		LDR r5, =1
133		ADDS r4, r4, r5
134		BLVS p_throw_overflow_error
135		STR r4, [sp]
136	L4:
137		LDR r4, [sp]
138		LDR r5, =127
139		CMP r4, r5
140		MOVLT r4, #1
141		MOVGE r4, #0
142		CMP r4, #1
143		BEQ L5
144		LDR r4, =13
145		STR r4, [sp, #-4]!
146		BL f_printLine
147		ADD sp, sp, #4
148		MOV r4, r0
149		STRB r4, [sp, #4]
150		ADD sp, sp, #5
151		LDR r0, =0
152		POP {pc}
153		.ltorg
154	p_print_string:
155		PUSH {lr}
156		LDR r1, [r0]
157		ADD r2, r0, #4
158		LDR r0, =msg_7
159		ADD r0, r0, #4
160		BL printf
161		MOV r0, #0
162		BL fflush
163		POP {pc}
164	p_throw_overflow_error:
165		LDR r0, =msg_8
166		BL p_throw_runtime_error
167	p_print_ln:
168		PUSH {lr}
169		LDR r0, =msg_9
170		ADD r0, r0, #4
171		BL puts
172		MOV r0, #0
173		BL fflush
174		POP {pc}
175	p_print_int:
176		PUSH {lr}
177		MOV r1, r0
178		LDR r0, =msg_10
179		ADD r0, r0, #4
180		BL printf
181		MOV r0, #0
182		BL fflush
183		POP {pc}
184	p_throw_runtime_error:
185		BL p_print_string
186		MOV r0, #-1
187		BL exit
188	
===========================================================
-- Finished
