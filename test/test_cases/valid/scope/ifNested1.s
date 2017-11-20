-- Compiling...
-- Printing Assembly...
ifNested1.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 7
4		.ascii	"correct"
5	msg_1:
6		.word 9
7		.ascii	"incorrect"
8	msg_2:
9		.word 9
10		.ascii	"incorrect"
11	msg_3:
12		.word 5
13		.ascii	"%.*s\0"
14	msg_4:
15		.word 1
16		.ascii	"\0"
17	
18	.text
19	
20	.global main
21	main:
22		PUSH {lr}
23		SUB sp, sp, #4
24		LDR r4, =13
25		STR r4, [sp]
26		LDR r4, [sp]
27		LDR r5, =13
28		CMP r4, r5
29		MOVEQ r4, #1
30		MOVNE r4, #0
31		CMP r4, #0
32		BEQ L0
33		LDR r4, [sp]
34		LDR r5, =5
35		CMP r4, r5
36		MOVGT r4, #1
37		MOVLE r4, #0
38		CMP r4, #0
39		BEQ L2
40		LDR r4, =msg_0
41		MOV r0, r4
42		BL p_print_string
43		BL p_print_ln
44		B L3
45	L2:
46		LDR r4, =msg_1
47		MOV r0, r4
48		BL p_print_string
49		BL p_print_ln
50	L3:
51		B L1
52	L0:
53		LDR r4, =msg_2
54		MOV r0, r4
55		BL p_print_string
56		BL p_print_ln
57	L1:
58		ADD sp, sp, #4
59		LDR r0, =0
60		POP {pc}
61		.ltorg
62	p_print_string:
63		PUSH {lr}
64		LDR r1, [r0]
65		ADD r2, r0, #4
66		LDR r0, =msg_3
67		ADD r0, r0, #4
68		BL printf
69		MOV r0, #0
70		BL fflush
71		POP {pc}
72	p_print_ln:
73		PUSH {lr}
74		LDR r0, =msg_4
75		ADD r0, r0, #4
76		BL puts
77		MOV r0, #0
78		BL fflush
79		POP {pc}
80	
===========================================================
-- Finished
