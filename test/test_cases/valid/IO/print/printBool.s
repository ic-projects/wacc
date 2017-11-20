-- Compiling...
-- Printing Assembly...
printBool.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 8
4		.ascii	"True is "
5	msg_1:
6		.word 9
7		.ascii	"False is "
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 5
13		.ascii	"true\0"
14	msg_4:
15		.word 6
16		.ascii	"false\0"
17	msg_5:
18		.word 1
19		.ascii	"\0"
20	
21	.text
22	
23	.global main
24	main:
25		PUSH {lr}
26		LDR r4, =msg_0
27		MOV r0, r4
28		BL p_print_string
29		MOV r4, #1
30		MOV r0, r4
31		BL p_print_bool
32		BL p_print_ln
33		LDR r4, =msg_1
34		MOV r0, r4
35		BL p_print_string
36		MOV r4, #0
37		MOV r0, r4
38		BL p_print_bool
39		BL p_print_ln
40		LDR r0, =0
41		POP {pc}
42		.ltorg
43	p_print_string:
44		PUSH {lr}
45		LDR r1, [r0]
46		ADD r2, r0, #4
47		LDR r0, =msg_2
48		ADD r0, r0, #4
49		BL printf
50		MOV r0, #0
51		BL fflush
52		POP {pc}
53	p_print_bool:
54		PUSH {lr}
55		CMP r0, #0
56		LDRNE r0, =msg_3
57		LDREQ r0, =msg_4
58		ADD r0, r0, #4
59		BL printf
60		MOV r0, #0
61		BL fflush
62		POP {pc}
63	p_print_ln:
64		PUSH {lr}
65		LDR r0, =msg_5
66		ADD r0, r0, #4
67		BL puts
68		MOV r0, #0
69		BL fflush
70		POP {pc}
71	
===========================================================
-- Finished
