-- Compiling...
-- Printing Assembly...
incFunction.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 82
4		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
5	msg_1:
6		.word 3
7		.ascii	"%d\0"
8	msg_2:
9		.word 1
10		.ascii	"\0"
11	msg_3:
12		.word 5
13		.ascii	"%.*s\0"
14	
15	.text
16	
17	.global main
18	f_inc:
19		PUSH {lr}
20		LDR r4, [sp, #4]
21		LDR r5, =1
22		ADDS r4, r4, r5
23		BLVS p_throw_overflow_error
24		MOV r0, r4
25		POP {pc}
26		POP {pc}
27		.ltorg
28	main:
29		PUSH {lr}
30		SUB sp, sp, #4
31		LDR r4, =0
32		STR r4, [sp]
33		LDR r4, [sp]
34		STR r4, [sp, #-4]!
35		BL f_inc
36		ADD sp, sp, #4
37		MOV r4, r0
38		STR r4, [sp]
39		LDR r4, [sp]
40		MOV r0, r4
41		BL p_print_int
42		BL p_print_ln
43		LDR r4, [sp]
44		STR r4, [sp, #-4]!
45		BL f_inc
46		ADD sp, sp, #4
47		MOV r4, r0
48		STR r4, [sp]
49		LDR r4, [sp]
50		STR r4, [sp, #-4]!
51		BL f_inc
52		ADD sp, sp, #4
53		MOV r4, r0
54		STR r4, [sp]
55		LDR r4, [sp]
56		STR r4, [sp, #-4]!
57		BL f_inc
58		ADD sp, sp, #4
59		MOV r4, r0
60		STR r4, [sp]
61		LDR r4, [sp]
62		MOV r0, r4
63		BL p_print_int
64		BL p_print_ln
65		ADD sp, sp, #4
66		LDR r0, =0
67		POP {pc}
68		.ltorg
69	p_throw_overflow_error:
70		LDR r0, =msg_0
71		BL p_throw_runtime_error
72	p_print_int:
73		PUSH {lr}
74		MOV r1, r0
75		LDR r0, =msg_1
76		ADD r0, r0, #4
77		BL printf
78		MOV r0, #0
79		BL fflush
80		POP {pc}
81	p_print_ln:
82		PUSH {lr}
83		LDR r0, =msg_2
84		ADD r0, r0, #4
85		BL puts
86		MOV r0, #0
87		BL fflush
88		POP {pc}
89	p_throw_runtime_error:
90		BL p_print_string
91		MOV r0, #-1
92		BL exit
93	p_print_string:
94		PUSH {lr}
95		LDR r1, [r0]
96		ADD r2, r0, #4
97		LDR r0, =msg_3
98		ADD r0, r0, #4
99		BL printf
100		MOV r0, #0
101		BL fflush
102		POP {pc}
103	
===========================================================
-- Finished
