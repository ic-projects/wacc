-- Compiling...
-- Printing Assembly...
longExpr3.s contents are:
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
12	main:
13		PUSH {lr}
14		SUB sp, sp, #4
15		LDR r4, =1
16		LDR r5, =2
17		SUBS r4, r4, r5
18		BLVS p_throw_overflow_error
19		LDR r5, =3
20		ADDS r4, r4, r5
21		BLVS p_throw_overflow_error
22		LDR r5, =4
23		SUBS r4, r4, r5
24		BLVS p_throw_overflow_error
25		LDR r5, =5
26		ADDS r4, r4, r5
27		BLVS p_throw_overflow_error
28		LDR r5, =6
29		SUBS r4, r4, r5
30		BLVS p_throw_overflow_error
31		LDR r5, =7
32		ADDS r4, r4, r5
33		BLVS p_throw_overflow_error
34		LDR r5, =8
35		SUBS r4, r4, r5
36		BLVS p_throw_overflow_error
37		LDR r5, =9
38		ADDS r4, r4, r5
39		BLVS p_throw_overflow_error
40		LDR r5, =10
41		SUBS r4, r4, r5
42		BLVS p_throw_overflow_error
43		LDR r5, =11
44		ADDS r4, r4, r5
45		BLVS p_throw_overflow_error
46		LDR r5, =12
47		SUBS r4, r4, r5
48		BLVS p_throw_overflow_error
49		LDR r5, =13
50		ADDS r4, r4, r5
51		BLVS p_throw_overflow_error
52		LDR r5, =14
53		SUBS r4, r4, r5
54		BLVS p_throw_overflow_error
55		LDR r5, =15
56		ADDS r4, r4, r5
57		BLVS p_throw_overflow_error
58		LDR r5, =16
59		SUBS r4, r4, r5
60		BLVS p_throw_overflow_error
61		LDR r5, =17
62		ADDS r4, r4, r5
63		BLVS p_throw_overflow_error
64		STR r4, [sp]
65		LDR r4, [sp]
66		MOV r0, r4
67		BL exit
68		ADD sp, sp, #4
69		LDR r0, =0
70		POP {pc}
71		.ltorg
72	p_throw_overflow_error:
73		LDR r0, =msg_0
74		BL p_throw_runtime_error
75	p_throw_runtime_error:
76		BL p_print_string
77		MOV r0, #-1
78		BL exit
79	p_print_string:
80		PUSH {lr}
81		LDR r1, [r0]
82		ADD r2, r0, #4
83		LDR r0, =msg_1
84		ADD r0, r0, #4
85		BL printf
86		MOV r0, #0
87		BL fflush
88		POP {pc}
89	
===========================================================
-- Finished
