-- Compiling...
-- Printing Assembly...
whitespace.s contents are:
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
12	main:
13		PUSH {lr}
14		SUB sp, sp, #4
15		LDR r4, =13
16		STR r4, [sp]
17		LDR r4, [sp]
18		LDR r5, =13
19		CMP r4, r5
20		MOVEQ r4, #1
21		MOVNE r4, #0
22		CMP r4, #0
23		BEQ L0
24		LDR r4, =1
25		STR r4, [sp]
26		B L1
27	L0:
28		LDR r4, =0
29		STR r4, [sp]
30	L1:
31		LDR r4, [sp]
32		MOV r0, r4
33		BL p_print_int
34		BL p_print_ln
35		ADD sp, sp, #4
36		LDR r0, =0
37		POP {pc}
38		.ltorg
39	p_print_int:
40		PUSH {lr}
41		MOV r1, r0
42		LDR r0, =msg_0
43		ADD r0, r0, #4
44		BL printf
45		MOV r0, #0
46		BL fflush
47		POP {pc}
48	p_print_ln:
49		PUSH {lr}
50		LDR r0, =msg_1
51		ADD r0, r0, #4
52		BL puts
53		MOV r0, #0
54		BL fflush
55		POP {pc}
56	
===========================================================
-- Finished
