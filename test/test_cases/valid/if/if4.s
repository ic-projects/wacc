-- Compiling...
-- Printing Assembly...
if4.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 9
4		.ascii	"incorrect"
5	msg_1:
6		.word 7
7		.ascii	"correct"
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 1
13		.ascii	"\0"
14	
15	.text
16	
17	.global main
18	main:
19		PUSH {lr}
20		SUB sp, sp, #2
21		MOV r4, #1
22		STRB r4, [sp, #1]
23		MOV r4, #0
24		STRB r4, [sp]
25		LDRSB r4, [sp, #1]
26		LDRSB r5, [sp]
27		AND r4, r4, r5
28		CMP r4, #0
29		BEQ L0
30		LDR r4, =msg_0
31		MOV r0, r4
32		BL p_print_string
33		BL p_print_ln
34		B L1
35	L0:
36		LDR r4, =msg_1
37		MOV r0, r4
38		BL p_print_string
39		BL p_print_ln
40	L1:
41		ADD sp, sp, #2
42		LDR r0, =0
43		POP {pc}
44		.ltorg
45	p_print_string:
46		PUSH {lr}
47		LDR r1, [r0]
48		ADD r2, r0, #4
49		LDR r0, =msg_2
50		ADD r0, r0, #4
51		BL printf
52		MOV r0, #0
53		BL fflush
54		POP {pc}
55	p_print_ln:
56		PUSH {lr}
57		LDR r0, =msg_3
58		ADD r0, r0, #4
59		BL puts
60		MOV r0, #0
61		BL fflush
62		POP {pc}
63	
===========================================================
-- Finished
