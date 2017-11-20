-- Compiling...
-- Printing Assembly...
whileBoolFlip.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 7
4		.ascii	"flip b!"
5	msg_1:
6		.word 11
7		.ascii	"end of loop"
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	msg_3:
12		.word 1
13		.ascii	"\0"
14	
15	.text
16	
17	.global main
18	main:
19		PUSH {lr}
20		SUB sp, sp, #1
21		MOV r4, #1
22		STRB r4, [sp]
23		B L0
24	L1:
25		LDR r4, =msg_0
26		MOV r0, r4
27		BL p_print_string
28		BL p_print_ln
29		LDRSB r4, [sp]
30		EOR r4, r4, #1
31		STRB r4, [sp]
32	L0:
33		LDRSB r4, [sp]
34		CMP r4, #1
35		BEQ L1
36		LDR r4, =msg_1
37		MOV r0, r4
38		BL p_print_string
39		BL p_print_ln
40		ADD sp, sp, #1
41		LDR r0, =0
42		POP {pc}
43		.ltorg
44	p_print_string:
45		PUSH {lr}
46		LDR r1, [r0]
47		ADD r2, r0, #4
48		LDR r0, =msg_2
49		ADD r0, r0, #4
50		BL printf
51		MOV r0, #0
52		BL fflush
53		POP {pc}
54	p_print_ln:
55		PUSH {lr}
56		LDR r0, =msg_3
57		ADD r0, r0, #4
58		BL puts
59		MOV r0, #0
60		BL fflush
61		POP {pc}
62	
===========================================================
-- Finished
