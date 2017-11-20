-- Compiling...
-- Printing Assembly...
plusExpr.s contents are:
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
18	main:
19		PUSH {lr}
20		SUB sp, sp, #8
21		LDR r4, =15
22		STR r4, [sp, #4]
23		LDR r4, =20
24		STR r4, [sp]
25		LDR r4, [sp, #4]
26		LDR r5, [sp]
27		ADDS r4, r4, r5
28		BLVS p_throw_overflow_error
29		MOV r0, r4
30		BL p_print_int
31		BL p_print_ln
32		ADD sp, sp, #8
33		LDR r0, =0
34		POP {pc}
35		.ltorg
36	p_throw_overflow_error:
37		LDR r0, =msg_0
38		BL p_throw_runtime_error
39	p_print_int:
40		PUSH {lr}
41		MOV r1, r0
42		LDR r0, =msg_1
43		ADD r0, r0, #4
44		BL printf
45		MOV r0, #0
46		BL fflush
47		POP {pc}
48	p_print_ln:
49		PUSH {lr}
50		LDR r0, =msg_2
51		ADD r0, r0, #4
52		BL puts
53		MOV r0, #0
54		BL fflush
55		POP {pc}
56	p_throw_runtime_error:
57		BL p_print_string
58		MOV r0, #-1
59		BL exit
60	p_print_string:
61		PUSH {lr}
62		LDR r1, [r0]
63		ADD r2, r0, #4
64		LDR r0, =msg_3
65		ADD r0, r0, #4
66		BL printf
67		MOV r0, #0
68		BL fflush
69		POP {pc}
70	
===========================================================
-- Finished
