-- Compiling...
-- Printing Assembly...
negBothDiv.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 45
4		.ascii	"DivideByZeroError: divide or modulo by zero\n\0"
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
21		LDR r4, =-4
22		STR r4, [sp, #4]
23		LDR r4, =-2
24		STR r4, [sp]
25		LDR r4, [sp, #4]
26		LDR r5, [sp]
27		MOV r0, r4
28		MOV r1, r5
29		BL p_check_divide_by_zero
30		BL __aeabi_idiv
31		MOV r4, r0
32		MOV r0, r4
33		BL p_print_int
34		BL p_print_ln
35		ADD sp, sp, #8
36		LDR r0, =0
37		POP {pc}
38		.ltorg
39	p_check_divide_by_zero:
40		PUSH {lr}
41		CMP r1, #0
42		LDREQ r0, =msg_0
43		BLEQ p_throw_runtime_error
44		POP {pc}
45	p_print_int:
46		PUSH {lr}
47		MOV r1, r0
48		LDR r0, =msg_1
49		ADD r0, r0, #4
50		BL printf
51		MOV r0, #0
52		BL fflush
53		POP {pc}
54	p_print_ln:
55		PUSH {lr}
56		LDR r0, =msg_2
57		ADD r0, r0, #4
58		BL puts
59		MOV r0, #0
60		BL fflush
61		POP {pc}
62	p_throw_runtime_error:
63		BL p_print_string
64		MOV r0, #-1
65		BL exit
66	p_print_string:
67		PUSH {lr}
68		LDR r1, [r0]
69		ADD r2, r0, #4
70		LDR r0, =msg_3
71		ADD r0, r0, #4
72		BL printf
73		MOV r0, #0
74		BL fflush
75		POP {pc}
76	
===========================================================
-- Finished
