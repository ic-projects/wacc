-- Compiling...
-- Printing Assembly...
printInt.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 22
4		.ascii	"An example integer is "
5	msg_1:
6		.word 5
7		.ascii	"%.*s\0"
8	msg_2:
9		.word 3
10		.ascii	"%d\0"
11	msg_3:
12		.word 1
13		.ascii	"\0"
14	
15	.text
16	
17	.global main
18	main:
19		PUSH {lr}
20		LDR r4, =msg_0
21		MOV r0, r4
22		BL p_print_string
23		LDR r4, =189
24		MOV r0, r4
25		BL p_print_int
26		BL p_print_ln
27		LDR r0, =0
28		POP {pc}
29		.ltorg
30	p_print_string:
31		PUSH {lr}
32		LDR r1, [r0]
33		ADD r2, r0, #4
34		LDR r0, =msg_1
35		ADD r0, r0, #4
36		BL printf
37		MOV r0, #0
38		BL fflush
39		POP {pc}
40	p_print_int:
41		PUSH {lr}
42		MOV r1, r0
43		LDR r0, =msg_2
44		ADD r0, r0, #4
45		BL printf
46		MOV r0, #0
47		BL fflush
48		POP {pc}
49	p_print_ln:
50		PUSH {lr}
51		LDR r0, =msg_3
52		ADD r0, r0, #4
53		BL puts
54		MOV r0, #0
55		BL fflush
56		POP {pc}
57	
===========================================================
-- Finished
