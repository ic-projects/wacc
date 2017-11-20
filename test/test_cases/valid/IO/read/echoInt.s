-- Compiling...
-- Printing Assembly...
echoInt.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 24
4		.ascii	"enter an integer to echo"
5	msg_1:
6		.word 5
7		.ascii	"%.*s\0"
8	msg_2:
9		.word 1
10		.ascii	"\0"
11	msg_3:
12		.word 3
13		.ascii	"%d\0"
14	msg_4:
15		.word 3
16		.ascii	"%d\0"
17	
18	.text
19	
20	.global main
21	main:
22		PUSH {lr}
23		SUB sp, sp, #4
24		LDR r4, =1
25		STR r4, [sp]
26		LDR r4, =msg_0
27		MOV r0, r4
28		BL p_print_string
29		BL p_print_ln
30		ADD r4, sp, #0
31		MOV r0, r4
32		BL p_read_int
33		LDR r4, [sp]
34		MOV r0, r4
35		BL p_print_int
36		BL p_print_ln
37		ADD sp, sp, #4
38		LDR r0, =0
39		POP {pc}
40		.ltorg
41	p_print_string:
42		PUSH {lr}
43		LDR r1, [r0]
44		ADD r2, r0, #4
45		LDR r0, =msg_1
46		ADD r0, r0, #4
47		BL printf
48		MOV r0, #0
49		BL fflush
50		POP {pc}
51	p_print_ln:
52		PUSH {lr}
53		LDR r0, =msg_2
54		ADD r0, r0, #4
55		BL puts
56		MOV r0, #0
57		BL fflush
58		POP {pc}
59	p_read_int:
60		PUSH {lr}
61		MOV r1, r0
62		LDR r0, =msg_3
63		ADD r0, r0, #4
64		BL scanf
65		POP {pc}
66	p_print_int:
67		PUSH {lr}
68		MOV r1, r0
69		LDR r0, =msg_4
70		ADD r0, r0, #4
71		BL printf
72		MOV r0, #0
73		BL fflush
74		POP {pc}
75	
===========================================================
-- Finished
