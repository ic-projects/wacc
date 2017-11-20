-- Compiling...
-- Printing Assembly...
modifyString.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 12
4		.ascii	"hello world!"
5	msg_1:
6		.word 3
7		.ascii	"Hi!"
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 1
13		.ascii	"\0"
14	msg_4:
15		.word 44
16		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
17	msg_5:
18		.word 45
19		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
20	
21	.text
22	
23	.global main
24	main:
25		PUSH {lr}
26		SUB sp, sp, #4
27		LDR r4, =msg_0
28		STR r4, [sp]
29		LDR r4, [sp]
30		MOV r0, r4
31		BL p_print_string
32		BL p_print_ln
33		MOV r4, #'H'
34		ADD r5, sp, #0
35		LDR r6, =0
36		LDR r5, [r5]
37		MOV r0, r6
38		MOV r1, r5
39		BL p_check_array_bounds
40		ADD r5, r5, #4
41		ADD r5, r5, r6
42		STRB r4, [r5]
43		LDR r4, [sp]
44		MOV r0, r4
45		BL p_print_string
46		BL p_print_ln
47		LDR r4, =msg_1
48		STR r4, [sp]
49		LDR r4, [sp]
50		MOV r0, r4
51		BL p_print_string
52		BL p_print_ln
53		ADD sp, sp, #4
54		LDR r0, =0
55		POP {pc}
56		.ltorg
57	p_print_string:
58		PUSH {lr}
59		LDR r1, [r0]
60		ADD r2, r0, #4
61		LDR r0, =msg_2
62		ADD r0, r0, #4
63		BL printf
64		MOV r0, #0
65		BL fflush
66		POP {pc}
67	p_print_ln:
68		PUSH {lr}
69		LDR r0, =msg_3
70		ADD r0, r0, #4
71		BL puts
72		MOV r0, #0
73		BL fflush
74		POP {pc}
75	p_check_array_bounds:
76		PUSH {lr}
77		CMP r0, #0
78		LDRLT r0, =msg_4
79		BLLT p_throw_runtime_error
80		LDR r1, [r1]
81		CMP r0, r1
82		LDRCS r0, =msg_5
83		BLCS p_throw_runtime_error
84		POP {pc}
85	p_throw_runtime_error:
86		BL p_print_string
87		MOV r0, #-1
88		BL exit
89	
===========================================================
-- Finished
