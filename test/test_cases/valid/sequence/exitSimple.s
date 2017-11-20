-- Compiling...
-- Printing Assembly...
exitSimple.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 22
4		.ascii	"Should not print this."
5	msg_1:
6		.word 5
7		.ascii	"%.*s\0"
8	msg_2:
9		.word 1
10		.ascii	"\0"
11	
12	.text
13	
14	.global main
15	main:
16		PUSH {lr}
17		LDR r4, =42
18		MOV r0, r4
19		BL exit
20		LDR r4, =msg_0
21		MOV r0, r4
22		BL p_print_string
23		BL p_print_ln
24		LDR r0, =0
25		POP {pc}
26		.ltorg
27	p_print_string:
28		PUSH {lr}
29		LDR r1, [r0]
30		ADD r2, r0, #4
31		LDR r0, =msg_1
32		ADD r0, r0, #4
33		BL printf
34		MOV r0, #0
35		BL fflush
36		POP {pc}
37	p_print_ln:
38		PUSH {lr}
39		LDR r0, =msg_2
40		ADD r0, r0, #4
41		BL puts
42		MOV r0, #0
43		BL fflush
44		POP {pc}
45	
===========================================================
-- Finished
