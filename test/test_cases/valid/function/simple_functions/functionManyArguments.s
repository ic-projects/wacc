-- Compiling...
-- Printing Assembly...
functionManyArguments.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 5
4		.ascii	"a is "
5	msg_1:
6		.word 5
7		.ascii	"b is "
8	msg_2:
9		.word 5
10		.ascii	"c is "
11	msg_3:
12		.word 5
13		.ascii	"d is "
14	msg_4:
15		.word 5
16		.ascii	"e is "
17	msg_5:
18		.word 5
19		.ascii	"f is "
20	msg_6:
21		.word 5
22		.ascii	"hello"
23	msg_7:
24		.word 10
25		.ascii	"answer is "
26	msg_8:
27		.word 5
28		.ascii	"%.*s\0"
29	msg_9:
30		.word 3
31		.ascii	"%d\0"
32	msg_10:
33		.word 1
34		.ascii	"\0"
35	msg_11:
36		.word 5
37		.ascii	"true\0"
38	msg_12:
39		.word 6
40		.ascii	"false\0"
41	msg_13:
42		.word 3
43		.ascii	"%p\0"
44	
45	.text
46	
47	.global main
48	f_doSomething:
49		PUSH {lr}
50		LDR r4, =msg_0
51		MOV r0, r4
52		BL p_print_string
53		LDR r4, [sp, #4]
54		MOV r0, r4
55		BL p_print_int
56		BL p_print_ln
57		LDR r4, =msg_1
58		MOV r0, r4
59		BL p_print_string
60		LDRSB r4, [sp, #8]
61		MOV r0, r4
62		BL p_print_bool
63		BL p_print_ln
64		LDR r4, =msg_2
65		MOV r0, r4
66		BL p_print_string
67		LDRSB r4, [sp, #9]
68		MOV r0, r4
69		BL putchar
70		BL p_print_ln
71		LDR r4, =msg_3
72		MOV r0, r4
73		BL p_print_string
74		LDR r4, [sp, #10]
75		MOV r0, r4
76		BL p_print_string
77		BL p_print_ln
78		LDR r4, =msg_4
79		MOV r0, r4
80		BL p_print_string
81		LDR r4, [sp, #14]
82		MOV r0, r4
83		BL p_print_reference
84		BL p_print_ln
85		LDR r4, =msg_5
86		MOV r0, r4
87		BL p_print_string
88		LDR r4, [sp, #18]
89		MOV r0, r4
90		BL p_print_reference
91		BL p_print_ln
92		MOV r4, #'g'
93		MOV r0, r4
94		POP {pc}
95		POP {pc}
96		.ltorg
97	main:
98		PUSH {lr}
99		SUB sp, sp, #9
100		LDR r0, =6
101		BL malloc
102		MOV r4, r0
103		MOV r5, #0
104		STRB r5, [r4, #4]
105		MOV r5, #1
106		STRB r5, [r4, #5]
107		LDR r5, =2
108		STR r5, [r4]
109		STR r4, [sp, #5]
110		LDR r0, =12
111		BL malloc
112		MOV r4, r0
113		LDR r5, =1
114		STR r5, [r4, #4]
115		LDR r5, =2
116		STR r5, [r4, #8]
117		LDR r5, =2
118		STR r5, [r4]
119		STR r4, [sp, #1]
120		LDR r4, [sp, #1]
121		STR r4, [sp, #-4]!
122		LDR r4, [sp, #9]
123		STR r4, [sp, #-4]!
124		LDR r4, =msg_6
125		STR r4, [sp, #-4]!
126		MOV r4, #'u'
127		STRB r4, [sp, #-1]!
128		MOV r4, #1
129		STRB r4, [sp, #-1]!
130		LDR r4, =42
131		STR r4, [sp, #-4]!
132		BL f_doSomething
133		ADD sp, sp, #18
134		MOV r4, r0
135		STRB r4, [sp]
136		LDR r4, =msg_7
137		MOV r0, r4
138		BL p_print_string
139		LDRSB r4, [sp]
140		MOV r0, r4
141		BL putchar
142		BL p_print_ln
143		ADD sp, sp, #9
144		LDR r0, =0
145		POP {pc}
146		.ltorg
147	p_print_string:
148		PUSH {lr}
149		LDR r1, [r0]
150		ADD r2, r0, #4
151		LDR r0, =msg_8
152		ADD r0, r0, #4
153		BL printf
154		MOV r0, #0
155		BL fflush
156		POP {pc}
157	p_print_int:
158		PUSH {lr}
159		MOV r1, r0
160		LDR r0, =msg_9
161		ADD r0, r0, #4
162		BL printf
163		MOV r0, #0
164		BL fflush
165		POP {pc}
166	p_print_ln:
167		PUSH {lr}
168		LDR r0, =msg_10
169		ADD r0, r0, #4
170		BL puts
171		MOV r0, #0
172		BL fflush
173		POP {pc}
174	p_print_bool:
175		PUSH {lr}
176		CMP r0, #0
177		LDRNE r0, =msg_11
178		LDREQ r0, =msg_12
179		ADD r0, r0, #4
180		BL printf
181		MOV r0, #0
182		BL fflush
183		POP {pc}
184	p_print_reference:
185		PUSH {lr}
186		MOV r1, r0
187		LDR r0, =msg_13
188		ADD r0, r0, #4
189		BL printf
190		MOV r0, #0
191		BL fflush
192		POP {pc}
193	
===========================================================
-- Finished
