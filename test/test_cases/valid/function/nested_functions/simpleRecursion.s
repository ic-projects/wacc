-- Compiling...
-- Printing Assembly...
simpleRecursion.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 82
4		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
5	msg_1:
6		.word 5
7		.ascii	"%.*s\0"
8	
9	.text
10	
11	.global main
12	f_rec:
13		PUSH {lr}
14		LDR r4, [sp, #4]
15		LDR r5, =0
16		CMP r4, r5
17		MOVEQ r4, #1
18		MOVNE r4, #0
19		CMP r4, #0
20		BEQ L0
21		B L1
22	L0:
23		SUB sp, sp, #4
24		LDR r4, [sp, #8]
25		LDR r5, =1
26		SUBS r4, r4, r5
27		BLVS p_throw_overflow_error
28		STR r4, [sp, #-4]!
29		BL f_rec
30		ADD sp, sp, #4
31		MOV r4, r0
32		STR r4, [sp]
33		ADD sp, sp, #4
34	L1:
35		LDR r4, =42
36		MOV r0, r4
37		POP {pc}
38		POP {pc}
39		.ltorg
40	main:
41		PUSH {lr}
42		SUB sp, sp, #4
43		LDR r4, =0
44		STR r4, [sp]
45		LDR r4, =8
46		STR r4, [sp, #-4]!
47		BL f_rec
48		ADD sp, sp, #4
49		MOV r4, r0
50		STR r4, [sp]
51		ADD sp, sp, #4
52		LDR r0, =0
53		POP {pc}
54		.ltorg
55	p_throw_overflow_error:
56		LDR r0, =msg_0
57		BL p_throw_runtime_error
58	p_throw_runtime_error:
59		BL p_print_string
60		MOV r0, #-1
61		BL exit
62	p_print_string:
63		PUSH {lr}
64		LDR r1, [r0]
65		ADD r2, r0, #4
66		LDR r0, =msg_1
67		ADD r0, r0, #4
68		BL printf
69		MOV r0, #0
70		BL fflush
71		POP {pc}
72	
===========================================================
-- Finished
