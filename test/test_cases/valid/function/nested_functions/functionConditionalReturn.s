-- Compiling...
-- Printing Assembly...
functionConditionalReturn.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 5
4		.ascii	"true\0"
5	msg_1:
6		.word 6
7		.ascii	"false\0"
8	msg_2:
9		.word 1
10		.ascii	"\0"
11	
12	.text
13	
14	.global main
15	f_f:
16		PUSH {lr}
17		MOV r4, #1
18		CMP r4, #0
19		BEQ L0
20		MOV r4, #1
21		MOV r0, r4
22		POP {pc}
23		B L1
24	L0:
25		MOV r4, #0
26		MOV r0, r4
27		POP {pc}
28	L1:
29		POP {pc}
30		.ltorg
31	main:
32		PUSH {lr}
33		SUB sp, sp, #1
34		BL f_f
35		MOV r4, r0
36		STRB r4, [sp]
37		LDRSB r4, [sp]
38		MOV r0, r4
39		BL p_print_bool
40		BL p_print_ln
41		ADD sp, sp, #1
42		LDR r0, =0
43		POP {pc}
44		.ltorg
45	p_print_bool:
46		PUSH {lr}
47		CMP r0, #0
48		LDRNE r0, =msg_0
49		LDREQ r0, =msg_1
50		ADD r0, r0, #4
51		BL printf
52		MOV r0, #0
53		BL fflush
54		POP {pc}
55	p_print_ln:
56		PUSH {lr}
57		LDR r0, =msg_2
58		ADD r0, r0, #4
59		BL puts
60		MOV r0, #0
61		BL fflush
62		POP {pc}
63	
===========================================================
-- Finished
