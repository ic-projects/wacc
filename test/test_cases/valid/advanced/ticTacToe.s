-- Compiling...
-- Printing Assembly...
ticTacToe.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 38
4		.ascii	"========= Tic Tac Toe ================"
5	msg_1:
6		.word 38
7		.ascii	"=  Because we know you want to win   ="
8	msg_2:
9		.word 38
10		.ascii	"======================================"
11	msg_3:
12		.word 38
13		.ascii	"=                                    ="
14	msg_4:
15		.word 38
16		.ascii	"= Who would you like to be?          ="
17	msg_5:
18		.word 38
19		.ascii	"=   x  (play first)                  ="
20	msg_6:
21		.word 38
22		.ascii	"=   o  (play second)                 ="
23	msg_7:
24		.word 38
25		.ascii	"=   q  (quit)                        ="
26	msg_8:
27		.word 38
28		.ascii	"=                                    ="
29	msg_9:
30		.word 38
31		.ascii	"======================================"
32	msg_10:
33		.word 39
34		.ascii	"Which symbol you would like to choose: "
35	msg_11:
36		.word 15
37		.ascii	"Goodbye safety."
38	msg_12:
39		.word 16
40		.ascii	"Invalid symbol: "
41	msg_13:
42		.word 17
43		.ascii	"Please try again."
44	msg_14:
45		.word 17
46		.ascii	"You have chosen: "
47	msg_15:
48		.word 6
49		.ascii	" 1 2 3"
50	msg_16:
51		.word 1
52		.ascii	"1"
53	msg_17:
54		.word 6
55		.ascii	" -+-+-"
56	msg_18:
57		.word 1
58		.ascii	"2"
59	msg_19:
60		.word 6
61		.ascii	" -+-+-"
62	msg_20:
63		.word 1
64		.ascii	"3"
65	msg_21:
66		.word 0
67		.ascii	""
68	msg_22:
69		.word 0
70		.ascii	""
71	msg_23:
72		.word 23
73		.ascii	"What is your next move?"
74	msg_24:
75		.word 12
76		.ascii	" row (1-3): "
77	msg_25:
78		.word 15
79		.ascii	" column (1-3): "
80	msg_26:
81		.word 0
82		.ascii	""
83	msg_27:
84		.word 39
85		.ascii	"Your move is invalid. Please try again."
86	msg_28:
87		.word 21
88		.ascii	"The AI played at row "
89	msg_29:
90		.word 8
91		.ascii	" column "
92	msg_30:
93		.word 31
94		.ascii	"AI is cleaning up its memory..."
95	msg_31:
96		.word 52
97		.ascii	"Internal Error: cannot find the next move for the AI"
98	msg_32:
99		.word 31
100		.ascii	"AI is cleaning up its memory..."
101	msg_33:
102		.word 50
103		.ascii	"Internal Error: symbol given is neither \'x\' or \'o\'"
104	msg_34:
105		.word 58
106		.ascii	"Initialising AI. Please wait, this may take a few minutes."
107	msg_35:
108		.word 9
109		.ascii	" has won!"
110	msg_36:
111		.word 10
112		.ascii	"Stalemate!"
113	msg_37:
114		.word 5
115		.ascii	"%.*s\0"
116	msg_38:
117		.word 1
118		.ascii	"\0"
119	msg_39:
120		.word 4
121		.ascii	" %c\0"
122	msg_40:
123		.word 50
124		.ascii	"NullReferenceError: dereference a null reference\n\0"
125	msg_41:
126		.word 3
127		.ascii	"%d\0"
128	msg_42:
129		.word 44
130		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
131	msg_43:
132		.word 45
133		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
134	msg_44:
135		.word 3
136		.ascii	"%d\0"
137	msg_45:
138		.word 50
139		.ascii	"NullReferenceError: dereference a null reference\n\0"
140	msg_46:
141		.word 82
142		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
143	
144	.text
145	
146	.global main
147	f_chooseSymbol:
148		PUSH {lr}
149		SUB sp, sp, #1
150		LDR r4, =msg_0
151		MOV r0, r4
152		BL p_print_string
153		BL p_print_ln
154		LDR r4, =msg_1
155		MOV r0, r4
156		BL p_print_string
157		BL p_print_ln
158		LDR r4, =msg_2
159		MOV r0, r4
160		BL p_print_string
161		BL p_print_ln
162		LDR r4, =msg_3
163		MOV r0, r4
164		BL p_print_string
165		BL p_print_ln
166		LDR r4, =msg_4
167		MOV r0, r4
168		BL p_print_string
169		BL p_print_ln
170		LDR r4, =msg_5
171		MOV r0, r4
172		BL p_print_string
173		BL p_print_ln
174		LDR r4, =msg_6
175		MOV r0, r4
176		BL p_print_string
177		BL p_print_ln
178		LDR r4, =msg_7
179		MOV r0, r4
180		BL p_print_string
181		BL p_print_ln
182		LDR r4, =msg_8
183		MOV r0, r4
184		BL p_print_string
185		BL p_print_ln
186		LDR r4, =msg_9
187		MOV r0, r4
188		BL p_print_string
189		BL p_print_ln
190		MOV r4, #0
191		STRB r4, [sp]
192		B L0
193	L1:
194		SUB sp, sp, #1
195		LDR r4, =msg_10
196		MOV r0, r4
197		BL p_print_string
198		MOV r4, #0
199		STRB r4, [sp]
200		ADD r4, sp, #0
201		MOV r0, r4
202		BL p_read_char
203		LDRSB r4, [sp]
204		MOV r5, #'x'
205		CMP r4, r5
206		MOVEQ r4, #1
207		MOVNE r4, #0
208		LDRSB r5, [sp]
209		MOV r6, #'X'
210		CMP r5, r6
211		MOVEQ r5, #1
212		MOVNE r5, #0
213		ORR r4, r4, r5
214		CMP r4, #0
215		BEQ L2
216		MOV r4, #'x'
217		STRB r4, [sp, #1]
218		B L3
219	L2:
220		LDRSB r4, [sp]
221		MOV r5, #'o'
222		CMP r4, r5
223		MOVEQ r4, #1
224		MOVNE r4, #0
225		LDRSB r5, [sp]
226		MOV r6, #'O'
227		CMP r5, r6
228		MOVEQ r5, #1
229		MOVNE r5, #0
230		ORR r4, r4, r5
231		CMP r4, #0
232		BEQ L4
233		MOV r4, #'o'
234		STRB r4, [sp, #1]
235		B L5
236	L4:
237		LDRSB r4, [sp]
238		MOV r5, #'q'
239		CMP r4, r5
240		MOVEQ r4, #1
241		MOVNE r4, #0
242		LDRSB r5, [sp]
243		MOV r6, #'Q'
244		CMP r5, r6
245		MOVEQ r5, #1
246		MOVNE r5, #0
247		ORR r4, r4, r5
248		CMP r4, #0
249		BEQ L6
250		LDR r4, =msg_11
251		MOV r0, r4
252		BL p_print_string
253		BL p_print_ln
254		LDR r4, =0
255		MOV r0, r4
256		BL exit
257		B L7
258	L6:
259		LDR r4, =msg_12
260		MOV r0, r4
261		BL p_print_string
262		LDRSB r4, [sp]
263		MOV r0, r4
264		BL putchar
265		BL p_print_ln
266		LDR r4, =msg_13
267		MOV r0, r4
268		BL p_print_string
269		BL p_print_ln
270	L7:
271	L5:
272	L3:
273		ADD sp, sp, #1
274	L0:
275		LDRSB r4, [sp]
276		MOV r5, #0
277		CMP r4, r5
278		MOVEQ r4, #1
279		MOVNE r4, #0
280		CMP r4, #1
281		BEQ L1
282		LDR r4, =msg_14
283		MOV r0, r4
284		BL p_print_string
285		LDRSB r4, [sp]
286		MOV r0, r4
287		BL putchar
288		BL p_print_ln
289		LDRSB r4, [sp]
290		MOV r0, r4
291		ADD sp, sp, #1
292		POP {pc}
293		POP {pc}
294		.ltorg
295	f_printBoard:
296		PUSH {lr}
297		SUB sp, sp, #17
298		LDR r4, [sp, #21]
299		MOV r0, r4
300		BL p_check_null_pointer
301		LDR r4, [r4]
302		LDR r4, [r4]
303		STR r4, [sp, #13]
304		LDR r4, [sp, #13]
305		MOV r0, r4
306		BL p_check_null_pointer
307		LDR r4, [r4]
308		LDR r4, [r4]
309		STR r4, [sp, #9]
310		LDR r4, [sp, #13]
311		MOV r0, r4
312		BL p_check_null_pointer
313		LDR r4, [r4, #4]
314		LDR r4, [r4]
315		STR r4, [sp, #5]
316		LDR r4, [sp, #21]
317		MOV r0, r4
318		BL p_check_null_pointer
319		LDR r4, [r4, #4]
320		LDR r4, [r4]
321		STR r4, [sp, #1]
322		LDR r4, =msg_15
323		MOV r0, r4
324		BL p_print_string
325		BL p_print_ln
326		LDR r4, =msg_16
327		MOV r0, r4
328		BL p_print_string
329		LDR r4, [sp, #9]
330		STR r4, [sp, #-4]!
331		BL f_printRow
332		ADD sp, sp, #4
333		MOV r4, r0
334		STRB r4, [sp]
335		LDR r4, =msg_17
336		MOV r0, r4
337		BL p_print_string
338		BL p_print_ln
339		LDR r4, =msg_18
340		MOV r0, r4
341		BL p_print_string
342		LDR r4, [sp, #5]
343		STR r4, [sp, #-4]!
344		BL f_printRow
345		ADD sp, sp, #4
346		MOV r4, r0
347		STRB r4, [sp]
348		LDR r4, =msg_19
349		MOV r0, r4
350		BL p_print_string
351		BL p_print_ln
352		LDR r4, =msg_20
353		MOV r0, r4
354		BL p_print_string
355		LDR r4, [sp, #1]
356		STR r4, [sp, #-4]!
357		BL f_printRow
358		ADD sp, sp, #4
359		MOV r4, r0
360		STRB r4, [sp]
361		LDR r4, =msg_21
362		MOV r0, r4
363		BL p_print_string
364		BL p_print_ln
365		MOV r4, #1
366		MOV r0, r4
367		ADD sp, sp, #17
368		POP {pc}
369		POP {pc}
370		.ltorg
371	f_printRow:
372		PUSH {lr}
373		SUB sp, sp, #8
374		LDR r4, [sp, #12]
375		MOV r0, r4
376		BL p_check_null_pointer
377		LDR r4, [r4]
378		LDR r4, [r4]
379		STR r4, [sp, #4]
380		LDR r4, [sp, #4]
381		MOV r0, r4
382		BL p_check_null_pointer
383		LDR r4, [r4]
384		LDRSB r4, [r4]
385		STRB r4, [sp, #3]
386		LDR r4, [sp, #4]
387		MOV r0, r4
388		BL p_check_null_pointer
389		LDR r4, [r4, #4]
390		LDRSB r4, [r4]
391		STRB r4, [sp, #2]
392		LDR r4, [sp, #12]
393		MOV r0, r4
394		BL p_check_null_pointer
395		LDR r4, [r4, #4]
396		LDRSB r4, [r4]
397		STRB r4, [sp, #1]
398		LDRSB r4, [sp, #3]
399		STRB r4, [sp, #-1]!
400		BL f_printCell
401		ADD sp, sp, #1
402		MOV r4, r0
403		STRB r4, [sp]
404		MOV r4, #'|'
405		MOV r0, r4
406		BL putchar
407		LDRSB r4, [sp, #2]
408		STRB r4, [sp, #-1]!
409		BL f_printCell
410		ADD sp, sp, #1
411		MOV r4, r0
412		STRB r4, [sp]
413		MOV r4, #'|'
414		MOV r0, r4
415		BL putchar
416		LDRSB r4, [sp, #1]
417		STRB r4, [sp, #-1]!
418		BL f_printCell
419		ADD sp, sp, #1
420		MOV r4, r0
421		STRB r4, [sp]
422		LDR r4, =msg_22
423		MOV r0, r4
424		BL p_print_string
425		BL p_print_ln
426		MOV r4, #1
427		MOV r0, r4
428		ADD sp, sp, #8
429		POP {pc}
430		POP {pc}
431		.ltorg
432	f_printCell:
433		PUSH {lr}
434		LDRSB r4, [sp, #4]
435		MOV r5, #0
436		CMP r4, r5
437		MOVEQ r4, #1
438		MOVNE r4, #0
439		CMP r4, #0
440		BEQ L8
441		MOV r4, #' '
442		MOV r0, r4
443		BL putchar
444		B L9
445	L8:
446		LDRSB r4, [sp, #4]
447		MOV r0, r4
448		BL putchar
449	L9:
450		MOV r4, #1
451		MOV r0, r4
452		POP {pc}
453		POP {pc}
454		.ltorg
455	f_askForAMoveHuman:
456		PUSH {lr}
457		SUB sp, sp, #9
458		MOV r4, #0
459		STRB r4, [sp, #8]
460		LDR r4, =0
461		STR r4, [sp, #4]
462		LDR r4, =0
463		STR r4, [sp]
464		B L10
465	L11:
466		LDR r4, =msg_23
467		MOV r0, r4
468		BL p_print_string
469		BL p_print_ln
470		LDR r4, =msg_24
471		MOV r0, r4
472		BL p_print_string
473		ADD r4, sp, #4
474		MOV r0, r4
475		BL p_read_int
476		LDR r4, =msg_25
477		MOV r0, r4
478		BL p_print_string
479		ADD r4, sp, #0
480		MOV r0, r4
481		BL p_read_int
482		LDR r4, [sp]
483		STR r4, [sp, #-4]!
484		LDR r4, [sp, #8]
485		STR r4, [sp, #-4]!
486		LDR r4, [sp, #21]
487		STR r4, [sp, #-4]!
488		BL f_validateMove
489		ADD sp, sp, #12
490		MOV r4, r0
491		STRB r4, [sp, #8]
492		LDRSB r4, [sp, #8]
493		CMP r4, #0
494		BEQ L12
495		LDR r4, =msg_26
496		MOV r0, r4
497		BL p_print_string
498		BL p_print_ln
499		LDR r4, [sp, #4]
500		ADD r5, sp, #17
501		LDR r6, =0
502		LDR r5, [r5]
503		MOV r0, r6
504		MOV r1, r5
505		BL p_check_array_bounds
506		ADD r5, r5, #4
507		ADD r5, r5, r6, LSL #2
508		STR r4, [r5]
509		LDR r4, [sp]
510		ADD r6, sp, #17
511		LDR r7, =1
512		LDR r6, [r6]
513		MOV r0, r7
514		MOV r1, r6
515		BL p_check_array_bounds
516		ADD r6, r6, #4
517		ADD r6, r6, r7, LSL #2
518		STR r4, [r6]
519		MOV r4, #1
520		MOV r0, r4
521		ADD sp, sp, #9
522		POP {pc}
523		B L13
524	L12:
525		LDR r4, =msg_27
526		MOV r0, r4
527		BL p_print_string
528		BL p_print_ln
529	L13:
530	L10:
531		LDRSB r4, [sp, #8]
532		EOR r4, r4, #1
533		CMP r4, #1
534		BEQ L11
535		MOV r4, #1
536		MOV r0, r4
537		ADD sp, sp, #9
538		POP {pc}
539		POP {pc}
540		.ltorg
541	f_validateMove:
542		PUSH {lr}
543		LDR r4, =1
544		LDR r5, [sp, #8]
545		CMP r4, r5
546		MOVLE r4, #1
547		MOVGT r4, #0
548		LDR r5, [sp, #8]
549		LDR r6, =3
550		CMP r5, r6
551		MOVLE r5, #1
552		MOVGT r5, #0
553		AND r4, r4, r5
554		LDR r5, =1
555		LDR r6, [sp, #12]
556		CMP r5, r6
557		MOVLE r5, #1
558		MOVGT r5, #0
559		AND r4, r4, r5
560		LDR r5, [sp, #12]
561		LDR r6, =3
562		CMP r5, r6
563		MOVLE r5, #1
564		MOVGT r5, #0
565		AND r4, r4, r5
566		CMP r4, #0
567		BEQ L14
568		SUB sp, sp, #1
569		LDR r4, [sp, #13]
570		STR r4, [sp, #-4]!
571		LDR r4, [sp, #13]
572		STR r4, [sp, #-4]!
573		LDR r4, [sp, #13]
574		STR r4, [sp, #-4]!
575		BL f_symbolAt
576		ADD sp, sp, #12
577		MOV r4, r0
578		STRB r4, [sp]
579		LDRSB r4, [sp]
580		MOV r5, #0
581		CMP r4, r5
582		MOVEQ r4, #1
583		MOVNE r4, #0
584		MOV r0, r4
585		ADD sp, sp, #1
586		POP {pc}
587		ADD sp, sp, #1
588		B L15
589	L14:
590		MOV r4, #0
591		MOV r0, r4
592		POP {pc}
593	L15:
594		POP {pc}
595		.ltorg
596	f_notifyMoveHuman:
597		PUSH {lr}
598		LDR r4, =msg_28
599		MOV r0, r4
600		BL p_print_string
601		LDR r4, [sp, #10]
602		MOV r0, r4
603		BL p_print_int
604		LDR r4, =msg_29
605		MOV r0, r4
606		BL p_print_string
607		LDR r4, [sp, #14]
608		MOV r0, r4
609		BL p_print_int
610		BL p_print_ln
611		MOV r4, #1
612		MOV r0, r4
613		POP {pc}
614		POP {pc}
615		.ltorg
616	f_initAI:
617		PUSH {lr}
618		SUB sp, sp, #16
619		LDR r0, =8
620		BL malloc
621		MOV r4, r0
622		LDRSB r5, [sp, #20]
623		LDR r0, =1
624		BL malloc
625		STRB r5, [r0]
626		STR r0, [r4]
627		LDR r5, =0
628		LDR r0, =4
629		BL malloc
630		STR r5, [r0]
631		STR r0, [r4, #4]
632		STR r4, [sp, #12]
633		LDRSB r4, [sp, #20]
634		STRB r4, [sp, #-1]!
635		BL f_generateAllPossibleStates
636		ADD sp, sp, #1
637		MOV r4, r0
638		STR r4, [sp, #8]
639		MOV r4, #'x'
640		STRB r4, [sp, #-1]!
641		LDRSB r4, [sp, #21]
642		STRB r4, [sp, #-1]!
643		LDR r4, [sp, #10]
644		STR r4, [sp, #-4]!
645		BL f_setValuesForAllStates
646		ADD sp, sp, #6
647		MOV r4, r0
648		STR r4, [sp, #4]
649		LDR r0, =8
650		BL malloc
651		MOV r4, r0
652		LDR r5, [sp, #12]
653		LDR r0, =4
654		BL malloc
655		STR r5, [r0]
656		STR r0, [r4]
657		LDR r5, [sp, #8]
658		LDR r0, =4
659		BL malloc
660		STR r5, [r0]
661		STR r0, [r4, #4]
662		STR r4, [sp]
663		LDR r4, [sp]
664		MOV r0, r4
665		ADD sp, sp, #16
666		POP {pc}
667		POP {pc}
668		.ltorg
669	f_generateAllPossibleStates:
670		PUSH {lr}
671		SUB sp, sp, #8
672		BL f_allocateNewBoard
673		MOV r4, r0
674		STR r4, [sp, #4]
675		LDR r4, [sp, #4]
676		STR r4, [sp, #-4]!
677		BL f_convertFromBoardToState
678		ADD sp, sp, #4
679		MOV r4, r0
680		STR r4, [sp]
681		MOV r4, #'x'
682		STRB r4, [sp, #-1]!
683		LDR r4, [sp, #1]
684		STR r4, [sp, #-4]!
685		BL f_generateNextStates
686		ADD sp, sp, #5
687		MOV r4, r0
688		STR r4, [sp]
689		LDR r4, [sp]
690		MOV r0, r4
691		ADD sp, sp, #8
692		POP {pc}
693		POP {pc}
694		.ltorg
695	f_convertFromBoardToState:
696		PUSH {lr}
697		SUB sp, sp, #12
698		BL f_generateEmptyPointerBoard
699		MOV r4, r0
700		STR r4, [sp, #8]
701		LDR r0, =8
702		BL malloc
703		MOV r4, r0
704		LDR r5, [sp, #16]
705		LDR r0, =4
706		BL malloc
707		STR r5, [r0]
708		STR r0, [r4]
709		LDR r5, [sp, #8]
710		LDR r0, =4
711		BL malloc
712		STR r5, [r0]
713		STR r0, [r4, #4]
714		STR r4, [sp, #4]
715		LDR r0, =8
716		BL malloc
717		MOV r4, r0
718		LDR r5, [sp, #4]
719		LDR r0, =4
720		BL malloc
721		STR r5, [r0]
722		STR r0, [r4]
723		LDR r5, =0
724		LDR r0, =4
725		BL malloc
726		STR r5, [r0]
727		STR r0, [r4, #4]
728		STR r4, [sp]
729		LDR r4, [sp]
730		MOV r0, r4
731		ADD sp, sp, #12
732		POP {pc}
733		POP {pc}
734		.ltorg
735	f_generateEmptyPointerBoard:
736		PUSH {lr}
737		SUB sp, sp, #20
738		BL f_generateEmptyPointerRow
739		MOV r4, r0
740		STR r4, [sp, #16]
741		BL f_generateEmptyPointerRow
742		MOV r4, r0
743		STR r4, [sp, #12]
744		BL f_generateEmptyPointerRow
745		MOV r4, r0
746		STR r4, [sp, #8]
747		LDR r0, =8
748		BL malloc
749		MOV r4, r0
750		LDR r5, [sp, #16]
751		LDR r0, =4
752		BL malloc
753		STR r5, [r0]
754		STR r0, [r4]
755		LDR r5, [sp, #12]
756		LDR r0, =4
757		BL malloc
758		STR r5, [r0]
759		STR r0, [r4, #4]
760		STR r4, [sp, #4]
761		LDR r0, =8
762		BL malloc
763		MOV r4, r0
764		LDR r5, [sp, #4]
765		LDR r0, =4
766		BL malloc
767		STR r5, [r0]
768		STR r0, [r4]
769		LDR r5, [sp, #8]
770		LDR r0, =4
771		BL malloc
772		STR r5, [r0]
773		STR r0, [r4, #4]
774		STR r4, [sp]
775		LDR r4, [sp]
776		MOV r0, r4
777		ADD sp, sp, #20
778		POP {pc}
779		POP {pc}
780		.ltorg
781	f_generateEmptyPointerRow:
782		PUSH {lr}
783		SUB sp, sp, #8
784		LDR r0, =8
785		BL malloc
786		MOV r4, r0
787		LDR r5, =0
788		LDR r0, =4
789		BL malloc
790		STR r5, [r0]
791		STR r0, [r4]
792		LDR r5, =0
793		LDR r0, =4
794		BL malloc
795		STR r5, [r0]
796		STR r0, [r4, #4]
797		STR r4, [sp, #4]
798		LDR r0, =8
799		BL malloc
800		MOV r4, r0
801		LDR r5, [sp, #4]
802		LDR r0, =4
803		BL malloc
804		STR r5, [r0]
805		STR r0, [r4]
806		LDR r5, =0
807		LDR r0, =4
808		BL malloc
809		STR r5, [r0]
810		STR r0, [r4, #4]
811		STR r4, [sp]
812		LDR r4, [sp]
813		MOV r0, r4
814		ADD sp, sp, #8
815		POP {pc}
816		POP {pc}
817		.ltorg
818	f_generateNextStates:
819		PUSH {lr}
820		SUB sp, sp, #14
821		LDR r4, [sp, #18]
822		MOV r0, r4
823		BL p_check_null_pointer
824		LDR r4, [r4]
825		LDR r4, [r4]
826		STR r4, [sp, #10]
827		LDR r4, [sp, #10]
828		MOV r0, r4
829		BL p_check_null_pointer
830		LDR r4, [r4]
831		LDR r4, [r4]
832		STR r4, [sp, #6]
833		LDR r4, [sp, #10]
834		MOV r0, r4
835		BL p_check_null_pointer
836		LDR r4, [r4, #4]
837		LDR r4, [r4]
838		STR r4, [sp, #2]
839		LDRSB r4, [sp, #22]
840		STRB r4, [sp, #-1]!
841		BL f_oppositeSymbol
842		ADD sp, sp, #1
843		MOV r4, r0
844		STRB r4, [sp, #1]
845		LDRSB r4, [sp, #1]
846		STRB r4, [sp, #-1]!
847		LDR r4, [sp, #7]
848		STR r4, [sp, #-4]!
849		BL f_hasWon
850		ADD sp, sp, #5
851		MOV r4, r0
852		STRB r4, [sp]
853		LDRSB r4, [sp]
854		CMP r4, #0
855		BEQ L16
856		LDR r4, [sp, #18]
857		MOV r0, r4
858		ADD sp, sp, #14
859		POP {pc}
860		B L17
861	L16:
862		SUB sp, sp, #1
863		LDRSB r4, [sp, #23]
864		STRB r4, [sp, #-1]!
865		LDR r4, [sp, #4]
866		STR r4, [sp, #-4]!
867		LDR r4, [sp, #12]
868		STR r4, [sp, #-4]!
869		BL f_generateNextStatesBoard
870		ADD sp, sp, #9
871		MOV r4, r0
872		STRB r4, [sp]
873		LDR r4, [sp, #19]
874		MOV r0, r4
875		ADD sp, sp, #15
876		POP {pc}
877		ADD sp, sp, #1
878	L17:
879		POP {pc}
880		.ltorg
881	f_generateNextStatesBoard:
882		PUSH {lr}
883		SUB sp, sp, #33
884		LDR r4, [sp, #37]
885		MOV r0, r4
886		BL p_check_null_pointer
887		LDR r4, [r4]
888		LDR r4, [r4]
889		STR r4, [sp, #29]
890		LDR r4, [sp, #29]
891		MOV r0, r4
892		BL p_check_null_pointer
893		LDR r4, [r4]
894		LDR r4, [r4]
895		STR r4, [sp, #25]
896		LDR r4, [sp, #29]
897		MOV r0, r4
898		BL p_check_null_pointer
899		LDR r4, [r4, #4]
900		LDR r4, [r4]
901		STR r4, [sp, #21]
902		LDR r4, [sp, #37]
903		MOV r0, r4
904		BL p_check_null_pointer
905		LDR r4, [r4, #4]
906		LDR r4, [r4]
907		STR r4, [sp, #17]
908		LDR r4, [sp, #41]
909		MOV r0, r4
910		BL p_check_null_pointer
911		LDR r4, [r4]
912		LDR r4, [r4]
913		STR r4, [sp, #13]
914		LDR r4, [sp, #13]
915		MOV r0, r4
916		BL p_check_null_pointer
917		LDR r4, [r4]
918		LDR r4, [r4]
919		STR r4, [sp, #9]
920		LDR r4, [sp, #13]
921		MOV r0, r4
922		BL p_check_null_pointer
923		LDR r4, [r4, #4]
924		LDR r4, [r4]
925		STR r4, [sp, #5]
926		LDR r4, [sp, #41]
927		MOV r0, r4
928		BL p_check_null_pointer
929		LDR r4, [r4, #4]
930		LDR r4, [r4]
931		STR r4, [sp, #1]
932		LDR r4, =1
933		STR r4, [sp, #-4]!
934		LDRSB r4, [sp, #49]
935		STRB r4, [sp, #-1]!
936		LDR r4, [sp, #14]
937		STR r4, [sp, #-4]!
938		LDR r4, [sp, #34]
939		STR r4, [sp, #-4]!
940		LDR r4, [sp, #50]
941		STR r4, [sp, #-4]!
942		BL f_generateNextStatesRow
943		ADD sp, sp, #17
944		MOV r4, r0
945		STRB r4, [sp]
946		LDR r4, =2
947		STR r4, [sp, #-4]!
948		LDRSB r4, [sp, #49]
949		STRB r4, [sp, #-1]!
950		LDR r4, [sp, #10]
951		STR r4, [sp, #-4]!
952		LDR r4, [sp, #30]
953		STR r4, [sp, #-4]!
954		LDR r4, [sp, #50]
955		STR r4, [sp, #-4]!
956		BL f_generateNextStatesRow
957		ADD sp, sp, #17
958		MOV r4, r0
959		STRB r4, [sp]
960		LDR r4, =3
961		STR r4, [sp, #-4]!
962		LDRSB r4, [sp, #49]
963		STRB r4, [sp, #-1]!
964		LDR r4, [sp, #6]
965		STR r4, [sp, #-4]!
966		LDR r4, [sp, #26]
967		STR r4, [sp, #-4]!
968		LDR r4, [sp, #50]
969		STR r4, [sp, #-4]!
970		BL f_generateNextStatesRow
971		ADD sp, sp, #17
972		MOV r4, r0
973		STRB r4, [sp]
974		MOV r4, #1
975		MOV r0, r4
976		ADD sp, sp, #33
977		POP {pc}
978		POP {pc}
979		.ltorg
980	f_generateNextStatesRow:
981		PUSH {lr}
982		SUB sp, sp, #11
983		LDR r4, [sp, #19]
984		MOV r0, r4
985		BL p_check_null_pointer
986		LDR r4, [r4]
987		LDR r4, [r4]
988		STR r4, [sp, #7]
989		LDR r4, [sp, #7]
990		MOV r0, r4
991		BL p_check_null_pointer
992		LDR r4, [r4]
993		LDRSB r4, [r4]
994		STRB r4, [sp, #6]
995		LDR r4, [sp, #7]
996		MOV r0, r4
997		BL p_check_null_pointer
998		LDR r4, [r4, #4]
999		LDRSB r4, [r4]
1000		STRB r4, [sp, #5]
1001		LDR r4, [sp, #19]
1002		MOV r0, r4
1003		BL p_check_null_pointer
1004		LDR r4, [r4, #4]
1005		LDRSB r4, [r4]
1006		STRB r4, [sp, #4]
1007		LDR r4, [sp, #23]
1008		MOV r0, r4
1009		BL p_check_null_pointer
1010		LDR r4, [r4]
1011		LDR r4, [r4]
1012		STR r4, [sp]
1013		LDR r4, =1
1014		STR r4, [sp, #-4]!
1015		LDR r4, [sp, #32]
1016		STR r4, [sp, #-4]!
1017		LDRSB r4, [sp, #35]
1018		STRB r4, [sp, #-1]!
1019		LDRSB r4, [sp, #15]
1020		STRB r4, [sp, #-1]!
1021		LDR r4, [sp, #25]
1022		STR r4, [sp, #-4]!
1023		BL f_generateNextStatesCell
1024		ADD sp, sp, #14
1025		MOV r4, r0
1026		LDR r5, [sp]
1027		MOV r0, r5
1028		BL p_check_null_pointer
1029		LDR r5, [r5]
1030		STR r4, [r5]
1031		LDR r4, =2
1032		STR r4, [sp, #-4]!
1033		LDR r4, [sp, #32]
1034		STR r4, [sp, #-4]!
1035		LDRSB r4, [sp, #35]
1036		STRB r4, [sp, #-1]!
1037		LDRSB r4, [sp, #14]
1038		STRB r4, [sp, #-1]!
1039		LDR r4, [sp, #25]
1040		STR r4, [sp, #-4]!
1041		BL f_generateNextStatesCell
1042		ADD sp, sp, #14
1043		MOV r4, r0
1044		LDR r5, [sp]
1045		MOV r0, r5
1046		BL p_check_null_pointer
1047		LDR r5, [r5, #4]
1048		STR r4, [r5]
1049		LDR r4, =3
1050		STR r4, [sp, #-4]!
1051		LDR r4, [sp, #32]
1052		STR r4, [sp, #-4]!
1053		LDRSB r4, [sp, #35]
1054		STRB r4, [sp, #-1]!
1055		LDRSB r4, [sp, #13]
1056		STRB r4, [sp, #-1]!
1057		LDR r4, [sp, #25]
1058		STR r4, [sp, #-4]!
1059		BL f_generateNextStatesCell
1060		ADD sp, sp, #14
1061		MOV r4, r0
1062		LDR r5, [sp, #23]
1063		MOV r0, r5
1064		BL p_check_null_pointer
1065		LDR r5, [r5, #4]
1066		STR r4, [r5]
1067		MOV r4, #1
1068		MOV r0, r4
1069		ADD sp, sp, #11
1070		POP {pc}
1071		POP {pc}
1072		.ltorg
1073	f_generateNextStatesCell:
1074		PUSH {lr}
1075		LDRSB r4, [sp, #8]
1076		MOV r5, #0
1077		CMP r4, r5
1078		MOVEQ r4, #1
1079		MOVNE r4, #0
1080		CMP r4, #0
1081		BEQ L18
1082		SUB sp, sp, #10
1083		LDR r4, [sp, #14]
1084		STR r4, [sp, #-4]!
1085		BL f_cloneBoard
1086		ADD sp, sp, #4
1087		MOV r4, r0
1088		STR r4, [sp, #6]
1089		LDR r4, [sp, #24]
1090		STR r4, [sp, #-4]!
1091		LDR r4, [sp, #24]
1092		STR r4, [sp, #-4]!
1093		LDRSB r4, [sp, #27]
1094		STRB r4, [sp, #-1]!
1095		LDR r4, [sp, #15]
1096		STR r4, [sp, #-4]!
1097		BL f_placeMove
1098		ADD sp, sp, #13
1099		MOV r4, r0
1100		STRB r4, [sp, #5]
1101		LDR r4, [sp, #6]
1102		STR r4, [sp, #-4]!
1103		BL f_convertFromBoardToState
1104		ADD sp, sp, #4
1105		MOV r4, r0
1106		STR r4, [sp, #1]
1107		LDRSB r4, [sp, #19]
1108		STRB r4, [sp, #-1]!
1109		BL f_oppositeSymbol
1110		ADD sp, sp, #1
1111		MOV r4, r0
1112		STRB r4, [sp]
1113		LDRSB r4, [sp]
1114		STRB r4, [sp, #-1]!
1115		LDR r4, [sp, #2]
1116		STR r4, [sp, #-4]!
1117		BL f_generateNextStates
1118		ADD sp, sp, #5
1119		MOV r4, r0
1120		STR r4, [sp, #1]
1121		LDR r4, [sp, #1]
1122		MOV r0, r4
1123		ADD sp, sp, #10
1124		POP {pc}
1125		ADD sp, sp, #10
1126		B L19
1127	L18:
1128		LDR r4, =0
1129		MOV r0, r4
1130		POP {pc}
1131	L19:
1132		POP {pc}
1133		.ltorg
1134	f_cloneBoard:
1135		PUSH {lr}
1136		SUB sp, sp, #5
1137		BL f_allocateNewBoard
1138		MOV r4, r0
1139		STR r4, [sp, #1]
1140		LDR r4, [sp, #1]
1141		STR r4, [sp, #-4]!
1142		LDR r4, [sp, #13]
1143		STR r4, [sp, #-4]!
1144		BL f_copyBoard
1145		ADD sp, sp, #8
1146		MOV r4, r0
1147		STRB r4, [sp]
1148		LDR r4, [sp, #1]
1149		MOV r0, r4
1150		ADD sp, sp, #5
1151		POP {pc}
1152		POP {pc}
1153		.ltorg
1154	f_copyBoard:
1155		PUSH {lr}
1156		SUB sp, sp, #33
1157		LDR r4, [sp, #37]
1158		MOV r0, r4
1159		BL p_check_null_pointer
1160		LDR r4, [r4]
1161		LDR r4, [r4]
1162		STR r4, [sp, #29]
1163		LDR r4, [sp, #29]
1164		MOV r0, r4
1165		BL p_check_null_pointer
1166		LDR r4, [r4]
1167		LDR r4, [r4]
1168		STR r4, [sp, #25]
1169		LDR r4, [sp, #29]
1170		MOV r0, r4
1171		BL p_check_null_pointer
1172		LDR r4, [r4, #4]
1173		LDR r4, [r4]
1174		STR r4, [sp, #21]
1175		LDR r4, [sp, #37]
1176		MOV r0, r4
1177		BL p_check_null_pointer
1178		LDR r4, [r4, #4]
1179		LDR r4, [r4]
1180		STR r4, [sp, #17]
1181		LDR r4, [sp, #41]
1182		MOV r0, r4
1183		BL p_check_null_pointer
1184		LDR r4, [r4]
1185		LDR r4, [r4]
1186		STR r4, [sp, #13]
1187		LDR r4, [sp, #13]
1188		MOV r0, r4
1189		BL p_check_null_pointer
1190		LDR r4, [r4]
1191		LDR r4, [r4]
1192		STR r4, [sp, #9]
1193		LDR r4, [sp, #13]
1194		MOV r0, r4
1195		BL p_check_null_pointer
1196		LDR r4, [r4, #4]
1197		LDR r4, [r4]
1198		STR r4, [sp, #5]
1199		LDR r4, [sp, #41]
1200		MOV r0, r4
1201		BL p_check_null_pointer
1202		LDR r4, [r4, #4]
1203		LDR r4, [r4]
1204		STR r4, [sp, #1]
1205		LDR r4, [sp, #9]
1206		STR r4, [sp, #-4]!
1207		LDR r4, [sp, #29]
1208		STR r4, [sp, #-4]!
1209		BL f_copyRow
1210		ADD sp, sp, #8
1211		MOV r4, r0
1212		STRB r4, [sp]
1213		LDR r4, [sp, #5]
1214		STR r4, [sp, #-4]!
1215		LDR r4, [sp, #25]
1216		STR r4, [sp, #-4]!
1217		BL f_copyRow
1218		ADD sp, sp, #8
1219		MOV r4, r0
1220		STRB r4, [sp]
1221		LDR r4, [sp, #1]
1222		STR r4, [sp, #-4]!
1223		LDR r4, [sp, #21]
1224		STR r4, [sp, #-4]!
1225		BL f_copyRow
1226		ADD sp, sp, #8
1227		MOV r4, r0
1228		STRB r4, [sp]
1229		MOV r4, #1
1230		MOV r0, r4
1231		ADD sp, sp, #33
1232		POP {pc}
1233		POP {pc}
1234		.ltorg
1235	f_copyRow:
1236		PUSH {lr}
1237		SUB sp, sp, #8
1238		LDR r4, [sp, #12]
1239		MOV r0, r4
1240		BL p_check_null_pointer
1241		LDR r4, [r4]
1242		LDR r4, [r4]
1243		STR r4, [sp, #4]
1244		LDR r4, [sp, #16]
1245		MOV r0, r4
1246		BL p_check_null_pointer
1247		LDR r4, [r4]
1248		LDR r4, [r4]
1249		STR r4, [sp]
1250		LDR r4, [sp, #4]
1251		MOV r0, r4
1252		BL p_check_null_pointer
1253		LDR r4, [r4]
1254		LDRSB r4, [r4]
1255		LDR r5, [sp]
1256		MOV r0, r5
1257		BL p_check_null_pointer
1258		LDR r5, [r5]
1259		STRB r4, [r5]
1260		LDR r4, [sp, #4]
1261		MOV r0, r4
1262		BL p_check_null_pointer
1263		LDR r4, [r4, #4]
1264		LDRSB r4, [r4]
1265		LDR r5, [sp]
1266		MOV r0, r5
1267		BL p_check_null_pointer
1268		LDR r5, [r5, #4]
1269		STRB r4, [r5]
1270		LDR r4, [sp, #12]
1271		MOV r0, r4
1272		BL p_check_null_pointer
1273		LDR r4, [r4, #4]
1274		LDRSB r4, [r4]
1275		LDR r5, [sp, #16]
1276		MOV r0, r5
1277		BL p_check_null_pointer
1278		LDR r5, [r5, #4]
1279		STRB r4, [r5]
1280		MOV r4, #1
1281		MOV r0, r4
1282		ADD sp, sp, #8
1283		POP {pc}
1284		POP {pc}
1285		.ltorg
1286	f_setValuesForAllStates:
1287		PUSH {lr}
1288		SUB sp, sp, #4
1289		LDR r4, =0
1290		STR r4, [sp]
1291		LDR r4, [sp, #8]
1292		LDR r5, =0
1293		CMP r4, r5
1294		MOVEQ r4, #1
1295		MOVNE r4, #0
1296		CMP r4, #0
1297		BEQ L20
1298		LDRSB r4, [sp, #13]
1299		LDRSB r5, [sp, #12]
1300		CMP r4, r5
1301		MOVEQ r4, #1
1302		MOVNE r4, #0
1303		CMP r4, #0
1304		BEQ L22
1305		LDR r4, =101
1306		STR r4, [sp]
1307		B L23
1308	L22:
1309		LDR r4, =-101
1310		STR r4, [sp]
1311	L23:
1312		B L21
1313	L20:
1314		SUB sp, sp, #14
1315		LDR r4, [sp, #22]
1316		MOV r0, r4
1317		BL p_check_null_pointer
1318		LDR r4, [r4]
1319		LDR r4, [r4]
1320		STR r4, [sp, #10]
1321		LDR r4, [sp, #10]
1322		MOV r0, r4
1323		BL p_check_null_pointer
1324		LDR r4, [r4]
1325		LDR r4, [r4]
1326		STR r4, [sp, #6]
1327		LDR r4, [sp, #10]
1328		MOV r0, r4
1329		BL p_check_null_pointer
1330		LDR r4, [r4, #4]
1331		LDR r4, [r4]
1332		STR r4, [sp, #2]
1333		LDRSB r4, [sp, #27]
1334		STRB r4, [sp, #-1]!
1335		BL f_oppositeSymbol
1336		ADD sp, sp, #1
1337		MOV r4, r0
1338		STRB r4, [sp, #1]
1339		LDRSB r4, [sp, #1]
1340		STRB r4, [sp, #-1]!
1341		LDR r4, [sp, #7]
1342		STR r4, [sp, #-4]!
1343		BL f_hasWon
1344		ADD sp, sp, #5
1345		MOV r4, r0
1346		STRB r4, [sp]
1347		LDRSB r4, [sp]
1348		CMP r4, #0
1349		BEQ L24
1350		LDRSB r4, [sp, #1]
1351		LDRSB r5, [sp, #26]
1352		CMP r4, r5
1353		MOVEQ r4, #1
1354		MOVNE r4, #0
1355		CMP r4, #0
1356		BEQ L26
1357		LDR r4, =100
1358		STR r4, [sp, #14]
1359		B L27
1360	L26:
1361		LDR r4, =-100
1362		STR r4, [sp, #14]
1363	L27:
1364		B L25
1365	L24:
1366		SUB sp, sp, #1
1367		LDR r4, [sp, #7]
1368		STR r4, [sp, #-4]!
1369		BL f_containEmptyCell
1370		ADD sp, sp, #4
1371		MOV r4, r0
1372		STRB r4, [sp]
1373		LDRSB r4, [sp]
1374		CMP r4, #0
1375		BEQ L28
1376		LDRSB r4, [sp, #2]
1377		STRB r4, [sp, #-1]!
1378		LDRSB r4, [sp, #28]
1379		STRB r4, [sp, #-1]!
1380		LDR r4, [sp, #5]
1381		STR r4, [sp, #-4]!
1382		BL f_calculateValuesFromNextStates
1383		ADD sp, sp, #6
1384		MOV r4, r0
1385		STR r4, [sp, #15]
1386		LDR r4, [sp, #15]
1387		LDR r5, =100
1388		CMP r4, r5
1389		MOVEQ r4, #1
1390		MOVNE r4, #0
1391		CMP r4, #0
1392		BEQ L30
1393		LDR r4, =90
1394		STR r4, [sp, #15]
1395		B L31
1396	L30:
1397	L31:
1398		B L29
1399	L28:
1400		LDR r4, =0
1401		STR r4, [sp, #15]
1402	L29:
1403		ADD sp, sp, #1
1404	L25:
1405		LDR r4, [sp, #14]
1406		LDR r5, [sp, #22]
1407		MOV r0, r5
1408		BL p_check_null_pointer
1409		LDR r5, [r5, #4]
1410		STR r4, [r5]
1411		ADD sp, sp, #14
1412	L21:
1413		LDR r4, [sp]
1414		MOV r0, r4
1415		ADD sp, sp, #4
1416		POP {pc}
1417		POP {pc}
1418		.ltorg
1419	f_calculateValuesFromNextStates:
1420		PUSH {lr}
1421		SUB sp, sp, #32
1422		LDR r4, [sp, #36]
1423		MOV r0, r4
1424		BL p_check_null_pointer
1425		LDR r4, [r4]
1426		LDR r4, [r4]
1427		STR r4, [sp, #28]
1428		LDR r4, [sp, #28]
1429		MOV r0, r4
1430		BL p_check_null_pointer
1431		LDR r4, [r4]
1432		LDR r4, [r4]
1433		STR r4, [sp, #24]
1434		LDR r4, [sp, #28]
1435		MOV r0, r4
1436		BL p_check_null_pointer
1437		LDR r4, [r4, #4]
1438		LDR r4, [r4]
1439		STR r4, [sp, #20]
1440		LDR r4, [sp, #36]
1441		MOV r0, r4
1442		BL p_check_null_pointer
1443		LDR r4, [r4, #4]
1444		LDR r4, [r4]
1445		STR r4, [sp, #16]
1446		LDRSB r4, [sp, #41]
1447		STRB r4, [sp, #-1]!
1448		LDRSB r4, [sp, #41]
1449		STRB r4, [sp, #-1]!
1450		LDR r4, [sp, #26]
1451		STR r4, [sp, #-4]!
1452		BL f_calculateValuesFromNextStatesRow
1453		ADD sp, sp, #6
1454		MOV r4, r0
1455		STR r4, [sp, #12]
1456		LDRSB r4, [sp, #41]
1457		STRB r4, [sp, #-1]!
1458		LDRSB r4, [sp, #41]
1459		STRB r4, [sp, #-1]!
1460		LDR r4, [sp, #22]
1461		STR r4, [sp, #-4]!
1462		BL f_calculateValuesFromNextStatesRow
1463		ADD sp, sp, #6
1464		MOV r4, r0
1465		STR r4, [sp, #8]
1466		LDRSB r4, [sp, #41]
1467		STRB r4, [sp, #-1]!
1468		LDRSB r4, [sp, #41]
1469		STRB r4, [sp, #-1]!
1470		LDR r4, [sp, #18]
1471		STR r4, [sp, #-4]!
1472		BL f_calculateValuesFromNextStatesRow
1473		ADD sp, sp, #6
1474		MOV r4, r0
1475		STR r4, [sp, #4]
1476		LDR r4, [sp, #4]
1477		STR r4, [sp, #-4]!
1478		LDR r4, [sp, #12]
1479		STR r4, [sp, #-4]!
1480		LDR r4, [sp, #20]
1481		STR r4, [sp, #-4]!
1482		LDRSB r4, [sp, #53]
1483		STRB r4, [sp, #-1]!
1484		LDRSB r4, [sp, #53]
1485		STRB r4, [sp, #-1]!
1486		BL f_combineValue
1487		ADD sp, sp, #14
1488		MOV r4, r0
1489		STR r4, [sp]
1490		LDR r4, [sp]
1491		MOV r0, r4
1492		ADD sp, sp, #32
1493		POP {pc}
1494		POP {pc}
1495		.ltorg
1496	f_calculateValuesFromNextStatesRow:
1497		PUSH {lr}
1498		SUB sp, sp, #32
1499		LDR r4, [sp, #36]
1500		MOV r0, r4
1501		BL p_check_null_pointer
1502		LDR r4, [r4]
1503		LDR r4, [r4]
1504		STR r4, [sp, #28]
1505		LDR r4, [sp, #28]
1506		MOV r0, r4
1507		BL p_check_null_pointer
1508		LDR r4, [r4]
1509		LDR r4, [r4]
1510		STR r4, [sp, #24]
1511		LDR r4, [sp, #28]
1512		MOV r0, r4
1513		BL p_check_null_pointer
1514		LDR r4, [r4, #4]
1515		LDR r4, [r4]
1516		STR r4, [sp, #20]
1517		LDR r4, [sp, #36]
1518		MOV r0, r4
1519		BL p_check_null_pointer
1520		LDR r4, [r4, #4]
1521		LDR r4, [r4]
1522		STR r4, [sp, #16]
1523		LDRSB r4, [sp, #41]
1524		STRB r4, [sp, #-1]!
1525		LDRSB r4, [sp, #41]
1526		STRB r4, [sp, #-1]!
1527		LDR r4, [sp, #26]
1528		STR r4, [sp, #-4]!
1529		BL f_setValuesForAllStates
1530		ADD sp, sp, #6
1531		MOV r4, r0
1532		STR r4, [sp, #12]
1533		LDRSB r4, [sp, #41]
1534		STRB r4, [sp, #-1]!
1535		LDRSB r4, [sp, #41]
1536		STRB r4, [sp, #-1]!
1537		LDR r4, [sp, #22]
1538		STR r4, [sp, #-4]!
1539		BL f_setValuesForAllStates
1540		ADD sp, sp, #6
1541		MOV r4, r0
1542		STR r4, [sp, #8]
1543		LDRSB r4, [sp, #41]
1544		STRB r4, [sp, #-1]!
1545		LDRSB r4, [sp, #41]
1546		STRB r4, [sp, #-1]!
1547		LDR r4, [sp, #18]
1548		STR r4, [sp, #-4]!
1549		BL f_setValuesForAllStates
1550		ADD sp, sp, #6
1551		MOV r4, r0
1552		STR r4, [sp, #4]
1553		LDR r4, [sp, #4]
1554		STR r4, [sp, #-4]!
1555		LDR r4, [sp, #12]
1556		STR r4, [sp, #-4]!
1557		LDR r4, [sp, #20]
1558		STR r4, [sp, #-4]!
1559		LDRSB r4, [sp, #53]
1560		STRB r4, [sp, #-1]!
1561		LDRSB r4, [sp, #53]
1562		STRB r4, [sp, #-1]!
1563		BL f_combineValue
1564		ADD sp, sp, #14
1565		MOV r4, r0
1566		STR r4, [sp]
1567		LDR r4, [sp]
1568		MOV r0, r4
1569		ADD sp, sp, #32
1570		POP {pc}
1571		POP {pc}
1572		.ltorg
1573	f_combineValue:
1574		PUSH {lr}
1575		SUB sp, sp, #4
1576		LDR r4, =0
1577		STR r4, [sp]
1578		LDRSB r4, [sp, #8]
1579		LDRSB r5, [sp, #9]
1580		CMP r4, r5
1581		MOVEQ r4, #1
1582		MOVNE r4, #0
1583		CMP r4, #0
1584		BEQ L32
1585		LDR r4, [sp, #18]
1586		STR r4, [sp, #-4]!
1587		LDR r4, [sp, #18]
1588		STR r4, [sp, #-4]!
1589		LDR r4, [sp, #18]
1590		STR r4, [sp, #-4]!
1591		BL f_min3
1592		ADD sp, sp, #12
1593		MOV r4, r0
1594		STR r4, [sp]
1595		B L33
1596	L32:
1597		LDR r4, [sp, #18]
1598		STR r4, [sp, #-4]!
1599		LDR r4, [sp, #18]
1600		STR r4, [sp, #-4]!
1601		LDR r4, [sp, #18]
1602		STR r4, [sp, #-4]!
1603		BL f_max3
1604		ADD sp, sp, #12
1605		MOV r4, r0
1606		STR r4, [sp]
1607	L33:
1608		LDR r4, [sp]
1609		MOV r0, r4
1610		ADD sp, sp, #4
1611		POP {pc}
1612		POP {pc}
1613		.ltorg
1614	f_min3:
1615		PUSH {lr}
1616		LDR r4, [sp, #4]
1617		LDR r5, [sp, #8]
1618		CMP r4, r5
1619		MOVLT r4, #1
1620		MOVGE r4, #0
1621		CMP r4, #0
1622		BEQ L34
1623		LDR r4, [sp, #4]
1624		LDR r5, [sp, #12]
1625		CMP r4, r5
1626		MOVLT r4, #1
1627		MOVGE r4, #0
1628		CMP r4, #0
1629		BEQ L36
1630		LDR r4, [sp, #4]
1631		MOV r0, r4
1632		POP {pc}
1633		B L37
1634	L36:
1635		LDR r4, [sp, #12]
1636		MOV r0, r4
1637		POP {pc}
1638	L37:
1639		B L35
1640	L34:
1641		LDR r4, [sp, #8]
1642		LDR r5, [sp, #12]
1643		CMP r4, r5
1644		MOVLT r4, #1
1645		MOVGE r4, #0
1646		CMP r4, #0
1647		BEQ L38
1648		LDR r4, [sp, #8]
1649		MOV r0, r4
1650		POP {pc}
1651		B L39
1652	L38:
1653		LDR r4, [sp, #12]
1654		MOV r0, r4
1655		POP {pc}
1656	L39:
1657	L35:
1658		POP {pc}
1659		.ltorg
1660	f_max3:
1661		PUSH {lr}
1662		LDR r4, [sp, #4]
1663		LDR r5, [sp, #8]
1664		CMP r4, r5
1665		MOVGT r4, #1
1666		MOVLE r4, #0
1667		CMP r4, #0
1668		BEQ L40
1669		LDR r4, [sp, #4]
1670		LDR r5, [sp, #12]
1671		CMP r4, r5
1672		MOVGT r4, #1
1673		MOVLE r4, #0
1674		CMP r4, #0
1675		BEQ L42
1676		LDR r4, [sp, #4]
1677		MOV r0, r4
1678		POP {pc}
1679		B L43
1680	L42:
1681		LDR r4, [sp, #12]
1682		MOV r0, r4
1683		POP {pc}
1684	L43:
1685		B L41
1686	L40:
1687		LDR r4, [sp, #8]
1688		LDR r5, [sp, #12]
1689		CMP r4, r5
1690		MOVGT r4, #1
1691		MOVLE r4, #0
1692		CMP r4, #0
1693		BEQ L44
1694		LDR r4, [sp, #8]
1695		MOV r0, r4
1696		POP {pc}
1697		B L45
1698	L44:
1699		LDR r4, [sp, #12]
1700		MOV r0, r4
1701		POP {pc}
1702	L45:
1703	L41:
1704		POP {pc}
1705		.ltorg
1706	f_destroyAI:
1707		PUSH {lr}
1708		SUB sp, sp, #9
1709		LDR r4, [sp, #13]
1710		MOV r0, r4
1711		BL p_check_null_pointer
1712		LDR r4, [r4]
1713		LDR r4, [r4]
1714		STR r4, [sp, #5]
1715		LDR r4, [sp, #13]
1716		MOV r0, r4
1717		BL p_check_null_pointer
1718		LDR r4, [r4, #4]
1719		LDR r4, [r4]
1720		STR r4, [sp, #1]
1721		LDR r4, [sp, #1]
1722		STR r4, [sp, #-4]!
1723		BL f_deleteStateTreeRecursively
1724		ADD sp, sp, #4
1725		MOV r4, r0
1726		STRB r4, [sp]
1727		LDR r4, [sp, #5]
1728		MOV r0, r4
1729		BL p_free_pair
1730		LDR r4, [sp, #13]
1731		MOV r0, r4
1732		BL p_free_pair
1733		MOV r4, #1
1734		MOV r0, r4
1735		ADD sp, sp, #9
1736		POP {pc}
1737		POP {pc}
1738		.ltorg
1739	f_askForAMoveAI:
1740		PUSH {lr}
1741		SUB sp, sp, #21
1742		LDR r4, [sp, #31]
1743		MOV r0, r4
1744		BL p_check_null_pointer
1745		LDR r4, [r4]
1746		LDR r4, [r4]
1747		STR r4, [sp, #17]
1748		LDR r4, [sp, #31]
1749		MOV r0, r4
1750		BL p_check_null_pointer
1751		LDR r4, [r4, #4]
1752		LDR r4, [r4]
1753		STR r4, [sp, #13]
1754		LDR r4, [sp, #13]
1755		MOV r0, r4
1756		BL p_check_null_pointer
1757		LDR r4, [r4]
1758		LDR r4, [r4]
1759		STR r4, [sp, #9]
1760		LDR r4, [sp, #9]
1761		MOV r0, r4
1762		BL p_check_null_pointer
1763		LDR r4, [r4, #4]
1764		LDR r4, [r4]
1765		STR r4, [sp, #5]
1766		LDR r4, [sp, #13]
1767		MOV r0, r4
1768		BL p_check_null_pointer
1769		LDR r4, [r4, #4]
1770		LDR r4, [r4]
1771		STR r4, [sp, #1]
1772		LDR r4, [sp, #35]
1773		STR r4, [sp, #-4]!
1774		LDR r4, [sp, #5]
1775		STR r4, [sp, #-4]!
1776		LDR r4, [sp, #13]
1777		STR r4, [sp, #-4]!
1778		BL f_findTheBestMove
1779		ADD sp, sp, #12
1780		MOV r4, r0
1781		STRB r4, [sp]
1782		LDR r4, =msg_30
1783		MOV r0, r4
1784		BL p_print_string
1785		BL p_print_ln
1786		ADD r4, sp, #35
1787		LDR r5, =1
1788		LDR r4, [r4]
1789		MOV r0, r5
1790		MOV r1, r4
1791		BL p_check_array_bounds
1792		ADD r4, r4, #4
1793		ADD r4, r4, r5, LSL #2
1794		LDR r4, [r4]
1795		STR r4, [sp, #-4]!
1796		ADD r4, sp, #39
1797		LDR r5, =0
1798		LDR r4, [r4]
1799		MOV r0, r5
1800		MOV r1, r4
1801		BL p_check_array_bounds
1802		ADD r4, r4, #4
1803		ADD r4, r4, r5, LSL #2
1804		LDR r4, [r4]
1805		STR r4, [sp, #-4]!
1806		LDR r4, [sp, #13]
1807		STR r4, [sp, #-4]!
1808		BL f_deleteAllOtherChildren
1809		ADD sp, sp, #12
1810		MOV r4, r0
1811		LDR r5, [sp, #31]
1812		MOV r0, r5
1813		BL p_check_null_pointer
1814		LDR r5, [r5, #4]
1815		STR r4, [r5]
1816		LDR r4, [sp, #13]
1817		STR r4, [sp, #-4]!
1818		BL f_deleteThisStateOnly
1819		ADD sp, sp, #4
1820		MOV r4, r0
1821		STRB r4, [sp]
1822		MOV r4, #1
1823		MOV r0, r4
1824		ADD sp, sp, #21
1825		POP {pc}
1826		POP {pc}
1827		.ltorg
1828	f_findTheBestMove:
1829		PUSH {lr}
1830		SUB sp, sp, #1
1831		LDR r4, [sp, #9]
1832		LDR r5, =90
1833		CMP r4, r5
1834		MOVEQ r4, #1
1835		MOVNE r4, #0
1836		CMP r4, #0
1837		BEQ L46
1838		SUB sp, sp, #1
1839		LDR r4, [sp, #14]
1840		STR r4, [sp, #-4]!
1841		LDR r4, =100
1842		STR r4, [sp, #-4]!
1843		LDR r4, [sp, #14]
1844		STR r4, [sp, #-4]!
1845		BL f_findMoveWithGivenValue
1846		ADD sp, sp, #12
1847		MOV r4, r0
1848		STRB r4, [sp]
1849		LDRSB r4, [sp]
1850		CMP r4, #0
1851		BEQ L48
1852		MOV r4, #1
1853		MOV r0, r4
1854		ADD sp, sp, #2
1855		POP {pc}
1856		B L49
1857	L48:
1858	L49:
1859		ADD sp, sp, #1
1860		B L47
1861	L46:
1862	L47:
1863		LDR r4, [sp, #13]
1864		STR r4, [sp, #-4]!
1865		LDR r4, [sp, #13]
1866		STR r4, [sp, #-4]!
1867		LDR r4, [sp, #13]
1868		STR r4, [sp, #-4]!
1869		BL f_findMoveWithGivenValue
1870		ADD sp, sp, #12
1871		MOV r4, r0
1872		STRB r4, [sp]
1873		LDRSB r4, [sp]
1874		CMP r4, #0
1875		BEQ L50
1876		MOV r4, #1
1877		MOV r0, r4
1878		ADD sp, sp, #1
1879		POP {pc}
1880		B L51
1881	L50:
1882		LDR r4, =msg_31
1883		MOV r0, r4
1884		BL p_print_string
1885		BL p_print_ln
1886		LDR r4, =-1
1887		MOV r0, r4
1888		BL exit
1889	L51:
1890		POP {pc}
1891		.ltorg
1892	f_findMoveWithGivenValue:
1893		PUSH {lr}
1894		SUB sp, sp, #17
1895		LDR r4, [sp, #21]
1896		MOV r0, r4
1897		BL p_check_null_pointer
1898		LDR r4, [r4]
1899		LDR r4, [r4]
1900		STR r4, [sp, #13]
1901		LDR r4, [sp, #13]
1902		MOV r0, r4
1903		BL p_check_null_pointer
1904		LDR r4, [r4]
1905		LDR r4, [r4]
1906		STR r4, [sp, #9]
1907		LDR r4, [sp, #13]
1908		MOV r0, r4
1909		BL p_check_null_pointer
1910		LDR r4, [r4, #4]
1911		LDR r4, [r4]
1912		STR r4, [sp, #5]
1913		LDR r4, [sp, #21]
1914		MOV r0, r4
1915		BL p_check_null_pointer
1916		LDR r4, [r4, #4]
1917		LDR r4, [r4]
1918		STR r4, [sp, #1]
1919		LDR r4, [sp, #29]
1920		STR r4, [sp, #-4]!
1921		LDR r4, [sp, #29]
1922		STR r4, [sp, #-4]!
1923		LDR r4, [sp, #17]
1924		STR r4, [sp, #-4]!
1925		BL f_findMoveWithGivenValueRow
1926		ADD sp, sp, #12
1927		MOV r4, r0
1928		STRB r4, [sp]
1929		LDRSB r4, [sp]
1930		CMP r4, #0
1931		BEQ L52
1932		LDR r4, =1
1933		ADD r5, sp, #29
1934		LDR r6, =0
1935		LDR r5, [r5]
1936		MOV r0, r6
1937		MOV r1, r5
1938		BL p_check_array_bounds
1939		ADD r5, r5, #4
1940		ADD r5, r5, r6, LSL #2
1941		STR r4, [r5]
1942		B L53
1943	L52:
1944		LDR r4, [sp, #29]
1945		STR r4, [sp, #-4]!
1946		LDR r4, [sp, #29]
1947		STR r4, [sp, #-4]!
1948		LDR r4, [sp, #13]
1949		STR r4, [sp, #-4]!
1950		BL f_findMoveWithGivenValueRow
1951		ADD sp, sp, #12
1952		MOV r4, r0
1953		STRB r4, [sp]
1954		LDRSB r4, [sp]
1955		CMP r4, #0
1956		BEQ L54
1957		LDR r4, =2
1958		ADD r6, sp, #29
1959		LDR r7, =0
1960		LDR r6, [r6]
1961		MOV r0, r7
1962		MOV r1, r6
1963		BL p_check_array_bounds
1964		ADD r6, r6, #4
1965		ADD r6, r6, r7, LSL #2
1966		STR r4, [r6]
1967		B L55
1968	L54:
1969		LDR r4, [sp, #29]
1970		STR r4, [sp, #-4]!
1971		LDR r4, [sp, #29]
1972		STR r4, [sp, #-4]!
1973		LDR r4, [sp, #9]
1974		STR r4, [sp, #-4]!
1975		BL f_findMoveWithGivenValueRow
1976		ADD sp, sp, #12
1977		MOV r4, r0
1978		STRB r4, [sp]
1979		LDRSB r4, [sp]
1980		CMP r4, #0
1981		BEQ L56
1982		LDR r4, =3
1983		ADD r7, sp, #29
1984		LDR r8, =0
1985		LDR r7, [r7]
1986		MOV r0, r8
1987		MOV r1, r7
1988		BL p_check_array_bounds
1989		ADD r7, r7, #4
1990		ADD r7, r7, r8, LSL #2
1991		STR r4, [r7]
1992		B L57
1993	L56:
1994		MOV r4, #0
1995		MOV r0, r4
1996		ADD sp, sp, #17
1997		POP {pc}
1998	L57:
1999	L55:
2000	L53:
2001		MOV r4, #1
2002		MOV r0, r4
2003		ADD sp, sp, #17
2004		POP {pc}
2005		POP {pc}
2006		.ltorg
2007	f_findMoveWithGivenValueRow:
2008		PUSH {lr}
2009		SUB sp, sp, #17
2010		LDR r4, [sp, #21]
2011		MOV r0, r4
2012		BL p_check_null_pointer
2013		LDR r4, [r4]
2014		LDR r4, [r4]
2015		STR r4, [sp, #13]
2016		LDR r4, [sp, #13]
2017		MOV r0, r4
2018		BL p_check_null_pointer
2019		LDR r4, [r4]
2020		LDR r4, [r4]
2021		STR r4, [sp, #9]
2022		LDR r4, [sp, #13]
2023		MOV r0, r4
2024		BL p_check_null_pointer
2025		LDR r4, [r4, #4]
2026		LDR r4, [r4]
2027		STR r4, [sp, #5]
2028		LDR r4, [sp, #21]
2029		MOV r0, r4
2030		BL p_check_null_pointer
2031		LDR r4, [r4, #4]
2032		LDR r4, [r4]
2033		STR r4, [sp, #1]
2034		LDR r4, [sp, #25]
2035		STR r4, [sp, #-4]!
2036		LDR r4, [sp, #13]
2037		STR r4, [sp, #-4]!
2038		BL f_hasGivenStateValue
2039		ADD sp, sp, #8
2040		MOV r4, r0
2041		STRB r4, [sp]
2042		LDRSB r4, [sp]
2043		CMP r4, #0
2044		BEQ L58
2045		LDR r4, =1
2046		ADD r5, sp, #29
2047		LDR r6, =1
2048		LDR r5, [r5]
2049		MOV r0, r6
2050		MOV r1, r5
2051		BL p_check_array_bounds
2052		ADD r5, r5, #4
2053		ADD r5, r5, r6, LSL #2
2054		STR r4, [r5]
2055		B L59
2056	L58:
2057		LDR r4, [sp, #25]
2058		STR r4, [sp, #-4]!
2059		LDR r4, [sp, #9]
2060		STR r4, [sp, #-4]!
2061		BL f_hasGivenStateValue
2062		ADD sp, sp, #8
2063		MOV r4, r0
2064		STRB r4, [sp]
2065		LDRSB r4, [sp]
2066		CMP r4, #0
2067		BEQ L60
2068		LDR r4, =2
2069		ADD r6, sp, #29
2070		LDR r7, =1
2071		LDR r6, [r6]
2072		MOV r0, r7
2073		MOV r1, r6
2074		BL p_check_array_bounds
2075		ADD r6, r6, #4
2076		ADD r6, r6, r7, LSL #2
2077		STR r4, [r6]
2078		B L61
2079	L60:
2080		LDR r4, [sp, #25]
2081		STR r4, [sp, #-4]!
2082		LDR r4, [sp, #5]
2083		STR r4, [sp, #-4]!
2084		BL f_hasGivenStateValue
2085		ADD sp, sp, #8
2086		MOV r4, r0
2087		STRB r4, [sp]
2088		LDRSB r4, [sp]
2089		CMP r4, #0
2090		BEQ L62
2091		LDR r4, =3
2092		ADD r7, sp, #29
2093		LDR r8, =1
2094		LDR r7, [r7]
2095		MOV r0, r8
2096		MOV r1, r7
2097		BL p_check_array_bounds
2098		ADD r7, r7, #4
2099		ADD r7, r7, r8, LSL #2
2100		STR r4, [r7]
2101		B L63
2102	L62:
2103		MOV r4, #0
2104		MOV r0, r4
2105		ADD sp, sp, #17
2106		POP {pc}
2107	L63:
2108	L61:
2109	L59:
2110		MOV r4, #1
2111		MOV r0, r4
2112		ADD sp, sp, #17
2113		POP {pc}
2114		POP {pc}
2115		.ltorg
2116	f_hasGivenStateValue:
2117		PUSH {lr}
2118		LDR r4, [sp, #4]
2119		LDR r5, =0
2120		CMP r4, r5
2121		MOVEQ r4, #1
2122		MOVNE r4, #0
2123		CMP r4, #0
2124		BEQ L64
2125		MOV r4, #0
2126		MOV r0, r4
2127		POP {pc}
2128		B L65
2129	L64:
2130		SUB sp, sp, #4
2131		LDR r4, [sp, #8]
2132		MOV r0, r4
2133		BL p_check_null_pointer
2134		LDR r4, [r4, #4]
2135		LDR r4, [r4]
2136		STR r4, [sp]
2137		LDR r4, [sp]
2138		LDR r5, [sp, #12]
2139		CMP r4, r5
2140		MOVEQ r4, #1
2141		MOVNE r4, #0
2142		MOV r0, r4
2143		ADD sp, sp, #4
2144		POP {pc}
2145		ADD sp, sp, #4
2146	L65:
2147		POP {pc}
2148		.ltorg
2149	f_notifyMoveAI:
2150		PUSH {lr}
2151		SUB sp, sp, #13
2152		LDR r4, [sp, #23]
2153		MOV r0, r4
2154		BL p_check_null_pointer
2155		LDR r4, [r4, #4]
2156		LDR r4, [r4]
2157		STR r4, [sp, #9]
2158		LDR r4, [sp, #9]
2159		MOV r0, r4
2160		BL p_check_null_pointer
2161		LDR r4, [r4]
2162		LDR r4, [r4]
2163		STR r4, [sp, #5]
2164		LDR r4, [sp, #5]
2165		MOV r0, r4
2166		BL p_check_null_pointer
2167		LDR r4, [r4, #4]
2168		LDR r4, [r4]
2169		STR r4, [sp, #1]
2170		LDR r4, =msg_32
2171		MOV r0, r4
2172		BL p_print_string
2173		BL p_print_ln
2174		LDR r4, [sp, #31]
2175		STR r4, [sp, #-4]!
2176		LDR r4, [sp, #31]
2177		STR r4, [sp, #-4]!
2178		LDR r4, [sp, #9]
2179		STR r4, [sp, #-4]!
2180		BL f_deleteAllOtherChildren
2181		ADD sp, sp, #12
2182		MOV r4, r0
2183		LDR r5, [sp, #23]
2184		MOV r0, r5
2185		BL p_check_null_pointer
2186		LDR r5, [r5, #4]
2187		STR r4, [r5]
2188		LDR r4, [sp, #9]
2189		STR r4, [sp, #-4]!
2190		BL f_deleteThisStateOnly
2191		ADD sp, sp, #4
2192		MOV r4, r0
2193		STRB r4, [sp]
2194		MOV r4, #1
2195		MOV r0, r4
2196		ADD sp, sp, #13
2197		POP {pc}
2198		POP {pc}
2199		.ltorg
2200	f_deleteAllOtherChildren:
2201		PUSH {lr}
2202		SUB sp, sp, #33
2203		LDR r4, [sp, #37]
2204		MOV r0, r4
2205		BL p_check_null_pointer
2206		LDR r4, [r4]
2207		LDR r4, [r4]
2208		STR r4, [sp, #29]
2209		LDR r4, [sp, #29]
2210		MOV r0, r4
2211		BL p_check_null_pointer
2212		LDR r4, [r4]
2213		LDR r4, [r4]
2214		STR r4, [sp, #25]
2215		LDR r4, [sp, #29]
2216		MOV r0, r4
2217		BL p_check_null_pointer
2218		LDR r4, [r4, #4]
2219		LDR r4, [r4]
2220		STR r4, [sp, #21]
2221		LDR r4, [sp, #37]
2222		MOV r0, r4
2223		BL p_check_null_pointer
2224		LDR r4, [r4, #4]
2225		LDR r4, [r4]
2226		STR r4, [sp, #17]
2227		LDR r4, =0
2228		STR r4, [sp, #13]
2229		LDR r4, =0
2230		STR r4, [sp, #9]
2231		LDR r4, =0
2232		STR r4, [sp, #5]
2233		LDR r4, [sp, #41]
2234		LDR r5, =1
2235		CMP r4, r5
2236		MOVEQ r4, #1
2237		MOVNE r4, #0
2238		CMP r4, #0
2239		BEQ L66
2240		LDR r4, [sp, #25]
2241		STR r4, [sp, #13]
2242		LDR r4, [sp, #21]
2243		STR r4, [sp, #9]
2244		LDR r4, [sp, #17]
2245		STR r4, [sp, #5]
2246		B L67
2247	L66:
2248		LDR r4, [sp, #25]
2249		STR r4, [sp, #9]
2250		LDR r4, [sp, #41]
2251		LDR r5, =2
2252		CMP r4, r5
2253		MOVEQ r4, #1
2254		MOVNE r4, #0
2255		CMP r4, #0
2256		BEQ L68
2257		LDR r4, [sp, #21]
2258		STR r4, [sp, #13]
2259		LDR r4, [sp, #17]
2260		STR r4, [sp, #5]
2261		B L69
2262	L68:
2263		LDR r4, [sp, #17]
2264		STR r4, [sp, #13]
2265		LDR r4, [sp, #21]
2266		STR r4, [sp, #5]
2267	L69:
2268	L67:
2269		LDR r4, [sp, #45]
2270		STR r4, [sp, #-4]!
2271		LDR r4, [sp, #17]
2272		STR r4, [sp, #-4]!
2273		BL f_deleteAllOtherChildrenRow
2274		ADD sp, sp, #8
2275		MOV r4, r0
2276		STR r4, [sp, #1]
2277		LDR r4, [sp, #9]
2278		STR r4, [sp, #-4]!
2279		BL f_deleteChildrenStateRecursivelyRow
2280		ADD sp, sp, #4
2281		MOV r4, r0
2282		STRB r4, [sp]
2283		LDR r4, [sp, #5]
2284		STR r4, [sp, #-4]!
2285		BL f_deleteChildrenStateRecursivelyRow
2286		ADD sp, sp, #4
2287		MOV r4, r0
2288		STRB r4, [sp]
2289		LDR r4, [sp, #1]
2290		MOV r0, r4
2291		ADD sp, sp, #33
2292		POP {pc}
2293		POP {pc}
2294		.ltorg
2295	f_deleteAllOtherChildrenRow:
2296		PUSH {lr}
2297		SUB sp, sp, #29
2298		LDR r4, [sp, #33]
2299		MOV r0, r4
2300		BL p_check_null_pointer
2301		LDR r4, [r4]
2302		LDR r4, [r4]
2303		STR r4, [sp, #25]
2304		LDR r4, [sp, #25]
2305		MOV r0, r4
2306		BL p_check_null_pointer
2307		LDR r4, [r4]
2308		LDR r4, [r4]
2309		STR r4, [sp, #21]
2310		LDR r4, [sp, #25]
2311		MOV r0, r4
2312		BL p_check_null_pointer
2313		LDR r4, [r4, #4]
2314		LDR r4, [r4]
2315		STR r4, [sp, #17]
2316		LDR r4, [sp, #33]
2317		MOV r0, r4
2318		BL p_check_null_pointer
2319		LDR r4, [r4, #4]
2320		LDR r4, [r4]
2321		STR r4, [sp, #13]
2322		LDR r4, =0
2323		STR r4, [sp, #9]
2324		LDR r4, =0
2325		STR r4, [sp, #5]
2326		LDR r4, =0
2327		STR r4, [sp, #1]
2328		LDR r4, [sp, #37]
2329		LDR r5, =1
2330		CMP r4, r5
2331		MOVEQ r4, #1
2332		MOVNE r4, #0
2333		CMP r4, #0
2334		BEQ L70
2335		LDR r4, [sp, #21]
2336		STR r4, [sp, #9]
2337		LDR r4, [sp, #17]
2338		STR r4, [sp, #5]
2339		LDR r4, [sp, #13]
2340		STR r4, [sp, #1]
2341		B L71
2342	L70:
2343		LDR r4, [sp, #21]
2344		STR r4, [sp, #5]
2345		LDR r4, [sp, #37]
2346		LDR r5, =2
2347		CMP r4, r5
2348		MOVEQ r4, #1
2349		MOVNE r4, #0
2350		CMP r4, #0
2351		BEQ L72
2352		LDR r4, [sp, #17]
2353		STR r4, [sp, #9]
2354		LDR r4, [sp, #13]
2355		STR r4, [sp, #1]
2356		B L73
2357	L72:
2358		LDR r4, [sp, #13]
2359		STR r4, [sp, #9]
2360		LDR r4, [sp, #17]
2361		STR r4, [sp, #1]
2362	L73:
2363	L71:
2364		LDR r4, [sp, #5]
2365		STR r4, [sp, #-4]!
2366		BL f_deleteStateTreeRecursively
2367		ADD sp, sp, #4
2368		MOV r4, r0
2369		STRB r4, [sp]
2370		LDR r4, [sp, #1]
2371		STR r4, [sp, #-4]!
2372		BL f_deleteStateTreeRecursively
2373		ADD sp, sp, #4
2374		MOV r4, r0
2375		STRB r4, [sp]
2376		LDR r4, [sp, #9]
2377		MOV r0, r4
2378		ADD sp, sp, #29
2379		POP {pc}
2380		POP {pc}
2381		.ltorg
2382	f_deleteStateTreeRecursively:
2383		PUSH {lr}
2384		LDR r4, [sp, #4]
2385		LDR r5, =0
2386		CMP r4, r5
2387		MOVEQ r4, #1
2388		MOVNE r4, #0
2389		CMP r4, #0
2390		BEQ L74
2391		MOV r4, #1
2392		MOV r0, r4
2393		POP {pc}
2394		B L75
2395	L74:
2396		SUB sp, sp, #13
2397		LDR r4, [sp, #17]
2398		MOV r0, r4
2399		BL p_check_null_pointer
2400		LDR r4, [r4]
2401		LDR r4, [r4]
2402		STR r4, [sp, #9]
2403		LDR r4, [sp, #9]
2404		MOV r0, r4
2405		BL p_check_null_pointer
2406		LDR r4, [r4]
2407		LDR r4, [r4]
2408		STR r4, [sp, #5]
2409		LDR r4, [sp, #9]
2410		MOV r0, r4
2411		BL p_check_null_pointer
2412		LDR r4, [r4, #4]
2413		LDR r4, [r4]
2414		STR r4, [sp, #1]
2415		LDR r4, [sp, #1]
2416		STR r4, [sp, #-4]!
2417		BL f_deleteChildrenStateRecursively
2418		ADD sp, sp, #4
2419		MOV r4, r0
2420		STRB r4, [sp]
2421		LDR r4, [sp, #17]
2422		STR r4, [sp, #-4]!
2423		BL f_deleteThisStateOnly
2424		ADD sp, sp, #4
2425		MOV r4, r0
2426		STRB r4, [sp]
2427		MOV r4, #1
2428		MOV r0, r4
2429		ADD sp, sp, #13
2430		POP {pc}
2431		ADD sp, sp, #13
2432	L75:
2433		POP {pc}
2434		.ltorg
2435	f_deleteThisStateOnly:
2436		PUSH {lr}
2437		SUB sp, sp, #13
2438		LDR r4, [sp, #17]
2439		MOV r0, r4
2440		BL p_check_null_pointer
2441		LDR r4, [r4]
2442		LDR r4, [r4]
2443		STR r4, [sp, #9]
2444		LDR r4, [sp, #9]
2445		MOV r0, r4
2446		BL p_check_null_pointer
2447		LDR r4, [r4]
2448		LDR r4, [r4]
2449		STR r4, [sp, #5]
2450		LDR r4, [sp, #9]
2451		MOV r0, r4
2452		BL p_check_null_pointer
2453		LDR r4, [r4, #4]
2454		LDR r4, [r4]
2455		STR r4, [sp, #1]
2456		LDR r4, [sp, #5]
2457		STR r4, [sp, #-4]!
2458		BL f_freeBoard
2459		ADD sp, sp, #4
2460		MOV r4, r0
2461		STRB r4, [sp]
2462		LDR r4, [sp, #1]
2463		STR r4, [sp, #-4]!
2464		BL f_freePointers
2465		ADD sp, sp, #4
2466		MOV r4, r0
2467		STRB r4, [sp]
2468		LDR r4, [sp, #9]
2469		MOV r0, r4
2470		BL p_free_pair
2471		LDR r4, [sp, #17]
2472		MOV r0, r4
2473		BL p_free_pair
2474		MOV r4, #1
2475		MOV r0, r4
2476		ADD sp, sp, #13
2477		POP {pc}
2478		POP {pc}
2479		.ltorg
2480	f_freePointers:
2481		PUSH {lr}
2482		SUB sp, sp, #17
2483		LDR r4, [sp, #21]
2484		MOV r0, r4
2485		BL p_check_null_pointer
2486		LDR r4, [r4]
2487		LDR r4, [r4]
2488		STR r4, [sp, #13]
2489		LDR r4, [sp, #13]
2490		MOV r0, r4
2491		BL p_check_null_pointer
2492		LDR r4, [r4]
2493		LDR r4, [r4]
2494		STR r4, [sp, #9]
2495		LDR r4, [sp, #13]
2496		MOV r0, r4
2497		BL p_check_null_pointer
2498		LDR r4, [r4, #4]
2499		LDR r4, [r4]
2500		STR r4, [sp, #5]
2501		LDR r4, [sp, #21]
2502		MOV r0, r4
2503		BL p_check_null_pointer
2504		LDR r4, [r4, #4]
2505		LDR r4, [r4]
2506		STR r4, [sp, #1]
2507		LDR r4, [sp, #9]
2508		STR r4, [sp, #-4]!
2509		BL f_freePointersRow
2510		ADD sp, sp, #4
2511		MOV r4, r0
2512		STRB r4, [sp]
2513		LDR r4, [sp, #5]
2514		STR r4, [sp, #-4]!
2515		BL f_freePointersRow
2516		ADD sp, sp, #4
2517		MOV r4, r0
2518		STRB r4, [sp]
2519		LDR r4, [sp, #1]
2520		STR r4, [sp, #-4]!
2521		BL f_freePointersRow
2522		ADD sp, sp, #4
2523		MOV r4, r0
2524		STRB r4, [sp]
2525		LDR r4, [sp, #13]
2526		MOV r0, r4
2527		BL p_free_pair
2528		LDR r4, [sp, #21]
2529		MOV r0, r4
2530		BL p_free_pair
2531		MOV r4, #1
2532		MOV r0, r4
2533		ADD sp, sp, #17
2534		POP {pc}
2535		POP {pc}
2536		.ltorg
2537	f_freePointersRow:
2538		PUSH {lr}
2539		SUB sp, sp, #4
2540		LDR r4, [sp, #8]
2541		MOV r0, r4
2542		BL p_check_null_pointer
2543		LDR r4, [r4]
2544		LDR r4, [r4]
2545		STR r4, [sp]
2546		LDR r4, [sp]
2547		MOV r0, r4
2548		BL p_free_pair
2549		LDR r4, [sp, #8]
2550		MOV r0, r4
2551		BL p_free_pair
2552		MOV r4, #1
2553		MOV r0, r4
2554		ADD sp, sp, #4
2555		POP {pc}
2556		POP {pc}
2557		.ltorg
2558	f_deleteChildrenStateRecursively:
2559		PUSH {lr}
2560		SUB sp, sp, #17
2561		LDR r4, [sp, #21]
2562		MOV r0, r4
2563		BL p_check_null_pointer
2564		LDR r4, [r4]
2565		LDR r4, [r4]
2566		STR r4, [sp, #13]
2567		LDR r4, [sp, #13]
2568		MOV r0, r4
2569		BL p_check_null_pointer
2570		LDR r4, [r4]
2571		LDR r4, [r4]
2572		STR r4, [sp, #9]
2573		LDR r4, [sp, #13]
2574		MOV r0, r4
2575		BL p_check_null_pointer
2576		LDR r4, [r4, #4]
2577		LDR r4, [r4]
2578		STR r4, [sp, #5]
2579		LDR r4, [sp, #21]
2580		MOV r0, r4
2581		BL p_check_null_pointer
2582		LDR r4, [r4, #4]
2583		LDR r4, [r4]
2584		STR r4, [sp, #1]
2585		LDR r4, [sp, #9]
2586		STR r4, [sp, #-4]!
2587		BL f_deleteChildrenStateRecursivelyRow
2588		ADD sp, sp, #4
2589		MOV r4, r0
2590		STRB r4, [sp]
2591		LDR r4, [sp, #5]
2592		STR r4, [sp, #-4]!
2593		BL f_deleteChildrenStateRecursivelyRow
2594		ADD sp, sp, #4
2595		MOV r4, r0
2596		STRB r4, [sp]
2597		LDR r4, [sp, #1]
2598		STR r4, [sp, #-4]!
2599		BL f_deleteChildrenStateRecursivelyRow
2600		ADD sp, sp, #4
2601		MOV r4, r0
2602		STRB r4, [sp]
2603		MOV r4, #1
2604		MOV r0, r4
2605		ADD sp, sp, #17
2606		POP {pc}
2607		POP {pc}
2608		.ltorg
2609	f_deleteChildrenStateRecursivelyRow:
2610		PUSH {lr}
2611		SUB sp, sp, #17
2612		LDR r4, [sp, #21]
2613		MOV r0, r4
2614		BL p_check_null_pointer
2615		LDR r4, [r4]
2616		LDR r4, [r4]
2617		STR r4, [sp, #13]
2618		LDR r4, [sp, #13]
2619		MOV r0, r4
2620		BL p_check_null_pointer
2621		LDR r4, [r4]
2622		LDR r4, [r4]
2623		STR r4, [sp, #9]
2624		LDR r4, [sp, #13]
2625		MOV r0, r4
2626		BL p_check_null_pointer
2627		LDR r4, [r4, #4]
2628		LDR r4, [r4]
2629		STR r4, [sp, #5]
2630		LDR r4, [sp, #21]
2631		MOV r0, r4
2632		BL p_check_null_pointer
2633		LDR r4, [r4, #4]
2634		LDR r4, [r4]
2635		STR r4, [sp, #1]
2636		LDR r4, [sp, #9]
2637		STR r4, [sp, #-4]!
2638		BL f_deleteStateTreeRecursively
2639		ADD sp, sp, #4
2640		MOV r4, r0
2641		STRB r4, [sp]
2642		LDR r4, [sp, #5]
2643		STR r4, [sp, #-4]!
2644		BL f_deleteStateTreeRecursively
2645		ADD sp, sp, #4
2646		MOV r4, r0
2647		STRB r4, [sp]
2648		LDR r4, [sp, #1]
2649		STR r4, [sp, #-4]!
2650		BL f_deleteStateTreeRecursively
2651		ADD sp, sp, #4
2652		MOV r4, r0
2653		STRB r4, [sp]
2654		MOV r4, #1
2655		MOV r0, r4
2656		ADD sp, sp, #17
2657		POP {pc}
2658		POP {pc}
2659		.ltorg
2660	f_askForAMove:
2661		PUSH {lr}
2662		LDRSB r4, [sp, #8]
2663		LDRSB r5, [sp, #9]
2664		CMP r4, r5
2665		MOVEQ r4, #1
2666		MOVNE r4, #0
2667		CMP r4, #0
2668		BEQ L76
2669		SUB sp, sp, #1
2670		LDR r4, [sp, #15]
2671		STR r4, [sp, #-4]!
2672		LDR r4, [sp, #9]
2673		STR r4, [sp, #-4]!
2674		BL f_askForAMoveHuman
2675		ADD sp, sp, #8
2676		MOV r4, r0
2677		STRB r4, [sp]
2678		ADD sp, sp, #1
2679		B L77
2680	L76:
2681		SUB sp, sp, #1
2682		LDR r4, [sp, #15]
2683		STR r4, [sp, #-4]!
2684		LDR r4, [sp, #15]
2685		STR r4, [sp, #-4]!
2686		LDRSB r4, [sp, #18]
2687		STRB r4, [sp, #-1]!
2688		LDRSB r4, [sp, #18]
2689		STRB r4, [sp, #-1]!
2690		LDR r4, [sp, #15]
2691		STR r4, [sp, #-4]!
2692		BL f_askForAMoveAI
2693		ADD sp, sp, #14
2694		MOV r4, r0
2695		STRB r4, [sp]
2696		ADD sp, sp, #1
2697	L77:
2698		MOV r4, #1
2699		MOV r0, r4
2700		POP {pc}
2701		POP {pc}
2702		.ltorg
2703	f_placeMove:
2704		PUSH {lr}
2705		SUB sp, sp, #4
2706		LDR r4, =0
2707		STR r4, [sp]
2708		LDR r4, [sp, #13]
2709		LDR r5, =2
2710		CMP r4, r5
2711		MOVLE r4, #1
2712		MOVGT r4, #0
2713		CMP r4, #0
2714		BEQ L78
2715		SUB sp, sp, #4
2716		LDR r4, [sp, #12]
2717		MOV r0, r4
2718		BL p_check_null_pointer
2719		LDR r4, [r4]
2720		LDR r4, [r4]
2721		STR r4, [sp]
2722		LDR r4, [sp, #17]
2723		LDR r5, =1
2724		CMP r4, r5
2725		MOVEQ r4, #1
2726		MOVNE r4, #0
2727		CMP r4, #0
2728		BEQ L80
2729		LDR r4, [sp]
2730		MOV r0, r4
2731		BL p_check_null_pointer
2732		LDR r4, [r4]
2733		LDR r4, [r4]
2734		STR r4, [sp, #4]
2735		B L81
2736	L80:
2737		LDR r4, [sp]
2738		MOV r0, r4
2739		BL p_check_null_pointer
2740		LDR r4, [r4, #4]
2741		LDR r4, [r4]
2742		STR r4, [sp, #4]
2743	L81:
2744		ADD sp, sp, #4
2745		B L79
2746	L78:
2747		LDR r4, [sp, #8]
2748		MOV r0, r4
2749		BL p_check_null_pointer
2750		LDR r4, [r4, #4]
2751		LDR r4, [r4]
2752		STR r4, [sp]
2753	L79:
2754		LDR r4, [sp, #17]
2755		LDR r5, =2
2756		CMP r4, r5
2757		MOVLE r4, #1
2758		MOVGT r4, #0
2759		CMP r4, #0
2760		BEQ L82
2761		SUB sp, sp, #4
2762		LDR r4, [sp, #4]
2763		MOV r0, r4
2764		BL p_check_null_pointer
2765		LDR r4, [r4]
2766		LDR r4, [r4]
2767		STR r4, [sp]
2768		LDR r4, [sp, #21]
2769		LDR r5, =1
2770		CMP r4, r5
2771		MOVEQ r4, #1
2772		MOVNE r4, #0
2773		CMP r4, #0
2774		BEQ L84
2775		LDRSB r4, [sp, #16]
2776		LDR r5, [sp]
2777		MOV r0, r5
2778		BL p_check_null_pointer
2779		LDR r5, [r5]
2780		STRB r4, [r5]
2781		B L85
2782	L84:
2783		LDRSB r4, [sp, #16]
2784		LDR r5, [sp]
2785		MOV r0, r5
2786		BL p_check_null_pointer
2787		LDR r5, [r5, #4]
2788		STRB r4, [r5]
2789	L85:
2790		ADD sp, sp, #4
2791		B L83
2792	L82:
2793		LDRSB r4, [sp, #12]
2794		LDR r5, [sp]
2795		MOV r0, r5
2796		BL p_check_null_pointer
2797		LDR r5, [r5, #4]
2798		STRB r4, [r5]
2799	L83:
2800		MOV r4, #1
2801		MOV r0, r4
2802		ADD sp, sp, #4
2803		POP {pc}
2804		POP {pc}
2805		.ltorg
2806	f_notifyMove:
2807		PUSH {lr}
2808		LDRSB r4, [sp, #8]
2809		LDRSB r5, [sp, #9]
2810		CMP r4, r5
2811		MOVEQ r4, #1
2812		MOVNE r4, #0
2813		CMP r4, #0
2814		BEQ L86
2815		SUB sp, sp, #1
2816		LDR r4, [sp, #19]
2817		STR r4, [sp, #-4]!
2818		LDR r4, [sp, #19]
2819		STR r4, [sp, #-4]!
2820		LDR r4, [sp, #19]
2821		STR r4, [sp, #-4]!
2822		LDRSB r4, [sp, #22]
2823		STRB r4, [sp, #-1]!
2824		LDRSB r4, [sp, #22]
2825		STRB r4, [sp, #-1]!
2826		LDR r4, [sp, #19]
2827		STR r4, [sp, #-4]!
2828		BL f_notifyMoveAI
2829		ADD sp, sp, #18
2830		MOV r4, r0
2831		STRB r4, [sp]
2832		ADD sp, sp, #1
2833		B L87
2834	L86:
2835		SUB sp, sp, #1
2836		LDR r4, [sp, #19]
2837		STR r4, [sp, #-4]!
2838		LDR r4, [sp, #19]
2839		STR r4, [sp, #-4]!
2840		LDRSB r4, [sp, #18]
2841		STRB r4, [sp, #-1]!
2842		LDRSB r4, [sp, #18]
2843		STRB r4, [sp, #-1]!
2844		LDR r4, [sp, #15]
2845		STR r4, [sp, #-4]!
2846		BL f_notifyMoveHuman
2847		ADD sp, sp, #14
2848		MOV r4, r0
2849		STRB r4, [sp]
2850		ADD sp, sp, #1
2851	L87:
2852		MOV r4, #1
2853		MOV r0, r4
2854		POP {pc}
2855		POP {pc}
2856		.ltorg
2857	f_oppositeSymbol:
2858		PUSH {lr}
2859		LDRSB r4, [sp, #4]
2860		MOV r5, #'x'
2861		CMP r4, r5
2862		MOVEQ r4, #1
2863		MOVNE r4, #0
2864		CMP r4, #0
2865		BEQ L88
2866		MOV r4, #'o'
2867		MOV r0, r4
2868		POP {pc}
2869		B L89
2870	L88:
2871		LDRSB r4, [sp, #4]
2872		MOV r5, #'o'
2873		CMP r4, r5
2874		MOVEQ r4, #1
2875		MOVNE r4, #0
2876		CMP r4, #0
2877		BEQ L90
2878		MOV r4, #'x'
2879		MOV r0, r4
2880		POP {pc}
2881		B L91
2882	L90:
2883		LDR r4, =msg_33
2884		MOV r0, r4
2885		BL p_print_string
2886		BL p_print_ln
2887		LDR r4, =-1
2888		MOV r0, r4
2889		BL exit
2890	L91:
2891	L89:
2892		POP {pc}
2893		.ltorg
2894	f_symbolAt:
2895		PUSH {lr}
2896		SUB sp, sp, #5
2897		LDR r4, =0
2898		STR r4, [sp, #1]
2899		LDR r4, [sp, #13]
2900		LDR r5, =2
2901		CMP r4, r5
2902		MOVLE r4, #1
2903		MOVGT r4, #0
2904		CMP r4, #0
2905		BEQ L92
2906		SUB sp, sp, #4
2907		LDR r4, [sp, #13]
2908		MOV r0, r4
2909		BL p_check_null_pointer
2910		LDR r4, [r4]
2911		LDR r4, [r4]
2912		STR r4, [sp]
2913		LDR r4, [sp, #17]
2914		LDR r5, =1
2915		CMP r4, r5
2916		MOVEQ r4, #1
2917		MOVNE r4, #0
2918		CMP r4, #0
2919		BEQ L94
2920		LDR r4, [sp]
2921		MOV r0, r4
2922		BL p_check_null_pointer
2923		LDR r4, [r4]
2924		LDR r4, [r4]
2925		STR r4, [sp, #5]
2926		B L95
2927	L94:
2928		LDR r4, [sp]
2929		MOV r0, r4
2930		BL p_check_null_pointer
2931		LDR r4, [r4, #4]
2932		LDR r4, [r4]
2933		STR r4, [sp, #5]
2934	L95:
2935		ADD sp, sp, #4
2936		B L93
2937	L92:
2938		LDR r4, [sp, #9]
2939		MOV r0, r4
2940		BL p_check_null_pointer
2941		LDR r4, [r4, #4]
2942		LDR r4, [r4]
2943		STR r4, [sp, #1]
2944	L93:
2945		MOV r4, #0
2946		STRB r4, [sp]
2947		LDR r4, [sp, #17]
2948		LDR r5, =2
2949		CMP r4, r5
2950		MOVLE r4, #1
2951		MOVGT r4, #0
2952		CMP r4, #0
2953		BEQ L96
2954		SUB sp, sp, #4
2955		LDR r4, [sp, #5]
2956		MOV r0, r4
2957		BL p_check_null_pointer
2958		LDR r4, [r4]
2959		LDR r4, [r4]
2960		STR r4, [sp]
2961		LDR r4, [sp, #21]
2962		LDR r5, =1
2963		CMP r4, r5
2964		MOVEQ r4, #1
2965		MOVNE r4, #0
2966		CMP r4, #0
2967		BEQ L98
2968		LDR r4, [sp]
2969		MOV r0, r4
2970		BL p_check_null_pointer
2971		LDR r4, [r4]
2972		LDRSB r4, [r4]
2973		STRB r4, [sp, #4]
2974		B L99
2975	L98:
2976		LDR r4, [sp]
2977		MOV r0, r4
2978		BL p_check_null_pointer
2979		LDR r4, [r4, #4]
2980		LDRSB r4, [r4]
2981		STRB r4, [sp, #4]
2982	L99:
2983		ADD sp, sp, #4
2984		B L97
2985	L96:
2986		LDR r4, [sp, #1]
2987		MOV r0, r4
2988		BL p_check_null_pointer
2989		LDR r4, [r4, #4]
2990		LDRSB r4, [r4]
2991		STRB r4, [sp]
2992	L97:
2993		LDRSB r4, [sp]
2994		MOV r0, r4
2995		ADD sp, sp, #5
2996		POP {pc}
2997		POP {pc}
2998		.ltorg
2999	f_containEmptyCell:
3000		PUSH {lr}
3001		SUB sp, sp, #19
3002		LDR r4, [sp, #23]
3003		MOV r0, r4
3004		BL p_check_null_pointer
3005		LDR r4, [r4]
3006		LDR r4, [r4]
3007		STR r4, [sp, #15]
3008		LDR r4, [sp, #15]
3009		MOV r0, r4
3010		BL p_check_null_pointer
3011		LDR r4, [r4]
3012		LDR r4, [r4]
3013		STR r4, [sp, #11]
3014		LDR r4, [sp, #15]
3015		MOV r0, r4
3016		BL p_check_null_pointer
3017		LDR r4, [r4, #4]
3018		LDR r4, [r4]
3019		STR r4, [sp, #7]
3020		LDR r4, [sp, #23]
3021		MOV r0, r4
3022		BL p_check_null_pointer
3023		LDR r4, [r4, #4]
3024		LDR r4, [r4]
3025		STR r4, [sp, #3]
3026		LDR r4, [sp, #11]
3027		STR r4, [sp, #-4]!
3028		BL f_containEmptyCellRow
3029		ADD sp, sp, #4
3030		MOV r4, r0
3031		STRB r4, [sp, #2]
3032		LDR r4, [sp, #7]
3033		STR r4, [sp, #-4]!
3034		BL f_containEmptyCellRow
3035		ADD sp, sp, #4
3036		MOV r4, r0
3037		STRB r4, [sp, #1]
3038		LDR r4, [sp, #3]
3039		STR r4, [sp, #-4]!
3040		BL f_containEmptyCellRow
3041		ADD sp, sp, #4
3042		MOV r4, r0
3043		STRB r4, [sp]
3044		LDRSB r4, [sp, #2]
3045		LDRSB r5, [sp, #1]
3046		ORR r4, r4, r5
3047		LDRSB r5, [sp]
3048		ORR r4, r4, r5
3049		MOV r0, r4
3050		ADD sp, sp, #19
3051		POP {pc}
3052		POP {pc}
3053		.ltorg
3054	f_containEmptyCellRow:
3055		PUSH {lr}
3056		SUB sp, sp, #7
3057		LDR r4, [sp, #11]
3058		MOV r0, r4
3059		BL p_check_null_pointer
3060		LDR r4, [r4]
3061		LDR r4, [r4]
3062		STR r4, [sp, #3]
3063		LDR r4, [sp, #3]
3064		MOV r0, r4
3065		BL p_check_null_pointer
3066		LDR r4, [r4]
3067		LDRSB r4, [r4]
3068		STRB r4, [sp, #2]
3069		LDR r4, [sp, #3]
3070		MOV r0, r4
3071		BL p_check_null_pointer
3072		LDR r4, [r4, #4]
3073		LDRSB r4, [r4]
3074		STRB r4, [sp, #1]
3075		LDR r4, [sp, #11]
3076		MOV r0, r4
3077		BL p_check_null_pointer
3078		LDR r4, [r4, #4]
3079		LDRSB r4, [r4]
3080		STRB r4, [sp]
3081		LDRSB r4, [sp, #2]
3082		MOV r5, #0
3083		CMP r4, r5
3084		MOVEQ r4, #1
3085		MOVNE r4, #0
3086		LDRSB r5, [sp, #1]
3087		MOV r6, #0
3088		CMP r5, r6
3089		MOVEQ r5, #1
3090		MOVNE r5, #0
3091		ORR r4, r4, r5
3092		LDRSB r5, [sp]
3093		MOV r6, #0
3094		CMP r5, r6
3095		MOVEQ r5, #1
3096		MOVNE r5, #0
3097		ORR r4, r4, r5
3098		MOV r0, r4
3099		ADD sp, sp, #7
3100		POP {pc}
3101		POP {pc}
3102		.ltorg
3103	f_hasWon:
3104		PUSH {lr}
3105		SUB sp, sp, #9
3106		LDR r4, =1
3107		STR r4, [sp, #-4]!
3108		LDR r4, =1
3109		STR r4, [sp, #-4]!
3110		LDR r4, [sp, #21]
3111		STR r4, [sp, #-4]!
3112		BL f_symbolAt
3113		ADD sp, sp, #12
3114		MOV r4, r0
3115		STRB r4, [sp, #8]
3116		LDR r4, =2
3117		STR r4, [sp, #-4]!
3118		LDR r4, =1
3119		STR r4, [sp, #-4]!
3120		LDR r4, [sp, #21]
3121		STR r4, [sp, #-4]!
3122		BL f_symbolAt
3123		ADD sp, sp, #12
3124		MOV r4, r0
3125		STRB r4, [sp, #7]
3126		LDR r4, =3
3127		STR r4, [sp, #-4]!
3128		LDR r4, =1
3129		STR r4, [sp, #-4]!
3130		LDR r4, [sp, #21]
3131		STR r4, [sp, #-4]!
3132		BL f_symbolAt
3133		ADD sp, sp, #12
3134		MOV r4, r0
3135		STRB r4, [sp, #6]
3136		LDR r4, =1
3137		STR r4, [sp, #-4]!
3138		LDR r4, =2
3139		STR r4, [sp, #-4]!
3140		LDR r4, [sp, #21]
3141		STR r4, [sp, #-4]!
3142		BL f_symbolAt
3143		ADD sp, sp, #12
3144		MOV r4, r0
3145		STRB r4, [sp, #5]
3146		LDR r4, =2
3147		STR r4, [sp, #-4]!
3148		LDR r4, =2
3149		STR r4, [sp, #-4]!
3150		LDR r4, [sp, #21]
3151		STR r4, [sp, #-4]!
3152		BL f_symbolAt
3153		ADD sp, sp, #12
3154		MOV r4, r0
3155		STRB r4, [sp, #4]
3156		LDR r4, =3
3157		STR r4, [sp, #-4]!
3158		LDR r4, =2
3159		STR r4, [sp, #-4]!
3160		LDR r4, [sp, #21]
3161		STR r4, [sp, #-4]!
3162		BL f_symbolAt
3163		ADD sp, sp, #12
3164		MOV r4, r0
3165		STRB r4, [sp, #3]
3166		LDR r4, =1
3167		STR r4, [sp, #-4]!
3168		LDR r4, =3
3169		STR r4, [sp, #-4]!
3170		LDR r4, [sp, #21]
3171		STR r4, [sp, #-4]!
3172		BL f_symbolAt
3173		ADD sp, sp, #12
3174		MOV r4, r0
3175		STRB r4, [sp, #2]
3176		LDR r4, =2
3177		STR r4, [sp, #-4]!
3178		LDR r4, =3
3179		STR r4, [sp, #-4]!
3180		LDR r4, [sp, #21]
3181		STR r4, [sp, #-4]!
3182		BL f_symbolAt
3183		ADD sp, sp, #12
3184		MOV r4, r0
3185		STRB r4, [sp, #1]
3186		LDR r4, =3
3187		STR r4, [sp, #-4]!
3188		LDR r4, =3
3189		STR r4, [sp, #-4]!
3190		LDR r4, [sp, #21]
3191		STR r4, [sp, #-4]!
3192		BL f_symbolAt
3193		ADD sp, sp, #12
3194		MOV r4, r0
3195		STRB r4, [sp]
3196		LDRSB r4, [sp, #8]
3197		LDRSB r5, [sp, #17]
3198		CMP r4, r5
3199		MOVEQ r4, #1
3200		MOVNE r4, #0
3201		LDRSB r5, [sp, #7]
3202		LDRSB r6, [sp, #17]
3203		CMP r5, r6
3204		MOVEQ r5, #1
3205		MOVNE r5, #0
3206		AND r4, r4, r5
3207		LDRSB r5, [sp, #6]
3208		LDRSB r6, [sp, #17]
3209		CMP r5, r6
3210		MOVEQ r5, #1
3211		MOVNE r5, #0
3212		AND r4, r4, r5
3213		LDRSB r5, [sp, #5]
3214		LDRSB r6, [sp, #17]
3215		CMP r5, r6
3216		MOVEQ r5, #1
3217		MOVNE r5, #0
3218		LDRSB r6, [sp, #4]
3219		LDRSB r7, [sp, #17]
3220		CMP r6, r7
3221		MOVEQ r6, #1
3222		MOVNE r6, #0
3223		AND r5, r5, r6
3224		LDRSB r6, [sp, #3]
3225		LDRSB r7, [sp, #17]
3226		CMP r6, r7
3227		MOVEQ r6, #1
3228		MOVNE r6, #0
3229		AND r5, r5, r6
3230		ORR r4, r4, r5
3231		LDRSB r5, [sp, #2]
3232		LDRSB r6, [sp, #17]
3233		CMP r5, r6
3234		MOVEQ r5, #1
3235		MOVNE r5, #0
3236		LDRSB r6, [sp, #1]
3237		LDRSB r7, [sp, #17]
3238		CMP r6, r7
3239		MOVEQ r6, #1
3240		MOVNE r6, #0
3241		AND r5, r5, r6
3242		LDRSB r6, [sp]
3243		LDRSB r7, [sp, #17]
3244		CMP r6, r7
3245		MOVEQ r6, #1
3246		MOVNE r6, #0
3247		AND r5, r5, r6
3248		ORR r4, r4, r5
3249		LDRSB r5, [sp, #8]
3250		LDRSB r6, [sp, #17]
3251		CMP r5, r6
3252		MOVEQ r5, #1
3253		MOVNE r5, #0
3254		LDRSB r6, [sp, #5]
3255		LDRSB r7, [sp, #17]
3256		CMP r6, r7
3257		MOVEQ r6, #1
3258		MOVNE r6, #0
3259		AND r5, r5, r6
3260		LDRSB r6, [sp, #2]
3261		LDRSB r7, [sp, #17]
3262		CMP r6, r7
3263		MOVEQ r6, #1
3264		MOVNE r6, #0
3265		AND r5, r5, r6
3266		ORR r4, r4, r5
3267		LDRSB r5, [sp, #7]
3268		LDRSB r6, [sp, #17]
3269		CMP r5, r6
3270		MOVEQ r5, #1
3271		MOVNE r5, #0
3272		LDRSB r6, [sp, #4]
3273		LDRSB r7, [sp, #17]
3274		CMP r6, r7
3275		MOVEQ r6, #1
3276		MOVNE r6, #0
3277		AND r5, r5, r6
3278		LDRSB r6, [sp, #1]
3279		LDRSB r7, [sp, #17]
3280		CMP r6, r7
3281		MOVEQ r6, #1
3282		MOVNE r6, #0
3283		AND r5, r5, r6
3284		ORR r4, r4, r5
3285		LDRSB r5, [sp, #6]
3286		LDRSB r6, [sp, #17]
3287		CMP r5, r6
3288		MOVEQ r5, #1
3289		MOVNE r5, #0
3290		LDRSB r6, [sp, #3]
3291		LDRSB r7, [sp, #17]
3292		CMP r6, r7
3293		MOVEQ r6, #1
3294		MOVNE r6, #0
3295		AND r5, r5, r6
3296		LDRSB r6, [sp]
3297		LDRSB r7, [sp, #17]
3298		CMP r6, r7
3299		MOVEQ r6, #1
3300		MOVNE r6, #0
3301		AND r5, r5, r6
3302		ORR r4, r4, r5
3303		LDRSB r5, [sp, #8]
3304		LDRSB r6, [sp, #17]
3305		CMP r5, r6
3306		MOVEQ r5, #1
3307		MOVNE r5, #0
3308		LDRSB r6, [sp, #4]
3309		LDRSB r7, [sp, #17]
3310		CMP r6, r7
3311		MOVEQ r6, #1
3312		MOVNE r6, #0
3313		AND r5, r5, r6
3314		LDRSB r6, [sp]
3315		LDRSB r7, [sp, #17]
3316		CMP r6, r7
3317		MOVEQ r6, #1
3318		MOVNE r6, #0
3319		AND r5, r5, r6
3320		ORR r4, r4, r5
3321		LDRSB r5, [sp, #6]
3322		LDRSB r6, [sp, #17]
3323		CMP r5, r6
3324		MOVEQ r5, #1
3325		MOVNE r5, #0
3326		LDRSB r6, [sp, #4]
3327		LDRSB r7, [sp, #17]
3328		CMP r6, r7
3329		MOVEQ r6, #1
3330		MOVNE r6, #0
3331		AND r5, r5, r6
3332		LDRSB r6, [sp, #2]
3333		LDRSB r7, [sp, #17]
3334		CMP r6, r7
3335		MOVEQ r6, #1
3336		MOVNE r6, #0
3337		AND r5, r5, r6
3338		ORR r4, r4, r5
3339		MOV r0, r4
3340		ADD sp, sp, #9
3341		POP {pc}
3342		POP {pc}
3343		.ltorg
3344	f_allocateNewBoard:
3345		PUSH {lr}
3346		SUB sp, sp, #20
3347		BL f_allocateNewRow
3348		MOV r4, r0
3349		STR r4, [sp, #16]
3350		BL f_allocateNewRow
3351		MOV r4, r0
3352		STR r4, [sp, #12]
3353		BL f_allocateNewRow
3354		MOV r4, r0
3355		STR r4, [sp, #8]
3356		LDR r0, =8
3357		BL malloc
3358		MOV r4, r0
3359		LDR r5, [sp, #16]
3360		LDR r0, =4
3361		BL malloc
3362		STR r5, [r0]
3363		STR r0, [r4]
3364		LDR r5, [sp, #12]
3365		LDR r0, =4
3366		BL malloc
3367		STR r5, [r0]
3368		STR r0, [r4, #4]
3369		STR r4, [sp, #4]
3370		LDR r0, =8
3371		BL malloc
3372		MOV r4, r0
3373		LDR r5, [sp, #4]
3374		LDR r0, =4
3375		BL malloc
3376		STR r5, [r0]
3377		STR r0, [r4]
3378		LDR r5, [sp, #8]
3379		LDR r0, =4
3380		BL malloc
3381		STR r5, [r0]
3382		STR r0, [r4, #4]
3383		STR r4, [sp]
3384		LDR r4, [sp]
3385		MOV r0, r4
3386		ADD sp, sp, #20
3387		POP {pc}
3388		POP {pc}
3389		.ltorg
3390	f_allocateNewRow:
3391		PUSH {lr}
3392		SUB sp, sp, #8
3393		LDR r0, =8
3394		BL malloc
3395		MOV r4, r0
3396		MOV r5, #0
3397		LDR r0, =1
3398		BL malloc
3399		STRB r5, [r0]
3400		STR r0, [r4]
3401		MOV r5, #0
3402		LDR r0, =1
3403		BL malloc
3404		STRB r5, [r0]
3405		STR r0, [r4, #4]
3406		STR r4, [sp, #4]
3407		LDR r0, =8
3408		BL malloc
3409		MOV r4, r0
3410		LDR r5, [sp, #4]
3411		LDR r0, =4
3412		BL malloc
3413		STR r5, [r0]
3414		STR r0, [r4]
3415		MOV r5, #0
3416		LDR r0, =1
3417		BL malloc
3418		STRB r5, [r0]
3419		STR r0, [r4, #4]
3420		STR r4, [sp]
3421		LDR r4, [sp]
3422		MOV r0, r4
3423		ADD sp, sp, #8
3424		POP {pc}
3425		POP {pc}
3426		.ltorg
3427	f_freeBoard:
3428		PUSH {lr}
3429		SUB sp, sp, #17
3430		LDR r4, [sp, #21]
3431		MOV r0, r4
3432		BL p_check_null_pointer
3433		LDR r4, [r4]
3434		LDR r4, [r4]
3435		STR r4, [sp, #13]
3436		LDR r4, [sp, #13]
3437		MOV r0, r4
3438		BL p_check_null_pointer
3439		LDR r4, [r4]
3440		LDR r4, [r4]
3441		STR r4, [sp, #9]
3442		LDR r4, [sp, #13]
3443		MOV r0, r4
3444		BL p_check_null_pointer
3445		LDR r4, [r4, #4]
3446		LDR r4, [r4]
3447		STR r4, [sp, #5]
3448		LDR r4, [sp, #21]
3449		MOV r0, r4
3450		BL p_check_null_pointer
3451		LDR r4, [r4, #4]
3452		LDR r4, [r4]
3453		STR r4, [sp, #1]
3454		LDR r4, [sp, #9]
3455		STR r4, [sp, #-4]!
3456		BL f_freeRow
3457		ADD sp, sp, #4
3458		MOV r4, r0
3459		STRB r4, [sp]
3460		LDR r4, [sp, #5]
3461		STR r4, [sp, #-4]!
3462		BL f_freeRow
3463		ADD sp, sp, #4
3464		MOV r4, r0
3465		STRB r4, [sp]
3466		LDR r4, [sp, #1]
3467		STR r4, [sp, #-4]!
3468		BL f_freeRow
3469		ADD sp, sp, #4
3470		MOV r4, r0
3471		STRB r4, [sp]
3472		LDR r4, [sp, #13]
3473		MOV r0, r4
3474		BL p_free_pair
3475		LDR r4, [sp, #21]
3476		MOV r0, r4
3477		BL p_free_pair
3478		MOV r4, #1
3479		MOV r0, r4
3480		ADD sp, sp, #17
3481		POP {pc}
3482		POP {pc}
3483		.ltorg
3484	f_freeRow:
3485		PUSH {lr}
3486		SUB sp, sp, #4
3487		LDR r4, [sp, #8]
3488		MOV r0, r4
3489		BL p_check_null_pointer
3490		LDR r4, [r4]
3491		LDR r4, [r4]
3492		STR r4, [sp]
3493		LDR r4, [sp]
3494		MOV r0, r4
3495		BL p_free_pair
3496		LDR r4, [sp, #8]
3497		MOV r0, r4
3498		BL p_free_pair
3499		MOV r4, #1
3500		MOV r0, r4
3501		ADD sp, sp, #4
3502		POP {pc}
3503		POP {pc}
3504		.ltorg
3505	f_printAiData:
3506		PUSH {lr}
3507		SUB sp, sp, #9
3508		LDR r4, [sp, #13]
3509		MOV r0, r4
3510		BL p_check_null_pointer
3511		LDR r4, [r4]
3512		LDR r4, [r4]
3513		STR r4, [sp, #5]
3514		LDR r4, [sp, #13]
3515		MOV r0, r4
3516		BL p_check_null_pointer
3517		LDR r4, [r4, #4]
3518		LDR r4, [r4]
3519		STR r4, [sp, #1]
3520		LDR r4, [sp, #1]
3521		STR r4, [sp, #-4]!
3522		BL f_printStateTreeRecursively
3523		ADD sp, sp, #4
3524		MOV r4, r0
3525		STRB r4, [sp]
3526		LDR r4, =0
3527		MOV r0, r4
3528		BL exit
3529		POP {pc}
3530		.ltorg
3531	f_printStateTreeRecursively:
3532		PUSH {lr}
3533		LDR r4, [sp, #4]
3534		LDR r5, =0
3535		CMP r4, r5
3536		MOVEQ r4, #1
3537		MOVNE r4, #0
3538		CMP r4, #0
3539		BEQ L100
3540		MOV r4, #1
3541		MOV r0, r4
3542		POP {pc}
3543		B L101
3544	L100:
3545		SUB sp, sp, #17
3546		LDR r4, [sp, #21]
3547		MOV r0, r4
3548		BL p_check_null_pointer
3549		LDR r4, [r4]
3550		LDR r4, [r4]
3551		STR r4, [sp, #13]
3552		LDR r4, [sp, #13]
3553		MOV r0, r4
3554		BL p_check_null_pointer
3555		LDR r4, [r4]
3556		LDR r4, [r4]
3557		STR r4, [sp, #9]
3558		LDR r4, [sp, #13]
3559		MOV r0, r4
3560		BL p_check_null_pointer
3561		LDR r4, [r4, #4]
3562		LDR r4, [r4]
3563		STR r4, [sp, #5]
3564		LDR r4, [sp, #21]
3565		MOV r0, r4
3566		BL p_check_null_pointer
3567		LDR r4, [r4, #4]
3568		LDR r4, [r4]
3569		STR r4, [sp, #1]
3570		MOV r4, #'v'
3571		MOV r0, r4
3572		BL putchar
3573		MOV r4, #'='
3574		MOV r0, r4
3575		BL putchar
3576		LDR r4, [sp, #1]
3577		MOV r0, r4
3578		BL p_print_int
3579		BL p_print_ln
3580		LDR r4, [sp, #9]
3581		STR r4, [sp, #-4]!
3582		BL f_printBoard
3583		ADD sp, sp, #4
3584		MOV r4, r0
3585		STRB r4, [sp]
3586		LDR r4, [sp, #5]
3587		STR r4, [sp, #-4]!
3588		BL f_printChildrenStateTree
3589		ADD sp, sp, #4
3590		MOV r4, r0
3591		STRB r4, [sp]
3592		MOV r4, #'p'
3593		MOV r0, r4
3594		BL putchar
3595		BL p_print_ln
3596		MOV r4, #1
3597		MOV r0, r4
3598		ADD sp, sp, #17
3599		POP {pc}
3600		ADD sp, sp, #17
3601	L101:
3602		POP {pc}
3603		.ltorg
3604	f_printChildrenStateTree:
3605		PUSH {lr}
3606		SUB sp, sp, #17
3607		LDR r4, [sp, #21]
3608		MOV r0, r4
3609		BL p_check_null_pointer
3610		LDR r4, [r4]
3611		LDR r4, [r4]
3612		STR r4, [sp, #13]
3613		LDR r4, [sp, #13]
3614		MOV r0, r4
3615		BL p_check_null_pointer
3616		LDR r4, [r4]
3617		LDR r4, [r4]
3618		STR r4, [sp, #9]
3619		LDR r4, [sp, #13]
3620		MOV r0, r4
3621		BL p_check_null_pointer
3622		LDR r4, [r4, #4]
3623		LDR r4, [r4]
3624		STR r4, [sp, #5]
3625		LDR r4, [sp, #21]
3626		MOV r0, r4
3627		BL p_check_null_pointer
3628		LDR r4, [r4, #4]
3629		LDR r4, [r4]
3630		STR r4, [sp, #1]
3631		LDR r4, [sp, #9]
3632		STR r4, [sp, #-4]!
3633		BL f_printChildrenStateTreeRow
3634		ADD sp, sp, #4
3635		MOV r4, r0
3636		STRB r4, [sp]
3637		LDR r4, [sp, #5]
3638		STR r4, [sp, #-4]!
3639		BL f_printChildrenStateTreeRow
3640		ADD sp, sp, #4
3641		MOV r4, r0
3642		STRB r4, [sp]
3643		LDR r4, [sp, #1]
3644		STR r4, [sp, #-4]!
3645		BL f_printChildrenStateTreeRow
3646		ADD sp, sp, #4
3647		MOV r4, r0
3648		STRB r4, [sp]
3649		MOV r4, #1
3650		MOV r0, r4
3651		ADD sp, sp, #17
3652		POP {pc}
3653		POP {pc}
3654		.ltorg
3655	f_printChildrenStateTreeRow:
3656		PUSH {lr}
3657		SUB sp, sp, #17
3658		LDR r4, [sp, #21]
3659		MOV r0, r4
3660		BL p_check_null_pointer
3661		LDR r4, [r4]
3662		LDR r4, [r4]
3663		STR r4, [sp, #13]
3664		LDR r4, [sp, #13]
3665		MOV r0, r4
3666		BL p_check_null_pointer
3667		LDR r4, [r4]
3668		LDR r4, [r4]
3669		STR r4, [sp, #9]
3670		LDR r4, [sp, #13]
3671		MOV r0, r4
3672		BL p_check_null_pointer
3673		LDR r4, [r4, #4]
3674		LDR r4, [r4]
3675		STR r4, [sp, #5]
3676		LDR r4, [sp, #21]
3677		MOV r0, r4
3678		BL p_check_null_pointer
3679		LDR r4, [r4, #4]
3680		LDR r4, [r4]
3681		STR r4, [sp, #1]
3682		LDR r4, [sp, #9]
3683		STR r4, [sp, #-4]!
3684		BL f_printStateTreeRecursively
3685		ADD sp, sp, #4
3686		MOV r4, r0
3687		STRB r4, [sp]
3688		LDR r4, [sp, #5]
3689		STR r4, [sp, #-4]!
3690		BL f_printStateTreeRecursively
3691		ADD sp, sp, #4
3692		MOV r4, r0
3693		STRB r4, [sp]
3694		LDR r4, [sp, #1]
3695		STR r4, [sp, #-4]!
3696		BL f_printStateTreeRecursively
3697		ADD sp, sp, #4
3698		MOV r4, r0
3699		STRB r4, [sp]
3700		MOV r4, #1
3701		MOV r0, r4
3702		ADD sp, sp, #17
3703		POP {pc}
3704		POP {pc}
3705		.ltorg
3706	main:
3707		PUSH {lr}
3708		SUB sp, sp, #17
3709		BL f_chooseSymbol
3710		MOV r4, r0
3711		STRB r4, [sp, #16]
3712		LDRSB r4, [sp, #16]
3713		STRB r4, [sp, #-1]!
3714		BL f_oppositeSymbol
3715		ADD sp, sp, #1
3716		MOV r4, r0
3717		STRB r4, [sp, #15]
3718		MOV r4, #'x'
3719		STRB r4, [sp, #14]
3720		BL f_allocateNewBoard
3721		MOV r4, r0
3722		STR r4, [sp, #10]
3723		LDR r4, =msg_34
3724		MOV r0, r4
3725		BL p_print_string
3726		BL p_print_ln
3727		LDRSB r4, [sp, #15]
3728		STRB r4, [sp, #-1]!
3729		BL f_initAI
3730		ADD sp, sp, #1
3731		MOV r4, r0
3732		STR r4, [sp, #6]
3733		LDR r4, =0
3734		STR r4, [sp, #2]
3735		MOV r4, #0
3736		STRB r4, [sp, #1]
3737		LDR r4, [sp, #10]
3738		STR r4, [sp, #-4]!
3739		BL f_printBoard
3740		ADD sp, sp, #4
3741		MOV r4, r0
3742		STRB r4, [sp]
3743		B L102
3744	L103:
3745		SUB sp, sp, #5
3746		LDR r0, =12
3747		BL malloc
3748		MOV r4, r0
3749		LDR r5, =0
3750		STR r5, [r4, #4]
3751		LDR r5, =0
3752		STR r5, [r4, #8]
3753		LDR r5, =2
3754		STR r5, [r4]
3755		STR r4, [sp, #1]
3756		LDR r4, [sp, #1]
3757		STR r4, [sp, #-4]!
3758		LDR r4, [sp, #15]
3759		STR r4, [sp, #-4]!
3760		LDRSB r4, [sp, #29]
3761		STRB r4, [sp, #-1]!
3762		LDRSB r4, [sp, #28]
3763		STRB r4, [sp, #-1]!
3764		LDR r4, [sp, #25]
3765		STR r4, [sp, #-4]!
3766		BL f_askForAMove
3767		ADD sp, sp, #14
3768		MOV r4, r0
3769		STRB r4, [sp, #5]
3770		ADD r4, sp, #1
3771		LDR r5, =1
3772		LDR r4, [r4]
3773		MOV r0, r5
3774		MOV r1, r4
3775		BL p_check_array_bounds
3776		ADD r4, r4, #4
3777		ADD r4, r4, r5, LSL #2
3778		LDR r4, [r4]
3779		STR r4, [sp, #-4]!
3780		ADD r4, sp, #5
3781		LDR r5, =0
3782		LDR r4, [r4]
3783		MOV r0, r5
3784		MOV r1, r4
3785		BL p_check_array_bounds
3786		ADD r4, r4, #4
3787		ADD r4, r4, r5, LSL #2
3788		LDR r4, [r4]
3789		STR r4, [sp, #-4]!
3790		LDRSB r4, [sp, #27]
3791		STRB r4, [sp, #-1]!
3792		LDR r4, [sp, #24]
3793		STR r4, [sp, #-4]!
3794		BL f_placeMove
3795		ADD sp, sp, #13
3796		MOV r4, r0
3797		STRB r4, [sp, #5]
3798		ADD r4, sp, #1
3799		LDR r5, =1
3800		LDR r4, [r4]
3801		MOV r0, r5
3802		MOV r1, r4
3803		BL p_check_array_bounds
3804		ADD r4, r4, #4
3805		ADD r4, r4, r5, LSL #2
3806		LDR r4, [r4]
3807		STR r4, [sp, #-4]!
3808		ADD r4, sp, #5
3809		LDR r5, =0
3810		LDR r4, [r4]
3811		MOV r0, r5
3812		MOV r1, r4
3813		BL p_check_array_bounds
3814		ADD r4, r4, #4
3815		ADD r4, r4, r5, LSL #2
3816		LDR r4, [r4]
3817		STR r4, [sp, #-4]!
3818		LDR r4, [sp, #19]
3819		STR r4, [sp, #-4]!
3820		LDRSB r4, [sp, #33]
3821		STRB r4, [sp, #-1]!
3822		LDRSB r4, [sp, #32]
3823		STRB r4, [sp, #-1]!
3824		LDR r4, [sp, #29]
3825		STR r4, [sp, #-4]!
3826		BL f_notifyMove
3827		ADD sp, sp, #18
3828		MOV r4, r0
3829		STRB r4, [sp, #5]
3830		LDR r4, [sp, #15]
3831		STR r4, [sp, #-4]!
3832		BL f_printBoard
3833		ADD sp, sp, #4
3834		MOV r4, r0
3835		STRB r4, [sp, #5]
3836		LDRSB r4, [sp, #19]
3837		STRB r4, [sp, #-1]!
3838		LDR r4, [sp, #16]
3839		STR r4, [sp, #-4]!
3840		BL f_hasWon
3841		ADD sp, sp, #5
3842		MOV r4, r0
3843		STRB r4, [sp]
3844		LDRSB r4, [sp]
3845		CMP r4, #0
3846		BEQ L104
3847		LDRSB r4, [sp, #19]
3848		STRB r4, [sp, #6]
3849		B L105
3850	L104:
3851	L105:
3852		LDRSB r4, [sp, #19]
3853		STRB r4, [sp, #-1]!
3854		BL f_oppositeSymbol
3855		ADD sp, sp, #1
3856		MOV r4, r0
3857		STRB r4, [sp, #19]
3858		LDR r4, [sp, #7]
3859		LDR r5, =1
3860		ADDS r4, r4, r5
3861		BLVS p_throw_overflow_error
3862		STR r4, [sp, #7]
3863		ADD sp, sp, #5
3864	L102:
3865		LDRSB r4, [sp, #1]
3866		MOV r5, #0
3867		CMP r4, r5
3868		MOVEQ r4, #1
3869		MOVNE r4, #0
3870		LDR r5, [sp, #2]
3871		LDR r6, =9
3872		CMP r5, r6
3873		MOVLT r5, #1
3874		MOVGE r5, #0
3875		AND r4, r4, r5
3876		CMP r4, #1
3877		BEQ L103
3878		LDR r4, [sp, #10]
3879		STR r4, [sp, #-4]!
3880		BL f_freeBoard
3881		ADD sp, sp, #4
3882		MOV r4, r0
3883		STRB r4, [sp]
3884		LDR r4, [sp, #6]
3885		STR r4, [sp, #-4]!
3886		BL f_destroyAI
3887		ADD sp, sp, #4
3888		MOV r4, r0
3889		STRB r4, [sp]
3890		LDRSB r4, [sp, #1]
3891		MOV r5, #0
3892		CMP r4, r5
3893		MOVNE r4, #1
3894		MOVEQ r4, #0
3895		CMP r4, #0
3896		BEQ L106
3897		LDRSB r4, [sp, #1]
3898		MOV r0, r4
3899		BL putchar
3900		LDR r4, =msg_35
3901		MOV r0, r4
3902		BL p_print_string
3903		BL p_print_ln
3904		B L107
3905	L106:
3906		LDR r4, =msg_36
3907		MOV r0, r4
3908		BL p_print_string
3909		BL p_print_ln
3910	L107:
3911		ADD sp, sp, #17
3912		LDR r0, =0
3913		POP {pc}
3914		.ltorg
3915	p_print_string:
3916		PUSH {lr}
3917		LDR r1, [r0]
3918		ADD r2, r0, #4
3919		LDR r0, =msg_37
3920		ADD r0, r0, #4
3921		BL printf
3922		MOV r0, #0
3923		BL fflush
3924		POP {pc}
3925	p_print_ln:
3926		PUSH {lr}
3927		LDR r0, =msg_38
3928		ADD r0, r0, #4
3929		BL puts
3930		MOV r0, #0
3931		BL fflush
3932		POP {pc}
3933	p_read_char:
3934		PUSH {lr}
3935		MOV r1, r0
3936		LDR r0, =msg_39
3937		ADD r0, r0, #4
3938		BL scanf
3939		POP {pc}
3940	p_check_null_pointer:
3941		PUSH {lr}
3942		CMP r0, #0
3943		LDREQ r0, =msg_40
3944		BLEQ p_throw_runtime_error
3945		POP {pc}
3946	p_read_int:
3947		PUSH {lr}
3948		MOV r1, r0
3949		LDR r0, =msg_41
3950		ADD r0, r0, #4
3951		BL scanf
3952		POP {pc}
3953	p_check_array_bounds:
3954		PUSH {lr}
3955		CMP r0, #0
3956		LDRLT r0, =msg_42
3957		BLLT p_throw_runtime_error
3958		LDR r1, [r1]
3959		CMP r0, r1
3960		LDRCS r0, =msg_43
3961		BLCS p_throw_runtime_error
3962		POP {pc}
3963	p_print_int:
3964		PUSH {lr}
3965		MOV r1, r0
3966		LDR r0, =msg_44
3967		ADD r0, r0, #4
3968		BL printf
3969		MOV r0, #0
3970		BL fflush
3971		POP {pc}
3972	p_free_pair:
3973		PUSH {lr}
3974		CMP r0, #0
3975		LDREQ r0, =msg_45
3976		BEQ p_throw_runtime_error
3977		PUSH {r0}
3978		LDR r0, [r0]
3979		BL free
3980		LDR r0, [sp]
3981		LDR r0, [r0, #4]
3982		BL free
3983		POP {r0}
3984		BL free
3985		POP {pc}
3986	p_throw_overflow_error:
3987		LDR r0, =msg_46
3988		BL p_throw_runtime_error
3989	p_throw_runtime_error:
3990		BL p_print_string
3991		MOV r0, #-1
3992		BL exit
3993	
===========================================================
-- Finished