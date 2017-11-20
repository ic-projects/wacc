-- Compiling...
-- Printing Assembly...
precedenceExpr1.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 82
4		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
5	msg_1:
6		.word 3
7		.ascii	"%d\0"
8	msg_2:
9		.word 1
10		.ascii	"\0"
11	msg_3:
12		.word 45
13		.ascii	"DivideByZeroError: divide or modulo by zero\n\0"
14	msg_4:
15		.word 5
16		.ascii	"%.*s\0"
17	
18	.text
19	
20	.global main
21	main:
22		PUSH {lr}
23		LDR r4, =1
24		LDR r5, =2
25		ADDS r4, r4, r5
26		BLVS p_throw_overflow_error
27		LDR r5, =3
28		ADDS r4, r4, r5
29		BLVS p_throw_overflow_error
30		LDR r5, =4
31		ADDS r4, r4, r5
32		BLVS p_throw_overflow_error
33		LDR r5, =5
34		ADDS r4, r4, r5
35		BLVS p_throw_overflow_error
36		MOV r0, r4
37		BL p_print_int
38		BL p_print_ln
39		LDR r4, =1
40		LDR r5, =2
41		SMULL r4, r5, r4, r5
42		CMP r5, r4, ASR #31
43		BLNE p_throw_overflow_error
44		LDR r5, =3
45		SMULL r4, r5, r4, r5
46		CMP r5, r4, ASR #31
47		BLNE p_throw_overflow_error
48		LDR r5, =4
49		SMULL r4, r5, r4, r5
50		CMP r5, r4, ASR #31
51		BLNE p_throw_overflow_error
52		LDR r5, =5
53		SMULL r4, r5, r4, r5
54		CMP r5, r4, ASR #31
55		BLNE p_throw_overflow_error
56		MOV r0, r4
57		BL p_print_int
58		BL p_print_ln
59		LDR r4, =1
60		LDR r5, =2
61		ADDS r4, r4, r5
62		BLVS p_throw_overflow_error
63		LDR r5, =3
64		LDR r6, =4
65		SMULL r5, r6, r5, r6
66		CMP r6, r5, ASR #31
67		BLNE p_throw_overflow_error
68		ADDS r4, r4, r5
69		BLVS p_throw_overflow_error
70		LDR r5, =5
71		ADDS r4, r4, r5
72		BLVS p_throw_overflow_error
73		MOV r0, r4
74		BL p_print_int
75		BL p_print_ln
76		LDR r4, =1
77		LDR r5, =2
78		ADDS r4, r4, r5
79		BLVS p_throw_overflow_error
80		LDR r5, =3
81		LDR r6, =4
82		SMULL r5, r6, r5, r6
83		CMP r6, r5, ASR #31
84		BLNE p_throw_overflow_error
85		LDR r6, =5
86		MOV r0, r5
87		MOV r1, r6
88		BL p_check_divide_by_zero
89		BL __aeabi_idiv
90		MOV r5, r0
91		SUBS r4, r4, r5
92		BLVS p_throw_overflow_error
93		LDR r5, =6
94		ADDS r4, r4, r5
95		BLVS p_throw_overflow_error
96		LDR r5, =7
97		SUBS r4, r4, r5
98		BLVS p_throw_overflow_error
99		MOV r0, r4
100		BL p_print_int
101		BL p_print_ln
102		LDR r4, =1
103		LDR r5, =2
104		ADDS r4, r4, r5
105		BLVS p_throw_overflow_error
106		LDR r5, =3
107		LDR r6, =4
108		SMULL r5, r6, r5, r6
109		CMP r6, r5, ASR #31
110		BLNE p_throw_overflow_error
111		LDR r6, =5
112		MOV r0, r5
113		MOV r1, r6
114		BL p_check_divide_by_zero
115		BL __aeabi_idiv
116		MOV r5, r0
117		SUBS r4, r4, r5
118		BLVS p_throw_overflow_error
119		LDR r5, =6
120		ADDS r4, r4, r5
121		BLVS p_throw_overflow_error
122		LDR r5, =7
123		SUBS r4, r4, r5
124		BLVS p_throw_overflow_error
125		MOV r0, r4
126		BL p_print_int
127		BL p_print_ln
128		LDR r4, =1
129		LDR r5, =2
130		SMULL r4, r5, r4, r5
131		CMP r5, r4, ASR #31
132		BLNE p_throw_overflow_error
133		LDR r5, =3
134		LDR r6, =4
135		SMULL r5, r6, r5, r6
136		CMP r6, r5, ASR #31
137		BLNE p_throw_overflow_error
138		ADDS r4, r4, r5
139		BLVS p_throw_overflow_error
140		MOV r0, r4
141		BL p_print_int
142		BL p_print_ln
143		LDR r4, =1
144		LDR r5, =2
145		MOV r0, r4
146		MOV r1, r5
147		BL p_check_divide_by_zero
148		BL __aeabi_idiv
149		MOV r4, r0
150		LDR r5, =3
151		LDR r6, =4
152		SMULL r5, r6, r5, r6
153		CMP r6, r5, ASR #31
154		BLNE p_throw_overflow_error
155		LDR r6, =5
156		SMULL r5, r6, r5, r6
157		CMP r6, r5, ASR #31
158		BLNE p_throw_overflow_error
159		ADDS r4, r4, r5
160		BLVS p_throw_overflow_error
161		MOV r0, r4
162		BL p_print_int
163		BL p_print_ln
164		LDR r4, =1
165		LDR r5, =2
166		LDR r6, =3
167		ADDS r5, r5, r6
168		BLVS p_throw_overflow_error
169		LDR r6, =4
170		SMULL r5, r6, r5, r6
171		CMP r6, r5, ASR #31
172		BLNE p_throw_overflow_error
173		ADDS r4, r4, r5
174		BLVS p_throw_overflow_error
175		LDR r5, =5
176		ADDS r4, r4, r5
177		BLVS p_throw_overflow_error
178		MOV r0, r4
179		BL p_print_int
180		BL p_print_ln
181		LDR r4, =1
182		LDR r5, =2
183		ADDS r4, r4, r5
184		BLVS p_throw_overflow_error
185		LDR r5, =3
186		ADDS r4, r4, r5
187		BLVS p_throw_overflow_error
188		LDR r5, =4
189		SMULL r4, r5, r4, r5
190		CMP r5, r4, ASR #31
191		BLNE p_throw_overflow_error
192		LDR r5, =5
193		ADDS r4, r4, r5
194		BLVS p_throw_overflow_error
195		MOV r0, r4
196		BL p_print_int
197		BL p_print_ln
198		LDR r4, =1
199		LDR r5, =2
200		LDR r6, =3
201		LDR r7, =4
202		SMULL r6, r7, r6, r7
203		CMP r7, r6, ASR #31
204		BLNE p_throw_overflow_error
205		ADDS r5, r5, r6
206		BLVS p_throw_overflow_error
207		LDR r6, =5
208		ADDS r5, r5, r6
209		BLVS p_throw_overflow_error
210		MOV r0, r4
211		MOV r1, r5
212		BL p_check_divide_by_zero
213		BL __aeabi_idiv
214		MOV r4, r0
215		MOV r0, r4
216		BL p_print_int
217		BL p_print_ln
218		LDR r0, =0
219		POP {pc}
220		.ltorg
221	p_throw_overflow_error:
222		LDR r0, =msg_0
223		BL p_throw_runtime_error
224	p_print_int:
225		PUSH {lr}
226		MOV r1, r0
227		LDR r0, =msg_1
228		ADD r0, r0, #4
229		BL printf
230		MOV r0, #0
231		BL fflush
232		POP {pc}
233	p_print_ln:
234		PUSH {lr}
235		LDR r0, =msg_2
236		ADD r0, r0, #4
237		BL puts
238		MOV r0, #0
239		BL fflush
240		POP {pc}
241	p_check_divide_by_zero:
242		PUSH {lr}
243		CMP r1, #0
244		LDREQ r0, =msg_3
245		BLEQ p_throw_runtime_error
246		POP {pc}
247	p_throw_runtime_error:
248		BL p_print_string
249		MOV r0, #-1
250		BL exit
251	p_print_string:
252		PUSH {lr}
253		LDR r1, [r0]
254		ADD r2, r0, #4
255		LDR r0, =msg_4
256		ADD r0, r0, #4
257		BL printf
258		MOV r0, #0
259		BL fflush
260		POP {pc}
261	
===========================================================
-- Finished
