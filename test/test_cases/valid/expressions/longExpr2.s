-- Compiling...
-- Printing Assembly...
longExpr2.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 82
4		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
5	msg_1:
6		.word 45
7		.ascii	"DivideByZeroError: divide or modulo by zero\n\0"
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	
12	.text
13	
14	.global main
15	main:
16		PUSH {lr}
17		SUB sp, sp, #4
18		LDR r4, =2
19		LDR r5, =3
20		ADDS r4, r4, r5
21		BLVS p_throw_overflow_error
22		LDR r5, =2
23		ADDS r4, r4, r5
24		BLVS p_throw_overflow_error
25		LDR r5, =1
26		ADDS r4, r4, r5
27		BLVS p_throw_overflow_error
28		LDR r5, =1
29		ADDS r4, r4, r5
30		BLVS p_throw_overflow_error
31		LDR r5, =1
32		ADDS r4, r4, r5
33		BLVS p_throw_overflow_error
34		LDR r5, =1
35		LDR r6, =2
36		ADDS r5, r5, r6
37		BLVS p_throw_overflow_error
38		LDR r6, =3
39		LDR r7, =4
40		LDR r8, =6
41		MOV r0, r7
42		MOV r1, r8
43		BL p_check_divide_by_zero
44		BL __aeabi_idiv
45		MOV r7, r0
46		SUBS r6, r6, r7
47		BLVS p_throw_overflow_error
48		SMULL r5, r6, r5, r6
49		CMP r6, r5, ASR #31
50		BLNE p_throw_overflow_error
51		LDR r6, =2
52		LDR r7, =18
53		LDR r8, =17
54		SUBS r7, r7, r8
55		BLVS p_throw_overflow_error
56		SMULL r6, r7, r6, r7
57		CMP r7, r6, ASR #31
58		BLNE p_throw_overflow_error
59		LDR r7, =3
60		LDR r8, =4
61		SMULL r7, r8, r7, r8
62		CMP r8, r7, ASR #31
63		BLNE p_throw_overflow_error
64		LDR r8, =4
65		MOV r0, r7
66		MOV r1, r8
67		BL p_check_divide_by_zero
68		BL __aeabi_idiv
69		MOV r7, r0
70		LDR r8, =6
71		ADDS r7, r7, r8
72		BLVS p_throw_overflow_error
73		ADDS r6, r6, r7
74		BLVS p_throw_overflow_error
75		MOV r0, r5
76		MOV r1, r6
77		BL p_check_divide_by_zero
78		BL __aeabi_idiv
79		MOV r5, r0
80		SUBS r4, r4, r5
81		BLVS p_throw_overflow_error
82		STR r4, [sp]
83		LDR r4, [sp]
84		MOV r0, r4
85		BL exit
86		ADD sp, sp, #4
87		LDR r0, =0
88		POP {pc}
89		.ltorg
90	p_throw_overflow_error:
91		LDR r0, =msg_0
92		BL p_throw_runtime_error
93	p_check_divide_by_zero:
94		PUSH {lr}
95		CMP r1, #0
96		LDREQ r0, =msg_1
97		BLEQ p_throw_runtime_error
98		POP {pc}
99	p_throw_runtime_error:
100		BL p_print_string
101		MOV r0, #-1
102		BL exit
103	p_print_string:
104		PUSH {lr}
105		LDR r1, [r0]
106		ADD r2, r0, #4
107		LDR r0, =msg_2
108		ADD r0, r0, #4
109		BL printf
110		MOV r0, #0
111		BL fflush
112		POP {pc}
113	
===========================================================
-- Finished
