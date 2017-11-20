-- Compiling...
-- Printing Assembly...
println.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 12
4		.ascii	"Hello World!"
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
17		LDR r4, =msg_0
18		MOV r0, r4
19		BL p_print_string
20		BL p_print_ln
21		LDR r0, =0
22		POP {pc}
23		.ltorg
24	p_print_string:
25		PUSH {lr}
26		LDR r1, [r0]
27		ADD r2, r0, #4
28		LDR r0, =msg_1
29		ADD r0, r0, #4
30		BL printf
31		MOV r0, #0
32		BL fflush
33		POP {pc}
34	p_print_ln:
35		PUSH {lr}
36		LDR r0, =msg_2
37		ADD r0, r0, #4
38		BL puts
39		MOV r0, #0
40		BL fflush
41		POP {pc}
42	
===========================================================
-- Finished
