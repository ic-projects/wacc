-- Compiling...
-- Printing Assembly...
functionUpdateParameter.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 5
4		.ascii	"x is "
5	msg_1:
6		.word 9
7		.ascii	"x is now "
8	msg_2:
9		.word 5
10		.ascii	"y is "
11	msg_3:
12		.word 11
13		.ascii	"y is still "
14	msg_4:
15		.word 5
16		.ascii	"%.*s\0"
17	msg_5:
18		.word 3
19		.ascii	"%d\0"
20	msg_6:
21		.word 1
22		.ascii	"\0"
23	
24	.text
25	
26	.global main
27	f_f:
28		PUSH {lr}
29		LDR r4, =msg_0
30		MOV r0, r4
31		BL p_print_string
32		LDR r4, [sp, #4]
33		MOV r0, r4
34		BL p_print_int
35		BL p_print_ln
36		LDR r4, =5
37		STR r4, [sp, #4]
38		LDR r4, =msg_1
39		MOV r0, r4
40		BL p_print_string
41		LDR r4, [sp, #4]
42		MOV r0, r4
43		BL p_print_int
44		BL p_print_ln
45		LDR r4, [sp, #4]
46		MOV r0, r4
47		POP {pc}
48		POP {pc}
49		.ltorg
50	main:
51		PUSH {lr}
52		SUB sp, sp, #8
53		LDR r4, =1
54		STR r4, [sp, #4]
55		LDR r4, =msg_2
56		MOV r0, r4
57		BL p_print_string
58		LDR r4, [sp, #4]
59		MOV r0, r4
60		BL p_print_int
61		BL p_print_ln
62		LDR r4, [sp, #4]
63		STR r4, [sp, #-4]!
64		BL f_f
65		ADD sp, sp, #4
66		MOV r4, r0
67		STR r4, [sp]
68		LDR r4, =msg_3
69		MOV r0, r4
70		BL p_print_string
71		LDR r4, [sp, #4]
72		MOV r0, r4
73		BL p_print_int
74		BL p_print_ln
75		ADD sp, sp, #8
76		LDR r0, =0
77		POP {pc}
78		.ltorg
79	p_print_string:
80		PUSH {lr}
81		LDR r1, [r0]
82		ADD r2, r0, #4
83		LDR r0, =msg_4
84		ADD r0, r0, #4
85		BL printf
86		MOV r0, #0
87		BL fflush
88		POP {pc}
89	p_print_int:
90		PUSH {lr}
91		MOV r1, r0
92		LDR r0, =msg_5
93		ADD r0, r0, #4
94		BL printf
95		MOV r0, #0
96		BL fflush
97		POP {pc}
98	p_print_ln:
99		PUSH {lr}
100		LDR r0, =msg_6
101		ADD r0, r0, #4
102		BL puts
103		MOV r0, #0
104		BL fflush
105		POP {pc}
106	
===========================================================
-- Finished
