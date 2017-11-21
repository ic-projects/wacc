-- Compiling...
-- Printing Assembly...
intinlineMultOverflow.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 3
4		.ascii	"%d\0"
5	msg_1:
6		.word 1
7		.ascii	"\0"
8	msg_2:
9		.word 82
10		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
11	msg_3:
12		.word 5
13		.ascii	"%.*s\0"
14	
15	.text
16	
17	.global main
18	main:
19		PUSH {lr}
20		LDR r4, =2147483
21		MOV r0, r4
22		BL p_print_int
23		BL p_print_ln
24		LDR r4, =2147483
25		LDR r5, =1000
26		SMULL r4, r5, r4, r5
27		CMP r5, r4, ASR #31
28		BLNE p_throw_overflow_error
29		MOV r0, r4
30		BL p_print_int
31		BL p_print_ln
32		LDR r4, =2147483
33		LDR r5, =1000
34		SMULL r4, r5, r4, r5
35		CMP r5, r4, ASR #31
36		BLNE p_throw_overflow_error
37		LDR r5, =1000
38		SMULL r4, r5, r4, r5
39		CMP r5, r4, ASR #31
40		BLNE p_throw_overflow_error
41		MOV r0, r4
42		BL p_print_int
43		BL p_print_ln
44		LDR r4, =2147483
45		LDR r5, =1000
46		SMULL r4, r5, r4, r5
47		CMP r5, r4, ASR #31
48		BLNE p_throw_overflow_error
49		LDR r5, =1000
50		SMULL r4, r5, r4, r5
51		CMP r5, r4, ASR #31
52		BLNE p_throw_overflow_error
53		LDR r5, =1000
54		SMULL r4, r5, r4, r5
55		CMP r5, r4, ASR #31
56		BLNE p_throw_overflow_error
57		MOV r0, r4
58		BL p_print_int
59		BL p_print_ln
60		LDR r0, =0
61		POP {pc}
62		.ltorg
63	p_print_int:
64		PUSH {lr}
65		MOV r1, r0
66		LDR r0, =msg_0
67		ADD r0, r0, #4
68		BL printf
69		MOV r0, #0
70		BL fflush
71		POP {pc}
72	p_print_ln:
73		PUSH {lr}
74		LDR r0, =msg_1
75		ADD r0, r0, #4
76		BL puts
77		MOV r0, #0
78		BL fflush
79		POP {pc}
80	p_throw_overflow_error:
81		LDR r0, =msg_2
82		BL p_throw_runtime_error
83	p_throw_runtime_error:
84		BL p_print_string
85		MOV r0, #-1
86		BL exit
87	p_print_string:
88		PUSH {lr}
89		LDR r1, [r0]
90		ADD r2, r0, #4
91		LDR r0, =msg_3
92		ADD r0, r0, #4
93		BL printf
94		MOV r0, #0
95		BL fflush
96		POP {pc}
97	
===========================================================
-- Finished
