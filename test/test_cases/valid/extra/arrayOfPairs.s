-- Compiling...
-- Printing Assembly...
arrayOfPairs.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 50
4		.ascii	"NullReferenceError: dereference a null reference\n\0"
5	msg_1:
6		.word 44
7		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
8	msg_2:
9		.word 45
10		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
11	msg_3:
12		.word 5
13		.ascii	"%.*s\0"
14	
15	.text
16	
17	.global main
18	f_f:
19		PUSH {lr}
20		LDR r4, [sp, #4]
21		MOV r0, r4
22		POP {pc}
23		POP {pc}
24		.ltorg
25	main:
26		PUSH {lr}
27		SUB sp, sp, #76
28		LDR r4, =0
29		STR r4, [sp, #72]
30		LDR r4, [sp, #72]
31		MOV r0, r4
32		BL p_check_null_pointer
33		LDR r4, [r4]
34		LDR r4, [r4]
35		STR r4, [sp, #68]
36		LDR r4, [sp, #72]
37		STR r4, [sp, #64]
38		LDR r0, =8
39		BL malloc
40		MOV r4, r0
41		LDR r5, [sp, #72]
42		LDR r0, =4
43		BL malloc
44		STR r5, [r0]
45		STR r0, [r4]
46		LDR r5, [sp, #68]
47		LDR r0, =4
48		BL malloc
49		STR r5, [r0]
50		STR r0, [r4, #4]
51		STR r4, [sp, #60]
52		LDR r0, =8
53		BL malloc
54		MOV r4, r0
55		LDR r5, [sp, #60]
56		LDR r0, =4
57		BL malloc
58		STR r5, [r0]
59		STR r0, [r4]
60		LDR r5, =0
61		LDR r0, =4
62		BL malloc
63		STR r5, [r0]
64		STR r0, [r4, #4]
65		STR r4, [sp, #56]
66		LDR r0, =8
67		BL malloc
68		MOV r4, r0
69		LDR r5, =0
70		LDR r0, =4
71		BL malloc
72		STR r5, [r0]
73		STR r0, [r4]
74		LDR r5, =0
75		LDR r0, =4
76		BL malloc
77		STR r5, [r0]
78		STR r0, [r4, #4]
79		STR r4, [sp, #52]
80		LDR r4, [sp, #52]
81		STR r4, [sp, #48]
82		LDR r0, =8
83		BL malloc
84		MOV r4, r0
85		LDR r5, =0
86		LDR r0, =4
87		BL malloc
88		STR r5, [r0]
89		STR r0, [r4]
90		LDR r5, =0
91		LDR r0, =4
92		BL malloc
93		STR r5, [r0]
94		STR r0, [r4, #4]
95		STR r4, [sp, #44]
96		LDR r0, =16
97		BL malloc
98		MOV r4, r0
99		LDR r5, [sp, #72]
100		STR r5, [r4, #4]
101		LDR r5, [sp, #68]
102		STR r5, [r4, #8]
103		LDR r5, [sp, #64]
104		STR r5, [r4, #12]
105		LDR r5, =3
106		STR r5, [r4]
107		STR r4, [sp, #40]
108		LDR r4, [sp, #40]
109		STR r4, [sp, #-4]!
110		BL f_f
111		ADD sp, sp, #4
112		MOV r4, r0
113		STR r4, [sp, #36]
114		LDR r0, =16
115		BL malloc
116		MOV r4, r0
117		LDR r5, =0
118		STR r5, [r4, #4]
119		LDR r5, =0
120		STR r5, [r4, #8]
121		LDR r5, [sp, #60]
122		STR r5, [r4, #12]
123		LDR r5, =3
124		STR r5, [r4]
125		STR r4, [sp, #32]
126		LDR r0, =16
127		BL malloc
128		MOV r4, r0
129		LDR r5, [sp, #36]
130		STR r5, [r4, #4]
131		LDR r5, [sp, #32]
132		STR r5, [r4, #8]
133		LDR r5, [sp, #40]
134		STR r5, [r4, #12]
135		LDR r5, =3
136		STR r5, [r4]
137		STR r4, [sp, #28]
138		LDR r0, =12
139		BL malloc
140		MOV r4, r0
141		ADD r5, sp, #28
142		LDR r6, =0
143		LDR r5, [r5]
144		MOV r0, r6
145		MOV r1, r5
146		BL p_check_array_bounds
147		ADD r5, r5, #4
148		ADD r5, r5, r6, LSL #2
149		ADD r6, sp, #28
150		LDR r7, [sp, #32]
151		LDR r7, [r7]
152		LDR r6, [r6]
153		MOV r0, r7
154		MOV r1, r6
155		BL p_check_array_bounds
156		ADD r6, r6, #4
157		ADD r6, r6, r7, LSL #2
158		LDR r6, [r6]
159		LDR r6, [r6]
160		LDR r5, [r5]
161		MOV r0, r6
162		MOV r1, r5
163		BL p_check_array_bounds
164		ADD r5, r5, #4
165		ADD r5, r5, r6, LSL #2
166		LDR r5, [r5]
167		STR r5, [r4, #4]
168		ADD r5, sp, #28
169		LDR r6, =1
170		LDR r5, [r5]
171		MOV r0, r6
172		MOV r1, r5
173		BL p_check_array_bounds
174		ADD r5, r5, #4
175		ADD r5, r5, r6, LSL #2
176		LDR r6, =2
177		LDR r5, [r5]
178		MOV r0, r6
179		MOV r1, r5
180		BL p_check_array_bounds
181		ADD r5, r5, #4
182		ADD r5, r5, r6, LSL #2
183		LDR r5, [r5]
184		STR r5, [r4, #8]
185		LDR r5, =2
186		STR r5, [r4]
187		STR r4, [sp, #24]
188		LDR r0, =12
189		BL malloc
190		MOV r4, r0
191		ADD r5, sp, #28
192		LDR r6, [sp, #32]
193		LDR r6, [r6]
194		LDR r5, [r5]
195		MOV r0, r6
196		MOV r1, r5
197		BL p_check_array_bounds
198		ADD r5, r5, #4
199		ADD r5, r5, r6, LSL #2
200		LDR r5, [r5]
201		STR r5, [r4, #4]
202		LDR r5, [sp, #36]
203		STR r5, [r4, #8]
204		LDR r5, =2
205		STR r5, [r4]
206		STR r4, [sp, #20]
207		LDR r0, =8
208		BL malloc
209		MOV r4, r0
210		ADD r5, sp, #20
211		LDR r6, =1
212		LDR r5, [r5]
213		MOV r0, r6
214		MOV r1, r5
215		BL p_check_array_bounds
216		ADD r5, r5, #4
217		ADD r5, r5, r6, LSL #2
218		LDR r6, =2
219		LDR r5, [r5]
220		MOV r0, r6
221		MOV r1, r5
222		BL p_check_array_bounds
223		ADD r5, r5, #4
224		ADD r5, r5, r6, LSL #2
225		LDR r5, [r5]
226		LDR r0, =4
227		BL malloc
228		STR r5, [r0]
229		STR r0, [r4]
230		LDR r5, =0
231		LDR r0, =4
232		BL malloc
233		STR r5, [r0]
234		STR r0, [r4, #4]
235		STR r4, [sp, #16]
236		LDR r0, =8
237		BL malloc
238		MOV r4, r0
239		LDR r5, =0
240		LDR r0, =4
241		BL malloc
242		STR r5, [r0]
243		STR r0, [r4]
244		ADD r5, sp, #24
245		ADD r6, sp, #28
246		ADD r7, sp, #28
247		ADD r8, sp, #28
248		LDR r9, [sp, #28]
249		LDR r9, [r9]
250		LDR r8, [r8]
251		MOV r0, r9
252		MOV r1, r8
253		BL p_check_array_bounds
254		ADD r8, r8, #4
255		ADD r8, r8, r9, LSL #2
256		LDR r8, [r8]
257		LDR r8, [r8]
258		LDR r7, [r7]
259		MOV r0, r8
260		MOV r1, r7
261		BL p_check_array_bounds
262		ADD r7, r7, #4
263		ADD r7, r7, r8, LSL #2
264		LDR r7, [r7]
265		LDR r7, [r7]
266		LDR r6, [r6]
267		MOV r0, r7
268		MOV r1, r6
269		BL p_check_array_bounds
270		ADD r6, r6, #4
271		ADD r6, r6, r7, LSL #2
272		LDR r6, [r6]
273		LDR r6, [r6]
274		LDR r5, [r5]
275		MOV r0, r6
276		MOV r1, r5
277		BL p_check_array_bounds
278		ADD r5, r5, #4
279		ADD r5, r5, r6, LSL #2
280		LDR r5, [r5]
281		LDR r0, =4
282		BL malloc
283		STR r5, [r0]
284		STR r0, [r4, #4]
285		STR r4, [sp, #12]
286		LDR r4, [sp, #72]
287		MOV r0, r4
288		BL p_check_null_pointer
289		LDR r4, [r4]
290		LDR r4, [r4]
291		STR r4, [sp, #8]
292		LDR r4, [sp, #8]
293		MOV r0, r4
294		BL p_check_null_pointer
295		LDR r4, [r4]
296		LDR r4, [r4]
297		STR r4, [sp, #4]
298		LDR r0, =8
299		BL malloc
300		MOV r4, r0
301		LDR r5, [sp, #4]
302		LDR r0, =4
303		BL malloc
304		STR r5, [r0]
305		STR r0, [r4]
306		ADD r5, sp, #20
307		LDR r6, =0
308		LDR r5, [r5]
309		MOV r0, r6
310		MOV r1, r5
311		BL p_check_array_bounds
312		ADD r5, r5, #4
313		ADD r5, r5, r6, LSL #2
314		LDR r5, [r5]
315		LDR r0, =4
316		BL malloc
317		STR r5, [r0]
318		STR r0, [r4, #4]
319		STR r4, [sp]
320		ADD sp, sp, #76
321		LDR r0, =0
322		POP {pc}
323		.ltorg
324	p_check_null_pointer:
325		PUSH {lr}
326		CMP r0, #0
327		LDREQ r0, =msg_0
328		BLEQ p_throw_runtime_error
329		POP {pc}
330	p_check_array_bounds:
331		PUSH {lr}
332		CMP r0, #0
333		LDRLT r0, =msg_1
334		BLLT p_throw_runtime_error
335		LDR r1, [r1]
336		CMP r0, r1
337		LDRCS r0, =msg_2
338		BLCS p_throw_runtime_error
339		POP {pc}
340	p_throw_runtime_error:
341		BL p_print_string
342		MOV r0, #-1
343		BL exit
344	p_print_string:
345		PUSH {lr}
346		LDR r1, [r0]
347		ADD r2, r0, #4
348		LDR r0, =msg_3
349		ADD r0, r0, #4
350		BL printf
351		MOV r0, #0
352		BL fflush
353		POP {pc}
354	
===========================================================
-- Finished
