-- Compiling...
-- Printing Assembly...
intWayOverflow.s contents are:
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
21		LDR r4, =2000000000
22		STR r4, [sp]
23		LDR r4, [sp]
24		MOV r0, r4
25		BL p_print_int
26		BL p_print_ln
27		LDR r4, [sp]
28		LDR r5, =2000000000
29		ADDS r4, r4, r5
30		BLVS p_throw_overflow_error
31		STR r4, [sp]
32		LDR r4, [sp]
33		MOV r0, r4
34		BL p_print_int
35		BL p_print_ln
36		ADD sp, sp, #4
37		LDR r0, =0
38		POP {pc}
39		.ltorg
40	p_print_int:
41		PUSH {lr}
42		MOV r1, r0
43		LDR r0, =msg_0
44		ADD r0, r0, #4
45		BL printf
46		MOV r0, #0
47		BL fflush
48		POP {pc}
49	p_print_ln:
50		PUSH {lr}
51		LDR r0, =msg_1
52		ADD r0, r0, #4
53		BL puts
54		MOV r0, #0
55		BL fflush
56		POP {pc}
57	p_throw_overflow_error:
58		LDR r0, =msg_2
59		BL p_throw_runtime_error
60	p_throw_runtime_error:
61		BL p_print_string
62		MOV r0, #-1
63		BL exit
64	p_print_string:
65		PUSH {lr}
66		LDR r1, [r0]
67		ADD r2, r0, #4
68		LDR r0, =msg_3
69		ADD r0, r0, #4
70		BL printf
71		MOV r0, #0
72		BL fflush
73		POP {pc}
74	
===========================================================
-- Finished
