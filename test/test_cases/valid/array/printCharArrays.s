-- Compiling...
-- Printing Assembly...
printCharArrays.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 3
4		.ascii	"%p\0"
5	msg_1:
6		.word 1
7		.ascii	"\0"
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 44
13		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
14	msg_4:
15		.word 45
16		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
17	msg_5:
18		.word 50
19		.ascii	"NullReferenceError: dereference a null reference\n\0"
20	msg_6:
21		.word 3
22		.ascii	"%d\0"
23	
24	.text
25	
26	.global main
27	main:
28		PUSH {lr}
29		SUB sp, sp, #20
30		LDR r0, =16
31		BL malloc
32		MOV r4, r0
33		LDR r5, =1
34		STR r5, [r4, #4]
35		LDR r5, =2
36		STR r5, [r4, #8]
37		LDR r5, =3
38		STR r5, [r4, #12]
39		LDR r5, =3
40		STR r5, [r4]
41		STR r4, [sp, #16]
42		LDR r0, =7
43		BL malloc
44		MOV r4, r0
45		MOV r5, #'a'
46		STRB r5, [r4, #4]
47		MOV r5, #'b'
48		STRB r5, [r4, #5]
49		MOV r5, #'c'
50		STRB r5, [r4, #6]
51		LDR r5, =3
52		STR r5, [r4]
53		STR r4, [sp, #12]
54		LDR r0, =8
55		BL malloc
56		MOV r4, r0
57		LDR r5, [sp, #16]
58		LDR r0, =4
59		BL malloc
60		STR r5, [r0]
61		STR r0, [r4]
62		LDR r5, [sp, #12]
63		LDR r0, =4
64		BL malloc
65		STR r5, [r0]
66		STR r0, [r4, #4]
67		STR r4, [sp, #8]
68		LDR r4, [sp, #16]
69		MOV r0, r4
70		BL p_print_reference
71		BL p_print_ln
72		LDR r4, [sp, #12]
73		MOV r0, r4
74		BL p_print_string
75		BL p_print_ln
76		ADD r4, sp, #12
77		LDR r5, =2
78		LDR r4, [r4]
79		MOV r0, r5
80		MOV r1, r4
81		BL p_check_array_bounds
82		ADD r4, r4, #4
83		ADD r4, r4, r5
84		LDRSB r4, [r4]
85		MOV r0, r4
86		BL putchar
87		BL p_print_ln
88		LDR r4, [sp, #8]
89		MOV r0, r4
90		BL p_check_null_pointer
91		LDR r4, [r4, #4]
92		LDR r4, [r4]
93		STR r4, [sp, #4]
94		LDR r4, [sp, #4]
95		MOV r0, r4
96		BL p_print_string
97		BL p_print_ln
98		ADD r4, sp, #4
99		LDR r5, =0
100		LDR r4, [r4]
101		MOV r0, r5
102		MOV r1, r4
103		BL p_check_array_bounds
104		ADD r4, r4, #4
105		ADD r4, r4, r5
106		LDRSB r4, [r4]
107		MOV r0, r4
108		BL putchar
109		BL p_print_ln
110		ADD r4, sp, #4
111		LDR r5, =1
112		LDR r4, [r4]
113		MOV r0, r5
114		MOV r1, r4
115		BL p_check_array_bounds
116		ADD r4, r4, #4
117		ADD r4, r4, r5
118		LDRSB r4, [r4]
119		MOV r0, r4
120		BL putchar
121		BL p_print_ln
122		ADD r4, sp, #4
123		LDR r5, =2
124		LDR r4, [r4]
125		MOV r0, r5
126		MOV r1, r4
127		BL p_check_array_bounds
128		ADD r4, r4, #4
129		ADD r4, r4, r5
130		LDRSB r4, [r4]
131		MOV r0, r4
132		BL putchar
133		BL p_print_ln
134		LDR r4, [sp, #8]
135		MOV r0, r4
136		BL p_check_null_pointer
137		LDR r4, [r4]
138		LDR r4, [r4]
139		STR r4, [sp]
140		ADD r4, sp, #0
141		LDR r5, =0
142		LDR r4, [r4]
143		MOV r0, r5
144		MOV r1, r4
145		BL p_check_array_bounds
146		ADD r4, r4, #4
147		ADD r4, r4, r5, LSL #2
148		LDR r4, [r4]
149		MOV r0, r4
150		BL p_print_int
151		BL p_print_ln
152		ADD r4, sp, #16
153		LDR r5, =0
154		LDR r4, [r4]
155		MOV r0, r5
156		MOV r1, r4
157		BL p_check_array_bounds
158		ADD r4, r4, #4
159		ADD r4, r4, r5, LSL #2
160		LDR r4, [r4]
161		MOV r0, r4
162		BL p_print_int
163		BL p_print_ln
164		ADD r4, sp, #0
165		LDR r5, =1
166		LDR r4, [r4]
167		MOV r0, r5
168		MOV r1, r4
169		BL p_check_array_bounds
170		ADD r4, r4, #4
171		ADD r4, r4, r5, LSL #2
172		LDR r4, [r4]
173		MOV r0, r4
174		BL p_print_int
175		BL p_print_ln
176		ADD r4, sp, #16
177		LDR r5, =1
178		LDR r4, [r4]
179		MOV r0, r5
180		MOV r1, r4
181		BL p_check_array_bounds
182		ADD r4, r4, #4
183		ADD r4, r4, r5, LSL #2
184		LDR r4, [r4]
185		MOV r0, r4
186		BL p_print_int
187		BL p_print_ln
188		ADD r4, sp, #0
189		LDR r5, =2
190		LDR r4, [r4]
191		MOV r0, r5
192		MOV r1, r4
193		BL p_check_array_bounds
194		ADD r4, r4, #4
195		ADD r4, r4, r5, LSL #2
196		LDR r4, [r4]
197		MOV r0, r4
198		BL p_print_int
199		BL p_print_ln
200		ADD r4, sp, #16
201		LDR r5, =2
202		LDR r4, [r4]
203		MOV r0, r5
204		MOV r1, r4
205		BL p_check_array_bounds
206		ADD r4, r4, #4
207		ADD r4, r4, r5, LSL #2
208		LDR r4, [r4]
209		MOV r0, r4
210		BL p_print_int
211		BL p_print_ln
212		ADD sp, sp, #20
213		LDR r0, =0
214		POP {pc}
215		.ltorg
216	p_print_reference:
217		PUSH {lr}
218		MOV r1, r0
219		LDR r0, =msg_0
220		ADD r0, r0, #4
221		BL printf
222		MOV r0, #0
223		BL fflush
224		POP {pc}
225	p_print_ln:
226		PUSH {lr}
227		LDR r0, =msg_1
228		ADD r0, r0, #4
229		BL puts
230		MOV r0, #0
231		BL fflush
232		POP {pc}
233	p_print_string:
234		PUSH {lr}
235		LDR r1, [r0]
236		ADD r2, r0, #4
237		LDR r0, =msg_2
238		ADD r0, r0, #4
239		BL printf
240		MOV r0, #0
241		BL fflush
242		POP {pc}
243	p_check_array_bounds:
244		PUSH {lr}
245		CMP r0, #0
246		LDRLT r0, =msg_3
247		BLLT p_throw_runtime_error
248		LDR r1, [r1]
249		CMP r0, r1
250		LDRCS r0, =msg_4
251		BLCS p_throw_runtime_error
252		POP {pc}
253	p_check_null_pointer:
254		PUSH {lr}
255		CMP r0, #0
256		LDREQ r0, =msg_5
257		BLEQ p_throw_runtime_error
258		POP {pc}
259	p_print_int:
260		PUSH {lr}
261		MOV r1, r0
262		LDR r0, =msg_6
263		ADD r0, r0, #4
264		BL printf
265		MOV r0, #0
266		BL fflush
267		POP {pc}
268	p_throw_runtime_error:
269		BL p_print_string
270		MOV r0, #-1
271		BL exit
272	
===========================================================
-- Finished
