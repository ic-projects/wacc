-- Compiling...
-- Printing Assembly...
elseexit.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 82
4		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
5	msg_1:
6		.word 5
7		.ascii	"%.*s\0"
8	
9	.text
10	
11	.global main
12	main:
13		PUSH {lr}
14		LDR r4, =2
15		LDR r5, =2
16		ADDS r4, r4, r5
17		BLVS p_throw_overflow_error
18		LDR r5, =3
19		CMP r4, r5
20		MOVEQ r4, #1
21		MOVNE r4, #0
22		CMP r4, #0
23		BEQ L0
24		LDR r4, =0
25		MOV r0, r4
26		BL exit
27		B L1
28	L0:
29		LDR r4, =1
30		MOV r0, r4
31		BL exit
32	L1:
33		LDR r0, =0
34		POP {pc}
35		.ltorg
36	p_throw_overflow_error:
37		LDR r0, =msg_0
38		BL p_throw_runtime_error
39	p_throw_runtime_error:
40		BL p_print_string
41		MOV r0, #-1
42		BL exit
43	p_print_string:
44		PUSH {lr}
45		LDR r1, [r0]
46		ADD r2, r0, #4
47		LDR r0, =msg_1
48		ADD r0, r0, #4
49		BL printf
50		MOV r0, #0
51		BL fflush
52		POP {pc}
53	
===========================================================
-- Finished
