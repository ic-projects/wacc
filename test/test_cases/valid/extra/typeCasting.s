-- Compiling...
-- Printing Assembly...
typeCasting.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 6
4		.ascii	"behold"
5	msg_1:
6		.word 50
7		.ascii	"NullReferenceError: dereference a null reference\n\0"
8	msg_2:
9		.word 3
10		.ascii	"%d\0"
11	msg_3:
12		.word 1
13		.ascii	"\0"
14	msg_4:
15		.word 5
16		.ascii	"true\0"
17	msg_5:
18		.word 6
19		.ascii	"false\0"
20	msg_6:
21		.word 5
22		.ascii	"%.*s\0"
23	
24	.text
25	
26	.global main
27	main:
28		PUSH {lr}
29		SUB sp, sp, #30
30		LDR r0, =8
31		BL malloc
32		MOV r4, r0
33		LDR r5, =msg_0
34		LDR r0, =4
35		BL malloc
36		STR r5, [r0]
37		STR r0, [r4]
38		LDR r5, =0
39		LDR r0, =4
40		BL malloc
41		STR r5, [r0]
42		STR r0, [r4, #4]
43		STR r4, [sp, #26]
44		LDR r0, =8
45		BL malloc
46		MOV r4, r0
47		LDR r5, [sp, #26]
48		LDR r0, =4
49		BL malloc
50		STR r5, [r0]
51		STR r0, [r4]
52		LDR r5, =0
53		LDR r0, =4
54		BL malloc
55		STR r5, [r0]
56		STR r0, [r4, #4]
57		STR r4, [sp, #22]
58		LDR r4, [sp, #22]
59		MOV r0, r4
60		BL p_check_null_pointer
61		LDR r4, [r4]
62		LDR r4, [r4]
63		STR r4, [sp, #18]
64		LDR r4, [sp, #18]
65		MOV r0, r4
66		BL p_check_null_pointer
67		LDR r4, [r4]
68		LDR r4, [r4]
69		STR r4, [sp, #14]
70		LDR r4, [sp, #14]
71		MOV r0, r4
72		BL p_print_int
73		BL p_print_ln
74		LDR r4, [sp, #22]
75		MOV r0, r4
76		BL p_check_null_pointer
77		LDR r4, [r4]
78		LDR r4, [r4]
79		STR r4, [sp, #10]
80		LDR r4, [sp, #10]
81		MOV r0, r4
82		BL p_check_null_pointer
83		LDR r4, [r4]
84		LDRSB r4, [r4]
85		STRB r4, [sp, #9]
86		LDRSB r4, [sp, #9]
87		MOV r0, r4
88		BL putchar
89		BL p_print_ln
90		LDR r4, [sp, #22]
91		MOV r0, r4
92		BL p_check_null_pointer
93		LDR r4, [r4]
94		LDR r4, [r4]
95		STR r4, [sp, #5]
96		LDR r4, [sp, #5]
97		MOV r0, r4
98		BL p_check_null_pointer
99		LDR r4, [r4]
100		LDRSB r4, [r4]
101		STRB r4, [sp, #4]
102		LDRSB r4, [sp, #4]
103		MOV r0, r4
104		BL p_print_bool
105		BL p_print_ln
106		LDR r0, =8
107		BL malloc
108		MOV r4, r0
109		LDR r5, =4
110		LDR r0, =4
111		BL malloc
112		STR r5, [r0]
113		STR r0, [r4]
114		LDR r5, =0
115		LDR r0, =4
116		BL malloc
117		STR r5, [r0]
118		STR r0, [r4, #4]
119		STR r4, [sp]
120		LDR r0, =8
121		BL malloc
122		MOV r4, r0
123		LDR r5, [sp]
124		LDR r0, =4
125		BL malloc
126		STR r5, [r0]
127		STR r0, [r4]
128		LDR r5, =0
129		LDR r0, =4
130		BL malloc
131		STR r5, [r0]
132		STR r0, [r4, #4]
133		STR r4, [sp, #22]
134		LDR r4, [sp, #22]
135		MOV r0, r4
136		BL p_check_null_pointer
137		LDR r4, [r4]
138		LDR r4, [r4]
139		STR r4, [sp, #18]
140		LDR r4, [sp, #18]
141		MOV r0, r4
142		BL p_check_null_pointer
143		LDR r4, [r4]
144		LDR r4, [r4]
145		STR r4, [sp, #14]
146		LDR r4, [sp, #14]
147		MOV r0, r4
148		BL p_print_int
149		BL p_print_ln
150		LDR r4, [sp, #22]
151		MOV r0, r4
152		BL p_check_null_pointer
153		LDR r4, [r4]
154		LDR r4, [r4]
155		STR r4, [sp, #10]
156		LDR r4, [sp, #10]
157		MOV r0, r4
158		BL p_check_null_pointer
159		LDR r4, [r4]
160		LDRSB r4, [r4]
161		STRB r4, [sp, #9]
162		LDRSB r4, [sp, #9]
163		MOV r0, r4
164		BL putchar
165		BL p_print_ln
166		LDR r4, [sp, #22]
167		MOV r0, r4
168		BL p_check_null_pointer
169		LDR r4, [r4]
170		LDR r4, [r4]
171		STR r4, [sp, #5]
172		LDR r4, [sp, #5]
173		MOV r0, r4
174		BL p_check_null_pointer
175		LDR r4, [r4]
176		LDRSB r4, [r4]
177		STRB r4, [sp, #4]
178		LDRSB r4, [sp, #4]
179		MOV r0, r4
180		BL p_print_bool
181		BL p_print_ln
182		ADD sp, sp, #30
183		LDR r0, =0
184		POP {pc}
185		.ltorg
186	p_check_null_pointer:
187		PUSH {lr}
188		CMP r0, #0
189		LDREQ r0, =msg_1
190		BLEQ p_throw_runtime_error
191		POP {pc}
192	p_print_int:
193		PUSH {lr}
194		MOV r1, r0
195		LDR r0, =msg_2
196		ADD r0, r0, #4
197		BL printf
198		MOV r0, #0
199		BL fflush
200		POP {pc}
201	p_print_ln:
202		PUSH {lr}
203		LDR r0, =msg_3
204		ADD r0, r0, #4
205		BL puts
206		MOV r0, #0
207		BL fflush
208		POP {pc}
209	p_print_bool:
210		PUSH {lr}
211		CMP r0, #0
212		LDRNE r0, =msg_4
213		LDREQ r0, =msg_5
214		ADD r0, r0, #4
215		BL printf
216		MOV r0, #0
217		BL fflush
218		POP {pc}
219	p_throw_runtime_error:
220		BL p_print_string
221		MOV r0, #-1
222		BL exit
223	p_print_string:
224		PUSH {lr}
225		LDR r1, [r0]
226		ADD r2, r0, #4
227		LDR r0, =msg_6
228		ADD r0, r0, #4
229		BL printf
230		MOV r0, #0
231		BL fflush
232		POP {pc}
233	
===========================================================
-- Finished
