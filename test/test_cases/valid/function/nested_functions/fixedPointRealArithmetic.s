-- Compiling...
-- Printing Assembly...
fixedPointRealArithmetic.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 24
4		.ascii	"Using fixed-point real: "
5	msg_1:
6		.word 3
7		.ascii	" / "
8	msg_2:
9		.word 3
10		.ascii	" * "
11	msg_3:
12		.word 3
13		.ascii	" = "
14	msg_4:
15		.word 82
16		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
17	msg_5:
18		.word 45
19		.ascii	"DivideByZeroError: divide or modulo by zero\n\0"
20	msg_6:
21		.word 5
22		.ascii	"%.*s\0"
23	msg_7:
24		.word 3
25		.ascii	"%d\0"
26	msg_8:
27		.word 1
28		.ascii	"\0"
29	
30	.text
31	
32	.global main
33	f_q:
34		PUSH {lr}
35		LDR r4, =14
36		MOV r0, r4
37		POP {pc}
38		POP {pc}
39		.ltorg
40	f_power:
41		PUSH {lr}
42		SUB sp, sp, #4
43		LDR r4, =1
44		STR r4, [sp]
45		B L0
46	L1:
47		LDR r4, [sp]
48		LDR r5, [sp, #8]
49		SMULL r4, r5, r4, r5
50		CMP r5, r4, ASR #31
51		BLNE p_throw_overflow_error
52		STR r4, [sp]
53		LDR r4, [sp, #12]
54		LDR r5, =1
55		SUBS r4, r4, r5
56		BLVS p_throw_overflow_error
57		STR r4, [sp, #12]
58	L0:
59		LDR r4, [sp, #12]
60		LDR r5, =0
61		CMP r4, r5
62		MOVGT r4, #1
63		MOVLE r4, #0
64		CMP r4, #1
65		BEQ L1
66		LDR r4, [sp]
67		MOV r0, r4
68		ADD sp, sp, #4
69		POP {pc}
70		POP {pc}
71		.ltorg
72	f_f:
73		PUSH {lr}
74		SUB sp, sp, #8
75		BL f_q
76		MOV r4, r0
77		STR r4, [sp, #4]
78		LDR r4, [sp, #4]
79		STR r4, [sp, #-4]!
80		LDR r4, =2
81		STR r4, [sp, #-4]!
82		BL f_power
83		ADD sp, sp, #8
84		MOV r4, r0
85		STR r4, [sp]
86		LDR r4, [sp]
87		MOV r0, r4
88		ADD sp, sp, #8
89		POP {pc}
90		POP {pc}
91		.ltorg
92	f_intToFixedPoint:
93		PUSH {lr}
94		SUB sp, sp, #4
95		BL f_f
96		MOV r4, r0
97		STR r4, [sp]
98		LDR r4, [sp, #8]
99		LDR r5, [sp]
100		SMULL r4, r5, r4, r5
101		CMP r5, r4, ASR #31
102		BLNE p_throw_overflow_error
103		MOV r0, r4
104		ADD sp, sp, #4
105		POP {pc}
106		POP {pc}
107		.ltorg
108	f_fixedPointToIntRoundDown:
109		PUSH {lr}
110		SUB sp, sp, #4
111		BL f_f
112		MOV r4, r0
113		STR r4, [sp]
114		LDR r4, [sp, #8]
115		LDR r5, [sp]
116		MOV r0, r4
117		MOV r1, r5
118		BL p_check_divide_by_zero
119		BL __aeabi_idiv
120		MOV r4, r0
121		MOV r0, r4
122		ADD sp, sp, #4
123		POP {pc}
124		POP {pc}
125		.ltorg
126	f_fixedPointToIntRoundNear:
127		PUSH {lr}
128		SUB sp, sp, #4
129		BL f_f
130		MOV r4, r0
131		STR r4, [sp]
132		LDR r4, [sp, #8]
133		LDR r5, =0
134		CMP r4, r5
135		MOVGE r4, #1
136		MOVLT r4, #0
137		CMP r4, #0
138		BEQ L2
139		LDR r4, [sp, #8]
140		LDR r5, [sp]
141		LDR r6, =2
142		MOV r0, r5
143		MOV r1, r6
144		BL p_check_divide_by_zero
145		BL __aeabi_idiv
146		MOV r5, r0
147		ADDS r4, r4, r5
148		BLVS p_throw_overflow_error
149		LDR r5, [sp]
150		MOV r0, r4
151		MOV r1, r5
152		BL p_check_divide_by_zero
153		BL __aeabi_idiv
154		MOV r4, r0
155		MOV r0, r4
156		ADD sp, sp, #4
157		POP {pc}
158		B L3
159	L2:
160		LDR r4, [sp, #8]
161		LDR r5, [sp]
162		LDR r6, =2
163		MOV r0, r5
164		MOV r1, r6
165		BL p_check_divide_by_zero
166		BL __aeabi_idiv
167		MOV r5, r0
168		SUBS r4, r4, r5
169		BLVS p_throw_overflow_error
170		LDR r5, [sp]
171		MOV r0, r4
172		MOV r1, r5
173		BL p_check_divide_by_zero
174		BL __aeabi_idiv
175		MOV r4, r0
176		MOV r0, r4
177		ADD sp, sp, #4
178		POP {pc}
179	L3:
180		POP {pc}
181		.ltorg
182	f_add:
183		PUSH {lr}
184		LDR r4, [sp, #4]
185		LDR r5, [sp, #8]
186		ADDS r4, r4, r5
187		BLVS p_throw_overflow_error
188		MOV r0, r4
189		POP {pc}
190		POP {pc}
191		.ltorg
192	f_subtract:
193		PUSH {lr}
194		LDR r4, [sp, #4]
195		LDR r5, [sp, #8]
196		SUBS r4, r4, r5
197		BLVS p_throw_overflow_error
198		MOV r0, r4
199		POP {pc}
200		POP {pc}
201		.ltorg
202	f_addByInt:
203		PUSH {lr}
204		SUB sp, sp, #4
205		BL f_f
206		MOV r4, r0
207		STR r4, [sp]
208		LDR r4, [sp, #8]
209		LDR r5, [sp, #12]
210		LDR r6, [sp]
211		SMULL r5, r6, r5, r6
212		CMP r6, r5, ASR #31
213		BLNE p_throw_overflow_error
214		ADDS r4, r4, r5
215		BLVS p_throw_overflow_error
216		MOV r0, r4
217		ADD sp, sp, #4
218		POP {pc}
219		POP {pc}
220		.ltorg
221	f_subtractByInt:
222		PUSH {lr}
223		SUB sp, sp, #4
224		BL f_f
225		MOV r4, r0
226		STR r4, [sp]
227		LDR r4, [sp, #8]
228		LDR r5, [sp, #12]
229		LDR r6, [sp]
230		SMULL r5, r6, r5, r6
231		CMP r6, r5, ASR #31
232		BLNE p_throw_overflow_error
233		SUBS r4, r4, r5
234		BLVS p_throw_overflow_error
235		MOV r0, r4
236		ADD sp, sp, #4
237		POP {pc}
238		POP {pc}
239		.ltorg
240	f_multiply:
241		PUSH {lr}
242		SUB sp, sp, #4
243		BL f_f
244		MOV r4, r0
245		STR r4, [sp]
246		LDR r4, [sp, #8]
247		LDR r5, [sp, #12]
248		SMULL r4, r5, r4, r5
249		CMP r5, r4, ASR #31
250		BLNE p_throw_overflow_error
251		LDR r5, [sp]
252		MOV r0, r4
253		MOV r1, r5
254		BL p_check_divide_by_zero
255		BL __aeabi_idiv
256		MOV r4, r0
257		MOV r0, r4
258		ADD sp, sp, #4
259		POP {pc}
260		POP {pc}
261		.ltorg
262	f_multiplyByInt:
263		PUSH {lr}
264		LDR r4, [sp, #4]
265		LDR r5, [sp, #8]
266		SMULL r4, r5, r4, r5
267		CMP r5, r4, ASR #31
268		BLNE p_throw_overflow_error
269		MOV r0, r4
270		POP {pc}
271		POP {pc}
272		.ltorg
273	f_divide:
274		PUSH {lr}
275		SUB sp, sp, #4
276		BL f_f
277		MOV r4, r0
278		STR r4, [sp]
279		LDR r4, [sp, #8]
280		LDR r5, [sp]
281		SMULL r4, r5, r4, r5
282		CMP r5, r4, ASR #31
283		BLNE p_throw_overflow_error
284		LDR r5, [sp, #12]
285		MOV r0, r4
286		MOV r1, r5
287		BL p_check_divide_by_zero
288		BL __aeabi_idiv
289		MOV r4, r0
290		MOV r0, r4
291		ADD sp, sp, #4
292		POP {pc}
293		POP {pc}
294		.ltorg
295	f_divideByInt:
296		PUSH {lr}
297		LDR r4, [sp, #4]
298		LDR r5, [sp, #8]
299		MOV r0, r4
300		MOV r1, r5
301		BL p_check_divide_by_zero
302		BL __aeabi_idiv
303		MOV r4, r0
304		MOV r0, r4
305		POP {pc}
306		POP {pc}
307		.ltorg
308	main:
309		PUSH {lr}
310		SUB sp, sp, #16
311		LDR r4, =10
312		STR r4, [sp, #12]
313		LDR r4, =3
314		STR r4, [sp, #8]
315		LDR r4, =msg_0
316		MOV r0, r4
317		BL p_print_string
318		LDR r4, [sp, #12]
319		MOV r0, r4
320		BL p_print_int
321		LDR r4, =msg_1
322		MOV r0, r4
323		BL p_print_string
324		LDR r4, [sp, #8]
325		MOV r0, r4
326		BL p_print_int
327		LDR r4, =msg_2
328		MOV r0, r4
329		BL p_print_string
330		LDR r4, [sp, #8]
331		MOV r0, r4
332		BL p_print_int
333		LDR r4, =msg_3
334		MOV r0, r4
335		BL p_print_string
336		LDR r4, [sp, #12]
337		STR r4, [sp, #-4]!
338		BL f_intToFixedPoint
339		ADD sp, sp, #4
340		MOV r4, r0
341		STR r4, [sp, #4]
342		LDR r4, [sp, #8]
343		STR r4, [sp, #-4]!
344		LDR r4, [sp, #8]
345		STR r4, [sp, #-4]!
346		BL f_divideByInt
347		ADD sp, sp, #8
348		MOV r4, r0
349		STR r4, [sp, #4]
350		LDR r4, [sp, #8]
351		STR r4, [sp, #-4]!
352		LDR r4, [sp, #8]
353		STR r4, [sp, #-4]!
354		BL f_multiplyByInt
355		ADD sp, sp, #8
356		MOV r4, r0
357		STR r4, [sp, #4]
358		LDR r4, [sp, #4]
359		STR r4, [sp, #-4]!
360		BL f_fixedPointToIntRoundNear
361		ADD sp, sp, #4
362		MOV r4, r0
363		STR r4, [sp]
364		LDR r4, [sp]
365		MOV r0, r4
366		BL p_print_int
367		BL p_print_ln
368		ADD sp, sp, #16
369		LDR r0, =0
370		POP {pc}
371		.ltorg
372	p_throw_overflow_error:
373		LDR r0, =msg_4
374		BL p_throw_runtime_error
375	p_check_divide_by_zero:
376		PUSH {lr}
377		CMP r1, #0
378		LDREQ r0, =msg_5
379		BLEQ p_throw_runtime_error
380		POP {pc}
381	p_print_string:
382		PUSH {lr}
383		LDR r1, [r0]
384		ADD r2, r0, #4
385		LDR r0, =msg_6
386		ADD r0, r0, #4
387		BL printf
388		MOV r0, #0
389		BL fflush
390		POP {pc}
391	p_print_int:
392		PUSH {lr}
393		MOV r1, r0
394		LDR r0, =msg_7
395		ADD r0, r0, #4
396		BL printf
397		MOV r0, #0
398		BL fflush
399		POP {pc}
400	p_print_ln:
401		PUSH {lr}
402		LDR r0, =msg_8
403		ADD r0, r0, #4
404		BL puts
405		MOV r0, #0
406		BL fflush
407		POP {pc}
408	p_throw_runtime_error:
409		BL p_print_string
410		MOV r0, #-1
411		BL exit
412	
===========================================================
-- Finished
