-- Compiling...
-- Printing Assembly...
functionForwardRef.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 3
4		.ascii	"%d\0"
5	msg_1:
6		.word 1
7		.ascii	"\0"
8	
9	.text
10	
11	.global main
12	f_f:
13		PUSH {lr}
14		SUB sp, sp, #4
15		BL f_g
16		MOV r4, r0
17		STR r4, [sp]
18		LDR r4, =0
19		MOV r0, r4
20		ADD sp, sp, #4
21		POP {pc}
22		POP {pc}
23		.ltorg
24	f_h:
25		PUSH {lr}
26		LDR r4, =0
27		MOV r0, r4
28		POP {pc}
29		POP {pc}
30		.ltorg
31	f_g:
32		PUSH {lr}
33		SUB sp, sp, #4
34		BL f_h
35		MOV r4, r0
36		STR r4, [sp]
37		LDR r4, =0
38		MOV r0, r4
39		ADD sp, sp, #4
40		POP {pc}
41		POP {pc}
42		.ltorg
43	main:
44		PUSH {lr}
45		SUB sp, sp, #4
46		BL f_f
47		MOV r4, r0
48		STR r4, [sp]
49		LDR r4, [sp]
50		MOV r0, r4
51		BL p_print_int
52		BL p_print_ln
53		ADD sp, sp, #4
54		LDR r0, =0
55		POP {pc}
56		.ltorg
57	p_print_int:
58		PUSH {lr}
59		MOV r1, r0
60		LDR r0, =msg_0
61		ADD r0, r0, #4
62		BL printf
63		MOV r0, #0
64		BL fflush
65		POP {pc}
66	p_print_ln:
67		PUSH {lr}
68		LDR r0, =msg_1
69		ADD r0, r0, #4
70		BL puts
71		MOV r0, #0
72		BL fflush
73		POP {pc}
74	
===========================================================
-- Finished
