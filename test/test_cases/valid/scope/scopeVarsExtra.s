-- Compiling...
-- Printing Assembly...
scopeVarsExtra.s contents are:
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
20		SUB sp, sp, #4
21		LDR r4, =2
22		STR r4, [sp]
23		LDR r4, [sp]
24		MOV r0, r4
25		BL p_print_int
26		BL p_print_ln
27		SUB sp, sp, #4
28		LDR r4, [sp, #4]
29		LDR r5, =1
30		ADDS r4, r4, r5
31		BLVS p_throw_overflow_error
32		STR r4, [sp, #4]
33		LDR r4, =4
34		STR r4, [sp]
35		LDR r4, [sp]
36		MOV r0, r4
37		BL p_print_int
38		BL p_print_ln
39		LDR r4, [sp]
40		LDR r5, =1
41		ADDS r4, r4, r5
42		BLVS p_throw_overflow_error
43		STR r4, [sp]
44		LDR r4, [sp]
45		MOV r0, r4
46		BL p_print_int
47		BL p_print_ln
48		ADD sp, sp, #4
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
74	p_throw_overflow_error:
75		LDR r0, =msg_2
76		BL p_throw_runtime_error
77	p_throw_runtime_error:
78		BL p_print_string
79		MOV r0, #-1
80		BL exit
81	p_print_string:
82		PUSH {lr}
83		LDR r1, [r0]
84		ADD r2, r0, #4
85		LDR r0, =msg_3
86		ADD r0, r0, #4
87		BL printf
88		MOV r0, #0
89		BL fflush
90		POP {pc}
91	
===========================================================
-- Finished
