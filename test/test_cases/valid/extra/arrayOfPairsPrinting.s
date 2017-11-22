-- Compiling...
-- Printing Assembly...
arrayOfPairsPrinting.s contents are:
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
9		.word 44
10		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
11	msg_3:
12		.word 45
13		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
14	msg_4:
15		.word 50
16		.ascii	"NullReferenceError: dereference a null reference\n\0"
17	msg_5:
18		.word 3
19		.ascii	"%d\0"
20	msg_6:
21		.word 5
22		.ascii	"%.*s\0"
23	
24	.text
25	
26	.global main
27	main:
28		PUSH {lr}
29		SUB sp, sp, #76
30		LDR r4, =0
31		STR r4, [sp, #72]
32		LDR r4, [sp, #72]
33		STR r4, [sp, #68]
34		LDR r4, [sp, #72]
35		STR r4, [sp, #64]
36		LDR r0, =8
37		BL malloc
38		MOV r4, r0
39		LDR r5, [sp, #72]
40		LDR r0, =4
41		BL malloc
42		STR r5, [r0]
43		STR r0, [r4]
44		LDR r5, [sp, #68]
45		LDR r0, =4
46		BL malloc
47		STR r5, [r0]
48		STR r0, [r4, #4]
49		STR r4, [sp, #60]
50		LDR r0, =8
51		BL malloc
52		MOV r4, r0
53		LDR r5, [sp, #60]
54		LDR r0, =4
55		BL malloc
56		STR r5, [r0]
57		STR r0, [r4]
58		LDR r5, =0
59		LDR r0, =4
60		BL malloc
61		STR r5, [r0]
62		STR r0, [r4, #4]
63		STR r4, [sp, #56]
64		LDR r0, =8
65		BL malloc
66		MOV r4, r0
67		LDR r5, =0
68		LDR r0, =4
69		BL malloc
70		STR r5, [r0]
71		STR r0, [r4]
72		LDR r5, =0
73		LDR r0, =4
74		BL malloc
75		STR r5, [r0]
76		STR r0, [r4, #4]
77		STR r4, [sp, #52]
78		LDR r4, [sp, #52]
79		STR r4, [sp, #48]
80		LDR r0, =8
81		BL malloc
82		MOV r4, r0
83		LDR r5, =0
84		LDR r0, =4
85		BL malloc
86		STR r5, [r0]
87		STR r0, [r4]
88		LDR r5, =0
89		LDR r0, =4
90		BL malloc
91		STR r5, [r0]
92		STR r0, [r4, #4]
93		STR r4, [sp, #44]
94		LDR r4, [sp, #44]
95		MOV r0, r4
96		BL p_print_reference
97		BL p_print_ln
98		LDR r0, =16
99		BL malloc
100		MOV r4, r0
101		LDR r5, [sp, #72]
102		STR r5, [r4, #4]
103		LDR r5, [sp, #68]
104		STR r5, [r4, #8]
105		LDR r5, [sp, #64]
106		STR r5, [r4, #12]
107		LDR r5, =3
108		STR r5, [r4]
109		STR r4, [sp, #40]
110		LDR r4, [sp, #40]
111		STR r4, [sp, #36]
112		LDR r0, =16
113		BL malloc
114		MOV r4, r0
115		LDR r5, =0
116		STR r5, [r4, #4]
117		LDR r5, =0
118		STR r5, [r4, #8]
119		LDR r5, [sp, #60]
120		STR r5, [r4, #12]
121		LDR r5, =3
122		STR r5, [r4]
123		STR r4, [sp, #32]
124		LDR r0, =16
125		BL malloc
126		MOV r4, r0
127		LDR r5, [sp, #36]
128		STR r5, [r4, #4]
129		LDR r5, [sp, #32]
130		STR r5, [r4, #8]
131		LDR r5, [sp, #40]
132		STR r5, [r4, #12]
133		LDR r5, =3
134		STR r5, [r4]
135		STR r4, [sp, #28]
136		LDR r4, [sp, #28]
137		MOV r0, r4
138		BL p_print_reference
139		BL p_print_ln
140		ADD r4, sp, #28
141		LDR r5, =0
142		LDR r4, [r4]
143		MOV r0, r5
144		MOV r1, r4
145		BL p_check_array_bounds
146		ADD r4, r4, #4
147		ADD r4, r4, r5, LSL #2
148		LDR r4, [r4]
149		MOV r0, r4
150		BL p_print_reference
151		BL p_print_ln
152		ADD r4, sp, #28
153		LDR r5, =1
154		LDR r4, [r4]
155		MOV r0, r5
156		MOV r1, r4
157		BL p_check_array_bounds
158		ADD r4, r4, #4
159		ADD r4, r4, r5, LSL #2
160		LDR r5, =2
161		LDR r4, [r4]
162		MOV r0, r5
163		MOV r1, r4
164		BL p_check_array_bounds
165		ADD r4, r4, #4
166		ADD r4, r4, r5, LSL #2
167		LDR r4, [r4]
168		MOV r0, r4
169		BL p_print_reference
170		BL p_print_ln
171		LDR r0, =8
172		BL malloc
173		MOV r4, r0
174		LDR r5, [sp, #32]
175		LDR r5, [r5]
176		LDR r0, =4
177		BL malloc
178		STR r5, [r0]
179		STR r0, [r4]
180		LDR r5, =1
181		LDR r0, =4
182		BL malloc
183		STR r5, [r0]
184		STR r0, [r4, #4]
185		STR r4, [sp, #24]
186		LDR r4, [sp, #24]
187		MOV r0, r4
188		BL p_check_null_pointer
189		LDR r4, [r4]
190		LDR r4, [r4]
191		STR r4, [sp, #20]
192		LDR r4, [sp, #20]
193		MOV r0, r4
194		BL p_print_int
195		BL p_print_ln
196		LDR r0, =12
197		BL malloc
198		MOV r4, r0
199		ADD r5, sp, #28
200		LDR r6, =0
201		LDR r5, [r5]
202		MOV r0, r6
203		MOV r1, r5
204		BL p_check_array_bounds
205		ADD r5, r5, #4
206		ADD r5, r5, r6, LSL #2
207		ADD r6, sp, #28
208		LDR r7, [sp, #32]
209		LDR r7, [r7]
210		LDR r6, [r6]
211		MOV r0, r7
212		MOV r1, r6
213		BL p_check_array_bounds
214		ADD r6, r6, #4
215		ADD r6, r6, r7, LSL #2
216		LDR r6, [r6]
217		LDR r6, [r6]
218		LDR r5, [r5]
219		MOV r0, r6
220		MOV r1, r5
221		BL p_check_array_bounds
222		ADD r5, r5, #4
223		ADD r5, r5, r6, LSL #2
224		LDR r5, [r5]
225		STR r5, [r4, #4]
226		ADD r5, sp, #28
227		LDR r6, =1
228		LDR r5, [r5]
229		MOV r0, r6
230		MOV r1, r5
231		BL p_check_array_bounds
232		ADD r5, r5, #4
233		ADD r5, r5, r6, LSL #2
234		LDR r6, =2
235		LDR r5, [r5]
236		MOV r0, r6
237		MOV r1, r5
238		BL p_check_array_bounds
239		ADD r5, r5, #4
240		ADD r5, r5, r6, LSL #2
241		LDR r5, [r5]
242		STR r5, [r4, #8]
243		LDR r5, =2
244		STR r5, [r4]
245		STR r4, [sp, #16]
246		LDR r0, =12
247		BL malloc
248		MOV r4, r0
249		ADD r5, sp, #28
250		LDR r6, [sp, #32]
251		LDR r6, [r6]
252		LDR r5, [r5]
253		MOV r0, r6
254		MOV r1, r5
255		BL p_check_array_bounds
256		ADD r5, r5, #4
257		ADD r5, r5, r6, LSL #2
258		LDR r5, [r5]
259		STR r5, [r4, #4]
260		LDR r5, [sp, #36]
261		STR r5, [r4, #8]
262		LDR r5, =2
263		STR r5, [r4]
264		STR r4, [sp, #12]
265		LDR r0, =8
266		BL malloc
267		MOV r4, r0
268		ADD r5, sp, #12
269		LDR r6, =1
270		LDR r5, [r5]
271		MOV r0, r6
272		MOV r1, r5
273		BL p_check_array_bounds
274		ADD r5, r5, #4
275		ADD r5, r5, r6, LSL #2
276		LDR r6, =2
277		LDR r5, [r5]
278		MOV r0, r6
279		MOV r1, r5
280		BL p_check_array_bounds
281		ADD r5, r5, #4
282		ADD r5, r5, r6, LSL #2
283		LDR r5, [r5]
284		LDR r0, =4
285		BL malloc
286		STR r5, [r0]
287		STR r0, [r4]
288		LDR r5, =0
289		LDR r0, =4
290		BL malloc
291		STR r5, [r0]
292		STR r0, [r4, #4]
293		STR r4, [sp, #8]
294		LDR r0, =8
295		BL malloc
296		MOV r4, r0
297		LDR r5, =0
298		LDR r0, =4
299		BL malloc
300		STR r5, [r0]
301		STR r0, [r4]
302		ADD r5, sp, #16
303		ADD r6, sp, #28
304		ADD r7, sp, #28
305		ADD r8, sp, #28
306		LDR r9, [sp, #28]
307		LDR r9, [r9]
308		LDR r8, [r8]
309		MOV r0, r9
310		MOV r1, r8
311		BL p_check_array_bounds
312		ADD r8, r8, #4
313		ADD r8, r8, r9, LSL #2
314		LDR r8, [r8]
315		LDR r8, [r8]
316		LDR r7, [r7]
317		MOV r0, r8
318		MOV r1, r7
319		BL p_check_array_bounds
320		ADD r7, r7, #4
321		ADD r7, r7, r8, LSL #2
322		LDR r7, [r7]
323		LDR r7, [r7]
324		LDR r6, [r6]
325		MOV r0, r7
326		MOV r1, r6
327		BL p_check_array_bounds
328		ADD r6, r6, #4
329		ADD r6, r6, r7, LSL #2
330		LDR r6, [r6]
331		LDR r6, [r6]
332		LDR r5, [r5]
333		MOV r0, r6
334		MOV r1, r5
335		BL p_check_array_bounds
336		ADD r5, r5, #4
337		ADD r5, r5, r6, LSL #2
338		LDR r5, [r5]
339		LDR r0, =4
340		BL malloc
341		STR r5, [r0]
342		STR r0, [r4, #4]
343		STR r4, [sp, #4]
344		LDR r0, =8
345		BL malloc
346		MOV r4, r0
347		LDR r5, [sp, #20]
348		LDR r0, =4
349		BL malloc
350		STR r5, [r0]
351		STR r0, [r4]
352		ADD r5, sp, #12
353		LDR r6, =0
354		LDR r5, [r5]
355		MOV r0, r6
356		MOV r1, r5
357		BL p_check_array_bounds
358		ADD r5, r5, #4
359		ADD r5, r5, r6, LSL #2
360		LDR r5, [r5]
361		LDR r0, =4
362		BL malloc
363		STR r5, [r0]
364		STR r0, [r4, #4]
365		STR r4, [sp]
366		ADD sp, sp, #76
367		LDR r0, =0
368		POP {pc}
369		.ltorg
370	p_print_reference:
371		PUSH {lr}
372		MOV r1, r0
373		LDR r0, =msg_0
374		ADD r0, r0, #4
375		BL printf
376		MOV r0, #0
377		BL fflush
378		POP {pc}
379	p_print_ln:
380		PUSH {lr}
381		LDR r0, =msg_1
382		ADD r0, r0, #4
383		BL puts
384		MOV r0, #0
385		BL fflush
386		POP {pc}
387	p_check_array_bounds:
388		PUSH {lr}
389		CMP r0, #0
390		LDRLT r0, =msg_2
391		BLLT p_throw_runtime_error
392		LDR r1, [r1]
393		CMP r0, r1
394		LDRCS r0, =msg_3
395		BLCS p_throw_runtime_error
396		POP {pc}
397	p_check_null_pointer:
398		PUSH {lr}
399		CMP r0, #0
400		LDREQ r0, =msg_4
401		BLEQ p_throw_runtime_error
402		POP {pc}
403	p_print_int:
404		PUSH {lr}
405		MOV r1, r0
406		LDR r0, =msg_5
407		ADD r0, r0, #4
408		BL printf
409		MOV r0, #0
410		BL fflush
411		POP {pc}
412	p_throw_runtime_error:
413		BL p_print_string
414		MOV r0, #-1
415		BL exit
416	p_print_string:
417		PUSH {lr}
418		LDR r1, [r0]
419		ADD r2, r0, #4
420		LDR r0, =msg_6
421		ADD r0, r0, #4
422		BL printf
423		MOV r0, #0
424		BL fflush
425		POP {pc}
426	
===========================================================
-- Finished
