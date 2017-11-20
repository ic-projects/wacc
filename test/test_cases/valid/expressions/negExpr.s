-- Compiling...
-- Printing Assembly...
negExpr.s contents are:
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
20		SUB sp, sp, #4
21		LDR r4, =42
22		STR r4, [sp]
23		LDR r4, [sp]
24		RSBS r4, r4, #0
25		BLVS p_throw_overflow_error
26		MOV r0, r4
27		BL p_print_int
28		BL p_print_ln
29		ADD sp, sp, #4
30		LDR r0, =0
31		POP {pc}
32		.ltorg
33	p_throw_overflow_error:
34		LDR r0, =msg_0
35		BL p_throw_runtime_error
36	p_print_int:
37		PUSH {lr}
38		MOV r1, r0
39		LDR r0, =msg_1
40		ADD r0, r0, #4
41		BL printf
42		MOV r0, #0
43		BL fflush
44		POP {pc}
45	p_print_ln:
46		PUSH {lr}
47		LDR r0, =msg_2
48		ADD r0, r0, #4
49		BL puts
50		MOV r0, #0
51		BL fflush
52		POP {pc}
53	p_throw_runtime_error:
54		BL p_print_string
55		MOV r0, #-1
56		BL exit
57	p_print_string:
58		PUSH {lr}
59		LDR r1, [r0]
60		ADD r2, r0, #4
61		LDR r0, =msg_3
62		ADD r0, r0, #4
63		BL printf
64		MOV r0, #0
65		BL fflush
66		POP {pc}
67	
===========================================================
-- Finished
