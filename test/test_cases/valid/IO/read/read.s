-- Compiling...
-- Printing Assembly...
read.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 31
4		.ascii	"input an integer to continue..."
5	msg_1:
6		.word 5
7		.ascii	"%.*s\0"
8	msg_2:
9		.word 1
10		.ascii	"\0"
11	msg_3:
12		.word 3
13		.ascii	"%d\0"
14	
15	.text
16	
17	.global main
18	main:
19		PUSH {lr}
20		SUB sp, sp, #4
21		LDR r4, =10
22		STR r4, [sp]
23		LDR r4, =msg_0
24		MOV r0, r4
25		BL p_print_string
26		BL p_print_ln
27		ADD r4, sp, #0
28		MOV r0, r4
29		BL p_read_int
30		ADD sp, sp, #4
31		LDR r0, =0
32		POP {pc}
33		.ltorg
34	p_print_string:
35		PUSH {lr}
36		LDR r1, [r0]
37		ADD r2, r0, #4
38		LDR r0, =msg_1
39		ADD r0, r0, #4
40		BL printf
41		MOV r0, #0
42		BL fflush
43		POP {pc}
44	p_print_ln:
45		PUSH {lr}
46		LDR r0, =msg_2
47		ADD r0, r0, #4
48		BL puts
49		MOV r0, #0
50		BL fflush
51		POP {pc}
52	p_read_int:
53		PUSH {lr}
54		MOV r1, r0
55		LDR r0, =msg_3
56		ADD r0, r0, #4
57		BL scanf
58		POP {pc}
59	
===========================================================
-- Finished
